# MRBetaContainer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**ExpiresAt** | **time.Time** |  | 
**Skills** | Pointer to [**[]MRBetaSkill**](MRBetaSkill.md) |  | [optional] 

## Methods

### NewMRBetaContainer

`func NewMRBetaContainer(id string, expiresAt time.Time, ) *MRBetaContainer`

NewMRBetaContainer instantiates a new MRBetaContainer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMRBetaContainerWithDefaults

`func NewMRBetaContainerWithDefaults() *MRBetaContainer`

NewMRBetaContainerWithDefaults instantiates a new MRBetaContainer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MRBetaContainer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MRBetaContainer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MRBetaContainer) SetId(v string)`

SetId sets Id field to given value.


### GetExpiresAt

`func (o *MRBetaContainer) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *MRBetaContainer) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *MRBetaContainer) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.


### GetSkills

`func (o *MRBetaContainer) GetSkills() []MRBetaSkill`

GetSkills returns the Skills field if non-nil, zero value otherwise.

### GetSkillsOk

`func (o *MRBetaContainer) GetSkillsOk() (*[]MRBetaSkill, bool)`

GetSkillsOk returns a tuple with the Skills field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkills

`func (o *MRBetaContainer) SetSkills(v []MRBetaSkill)`

SetSkills sets Skills field to given value.

### HasSkills

`func (o *MRBetaContainer) HasSkills() bool`

HasSkills returns a boolean if a field has been set.

### SetSkillsNil

`func (o *MRBetaContainer) SetSkillsNil(b bool)`

 SetSkillsNil sets the value for Skills to be an explicit nil

### UnsetSkills
`func (o *MRBetaContainer) UnsetSkills()`

UnsetSkills ensures that no value is present for Skills, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


