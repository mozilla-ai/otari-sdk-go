# MRBetaDocumentBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Citations** | Pointer to [**NullableMRBetaCitationConfig**](MRBetaCitationConfig.md) |  | [optional] 
**Source** | [**Source**](Source.md) |  | 
**Title** | Pointer to **NullableString** |  | [optional] 
**Type** | **string** |  | 

## Methods

### NewMRBetaDocumentBlock

`func NewMRBetaDocumentBlock(source Source, type_ string, ) *MRBetaDocumentBlock`

NewMRBetaDocumentBlock instantiates a new MRBetaDocumentBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMRBetaDocumentBlockWithDefaults

`func NewMRBetaDocumentBlockWithDefaults() *MRBetaDocumentBlock`

NewMRBetaDocumentBlockWithDefaults instantiates a new MRBetaDocumentBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCitations

`func (o *MRBetaDocumentBlock) GetCitations() MRBetaCitationConfig`

GetCitations returns the Citations field if non-nil, zero value otherwise.

### GetCitationsOk

`func (o *MRBetaDocumentBlock) GetCitationsOk() (*MRBetaCitationConfig, bool)`

GetCitationsOk returns a tuple with the Citations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitations

`func (o *MRBetaDocumentBlock) SetCitations(v MRBetaCitationConfig)`

SetCitations sets Citations field to given value.

### HasCitations

`func (o *MRBetaDocumentBlock) HasCitations() bool`

HasCitations returns a boolean if a field has been set.

### SetCitationsNil

`func (o *MRBetaDocumentBlock) SetCitationsNil(b bool)`

 SetCitationsNil sets the value for Citations to be an explicit nil

### UnsetCitations
`func (o *MRBetaDocumentBlock) UnsetCitations()`

UnsetCitations ensures that no value is present for Citations, not even an explicit nil
### GetSource

`func (o *MRBetaDocumentBlock) GetSource() Source`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *MRBetaDocumentBlock) GetSourceOk() (*Source, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *MRBetaDocumentBlock) SetSource(v Source)`

SetSource sets Source field to given value.


### GetTitle

`func (o *MRBetaDocumentBlock) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *MRBetaDocumentBlock) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *MRBetaDocumentBlock) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *MRBetaDocumentBlock) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *MRBetaDocumentBlock) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *MRBetaDocumentBlock) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetType

`func (o *MRBetaDocumentBlock) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MRBetaDocumentBlock) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MRBetaDocumentBlock) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


