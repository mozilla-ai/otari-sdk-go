# PlaygroundMessageCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | **string** |  | 
**Reasoning** | Pointer to **NullableString** |  | [optional] 
**Role** | **string** |  | 

## Methods

### NewPlaygroundMessageCreate

`func NewPlaygroundMessageCreate(content string, role string, ) *PlaygroundMessageCreate`

NewPlaygroundMessageCreate instantiates a new PlaygroundMessageCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlaygroundMessageCreateWithDefaults

`func NewPlaygroundMessageCreateWithDefaults() *PlaygroundMessageCreate`

NewPlaygroundMessageCreateWithDefaults instantiates a new PlaygroundMessageCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *PlaygroundMessageCreate) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *PlaygroundMessageCreate) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *PlaygroundMessageCreate) SetContent(v string)`

SetContent sets Content field to given value.


### GetReasoning

`func (o *PlaygroundMessageCreate) GetReasoning() string`

GetReasoning returns the Reasoning field if non-nil, zero value otherwise.

### GetReasoningOk

`func (o *PlaygroundMessageCreate) GetReasoningOk() (*string, bool)`

GetReasoningOk returns a tuple with the Reasoning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasoning

`func (o *PlaygroundMessageCreate) SetReasoning(v string)`

SetReasoning sets Reasoning field to given value.

### HasReasoning

`func (o *PlaygroundMessageCreate) HasReasoning() bool`

HasReasoning returns a boolean if a field has been set.

### SetReasoningNil

`func (o *PlaygroundMessageCreate) SetReasoningNil(b bool)`

 SetReasoningNil sets the value for Reasoning to be an explicit nil

### UnsetReasoning
`func (o *PlaygroundMessageCreate) UnsetReasoning()`

UnsetReasoning ensures that no value is present for Reasoning, not even an explicit nil
### GetRole

`func (o *PlaygroundMessageCreate) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *PlaygroundMessageCreate) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *PlaygroundMessageCreate) SetRole(v string)`

SetRole sets Role field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


