# KeyLimitPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LimitRequests** | Pointer to **NullableInt32** |  | [optional] 
**LimitTokens** | Pointer to **NullableInt32** |  | [optional] 
**LimitUsd** | **NullableFloat32** | Dollars allowed per period; null when uncapped or shared | 
**Provider** | **NullableString** | The provider this cap is narrowed to, or null when it caps every one | 
**RemainingRequests** | Pointer to **NullableInt32** |  | [optional] 
**RemainingTokens** | Pointer to **NullableInt32** |  | [optional] 
**RemainingUsd** | **NullableFloat32** | Dollars left this period; null when dollars are uncapped | 
**ResetsAt** | **NullableTime** | When the period ends and use starts again from zero; null for a cap that never resets | 
**Scope** | **string** | Where the cap sits: &#39;user&#39; for the budget the key&#39;s user is on, or the scope of a spend ceiling: &#39;api_token&#39;, &#39;workspace_member&#39;, &#39;org_member&#39;, &#39;workspace&#39; or &#39;organization&#39; | 
**UsedRequests** | Pointer to **NullableInt32** |  | [optional] 
**UsedTokens** | Pointer to **NullableInt32** |  | [optional] 
**UsedUsd** | **NullableFloat32** | Dollars spent and held this period; null when uncapped or shared | 

## Methods

### NewKeyLimitPublic

`func NewKeyLimitPublic(limitUsd NullableFloat32, provider NullableString, remainingUsd NullableFloat32, resetsAt NullableTime, scope string, usedUsd NullableFloat32, ) *KeyLimitPublic`

NewKeyLimitPublic instantiates a new KeyLimitPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyLimitPublicWithDefaults

`func NewKeyLimitPublicWithDefaults() *KeyLimitPublic`

NewKeyLimitPublicWithDefaults instantiates a new KeyLimitPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimitRequests

`func (o *KeyLimitPublic) GetLimitRequests() int32`

GetLimitRequests returns the LimitRequests field if non-nil, zero value otherwise.

### GetLimitRequestsOk

`func (o *KeyLimitPublic) GetLimitRequestsOk() (*int32, bool)`

GetLimitRequestsOk returns a tuple with the LimitRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitRequests

`func (o *KeyLimitPublic) SetLimitRequests(v int32)`

SetLimitRequests sets LimitRequests field to given value.

### HasLimitRequests

`func (o *KeyLimitPublic) HasLimitRequests() bool`

HasLimitRequests returns a boolean if a field has been set.

### SetLimitRequestsNil

`func (o *KeyLimitPublic) SetLimitRequestsNil(b bool)`

 SetLimitRequestsNil sets the value for LimitRequests to be an explicit nil

### UnsetLimitRequests
`func (o *KeyLimitPublic) UnsetLimitRequests()`

UnsetLimitRequests ensures that no value is present for LimitRequests, not even an explicit nil
### GetLimitTokens

`func (o *KeyLimitPublic) GetLimitTokens() int32`

GetLimitTokens returns the LimitTokens field if non-nil, zero value otherwise.

### GetLimitTokensOk

`func (o *KeyLimitPublic) GetLimitTokensOk() (*int32, bool)`

GetLimitTokensOk returns a tuple with the LimitTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitTokens

`func (o *KeyLimitPublic) SetLimitTokens(v int32)`

SetLimitTokens sets LimitTokens field to given value.

### HasLimitTokens

`func (o *KeyLimitPublic) HasLimitTokens() bool`

HasLimitTokens returns a boolean if a field has been set.

### SetLimitTokensNil

`func (o *KeyLimitPublic) SetLimitTokensNil(b bool)`

 SetLimitTokensNil sets the value for LimitTokens to be an explicit nil

### UnsetLimitTokens
`func (o *KeyLimitPublic) UnsetLimitTokens()`

UnsetLimitTokens ensures that no value is present for LimitTokens, not even an explicit nil
### GetLimitUsd

`func (o *KeyLimitPublic) GetLimitUsd() float32`

GetLimitUsd returns the LimitUsd field if non-nil, zero value otherwise.

### GetLimitUsdOk

