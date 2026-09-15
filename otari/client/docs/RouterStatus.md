# RouterStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alpha** | **float32** |  | 
**ConfidenceFloor** | **float32** |  | 
**DefaultPool** | [**PoolStatus**](PoolStatus.md) |  | 
**EmbeddingModel** | **string** |  | 
**Granularity** | **string** |  | 
**K** | **int32** |  | 
**Policies** | [**[]LearnedPolicy**](LearnedPolicy.md) |  | 
**SeedCount** | **int32** |  | 
**Tasks** | [**[]TaskPool**](TaskPool.md) |  | 
**UserId** | **string** |  | 
**WorkspaceId** | **string** |  | 

## Methods

### NewRouterStatus

`func NewRouterStatus(alpha float32, confidenceFloor float32, defaultPool PoolStatus, embeddingModel string, granularity string, k int32, policies []LearnedPolicy, seedCount int32, tasks []TaskPool, userId string, workspaceId string, ) *RouterStatus`

NewRouterStatus instantiates a new RouterStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRouterStatusWithDefaults

`func NewRouterStatusWithDefaults() *RouterStatus`

NewRouterStatusWithDefaults instantiates a new RouterStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlpha

`func (o *RouterStatus) GetAlpha() float32`

GetAlpha returns the Alpha field if non-nil, zero value otherwise.

### GetAlphaOk

`func (o *RouterStatus) GetAlphaOk() (*float32, bool)`

GetAlphaOk returns a tuple with the Alpha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlpha

`func (o *RouterStatus) SetAlpha(v float32)`

SetAlpha sets Alpha field to given value.


### GetConfidenceFloor

`func (o *RouterStatus) GetConfidenceFloor() float32`

GetConfidenceFloor returns the ConfidenceFloor field if non-nil, zero value otherwise.

### GetConfidenceFloorOk

`func (o *RouterStatus) GetConfidenceFloorOk() (*float32, bool)`

GetConfidenceFloorOk returns a tuple with the ConfidenceFloor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidenceFloor

`func (o *RouterStatus) SetConfidenceFloor(v float32)`

SetConfidenceFloor sets ConfidenceFloor field to given value.


### GetDefaultPool

`func (o *RouterStatus) GetDefaultPool() PoolStatus`

GetDefaultPool returns the DefaultPool field if non-nil, zero value otherwise.

### GetDefaultPoolOk

`func (o *RouterStatus) GetDefaultPoolOk() (*PoolStatus, bool)`

GetDefaultPoolOk returns a tuple with the DefaultPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPool

`func (o *RouterStatus) SetDefaultPool(v PoolStatus)`

SetDefaultPool sets DefaultPool field to given value.


### GetEmbeddingModel

`func (o *RouterStatus) GetEmbeddingModel() string`

GetEmbeddingModel returns the EmbeddingModel field if non-nil, zero value otherwise.

### GetEmbeddingModelOk

`func (o *RouterStatus) GetEmbeddingModelOk() (*string, bool)`

GetEmbeddingModelOk returns a tuple with the EmbeddingModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbeddingModel

`func (o *RouterStatus) SetEmbeddingModel(v string)`

SetEmbeddingModel sets EmbeddingModel field to given value.


### GetGranularity

`func (o *RouterStatus) GetGranularity() string`

GetGranularity returns the Granularity field if non-nil, zero value otherwise.

### GetGranularityOk

`func (o *RouterStatus) GetGranularityOk() (*string, bool)`

GetGranularityOk returns a tuple with the Granularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGranularity

`func (o *RouterStatus) SetGranularity(v string)`

SetGranularity sets Granularity field to given value.


### GetK

`func (o *RouterStatus) GetK() int32`

GetK returns the K field if non-nil, zero value otherwise.

### GetKOk

`func (o *RouterStatus) GetKOk() (*int32, bool)`

GetKOk returns a tuple with the K field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetK

`func (o *RouterStatus) SetK(v int32)`

SetK sets K field to given value.


### GetPolicies

`func (o *RouterStatus) GetPolicies() []LearnedPolicy`

GetPolicies returns the Policies field if non-nil, zero value otherwise.

### GetPoliciesOk

`func (o *RouterStatus) GetPoliciesOk() (*[]LearnedPolicy, bool)`

GetPoliciesOk returns a tuple with the Policies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicies

`func (o *RouterStatus) SetPolicies(v []LearnedPolicy)`

SetPolicies sets Policies field to given value.


### GetSeedCount

`func (o *RouterStatus) GetSeedCount() int32`

GetSeedCount returns the SeedCount field if non-nil, zero value otherwise.

### GetSeedCountOk

`func (o *RouterStatus) GetSeedCountOk() (*int32, bool)`

GetSeedCountOk returns a tuple with the SeedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeedCount

`func (o *RouterStatus) SetSeedCount(v int32)`

SetSeedCount sets SeedCount field to given value.


### GetTasks

`func (o *RouterStatus) GetTasks() []TaskPool`

GetTasks returns the Tasks field if non-nil, zero value otherwise.

### GetTasksOk

`func (o *RouterStatus) GetTasksOk() (*[]TaskPool, bool)`

GetTasksOk returns a tuple with the Tasks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTasks

`func (o *RouterStatus) SetTasks(v []TaskPool)`

SetTasks sets Tasks field to given value.


### GetUserId

`func (o *RouterStatus) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *RouterStatus) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *RouterStatus) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetWorkspaceId

`func (o *RouterStatus) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *RouterStatus) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *RouterStatus) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


