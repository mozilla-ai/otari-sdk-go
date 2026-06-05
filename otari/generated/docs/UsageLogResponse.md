# UsageLogResponse

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

### NewUsageLogResponse

`func NewUsageLogResponse(apiKeyId NullableString, completionTokens NullableInt32, cost NullableFloat32, endpoint string, errorMessage NullableString, id string, model string, promptTokens NullableInt32, provider NullableString, status string, timestamp string, totalTokens NullableInt32, userId NullableString, ) *UsageLogResponse`

NewUsageLogResponse instantiates a new UsageLogResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsageLogResponseWithDefaults

`func NewUsageLogResponseWithDefaults() *UsageLogResponse`

NewUsageLogResponseWithDefaults instantiates a new UsageLogResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKeyId

`func (o *UsageLogResponse) GetApiKeyId() string`

GetApiKeyId returns the ApiKeyId field if non-nil, zero value otherwise.

### GetApiKeyIdOk

`func (o *UsageLogResponse) GetApiKeyIdOk() (*string, bool)`

GetApiKeyIdOk returns a tuple with the ApiKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKeyId

`func (o *UsageLogResponse) SetApiKeyId(v string)`

SetApiKeyId sets ApiKeyId field to given value.


### SetApiKeyIdNil

`func (o *UsageLogResponse) SetApiKeyIdNil(b bool)`

 SetApiKeyIdNil sets the value for ApiKeyId to be an explicit nil

### UnsetApiKeyId
`func (o *UsageLogResponse) UnsetApiKeyId()`

UnsetApiKeyId ensures that no value is present for ApiKeyId, not even an explicit nil
### GetCompletionTokens

`func (o *UsageLogResponse) GetCompletionTokens() int32`

GetCompletionTokens returns the CompletionTokens field if non-nil, zero value otherwise.

### GetCompletionTokensOk

`func (o *UsageLogResponse) GetCompletionTokensOk() (*int32, bool)`

GetCompletionTokensOk returns a tuple with the CompletionTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionTokens

`func (o *UsageLogResponse) SetCompletionTokens(v int32)`

SetCompletionTokens sets CompletionTokens field to given value.


### SetCompletionTokensNil

`func (o *UsageLogResponse) SetCompletionTokensNil(b bool)`

 SetCompletionTokensNil sets the value for CompletionTokens to be an explicit nil

### UnsetCompletionTokens
`func (o *UsageLogResponse) UnsetCompletionTokens()`

UnsetCompletionTokens ensures that no value is present for CompletionTokens, not even an explicit nil
### GetCost

`func (o *UsageLogResponse) GetCost() float32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *UsageLogResponse) GetCostOk() (*float32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *UsageLogResponse) SetCost(v float32)`

SetCost sets Cost field to given value.


### SetCostNil

`func (o *UsageLogResponse) SetCostNil(b bool)`

 SetCostNil sets the value for Cost to be an explicit nil

### UnsetCost
`func (o *UsageLogResponse) UnsetCost()`

UnsetCost ensures that no value is present for Cost, not even an explicit nil
### GetEndpoint

`func (o *UsageLogResponse) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *UsageLogResponse) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *UsageLogResponse) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.


### GetErrorMessage

`func (o *UsageLogResponse) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *UsageLogResponse) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *UsageLogResponse) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.


### SetErrorMessageNil

`func (o *UsageLogResponse) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *UsageLogResponse) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetId

`func (o *UsageLogResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UsageLogResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UsageLogResponse) SetId(v string)`

SetId sets Id field to given value.


### GetModel

`func (o *UsageLogResponse) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *UsageLogResponse) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *UsageLogResponse) SetModel(v string)`

SetModel sets Model field to given value.


### GetPromptTokens

`func (o *UsageLogResponse) GetPromptTokens() int32`

GetPromptTokens returns the PromptTokens field if non-nil, zero value otherwise.

### GetPromptTokensOk

`func (o *UsageLogResponse) GetPromptTokensOk() (*int32, bool)`

GetPromptTokensOk returns a tuple with the PromptTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptTokens

`func (o *UsageLogResponse) SetPromptTokens(v int32)`

SetPromptTokens sets PromptTokens field to given value.


### SetPromptTokensNil

`func (o *UsageLogResponse) SetPromptTokensNil(b bool)`

 SetPromptTokensNil sets the value for PromptTokens to be an explicit nil

### UnsetPromptTokens
`func (o *UsageLogResponse) UnsetPromptTokens()`

UnsetPromptTokens ensures that no value is present for PromptTokens, not even an explicit nil
### GetProvider

`func (o *UsageLogResponse) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *UsageLogResponse) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *UsageLogResponse) SetProvider(v string)`

SetProvider sets Provider field to given value.


### SetProviderNil

`func (o *UsageLogResponse) SetProviderNil(b bool)`

 SetProviderNil sets the value for Provider to be an explicit nil

### UnsetProvider
`func (o *UsageLogResponse) UnsetProvider()`

UnsetProvider ensures that no value is present for Provider, not even an explicit nil
### GetStatus

`func (o *UsageLogResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UsageLogResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UsageLogResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTimestamp

`func (o *UsageLogResponse) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *UsageLogResponse) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *UsageLogResponse) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.


### GetTotalTokens

`func (o *UsageLogResponse) GetTotalTokens() int32`

GetTotalTokens returns the TotalTokens field if non-nil, zero value otherwise.

### GetTotalTokensOk

`func (o *UsageLogResponse) GetTotalTokensOk() (*int32, bool)`

GetTotalTokensOk returns a tuple with the TotalTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTokens

`func (o *UsageLogResponse) SetTotalTokens(v int32)`

SetTotalTokens sets TotalTokens field to given value.


### SetTotalTokensNil

`func (o *UsageLogResponse) SetTotalTokensNil(b bool)`

 SetTotalTokensNil sets the value for TotalTokens to be an explicit nil

### UnsetTotalTokens
`func (o *UsageLogResponse) UnsetTotalTokens()`

UnsetTotalTokens ensures that no value is present for TotalTokens, not even an explicit nil
### GetUserId

`func (o *UsageLogResponse) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UsageLogResponse) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UsageLogResponse) SetUserId(v string)`

SetUserId sets UserId field to given value.


### SetUserIdNil

`func (o *UsageLogResponse) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *UsageLogResponse) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


