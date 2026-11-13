# CCServerToolUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WebSearchRequests** | Pointer to **NullableInt32** | Filter to a single failure status code (e.g. 429 for provider rate limits, 402 for missing-pricing rejections). Only error rows carry one, so this filter also restricts to status&#x3D;&#39;error&#39; unless &#39;status&#39; is given explicitly | [optional] 

## Methods

### NewCCServerToolUsage

`func NewCCServerToolUsage() *CCServerToolUsage`

NewCCServerToolUsage instantiates a new CCServerToolUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCCServerToolUsageWithDefaults

`func NewCCServerToolUsageWithDefaults() *CCServerToolUsage`

NewCCServerToolUsageWithDefaults instantiates a new CCServerToolUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWebSearchRequests

`func (o *CCServerToolUsage) GetWebSearchRequests() int32`

GetWebSearchRequests returns the WebSearchRequests field if non-nil, zero value otherwise.

### GetWebSearchRequestsOk

`func (o *CCServerToolUsage) GetWebSearchRequestsOk() (*int32, bool)`

GetWebSearchRequestsOk returns a tuple with the WebSearchRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebSearchRequests

`func (o *CCServerToolUsage) SetWebSearchRequests(v int32)`

SetWebSearchRequests sets WebSearchRequests field to given value.

### HasWebSearchRequests

`func (o *CCServerToolUsage) HasWebSearchRequests() bool`

HasWebSearchRequests returns a boolean if a field has been set.

### SetWebSearchRequestsNil

`func (o *CCServerToolUsage) SetWebSearchRequestsNil(b bool)`

 SetWebSearchRequestsNil sets the value for WebSearchRequests to be an explicit nil

### UnsetWebSearchRequests
`func (o *CCServerToolUsage) UnsetWebSearchRequests()`

UnsetWebSearchRequests ensures that no value is present for WebSearchRequests, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


