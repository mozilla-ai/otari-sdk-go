# WorkspaceMcpServerGrantPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKeyId** | **NullableString** | The key a key grant names; null on a user&#39;s grant | 
**CreatedAt** | **string** |  | 
**Id** | **string** |  | 
**KeyActive** | Pointer to **NullableBool** | Whether a key grant&#39;s key is active | [optional] 
**KeyName** | Pointer to **NullableString** |  | [optional] 
**Subject** | **string** |  | 
**Tools** | **[]string** | Tool patterns: names, &#39;*&#39;, &#39;prefix*&#39; or &#39;@read&#39; for the server&#39;s read_tools; [] denies | 
**UpdatedAt** | **string** |  | 
**UserId** | **NullableString** | The user a grant names, or the owner of a key grant&#39;s key | 

## Methods

### NewWorkspaceMcpServerGrantPublic

`func NewWorkspaceMcpServerGrantPublic(apiKeyId NullableString, createdAt string, id string, subject string, tools []string, updatedAt string, userId NullableString, ) *WorkspaceMcpServerGrantPublic`

NewWorkspaceMcpServerGrantPublic instantiates a new WorkspaceMcpServerGrantPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkspaceMcpServerGrantPublicWithDefaults

`func NewWorkspaceMcpServerGrantPublicWithDefaults() *WorkspaceMcpServerGrantPublic`

NewWorkspaceMcpServerGrantPublicWithDefaults instantiates a new WorkspaceMcpServerGrantPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKeyId

`func (o *WorkspaceMcpServerGrantPublic) GetApiKeyId() string`

GetApiKeyId returns the ApiKeyId field if non-nil, zero value otherwise.

### GetApiKeyIdOk

`func (o *WorkspaceMcpServerGrantPublic) GetApiKeyIdOk() (*string, bool)`

GetApiKeyIdOk returns a tuple with the ApiKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKeyId

`func (o *WorkspaceMcpServerGrantPublic) SetApiKeyId(v string)`

SetApiKeyId sets ApiKeyId field to given value.


### SetApiKeyIdNil

`func (o *WorkspaceMcpServerGrantPublic) SetApiKeyIdNil(b bool)`

 SetApiKeyIdNil sets the value for ApiKeyId to be an explicit nil

### UnsetApiKeyId
`func (o *WorkspaceMcpServerGrantPublic) UnsetApiKeyId()`

UnsetApiKeyId ensures that no value is present for ApiKeyId, not even an explicit nil
### GetCreatedAt

`func (o *WorkspaceMcpServerGrantPublic) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *WorkspaceMcpServerGrantPublic) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *WorkspaceMcpServerGrantPublic) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetId

`func (o *WorkspaceMcpServerGrantPublic) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkspaceMcpServerGrantPublic) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkspaceMcpServerGrantPublic) SetId(v string)`

SetId sets Id field to given value.


### GetKeyActive

`func (o *WorkspaceMcpServerGrantPublic) GetKeyActive() bool`

GetKeyActive returns the KeyActive field if non-nil, zero value otherwise.

### GetKeyActiveOk

`func (o *WorkspaceMcpServerGrantPublic) GetKeyActiveOk() (*bool, bool)`

GetKeyActiveOk returns a tuple with the KeyActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyActive

`func (o *WorkspaceMcpServerGrantPublic) SetKeyActive(v bool)`

SetKeyActive sets KeyActive field to given value.

### HasKeyActive

`func (o *WorkspaceMcpServerGrantPublic) HasKeyActive() bool`

HasKeyActive returns a boolean if a field has been set.

### SetKeyActiveNil

`func (o *WorkspaceMcpServerGrantPublic) SetKeyActiveNil(b bool)`

 SetKeyActiveNil sets the value for KeyActive to be an explicit nil

### UnsetKeyActive
`func (o *WorkspaceMcpServerGrantPublic) UnsetKeyActive()`

UnsetKeyActive ensures that no value is present for KeyActive, not even an explicit nil
### GetKeyName

`func (o *WorkspaceMcpServerGrantPublic) GetKeyName() string`

GetKeyName returns the KeyName field if non-nil, zero value otherwise.

### GetKeyNameOk

`func (o *WorkspaceMcpServerGrantPublic) GetKeyNameOk() (*string, bool)`

GetKeyNameOk returns a tuple with the KeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyName

`func (o *WorkspaceMcpServerGrantPublic) SetKeyName(v string)`

SetKeyName sets KeyName field to given value.

### HasKeyName

`func (o *WorkspaceMcpServerGrantPublic) HasKeyName() bool`

HasKeyName returns a boolean if a field has been set.

### SetKeyNameNil

`func (o *WorkspaceMcpServerGrantPublic) SetKeyNameNil(b bool)`

 SetKeyNameNil sets the value for KeyName to be an explicit nil

### UnsetKeyName
`func (o *WorkspaceMcpServerGrantPublic) UnsetKeyName()`

UnsetKeyName ensures that no value is present for KeyName, not even an explicit nil
### GetSubject

`func (o *WorkspaceMcpServerGrantPublic) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *WorkspaceMcpServerGrantPublic) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *WorkspaceMcpServerGrantPublic) SetSubject(v string)`

SetSubject sets Subject field to given value.


### GetTools

`func (o *WorkspaceMcpServerGrantPublic) GetTools() []string`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *WorkspaceMcpServerGrantPublic) GetToolsOk() (*[]string, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *WorkspaceMcpServerGrantPublic) SetTools(v []string)`

SetTools sets Tools field to given value.


### GetUpdatedAt

`func (o *WorkspaceMcpServerGrantPublic) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *WorkspaceMcpServerGrantPublic) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *WorkspaceMcpServerGrantPublic) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetUserId

`func (o *WorkspaceMcpServerGrantPublic) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *WorkspaceMcpServerGrantPublic) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *WorkspaceMcpServerGrantPublic) SetUserId(v string)`

SetUserId sets UserId field to given value.


### SetUserIdNil

`func (o *WorkspaceMcpServerGrantPublic) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *WorkspaceMcpServerGrantPublic) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


