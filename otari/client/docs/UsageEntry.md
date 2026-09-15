# UsageEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKeyId** | **NullableString** |  | 
**ApiKeyName** | Pointer to **NullableString** |  | [optional] 
**AttemptCount** | Pointer to **NullableInt32** |  | [optional] 
**AttemptPosition** | Pointer to **NullableInt32** |  | [optional] 
**BillingMeters** | [**NullableBillingMeters**](BillingMeters.md) |  | 
**BulkEditable** | **bool** |  | 
**CacheReadTokens** | **NullableInt32** |  | 
**CacheWrite1hTokens** | **NullableInt32** |  | 
**CacheWriteTokens** | **NullableInt32** |  | 
**CompletionTokens** | **NullableInt32** |  | 
**Cost** | **NullableFloat32** |  | 
**CountsTowardBudget** | **bool** |  | 
**Endpoint** | **string** |  | 
**ErrorMessage** | **NullableString** |  | 
**Id** | **string** |  | 
**LatencyMs** | **NullableInt32** |  | 
**Model** | **string** |  | 
**PolicyName** | Pointer to **NullableString** |  | [optional] 
**PricingBreakdown** | [**[]UsageEntryPricingBreakdownInner**](UsageEntryPricingBreakdownInner.md) |  | 
**PromptTokens** | **NullableInt32** |  | 
**Provider** | **NullableString** |  | 
**RequestGroupId** | Pointer to **NullableString** |  | [optional] 
**SelectionReason** | Pointer to **NullableString** |  | [optional] 
**Source** | **string** |  | 
**SourceLabel** | **NullableString** |  | 
**Status** | **string** |  | 
**StatusCode** | **NullableInt32** |  | 
**Timestamp** | **string** |  | 
**TotalTokens** | **NullableInt32** |  | 
**UserAlias** | Pointer to **NullableString** |  | [optional] 
**UserId** | **NullableString** |  | 

## Methods

### NewUsageEntry

`func NewUsageEntry(apiKeyId NullableString, billingMeters NullableBillingMeters, bulkEditable bool, cacheReadTokens NullableInt32, cacheWrite1hTokens NullableInt32, cacheWriteTokens NullableInt32, completionTokens NullableInt32, cost NullableFloat32, countsTowardBudget bool, endpoint string, errorMessage NullableString, id string, latencyMs NullableInt32, model string, pricingBreakdown []UsageEntryPricingBreakdownInner, promptTokens NullableInt32, provider NullableString, source string, sourceLabel NullableString, status string, statusCode NullableInt32, timestamp string, totalTokens NullableInt32, userId NullableString, ) *UsageEntry`

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
### GetApiKeyName

`func (o *UsageEntry) GetApiKeyName() string`

GetApiKeyName returns the ApiKeyName field if non-nil, zero value otherwise.

### GetApiKeyNameOk

`func (o *UsageEntry) GetApiKeyNameOk() (*string, bool)`

GetApiKeyNameOk returns a tuple with the ApiKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKeyName

`func (o *UsageEntry) SetApiKeyName(v string)`

SetApiKeyName sets ApiKeyName field to given value.

### HasApiKeyName

`func (o *UsageEntry) HasApiKeyName() bool`

HasApiKeyName returns a boolean if a field has been set.

### SetApiKeyNameNil

`func (o *UsageEntry) SetApiKeyNameNil(b bool)`

 SetApiKeyNameNil sets the value for ApiKeyName to be an explicit nil

### UnsetApiKeyName
`func (o *UsageEntry) UnsetApiKeyName()`

UnsetApiKeyName ensures that no value is present for ApiKeyName, not even an explicit nil
### GetAttemptCount

`func (o *UsageEntry) GetAttemptCount() int32`

GetAttemptCount returns the AttemptCount field if non-nil, zero value otherwise.

### GetAttemptCountOk

`func (o *UsageEntry) GetAttemptCountOk() (*int32, bool)`

GetAttemptCountOk returns a tuple with the AttemptCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttemptCount

`func (o *UsageEntry) SetAttemptCount(v int32)`

SetAttemptCount sets AttemptCount field to given value.

### HasAttemptCount

`func (o *UsageEntry) HasAttemptCount() bool`

HasAttemptCount returns a boolean if a field has been set.

### SetAttemptCountNil

`func (o *UsageEntry) SetAttemptCountNil(b bool)`

 SetAttemptCountNil sets the value for AttemptCount to be an explicit nil

