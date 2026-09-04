# RankRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Examples** | [**[]ScoredExample**](ScoredExample.md) | The scored prompts to record. | 
**UserId** | **string** | Whose routing memory these examples belong to. | 
**WorkspaceId** | Pointer to **NullableString** | Which workspace&#39;s routing memory these examples belong to. Omit for the deployment&#39;s default workspace. Only requests billing to that workspace vote over them. | [optional] 

## Methods

### NewRankRequest

`func NewRankRequest(examples []ScoredExample, userId string, ) *RankRequest`

NewRankRequest instantiates a new RankRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRankRequestWithDefaults

`func NewRankRequestWithDefaults() *RankRequest`

NewRankRequestWithDefaults instantiates a new RankRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExamples

`func (o *RankRequest) GetExamples() []ScoredExample`

GetExamples returns the Examples field if non-nil, zero value otherwise.

### GetExamplesOk

`func (o *RankRequest) GetExamplesOk() (*[]ScoredExample, bool)`

GetExamplesOk returns a tuple with the Examples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExamples

`func (o *RankRequest) SetExamples(v []ScoredExample)`

SetExamples sets Examples field to given value.


### GetUserId

`func (o *RankRequest) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *RankRequest) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *RankRequest) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetWorkspaceId

`func (o *RankRequest) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *RankRequest) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *RankRequest) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.

### HasWorkspaceId

`func (o *RankRequest) HasWorkspaceId() bool`

HasWorkspaceId returns a boolean if a field has been set.

### SetWorkspaceIdNil

`func (o *RankRequest) SetWorkspaceIdNil(b bool)`

 SetWorkspaceIdNil sets the value for WorkspaceId to be an explicit nil

### UnsetWorkspaceId
`func (o *RankRequest) UnsetWorkspaceId()`

UnsetWorkspaceId ensures that no value is present for WorkspaceId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


