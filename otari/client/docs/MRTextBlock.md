# MRTextBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Citations** | Pointer to [**[]MRTextBlockCitationsInner**](MRTextBlockCitationsInner.md) |  | [optional] 
**Text** | **string** |  | 
**Type** | **string** |  | 

## Methods

### NewMRTextBlock

`func NewMRTextBlock(text string, type_ string, ) *MRTextBlock`

NewMRTextBlock instantiates a new MRTextBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMRTextBlockWithDefaults

`func NewMRTextBlockWithDefaults() *MRTextBlock`

NewMRTextBlockWithDefaults instantiates a new MRTextBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCitations

`func (o *MRTextBlock) GetCitations() []MRTextBlockCitationsInner`

GetCitations returns the Citations field if non-nil, zero value otherwise.

### GetCitationsOk

`func (o *MRTextBlock) GetCitationsOk() (*[]MRTextBlockCitationsInner, bool)`

GetCitationsOk returns a tuple with the Citations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitations

`func (o *MRTextBlock) SetCitations(v []MRTextBlockCitationsInner)`

SetCitations sets Citations field to given value.

### HasCitations

`func (o *MRTextBlock) HasCitations() bool`

HasCitations returns a boolean if a field has been set.

### SetCitationsNil

`func (o *MRTextBlock) SetCitationsNil(b bool)`

 SetCitationsNil sets the value for Citations to be an explicit nil

### UnsetCitations
`func (o *MRTextBlock) UnsetCitations()`

UnsetCitations ensures that no value is present for Citations, not even an explicit nil
### GetText

`func (o *MRTextBlock) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *MRTextBlock) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *MRTextBlock) SetText(v string)`

SetText sets Text field to given value.


### GetType

`func (o *MRTextBlock) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MRTextBlock) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MRTextBlock) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


