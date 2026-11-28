# CCKPromptTokensDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AudioTokens** | Pointer to **NullableInt32** |  | [optional] 
**CacheWriteTokens** | Pointer to **NullableInt32** |  | [optional] 
**CachedTokens** | Pointer to **NullableInt32** |  | [optional] 
**ImageTokens** | Pointer to **NullableInt32** |  | [optional] 
**TextTokens** | Pointer to **NullableInt32** |  | [optional] 
**CacheCreationTokenDetails** | Pointer to [**NullableCCKCacheCreationTokenDetails**](CCKCacheCreationTokenDetails.md) |  | [optional] 
**CachedAudioTokens** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewCCKPromptTokensDetails

`func NewCCKPromptTokensDetails() *CCKPromptTokensDetails`

NewCCKPromptTokensDetails instantiates a new CCKPromptTokensDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCCKPromptTokensDetailsWithDefaults

`func NewCCKPromptTokensDetailsWithDefaults() *CCKPromptTokensDetails`

NewCCKPromptTokensDetailsWithDefaults instantiates a new CCKPromptTokensDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudioTokens

`func (o *CCKPromptTokensDetails) GetAudioTokens() int32`

GetAudioTokens returns the AudioTokens field if non-nil, zero value otherwise.

### GetAudioTokensOk

`func (o *CCKPromptTokensDetails) GetAudioTokensOk() (*int32, bool)`

GetAudioTokensOk returns a tuple with the AudioTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudioTokens

`func (o *CCKPromptTokensDetails) SetAudioTokens(v int32)`

SetAudioTokens sets AudioTokens field to given value.

### HasAudioTokens

`func (o *CCKPromptTokensDetails) HasAudioTokens() bool`

HasAudioTokens returns a boolean if a field has been set.

### SetAudioTokensNil

`func (o *CCKPromptTokensDetails) SetAudioTokensNil(b bool)`

 SetAudioTokensNil sets the value for AudioTokens to be an explicit nil

### UnsetAudioTokens
`func (o *CCKPromptTokensDetails) UnsetAudioTokens()`

UnsetAudioTokens ensures that no value is present for AudioTokens, not even an explicit nil
### GetCacheWriteTokens

`func (o *CCKPromptTokensDetails) GetCacheWriteTokens() int32`

GetCacheWriteTokens returns the CacheWriteTokens field if non-nil, zero value otherwise.

### GetCacheWriteTokensOk

`func (o *CCKPromptTokensDetails) GetCacheWriteTokensOk() (*int32, bool)`

GetCacheWriteTokensOk returns a tuple with the CacheWriteTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheWriteTokens

`func (o *CCKPromptTokensDetails) SetCacheWriteTokens(v int32)`

SetCacheWriteTokens sets CacheWriteTokens field to given value.

### HasCacheWriteTokens

`func (o *CCKPromptTokensDetails) HasCacheWriteTokens() bool`

HasCacheWriteTokens returns a boolean if a field has been set.

### SetCacheWriteTokensNil

`func (o *CCKPromptTokensDetails) SetCacheWriteTokensNil(b bool)`

 SetCacheWriteTokensNil sets the value for CacheWriteTokens to be an explicit nil

### UnsetCacheWriteTokens
`func (o *CCKPromptTokensDetails) UnsetCacheWriteTokens()`

UnsetCacheWriteTokens ensures that no value is present for CacheWriteTokens, not even an explicit nil
### GetCachedTokens

`func (o *CCKPromptTokensDetails) GetCachedTokens() int32`

GetCachedTokens returns the CachedTokens field if non-nil, zero value otherwise.

### GetCachedTokensOk

`func (o *CCKPromptTokensDetails) GetCachedTokensOk() (*int32, bool)`

GetCachedTokensOk returns a tuple with the CachedTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCachedTokens

`func (o *CCKPromptTokensDetails) SetCachedTokens(v int32)`

SetCachedTokens sets CachedTokens field to given value.

### HasCachedTokens

`func (o *CCKPromptTokensDetails) HasCachedTokens() bool`

HasCachedTokens returns a boolean if a field has been set.

### SetCachedTokensNil

`func (o *CCKPromptTokensDetails) SetCachedTokensNil(b bool)`

 SetCachedTokensNil sets the value for CachedTokens to be an explicit nil

### UnsetCachedTokens
`func (o *CCKPromptTokensDetails) UnsetCachedTokens()`

UnsetCachedTokens ensures that no value is present for CachedTokens, not even an explicit nil
### GetImageTokens

`func (o *CCKPromptTokensDetails) GetImageTokens() int32`

GetImageTokens returns the ImageTokens field if non-nil, zero value otherwise.

