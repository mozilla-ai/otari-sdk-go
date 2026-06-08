# UsageEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKeyId** | **NullableString** |  | 
**CompletionTokens** | **NullableInt32** |  | 
**Cost** | **NullableFloat32** |  | 
**Endpoint** | **string** |  | 
**ErrorMessage** | **NullableString** |  | 
**Id** | **string** |  | 
**Model** | **string** |  | 
**PromptTokens** | **NullableInt32** |  | 
**Provider** | **NullableString** |  | 
**Status** | **string** |  | 
**Timestamp** | **string** |  | 
**TotalTokens** | **NullableInt32** |  | 
**UserId** | **NullableString** |  | 

## Methods

### NewUsageEntry

`func NewUsageEntry(apiKeyId NullableString, completionTokens NullableInt32, cost NullableFloat32, endpoint string, errorMessage NullableString, id string, model string, promptTokens NullableInt32, provider NullableString, status string, timestamp string, totalTokens NullableInt32, userId NullableString, ) *UsageEntry`

NewUsageEntry instantiates a new UsageEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageEntryWithDefaults

`func NewUsageEntryWithDefaults() *UsageEntry`

NewUsageEntryWithDefaults instantiates a new UsageEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKeyId

`func (o *UsageEntry) GetApiKeyId() string`

GetApiKeyId returns the ApiKeyId field if non-nil, zero value otherwise.

### GetApiKeyIdOk

`func (o *UsageEntry) GetApiKeyIdOk() (*string, bool)`

GetApiKeyIdOk returns a tuple with the ApiKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKeyId

`func (o *UsageEntry) SetApiKeyId(v string)`

SetApiKeyId sets ApiKeyId field to given value.


### SetApiKeyIdNil

`func (o *UsageEntry) SetApiKeyIdNil(b bool)`

 SetApiKeyIdNil sets the value for ApiKeyId to be an explicit nil

### UnsetApiKeyId
`func (o *UsageEntry) UnsetApiKeyId()`

UnsetApiKeyId ensures that no value is present for ApiKeyId, not even an explicit nil
### GetCompletionTokens

`func (o *UsageEntry) GetCompletionTokens() int32`

GetCompletionTokens returns the CompletionTokens field if non-nil, zero value otherwise.

### GetCompletionTokensOk

`func (o *UsageEntry) GetCompletionTokensOk() (*int32, bool)`

GetCompletionTokensOk returns a tuple with the CompletionTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionTokens

`func (o *UsageEntry) SetCompletionTokens(v int32)`

SetCompletionTokens sets CompletionTokens field to given value.


### SetCompletionTokensNil

`func (o *UsageEntry) SetCompletionTokensNil(b bool)`

 SetCompletionTokensNil sets the value for CompletionTokens to be an explicit nil

### UnsetCompletionTokens
`func (o *UsageEntry) UnsetCompletionTokens()`

UnsetCompletionTokens ensures that no value is present for CompletionTokens, not even an explicit nil
### GetCost

`func (o *UsageEntry) GetCost() float32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *UsageEntry) GetCostOk() (*float32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *UsageEntry) SetCost(v float32)`

SetCost sets Cost field to given value.


### SetCostNil

`func (o *UsageEntry) SetCostNil(b bool)`

 SetCostNil sets the value for Cost to be an explicit nil

### UnsetCost
`func (o *UsageEntry) UnsetCost()`

UnsetCost ensures that no value is present for Cost, not even an explicit nil
### GetEndpoint

`func (o *UsageEntry) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *UsageEntry) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *UsageEntry) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.


### GetErrorMessage

`func (o *UsageEntry) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *UsageEntry) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *UsageEntry) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.


### SetErrorMessageNil

`func (o *UsageEntry) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *UsageEntry) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetId

`func (o *UsageEntry) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UsageEntry) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UsageEntry) SetId(v string)`

SetId sets Id field to given value.


### GetModel

`func (o *UsageEntry) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *UsageEntry) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *UsageEntry) SetModel(v string)`

SetModel sets Model field to given value.


### GetPromptTokens

`func (o *UsageEntry) GetPromptTokens() int32`

GetPromptTokens returns the PromptTokens field if non-nil, zero value otherwise.

### GetPromptTokensOk

`func (o *UsageEntry) GetPromptTokensOk() (*int32, bool)`

GetPromptTokensOk returns a tuple with the PromptTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptTokens

`func (o *UsageEntry) SetPromptTokens(v int32)`

SetPromptTokens sets PromptTokens field to given value.


### SetPromptTokensNil

`func (o *UsageEntry) SetPromptTokensNil(b bool)`

 SetPromptTokensNil sets the value for PromptTokens to be an explicit nil

### UnsetPromptTokens
`func (o *UsageEntry) UnsetPromptTokens()`

UnsetPromptTokens ensures that no value is present for PromptTokens, not even an explicit nil
### GetProvider

`func (o *UsageEntry) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *UsageEntry) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *UsageEntry) SetProvider(v string)`

SetProvider sets Provider field to given value.


### SetProviderNil

`func (o *UsageEntry) SetProviderNil(b bool)`

 SetProviderNil sets the value for Provider to be an explicit nil

### UnsetProvider
`func (o *UsageEntry) UnsetProvider()`

UnsetProvider ensures that no value is present for Provider, not even an explicit nil
### GetStatus

`func (o *UsageEntry) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UsageEntry) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UsageEntry) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTimestamp

`func (o *UsageEntry) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *UsageEntry) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *UsageEntry) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.


### GetTotalTokens

`func (o *UsageEntry) GetTotalTokens() int32`

GetTotalTokens returns the TotalTokens field if non-nil, zero value otherwise.

### GetTotalTokensOk

`func (o *UsageEntry) GetTotalTokensOk() (*int32, bool)`

GetTotalTokensOk returns a tuple with the TotalTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTokens

`func (o *UsageEntry) SetTotalTokens(v int32)`

SetTotalTokens sets TotalTokens field to given value.


### SetTotalTokensNil

`func (o *UsageEntry) SetTotalTokensNil(b bool)`

 SetTotalTokensNil sets the value for TotalTokens to be an explicit nil

### UnsetTotalTokens
`func (o *UsageEntry) UnsetTotalTokens()`

UnsetTotalTokens ensures that no value is present for TotalTokens, not even an explicit nil
### GetUserId

`func (o *UsageEntry) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UsageEntry) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UsageEntry) SetUserId(v string)`

SetUserId sets UserId field to given value.


### SetUserIdNil

`func (o *UsageEntry) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *UsageEntry) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


