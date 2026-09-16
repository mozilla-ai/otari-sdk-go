# SelectorIndexResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Models** | **int32** | Slugs that resolve to an offering. | 
**Offerings** | **int32** | Selectors the deployment serves. | 
**ShortSelectors** | **int32** | Offerings with an unambiguous short spelling. | 

## Methods

### NewSelectorIndexResponse

`func NewSelectorIndexResponse(models int32, offerings int32, shortSelectors int32, ) *SelectorIndexResponse`

NewSelectorIndexResponse instantiates a new SelectorIndexResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelectorIndexResponseWithDefaults

`func NewSelectorIndexResponseWithDefaults() *SelectorIndexResponse`

NewSelectorIndexResponseWithDefaults instantiates a new SelectorIndexResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModels

`func (o *SelectorIndexResponse) GetModels() int32`

GetModels returns the Models field if non-nil, zero value otherwise.

### GetModelsOk

`func (o *SelectorIndexResponse) GetModelsOk() (*int32, bool)`

GetModelsOk returns a tuple with the Models field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModels

`func (o *SelectorIndexResponse) SetModels(v int32)`

SetModels sets Models field to given value.


### GetOfferings

`func (o *SelectorIndexResponse) GetOfferings() int32`

GetOfferings returns the Offerings field if non-nil, zero value otherwise.

### GetOfferingsOk

`func (o *SelectorIndexResponse) GetOfferingsOk() (*int32, bool)`

GetOfferingsOk returns a tuple with the Offerings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferings

`func (o *SelectorIndexResponse) SetOfferings(v int32)`

SetOfferings sets Offerings field to given value.


### GetShortSelectors

`func (o *SelectorIndexResponse) GetShortSelectors() int32`

GetShortSelectors returns the ShortSelectors field if non-nil, zero value otherwise.

### GetShortSelectorsOk

`func (o *SelectorIndexResponse) GetShortSelectorsOk() (*int32, bool)`

GetShortSelectorsOk returns a tuple with the ShortSelectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortSelectors

`func (o *SelectorIndexResponse) SetShortSelectors(v int32)`

SetShortSelectors sets ShortSelectors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


