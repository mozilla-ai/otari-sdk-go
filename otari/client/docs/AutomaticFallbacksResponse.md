# AutomaticFallbacksResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Models** | [**[]AutomaticFallbackPublic**](AutomaticFallbackPublic.md) | Most requested first. | 
**WindowDays** | **int32** | How many days of traffic the counts cover. | 

## Methods

### NewAutomaticFallbacksResponse

`func NewAutomaticFallbacksResponse(models []AutomaticFallbackPublic, windowDays int32, ) *AutomaticFallbacksResponse`

NewAutomaticFallbacksResponse instantiates a new AutomaticFallbacksResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutomaticFallbacksResponseWithDefaults

`func NewAutomaticFallbacksResponseWithDefaults() *AutomaticFallbacksResponse`

NewAutomaticFallbacksResponseWithDefaults instantiates a new AutomaticFallbacksResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModels

`func (o *AutomaticFallbacksResponse) GetModels() []AutomaticFallbackPublic`

GetModels returns the Models field if non-nil, zero value otherwise.

### GetModelsOk

`func (o *AutomaticFallbacksResponse) GetModelsOk() (*[]AutomaticFallbackPublic, bool)`

GetModelsOk returns a tuple with the Models field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModels

`func (o *AutomaticFallbacksResponse) SetModels(v []AutomaticFallbackPublic)`

SetModels sets Models field to given value.


### GetWindowDays

`func (o *AutomaticFallbacksResponse) GetWindowDays() int32`

GetWindowDays returns the WindowDays field if non-nil, zero value otherwise.

### GetWindowDaysOk

`func (o *AutomaticFallbacksResponse) GetWindowDaysOk() (*int32, bool)`

GetWindowDaysOk returns a tuple with the WindowDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowDays

`func (o *AutomaticFallbacksResponse) SetWindowDays(v int32)`

SetWindowDays sets WindowDays field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


