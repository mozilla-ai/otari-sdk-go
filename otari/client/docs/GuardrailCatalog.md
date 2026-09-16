# GuardrailCatalog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | **bool** | Whether the guardrails service answered with its profiles | 
**Profiles** | Pointer to [**[]GuardrailProfileSpec**](GuardrailProfileSpec.md) |  | [optional] 
**Reason** | Pointer to **NullableString** | Why the catalog is unavailable, in terms a tenant can act on | [optional] 

## Methods

### NewGuardrailCatalog

`func NewGuardrailCatalog(available bool, ) *GuardrailCatalog`

NewGuardrailCatalog instantiates a new GuardrailCatalog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGuardrailCatalogWithDefaults

`func NewGuardrailCatalogWithDefaults() *GuardrailCatalog`

NewGuardrailCatalogWithDefaults instantiates a new GuardrailCatalog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *GuardrailCatalog) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *GuardrailCatalog) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *GuardrailCatalog) SetAvailable(v bool)`

SetAvailable sets Available field to given value.


### GetProfiles

`func (o *GuardrailCatalog) GetProfiles() []GuardrailProfileSpec`

GetProfiles returns the Profiles field if non-nil, zero value otherwise.

### GetProfilesOk

`func (o *GuardrailCatalog) GetProfilesOk() (*[]GuardrailProfileSpec, bool)`

GetProfilesOk returns a tuple with the Profiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfiles

`func (o *GuardrailCatalog) SetProfiles(v []GuardrailProfileSpec)`

SetProfiles sets Profiles field to given value.

### HasProfiles

`func (o *GuardrailCatalog) HasProfiles() bool`

HasProfiles returns a boolean if a field has been set.

### GetReason

`func (o *GuardrailCatalog) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *GuardrailCatalog) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *GuardrailCatalog) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *GuardrailCatalog) HasReason() bool`

HasReason returns a boolean if a field has been set.

### SetReasonNil

`func (o *GuardrailCatalog) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *GuardrailCatalog) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


