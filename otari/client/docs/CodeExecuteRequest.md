# CodeExecuteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | Python source to run in the sandbox | 
**Container** | Pointer to **NullableString** | &#39;auto&#39; to hold the sandbox after the run, or a container id an earlier response named in Otari-Container-Id to run in that sandbox again, with its files and interpreter state. Omitted, the sandbox is released when the run ends. | [optional] 
**Environment** | Pointer to **NullableString** | The sandbox environment to run in: a name GET /api/v1/code/environments lists. Omitted, the workspace&#39;s pinned environment, else the deployment&#39;s default. A held container keeps the environment it was created in. | [optional] 
**FileIds** | Pointer to **[]string** | Uploads from POST /api/v1/files to put in the working directory before the code runs, each under its own filename. | [optional] 
**TimeoutSeconds** | Pointer to **NullableInt32** | How long the code may run. Never more than the workspace&#39;s code execution ceiling, or 60 seconds when it sets none. | [optional] 
**User** | Pointer to **NullableString** | User ID for usage attribution | [optional] 

## Methods

### NewCodeExecuteRequest

`func NewCodeExecuteRequest(code string, ) *CodeExecuteRequest`

NewCodeExecuteRequest instantiates a new CodeExecuteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCodeExecuteRequestWithDefaults

`func NewCodeExecuteRequestWithDefaults() *CodeExecuteRequest`

NewCodeExecuteRequestWithDefaults instantiates a new CodeExecuteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *CodeExecuteRequest) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CodeExecuteRequest) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CodeExecuteRequest) SetCode(v string)`

SetCode sets Code field to given value.


### GetContainer

`func (o *CodeExecuteRequest) GetContainer() string`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *CodeExecuteRequest) GetContainerOk() (*string, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *CodeExecuteRequest) SetContainer(v string)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *CodeExecuteRequest) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### SetContainerNil

`func (o *CodeExecuteRequest) SetContainerNil(b bool)`

 SetContainerNil sets the value for Container to be an explicit nil

### UnsetContainer
`func (o *CodeExecuteRequest) UnsetContainer()`

UnsetContainer ensures that no value is present for Container, not even an explicit nil
### GetEnvironment

`func (o *CodeExecuteRequest) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *CodeExecuteRequest) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *CodeExecuteRequest) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *CodeExecuteRequest) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### SetEnvironmentNil

`func (o *CodeExecuteRequest) SetEnvironmentNil(b bool)`

 SetEnvironmentNil sets the value for Environment to be an explicit nil

### UnsetEnvironment
`func (o *CodeExecuteRequest) UnsetEnvironment()`

UnsetEnvironment ensures that no value is present for Environment, not even an explicit nil
### GetFileIds

`func (o *CodeExecuteRequest) GetFileIds() []string`

GetFileIds returns the FileIds field if non-nil, zero value otherwise.

### GetFileIdsOk

`func (o *CodeExecuteRequest) GetFileIdsOk() (*[]string, bool)`

GetFileIdsOk returns a tuple with the FileIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileIds

`func (o *CodeExecuteRequest) SetFileIds(v []string)`

SetFileIds sets FileIds field to given value.

### HasFileIds

`func (o *CodeExecuteRequest) HasFileIds() bool`

HasFileIds returns a boolean if a field has been set.

### SetFileIdsNil

`func (o *CodeExecuteRequest) SetFileIdsNil(b bool)`

 SetFileIdsNil sets the value for FileIds to be an explicit nil

### UnsetFileIds
`func (o *CodeExecuteRequest) UnsetFileIds()`

UnsetFileIds ensures that no value is present for FileIds, not even an explicit nil
### GetTimeoutSeconds

`func (o *CodeExecuteRequest) GetTimeoutSeconds() int32`

GetTimeoutSeconds returns the TimeoutSeconds field if non-nil, zero value otherwise.

### GetTimeoutSecondsOk

`func (o *CodeExecuteRequest) GetTimeoutSecondsOk() (*int32, bool)`

GetTimeoutSecondsOk returns a tuple with the TimeoutSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeoutSeconds

`func (o *CodeExecuteRequest) SetTimeoutSeconds(v int32)`

SetTimeoutSeconds sets TimeoutSeconds field to given value.

### HasTimeoutSeconds

`func (o *CodeExecuteRequest) HasTimeoutSeconds() bool`

HasTimeoutSeconds returns a boolean if a field has been set.

### SetTimeoutSecondsNil

`func (o *CodeExecuteRequest) SetTimeoutSecondsNil(b bool)`

 SetTimeoutSecondsNil sets the value for TimeoutSeconds to be an explicit nil

### UnsetTimeoutSeconds
`func (o *CodeExecuteRequest) UnsetTimeoutSeconds()`

UnsetTimeoutSeconds ensures that no value is present for TimeoutSeconds, not even an explicit nil
### GetUser

`func (o *CodeExecuteRequest) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *CodeExecuteRequest) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *CodeExecuteRequest) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *CodeExecuteRequest) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *CodeExecuteRequest) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *CodeExecuteRequest) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


