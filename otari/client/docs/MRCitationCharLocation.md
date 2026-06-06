# MRCitationCharLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CitedText** | **string** |  | 
**DocumentIndex** | **int32** |  | 
**DocumentTitle** | Pointer to **NullableString** |  | [optional] 
**EndCharIndex** | **int32** |  | 
**FileId** | Pointer to **NullableString** |  | [optional] 
**StartCharIndex** | **int32** |  | 
**Type** | **string** |  | 

## Methods

### NewMRCitationCharLocation

`func NewMRCitationCharLocation(citedText string, documentIndex int32, endCharIndex int32, startCharIndex int32, type_ string, ) *MRCitationCharLocation`

NewMRCitationCharLocation instantiates a new MRCitationCharLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMRCitationCharLocationWithDefaults

`func NewMRCitationCharLocationWithDefaults() *MRCitationCharLocation`

NewMRCitationCharLocationWithDefaults instantiates a new MRCitationCharLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCitedText

`func (o *MRCitationCharLocation) GetCitedText() string`

GetCitedText returns the CitedText field if non-nil, zero value otherwise.

### GetCitedTextOk

`func (o *MRCitationCharLocation) GetCitedTextOk() (*string, bool)`

GetCitedTextOk returns a tuple with the CitedText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitedText

`func (o *MRCitationCharLocation) SetCitedText(v string)`

SetCitedText sets CitedText field to given value.


### GetDocumentIndex

`func (o *MRCitationCharLocation) GetDocumentIndex() int32`

GetDocumentIndex returns the DocumentIndex field if non-nil, zero value otherwise.

### GetDocumentIndexOk

`func (o *MRCitationCharLocation) GetDocumentIndexOk() (*int32, bool)`

GetDocumentIndexOk returns a tuple with the DocumentIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentIndex

`func (o *MRCitationCharLocation) SetDocumentIndex(v int32)`

SetDocumentIndex sets DocumentIndex field to given value.


### GetDocumentTitle

`func (o *MRCitationCharLocation) GetDocumentTitle() string`

GetDocumentTitle returns the DocumentTitle field if non-nil, zero value otherwise.

### GetDocumentTitleOk

`func (o *MRCitationCharLocation) GetDocumentTitleOk() (*string, bool)`

GetDocumentTitleOk returns a tuple with the DocumentTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentTitle

`func (o *MRCitationCharLocation) SetDocumentTitle(v string)`

SetDocumentTitle sets DocumentTitle field to given value.

### HasDocumentTitle

`func (o *MRCitationCharLocation) HasDocumentTitle() bool`

HasDocumentTitle returns a boolean if a field has been set.

### SetDocumentTitleNil

`func (o *MRCitationCharLocation) SetDocumentTitleNil(b bool)`

 SetDocumentTitleNil sets the value for DocumentTitle to be an explicit nil

### UnsetDocumentTitle
`func (o *MRCitationCharLocation) UnsetDocumentTitle()`

UnsetDocumentTitle ensures that no value is present for DocumentTitle, not even an explicit nil
### GetEndCharIndex

`func (o *MRCitationCharLocation) GetEndCharIndex() int32`

GetEndCharIndex returns the EndCharIndex field if non-nil, zero value otherwise.

### GetEndCharIndexOk

`func (o *MRCitationCharLocation) GetEndCharIndexOk() (*int32, bool)`

GetEndCharIndexOk returns a tuple with the EndCharIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndCharIndex

`func (o *MRCitationCharLocation) SetEndCharIndex(v int32)`

SetEndCharIndex sets EndCharIndex field to given value.


### GetFileId

`func (o *MRCitationCharLocation) GetFileId() string`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *MRCitationCharLocation) GetFileIdOk() (*string, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *MRCitationCharLocation) SetFileId(v string)`

SetFileId sets FileId field to given value.

### HasFileId

`func (o *MRCitationCharLocation) HasFileId() bool`

HasFileId returns a boolean if a field has been set.

### SetFileIdNil

`func (o *MRCitationCharLocation) SetFileIdNil(b bool)`

 SetFileIdNil sets the value for FileId to be an explicit nil

### UnsetFileId
`func (o *MRCitationCharLocation) UnsetFileId()`

UnsetFileId ensures that no value is present for FileId, not even an explicit nil
### GetStartCharIndex

`func (o *MRCitationCharLocation) GetStartCharIndex() int32`

GetStartCharIndex returns the StartCharIndex field if non-nil, zero value otherwise.

### GetStartCharIndexOk

`func (o *MRCitationCharLocation) GetStartCharIndexOk() (*int32, bool)`

GetStartCharIndexOk returns a tuple with the StartCharIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartCharIndex

`func (o *MRCitationCharLocation) SetStartCharIndex(v int32)`

SetStartCharIndex sets StartCharIndex field to given value.


### GetType

`func (o *MRCitationCharLocation) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MRCitationCharLocation) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MRCitationCharLocation) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