### GetImageTokensOk

`func (o *CCKPromptTokensDetails) GetImageTokensOk() (*int32, bool)`

GetImageTokensOk returns a tuple with the ImageTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageTokens

`func (o *CCKPromptTokensDetails) SetImageTokens(v int32)`

SetImageTokens sets ImageTokens field to given value.

### HasImageTokens

`func (o *CCKPromptTokensDetails) HasImageTokens() bool`

HasImageTokens returns a boolean if a field has been set.

### SetImageTokensNil

`func (o *CCKPromptTokensDetails) SetImageTokensNil(b bool)`

 SetImageTokensNil sets the value for ImageTokens to be an explicit nil

### UnsetImageTokens
`func (o *CCKPromptTokensDetails) UnsetImageTokens()`

UnsetImageTokens ensures that no value is present for ImageTokens, not even an explicit nil
### GetTextTokens

`func (o *CCKPromptTokensDetails) GetTextTokens() int32`

GetTextTokens returns the TextTokens field if non-nil, zero value otherwise.

### GetTextTokensOk

`func (o *CCKPromptTokensDetails) GetTextTokensOk() (*int32, bool)`

GetTextTokensOk returns a tuple with the TextTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextTokens

`func (o *CCKPromptTokensDetails) SetTextTokens(v int32)`

SetTextTokens sets TextTokens field to given value.

### HasTextTokens

`func (o *CCKPromptTokensDetails) HasTextTokens() bool`

HasTextTokens returns a boolean if a field has been set.

### SetTextTokensNil

`func (o *CCKPromptTokensDetails) SetTextTokensNil(b bool)`

 SetTextTokensNil sets the value for TextTokens to be an explicit nil

### UnsetTextTokens
`func (o *CCKPromptTokensDetails) UnsetTextTokens()`

UnsetTextTokens ensures that no value is present for TextTokens, not even an explicit nil
### GetCacheCreationTokenDetails

`func (o *CCKPromptTokensDetails) GetCacheCreationTokenDetails() CCKCacheCreationTokenDetails`

GetCacheCreationTokenDetails returns the CacheCreationTokenDetails field if non-nil, zero value otherwise.

### GetCacheCreationTokenDetailsOk

`func (o *CCKPromptTokensDetails) GetCacheCreationTokenDetailsOk() (*CCKCacheCreationTokenDetails, bool)`

GetCacheCreationTokenDetailsOk returns a tuple with the CacheCreationTokenDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheCreationTokenDetails

`func (o *CCKPromptTokensDetails) SetCacheCreationTokenDetails(v CCKCacheCreationTokenDetails)`

SetCacheCreationTokenDetails sets CacheCreationTokenDetails field to given value.

### HasCacheCreationTokenDetails

`func (o *CCKPromptTokensDetails) HasCacheCreationTokenDetails() bool`

HasCacheCreationTokenDetails returns a boolean if a field has been set.

### SetCacheCreationTokenDetailsNil

`func (o *CCKPromptTokensDetails) SetCacheCreationTokenDetailsNil(b bool)`

 SetCacheCreationTokenDetailsNil sets the value for CacheCreationTokenDetails to be an explicit nil

### UnsetCacheCreationTokenDetails
`func (o *CCKPromptTokensDetails) UnsetCacheCreationTokenDetails()`

UnsetCacheCreationTokenDetails ensures that no value is present for CacheCreationTokenDetails, not even an explicit nil
### GetCachedAudioTokens

`func (o *CCKPromptTokensDetails) GetCachedAudioTokens() int32`

GetCachedAudioTokens returns the CachedAudioTokens field if non-nil, zero value otherwise.

### GetCachedAudioTokensOk

`func (o *CCKPromptTokensDetails) GetCachedAudioTokensOk() (*int32, bool)`

GetCachedAudioTokensOk returns a tuple with the CachedAudioTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCachedAudioTokens

`func (o *CCKPromptTokensDetails) SetCachedAudioTokens(v int32)`

SetCachedAudioTokens sets CachedAudioTokens field to given value.

### HasCachedAudioTokens

`func (o *CCKPromptTokensDetails) HasCachedAudioTokens() bool`

HasCachedAudioTokens returns a boolean if a field has been set.

### SetCachedAudioTokensNil

`func (o *CCKPromptTokensDetails) SetCachedAudioTokensNil(b bool)`

 SetCachedAudioTokensNil sets the value for CachedAudioTokens to be an explicit nil

### UnsetCachedAudioTokens
`func (o *CCKPromptTokensDetails) UnsetCachedAudioTokens()`

UnsetCachedAudioTokens ensures that no value is present for CachedAudioTokens, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


