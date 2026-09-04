# InFlightEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKeyId** | **NullableString** |  | 
**ElapsedMs** | **int32** |  | 
**Endpoint** | **string** |  | 
**Id** | **string** |  | 
**Model** | **string** |  | 
**PolicyName** | **NullableString** |  | 
**Provider** | **NullableString** |  | 
**StartedAt** | **time.Time** |  | 
**UserId** | **NullableString** |  | 

## Methods

### NewInFlightEntry

`func NewInFlightEntry(apiKeyId NullableString, elapsedMs int32, endpoint string, id string, model string, policyName NullableString, provider NullableString, startedAt time.Time, userId NullableString, ) *InFlightEntry`

NewInFlightEntry instantiates a new InFlightEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInFlightEntryWithDefaults

`func NewInFlightEntryWithDefaults() *InFlightEntry`

NewInFlightEntryWithDefaults instantiates a new InFlightEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKeyId

`func (o *InFlightEntry) GetApiKeyId() string`

GetApiKeyId returns the ApiKeyId field if non-nil, zero value otherwise.

### GetApiKeyIdOk

`func (o *InFlightEntry) GetApiKeyIdOk() (*string, bool)`

GetApiKeyIdOk returns a tuple with the ApiKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKeyId

`func (o *InFlightEntry) SetApiKeyId(v string)`

SetApiKeyId sets ApiKeyId field to given value.


### SetApiKeyIdNil

`func (o *InFlightEntry) SetApiKeyIdNil(b bool)`

 SetApiKeyIdNil sets the value for ApiKeyId to be an explicit nil

### UnsetApiKeyId
`func (o *InFlightEntry) UnsetApiKeyId()`

UnsetApiKeyId ensures that no value is present for ApiKeyId, not even an explicit nil
### GetElapsedMs

`func (o *InFlightEntry) GetElapsedMs() int32`

GetElapsedMs returns the ElapsedMs field if non-nil, zero value otherwise.

### GetElapsedMsOk

`func (o *InFlightEntry) GetElapsedMsOk() (*int32, bool)`

GetElapsedMsOk returns a tuple with the ElapsedMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElapsedMs

`func (o *InFlightEntry) SetElapsedMs(v int32)`

SetElapsedMs sets ElapsedMs field to given value.


### GetEndpoint

`func (o *InFlightEntry) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *InFlightEntry) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *InFlightEntry) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.


### GetId

`func (o *InFlightEntry) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InFlightEntry) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InFlightEntry) SetId(v string)`

SetId sets Id field to given value.


### GetModel

`func (o *InFlightEntry) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *InFlightEntry) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *InFlightEntry) SetModel(v string)`

SetModel sets Model field to given value.


### GetPolicyName

`func (o *InFlightEntry) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *InFlightEntry) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *InFlightEntry) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.


### SetPolicyNameNil

`func (o *InFlightEntry) SetPolicyNameNil(b bool)`

 SetPolicyNameNil sets the value for PolicyName to be an explicit nil

### UnsetPolicyName
`func (o *InFlightEntry) UnsetPolicyName()`

UnsetPolicyName ensures that no value is present for PolicyName, not even an explicit nil
### GetProvider

`func (o *InFlightEntry) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *InFlightEntry) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *InFlightEntry) SetProvider(v string)`

SetProvider sets Provider field to given value.


### SetProviderNil

`func (o *InFlightEntry) SetProviderNil(b bool)`

 SetProviderNil sets the value for Provider to be an explicit nil

### UnsetProvider
`func (o *InFlightEntry) UnsetProvider()`

UnsetProvider ensures that no value is present for Provider, not even an explicit nil
### GetStartedAt

`func (o *InFlightEntry) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *InFlightEntry) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *InFlightEntry) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.


### GetUserId

`func (o *InFlightEntry) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *InFlightEntry) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *InFlightEntry) SetUserId(v string)`

SetUserId sets UserId field to given value.


### SetUserIdNil

`func (o *InFlightEntry) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *InFlightEntry) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


