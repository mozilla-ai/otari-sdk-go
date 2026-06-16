# ChatCompletion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Choices** | [**[]CCChoice**](CCChoice.md) |  | 
**Created** | **int32** |  | 
**Model** | **string** |  | 
**Object** | **string** |  | 
**ServiceTier** | Pointer to **NullableString** |  | [optional] 
**SystemFingerprint** | Pointer to **NullableString** | Filter models by provider name | [optional] 
**Usage** | Pointer to [**NullableCCCompletionUsage**](CCCompletionUsage.md) |  | [optional] 

## Methods

### NewChatCompletion

`func NewChatCompletion(id string, choices []CCChoice, created int32, model string, object string, ) *ChatCompletion`

NewChatCompletion instantiates a new ChatCompletion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatCompletionWithDefaults

`func NewChatCompletionWithDefaults() *ChatCompletion`

NewChatCompletionWithDefaults instantiates a new ChatCompletion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ChatCompletion) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ChatCompletion) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ChatCompletion) SetId(v string)`

SetId sets Id field to given value.


### GetChoices

`func (o *ChatCompletion) GetChoices() []CCChoice`

GetChoices returns the Choices field if non-nil, zero value otherwise.

### GetChoicesOk

`func (o *ChatCompletion) GetChoicesOk() (*[]CCChoice, bool)`

GetChoicesOk returns a tuple with the Choices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChoices

`func (o *ChatCompletion) SetChoices(v []CCChoice)`

SetChoices sets Choices field to given value.


### GetCreated

`func (o *ChatCompletion) GetCreated() int32`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ChatCompletion) GetCreatedOk() (*int32, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ChatCompletion) SetCreated(v int32)`

SetCreated sets Created field to given value.


### GetModel

`func (o *ChatCompletion) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ChatCompletion) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ChatCompletion) SetModel(v string)`

SetModel sets Model field to given value.


### GetObject

`func (o *ChatCompletion) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ChatCompletion) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ChatCompletion) SetObject(v string)`

SetObject sets Object field to given value.


### GetServiceTier

`func (o *ChatCompletion) GetServiceTier() string`

GetServiceTier returns the ServiceTier field if non-nil, zero value otherwise.

### GetServiceTierOk

`func (o *ChatCompletion) GetServiceTierOk() (*string, bool)`

GetServiceTierOk returns a tuple with the ServiceTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceTier

`func (o *ChatCompletion) SetServiceTier(v string)`

SetServiceTier sets ServiceTier field to given value.

### HasServiceTier

`func (o *ChatCompletion) HasServiceTier() bool`

HasServiceTier returns a boolean if a field has been set.

### SetServiceTierNil

`func (o *ChatCompletion) SetServiceTierNil(b bool)`

 SetServiceTierNil sets the value for ServiceTier to be an explicit nil

### UnsetServiceTier
`func (o *ChatCompletion) UnsetServiceTier()`

UnsetServiceTier ensures that no value is present for ServiceTier, not even an explicit nil
### GetSystemFingerprint

`func (o *ChatCompletion) GetSystemFingerprint() string`

GetSystemFingerprint returns the SystemFingerprint field if non-nil, zero value otherwise.

### GetSystemFingerprintOk

`func (o *ChatCompletion) GetSystemFingerprintOk() (*string, bool)`

GetSystemFingerprintOk returns a tuple with the SystemFingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemFingerprint

`func (o *ChatCompletion) SetSystemFingerprint(v string)`

SetSystemFingerprint sets SystemFingerprint field to given value.

### HasSystemFingerprint

`func (o *ChatCompletion) HasSystemFingerprint() bool`

HasSystemFingerprint returns a boolean if a field has been set.

### SetSystemFingerprintNil

`func (o *ChatCompletion) SetSystemFingerprintNil(b bool)`

 SetSystemFingerprintNil sets the value for SystemFingerprint to be an explicit nil

### UnsetSystemFingerprint
`func (o *ChatCompletion) UnsetSystemFingerprint()`

UnsetSystemFingerprint ensures that no value is present for SystemFingerprint, not even an explicit nil
### GetUsage

`func (o *ChatCompletion) GetUsage() CCCompletionUsage`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *ChatCompletion) GetUsageOk() (*CCCompletionUsage, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *ChatCompletion) SetUsage(v CCCompletionUsage)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *ChatCompletion) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### SetUsageNil

`func (o *ChatCompletion) SetUsageNil(b bool)`

 SetUsageNil sets the value for Usage to be an explicit nil

### UnsetUsage
`func (o *ChatCompletion) UnsetUsage()`

UnsetUsage ensures that no value is present for Usage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


