# OrganizationSsoEnforcementRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BreakGlassUserId** | Pointer to **NullableString** | One of the organization&#39;s own owners who keeps a non-SSO sign-in route when required is true. | [optional] 
**BreakGlassWaived** | Pointer to **bool** | Whether the organization explicitly waives a break-glass owner when enforcement is required. | [optional] [default to false]
**Required** | **bool** | Whether this organization requires sign-in through its own connection. | 

## Methods

### NewOrganizationSsoEnforcementRequest

`func NewOrganizationSsoEnforcementRequest(required bool, ) *OrganizationSsoEnforcementRequest`

NewOrganizationSsoEnforcementRequest instantiates a new OrganizationSsoEnforcementRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationSsoEnforcementRequestWithDefaults

`func NewOrganizationSsoEnforcementRequestWithDefaults() *OrganizationSsoEnforcementRequest`

NewOrganizationSsoEnforcementRequestWithDefaults instantiates a new OrganizationSsoEnforcementRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBreakGlassUserId

`func (o *OrganizationSsoEnforcementRequest) GetBreakGlassUserId() string`

GetBreakGlassUserId returns the BreakGlassUserId field if non-nil, zero value otherwise.

### GetBreakGlassUserIdOk

`func (o *OrganizationSsoEnforcementRequest) GetBreakGlassUserIdOk() (*string, bool)`

GetBreakGlassUserIdOk returns a tuple with the BreakGlassUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakGlassUserId

`func (o *OrganizationSsoEnforcementRequest) SetBreakGlassUserId(v string)`

SetBreakGlassUserId sets BreakGlassUserId field to given value.

### HasBreakGlassUserId

`func (o *OrganizationSsoEnforcementRequest) HasBreakGlassUserId() bool`

HasBreakGlassUserId returns a boolean if a field has been set.

### SetBreakGlassUserIdNil

`func (o *OrganizationSsoEnforcementRequest) SetBreakGlassUserIdNil(b bool)`

 SetBreakGlassUserIdNil sets the value for BreakGlassUserId to be an explicit nil

### UnsetBreakGlassUserId
`func (o *OrganizationSsoEnforcementRequest) UnsetBreakGlassUserId()`

UnsetBreakGlassUserId ensures that no value is present for BreakGlassUserId, not even an explicit nil
### GetBreakGlassWaived

`func (o *OrganizationSsoEnforcementRequest) GetBreakGlassWaived() bool`

GetBreakGlassWaived returns the BreakGlassWaived field if non-nil, zero value otherwise.

### GetBreakGlassWaivedOk

`func (o *OrganizationSsoEnforcementRequest) GetBreakGlassWaivedOk() (*bool, bool)`

GetBreakGlassWaivedOk returns a tuple with the BreakGlassWaived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakGlassWaived

`func (o *OrganizationSsoEnforcementRequest) SetBreakGlassWaived(v bool)`

SetBreakGlassWaived sets BreakGlassWaived field to given value.

### HasBreakGlassWaived

`func (o *OrganizationSsoEnforcementRequest) HasBreakGlassWaived() bool`

HasBreakGlassWaived returns a boolean if a field has been set.

### GetRequired

`func (o *OrganizationSsoEnforcementRequest) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *OrganizationSsoEnforcementRequest) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *OrganizationSsoEnforcementRequest) SetRequired(v bool)`

SetRequired sets Required field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


