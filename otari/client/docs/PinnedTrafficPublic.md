# PinnedTrafficPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKeyId** | **NullableString** | The key, or null for requests made without one or whose key was deleted. | 
**KeyName** | **NullableString** | The key&#39;s name, or null where it has none. | 
**Provider** | **string** | The provider instance the requests reached. | 
**Requests** | **int32** |  | 

## Methods

### NewPinnedTrafficPublic

`func NewPinnedTrafficPublic(apiKeyId NullableString, keyName NullableString, provider string, requests int32, ) *PinnedTrafficPublic`

NewPinnedTrafficPublic instantiates a new PinnedTrafficPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPinnedTrafficPublicWithDefaults

`func NewPinnedTrafficPublicWithDefaults() *PinnedTrafficPublic`

NewPinnedTrafficPublicWithDefaults instantiates a new PinnedTrafficPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKeyId

`func (o *PinnedTrafficPublic) GetApiKeyId() string`

GetApiKeyId returns the ApiKeyId field if non-nil, zero value otherwise.

### GetApiKeyIdOk

`func (o *PinnedTrafficPublic) GetApiKeyIdOk() (*string, bool)`

GetApiKeyIdOk returns a tuple with the ApiKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKeyId

`func (o *PinnedTrafficPublic) SetApiKeyId(v string)`

SetApiKeyId sets ApiKeyId field to given value.


### SetApiKeyIdNil

`func (o *PinnedTrafficPublic) SetApiKeyIdNil(b bool)`

 SetApiKeyIdNil sets the value for ApiKeyId to be an explicit nil

### UnsetApiKeyId
`func (o *PinnedTrafficPublic) UnsetApiKeyId()`

UnsetApiKeyId ensures that no value is present for ApiKeyId, not even an explicit nil
### GetKeyName

`func (o *PinnedTrafficPublic) GetKeyName() string`

GetKeyName returns the KeyName field if non-nil, zero value otherwise.

### GetKeyNameOk

`func (o *PinnedTrafficPublic) GetKeyNameOk() (*string, bool)`

GetKeyNameOk returns a tuple with the KeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyName

`func (o *PinnedTrafficPublic) SetKeyName(v string)`

SetKeyName sets KeyName field to given value.


### SetKeyNameNil

`func (o *PinnedTrafficPublic) SetKeyNameNil(b bool)`

 SetKeyNameNil sets the value for KeyName to be an explicit nil

### UnsetKeyName
`func (o *PinnedTrafficPublic) UnsetKeyName()`

UnsetKeyName ensures that no value is present for KeyName, not even an explicit nil
### GetProvider

`func (o *PinnedTrafficPublic) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *PinnedTrafficPublic) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *PinnedTrafficPublic) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetRequests

`func (o *PinnedTrafficPublic) GetRequests() int32`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *PinnedTrafficPublic) GetRequestsOk() (*int32, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *PinnedTrafficPublic) SetRequests(v int32)`

SetRequests sets Requests field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


