# MRBetaTextBlockCitationsInner

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

### NewMRBetaTextBlockCitationsInner

`func NewMRBetaTextBlockCitationsInner(citedText string, documentIndex int32, endCharIndex int32, startCharIndex int32, type_ string, endPageNumber int32, startPageNumber int32, endBlockIndex int32, startBlockIndex int32, encryptedIndex string, url string, searchResultIndex int32, source string, ) *MRBetaTextBlockCitationsInner`

NewMRBetaTextBlockCitationsInner instantiates a new MRBetaTextBlockCitationsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMRBetaTextBlockCitationsInnerWithDefaults

`func NewMRBetaTextBlockCitationsInnerWithDefaults() *MRBetaTextBlockCitationsInner`

NewMRBetaTextBlockCitationsInnerWithDefaults instantiates a new MRBetaTextBlockCitationsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCitedText

`func (o *MRBetaTextBlockCitationsInner) GetCitedText() string`

GetCitedText returns the CitedText field if non-nil, zero value otherwise.

### GetCitedTextOk

`func (o *MRBetaTextBlockCitationsInner) GetCitedTextOk() (*string, bool)`

GetCitedTextOk returns a tuple with the CitedText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitedText

`func (o *MRBetaTextBlockCitationsInner) SetCitedText(v string)`

SetCitedText sets CitedText field to given value.


### GetDocumentIndex

`func (o *MRBetaTextBlockCitationsInner) GetDocumentIndex() int32`

GetDocumentIndex returns the DocumentIndex field if non-nil, zero value otherwise.

### GetDocumentIndexOk

`func (o *MRBetaTextBlockCitationsInner) GetDocumentIndexOk() (*int32, bool)`

GetDocumentIndexOk returns a tuple with the DocumentIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentIndex

`func (o *MRBetaTextBlockCitationsInner) SetDocumentIndex(v int32)`

SetDocumentIndex sets DocumentIndex field to given value.


### GetDocumentTitle

`func (o *MRBetaTextBlockCitationsInner) GetDocumentTitle() string`

GetDocumentTitle returns the DocumentTitle field if non-nil, zero value otherwise.

### GetDocumentTitleOk

`func (o *MRBetaTextBlockCitationsInner) GetDocumentTitleOk() (*string, bool)`

GetDocumentTitleOk returns a tuple with the DocumentTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentTitle

`func (o *MRBetaTextBlockCitationsInner) SetDocumentTitle(v string)`

SetDocumentTitle sets DocumentTitle field to given value.

### HasDocumentTitle

`func (o *MRBetaTextBlockCitationsInner) HasDocumentTitle() bool`

HasDocumentTitle returns a boolean if a field has been set.

### GetEndCharIndex

`func (o *MRBetaTextBlockCitationsInner) GetEndCharIndex() int32`

GetEndCharIndex returns the EndCharIndex field if non-nil, zero value otherwise.

### GetEndCharIndexOk

`func (o *MRBetaTextBlockCitationsInner) GetEndCharIndexOk() (*int32, bool)`

GetEndCharIndexOk returns a tuple with the EndCharIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndCharIndex

`func (o *MRBetaTextBlockCitationsInner) SetEndCharIndex(v int32)`

SetEndCharIndex sets EndCharIndex field to given value.


### GetFileId

