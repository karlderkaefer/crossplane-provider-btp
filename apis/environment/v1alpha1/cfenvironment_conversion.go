package v1alpha1

func (*CloudFoundryEnvironment) Hub() {}

// import (
// 	"fmt"
// 	v1beta1 "github.com/sap/crossplane-provider-btp/apis/environment/v1beta1"
// 	"sigs.k8s.io/controller-runtime/pkg/conversion"
// )

// func (src *CloudFoundryEnvironment) ConvertTo(dst conversion.Hub) error {
// 	dHub, ok := dst.(*v1beta1.CloudFoundryEnvironment)
// 	if !ok {
// 		return fmt.Errorf("Unsupported conversion")
// 	}
// 	// Implement conversion logic here from v1beta1 to v1alpha1
// 	if src.Spec.ForProvider.Landscape != "" {
// 		dHub.Spec.ForProvider.Landscape = src.Spec.ForProvider.Landscape
// 	}
// 	if src.Spec.ForProvider.Name != "" {
// 		dHub.Spec.ForProvider.Name = src.Spec.ForProvider.Name
// 	}
// 	if src.Spec.ForProvider.OrgName != "" {
// 		dHub.Spec.ForProvider.Org.OrgName = src.Spec.ForProvider.OrgName
// 	}
// 	dHub.Spec.ForProvider.Managers = src.Spec.ForProvider.Managers
// 	return nil
// }

// func (dst *CloudFoundryEnvironment) ConvertFrom(src conversion.Hub) error {
// 	sHub, ok := src.(*v1beta1.CloudFoundryEnvironment)
// 	if !ok {
// 		return fmt.Errorf("Unsupported conversion")
// 	}
// 	dst.Spec.ForProvider.Landscape = sHub.Spec.ForProvider.Landscape
// 	dst.Spec.ForProvider.Name = sHub.Spec.ForProvider.Name
// 	dst.Spec.ForProvider.OrgName = sHub.Spec.ForProvider.Org.OrgName
// 	dst.Spec.ForProvider.Managers = sHub.Spec.ForProvider.Managers
// 	return nil
// }
