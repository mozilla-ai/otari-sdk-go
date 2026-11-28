# ChatCompletionChunk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Choices** | [**[]CCKChunkChoice**](CCKChunkChoice.md) |  | 
**Created** | **int32** |  | 
**Model** | **string** |  | 
**Object** | **string** |  | 
**Moderation** | Pointer to [**NullableCCKModeration**](CCKModeration.md) |  | [optional] 
**Obfuscation** | Pointer to **NullableString** | Filter to a single event type or metric name (e.g. &#39;tool_result&#39;, &#39;claude_code.commit.count&#39;) | [optional] 
**ServiceTier** | Pointer to **NullableString** |  | [optional] 
**SystemFingerprint** | Pointer to **NullableString** |  | [optional] 
**Usage** | Pointer to [**NullableCCKCompletionUsage**](CCKCompletionUsage.md) |  | [optional] 

## Methods

### NewChatCompletionChunk

`func NewChatCompletionChunk(id string, choices []CCKChunkChoice, created int32, model string, object string, ) *ChatCompletionChunk`

NewChatCompletionChunk instantiates a new ChatCompletionChunk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatCompletionChunkWithDefaults

`func NewChatCompletionChunkWithDefaults() *ChatCompletionChunk`

NewChatCompletionChunkWithDefaults instantiates a new ChatCompletionChunk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ChatCompletionChunk) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ChatCompletionChunk) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ChatCompletionChunk) SetId(v string)`

SetId sets Id field to given value.


### GetChoices

`func (o *ChatCompletionChunk) GetChoices() []CCKChunkChoice`

GetChoices returns the Choices field if non-nil, zero value otherwise.

### GetChoicesOk

`func (o *ChatCompletionChunk) GetChoicesOk() (*[]CCKChunkChoice, bool)`

GetChoicesOk returns a tuple with the Choices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChoices

`func (o *ChatCompletionChunk) SetChoices(v []CCKChunkChoice)`

SetChoices sets Choices field to given value.


### GetCreated

`func (o *ChatCompletionChunk) GetCreated() int32`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ChatCompletionChunk) GetCreatedOk() (*int32, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ChatCompletionChunk) SetCreated(v int32)`

SetCreated sets Created field to given value.


### GetModel

`func (o *ChatCompletionChunk) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ChatCompletionChunk) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ChatCompletionChunk) SetModel(v string)`

SetModel sets Model field to given value.


### GetObject

`func (o *ChatCompletionChunk) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ChatCompletionChunk) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ChatCompletionChunk) SetObject(v string)`

SetObject sets Object field to given value.


### GetModeration

`func (o *ChatCompletionChunk) GetModeration() CCKModeration`

GetModeration returns the Moderation field if non-nil, zero value otherwise.

### GetModerationOk

`func (o *ChatCompletionChunk) GetModerationOk() (*CCKModeration, bool)`

GetModerationOk returns a tuple with the Moderation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModeration

`func (o *ChatCompletionChunk) SetModeration(v CCKModeration)`

SetModeration sets Moderation field to given value.

### HasModeration

`func (o *ChatCompletionChunk) HasModeration() bool`

HasModeration returns a boolean if a field has been set.

### SetModerationNil

`func (o *ChatCompletionChunk) SetModerationNil(b bool)`

 SetModerationNil sets the value for Moderation to be an explicit nil

### UnsetModeration
`func (o *ChatCompletionChunk) UnsetModeration()`

UnsetModeration ensures that no value is present for Moderation, not even an explicit nil
### GetObfuscation

`func (o *ChatCompletionChunk) GetObfuscation() string`

GetObfuscation returns the Obfuscation field if non-nil, zero value otherwise.

### GetObfuscationOk

`func (o *ChatCompletionChunk) GetObfuscationOk() (*string, bool)`

GetObfuscationOk returns a tuple with the Obfuscation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObfuscation

`func (o *ChatCompletionChunk) SetObfuscation(v string)`

SetObfuscation sets Obfuscation field to given value.

### HasObfuscation

`func (o *ChatCompletionChunk) HasObfuscation() bool`

HasObfuscation returns a boolean if a field has been set.

### SetObfuscationNil

`func (o *ChatCompletionChunk) SetObfuscationNil(b bool)`

 SetObfuscationNil sets the value for Obfuscation to be an explicit nil

### UnsetObfuscation
`func (o *ChatCompletionChunk) UnsetObfuscation()`

UnsetObfuscation ensures that no value is present for Obfuscation, not even an explicit nil
### GetServiceTier

`func (o *ChatCompletionChunk) GetServiceTier() string`

GetServiceTier returns the ServiceTier field if non-nil, zero value otherwise.

### GetServiceTierOk

`func (o *ChatCompletionChunk) GetServiceTierOk() (*string, bool)`

GetServiceTierOk returns a tuple with the ServiceTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceTier

`func (o *ChatCompletionChunk) SetServiceTier(v string)`

SetServiceTier sets ServiceTier field to given value.

### HasServiceTier

`func (o *ChatCompletionChunk) HasServiceTier() bool`

HasServiceTier returns a boolean if a field has been set.

### SetServiceTierNil

`func (o *ChatCompletionChunk) SetServiceTierNil(b bool)`

 SetServiceTierNil sets the value for ServiceTier to be an explicit nil

### UnsetServiceTier
`func (o *ChatCompletionChunk) UnsetServiceTier()`

UnsetServiceTier ensures that no value is present for ServiceTier, not even an explicit nil
### GetSystemFingerprint

`func (o *ChatCompletionChunk) GetSystemFingerprint() string`

GetSystemFingerprint returns the SystemFingerprint field if non-nil, zero value otherwise.

### GetSystemFingerprintOk

`func (o *ChatCompletionChunk) GetSystemFingerprintOk() (*string, bool)`

GetSystemFingerprintOk returns a tuple with the SystemFingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemFingerprint

`func (o *ChatCompletionChunk) SetSystemFingerprint(v string)`

SetSystemFingerprint sets SystemFingerprint field to given value.

### HasSystemFingerprint

`func (o *ChatCompletionChunk) HasSystemFingerprint() bool`

HasSystemFingerprint returns a boolean if a field has been set.

### SetSystemFingerprintNil

`func (o *ChatCompletionChunk) SetSystemFingerprintNil(b bool)`

 SetSystemFingerprintNil sets the value for SystemFingerprint to be an explicit nil

### UnsetSystemFingerprint
`func (o *ChatCompletionChunk) UnsetSystemFingerprint()`

UnsetSystemFingerprint ensures that no value is present for SystemFingerprint, not even an explicit nil
### GetUsage

`func (o *ChatCompletionChunk) GetUsage() CCKCompletionUsage`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *ChatCompletionChunk) GetUsageOk() (*CCKCompletionUsage, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *ChatCompletionChunk) SetUsage(v CCKCompletionUsage)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *ChatCompletionChunk) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### SetUsageNil

`func (o *ChatCompletionChunk) SetUsageNil(b bool)`

 SetUsageNil sets the value for Usage to be an explicit nil

### UnsetUsage
`func (o *ChatCompletionChunk) UnsetUsage()`

UnsetUsage ensures that no value is present for Usage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


