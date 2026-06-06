# MRDocumentBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Citations** | Pointer to [**NullableMRCitationsConfig**](MRCitationsConfig.md) |  | [optional] 
**Source** | [**Source**](Source.md) |  | 
**Title** | Pointer to **NullableString** |  | [optional] 
**Type** | **string** |  | 

## Methods

### NewMRDocumentBlock

`func NewMRDocumentBlock(source Source, type_ string, ) *MRDocumentBlock`

NewMRDocumentBlock instantiates a new MRDocumentBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMRDocumentBlockWithDefaults

`func NewMRDocumentBlockWithDefaults() *MRDocumentBlock`

NewMRDocumentBlockWithDefaults instantiates a new MRDocumentBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCitations

`func (o *MRDocumentBlock) GetCitations() MRCitationsConfig`

GetCitations returns the Citations field if non-nil, zero value otherwise.

### GetCitationsOk

`func (o *MRDocumentBlock) GetCitationsOk() (*MRCitationsConfig, bool)`

GetCitationsOk returns a tuple with the Citations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitations

`func (o *MRDocumentBlock) SetCitations(v MRCitationsConfig)`

SetCitations sets Citations field to given value.

### HasCitations

`func (o *MRDocumentBlock) HasCitations() bool`

HasCitations returns a boolean if a field has been set.

### SetCitationsNil

`func (o *MRDocumentBlock) SetCitationsNil(b bool)`

 SetCitationsNil sets the value for Citations to be an explicit nil

### UnsetCitations
`func (o *MRDocumentBlock) UnsetCitations()`

UnsetCitations ensures that no value is present for Citations, not even an explicit nil
### GetSource

`func (o *MRDocumentBlock) GetSource() Source`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *MRDocumentBlock) GetSourceOk() (*Source, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *MRDocumentBlock) SetSource(v Source)`

SetSource sets Source field to given value.


### GetTitle

`func (o *MRDocumentBlock) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *MRDocumentBlock) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *MRDocumentBlock) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *MRDocumentBlock) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *MRDocumentBlock) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *MRDocumentBlock) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetType

`func (o *MRDocumentBlock) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MRDocumentBlock) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MRDocumentBlock) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