### UnsetAttemptCount
`func (o *UsageEntry) UnsetAttemptCount()`

UnsetAttemptCount ensures that no value is present for AttemptCount, not even an explicit nil
### GetAttemptPosition

`func (o *UsageEntry) GetAttemptPosition() int32`

GetAttemptPosition returns the AttemptPosition field if non-nil, zero value otherwise.

### GetAttemptPositionOk

`func (o *UsageEntry) GetAttemptPositionOk() (*int32, bool)`

GetAttemptPositionOk returns a tuple with the AttemptPosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttemptPosition

`func (o *UsageEntry) SetAttemptPosition(v int32)`

SetAttemptPosition sets AttemptPosition field to given value.

### HasAttemptPosition

`func (o *UsageEntry) HasAttemptPosition() bool`

HasAttemptPosition returns a boolean if a field has been set.

### SetAttemptPositionNil

`func (o *UsageEntry) SetAttemptPositionNil(b bool)`

 SetAttemptPositionNil sets the value for AttemptPosition to be an explicit nil

### UnsetAttemptPosition
`func (o *UsageEntry) UnsetAttemptPosition()`

UnsetAttemptPosition ensures that no value is present for AttemptPosition, not even an explicit nil
### GetBillingMeters

`func (o *UsageEntry) GetBillingMeters() BillingMeters`

GetBillingMeters returns the BillingMeters field if non-nil, zero value otherwise.

### GetBillingMetersOk

`func (o *UsageEntry) GetBillingMetersOk() (*BillingMeters, bool)`

GetBillingMetersOk returns a tuple with the BillingMeters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingMeters

`func (o *UsageEntry) SetBillingMeters(v BillingMeters)`

SetBillingMeters sets BillingMeters field to given value.


### SetBillingMetersNil

`func (o *UsageEntry) SetBillingMetersNil(b bool)`

 SetBillingMetersNil sets the value for BillingMeters to be an explicit nil

### UnsetBillingMeters
`func (o *UsageEntry) UnsetBillingMeters()`

UnsetBillingMeters ensures that no value is present for BillingMeters, not even an explicit nil
### GetBulkEditable

`func (o *UsageEntry) GetBulkEditable() bool`

GetBulkEditable returns the BulkEditable field if non-nil, zero value otherwise.

### GetBulkEditableOk

`func (o *UsageEntry) GetBulkEditableOk() (*bool, bool)`

GetBulkEditableOk returns a tuple with the BulkEditable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulkEditable

`func (o *UsageEntry) SetBulkEditable(v bool)`

SetBulkEditable sets BulkEditable field to given value.


### GetCacheReadTokens

`func (o *UsageEntry) GetCacheReadTokens() int32`

GetCacheReadTokens returns the CacheReadTokens field if non-nil, zero value otherwise.

### GetCacheReadTokensOk

`func (o *UsageEntry) GetCacheReadTokensOk() (*int32, bool)`

GetCacheReadTokensOk returns a tuple with the CacheReadTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheReadTokens

`func (o *UsageEntry) SetCacheReadTokens(v int32)`

SetCacheReadTokens sets CacheReadTokens field to given value.


### SetCacheReadTokensNil

`func (o *UsageEntry) SetCacheReadTokensNil(b bool)`

 SetCacheReadTokensNil sets the value for CacheReadTokens to be an explicit nil

### UnsetCacheReadTokens
`func (o *UsageEntry) UnsetCacheReadTokens()`

UnsetCacheReadTokens ensures that no value is present for CacheReadTokens, not even an explicit nil
### GetCacheWrite1hTokens

`func (o *UsageEntry) GetCacheWrite1hTokens() int32`

GetCacheWrite1hTokens returns the CacheWrite1hTokens field if non-nil, zero value otherwise.

### GetCacheWrite1hTokensOk

`func (o *UsageEntry) GetCacheWrite1hTokensOk() (*int32, bool)`

GetCacheWrite1hTokensOk returns a tuple with the CacheWrite1hTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheWrite1hTokens

`func (o *UsageEntry) SetCacheWrite1hTokens(v int32)`

SetCacheWrite1hTokens sets CacheWrite1hTokens field to given value.


### SetCacheWrite1hTokensNil

`func (o *UsageEntry) SetCacheWrite1hTokensNil(b bool)`

 SetCacheWrite1hTokensNil sets the value for CacheWrite1hTokens to be an explicit nil