`func (o *KeyLimitPublic) GetLimitUsdOk() (*float32, bool)`

GetLimitUsdOk returns a tuple with the LimitUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitUsd

`func (o *KeyLimitPublic) SetLimitUsd(v float32)`

SetLimitUsd sets LimitUsd field to given value.


### SetLimitUsdNil

`func (o *KeyLimitPublic) SetLimitUsdNil(b bool)`

 SetLimitUsdNil sets the value for LimitUsd to be an explicit nil

### UnsetLimitUsd
`func (o *KeyLimitPublic) UnsetLimitUsd()`

UnsetLimitUsd ensures that no value is present for LimitUsd, not even an explicit nil
### GetProvider

`func (o *KeyLimitPublic) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *KeyLimitPublic) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *KeyLimitPublic) SetProvider(v string)`

SetProvider sets Provider field to given value.


### SetProviderNil

`func (o *KeyLimitPublic) SetProviderNil(b bool)`

 SetProviderNil sets the value for Provider to be an explicit nil

### UnsetProvider
`func (o *KeyLimitPublic) UnsetProvider()`

UnsetProvider ensures that no value is present for Provider, not even an explicit nil
### GetRemainingRequests

`func (o *KeyLimitPublic) GetRemainingRequests() int32`

GetRemainingRequests returns the RemainingRequests field if non-nil, zero value otherwise.

### GetRemainingRequestsOk

`func (o *KeyLimitPublic) GetRemainingRequestsOk() (*int32, bool)`

GetRemainingRequestsOk returns a tuple with the RemainingRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingRequests

`func (o *KeyLimitPublic) SetRemainingRequests(v int32)`

SetRemainingRequests sets RemainingRequests field to given value.

### HasRemainingRequests

`func (o *KeyLimitPublic) HasRemainingRequests() bool`

HasRemainingRequests returns a boolean if a field has been set.

### SetRemainingRequestsNil

`func (o *KeyLimitPublic) SetRemainingRequestsNil(b bool)`

 SetRemainingRequestsNil sets the value for RemainingRequests to be an explicit nil

### UnsetRemainingRequests
`func (o *KeyLimitPublic) UnsetRemainingRequests()`

UnsetRemainingRequests ensures that no value is present for RemainingRequests, not even an explicit nil
### GetRemainingTokens

`func (o *KeyLimitPublic) GetRemainingTokens() int32`

GetRemainingTokens returns the RemainingTokens field if non-nil, zero value otherwise.

### GetRemainingTokensOk

`func (o *KeyLimitPublic) GetRemainingTokensOk() (*int32, bool)`

GetRemainingTokensOk returns a tuple with the RemainingTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingTokens

`func (o *KeyLimitPublic) SetRemainingTokens(v int32)`

SetRemainingTokens sets RemainingTokens field to given value.

### HasRemainingTokens

`func (o *KeyLimitPublic) HasRemainingTokens() bool`

HasRemainingTokens returns a boolean if a field has been set.

### SetRemainingTokensNil

`func (o *KeyLimitPublic) SetRemainingTokensNil(b bool)`

 SetRemainingTokensNil sets the value for RemainingTokens to be an explicit nil

### UnsetRemainingTokens
`func (o *KeyLimitPublic) UnsetRemainingTokens()`

UnsetRemainingTokens ensures that no value is present for RemainingTokens, not even an explicit nil
### GetRemainingUsd

`func (o *KeyLimitPublic) GetRemainingUsd() float32`

GetRemainingUsd returns the RemainingUsd field if non-nil, zero value otherwise.

### GetRemainingUsdOk

`func (o *KeyLimitPublic) GetRemainingUsdOk() (*float32, bool)`

GetRemainingUsdOk returns a tuple with the RemainingUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingUsd

`func (o *KeyLimitPublic) SetRemainingUsd(v float32)`

SetRemainingUsd sets RemainingUsd field to given value.


### SetRemainingUsdNil

`func (o *KeyLimitPublic) SetRemainingUsdNil(b bool)`

 SetRemainingUsdNil sets the value for RemainingUsd to be an explicit nil

