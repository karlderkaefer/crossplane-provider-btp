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

// checks if the RegistrationDetailsResponseObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RegistrationDetailsResponseObject{}

// RegistrationDetailsResponseObject struct for RegistrationDetailsResponseObject
type RegistrationDetailsResponseObject struct {
	// The ID returned by an xsuaa service instance after the app provider has connected the multitenant application to an xsuaa service instance.
	AppId *string `json:"appId,omitempty"`
	// The unique registration name of the deployed multitenant application as defined by the app developer.
	AppName *string `json:"appName,omitempty"`
	// The plan used to register the multitenant application or reusable service. <b>- saasApplication:</b> Registered entity is a multitenant application. <b>- saasService:</b> Registered entity is a reuse service.
	AppType *string `json:"appType,omitempty"`
	// Any callback URLs that the multitenant application exposes.
	AppUrls *string `json:"appUrls,omitempty"`
	// Name of the cloud automation solution associated with the multitenant application.
	AutomationSolutionName *string `json:"automationSolutionName,omitempty"`
	// The category to which the application is grouped in the Subscriptions page in the cockpit. If left empty, it gets assigned to the default category.
	Category *string `json:"category,omitempty"`
	// The unique commercial registration name of the deployed multitenant application as defined by the app developer.
	CommercialAppName *string `json:"commercialAppName,omitempty"`
	// The description of the multitenant application for customer-facing UIs.
	Description *string `json:"description,omitempty"`
	// The display name of the application for customer-facing UIs.
	DisplayName *string `json:"displayName,omitempty"`
	// A flag to determine wheater to fail subscription when automation solution fails or not.
	FailSubscriptionOnAutomationFailure *bool `json:"failSubscriptionOnAutomationFailure,omitempty"`
	// Name of the formations solution associated with the multitenant application.
	// Deprecated
	FormationSolutionName *string `json:"formationSolutionName,omitempty"`
	// ID of the global account associated with the multitenant application.
	GlobalAccountId *string `json:"globalAccountId,omitempty"`
	// Error message to return when automation solution fails.
	MessageOnAutomationFailure *string `json:"messageOnAutomationFailure,omitempty"`
	// The unique ID of the Cloud Foundry org where the app provider has deployed and registered the multitenant application.
	OrganizationGuid *string `json:"organizationGuid,omitempty"`
	ParamsSchema *EntitledApplicationsResponseObjectParamsSchema `json:"paramsSchema,omitempty"`
	// Whether the parameters are transferred to the application’s dependencies.
	PropagateParams *bool `json:"propagateParams,omitempty"`
	// The unique ID of the tenant that provides the multitenant application.
	ProviderTenantId *string `json:"providerTenantId,omitempty"`
	// The ID of the multitenant application that is registered to the SAP SaaS Provisioning service registry.
	ServiceInstanceId *string `json:"serviceInstanceId,omitempty"`
	// The unique ID of the Cloud Foundry space where the app provider has deployed and registered the multitenant application.
	SpaceGuid *string `json:"spaceGuid,omitempty"`
	// The xsappname configured in the security descriptor file used to create the xsuaa service instance for the multitenant application.
	Xsappname *string `json:"xsappname,omitempty"`
}

// NewRegistrationDetailsResponseObject instantiates a new RegistrationDetailsResponseObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegistrationDetailsResponseObject() *RegistrationDetailsResponseObject {
	this := RegistrationDetailsResponseObject{}
	return &this
}

// NewRegistrationDetailsResponseObjectWithDefaults instantiates a new RegistrationDetailsResponseObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegistrationDetailsResponseObjectWithDefaults() *RegistrationDetailsResponseObject {
	this := RegistrationDetailsResponseObject{}
	return &this
}

// GetAppId returns the AppId field value if set, zero value otherwise.
func (o *RegistrationDetailsResponseObject) GetAppId() string {
	if o == nil || IsNil(o.AppId) {
		var ret string
		return ret
	}
	return *o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationDetailsResponseObject) GetAppIdOk() (*string, bool) {
	if o == nil || IsNil(o.AppId) {
		return nil, false
	}
	return o.AppId, true
}

// HasAppId returns a boolean if a field has been set.
func (o *RegistrationDetailsResponseObject) HasAppId() bool {
	if o != nil && !IsNil(o.AppId) {
		return true
	}

	return false
}

