# Annotations

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Audience** | Pointer to **[]string** |  | [optional] 
**Priority** | Pointer to **NullableFloat32** |  | [optional] 

## Methods

### NewAnnotations

`func NewAnnotations() *Annotations`

NewAnnotations instantiates a new Annotations object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnnotationsWithDefaults

`func NewAnnotationsWithDefaults() *Annotations`

NewAnnotationsWithDefaults instantiates a new Annotations object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudience

`func (o *Annotations) GetAudience() []string`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *Annotations) GetAudienceOk() (*[]string, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *Annotations) SetAudience(v []string)`

SetAudience sets Audience field to given value.

### HasAudience

`func (o *Annotations) HasAudience() bool`

HasAudience returns a boolean if a field has been set.

### SetAudienceNil

`func (o *Annotations) SetAudienceNil(b bool)`

 SetAudienceNil sets the value for Audience to be an explicit nil

### UnsetAudience
`func (o *Annotations) UnsetAudience()`

UnsetAudience ensures that no value is present for Audience, not even an explicit nil
### GetPriority

`func (o *Annotations) GetPriority() float32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *Annotations) GetPriorityOk() (*float32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *Annotations) SetPriority(v float32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *Annotations) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

### SetPriorityNil

`func (o *Annotations) SetPriorityNil(b bool)`

 SetPriorityNil sets the value for Priority to be an explicit nil

### UnsetPriority
`func (o *Annotations) UnsetPriority()`

UnsetPriority ensures that no value is present for Priority, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