### UnsetRemainingUsd
`func (o *KeyLimitPublic) UnsetRemainingUsd()`

UnsetRemainingUsd ensures that no value is present for RemainingUsd, not even an explicit nil
### GetResetsAt

`func (o *KeyLimitPublic) GetResetsAt() time.Time`

GetResetsAt returns the ResetsAt field if non-nil, zero value otherwise.

### GetResetsAtOk

`func (o *KeyLimitPublic) GetResetsAtOk() (*time.Time, bool)`

GetResetsAtOk returns a tuple with the ResetsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetsAt

`func (o *KeyLimitPublic) SetResetsAt(v time.Time)`

SetResetsAt sets ResetsAt field to given value.


### SetResetsAtNil

`func (o *KeyLimitPublic) SetResetsAtNil(b bool)`

 SetResetsAtNil sets the value for ResetsAt to be an explicit nil

### UnsetResetsAt
`func (o *KeyLimitPublic) UnsetResetsAt()`

UnsetResetsAt ensures that no value is present for ResetsAt, not even an explicit nil
### GetScope

`func (o *KeyLimitPublic) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *KeyLimitPublic) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *KeyLimitPublic) SetScope(v string)`

SetScope sets Scope field to given value.


### GetUsedRequests

`func (o *KeyLimitPublic) GetUsedRequests() int32`

GetUsedRequests returns the UsedRequests field if non-nil, zero value otherwise.

### GetUsedRequestsOk

`func (o *KeyLimitPublic) GetUsedRequestsOk() (*int32, bool)`

GetUsedRequestsOk returns a tuple with the UsedRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedRequests

`func (o *KeyLimitPublic) SetUsedRequests(v int32)`

SetUsedRequests sets UsedRequests field to given value.

### HasUsedRequests

`func (o *KeyLimitPublic) HasUsedRequests() bool`

HasUsedRequests returns a boolean if a field has been set.

### SetUsedRequestsNil

`func (o *KeyLimitPublic) SetUsedRequestsNil(b bool)`

 SetUsedRequestsNil sets the value for UsedRequests to be an explicit nil

### UnsetUsedRequests
`func (o *KeyLimitPublic) UnsetUsedRequests()`

UnsetUsedRequests ensures that no value is present for UsedRequests, not even an explicit nil
### GetUsedTokens

`func (o *KeyLimitPublic) GetUsedTokens() int32`

GetUsedTokens returns the UsedTokens field if non-nil, zero value otherwise.

### GetUsedTokensOk

`func (o *KeyLimitPublic) GetUsedTokensOk() (*int32, bool)`

GetUsedTokensOk returns a tuple with the UsedTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedTokens

`func (o *KeyLimitPublic) SetUsedTokens(v int32)`

SetUsedTokens sets UsedTokens field to given value.

### HasUsedTokens

`func (o *KeyLimitPublic) HasUsedTokens() bool`

HasUsedTokens returns a boolean if a field has been set.

### SetUsedTokensNil

`func (o *KeyLimitPublic) SetUsedTokensNil(b bool)`

 SetUsedTokensNil sets the value for UsedTokens to be an explicit nil

### UnsetUsedTokens
`func (o *KeyLimitPublic) UnsetUsedTokens()`

UnsetUsedTokens ensures that no value is present for UsedTokens, not even an explicit nil
### GetUsedUsd

`func (o *KeyLimitPublic) GetUsedUsd() float32`

GetUsedUsd returns the UsedUsd field if non-nil, zero value otherwise.

### GetUsedUsdOk

`func (o *KeyLimitPublic) GetUsedUsdOk() (*float32, bool)`

GetUsedUsdOk returns a tuple with the UsedUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedUsd

`func (o *KeyLimitPublic) SetUsedUsd(v float32)`

SetUsedUsd sets UsedUsd field to given value.


### SetUsedUsdNil

`func (o *KeyLimitPublic) SetUsedUsdNil(b bool)`

 SetUsedUsdNil sets the value for UsedUsd to be an explicit nil

### UnsetUsedUsd
`func (o *KeyLimitPublic) UnsetUsedUsd()`

UnsetUsedUsd ensures that no value is present for UsedUsd, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