// SetAppId gets a reference to the given string and assigns it to the AppId field.
func (o *RegistrationDetailsResponseObject) SetAppId(v string) {
	o.AppId = &v
}

// GetAppName returns the AppName field value if set, zero value otherwise.
func (o *RegistrationDetailsResponseObject) GetAppName() string {
	if o == nil || IsNil(o.AppName) {
		var ret string
		return ret
	}
	return *o.AppName
}

// GetAppNameOk returns a tuple with the AppName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationDetailsResponseObject) GetAppNameOk() (*string, bool) {
	if o == nil || IsNil(o.AppName) {
		return nil, false
	}
	return o.AppName, true
}

// HasAppName returns a boolean if a field has been set.
func (o *RegistrationDetailsResponseObject) HasAppName() bool {
	if o != nil && !IsNil(o.AppName) {
		return true
	}

	return false
}

// SetAppName gets a reference to the given string and assigns it to the AppName field.
func (o *RegistrationDetailsResponseObject) SetAppName(v string) {
	o.AppName = &v
}

// GetAppType returns the AppType field value if set, zero value otherwise.
func (o *RegistrationDetailsResponseObject) GetAppType() string {
	if o == nil || IsNil(o.AppType) {
		var ret string
		return ret
	}
	return *o.AppType
}

// GetAppTypeOk returns a tuple with the AppType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationDetailsResponseObject) GetAppTypeOk() (*string, bool) {
	if o == nil || IsNil(o.AppType) {
		return nil, false
	}
	return o.AppType, true
}

// HasAppType returns a boolean if a field has been set.
func (o *RegistrationDetailsResponseObject) HasAppType() bool {
	if o != nil && !IsNil(o.AppType) {
		return true
	}

	return false
}

// SetAppType gets a reference to the given string and assigns it to the AppType field.
func (o *RegistrationDetailsResponseObject) SetAppType(v string) {
	o.AppType = &v
}

// GetAppUrls returns the AppUrls field value if set, zero value otherwise.
func (o *RegistrationDetailsResponseObject) GetAppUrls() string {
	if o == nil || IsNil(o.AppUrls) {
		var ret string
		return ret
	}
	return *o.AppUrls
}

// GetAppUrlsOk returns a tuple with the AppUrls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationDetailsResponseObject) GetAppUrlsOk() (*string, bool) {
	if o == nil || IsNil(o.AppUrls) {
		return nil, false
	}
	return o.AppUrls, true
}

// HasAppUrls returns a boolean if a field has been set.
func (o *RegistrationDetailsResponseObject) HasAppUrls() bool {
	if o != nil && !IsNil(o.AppUrls) {
		return true
	}

	return false
}

// SetAppUrls gets a reference to the given string and assigns it to the AppUrls field.
func (o *RegistrationDetailsResponseObject) SetAppUrls(v string) {
	o.AppUrls = &v
}

// GetAutomationSolutionName returns the AutomationSolutionName field value if set, zero value otherwise.
func (o *RegistrationDetailsResponseObject) GetAutomationSolutionName() string {
	if o == nil || IsNil(o.AutomationSolutionName) {
		var ret string
		return ret
	}
	return *o.AutomationSolutionName
}

// GetAutomationSolutionNameOk returns a tuple with the AutomationSolutionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationDetailsResponseObject) GetAutomationSolutionNameOk() (*string, bool) {
	if o == nil || IsNil(o.AutomationSolutionName) {
		return nil, false
	}
	return o.AutomationSolutionName, true
}

// HasAutomationSolutionName returns a boolean if a field has been set.
func (o *RegistrationDetailsResponseObject) HasAutomationSolutionName() bool {
	if o != nil && !IsNil(o.AutomationSolutionName) {
		return true
	}

	return false
}

// SetAutomationSolutionName gets a reference to the given string and assigns it to the AutomationSolutionName field.
func (o *RegistrationDetailsResponseObject) SetAutomationSolutionName(v string) {
	o.AutomationSolutionName = &v
}

// GetCategory returns the Category field value if set, zero value otherwise.
func (o *RegistrationDetailsResponseObject) GetCategory() string {
	if o == nil || IsNil(o.Category) {
		var ret string
		return ret
	}
	return *o.Category
}

