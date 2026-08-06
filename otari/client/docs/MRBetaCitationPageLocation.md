# MRBetaCitationPageLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CitedText** | **string** |  | 
**DocumentIndex** | **int32** |  | 
**DocumentTitle** | Pointer to **NullableString** |  | [optional] 
**EndPageNumber** | **int32** |  | 
**FileId** | Pointer to **NullableString** |  | [optional] 
**StartPageNumber** | **int32** |  | 
**Type** | **string** |  | 

## Methods

### NewMRBetaCitationPageLocation

`func NewMRBetaCitationPageLocation(citedText string, documentIndex int32, endPageNumber int32, startPageNumber int32, type_ string, ) *MRBetaCitationPageLocation`

NewMRBetaCitationPageLocation instantiates a new MRBetaCitationPageLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMRBetaCitationPageLocationWithDefaults

`func NewMRBetaCitationPageLocationWithDefaults() *MRBetaCitationPageLocation`

NewMRBetaCitationPageLocationWithDefaults instantiates a new MRBetaCitationPageLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCitedText

`func (o *MRBetaCitationPageLocation) GetCitedText() string`

GetCitedText returns the CitedText field if non-nil, zero value otherwise.

### GetCitedTextOk

`func (o *MRBetaCitationPageLocation) GetCitedTextOk() (*string, bool)`

GetCitedTextOk returns a tuple with the CitedText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitedText

`func (o *MRBetaCitationPageLocation) SetCitedText(v string)`

SetCitedText sets CitedText field to given value.


### GetDocumentIndex

`func (o *MRBetaCitationPageLocation) GetDocumentIndex() int32`

GetDocumentIndex returns the DocumentIndex field if non-nil, zero value otherwise.

### GetDocumentIndexOk

`func (o *MRBetaCitationPageLocation) GetDocumentIndexOk() (*int32, bool)`

GetDocumentIndexOk returns a tuple with the DocumentIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentIndex

`func (o *MRBetaCitationPageLocation) SetDocumentIndex(v int32)`

SetDocumentIndex sets DocumentIndex field to given value.


### GetDocumentTitle

`func (o *MRBetaCitationPageLocation) GetDocumentTitle() string`

GetDocumentTitle returns the DocumentTitle field if non-nil, zero value otherwise.

### GetDocumentTitleOk

`func (o *MRBetaCitationPageLocation) GetDocumentTitleOk() (*string, bool)`

GetDocumentTitleOk returns a tuple with the DocumentTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentTitle

`func (o *MRBetaCitationPageLocation) SetDocumentTitle(v string)`

SetDocumentTitle sets DocumentTitle field to given value.

### HasDocumentTitle

`func (o *MRBetaCitationPageLocation) HasDocumentTitle() bool`

HasDocumentTitle returns a boolean if a field has been set.

### SetDocumentTitleNil

`func (o *MRBetaCitationPageLocation) SetDocumentTitleNil(b bool)`

 SetDocumentTitleNil sets the value for DocumentTitle to be an explicit nil

### UnsetDocumentTitle
`func (o *MRBetaCitationPageLocation) UnsetDocumentTitle()`

UnsetDocumentTitle ensures that no value is present for DocumentTitle, not even an explicit nil
### GetEndPageNumber

`func (o *MRBetaCitationPageLocation) GetEndPageNumber() int32`

GetEndPageNumber returns the EndPageNumber field if non-nil, zero value otherwise.

### GetEndPageNumberOk

`func (o *MRBetaCitationPageLocation) GetEndPageNumberOk() (*int32, bool)`

GetEndPageNumberOk returns a tuple with the EndPageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPageNumber

`func (o *MRBetaCitationPageLocation) SetEndPageNumber(v int32)`

SetEndPageNumber sets EndPageNumber field to given value.


### GetFileId

`func (o *MRBetaCitationPageLocation) GetFileId() string`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *MRBetaCitationPageLocation) GetFileIdOk() (*string, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *MRBetaCitationPageLocation) SetFileId(v string)`

SetFileId sets FileId field to given value.

### HasFileId

`func (o *MRBetaCitationPageLocation) HasFileId() bool`

HasFileId returns a boolean if a field has been set.

### SetFileIdNil

`func (o *MRBetaCitationPageLocation) SetFileIdNil(b bool)`

 SetFileIdNil sets the value for FileId to be an explicit nil

### UnsetFileId
`func (o *MRBetaCitationPageLocation) UnsetFileId()`

UnsetFileId ensures that no value is present for FileId, not even an explicit nil
### GetStartPageNumber

`func (o *MRBetaCitationPageLocation) GetStartPageNumber() int32`

GetStartPageNumber returns the StartPageNumber field if non-nil, zero value otherwise.

### GetStartPageNumberOk

`func (o *MRBetaCitationPageLocation) GetStartPageNumberOk() (*int32, bool)`

GetStartPageNumberOk returns a tuple with the StartPageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartPageNumber

`func (o *MRBetaCitationPageLocation) SetStartPageNumber(v int32)`

SetStartPageNumber sets StartPageNumber field to given value.


### GetType

`func (o *MRBetaCitationPageLocation) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MRBetaCitationPageLocation) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MRBetaCitationPageLocation) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


