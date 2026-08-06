# MRBetaCitationContentBlockLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CitedText** | **string** |  | 
**DocumentIndex** | **int32** |  | 
**DocumentTitle** | Pointer to **NullableString** |  | [optional] 
**EndBlockIndex** | **int32** |  | 
**FileId** | Pointer to **NullableString** |  | [optional] 
**StartBlockIndex** | **int32** |  | 
**Type** | **string** |  | 

## Methods

### NewMRBetaCitationContentBlockLocation

`func NewMRBetaCitationContentBlockLocation(citedText string, documentIndex int32, endBlockIndex int32, startBlockIndex int32, type_ string, ) *MRBetaCitationContentBlockLocation`

NewMRBetaCitationContentBlockLocation instantiates a new MRBetaCitationContentBlockLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMRBetaCitationContentBlockLocationWithDefaults

`func NewMRBetaCitationContentBlockLocationWithDefaults() *MRBetaCitationContentBlockLocation`

NewMRBetaCitationContentBlockLocationWithDefaults instantiates a new MRBetaCitationContentBlockLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCitedText

`func (o *MRBetaCitationContentBlockLocation) GetCitedText() string`

GetCitedText returns the CitedText field if non-nil, zero value otherwise.

### GetCitedTextOk

`func (o *MRBetaCitationContentBlockLocation) GetCitedTextOk() (*string, bool)`

GetCitedTextOk returns a tuple with the CitedText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitedText

`func (o *MRBetaCitationContentBlockLocation) SetCitedText(v string)`

SetCitedText sets CitedText field to given value.


### GetDocumentIndex

`func (o *MRBetaCitationContentBlockLocation) GetDocumentIndex() int32`

GetDocumentIndex returns the DocumentIndex field if non-nil, zero value otherwise.

### GetDocumentIndexOk

`func (o *MRBetaCitationContentBlockLocation) GetDocumentIndexOk() (*int32, bool)`

GetDocumentIndexOk returns a tuple with the DocumentIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentIndex

`func (o *MRBetaCitationContentBlockLocation) SetDocumentIndex(v int32)`

SetDocumentIndex sets DocumentIndex field to given value.


### GetDocumentTitle

`func (o *MRBetaCitationContentBlockLocation) GetDocumentTitle() string`

GetDocumentTitle returns the DocumentTitle field if non-nil, zero value otherwise.

### GetDocumentTitleOk

`func (o *MRBetaCitationContentBlockLocation) GetDocumentTitleOk() (*string, bool)`

GetDocumentTitleOk returns a tuple with the DocumentTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentTitle

`func (o *MRBetaCitationContentBlockLocation) SetDocumentTitle(v string)`

SetDocumentTitle sets DocumentTitle field to given value.

### HasDocumentTitle

`func (o *MRBetaCitationContentBlockLocation) HasDocumentTitle() bool`

HasDocumentTitle returns a boolean if a field has been set.

### SetDocumentTitleNil

`func (o *MRBetaCitationContentBlockLocation) SetDocumentTitleNil(b bool)`

 SetDocumentTitleNil sets the value for DocumentTitle to be an explicit nil

### UnsetDocumentTitle
`func (o *MRBetaCitationContentBlockLocation) UnsetDocumentTitle()`

UnsetDocumentTitle ensures that no value is present for DocumentTitle, not even an explicit nil
### GetEndBlockIndex

`func (o *MRBetaCitationContentBlockLocation) GetEndBlockIndex() int32`

GetEndBlockIndex returns the EndBlockIndex field if non-nil, zero value otherwise.

### GetEndBlockIndexOk

`func (o *MRBetaCitationContentBlockLocation) GetEndBlockIndexOk() (*int32, bool)`

GetEndBlockIndexOk returns a tuple with the EndBlockIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndBlockIndex

`func (o *MRBetaCitationContentBlockLocation) SetEndBlockIndex(v int32)`

SetEndBlockIndex sets EndBlockIndex field to given value.


### GetFileId

`func (o *MRBetaCitationContentBlockLocation) GetFileId() string`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *MRBetaCitationContentBlockLocation) GetFileIdOk() (*string, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *MRBetaCitationContentBlockLocation) SetFileId(v string)`

SetFileId sets FileId field to given value.

### HasFileId

`func (o *MRBetaCitationContentBlockLocation) HasFileId() bool`

HasFileId returns a boolean if a field has been set.

### SetFileIdNil

`func (o *MRBetaCitationContentBlockLocation) SetFileIdNil(b bool)`

 SetFileIdNil sets the value for FileId to be an explicit nil

### UnsetFileId
`func (o *MRBetaCitationContentBlockLocation) UnsetFileId()`

UnsetFileId ensures that no value is present for FileId, not even an explicit nil
### GetStartBlockIndex

`func (o *MRBetaCitationContentBlockLocation) GetStartBlockIndex() int32`

GetStartBlockIndex returns the StartBlockIndex field if non-nil, zero value otherwise.

### GetStartBlockIndexOk

`func (o *MRBetaCitationContentBlockLocation) GetStartBlockIndexOk() (*int32, bool)`

GetStartBlockIndexOk returns a tuple with the StartBlockIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartBlockIndex

`func (o *MRBetaCitationContentBlockLocation) SetStartBlockIndex(v int32)`

SetStartBlockIndex sets StartBlockIndex field to given value.


### GetType

`func (o *MRBetaCitationContentBlockLocation) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MRBetaCitationContentBlockLocation) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MRBetaCitationContentBlockLocation) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