// GetCategoryOk returns a tuple with the Category field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationDetailsResponseObject) GetCategoryOk() (*string, bool) {
	if o == nil || IsNil(o.Category) {
		return nil, false
	}
	return o.Category, true
}

// HasCategory returns a boolean if a field has been set.
func (o *RegistrationDetailsResponseObject) HasCategory() bool {
	if o != nil && !IsNil(o.Category) {
		return true
	}

	return false
}

// SetCategory gets a reference to the given string and assigns it to the Category field.
func (o *RegistrationDetailsResponseObject) SetCategory(v string) {
	o.Category = &v
}

// GetCommercialAppName returns the CommercialAppName field value if set, zero value otherwise.
func (o *RegistrationDetailsResponseObject) GetCommercialAppName() string {
	if o == nil || IsNil(o.CommercialAppName) {
		var ret string
		return ret
	}
	return *o.CommercialAppName
}

// GetCommercialAppNameOk returns a tuple with the CommercialAppName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationDetailsResponseObject) GetCommercialAppNameOk() (*string, bool) {
	if o == nil || IsNil(o.CommercialAppName) {
		return nil, false
	}
	return o.CommercialAppName, true
}

// HasCommercialAppName returns a boolean if a field has been set.
func (o *RegistrationDetailsResponseObject) HasCommercialAppName() bool {
	if o != nil && !IsNil(o.CommercialAppName) {
		return true
	}

	return false
}

// SetCommercialAppName gets a reference to the given string and assigns it to the CommercialAppName field.
func (o *RegistrationDetailsResponseObject) SetCommercialAppName(v string) {
	o.CommercialAppName = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RegistrationDetailsResponseObject) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationDetailsResponseObject) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RegistrationDetailsResponseObject) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RegistrationDetailsResponseObject) SetDescription(v string) {
	o.Description = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *RegistrationDetailsResponseObject) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationDetailsResponseObject) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// HasDisplayName returns a boolean if a field has been set.
func (o *RegistrationDetailsResponseObject) HasDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *RegistrationDetailsResponseObject) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetFailSubscriptionOnAutomationFailure returns the FailSubscriptionOnAutomationFailure field value if set, zero value otherwise.
func (o *RegistrationDetailsResponseObject) GetFailSubscriptionOnAutomationFailure() bool {
	if o == nil || IsNil(o.FailSubscriptionOnAutomationFailure) {
		var ret bool
		return ret
	}
	return *o.FailSubscriptionOnAutomationFailure
}

// GetFailSubscriptionOnAutomationFailureOk returns a tuple with the FailSubscriptionOnAutomationFailure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationDetailsResponseObject) GetFailSubscriptionOnAutomationFailureOk() (*bool, bool) {
	if o == nil || IsNil(o.FailSubscriptionOnAutomationFailure) {
		return nil, false
	}
	return o.FailSubscriptionOnAutomationFailure, true
}

// HasFailSubscriptionOnAutomationFailure returns a boolean if a field has been set.
func (o *RegistrationDetailsResponseObject) HasFailSubscriptionOnAutomationFailure() bool {
	if o != nil && !IsNil(o.FailSubscriptionOnAutomationFailure) {
		return true
	}

	return false
}

// SetFailSubscriptionOnAutomationFailure gets a reference to the given bool and assigns it to the FailSubscriptionOnAutomationFailure field.
func (o *RegistrationDetailsResponseObject) SetFailSubscriptionOnAutomationFailure(v bool) {
	o.FailSubscriptionOnAutomationFailure = &v
}

// GetFormationSolutionName returns the FormationSolutionName field value if set, zero value otherwise.
// Deprecated
func (o *RegistrationDetailsResponseObject) GetFormationSolutionName() string {
	if o == nil || IsNil(o.FormationSolutionName) {
		var ret string
		return ret
	}
	return *o.FormationSolutionName
}

// GetFormationSolutionNameOk returns a tuple with the FormationSolutionName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *RegistrationDetailsResponseObject) GetFormationSolutionNameOk() (*string, bool) {
	if o == nil || IsNil(o.FormationSolutionName) {
		return nil, false
	}
	return o.FormationSolutionName, true
}

// HasFormationSolutionName returns a boolean if a field has been set.
func (o *RegistrationDetailsResponseObject) HasFormationSolutionName() bool {
	if o != nil && !IsNil(o.FormationSolutionName) {
		return true
	}

	return false
}