### UnsetCacheWrite1hTokens
`func (o *UsageEntry) UnsetCacheWrite1hTokens()`

UnsetCacheWrite1hTokens ensures that no value is present for CacheWrite1hTokens, not even an explicit nil
### GetCacheWriteTokens

`func (o *UsageEntry) GetCacheWriteTokens() int32`

GetCacheWriteTokens returns the CacheWriteTokens field if non-nil, zero value otherwise.

### GetCacheWriteTokensOk

`func (o *UsageEntry) GetCacheWriteTokensOk() (*int32, bool)`

GetCacheWriteTokensOk returns a tuple with the CacheWriteTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCacheWriteTokens

`func (o *UsageEntry) SetCacheWriteTokens(v int32)`

SetCacheWriteTokens sets CacheWriteTokens field to given value.


### SetCacheWriteTokensNil

`func (o *UsageEntry) SetCacheWriteTokensNil(b bool)`

 SetCacheWriteTokensNil sets the value for CacheWriteTokens to be an explicit nil

### UnsetCacheWriteTokens
`func (o *UsageEntry) UnsetCacheWriteTokens()`

UnsetCacheWriteTokens ensures that no value is present for CacheWriteTokens, not even an explicit nil
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
### GetCountsTowardBudget

`func (o *UsageEntry) GetCountsTowardBudget() bool`

GetCountsTowardBudget returns the CountsTowardBudget field if non-nil, zero value otherwise.

### GetCountsTowardBudgetOk

`func (o *UsageEntry) GetCountsTowardBudgetOk() (*bool, bool)`

GetCountsTowardBudgetOk returns a tuple with the CountsTowardBudget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountsTowardBudget

`func (o *UsageEntry) SetCountsTowardBudget(v bool)`

SetCountsTowardBudget sets CountsTowardBudget field to given value.


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


### GetLatencyMs

`func (o *UsageEntry) GetLatencyMs() int32`

GetLatencyMs returns the LatencyMs field if non-nil, zero value otherwise.

### GetLatencyMsOk

`func (o *UsageEntry) GetLatencyMsOk() (*int32, bool)`

GetLatencyMsOk returns a tuple with the LatencyMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatencyMs

`func (o *UsageEntry) SetLatencyMs(v int32)`

SetLatencyMs sets LatencyMs field to given value.


### SetLatencyMsNil

`func (o *UsageEntry) SetLatencyMsNil(b bool)`

 SetLatencyMsNil sets the value for LatencyMs to be an explicit nil

### UnsetLatencyMs
`func (o *UsageEntry) UnsetLatencyMs()`

UnsetLatencyMs ensures that no value is present for LatencyMs, not even an explicit nil
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


### GetPolicyName

