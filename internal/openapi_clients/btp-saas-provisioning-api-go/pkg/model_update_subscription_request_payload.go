/*
SaaS Provisioning Service

The SAP SaaS Provisioning service provides REST APIs that are responsible for the registration and provisioning of multitenant applications and services.   Use the APIs in this service to perform various operations related to your multitenant applications and services. For example, to get application registration details, subscribe a tenant to your application, unsubscribe a tenant from your application, retrieve all your application subscriptions, update subscription dependencies, and to get subscription job information. Note: \"Application Operations for App Providers\" APIs are intended for maintenance activities, not for runtime flows.  See also: * [Authorization](https://help.sap.com/viewer/65de2977205c403bbc107264b8eccf4b/latest/en-US/3670474a58c24ac2b082e76cbbd9dc19.html) * [Rate Limiting](https://help.sap.com/viewer/65de2977205c403bbc107264b8eccf4b/latest/en-US/77b217b3f57a45b987eb7fbc3305ce1e.html) * [Error Response Format](https://help.sap.com/viewer/65de2977205c403bbc107264b8eccf4b/latest/en-US/77fef2fb104b4b1795e2e6cee790e8b8.html) * [Asynchronous Jobs](https://help.sap.com/viewer/65de2977205c403bbc107264b8eccf4b/latest/en-US/0a0a6ab0ad114d72a6611c1c6b21683e.html)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the UpdateSubscriptionRequestPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateSubscriptionRequestPayload{}

// UpdateSubscriptionRequestPayload Create the request to update parameters in an existing subscription from a subaccount.
type UpdateSubscriptionRequestPayload struct {
	// The new plan of the multitenant application to update in the existing subscription.
	PlanName *string `json:"planName,omitempty"`
	// Additional subscription parameters determined by the application provider.
	SubscriptionParams map[string]map[string]interface{} `json:"subscriptionParams,omitempty"`
}

// NewUpdateSubscriptionRequestPayload instantiates a new UpdateSubscriptionRequestPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateSubscriptionRequestPayload() *UpdateSubscriptionRequestPayload {
	this := UpdateSubscriptionRequestPayload{}
	return &this
}

// NewUpdateSubscriptionRequestPayloadWithDefaults instantiates a new UpdateSubscriptionRequestPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateSubscriptionRequestPayloadWithDefaults() *UpdateSubscriptionRequestPayload {
	this := UpdateSubscriptionRequestPayload{}
	return &this
}

// GetPlanName returns the PlanName field value if set, zero value otherwise.
func (o *UpdateSubscriptionRequestPayload) GetPlanName() string {
	if o == nil || IsNil(o.PlanName) {
		var ret string
		return ret
	}
	return *o.PlanName
}

// GetPlanNameOk returns a tuple with the PlanName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSubscriptionRequestPayload) GetPlanNameOk() (*string, bool) {
	if o == nil || IsNil(o.PlanName) {
		return nil, false
	}
	return o.PlanName, true
}

// HasPlanName returns a boolean if a field has been set.
func (o *UpdateSubscriptionRequestPayload) HasPlanName() bool {
	if o != nil && !IsNil(o.PlanName) {
		return true
	}

	return false
}

// SetPlanName gets a reference to the given string and assigns it to the PlanName field.
func (o *UpdateSubscriptionRequestPayload) SetPlanName(v string) {
	o.PlanName = &v
}

// GetSubscriptionParams returns the SubscriptionParams field value if set, zero value otherwise.
func (o *UpdateSubscriptionRequestPayload) GetSubscriptionParams() map[string]map[string]interface{} {
	if o == nil || IsNil(o.SubscriptionParams) {
		var ret map[string]map[string]interface{}
		return ret
	}
	return o.SubscriptionParams
}

// GetSubscriptionParamsOk returns a tuple with the SubscriptionParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateSubscriptionRequestPayload) GetSubscriptionParamsOk() (map[string]map[string]interface{}, bool) {
	if o == nil || IsNil(o.SubscriptionParams) {
		return map[string]map[string]interface{}{}, false
	}
	return o.SubscriptionParams, true
}

// HasSubscriptionParams returns a boolean if a field has been set.
func (o *UpdateSubscriptionRequestPayload) HasSubscriptionParams() bool {
	if o != nil && !IsNil(o.SubscriptionParams) {
		return true
	}

	return false
}

// SetSubscriptionParams gets a reference to the given map[string]map[string]interface{} and assigns it to the SubscriptionParams field.
func (o *UpdateSubscriptionRequestPayload) SetSubscriptionParams(v map[string]map[string]interface{}) {
	o.SubscriptionParams = v
}

func (o UpdateSubscriptionRequestPayload) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateSubscriptionRequestPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PlanName) {
		toSerialize["planName"] = o.PlanName
	}
	if !IsNil(o.SubscriptionParams) {
		toSerialize["subscriptionParams"] = o.SubscriptionParams
	}
	return toSerialize, nil
}

type NullableUpdateSubscriptionRequestPayload struct {
	value *UpdateSubscriptionRequestPayload
	isSet bool
}

func (v NullableUpdateSubscriptionRequestPayload) Get() *UpdateSubscriptionRequestPayload {
	return v.value
}

func (v *NullableUpdateSubscriptionRequestPayload) Set(val *UpdateSubscriptionRequestPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateSubscriptionRequestPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateSubscriptionRequestPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateSubscriptionRequestPayload(val *UpdateSubscriptionRequestPayload) *NullableUpdateSubscriptionRequestPayload {
	return &NullableUpdateSubscriptionRequestPayload{value: val, isSet: true}
}

func (v NullableUpdateSubscriptionRequestPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateSubscriptionRequestPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