// SetFormationSolutionName gets a reference to the given string and assigns it to the FormationSolutionName field.
// Deprecated
func (o *RegistrationDetailsResponseObject) SetFormationSolutionName(v string) {
	o.FormationSolutionName = &v
}

// GetGlobalAccountId returns the GlobalAccountId field value if set, zero value otherwise.
func (o *RegistrationDetailsResponseObject) GetGlobalAccountId() string {
	if o == nil || IsNil(o.GlobalAccountId) {
		var ret string
		return ret
	}
	return *o.GlobalAccountId
}

// GetGlobalAccountIdOk returns a tuple with the GlobalAccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationDetailsResponseObject) GetGlobalAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.GlobalAccountId) {
		return nil, false
	}
	return o.GlobalAccountId, true
}

// HasGlobalAccountId returns a boolean if a field has been set.
func (o *RegistrationDetailsResponseObject) HasGlobalAccountId() bool {
	if o != nil && !IsNil(o.GlobalAccountId) {
		return true
	}

	return false
}

// SetGlobalAccountId gets a reference to the given string and assigns it to the GlobalAccountId field.
func (o *RegistrationDetailsResponseObject) SetGlobalAccountId(v string) {
	o.GlobalAccountId = &v
}

// GetMessageOnAutomationFailure returns the MessageOnAutomationFailure field value if set, zero value otherwise.
func (o *RegistrationDetailsResponseObject) GetMessageOnAutomationFailure() string {
	if o == nil || IsNil(o.MessageOnAutomationFailure) {
		var ret string
		return ret
	}
	return *o.MessageOnAutomationFailure
}

// GetMessageOnAutomationFailureOk returns a tuple with the MessageOnAutomationFailure field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationDetailsResponseObject) GetMessageOnAutomationFailureOk() (*string, bool) {
	if o == nil || IsNil(o.MessageOnAutomationFailure) {
		return nil, false
	}
	return o.MessageOnAutomationFailure, true
}

// HasMessageOnAutomationFailure returns a boolean if a field has been set.
func (o *RegistrationDetailsResponseObject) HasMessageOnAutomationFailure() bool {
	if o != nil && !IsNil(o.MessageOnAutomationFailure) {
		return true
	}

	return false
}

// SetMessageOnAutomationFailure gets a reference to the given string and assigns it to the MessageOnAutomationFailure field.
func (o *RegistrationDetailsResponseObject) SetMessageOnAutomationFailure(v string) {
	o.MessageOnAutomationFailure = &v
}

// GetOrganizationGuid returns the OrganizationGuid field value if set, zero value otherwise.
func (o *RegistrationDetailsResponseObject) GetOrganizationGuid() string {
	if o == nil || IsNil(o.OrganizationGuid) {
		var ret string
		return ret
	}
	return *o.OrganizationGuid
}

// GetOrganizationGuidOk returns a tuple with the OrganizationGuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationDetailsResponseObject) GetOrganizationGuidOk() (*string, bool) {
	if o == nil || IsNil(o.OrganizationGuid) {
		return nil, false
	}
	return o.OrganizationGuid, true
}

// HasOrganizationGuid returns a boolean if a field has been set.
func (o *RegistrationDetailsResponseObject) HasOrganizationGuid() bool {
	if o != nil && !IsNil(o.OrganizationGuid) {
		return true
	}

	return false
}

// SetOrganizationGuid gets a reference to the given string and assigns it to the OrganizationGuid field.
func (o *RegistrationDetailsResponseObject) SetOrganizationGuid(v string) {
	o.OrganizationGuid = &v
}

// GetParamsSchema returns the ParamsSchema field value if set, zero value otherwise.
func (o *RegistrationDetailsResponseObject) GetParamsSchema() EntitledApplicationsResponseObjectParamsSchema {
	if o == nil || IsNil(o.ParamsSchema) {
		var ret EntitledApplicationsResponseObjectParamsSchema
		return ret
	}
	return *o.ParamsSchema
}

// GetParamsSchemaOk returns a tuple with the ParamsSchema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationDetailsResponseObject) GetParamsSchemaOk() (*EntitledApplicationsResponseObjectParamsSchema, bool) {
	if o == nil || IsNil(o.ParamsSchema) {
		return nil, false
	}
	return o.ParamsSchema, true
}

