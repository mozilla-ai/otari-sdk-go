# RerankMeta

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BilledUnits** | Pointer to **map[string]float32** |  | [optional] 
**Tokens** | Pointer to **map[string]int32** |  | [optional] 

## Methods

### NewRerankMeta

`func NewRerankMeta() *RerankMeta`

NewRerankMeta instantiates a new RerankMeta object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRerankMetaWithDefaults

`func NewRerankMetaWithDefaults() *RerankMeta`

NewRerankMetaWithDefaults instantiates a new RerankMeta object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBilledUnits

`func (o *RerankMeta) GetBilledUnits() map[string]float32`

GetBilledUnits returns the BilledUnits field if non-nil, zero value otherwise.

### GetBilledUnitsOk

`func (o *RerankMeta) GetBilledUnitsOk() (*map[string]float32, bool)`

GetBilledUnitsOk returns a tuple with the BilledUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBilledUnits

`func (o *RerankMeta) SetBilledUnits(v map[string]float32)`

SetBilledUnits sets BilledUnits field to given value.

### HasBilledUnits

`func (o *RerankMeta) HasBilledUnits() bool`

HasBilledUnits returns a boolean if a field has been set.

### SetBilledUnitsNil

`func (o *RerankMeta) SetBilledUnitsNil(b bool)`

 SetBilledUnitsNil sets the value for BilledUnits to be an explicit nil

### UnsetBilledUnits
`func (o *RerankMeta) UnsetBilledUnits()`

UnsetBilledUnits ensures that no value is present for BilledUnits, not even an explicit nil
### GetTokens

`func (o *RerankMeta) GetTokens() map[string]int32`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *RerankMeta) GetTokensOk() (*map[string]int32, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *RerankMeta) SetTokens(v map[string]int32)`

SetTokens sets Tokens field to given value.

### HasTokens

`func (o *RerankMeta) HasTokens() bool`

HasTokens returns a boolean if a field has been set.

### SetTokensNil

`func (o *RerankMeta) SetTokensNil(b bool)`

 SetTokensNil sets the value for Tokens to be an explicit nil

### UnsetTokens
`func (o *RerankMeta) UnsetTokens()`

UnsetTokens ensures that no value is present for Tokens, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