`func (o *MRBetaTextBlockCitationsInner) GetFileId() string`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *MRBetaTextBlockCitationsInner) GetFileIdOk() (*string, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *MRBetaTextBlockCitationsInner) SetFileId(v string)`

SetFileId sets FileId field to given value.

### HasFileId

`func (o *MRBetaTextBlockCitationsInner) HasFileId() bool`

HasFileId returns a boolean if a field has been set.

### GetStartCharIndex

`func (o *MRBetaTextBlockCitationsInner) GetStartCharIndex() int32`

GetStartCharIndex returns the StartCharIndex field if non-nil, zero value otherwise.

### GetStartCharIndexOk

`func (o *MRBetaTextBlockCitationsInner) GetStartCharIndexOk() (*int32, bool)`

GetStartCharIndexOk returns a tuple with the StartCharIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartCharIndex

`func (o *MRBetaTextBlockCitationsInner) SetStartCharIndex(v int32)`

SetStartCharIndex sets StartCharIndex field to given value.


### GetType

`func (o *MRBetaTextBlockCitationsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MRBetaTextBlockCitationsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MRBetaTextBlockCitationsInner) SetType(v string)`

SetType sets Type field to given value.


### GetEndPageNumber

`func (o *MRBetaTextBlockCitationsInner) GetEndPageNumber() int32`

GetEndPageNumber returns the EndPageNumber field if non-nil, zero value otherwise.

### GetEndPageNumberOk

`func (o *MRBetaTextBlockCitationsInner) GetEndPageNumberOk() (*int32, bool)`

GetEndPageNumberOk returns a tuple with the EndPageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndPageNumber

`func (o *MRBetaTextBlockCitationsInner) SetEndPageNumber(v int32)`

SetEndPageNumber sets EndPageNumber field to given value.


### GetStartPageNumber

`func (o *MRBetaTextBlockCitationsInner) GetStartPageNumber() int32`

GetStartPageNumber returns the StartPageNumber field if non-nil, zero value otherwise.

### GetStartPageNumberOk

`func (o *MRBetaTextBlockCitationsInner) GetStartPageNumberOk() (*int32, bool)`

GetStartPageNumberOk returns a tuple with the StartPageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartPageNumber

`func (o *MRBetaTextBlockCitationsInner) SetStartPageNumber(v int32)`

SetStartPageNumber sets StartPageNumber field to given value.


### GetEndBlockIndex

`func (o *MRBetaTextBlockCitationsInner) GetEndBlockIndex() int32`

GetEndBlockIndex returns the EndBlockIndex field if non-nil, zero value otherwise.

### GetEndBlockIndexOk

`func (o *MRBetaTextBlockCitationsInner) GetEndBlockIndexOk() (*int32, bool)`

GetEndBlockIndexOk returns a tuple with the EndBlockIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndBlockIndex

`func (o *MRBetaTextBlockCitationsInner) SetEndBlockIndex(v int32)`

SetEndBlockIndex sets EndBlockIndex field to given value.


### GetStartBlockIndex

`func (o *MRBetaTextBlockCitationsInner) GetStartBlockIndex() int32`

GetStartBlockIndex returns the StartBlockIndex field if non-nil, zero value otherwise.

### GetStartBlockIndexOk

`func (o *MRBetaTextBlockCitationsInner) GetStartBlockIndexOk() (*int32, bool)`

GetStartBlockIndexOk returns a tuple with the StartBlockIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartBlockIndex

`func (o *MRBetaTextBlockCitationsInner) SetStartBlockIndex(v int32)`

SetStartBlockIndex sets StartBlockIndex field to given value.


### GetEncryptedIndex

`func (o *MRBetaTextBlockCitationsInner) GetEncryptedIndex() string`

GetEncryptedIndex returns the EncryptedIndex field if non-nil, zero value otherwise.

### GetEncryptedIndexOk

`func (o *MRBetaTextBlockCitationsInner) GetEncryptedIndexOk() (*string, bool)`

GetEncryptedIndexOk returns a tuple with the EncryptedIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedIndex

`func (o *MRBetaTextBlockCitationsInner) SetEncryptedIndex(v string)`

SetEncryptedIndex sets EncryptedIndex field to given value.


### GetTitle

`func (o *MRBetaTextBlockCitationsInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *MRBetaTextBlockCitationsInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *MRBetaTextBlockCitationsInner) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *MRBetaTextBlockCitationsInner) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUrl

`func (o *MRBetaTextBlockCitationsInner) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *MRBetaTextBlockCitationsInner) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *MRBetaTextBlockCitationsInner) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetSearchResultIndex

`func (o *MRBetaTextBlockCitationsInner) GetSearchResultIndex() int32`

GetSearchResultIndex returns the SearchResultIndex field if non-nil, zero value otherwise.

### GetSearchResultIndexOk

`func (o *MRBetaTextBlockCitationsInner) GetSearchResultIndexOk() (*int32, bool)`

GetSearchResultIndexOk returns a tuple with the SearchResultIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchResultIndex

`func (o *MRBetaTextBlockCitationsInner) SetSearchResultIndex(v int32)`

SetSearchResultIndex sets SearchResultIndex field to given value.


### GetSource

`func (o *MRBetaTextBlockCitationsInner) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *MRBetaTextBlockCitationsInner) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *MRBetaTextBlockCitationsInner) SetSource(v string)`

SetSource sets Source field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


