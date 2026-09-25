# SelectorIndexResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Models** | **int32** | Slugs that resolve to an offering. | 
**Offerings** | **int32** | Selectors the deployment serves. | 
**PinnedSelectors** | **int32** | Pinned spellings, one per instance a model is offered on. | 

## Methods

### NewSelectorIndexResponse

`func NewSelectorIndexResponse(models int32, offerings int32, pinnedSelectors int32, ) *SelectorIndexResponse`

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


### GetPinnedSelectors

`func (o *SelectorIndexResponse) GetPinnedSelectors() int32`

GetPinnedSelectors returns the PinnedSelectors field if non-nil, zero value otherwise.

### GetPinnedSelectorsOk

`func (o *SelectorIndexResponse) GetPinnedSelectorsOk() (*int32, bool)`

GetPinnedSelectorsOk returns a tuple with the PinnedSelectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinnedSelectors

`func (o *SelectorIndexResponse) SetPinnedSelectors(v int32)`

SetPinnedSelectors sets PinnedSelectors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


