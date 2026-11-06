# OrgWebSearchKeyUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiBase** | Pointer to **NullableString** |  | [optional] 
**ApiKey** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Options** | Pointer to **map[string]interface{}** | Tags for cost attribution, recorded on the request&#39;s usage rows and filterable in the usage API: up to 16 string pairs, keys up to 64 characters and values up to 512. A null value is ignored. LiteLLM&#39;s nested &#x60;spend_logs_metadata&#x60; object is also read, and wins over a flat key of the same name; it is never forwarded to the provider. | [optional] 

## Methods

### NewOrgWebSearchKeyUpdateRequest

`func NewOrgWebSearchKeyUpdateRequest() *OrgWebSearchKeyUpdateRequest`

NewOrgWebSearchKeyUpdateRequest instantiates a new OrgWebSearchKeyUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgWebSearchKeyUpdateRequestWithDefaults

`func NewOrgWebSearchKeyUpdateRequestWithDefaults() *OrgWebSearchKeyUpdateRequest`

NewOrgWebSearchKeyUpdateRequestWithDefaults instantiates a new OrgWebSearchKeyUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiBase

`func (o *OrgWebSearchKeyUpdateRequest) GetApiBase() string`

GetApiBase returns the ApiBase field if non-nil, zero value otherwise.

### GetApiBaseOk

`func (o *OrgWebSearchKeyUpdateRequest) GetApiBaseOk() (*string, bool)`

GetApiBaseOk returns a tuple with the ApiBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiBase

`func (o *OrgWebSearchKeyUpdateRequest) SetApiBase(v string)`

SetApiBase sets ApiBase field to given value.

### HasApiBase

`func (o *OrgWebSearchKeyUpdateRequest) HasApiBase() bool`

HasApiBase returns a boolean if a field has been set.

### SetApiBaseNil

`func (o *OrgWebSearchKeyUpdateRequest) SetApiBaseNil(b bool)`

 SetApiBaseNil sets the value for ApiBase to be an explicit nil

### UnsetApiBase
`func (o *OrgWebSearchKeyUpdateRequest) UnsetApiBase()`

UnsetApiBase ensures that no value is present for ApiBase, not even an explicit nil
### GetApiKey

`func (o *OrgWebSearchKeyUpdateRequest) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *OrgWebSearchKeyUpdateRequest) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *OrgWebSearchKeyUpdateRequest) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *OrgWebSearchKeyUpdateRequest) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### SetApiKeyNil

`func (o *OrgWebSearchKeyUpdateRequest) SetApiKeyNil(b bool)`

 SetApiKeyNil sets the value for ApiKey to be an explicit nil

### UnsetApiKey
`func (o *OrgWebSearchKeyUpdateRequest) UnsetApiKey()`

UnsetApiKey ensures that no value is present for ApiKey, not even an explicit nil
### GetName

`func (o *OrgWebSearchKeyUpdateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrgWebSearchKeyUpdateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrgWebSearchKeyUpdateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OrgWebSearchKeyUpdateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *OrgWebSearchKeyUpdateRequest) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *OrgWebSearchKeyUpdateRequest) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetOptions

`func (o *OrgWebSearchKeyUpdateRequest) GetOptions() map[string]interface{}`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *OrgWebSearchKeyUpdateRequest) GetOptionsOk() (*map[string]interface{}, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *OrgWebSearchKeyUpdateRequest) SetOptions(v map[string]interface{})`

SetOptions sets Options field to given value.

### HasOptions

`func (o *OrgWebSearchKeyUpdateRequest) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *OrgWebSearchKeyUpdateRequest) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *OrgWebSearchKeyUpdateRequest) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


