# ProviderEndpointPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiBase** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**DefaultParams** | Pointer to **map[string]interface{}** | Tags for cost attribution, recorded on the request&#39;s usage rows and filterable in the usage API: up to 16 string pairs, keys up to 64 characters and values up to 512. A null value is ignored. LiteLLM&#39;s nested &#x60;spend_logs_metadata&#x60; object is also read, and wins over a flat key of the same name; it is never forwarded to the provider. | [optional] 
**Id** | **string** |  | 
**Last4** | Pointer to **NullableString** |  | [optional] 
**Name** | **string** |  | 
**Provider** | **string** |  | 
**UpdatedAt** | Pointer to **NullableTime** |  | [optional] 
**UserId** | Pointer to **NullableString** |  | [optional] 
**WorkspaceId** | **string** |  | 

## Methods

### NewProviderEndpointPublic

`func NewProviderEndpointPublic(apiBase string, createdAt time.Time, id string, name string, provider string, workspaceId string, ) *ProviderEndpointPublic`

NewProviderEndpointPublic instantiates a new ProviderEndpointPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderEndpointPublicWithDefaults

`func NewProviderEndpointPublicWithDefaults() *ProviderEndpointPublic`

NewProviderEndpointPublicWithDefaults instantiates a new ProviderEndpointPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiBase

`func (o *ProviderEndpointPublic) GetApiBase() string`

GetApiBase returns the ApiBase field if non-nil, zero value otherwise.

### GetApiBaseOk

`func (o *ProviderEndpointPublic) GetApiBaseOk() (*string, bool)`

GetApiBaseOk returns a tuple with the ApiBase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiBase

`func (o *ProviderEndpointPublic) SetApiBase(v string)`

SetApiBase sets ApiBase field to given value.


### GetCreatedAt

`func (o *ProviderEndpointPublic) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ProviderEndpointPublic) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ProviderEndpointPublic) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetDefaultParams

`func (o *ProviderEndpointPublic) GetDefaultParams() map[string]interface{}`

GetDefaultParams returns the DefaultParams field if non-nil, zero value otherwise.

### GetDefaultParamsOk

`func (o *ProviderEndpointPublic) GetDefaultParamsOk() (*map[string]interface{}, bool)`

GetDefaultParamsOk returns a tuple with the DefaultParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultParams

`func (o *ProviderEndpointPublic) SetDefaultParams(v map[string]interface{})`

SetDefaultParams sets DefaultParams field to given value.

### HasDefaultParams

`func (o *ProviderEndpointPublic) HasDefaultParams() bool`

HasDefaultParams returns a boolean if a field has been set.

### SetDefaultParamsNil

`func (o *ProviderEndpointPublic) SetDefaultParamsNil(b bool)`

 SetDefaultParamsNil sets the value for DefaultParams to be an explicit nil

### UnsetDefaultParams
`func (o *ProviderEndpointPublic) UnsetDefaultParams()`

UnsetDefaultParams ensures that no value is present for DefaultParams, not even an explicit nil
### GetId

`func (o *ProviderEndpointPublic) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProviderEndpointPublic) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProviderEndpointPublic) SetId(v string)`

SetId sets Id field to given value.


### GetLast4

`func (o *ProviderEndpointPublic) GetLast4() string`

GetLast4 returns the Last4 field if non-nil, zero value otherwise.

### GetLast4Ok

`func (o *ProviderEndpointPublic) GetLast4Ok() (*string, bool)`

GetLast4Ok returns a tuple with the Last4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast4

`func (o *ProviderEndpointPublic) SetLast4(v string)`

SetLast4 sets Last4 field to given value.

### HasLast4

`func (o *ProviderEndpointPublic) HasLast4() bool`

HasLast4 returns a boolean if a field has been set.

### SetLast4Nil

`func (o *ProviderEndpointPublic) SetLast4Nil(b bool)`

 SetLast4Nil sets the value for Last4 to be an explicit nil

### UnsetLast4
`func (o *ProviderEndpointPublic) UnsetLast4()`

UnsetLast4 ensures that no value is present for Last4, not even an explicit nil
### GetName

`func (o *ProviderEndpointPublic) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ProviderEndpointPublic) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ProviderEndpointPublic) SetName(v string)`

SetName sets Name field to given value.


### GetProvider

`func (o *ProviderEndpointPublic) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ProviderEndpointPublic) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ProviderEndpointPublic) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetUpdatedAt

`func (o *ProviderEndpointPublic) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ProviderEndpointPublic) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ProviderEndpointPublic) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ProviderEndpointPublic) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *ProviderEndpointPublic) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *ProviderEndpointPublic) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetUserId

`func (o *ProviderEndpointPublic) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ProviderEndpointPublic) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ProviderEndpointPublic) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ProviderEndpointPublic) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *ProviderEndpointPublic) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *ProviderEndpointPublic) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetWorkspaceId

`func (o *ProviderEndpointPublic) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *ProviderEndpointPublic) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *ProviderEndpointPublic) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


