# CurrentPricingPage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** |  | 
**Data** | [**[]PricingResponse**](PricingResponse.md) |  | 

## Methods

### NewCurrentPricingPage

`func NewCurrentPricingPage(count int32, data []PricingResponse, ) *CurrentPricingPage`

NewCurrentPricingPage instantiates a new CurrentPricingPage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCurrentPricingPageWithDefaults

`func NewCurrentPricingPageWithDefaults() *CurrentPricingPage`

NewCurrentPricingPageWithDefaults instantiates a new CurrentPricingPage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *CurrentPricingPage) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *CurrentPricingPage) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *CurrentPricingPage) SetCount(v int32)`

SetCount sets Count field to given value.


### GetData

`func (o *CurrentPricingPage) GetData() []PricingResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CurrentPricingPage) GetDataOk() (*[]PricingResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CurrentPricingPage) SetData(v []PricingResponse)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


