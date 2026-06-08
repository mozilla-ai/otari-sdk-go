# MRTextBlockCitationsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CitedText** | **string** |  | 
**DocumentIndex** | **int32** |  | 
**DocumentTitle** | Pointer to **string** |  | [optional] 
**EndCharIndex** | **int32** |  | 
**FileId** | Pointer to **string** |  | [optional] 
**StartCharIndex** | **int32** |  | 
**Type** | **string** |  | 
**EndPageNumber** | **int32** |  | 
**StartPageNumber** | **int32** |  | 
**EndBlockIndex** | **int32** |  | 
**StartBlockIndex** | **int32** |  | 
**EncryptedIndex** | **string** |  | 
**Title** | Pointer to **string** |  | [optional] 
**Url** | **string** |  | 
**SearchResultIndex** | **int32** |  | 
**Source** | **string** |  | 

## Methods

### NewMRTextBlockCitationsInner

`func NewMRTextBlockCitationsInner(citedText string, documentIndex int32, endCharIndex int32, startCharIndex int32, type_ string, endPageNumber int32, startPageNumber int32, endBlockIndex int32, startBlockIndex int32, encryptedIndex string, url string, searchResultIndex int32, source string, ) *MRTextBlockCitationsInner`

NewMRTextBlockCitationsInner instantiates a new MRTextBlockCitationsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMRTextBlockCitationsInnerWithDefaults

`func NewMRTextBlockCitationsInnerWithDefaults() *MRTextBlockCitationsInner`

NewMRTextBlockCitationsInnerWithDefaults instantiates a new MRTextBlockCitationsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCitedText

`func (o *MRTextBlockCitationsInner) GetCitedText() string`

GetCitedText returns the CitedText field if non-nil, zero value otherwise.

### GetCitedTextOk

`func (o *MRTextBlockCitationsInner) GetCitedTextOk() (*string, bool)`

GetCitedTextOk returns a tuple with the CitedText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitedText

`func (o *MRTextBlockCitationsInner) SetCitedText(v string)`

SetCitedText sets CitedText field to given value.


### GetDocumentIndex

`func (o *MRTextBlockCitationsInner) GetDocumentIndex() int32`

GetDocumentIndex returns the DocumentIndex field if non-nil, zero value otherwise.

### GetDocumentIndexOk

`func (o *MRTextBlockCitationsInner) GetDocumentIndexOk() (*int32, bool)`

GetDocumentIndexOk returns a tuple with the DocumentIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentIndex

`func (o *MRTextBlockCitationsInner) SetDocumentIndex(v int32)`

SetDocumentIndex sets DocumentIndex field to given value.


### GetDocumentTitle

`func (o *MRTextBlockCitationsInner) GetDocumentTitle() string`

GetDocumentTitle returns the DocumentTitle field if non-nil, zero value otherwise.

### GetDocumentTitleOk

`func (o *MRTextBlockCitationsInner) GetDocumentTitleOk() (*string, bool)`

GetDocumentTitleOk returns a tuple with the DocumentTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentTitle

`func (o *MRTextBlockCitationsInner) SetDocumentTitle(v string)`

SetDocumentTitle sets DocumentTitle field to given value.

### HasDocumentTitle

`func (o *MRTextBlockCitationsInner) HasDocumentTitle() bool`

HasDocumentTitle returns a boolean if a field has been set.

### GetEndCharIndex

`func (o *MRTextBlockCitationsInner) GetEndCharIndex() int32`

GetEndCharIndex returns the EndCharIndex field if non-nil, zero value otherwise.

### GetEndCharIndexOk

`func (o *MRTextBlockCitationsInner) GetEndCharIndexOk() (*int32, bool)`

GetEndCharIndexOk returns a tuple with the EndCharIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndCharIndex

`func (o *MRTextBlockCitationsInner) SetEndCharIndex(v int32)`

SetEndCharIndex sets EndCharIndex field to given value.


### GetFileId

