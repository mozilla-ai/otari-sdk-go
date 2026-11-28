# ContentShareCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Audience** | **string** |  | 
**Note** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewContentShareCreate

`func NewContentShareCreate(audience string, ) *ContentShareCreate`

NewContentShareCreate instantiates a new ContentShareCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentShareCreateWithDefaults

`func NewContentShareCreateWithDefaults() *ContentShareCreate`

NewContentShareCreateWithDefaults instantiates a new ContentShareCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudience

`func (o *ContentShareCreate) GetAudience() string`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *ContentShareCreate) GetAudienceOk() (*string, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *ContentShareCreate) SetAudience(v string)`

SetAudience sets Audience field to given value.


### GetNote

`func (o *ContentShareCreate) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *ContentShareCreate) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *ContentShareCreate) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *ContentShareCreate) HasNote() bool`

HasNote returns a boolean if a field has been set.

### SetNoteNil

`func (o *ContentShareCreate) SetNoteNil(b bool)`

 SetNoteNil sets the value for Note to be an explicit nil

### UnsetNote
`func (o *ContentShareCreate) UnsetNote()`

UnsetNote ensures that no value is present for Note, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


