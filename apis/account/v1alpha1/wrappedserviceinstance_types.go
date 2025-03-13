/*
Copyright 2022 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	"reflect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

// WrappedServiceInstanceParameters are the configurable fields of a WrappedServiceInstance.
type WrappedServiceInstanceParameters struct {
	// Reference to a ServicePlan in account to populate serviceplanId.
	// +kubebuilder:validation:Optional
	ServicePlanRef *xpv1.Reference `json:"servicePlanRef,omitempty" tf:"-"`

	// Selector for a ServicePlan in account to populate serviceplanId.
	// +kubebuilder:validation:Optional
	ServicePlanSelector *xpv1.Selector `json:"servicePlanSelector,omitempty" tf:"-"`

	// (String) The ID of the service plan.
	// The ID of the service plan.
	// +crossplane:generate:reference:type=github.com/sap/crossplane-provider-btp/apis/account/v1alpha1.ServicePlan
	// +crossplane:generate:reference:extractor=github.com/sap/crossplane-provider-btp/apis/account/v1alpha1.ServicePlanId()
	// +crossplane:generate:reference:refFieldName=ServicePlanRef
	// +crossplane:generate:reference:selectorFieldName=ServicePlanSelector
	ServiceplanID *string `json:"serviceplanId,omitempty" tf:"serviceplan_id,omitempty"`

	// (String) The ID of the subaccount.
	// The ID of the subaccount.
	// +crossplane:generate:reference:type=github.com/sap/crossplane-provider-btp/apis/account/v1alpha1.Subaccount
	// +crossplane:generate:reference:extractor=github.com/sap/crossplane-provider-btp/apis/account/v1alpha1.SubaccountUuid()
	// +crossplane:generate:reference:refFieldName=SubaccountRef
	// +crossplane:generate:reference:selectorFieldName=SubaccountSelector
	SubaccountID *string `json:"subaccountId,omitempty" tf:"subaccount_id,omitempty"`

	// Reference to a Subaccount in account to populate subaccountId.
	// +kubebuilder:validation:Optional
	SubaccountRef *xpv1.Reference `json:"subaccountRef,omitempty" tf:"-"`

	// Selector for a Subaccount in account to populate subaccountId.
	// +kubebuilder:validation:Optional
	SubaccountSelector *xpv1.Selector `json:"subaccountSelector,omitempty" tf:"-"`
}

// WrappedServiceInstanceObservation are the observable fields of a WrappedServiceInstance.
type WrappedServiceInstanceObservation struct {
	ObservableField string `json:"observableField,omitempty"`
}

// A WrappedServiceInstanceSpec defines the desired state of a WrappedServiceInstance.
type WrappedServiceInstanceSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       WrappedServiceInstanceParameters `json:"forProvider"`
}

// A WrappedServiceInstanceStatus represents the observed state of a WrappedServiceInstance.
type WrappedServiceInstanceStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          WrappedServiceInstanceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// A WrappedServiceInstance is an example API type.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,btp}
type WrappedServiceInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   WrappedServiceInstanceSpec   `json:"spec"`
	Status WrappedServiceInstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// WrappedServiceInstanceList contains a list of WrappedServiceInstance
type WrappedServiceInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WrappedServiceInstance `json:"items"`
}

// WrappedServiceInstance type metadata.
var (
	WrappedServiceInstanceKind             = reflect.TypeOf(WrappedServiceInstance{}).Name()
	WrappedServiceInstanceGroupKind        = schema.GroupKind{Group: CRDGroup, Kind: WrappedServiceInstanceKind}.String()
	WrappedServiceInstanceKindAPIVersion   = WrappedServiceInstanceKind + "." + CRDGroupVersion.String()
	WrappedServiceInstanceGroupVersionKind = CRDGroupVersion.WithKind(WrappedServiceInstanceKind)
)

func init() {
	SchemeBuilder.Register(&WrappedServiceInstance{}, &WrappedServiceInstanceList{})
}
