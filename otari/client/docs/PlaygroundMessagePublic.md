# PlaygroundMessagePublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | **string** |  | 
**Reasoning** | Pointer to **NullableString** |  | [optional] 
**Role** | **string** |  | 

## Methods

### NewPlaygroundMessagePublic

`func NewPlaygroundMessagePublic(content string, role string, ) *PlaygroundMessagePublic`

NewPlaygroundMessagePublic instantiates a new PlaygroundMessagePublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlaygroundMessagePublicWithDefaults

`func NewPlaygroundMessagePublicWithDefaults() *PlaygroundMessagePublic`

NewPlaygroundMessagePublicWithDefaults instantiates a new PlaygroundMessagePublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *PlaygroundMessagePublic) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *PlaygroundMessagePublic) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *PlaygroundMessagePublic) SetContent(v string)`

SetContent sets Content field to given value.


### GetReasoning

`func (o *PlaygroundMessagePublic) GetReasoning() string`

GetReasoning returns the Reasoning field if non-nil, zero value otherwise.

### GetReasoningOk

`func (o *PlaygroundMessagePublic) GetReasoningOk() (*string, bool)`

GetReasoningOk returns a tuple with the Reasoning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasoning

`func (o *PlaygroundMessagePublic) SetReasoning(v string)`

SetReasoning sets Reasoning field to given value.

### HasReasoning

`func (o *PlaygroundMessagePublic) HasReasoning() bool`

HasReasoning returns a boolean if a field has been set.

### SetReasoningNil

`func (o *PlaygroundMessagePublic) SetReasoningNil(b bool)`

 SetReasoningNil sets the value for Reasoning to be an explicit nil

### UnsetReasoning
`func (o *PlaygroundMessagePublic) UnsetReasoning()`

UnsetReasoning ensures that no value is present for Reasoning, not even an explicit nil
### GetRole

`func (o *PlaygroundMessagePublic) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *PlaygroundMessagePublic) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *PlaygroundMessagePublic) SetRole(v string)`

SetRole sets Role field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


