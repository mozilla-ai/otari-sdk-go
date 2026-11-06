# FetchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentType** | Pointer to **string** | The page&#39;s content type; empty when the provider does not report one. | [optional] [default to ""]
**FetchTool** | **string** | The configured fetch tool that served the request | 
**FinalUrl** | **string** | Where the page was found after redirects. The URL asked for when the provider does not report where a redirect led (reports_final_url in the provider catalog). | 
**Object** | Pointer to **string** |  | [optional] [default to "fetch"]
**Published** | Pointer to **NullableTime** | When the page was published, where the provider reports it. | [optional] 
**Raw** | Pointer to **map[string]interface{}** | Tags for cost attribution, recorded on the request&#39;s usage rows and filterable in the usage API: up to 16 string pairs, keys up to 64 characters and values up to 512. A null value is ignored. LiteLLM&#39;s nested &#x60;spend_logs_metadata&#x60; object is also read, and wins over a flat key of the same name; it is never forwarded to the provider. | [optional] 
**Text** | **string** | The page&#39;s text, as the provider extracted it. | 
**Title** | Pointer to **string** | The page&#39;s title; empty when the provider gave none. | [optional] [default to ""]
**Truncated** | **bool** | Whether text is less than the whole page: the body or the extracted text hit a limit of the provider&#39;s, or max_chars cut it. | 
**Url** | **string** | The URL fetched, as asked, in canonical form. | 

## Methods

### NewFetchResponse

`func NewFetchResponse(fetchTool string, finalUrl string, text string, truncated bool, url string, ) *FetchResponse`

NewFetchResponse instantiates a new FetchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFetchResponseWithDefaults

`func NewFetchResponseWithDefaults() *FetchResponse`

NewFetchResponseWithDefaults instantiates a new FetchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentType

`func (o *FetchResponse) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *FetchResponse) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *FetchResponse) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *FetchResponse) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetFetchTool

`func (o *FetchResponse) GetFetchTool() string`

GetFetchTool returns the FetchTool field if non-nil, zero value otherwise.

### GetFetchToolOk

`func (o *FetchResponse) GetFetchToolOk() (*string, bool)`

GetFetchToolOk returns a tuple with the FetchTool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFetchTool

`func (o *FetchResponse) SetFetchTool(v string)`

SetFetchTool sets FetchTool field to given value.


### GetFinalUrl

`func (o *FetchResponse) GetFinalUrl() string`

GetFinalUrl returns the FinalUrl field if non-nil, zero value otherwise.

### GetFinalUrlOk

`func (o *FetchResponse) GetFinalUrlOk() (*string, bool)`

GetFinalUrlOk returns a tuple with the FinalUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalUrl

`func (o *FetchResponse) SetFinalUrl(v string)`

SetFinalUrl sets FinalUrl field to given value.


### GetObject

`func (o *FetchResponse) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *FetchResponse) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *FetchResponse) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *FetchResponse) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetPublished

`func (o *FetchResponse) GetPublished() time.Time`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *FetchResponse) GetPublishedOk() (*time.Time, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *FetchResponse) SetPublished(v time.Time)`

SetPublished sets Published field to given value.

### HasPublished

`func (o *FetchResponse) HasPublished() bool`

HasPublished returns a boolean if a field has been set.

### SetPublishedNil

`func (o *FetchResponse) SetPublishedNil(b bool)`

 SetPublishedNil sets the value for Published to be an explicit nil

### UnsetPublished
`func (o *FetchResponse) UnsetPublished()`

UnsetPublished ensures that no value is present for Published, not even an explicit nil
### GetRaw

`func (o *FetchResponse) GetRaw() map[string]interface{}`

GetRaw returns the Raw field if non-nil, zero value otherwise.

### GetRawOk

`func (o *FetchResponse) GetRawOk() (*map[string]interface{}, bool)`

GetRawOk returns a tuple with the Raw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaw

`func (o *FetchResponse) SetRaw(v map[string]interface{})`

SetRaw sets Raw field to given value.

### HasRaw

`func (o *FetchResponse) HasRaw() bool`

HasRaw returns a boolean if a field has been set.

### SetRawNil

`func (o *FetchResponse) SetRawNil(b bool)`

 SetRawNil sets the value for Raw to be an explicit nil

### UnsetRaw
`func (o *FetchResponse) UnsetRaw()`

UnsetRaw ensures that no value is present for Raw, not even an explicit nil
### GetText

`func (o *FetchResponse) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *FetchResponse) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *FetchResponse) SetText(v string)`

SetText sets Text field to given value.


### GetTitle

`func (o *FetchResponse) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *FetchResponse) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *FetchResponse) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *FetchResponse) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetTruncated

`func (o *FetchResponse) GetTruncated() bool`

GetTruncated returns the Truncated field if non-nil, zero value otherwise.

### GetTruncatedOk

`func (o *FetchResponse) GetTruncatedOk() (*bool, bool)`

GetTruncatedOk returns a tuple with the Truncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncated

`func (o *FetchResponse) SetTruncated(v bool)`

SetTruncated sets Truncated field to given value.


### GetUrl

`func (o *FetchResponse) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *FetchResponse) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *FetchResponse) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


