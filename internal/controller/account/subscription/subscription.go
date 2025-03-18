package subscription

import (
	"context"
	"encoding/json"
	"fmt"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/pkg/errors"
	"github.com/sap/crossplane-provider-btp/apis/account/v1alpha1"
	providerv1alpha1 "github.com/sap/crossplane-provider-btp/apis/v1alpha1"
	"github.com/sap/crossplane-provider-btp/btp"
	"github.com/sap/crossplane-provider-btp/internal/clients/subscription"
	"github.com/sap/crossplane-provider-btp/internal/tracking"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
)

const (
	errNotSubscription = "managed resource is not a Subscription custom resource"
	errTrackPCUsage    = "cannot track ProviderConfig usage"

	errExtractSecretKey     = "no Cloud Management Secret Found"
	errGetCredentialsSecret = "could not get secret of local cloud management"
	errCredentialsCorrupted = "secret credentials data not in the expected format"
)

// api handler creation logic based on a bytemap extracted from a secrets data
var newSubscriptionClientFn = func(ctx context.Context, cisSecretData map[string][]byte) (subscription.SubscriptionApiHandlerI, error) {
	if len(cisSecretData) == 0 {
		return nil, errors.New(errCredentialsCorrupted)
	}

	cisBinding := cisSecretData[providerv1alpha1.RawBindingKey]

	var cisCredential btp.CISCredential
	if err := json.Unmarshal(cisBinding, &cisCredential); err != nil {
		return nil, errors.Wrap(err, errCredentialsCorrupted)
	}

	apiHandler := subscription.NewSubscriptionApiHandler(ctx,
		cisCredential.Uaa.Clientid,
		cisCredential.Uaa.Clientsecret,
		fmt.Sprintf("%s/oauth/token", cisCredential.Uaa.Url),
		cisCredential.Endpoints.SaasRegistryServiceUrl,
	)

	return apiHandler, nil
}

type connector struct {
	kube         client.Client
	usage        resource.Tracker
	newServiceFn func(ctx context.Context, cisSecretData map[string][]byte) (subscription.SubscriptionApiHandlerI, error)
}

func (c *connector) Connect(ctx context.Context, mg resource.Managed) (managed.ExternalClient, error) {
	cr, ok := mg.(*v1alpha1.Subscription)
	if !ok {
		return nil, errors.New(errNotSubscription)
	}

	if err := c.usage.Track(ctx, mg); err != nil {
		return nil, errors.Wrap(err, errTrackPCUsage)
	}

	secretName := cr.Spec.CloudManagementSecret
	namespace := cr.Spec.CloudManagementSecretNamespace
	creds, errGet := c.loadSecret(ctx, secretName, namespace)
	if errGet != nil {
		return nil, errGet
	}

	svc, errInit := c.newServiceFn(ctx, creds)
	if errInit != nil {
		return nil, errInit
	}

	return &external{
		kube:       c.kube,
		apiHandler: svc,
		typeMapper: subscription.NewSubscriptionTypeMapper(),
	}, nil
}

func (c *connector) loadSecret(ctx context.Context, name string, namespace string) (map[string][]byte, error) {
	if name == "" || namespace == "" {
		return nil, errors.New(errExtractSecretKey)
	}

	secret := &corev1.Secret{}
	if err := c.kube.Get(
		ctx, types.NamespacedName{
			Namespace: namespace,
			Name:      name,
		}, secret,
	); err != nil {
		return nil, errors.Wrap(err, errGetCredentialsSecret)
	}
	return secret.Data, nil
}

type external struct {
	kube       client.Client
	apiHandler subscription.SubscriptionApiHandlerI
	typeMapper subscription.SubscriptionTypeMapperI
	tracker    tracking.ReferenceResolverTracker
}

