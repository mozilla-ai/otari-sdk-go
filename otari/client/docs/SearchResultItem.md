# SearchResultItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | Pointer to **NullableString** |  | [optional] 
**Snippet** | Pointer to **NullableString** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Url** | **string** |  | 

## Methods

### NewSearchResultItem

`func NewSearchResultItem(url string, ) *SearchResultItem`

NewSearchResultItem instantiates a new SearchResultItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchResultItemWithDefaults

`func NewSearchResultItemWithDefaults() *SearchResultItem`

NewSearchResultItemWithDefaults instantiates a new SearchResultItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *SearchResultItem) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *SearchResultItem) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *SearchResultItem) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *SearchResultItem) HasDate() bool`

HasDate returns a boolean if a field has been set.

### SetDateNil

`func (o *SearchResultItem) SetDateNil(b bool)`

 SetDateNil sets the value for Date to be an explicit nil

### UnsetDate
`func (o *SearchResultItem) UnsetDate()`

UnsetDate ensures that no value is present for Date, not even an explicit nil
### GetSnippet

`func (o *SearchResultItem) GetSnippet() string`

GetSnippet returns the Snippet field if non-nil, zero value otherwise.

### GetSnippetOk

`func (o *SearchResultItem) GetSnippetOk() (*string, bool)`

GetSnippetOk returns a tuple with the Snippet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnippet

`func (o *SearchResultItem) SetSnippet(v string)`

SetSnippet sets Snippet field to given value.

### HasSnippet

`func (o *SearchResultItem) HasSnippet() bool`

HasSnippet returns a boolean if a field has been set.

### SetSnippetNil

`func (o *SearchResultItem) SetSnippetNil(b bool)`

 SetSnippetNil sets the value for Snippet to be an explicit nil

### UnsetSnippet
`func (o *SearchResultItem) UnsetSnippet()`

UnsetSnippet ensures that no value is present for Snippet, not even an explicit nil
### GetTitle

`func (o *SearchResultItem) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SearchResultItem) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SearchResultItem) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *SearchResultItem) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *SearchResultItem) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *SearchResultItem) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetUrl

`func (o *SearchResultItem) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *SearchResultItem) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *SearchResultItem) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


