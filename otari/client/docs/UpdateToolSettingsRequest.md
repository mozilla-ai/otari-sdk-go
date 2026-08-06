# UpdateToolSettingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GuardrailsUrl** | Pointer to **NullableString** |  | [optional] 
**SandboxPurposeHint** | Pointer to **NullableString** |  | [optional] 
**SandboxUrl** | Pointer to **NullableString** |  | [optional] 
**WebSearchEngines** | Pointer to **NullableString** |  | [optional] 
**WebSearchExtract** | Pointer to **NullableBool** |  | [optional] 
**WebSearchMaxResults** | Pointer to **NullableInt32** |  | [optional] 
**WebSearchPurposeHint** | Pointer to **NullableString** |  | [optional] 
**WebSearchUrl** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewUpdateToolSettingsRequest

`func NewUpdateToolSettingsRequest() *UpdateToolSettingsRequest`

NewUpdateToolSettingsRequest instantiates a new UpdateToolSettingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateToolSettingsRequestWithDefaults

`func NewUpdateToolSettingsRequestWithDefaults() *UpdateToolSettingsRequest`

NewUpdateToolSettingsRequestWithDefaults instantiates a new UpdateToolSettingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGuardrailsUrl

`func (o *UpdateToolSettingsRequest) GetGuardrailsUrl() string`

GetGuardrailsUrl returns the GuardrailsUrl field if non-nil, zero value otherwise.

### GetGuardrailsUrlOk

`func (o *UpdateToolSettingsRequest) GetGuardrailsUrlOk() (*string, bool)`

GetGuardrailsUrlOk returns a tuple with the GuardrailsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuardrailsUrl

`func (o *UpdateToolSettingsRequest) SetGuardrailsUrl(v string)`

SetGuardrailsUrl sets GuardrailsUrl field to given value.

### HasGuardrailsUrl

`func (o *UpdateToolSettingsRequest) HasGuardrailsUrl() bool`

HasGuardrailsUrl returns a boolean if a field has been set.

### SetGuardrailsUrlNil

`func (o *UpdateToolSettingsRequest) SetGuardrailsUrlNil(b bool)`

 SetGuardrailsUrlNil sets the value for GuardrailsUrl to be an explicit nil

### UnsetGuardrailsUrl
`func (o *UpdateToolSettingsRequest) UnsetGuardrailsUrl()`

UnsetGuardrailsUrl ensures that no value is present for GuardrailsUrl, not even an explicit nil
### GetSandboxPurposeHint

`func (o *UpdateToolSettingsRequest) GetSandboxPurposeHint() string`

GetSandboxPurposeHint returns the SandboxPurposeHint field if non-nil, zero value otherwise.

### GetSandboxPurposeHintOk

`func (o *UpdateToolSettingsRequest) GetSandboxPurposeHintOk() (*string, bool)`

GetSandboxPurposeHintOk returns a tuple with the SandboxPurposeHint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSandboxPurposeHint

`func (o *UpdateToolSettingsRequest) SetSandboxPurposeHint(v string)`

SetSandboxPurposeHint sets SandboxPurposeHint field to given value.

### HasSandboxPurposeHint

`func (o *UpdateToolSettingsRequest) HasSandboxPurposeHint() bool`

HasSandboxPurposeHint returns a boolean if a field has been set.

### SetSandboxPurposeHintNil

`func (o *UpdateToolSettingsRequest) SetSandboxPurposeHintNil(b bool)`

 SetSandboxPurposeHintNil sets the value for SandboxPurposeHint to be an explicit nil

### UnsetSandboxPurposeHint
`func (o *UpdateToolSettingsRequest) UnsetSandboxPurposeHint()`

UnsetSandboxPurposeHint ensures that no value is present for SandboxPurposeHint, not even an explicit nil
### GetSandboxUrl

`func (o *UpdateToolSettingsRequest) GetSandboxUrl() string`

GetSandboxUrl returns the SandboxUrl field if non-nil, zero value otherwise.

### GetSandboxUrlOk

`func (o *UpdateToolSettingsRequest) GetSandboxUrlOk() (*string, bool)`

GetSandboxUrlOk returns a tuple with the SandboxUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSandboxUrl

`func (o *UpdateToolSettingsRequest) SetSandboxUrl(v string)`

SetSandboxUrl sets SandboxUrl field to given value.

### HasSandboxUrl

`func (o *UpdateToolSettingsRequest) HasSandboxUrl() bool`

HasSandboxUrl returns a boolean if a field has been set.

### SetSandboxUrlNil

`func (o *UpdateToolSettingsRequest) SetSandboxUrlNil(b bool)`

 SetSandboxUrlNil sets the value for SandboxUrl to be an explicit nil

### UnsetSandboxUrl
`func (o *UpdateToolSettingsRequest) UnsetSandboxUrl()`

UnsetSandboxUrl ensures that no value is present for SandboxUrl, not even an explicit nil
### GetWebSearchEngines

`func (o *UpdateToolSettingsRequest) GetWebSearchEngines() string`

GetWebSearchEngines returns the WebSearchEngines field if non-nil, zero value otherwise.

### GetWebSearchEnginesOk

`func (o *UpdateToolSettingsRequest) GetWebSearchEnginesOk() (*string, bool)`

GetWebSearchEnginesOk returns a tuple with the WebSearchEngines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebSearchEngines

`func (o *UpdateToolSettingsRequest) SetWebSearchEngines(v string)`

SetWebSearchEngines sets WebSearchEngines field to given value.

### HasWebSearchEngines

`func (o *UpdateToolSettingsRequest) HasWebSearchEngines() bool`

