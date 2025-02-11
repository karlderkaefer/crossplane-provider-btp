package wrappedserviceinstance

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	"github.com/crossplane/crossplane-runtime/pkg/resource"

	"github.com/sap/crossplane-provider-btp/apis/account/v1alpha1"
	apisv1alpha1 "github.com/sap/crossplane-provider-btp/apis/v1alpha1"
)

const (
	errNotWrappedServiceInstance = "managed resource is not a WrappedServiceInstance custom resource"
	errTrackPCUsage              = "cannot track ProviderConfig usage"
	errGetPC                     = "cannot get ProviderConfig"
	errGetCreds                  = "cannot get credentials"

	errCreate = "cannot create external resource"

	errNewClient = "cannot create new Service"
)

// AsyncServiceInstanceHandler defines async CRUD operations on a ServiceInstance
type AsyncServiceInstanceHandler interface {
	CreateResource() error
}

// A connector is expected to produce an ExternalClient when its Connect method
// is called.
type connector struct {
	kube         client.Client
	usage        resource.Tracker
	newServiceFn func(creds []byte) (interface{}, error)
}

// Connect typically produces an ExternalClient by:
// 1. Tracking that the managed resource is using a ProviderConfig.
// 2. Getting the managed resource's ProviderConfig.
// 3. Getting the credentials specified by the ProviderConfig.
// 4. Using the credentials to form a client.
func (c *connector) Connect(ctx context.Context, mg resource.Managed) (managed.ExternalClient, error) {
	cr, ok := mg.(*v1alpha1.WrappedServiceInstance)
	if !ok {
		return nil, errors.New(errNotWrappedServiceInstance)
	}

	if err := c.usage.Track(ctx, mg); err != nil {
		return nil, errors.Wrap(err, errTrackPCUsage)
	}

	pc := &apisv1alpha1.ProviderConfig{}
	if err := c.kube.Get(ctx, types.NamespacedName{Name: cr.GetProviderConfigReference().Name}, pc); err != nil {
		return nil, errors.Wrap(err, errGetPC)
	}

	return &external{}, nil
}

// An ExternalClient observes, then either creates, updates, or deletes an
// external resource to ensure it reflects the managed resource's desired state.
type external struct {
	instanceHandler AsyncServiceInstanceHandler
}

func (c *external) Observe(ctx context.Context, mg resource.Managed) (managed.ExternalObservation, error) {
	_, ok := mg.(*v1alpha1.WrappedServiceInstance)
	if !ok {
		return managed.ExternalObservation{}, errors.New(errNotWrappedServiceInstance)
	}

	return managed.ExternalObservation{
		ResourceExists: false,
	}, nil
}

func (c *external) Create(ctx context.Context, mg resource.Managed) (managed.ExternalCreation, error) {
	_, ok := mg.(*v1alpha1.WrappedServiceInstance)
	if !ok {
		return managed.ExternalCreation{}, errors.New(errNotWrappedServiceInstance)
	}
	//TODO: apply some concepts for asnyc handling (annotations etc., see upjet reconciler)
	err := c.instanceHandler.CreateResource()

	return managed.ExternalCreation{
		ConnectionDetails: managed.ConnectionDetails{},
	}, errors.Wrap(err, errCreate)
}

func (c *external) Update(ctx context.Context, mg resource.Managed) (managed.ExternalUpdate, error) {
	cr, ok := mg.(*v1alpha1.WrappedServiceInstance)
	if !ok {
		return managed.ExternalUpdate{}, errors.New(errNotWrappedServiceInstance)
	}

	fmt.Printf("Updating: %+v", cr)

	return managed.ExternalUpdate{
		// Optionally return any details that may be required to connect to the
		// external resource. These will be stored as the connection secret.
		ConnectionDetails: managed.ConnectionDetails{},
	}, nil
}

func (c *external) Delete(ctx context.Context, mg resource.Managed) error {
	cr, ok := mg.(*v1alpha1.WrappedServiceInstance)
	if !ok {
		return errors.New(errNotWrappedServiceInstance)
	}

	fmt.Printf("Deleting: %+v", cr)

	return nil
}
