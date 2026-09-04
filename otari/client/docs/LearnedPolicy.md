# LearnedPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Backend** | **string** |  | 
**Candidates** | **[]string** |  | 
**DefaultTarget** | **string** |  | 
**Name** | **string** |  | 

## Methods

### NewLearnedPolicy

`func NewLearnedPolicy(backend string, candidates []string, defaultTarget string, name string, ) *LearnedPolicy`

NewLearnedPolicy instantiates a new LearnedPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLearnedPolicyWithDefaults

`func NewLearnedPolicyWithDefaults() *LearnedPolicy`

NewLearnedPolicyWithDefaults instantiates a new LearnedPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackend

`func (o *LearnedPolicy) GetBackend() string`

GetBackend returns the Backend field if non-nil, zero value otherwise.

### GetBackendOk

`func (o *LearnedPolicy) GetBackendOk() (*string, bool)`

GetBackendOk returns a tuple with the Backend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackend

`func (o *LearnedPolicy) SetBackend(v string)`

SetBackend sets Backend field to given value.


### GetCandidates

`func (o *LearnedPolicy) GetCandidates() []string`

GetCandidates returns the Candidates field if non-nil, zero value otherwise.

### GetCandidatesOk

`func (o *LearnedPolicy) GetCandidatesOk() (*[]string, bool)`

GetCandidatesOk returns a tuple with the Candidates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCandidates

`func (o *LearnedPolicy) SetCandidates(v []string)`

SetCandidates sets Candidates field to given value.


### GetDefaultTarget

`func (o *LearnedPolicy) GetDefaultTarget() string`

GetDefaultTarget returns the DefaultTarget field if non-nil, zero value otherwise.

### GetDefaultTargetOk

`func (o *LearnedPolicy) GetDefaultTargetOk() (*string, bool)`

GetDefaultTargetOk returns a tuple with the DefaultTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTarget

`func (o *LearnedPolicy) SetDefaultTarget(v string)`

SetDefaultTarget sets DefaultTarget field to given value.


### GetName

`func (o *LearnedPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LearnedPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LearnedPolicy) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


