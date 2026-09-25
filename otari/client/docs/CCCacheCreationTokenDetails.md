# CCCacheCreationTokenDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ephemeral5mInputTokens** | Pointer to **NullableInt32** | Filter to a single failure status code (e.g. 429 for provider rate limits, 402 for missing-pricing rejections). Only error rows carry one, so this filter also restricts to status&#x3D;&#39;error&#39; unless &#39;status&#39; is given explicitly | [optional] 
**Ephemeral1hInputTokens** | Pointer to **NullableInt32** | Filter to a single failure status code (e.g. 429 for provider rate limits, 402 for missing-pricing rejections). Only error rows carry one, so this filter also restricts to status&#x3D;&#39;error&#39; unless &#39;status&#39; is given explicitly | [optional] 

## Methods

### NewCCCacheCreationTokenDetails

`func NewCCCacheCreationTokenDetails() *CCCacheCreationTokenDetails`

NewCCCacheCreationTokenDetails instantiates a new CCCacheCreationTokenDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCCCacheCreationTokenDetailsWithDefaults

`func NewCCCacheCreationTokenDetailsWithDefaults() *CCCacheCreationTokenDetails`

NewCCCacheCreationTokenDetailsWithDefaults instantiates a new CCCacheCreationTokenDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEphemeral5mInputTokens

`func (o *CCCacheCreationTokenDetails) GetEphemeral5mInputTokens() int32`

GetEphemeral5mInputTokens returns the Ephemeral5mInputTokens field if non-nil, zero value otherwise.

### GetEphemeral5mInputTokensOk

`func (o *CCCacheCreationTokenDetails) GetEphemeral5mInputTokensOk() (*int32, bool)`

GetEphemeral5mInputTokensOk returns a tuple with the Ephemeral5mInputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEphemeral5mInputTokens

`func (o *CCCacheCreationTokenDetails) SetEphemeral5mInputTokens(v int32)`

SetEphemeral5mInputTokens sets Ephemeral5mInputTokens field to given value.

### HasEphemeral5mInputTokens

`func (o *CCCacheCreationTokenDetails) HasEphemeral5mInputTokens() bool`

HasEphemeral5mInputTokens returns a boolean if a field has been set.

### SetEphemeral5mInputTokensNil

`func (o *CCCacheCreationTokenDetails) SetEphemeral5mInputTokensNil(b bool)`

 SetEphemeral5mInputTokensNil sets the value for Ephemeral5mInputTokens to be an explicit nil

### UnsetEphemeral5mInputTokens
`func (o *CCCacheCreationTokenDetails) UnsetEphemeral5mInputTokens()`

UnsetEphemeral5mInputTokens ensures that no value is present for Ephemeral5mInputTokens, not even an explicit nil
### GetEphemeral1hInputTokens

`func (o *CCCacheCreationTokenDetails) GetEphemeral1hInputTokens() int32`

GetEphemeral1hInputTokens returns the Ephemeral1hInputTokens field if non-nil, zero value otherwise.

### GetEphemeral1hInputTokensOk

`func (o *CCCacheCreationTokenDetails) GetEphemeral1hInputTokensOk() (*int32, bool)`

GetEphemeral1hInputTokensOk returns a tuple with the Ephemeral1hInputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEphemeral1hInputTokens

`func (o *CCCacheCreationTokenDetails) SetEphemeral1hInputTokens(v int32)`

SetEphemeral1hInputTokens sets Ephemeral1hInputTokens field to given value.

### HasEphemeral1hInputTokens

`func (o *CCCacheCreationTokenDetails) HasEphemeral1hInputTokens() bool`

HasEphemeral1hInputTokens returns a boolean if a field has been set.

### SetEphemeral1hInputTokensNil

`func (o *CCCacheCreationTokenDetails) SetEphemeral1hInputTokensNil(b bool)`

 SetEphemeral1hInputTokensNil sets the value for Ephemeral1hInputTokens to be an explicit nil

### UnsetEphemeral1hInputTokens
`func (o *CCCacheCreationTokenDetails) UnsetEphemeral1hInputTokens()`

UnsetEphemeral1hInputTokens ensures that no value is present for Ephemeral1hInputTokens, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