// HasParamsSchema returns a boolean if a field has been set.
func (o *RegistrationDetailsResponseObject) HasParamsSchema() bool {
	if o != nil && !IsNil(o.ParamsSchema) {
		return true
	}

	return false
}

// SetParamsSchema gets a reference to the given EntitledApplicationsResponseObjectParamsSchema and assigns it to the ParamsSchema field.
func (o *RegistrationDetailsResponseObject) SetParamsSchema(v EntitledApplicationsResponseObjectParamsSchema) {
	o.ParamsSchema = &v
}

// GetPropagateParams returns the PropagateParams field value if set, zero value otherwise.
func (o *RegistrationDetailsResponseObject) GetPropagateParams() bool {
	if o == nil || IsNil(o.PropagateParams) {
		var ret bool
		return ret
	}
	return *o.PropagateParams
}

// GetPropagateParamsOk returns a tuple with the PropagateParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationDetailsResponseObject) GetPropagateParamsOk() (*bool, bool) {
	if o == nil || IsNil(o.PropagateParams) {
		return nil, false
	}
	return o.PropagateParams, true
}

// HasPropagateParams returns a boolean if a field has been set.
func (o *RegistrationDetailsResponseObject) HasPropagateParams() bool {
	if o != nil && !IsNil(o.PropagateParams) {
		return true
	}

	return false
}

// SetPropagateParams gets a reference to the given bool and assigns it to the PropagateParams field.
func (o *RegistrationDetailsResponseObject) SetPropagateParams(v bool) {
	o.PropagateParams = &v
}

// GetProviderTenantId returns the ProviderTenantId field value if set, zero value otherwise.
func (o *RegistrationDetailsResponseObject) GetProviderTenantId() string {
	if o == nil || IsNil(o.ProviderTenantId) {
		var ret string
		return ret
	}
	return *o.ProviderTenantId
}

// GetProviderTenantIdOk returns a tuple with the ProviderTenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationDetailsResponseObject) GetProviderTenantIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProviderTenantId) {
		return nil, false
	}
	return o.ProviderTenantId, true
}

// HasProviderTenantId returns a boolean if a field has been set.
func (o *RegistrationDetailsResponseObject) HasProviderTenantId() bool {
	if o != nil && !IsNil(o.ProviderTenantId) {
		return true
	}

	return false
}

// SetProviderTenantId gets a reference to the given string and assigns it to the ProviderTenantId field.
func (o *RegistrationDetailsResponseObject) SetProviderTenantId(v string) {
	o.ProviderTenantId = &v
}

// GetServiceInstanceId returns the ServiceInstanceId field value if set, zero value otherwise.
func (o *RegistrationDetailsResponseObject) GetServiceInstanceId() string {
	if o == nil || IsNil(o.ServiceInstanceId) {
		var ret string
		return ret
	}
	return *o.ServiceInstanceId
}

// GetServiceInstanceIdOk returns a tuple with the ServiceInstanceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationDetailsResponseObject) GetServiceInstanceIdOk() (*string, bool) {
	if o == nil || IsNil(o.ServiceInstanceId) {
		return nil, false
	}
	return o.ServiceInstanceId, true
}

// HasServiceInstanceId returns a boolean if a field has been set.
func (o *RegistrationDetailsResponseObject) HasServiceInstanceId() bool {
	if o != nil && !IsNil(o.ServiceInstanceId) {
		return true
	}

	return false
}

// SetServiceInstanceId gets a reference to the given string and assigns it to the ServiceInstanceId field.
func (o *RegistrationDetailsResponseObject) SetServiceInstanceId(v string) {
	o.ServiceInstanceId = &v
}

// GetSpaceGuid returns the SpaceGuid field value if set, zero value otherwise.
func (o *RegistrationDetailsResponseObject) GetSpaceGuid() string {
	if o == nil || IsNil(o.SpaceGuid) {
		var ret string
		return ret
	}
	return *o.SpaceGuid
}

// GetSpaceGuidOk returns a tuple with the SpaceGuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationDetailsResponseObject) GetSpaceGuidOk() (*string, bool) {
	if o == nil || IsNil(o.SpaceGuid) {
		return nil, false
	}
	return o.SpaceGuid, true
}

// HasSpaceGuid returns a boolean if a field has been set.
func (o *RegistrationDetailsResponseObject) HasSpaceGuid() bool {
	if o != nil && !IsNil(o.SpaceGuid) {
		return true
	}

	return false
}