`func (o *UsageEntry) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *UsageEntry) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *UsageEntry) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.

### HasPolicyName

`func (o *UsageEntry) HasPolicyName() bool`

HasPolicyName returns a boolean if a field has been set.

### SetPolicyNameNil

`func (o *UsageEntry) SetPolicyNameNil(b bool)`

 SetPolicyNameNil sets the value for PolicyName to be an explicit nil

### UnsetPolicyName
`func (o *UsageEntry) UnsetPolicyName()`

UnsetPolicyName ensures that no value is present for PolicyName, not even an explicit nil
### GetPricingBreakdown

`func (o *UsageEntry) GetPricingBreakdown() []UsageEntryPricingBreakdownInner`

GetPricingBreakdown returns the PricingBreakdown field if non-nil, zero value otherwise.

### GetPricingBreakdownOk

`func (o *UsageEntry) GetPricingBreakdownOk() (*[]UsageEntryPricingBreakdownInner, bool)`

GetPricingBreakdownOk returns a tuple with the PricingBreakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingBreakdown

`func (o *UsageEntry) SetPricingBreakdown(v []UsageEntryPricingBreakdownInner)`

SetPricingBreakdown sets PricingBreakdown field to given value.


### SetPricingBreakdownNil

`func (o *UsageEntry) SetPricingBreakdownNil(b bool)`

 SetPricingBreakdownNil sets the value for PricingBreakdown to be an explicit nil

### UnsetPricingBreakdown
`func (o *UsageEntry) UnsetPricingBreakdown()`

UnsetPricingBreakdown ensures that no value is present for PricingBreakdown, not even an explicit nil
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
### GetRequestGroupId

`func (o *UsageEntry) GetRequestGroupId() string`

GetRequestGroupId returns the RequestGroupId field if non-nil, zero value otherwise.

### GetRequestGroupIdOk

`func (o *UsageEntry) GetRequestGroupIdOk() (*string, bool)`

GetRequestGroupIdOk returns a tuple with the RequestGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestGroupId

`func (o *UsageEntry) SetRequestGroupId(v string)`

SetRequestGroupId sets RequestGroupId field to given value.

### HasRequestGroupId

`func (o *UsageEntry) HasRequestGroupId() bool`

HasRequestGroupId returns a boolean if a field has been set.

### SetRequestGroupIdNil

`func (o *UsageEntry) SetRequestGroupIdNil(b bool)`

 SetRequestGroupIdNil sets the value for RequestGroupId to be an explicit nil

### UnsetRequestGroupId
`func (o *UsageEntry) UnsetRequestGroupId()`

UnsetRequestGroupId ensures that no value is present for RequestGroupId, not even an explicit nil
### GetSelectionReason

`func (o *UsageEntry) GetSelectionReason() string`

GetSelectionReason returns the SelectionReason field if non-nil, zero value otherwise.

### GetSelectionReasonOk

`func (o *UsageEntry) GetSelectionReasonOk() (*string, bool)`

GetSelectionReasonOk returns a tuple with the SelectionReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectionReason

`func (o *UsageEntry) SetSelectionReason(v string)`

SetSelectionReason sets SelectionReason field to given value.

### HasSelectionReason

`func (o *UsageEntry) HasSelectionReason() bool`

HasSelectionReason returns a boolean if a field has been set.

### SetSelectionReasonNil

`func (o *UsageEntry) SetSelectionReasonNil(b bool)`

 SetSelectionReasonNil sets the value for SelectionReason to be an explicit nil

### UnsetSelectionReason
`func (o *UsageEntry) UnsetSelectionReason()`

UnsetSelectionReason ensures that no value is present for SelectionReason, not even an explicit nil
### GetSource

`func (o *UsageEntry) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *UsageEntry) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *UsageEntry) SetSource(v string)`

SetSource sets Source field to given value.


### GetSourceLabel

`func (o *UsageEntry) GetSourceLabel() string`

GetSourceLabel returns the SourceLabel field if non-nil, zero value otherwise.

### GetSourceLabelOk

`func (o *UsageEntry) GetSourceLabelOk() (*string, bool)`

GetSourceLabelOk returns a tuple with the SourceLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceLabel

`func (o *UsageEntry) SetSourceLabel(v string)`

SetSourceLabel sets SourceLabel field to given value.


### SetSourceLabelNil

`func (o *UsageEntry) SetSourceLabelNil(b bool)`

 SetSourceLabelNil sets the value for SourceLabel to be an explicit nil

### UnsetSourceLabel
`func (o *UsageEntry) UnsetSourceLabel()`

UnsetSourceLabel ensures that no value is present for SourceLabel, not even an explicit nil
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


### GetStatusCode

`func (o *UsageEntry) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *UsageEntry) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *UsageEntry) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.


### SetStatusCodeNil

`func (o *UsageEntry) SetStatusCodeNil(b bool)`

 SetStatusCodeNil sets the value for StatusCode to be an explicit nil

### UnsetStatusCode
`func (o *UsageEntry) UnsetStatusCode()`

UnsetStatusCode ensures that no value is present for StatusCode, not even an explicit nil
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
### GetUserAlias

`func (o *UsageEntry) GetUserAlias() string`

GetUserAlias returns the UserAlias field if non-nil, zero value otherwise.

### GetUserAliasOk

`func (o *UsageEntry) GetUserAliasOk() (*string, bool)`

GetUserAliasOk returns a tuple with the UserAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAlias

`func (o *UsageEntry) SetUserAlias(v string)`

SetUserAlias sets UserAlias field to given value.

### HasUserAlias

`func (o *UsageEntry) HasUserAlias() bool`

HasUserAlias returns a boolean if a field has been set.

### SetUserAliasNil

`func (o *UsageEntry) SetUserAliasNil(b bool)`

 SetUserAliasNil sets the value for UserAlias to be an explicit nil

### UnsetUserAlias
`func (o *UsageEntry) UnsetUserAlias()`

UnsetUserAlias ensures that no value is present for UserAlias, not even an explicit nil
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


