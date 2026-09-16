# WebSearchBackendResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to **string** |  | [optional] [default to ""]
**ExtractedContent** | Pointer to **NullableString** | The page&#39;s own text, when the provider returned it, so the caller can skip fetching the page. | [optional] 
**PublishedDate** | Pointer to **NullableString** | The provider&#39;s own recency string for the page, forwarded unparsed. Declared so a search over this hop renders the same date an in-process one does. | [optional] 
**Title** | Pointer to **string** |  | [optional] [default to ""]
**Url** | **string** |  | 

## Methods

### NewWebSearchBackendResult

`func NewWebSearchBackendResult(url string, ) *WebSearchBackendResult`

NewWebSearchBackendResult instantiates a new WebSearchBackendResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebSearchBackendResultWithDefaults

`func NewWebSearchBackendResultWithDefaults() *WebSearchBackendResult`

NewWebSearchBackendResultWithDefaults instantiates a new WebSearchBackendResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *WebSearchBackendResult) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *WebSearchBackendResult) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *WebSearchBackendResult) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *WebSearchBackendResult) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetExtractedContent

`func (o *WebSearchBackendResult) GetExtractedContent() string`

GetExtractedContent returns the ExtractedContent field if non-nil, zero value otherwise.

### GetExtractedContentOk

`func (o *WebSearchBackendResult) GetExtractedContentOk() (*string, bool)`

GetExtractedContentOk returns a tuple with the ExtractedContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtractedContent

`func (o *WebSearchBackendResult) SetExtractedContent(v string)`

SetExtractedContent sets ExtractedContent field to given value.

### HasExtractedContent

`func (o *WebSearchBackendResult) HasExtractedContent() bool`

HasExtractedContent returns a boolean if a field has been set.

### SetExtractedContentNil

`func (o *WebSearchBackendResult) SetExtractedContentNil(b bool)`

 SetExtractedContentNil sets the value for ExtractedContent to be an explicit nil

### UnsetExtractedContent
`func (o *WebSearchBackendResult) UnsetExtractedContent()`

UnsetExtractedContent ensures that no value is present for ExtractedContent, not even an explicit nil
### GetPublishedDate

`func (o *WebSearchBackendResult) GetPublishedDate() string`

GetPublishedDate returns the PublishedDate field if non-nil, zero value otherwise.

### GetPublishedDateOk

`func (o *WebSearchBackendResult) GetPublishedDateOk() (*string, bool)`

GetPublishedDateOk returns a tuple with the PublishedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedDate

`func (o *WebSearchBackendResult) SetPublishedDate(v string)`

SetPublishedDate sets PublishedDate field to given value.

### HasPublishedDate

`func (o *WebSearchBackendResult) HasPublishedDate() bool`

HasPublishedDate returns a boolean if a field has been set.

### SetPublishedDateNil

`func (o *WebSearchBackendResult) SetPublishedDateNil(b bool)`

 SetPublishedDateNil sets the value for PublishedDate to be an explicit nil

### UnsetPublishedDate
`func (o *WebSearchBackendResult) UnsetPublishedDate()`

UnsetPublishedDate ensures that no value is present for PublishedDate, not even an explicit nil
### GetTitle

`func (o *WebSearchBackendResult) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *WebSearchBackendResult) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *WebSearchBackendResult) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *WebSearchBackendResult) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUrl

`func (o *WebSearchBackendResult) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *WebSearchBackendResult) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *WebSearchBackendResult) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


