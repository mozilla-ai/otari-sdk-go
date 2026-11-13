# RatesDocumentInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Multipliers** | Pointer to [**[]RateMultiplier**](RateMultiplier.md) |  | [optional] 
**WebSearch** | Pointer to [**NullableUnitRateInput**](UnitRateInput.md) |  | [optional] 

## Methods

### NewRatesDocumentInput

`func NewRatesDocumentInput() *RatesDocumentInput`

NewRatesDocumentInput instantiates a new RatesDocumentInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRatesDocumentInputWithDefaults

`func NewRatesDocumentInputWithDefaults() *RatesDocumentInput`

NewRatesDocumentInputWithDefaults instantiates a new RatesDocumentInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMultipliers

`func (o *RatesDocumentInput) GetMultipliers() []RateMultiplier`

GetMultipliers returns the Multipliers field if non-nil, zero value otherwise.

### GetMultipliersOk

`func (o *RatesDocumentInput) GetMultipliersOk() (*[]RateMultiplier, bool)`

GetMultipliersOk returns a tuple with the Multipliers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultipliers

`func (o *RatesDocumentInput) SetMultipliers(v []RateMultiplier)`

SetMultipliers sets Multipliers field to given value.

### HasMultipliers

`func (o *RatesDocumentInput) HasMultipliers() bool`

HasMultipliers returns a boolean if a field has been set.

### GetWebSearch

`func (o *RatesDocumentInput) GetWebSearch() UnitRateInput`

GetWebSearch returns the WebSearch field if non-nil, zero value otherwise.

### GetWebSearchOk

`func (o *RatesDocumentInput) GetWebSearchOk() (*UnitRateInput, bool)`

GetWebSearchOk returns a tuple with the WebSearch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebSearch

`func (o *RatesDocumentInput) SetWebSearch(v UnitRateInput)`

SetWebSearch sets WebSearch field to given value.

### HasWebSearch

`func (o *RatesDocumentInput) HasWebSearch() bool`

HasWebSearch returns a boolean if a field has been set.

### SetWebSearchNil

`func (o *RatesDocumentInput) SetWebSearchNil(b bool)`

 SetWebSearchNil sets the value for WebSearch to be an explicit nil

### UnsetWebSearch
`func (o *RatesDocumentInput) UnsetWebSearch()`

UnsetWebSearch ensures that no value is present for WebSearch, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