`func (o *MRTextBlockCitationsInner) GetFileId() string`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *MRTextBlockCitationsInner) GetFileIdOk() (*string, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *MRTextBlockCitationsInner) SetFileId(v string)`

SetFileId sets FileId field to given value.

### HasFileId

`func (o *MRTextBlockCitationsInner) HasFileId() bool`

HasFileId returns a boolean if a field has been set.

### GetStartCharIndex

`func (o *MRTextBlockCitationsInner) GetStartCharIndex() int32`

GetStartCharIndex returns the StartCharIndex field if non-nil, zero value otherwise.

### GetStartCharIndexOk

`func (o *MRTextBlockCitationsInner) GetStartCharIndexOk() (*int32, bool)`

GetStartCharIndexOk returns a tuple with the StartCharIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartCharIndex

`func (o *MRTextBlockCitationsInner) SetStartCharIndex(v int32)`

SetStartCharIndex sets StartCharIndex field to given value.


### GetType

`func (o *MRTextBlockCitationsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MRTextBlockCitationsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MRTextBlockCitationsInner) SetType(v string)`

SetType sets Type field to given value.


### GetEndPageNumber

`func (o *MRTextBlockCitationsInner) GetEndPageNumber() int32`

GetEndPageNumber returns the EndPageNumber field if non-nil, zero value otherwise.

### GetEndPageNumberOk

`func (o *MRTextBlockCitationsInner) GetEndPageNumberOk() (*int32, bool)`

GetEndPageNumberOk returns a tuple with the EndPageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPageNumber

`func (o *MRTextBlockCitationsInner) SetEndPageNumber(v int32)`

SetEndPageNumber sets EndPageNumber field to given value.


### GetStartPageNumber

`func (o *MRTextBlockCitationsInner) GetStartPageNumber() int32`

GetStartPageNumber returns the StartPageNumber field if non-nil, zero value otherwise.

### GetStartPageNumberOk

`func (o *MRTextBlockCitationsInner) GetStartPageNumberOk() (*int32, bool)`

GetStartPageNumberOk returns a tuple with the StartPageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartPageNumber

`func (o *MRTextBlockCitationsInner) SetStartPageNumber(v int32)`

SetStartPageNumber sets StartPageNumber field to given value.


### GetEndBlockIndex

`func (o *MRTextBlockCitationsInner) GetEndBlockIndex() int32`

GetEndBlockIndex returns the EndBlockIndex field if non-nil, zero value otherwise.

### GetEndBlockIndexOk

`func (o *MRTextBlockCitationsInner) GetEndBlockIndexOk() (*int32, bool)`

GetEndBlockIndexOk returns a tuple with the EndBlockIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndBlockIndex

`func (o *MRTextBlockCitationsInner) SetEndBlockIndex(v int32)`

SetEndBlockIndex sets EndBlockIndex field to given value.


### GetStartBlockIndex

`func (o *MRTextBlockCitationsInner) GetStartBlockIndex() int32`

GetStartBlockIndex returns the StartBlockIndex field if non-nil, zero value otherwise.

### GetStartBlockIndexOk

`func (o *MRTextBlockCitationsInner) GetStartBlockIndexOk() (*int32, bool)`

GetStartBlockIndexOk returns a tuple with the StartBlockIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartBlockIndex

`func (o *MRTextBlockCitationsInner) SetStartBlockIndex(v int32)`

SetStartBlockIndex sets StartBlockIndex field to given value.


### GetEncryptedIndex

`func (o *MRTextBlockCitationsInner) GetEncryptedIndex() string`

GetEncryptedIndex returns the EncryptedIndex field if non-nil, zero value otherwise.

### GetEncryptedIndexOk

`func (o *MRTextBlockCitationsInner) GetEncryptedIndexOk() (*string, bool)`

GetEncryptedIndexOk returns a tuple with the EncryptedIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedIndex

`func (o *MRTextBlockCitationsInner) SetEncryptedIndex(v string)`

SetEncryptedIndex sets EncryptedIndex field to given value.


### GetTitle

`func (o *MRTextBlockCitationsInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *MRTextBlockCitationsInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *MRTextBlockCitationsInner) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *MRTextBlockCitationsInner) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUrl

`func (o *MRTextBlockCitationsInner) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *MRTextBlockCitationsInner) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *MRTextBlockCitationsInner) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetSearchResultIndex

`func (o *MRTextBlockCitationsInner) GetSearchResultIndex() int32`

GetSearchResultIndex returns the SearchResultIndex field if non-nil, zero value otherwise.

### GetSearchResultIndexOk

`func (o *MRTextBlockCitationsInner) GetSearchResultIndexOk() (*int32, bool)`

GetSearchResultIndexOk returns a tuple with the SearchResultIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchResultIndex

`func (o *MRTextBlockCitationsInner) SetSearchResultIndex(v int32)`

SetSearchResultIndex sets SearchResultIndex field to given value.


### GetSource

`func (o *MRTextBlockCitationsInner) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *MRTextBlockCitationsInner) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *MRTextBlockCitationsInner) SetSource(v string)`

SetSource sets Source field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


