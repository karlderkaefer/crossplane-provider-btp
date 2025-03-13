package wrappedserviceinstance

import (
	"context"
	"fmt"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	tjcontroller "github.com/crossplane/upjet/pkg/controller"
	ujresource "github.com/crossplane/upjet/pkg/resource"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/crossplane/crossplane-runtime/pkg/meta"
	"github.com/crossplane/crossplane-runtime/pkg/reconciler/managed"
	"github.com/crossplane/crossplane-runtime/pkg/resource"

	"github.com/sap/crossplane-provider-btp/apis/account/v1alpha1"
	apisv1alpha1 "github.com/sap/crossplane-provider-btp/apis/v1alpha1"
	"github.com/sap/crossplane-provider-btp/internal"
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

type connector struct {
	kube         client.Client
	usage        resource.Tracker
	newServiceFn func(creds []byte) (interface{}, error)

	tfConnector *tjcontroller.Connector
}

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

	return &external{
		tfConnector: c.tfConnector,
		kube:        c.kube,
	}, nil
}

// An ExternalClient observes, then either creates, updates, or deletes an
// external resource to ensure it reflects the managed resource's desired state.
type external struct {
	kube            client.Client
	instanceHandler AsyncServiceInstanceHandler
	tfConnector     *tjcontroller.Connector
}

func (c *external) Observe(ctx context.Context, mg resource.Managed) (managed.ExternalObservation, error) {
	cr, ok := mg.(*v1alpha1.WrappedServiceInstance)
	if !ok {
		return managed.ExternalObservation{}, errors.New(errNotWrappedServiceInstance)
	}

	serviceInstance := serviceInstanceCr(cr)

	connect, err := c.tfConnector.Connect(ctx, serviceInstance)
	if err != nil {
		return managed.ExternalObservation{}, err
	}

	// will return true, true, in case of in memory running async operations
	obs, err := connect.Observe(ctx, serviceInstance)
	if err != nil {
		return managed.ExternalObservation{}, err
	}

	if !obs.ResourceExists {
		return managed.ExternalObservation{ResourceExists: false}, nil
	}

	// pull data back from tf state of embedded resource
	if serviceInstance.GetCondition(ujresource.TypeAsyncOperation).Reason == ujresource.ReasonFinished {
		cr.SetConditions(xpv1.Available(), ujresource.AsyncOperationFinishedCondition())
		cr.Status.AtProvider.ObservableField = internal.Val(serviceInstance.Status.AtProvider.ID)
		c.kube.Status().Update(ctx, cr)
		meta.SetExternalName(cr, meta.GetExternalName(serviceInstance))
		c.kube.Update(ctx, cr)
	}

	return managed.ExternalObservation{
		ResourceExists:   true,
		ResourceUpToDate: true,
	}, nil

}

func (c *external) Create(ctx context.Context, mg resource.Managed) (managed.ExternalCreation, error) {
	cr, ok := mg.(*v1alpha1.WrappedServiceInstance)
	if !ok {
		return managed.ExternalCreation{}, errors.New(errNotWrappedServiceInstance)
	}

	serviceInstance := serviceInstanceCr(cr)

	connect, err := c.tfConnector.Connect(ctx, serviceInstance)
	if err != nil {
		return managed.ExternalCreation{}, err
	}

	cr.SetConditions(xpv1.Creating())
	_, err = connect.Create(ctx, serviceInstance)
	if err != nil {
		return managed.ExternalCreation{}, err
	}

	return managed.ExternalCreation{
		ConnectionDetails: managed.ConnectionDetails{},
	}, nil
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

func serviceInstanceCr(wsi *v1alpha1.WrappedServiceInstance) *v1alpha1.SubaccountServiceInstance {
	sInstance := &v1alpha1.SubaccountServiceInstance{
		TypeMeta: metav1.TypeMeta{
			Kind:       v1alpha1.SubaccountServiceInstance_Kind,
			APIVersion: v1alpha1.CRDGroupVersion.String(),
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:              wsi.Name,
			UID:               wsi.UID + "-service-instance",
			DeletionTimestamp: wsi.DeletionTimestamp,
		},
		Spec: v1alpha1.SubaccountServiceInstanceSpec{
			ResourceSpec: xpv1.ResourceSpec{
				ProviderConfigReference: &xpv1.Reference{
					Name: wsi.GetProviderConfigReference().Name,
				},
				ManagementPolicies: []xpv1.ManagementAction{xpv1.ManagementActionAll},
			},
			ForProvider: v1alpha1.SubaccountServiceInstanceParameters{
				Name:          &wsi.Name,
				ServiceplanID: wsi.Spec.ForProvider.ServiceplanID,
				SubaccountID:  wsi.Spec.ForProvider.SubaccountID,
			},
			InitProvider: v1alpha1.SubaccountServiceInstanceInitParameters{},
		},
		Status: v1alpha1.SubaccountServiceInstanceStatus{},
	}
	meta.SetExternalName(sInstance, meta.GetExternalName(wsi))
	return sInstance
}
