# OrganizationSsoEnforcementPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BreakGlassUserId** | **NullableString** | One of the organization&#39;s own owners who keeps a non-SSO sign-in route, where the organization named one. | 
**BreakGlassWaived** | **bool** | Whether the organization has explicitly waived naming a break-glass owner. | 
**Required** | **bool** | Whether this organization requires sign-in through its own connection. | 

## Methods

### NewOrganizationSsoEnforcementPublic

`func NewOrganizationSsoEnforcementPublic(breakGlassUserId NullableString, breakGlassWaived bool, required bool, ) *OrganizationSsoEnforcementPublic`

NewOrganizationSsoEnforcementPublic instantiates a new OrganizationSsoEnforcementPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationSsoEnforcementPublicWithDefaults

`func NewOrganizationSsoEnforcementPublicWithDefaults() *OrganizationSsoEnforcementPublic`

NewOrganizationSsoEnforcementPublicWithDefaults instantiates a new OrganizationSsoEnforcementPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBreakGlassUserId

`func (o *OrganizationSsoEnforcementPublic) GetBreakGlassUserId() string`

GetBreakGlassUserId returns the BreakGlassUserId field if non-nil, zero value otherwise.

### GetBreakGlassUserIdOk

`func (o *OrganizationSsoEnforcementPublic) GetBreakGlassUserIdOk() (*string, bool)`

GetBreakGlassUserIdOk returns a tuple with the BreakGlassUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakGlassUserId

`func (o *OrganizationSsoEnforcementPublic) SetBreakGlassUserId(v string)`

SetBreakGlassUserId sets BreakGlassUserId field to given value.


### SetBreakGlassUserIdNil

`func (o *OrganizationSsoEnforcementPublic) SetBreakGlassUserIdNil(b bool)`

 SetBreakGlassUserIdNil sets the value for BreakGlassUserId to be an explicit nil

### UnsetBreakGlassUserId
`func (o *OrganizationSsoEnforcementPublic) UnsetBreakGlassUserId()`

UnsetBreakGlassUserId ensures that no value is present for BreakGlassUserId, not even an explicit nil
### GetBreakGlassWaived

`func (o *OrganizationSsoEnforcementPublic) GetBreakGlassWaived() bool`

GetBreakGlassWaived returns the BreakGlassWaived field if non-nil, zero value otherwise.

### GetBreakGlassWaivedOk

`func (o *OrganizationSsoEnforcementPublic) GetBreakGlassWaivedOk() (*bool, bool)`

GetBreakGlassWaivedOk returns a tuple with the BreakGlassWaived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakGlassWaived

`func (o *OrganizationSsoEnforcementPublic) SetBreakGlassWaived(v bool)`

SetBreakGlassWaived sets BreakGlassWaived field to given value.


### GetRequired

`func (o *OrganizationSsoEnforcementPublic) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *OrganizationSsoEnforcementPublic) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *OrganizationSsoEnforcementPublic) SetRequired(v bool)`

SetRequired sets Required field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


