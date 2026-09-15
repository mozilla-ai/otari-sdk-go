# ExplainResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Candidates** | [**[]CandidateResponse**](CandidateResponse.md) |  | 
**Dropped** | [**[]DroppedResponse**](DroppedResponse.md) |  | 
**Guardrails** | **[]map[string]interface{}** |  | 
**IsDynamic** | **bool** |  | 
**Name** | **string** |  | 
**RouterBackend** | Pointer to **NullableString** |  | [optional] 
**RouterCandidates** | Pointer to **[]string** |  | [optional] 
**RouterWeights** | Pointer to **map[string]float32** | For a weighted policy, the percentage of traffic each candidate receives, normalized over the candidates this caller may use. Empty for every other policy, and for a weighted policy whose whole split this caller may not use: a split over no candidate is not a split, and each filtered candidate is named in &#x60;dropped&#x60; instead. A weighted split needs no request state, so unlike a learned router&#39;s ranking it is knowable here: the plan above is the real ordering by share, not the decline path. | [optional] 
**SelectionReason** | **string** |  | 

## Methods

### NewExplainResponse

`func NewExplainResponse(candidates []CandidateResponse, dropped []DroppedResponse, guardrails []*map[string]interface{}, isDynamic bool, name string, selectionReason string, ) *ExplainResponse`

NewExplainResponse instantiates a new ExplainResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExplainResponseWithDefaults

`func NewExplainResponseWithDefaults() *ExplainResponse`

NewExplainResponseWithDefaults instantiates a new ExplainResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCandidates

`func (o *ExplainResponse) GetCandidates() []CandidateResponse`

GetCandidates returns the Candidates field if non-nil, zero value otherwise.

### GetCandidatesOk

`func (o *ExplainResponse) GetCandidatesOk() (*[]CandidateResponse, bool)`

GetCandidatesOk returns a tuple with the Candidates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCandidates

`func (o *ExplainResponse) SetCandidates(v []CandidateResponse)`

SetCandidates sets Candidates field to given value.


### GetDropped

`func (o *ExplainResponse) GetDropped() []DroppedResponse`

GetDropped returns the Dropped field if non-nil, zero value otherwise.

### GetDroppedOk

`func (o *ExplainResponse) GetDroppedOk() (*[]DroppedResponse, bool)`

GetDroppedOk returns a tuple with the Dropped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDropped

`func (o *ExplainResponse) SetDropped(v []DroppedResponse)`

SetDropped sets Dropped field to given value.


### GetGuardrails

`func (o *ExplainResponse) GetGuardrails() []*map[string]interface{}`

GetGuardrails returns the Guardrails field if non-nil, zero value otherwise.

### GetGuardrailsOk

`func (o *ExplainResponse) GetGuardrailsOk() (*[]*map[string]interface{}, bool)`

GetGuardrailsOk returns a tuple with the Guardrails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuardrails

`func (o *ExplainResponse) SetGuardrails(v []*map[string]interface{})`

SetGuardrails sets Guardrails field to given value.


### GetIsDynamic

`func (o *ExplainResponse) GetIsDynamic() bool`

GetIsDynamic returns the IsDynamic field if non-nil, zero value otherwise.

### GetIsDynamicOk

`func (o *ExplainResponse) GetIsDynamicOk() (*bool, bool)`

GetIsDynamicOk returns a tuple with the IsDynamic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDynamic

`func (o *ExplainResponse) SetIsDynamic(v bool)`

SetIsDynamic sets IsDynamic field to given value.


### GetName

`func (o *ExplainResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ExplainResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ExplainResponse) SetName(v string)`

SetName sets Name field to given value.


### GetRouterBackend

`func (o *ExplainResponse) GetRouterBackend() string`

GetRouterBackend returns the RouterBackend field if non-nil, zero value otherwise.

### GetRouterBackendOk

`func (o *ExplainResponse) GetRouterBackendOk() (*string, bool)`

GetRouterBackendOk returns a tuple with the RouterBackend field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouterBackend

`func (o *ExplainResponse) SetRouterBackend(v string)`

SetRouterBackend sets RouterBackend field to given value.

### HasRouterBackend

`func (o *ExplainResponse) HasRouterBackend() bool`

HasRouterBackend returns a boolean if a field has been set.

### SetRouterBackendNil

`func (o *ExplainResponse) SetRouterBackendNil(b bool)`

 SetRouterBackendNil sets the value for RouterBackend to be an explicit nil

### UnsetRouterBackend
`func (o *ExplainResponse) UnsetRouterBackend()`

UnsetRouterBackend ensures that no value is present for RouterBackend, not even an explicit nil
### GetRouterCandidates

`func (o *ExplainResponse) GetRouterCandidates() []string`

GetRouterCandidates returns the RouterCandidates field if non-nil, zero value otherwise.

### GetRouterCandidatesOk

`func (o *ExplainResponse) GetRouterCandidatesOk() (*[]string, bool)`

GetRouterCandidatesOk returns a tuple with the RouterCandidates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouterCandidates

`func (o *ExplainResponse) SetRouterCandidates(v []string)`

SetRouterCandidates sets RouterCandidates field to given value.

### HasRouterCandidates

`func (o *ExplainResponse) HasRouterCandidates() bool`

HasRouterCandidates returns a boolean if a field has been set.

### GetRouterWeights

`func (o *ExplainResponse) GetRouterWeights() map[string]float32`

GetRouterWeights returns the RouterWeights field if non-nil, zero value otherwise.

### GetRouterWeightsOk

`func (o *ExplainResponse) GetRouterWeightsOk() (*map[string]float32, bool)`

GetRouterWeightsOk returns a tuple with the RouterWeights field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouterWeights

`func (o *ExplainResponse) SetRouterWeights(v map[string]float32)`

SetRouterWeights sets RouterWeights field to given value.

### HasRouterWeights

`func (o *ExplainResponse) HasRouterWeights() bool`

HasRouterWeights returns a boolean if a field has been set.

### GetSelectionReason

`func (o *ExplainResponse) GetSelectionReason() string`

GetSelectionReason returns the SelectionReason field if non-nil, zero value otherwise.

### GetSelectionReasonOk

`func (o *ExplainResponse) GetSelectionReasonOk() (*string, bool)`

GetSelectionReasonOk returns a tuple with the SelectionReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectionReason

`func (o *ExplainResponse) SetSelectionReason(v string)`

SetSelectionReason sets SelectionReason field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