func (c *external) Observe(ctx context.Context, mg resource.Managed) (managed.ExternalObservation, error) {
	cr, ok := mg.(*v1alpha1.Subscription)
	if !ok {
		return managed.ExternalObservation{}, errors.New(errNotSubscription)
	}

	apiRes, err := c.loadSubscription(ctx, cr)
	if err != nil {
		return managed.ExternalObservation{}, err
	}
	if apiRes == nil {
		return managed.ExternalObservation{ResourceExists: false}, nil
	}

	c.syncStatus(apiRes, cr)

	if c.typeMapper.IsAvailable(cr) {
		cr.SetConditions(xpv1.Available())
	} else {
		cr.SetConditions(xpv1.Unavailable())
	}

	return managed.ExternalObservation{
		ResourceExists:    true,
		ResourceUpToDate:  c.isUpToDate(apiRes, cr),
		ConnectionDetails: managed.ConnectionDetails{},
	}, nil
}

func (c *external) Create(ctx context.Context, mg resource.Managed) (managed.ExternalCreation, error) {
	cr, ok := mg.(*v1alpha1.Subscription)
	if !ok {
		return managed.ExternalCreation{}, errors.New(errNotSubscription)
	}

	cr.SetConditions(xpv1.Creating())
	externalName, clientErr := c.apiHandler.CreateSubscription(ctx, c.typeMapper.ConvertToCreatePayload(cr))
	if clientErr != nil {
		return managed.ExternalCreation{}, clientErr
	}

	// set external ID as name to allow proper importing
	meta.SetExternalName(cr, externalName)

	return managed.ExternalCreation{
		ConnectionDetails: managed.ConnectionDetails{},
	}, nil
}

// Update implemented, but actually not used right now
func (c *external) Update(ctx context.Context, mg resource.Managed) (managed.ExternalUpdate, error) {
	cr, ok := mg.(*v1alpha1.Subscription)
	if !ok {
		return managed.ExternalUpdate{}, errors.New(errNotSubscription)
	}

	err := c.apiHandler.UpdateSubscription(ctx, meta.GetExternalName(cr), c.typeMapper.ConvertToUpdatePayload(cr))
	if err != nil {
		return managed.ExternalUpdate{}, err
	}

	return managed.ExternalUpdate{
		// Optionally return any details that may be required to connect to the
		// external resource. These will be stored as the connection secret.
		ConnectionDetails: managed.ConnectionDetails{},
	}, nil
}

func (c *external) Delete(ctx context.Context, mg resource.Managed) (managed.ExternalDelete, error) {
	cr, ok := mg.(*v1alpha1.Subscription)
	if !ok {
		return managed.ExternalDelete{}, errors.New(errNotSubscription)
	}

	cr.SetConditions(xpv1.Deleting())

	if !c.typeMapper.IsAvailable(cr) {
		// api will return 500 if called multiple times, so we will ensure to call it only once
		return managed.ExternalDelete{}, nil
	}
	return managed.ExternalDelete{}, c.apiHandler.DeleteSubscription(ctx, meta.GetExternalName(cr))
}

func (e *external) Disconnect(ctx context.Context) error {
	return nil
}

// loadSubscription gets a Subscription using the APIHandler if a proper externalName has been set, otherwise returns nil
func (c *external) loadSubscription(ctx context.Context, cr *v1alpha1.Subscription) (*subscription.SubscriptionGet, error) {
	externalName := meta.GetExternalName(cr)
	if externalName == cr.Name {
		// in case a subscription has never been created (or imported) the externalName will be set from the resource name
		// -> resource needs creation in this case
		return nil, nil
	}
	return c.apiHandler.GetSubscription(ctx, meta.GetExternalName(cr))
}

// syncStatus delegates saving the observation based on external resource to the typemapper
func (c *external) syncStatus(apiRes *subscription.SubscriptionGet, cr *v1alpha1.Subscription) {
	c.typeMapper.SyncStatus(apiRes, &cr.Status.AtProvider)
}

// isUpToDate delegates comparision of cr data and api resource to the typemapper
func (c *external) isUpToDate(apiRes *subscription.SubscriptionGet, cr *v1alpha1.Subscription) bool {
	return c.typeMapper.IsUpToDate(cr, apiRes)
}