// SetSpaceGuid gets a reference to the given string and assigns it to the SpaceGuid field.
func (o *RegistrationDetailsResponseObject) SetSpaceGuid(v string) {
	o.SpaceGuid = &v
}

// GetXsappname returns the Xsappname field value if set, zero value otherwise.
func (o *RegistrationDetailsResponseObject) GetXsappname() string {
	if o == nil || IsNil(o.Xsappname) {
		var ret string
		return ret
	}
	return *o.Xsappname
}

// GetXsappnameOk returns a tuple with the Xsappname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationDetailsResponseObject) GetXsappnameOk() (*string, bool) {
	if o == nil || IsNil(o.Xsappname) {
		return nil, false
	}
	return o.Xsappname, true
}

// HasXsappname returns a boolean if a field has been set.
func (o *RegistrationDetailsResponseObject) HasXsappname() bool {
	if o != nil && !IsNil(o.Xsappname) {
		return true
	}

	return false
}

// SetXsappname gets a reference to the given string and assigns it to the Xsappname field.
func (o *RegistrationDetailsResponseObject) SetXsappname(v string) {
	o.Xsappname = &v
}

func (o RegistrationDetailsResponseObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RegistrationDetailsResponseObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AppId) {
		toSerialize["appId"] = o.AppId
	}
	if !IsNil(o.AppName) {
		toSerialize["appName"] = o.AppName
	}
	if !IsNil(o.AppType) {
		toSerialize["appType"] = o.AppType
	}
	if !IsNil(o.AppUrls) {
		toSerialize["appUrls"] = o.AppUrls
	}
	if !IsNil(o.AutomationSolutionName) {
		toSerialize["automationSolutionName"] = o.AutomationSolutionName
	}
	if !IsNil(o.Category) {
		toSerialize["category"] = o.Category
	}
	if !IsNil(o.CommercialAppName) {
		toSerialize["commercialAppName"] = o.CommercialAppName
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.FailSubscriptionOnAutomationFailure) {
		toSerialize["failSubscriptionOnAutomationFailure"] = o.FailSubscriptionOnAutomationFailure
	}
	if !IsNil(o.FormationSolutionName) {
		toSerialize["formationSolutionName"] = o.FormationSolutionName
	}
	if !IsNil(o.GlobalAccountId) {
		toSerialize["globalAccountId"] = o.GlobalAccountId
	}
	if !IsNil(o.MessageOnAutomationFailure) {
		toSerialize["messageOnAutomationFailure"] = o.MessageOnAutomationFailure
	}
	if !IsNil(o.OrganizationGuid) {
		toSerialize["organizationGuid"] = o.OrganizationGuid
	}
	if !IsNil(o.ParamsSchema) {
		toSerialize["paramsSchema"] = o.ParamsSchema
	}
	if !IsNil(o.PropagateParams) {
		toSerialize["propagateParams"] = o.PropagateParams
	}
	if !IsNil(o.ProviderTenantId) {
		toSerialize["providerTenantId"] = o.ProviderTenantId
	}
	if !IsNil(o.ServiceInstanceId) {
		toSerialize["serviceInstanceId"] = o.ServiceInstanceId
	}
	if !IsNil(o.SpaceGuid) {
		toSerialize["spaceGuid"] = o.SpaceGuid
	}
	if !IsNil(o.Xsappname) {
		toSerialize["xsappname"] = o.Xsappname
	}
	return toSerialize, nil
}

type NullableRegistrationDetailsResponseObject struct {
	value *RegistrationDetailsResponseObject
	isSet bool
}

func (v NullableRegistrationDetailsResponseObject) Get() *RegistrationDetailsResponseObject {
	return v.value
}

func (v *NullableRegistrationDetailsResponseObject) Set(val *RegistrationDetailsResponseObject) {
	v.value = val
	v.isSet = true
}

func (v NullableRegistrationDetailsResponseObject) IsSet() bool {
	return v.isSet
}

func (v *NullableRegistrationDetailsResponseObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegistrationDetailsResponseObject(val *RegistrationDetailsResponseObject) *NullableRegistrationDetailsResponseObject {
	return &NullableRegistrationDetailsResponseObject{value: val, isSet: true}
}

func (v NullableRegistrationDetailsResponseObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegistrationDetailsResponseObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


