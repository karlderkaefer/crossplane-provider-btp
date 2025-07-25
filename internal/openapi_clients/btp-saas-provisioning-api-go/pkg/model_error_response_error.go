/*
SaaS Provisioning Service

The SAP SaaS Provisioning service provides REST APIs that are responsible for the registration and provisioning of multitenant applications and services.   Use the APIs in this service to perform various operations related to your multitenant applications and services. For example, to get application registration details, subscribe a tenant to your application, unsubscribe a tenant from your application, retrieve all your application subscriptions, update subscription dependencies, and to get subscription job information. Note: \"Application Operations for App Providers\" APIs are intended for maintenance activities, not for runtime flows.  See also: * [Authorization](https://help.sap.com/viewer/65de2977205c403bbc107264b8eccf4b/latest/en-US/3670474a58c24ac2b082e76cbbd9dc19.html) * [Rate Limiting](https://help.sap.com/viewer/65de2977205c403bbc107264b8eccf4b/latest/en-US/77b217b3f57a45b987eb7fbc3305ce1e.html) * [Error Response Format](https://help.sap.com/viewer/65de2977205c403bbc107264b8eccf4b/latest/en-US/77fef2fb104b4b1795e2e6cee790e8b8.html) * [Asynchronous Jobs](https://help.sap.com/viewer/65de2977205c403bbc107264b8eccf4b/latest/en-US/0a0a6ab0ad114d72a6611c1c6b21683e.html)

API version: 1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ErrorResponseError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ErrorResponseError{}

// ErrorResponseError struct for ErrorResponseError
type ErrorResponseError struct {
	// Technical code of the error as a reference for support
	Code int32 `json:"code"`
	// Log correlation ID to track the event
	CorrelationID string `json:"correlationID"`
	Description *string `json:"description,omitempty"`
	// Nesting of error responses
	Details []NestingErrorDetailsResponseObject `json:"details,omitempty"`
	// User-friendly description of the error.
	Message string `json:"message"`
	// Describes a data element (for example, a resource path: /online-store/v1/products/123)
	Target *string `json:"target,omitempty"`
}

type _ErrorResponseError ErrorResponseError

// NewErrorResponseError instantiates a new ErrorResponseError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorResponseError(code int32, correlationID string, message string) *ErrorResponseError {
	this := ErrorResponseError{}
	this.Code = code
	this.CorrelationID = correlationID
	this.Message = message
	return &this
}

// NewErrorResponseErrorWithDefaults instantiates a new ErrorResponseError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorResponseErrorWithDefaults() *ErrorResponseError {
	this := ErrorResponseError{}
	return &this
}

// GetCode returns the Code field value
func (o *ErrorResponseError) GetCode() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Code
}

// GetCodeOk returns a tuple with the Code field value
// and a boolean to check if the value has been set.
func (o *ErrorResponseError) GetCodeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Code, true
}

// SetCode sets field value
func (o *ErrorResponseError) SetCode(v int32) {
	o.Code = v
}

// GetCorrelationID returns the CorrelationID field value
func (o *ErrorResponseError) GetCorrelationID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CorrelationID
}

// GetCorrelationIDOk returns a tuple with the CorrelationID field value
// and a boolean to check if the value has been set.
func (o *ErrorResponseError) GetCorrelationIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CorrelationID, true
}

// SetCorrelationID sets field value
func (o *ErrorResponseError) SetCorrelationID(v string) {
	o.CorrelationID = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ErrorResponseError) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorResponseError) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ErrorResponseError) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ErrorResponseError) SetDescription(v string) {
	o.Description = &v
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *ErrorResponseError) GetDetails() []NestingErrorDetailsResponseObject {
	if o == nil || IsNil(o.Details) {
		var ret []NestingErrorDetailsResponseObject
		return ret
	}
	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorResponseError) GetDetailsOk() ([]NestingErrorDetailsResponseObject, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *ErrorResponseError) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given []NestingErrorDetailsResponseObject and assigns it to the Details field.
func (o *ErrorResponseError) SetDetails(v []NestingErrorDetailsResponseObject) {
	o.Details = v
}

// GetMessage returns the Message field value
func (o *ErrorResponseError) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *ErrorResponseError) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *ErrorResponseError) SetMessage(v string) {
	o.Message = v
}

// GetTarget returns the Target field value if set, zero value otherwise.
func (o *ErrorResponseError) GetTarget() string {
	if o == nil || IsNil(o.Target) {
		var ret string
		return ret
	}
	return *o.Target
}

// GetTargetOk returns a tuple with the Target field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorResponseError) GetTargetOk() (*string, bool) {
	if o == nil || IsNil(o.Target) {
		return nil, false
	}
	return o.Target, true
}

// HasTarget returns a boolean if a field has been set.
func (o *ErrorResponseError) HasTarget() bool {
	if o != nil && !IsNil(o.Target) {
		return true
	}

	return false
}

// SetTarget gets a reference to the given string and assigns it to the Target field.
func (o *ErrorResponseError) SetTarget(v string) {
	o.Target = &v
}

func (o ErrorResponseError) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ErrorResponseError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["code"] = o.Code
	toSerialize["correlationID"] = o.CorrelationID
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	toSerialize["message"] = o.Message
	if !IsNil(o.Target) {
		toSerialize["target"] = o.Target
	}
	return toSerialize, nil
}

func (o *ErrorResponseError) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"code",
		"correlationID",
		"message",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varErrorResponseError := _ErrorResponseError{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varErrorResponseError)

	if err != nil {
		return err
	}

	*o = ErrorResponseError(varErrorResponseError)

	return err
}

type NullableErrorResponseError struct {
	value *ErrorResponseError
	isSet bool
}

func (v NullableErrorResponseError) Get() *ErrorResponseError {
	return v.value
}

func (v *NullableErrorResponseError) Set(val *ErrorResponseError) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorResponseError) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorResponseError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorResponseError(val *ErrorResponseError) *NullableErrorResponseError {
	return &NullableErrorResponseError{value: val, isSet: true}
}

func (v NullableErrorResponseError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorResponseError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


