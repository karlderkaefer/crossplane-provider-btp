package v1beta1


// func (*CloudFoundryEnvironment) Hub() {}

import (
	"fmt"
	v1alpha1 "github.com/sap/crossplane-provider-btp/apis/environment/v1alpha1"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

func (src *CloudFoundryEnvironment) ConvertTo(dst conversion.Hub) error {
	dHub, ok := dst.(*v1alpha1.CloudFoundryEnvironment)
	if !ok {
		return fmt.Errorf("Unsupported conversion")
	}	
	// Implement conversion logic here from v1beta1 to v1alpha1
	if src.Spec.ForProvider.Landscape != "" {
		dHub.Spec.ForProvider.Landscape = src.Spec.ForProvider.Landscape
	}
	if src.Spec.ForProvider.Name != "" {
		dHub.Spec.ForProvider.Name = src.Spec.ForProvider.Name
	}
	if src.Spec.ForProvider.Org.OrgName != "" {
		dHub.Spec.ForProvider.OrgName = src.Spec.ForProvider.Org.OrgName
	}
	dHub.Spec.ForProvider.Managers = src.Spec.ForProvider.Managers
	dHub.Spec.CloudManagementRef = src.Spec.CloudManagementRef
	dHub.Spec.CloudManagementSelector = src.Spec.CloudManagementSelector
	dHub.Spec.CloudManagementSecret = src.Spec.CloudManagementSecret
	dHub.Spec.CloudManagementSecretNamespace = src.Spec.CloudManagementSecretNamespace
	dHub.Spec.SubaccountRef = src.Spec.SubaccountRef
	dHub.Spec.SubaccountSelector = src.Spec.SubaccountSelector
	dHub.Spec.SubaccountGuid = src.Spec.SubaccountGuid
	dHub.ObjectMeta = src.ObjectMeta
	
	return nil
}

func (dst *CloudFoundryEnvironment) ConvertFrom(src conversion.Hub) error {
	sHub, ok := src.(*v1alpha1.CloudFoundryEnvironment)
	if !ok {
		return fmt.Errorf("Unsupported conversion")
	}
	dst.Spec.ForProvider.Landscape = sHub.Spec.ForProvider.Landscape
	dst.Spec.ForProvider.Name = sHub.Spec.ForProvider.Name
	dst.Spec.ForProvider.Org.OrgName = sHub.Spec.ForProvider.OrgName
	dst.Spec.ForProvider.Managers = sHub.Spec.ForProvider.Managers
	dst.Spec.CloudManagementRef = sHub.Spec.CloudManagementRef
	dst.Spec.CloudManagementSelector = sHub.Spec.CloudManagementSelector
	dst.Spec.CloudManagementSecret = sHub.Spec.CloudManagementSecret
	dst.Spec.CloudManagementSecretNamespace = sHub.Spec.CloudManagementSecretNamespace
	dst.Spec.SubaccountRef = sHub.Spec.SubaccountRef
	dst.Spec.SubaccountSelector = sHub.Spec.SubaccountSelector
	dst.Spec.SubaccountGuid = sHub.Spec.SubaccountGuid
	dst.ObjectMeta = sHub.ObjectMeta
	// Implement conversion logic here from v1alpha1 to v1beta1
	return nil
}