HasWebSearchEngines returns a boolean if a field has been set.

### SetWebSearchEnginesNil

`func (o *UpdateToolSettingsRequest) SetWebSearchEnginesNil(b bool)`

 SetWebSearchEnginesNil sets the value for WebSearchEngines to be an explicit nil

### UnsetWebSearchEngines
`func (o *UpdateToolSettingsRequest) UnsetWebSearchEngines()`

UnsetWebSearchEngines ensures that no value is present for WebSearchEngines, not even an explicit nil
### GetWebSearchExtract

`func (o *UpdateToolSettingsRequest) GetWebSearchExtract() bool`

GetWebSearchExtract returns the WebSearchExtract field if non-nil, zero value otherwise.

### GetWebSearchExtractOk

`func (o *UpdateToolSettingsRequest) GetWebSearchExtractOk() (*bool, bool)`

GetWebSearchExtractOk returns a tuple with the WebSearchExtract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebSearchExtract

`func (o *UpdateToolSettingsRequest) SetWebSearchExtract(v bool)`

SetWebSearchExtract sets WebSearchExtract field to given value.

### HasWebSearchExtract

`func (o *UpdateToolSettingsRequest) HasWebSearchExtract() bool`

HasWebSearchExtract returns a boolean if a field has been set.

### SetWebSearchExtractNil

`func (o *UpdateToolSettingsRequest) SetWebSearchExtractNil(b bool)`

 SetWebSearchExtractNil sets the value for WebSearchExtract to be an explicit nil

### UnsetWebSearchExtract
`func (o *UpdateToolSettingsRequest) UnsetWebSearchExtract()`

UnsetWebSearchExtract ensures that no value is present for WebSearchExtract, not even an explicit nil
### GetWebSearchMaxResults

`func (o *UpdateToolSettingsRequest) GetWebSearchMaxResults() int32`

GetWebSearchMaxResults returns the WebSearchMaxResults field if non-nil, zero value otherwise.

### GetWebSearchMaxResultsOk

`func (o *UpdateToolSettingsRequest) GetWebSearchMaxResultsOk() (*int32, bool)`

GetWebSearchMaxResultsOk returns a tuple with the WebSearchMaxResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebSearchMaxResults

`func (o *UpdateToolSettingsRequest) SetWebSearchMaxResults(v int32)`

SetWebSearchMaxResults sets WebSearchMaxResults field to given value.

### HasWebSearchMaxResults

`func (o *UpdateToolSettingsRequest) HasWebSearchMaxResults() bool`

HasWebSearchMaxResults returns a boolean if a field has been set.

### SetWebSearchMaxResultsNil

`func (o *UpdateToolSettingsRequest) SetWebSearchMaxResultsNil(b bool)`

 SetWebSearchMaxResultsNil sets the value for WebSearchMaxResults to be an explicit nil

### UnsetWebSearchMaxResults
`func (o *UpdateToolSettingsRequest) UnsetWebSearchMaxResults()`

UnsetWebSearchMaxResults ensures that no value is present for WebSearchMaxResults, not even an explicit nil
### GetWebSearchPurposeHint

`func (o *UpdateToolSettingsRequest) GetWebSearchPurposeHint() string`

GetWebSearchPurposeHint returns the WebSearchPurposeHint field if non-nil, zero value otherwise.

### GetWebSearchPurposeHintOk

`func (o *UpdateToolSettingsRequest) GetWebSearchPurposeHintOk() (*string, bool)`

GetWebSearchPurposeHintOk returns a tuple with the WebSearchPurposeHint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebSearchPurposeHint

`func (o *UpdateToolSettingsRequest) SetWebSearchPurposeHint(v string)`

SetWebSearchPurposeHint sets WebSearchPurposeHint field to given value.

### HasWebSearchPurposeHint

`func (o *UpdateToolSettingsRequest) HasWebSearchPurposeHint() bool`

HasWebSearchPurposeHint returns a boolean if a field has been set.

### SetWebSearchPurposeHintNil

`func (o *UpdateToolSettingsRequest) SetWebSearchPurposeHintNil(b bool)`

 SetWebSearchPurposeHintNil sets the value for WebSearchPurposeHint to be an explicit nil

### UnsetWebSearchPurposeHint
`func (o *UpdateToolSettingsRequest) UnsetWebSearchPurposeHint()`

UnsetWebSearchPurposeHint ensures that no value is present for WebSearchPurposeHint, not even an explicit nil
### GetWebSearchUrl

`func (o *UpdateToolSettingsRequest) GetWebSearchUrl() string`

GetWebSearchUrl returns the WebSearchUrl field if non-nil, zero value otherwise.

### GetWebSearchUrlOk

`func (o *UpdateToolSettingsRequest) GetWebSearchUrlOk() (*string, bool)`

GetWebSearchUrlOk returns a tuple with the WebSearchUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebSearchUrl

`func (o *UpdateToolSettingsRequest) SetWebSearchUrl(v string)`

SetWebSearchUrl sets WebSearchUrl field to given value.

### HasWebSearchUrl

`func (o *UpdateToolSettingsRequest) HasWebSearchUrl() bool`

HasWebSearchUrl returns a boolean if a field has been set.

### SetWebSearchUrlNil

`func (o *UpdateToolSettingsRequest) SetWebSearchUrlNil(b bool)`

 SetWebSearchUrlNil sets the value for WebSearchUrl to be an explicit nil

### UnsetWebSearchUrl
`func (o *UpdateToolSettingsRequest) UnsetWebSearchUrl()`

UnsetWebSearchUrl ensures that no value is present for WebSearchUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


