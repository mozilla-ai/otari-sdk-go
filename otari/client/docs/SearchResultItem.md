# SearchResultItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | Pointer to **NullableString** | The page&#39;s date, as the provider gave it. | [optional] 
**LastUpdated** | Pointer to **NullableString** | Always null: no provider reports it. | [optional] 
**Snippet** | **string** | The page text the provider returned, else its excerpt, cut to the page budget. | 
**Title** | **string** | The page&#39;s title; empty when the provider gave none. | 
**Url** | **string** |  | 

## Methods

### NewSearchResultItem

`func NewSearchResultItem(snippet string, title string, url string, ) *SearchResultItem`

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
### GetLastUpdated

`func (o *SearchResultItem) GetLastUpdated() string`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *SearchResultItem) GetLastUpdatedOk() (*string, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *SearchResultItem) SetLastUpdated(v string)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *SearchResultItem) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### SetLastUpdatedNil

`func (o *SearchResultItem) SetLastUpdatedNil(b bool)`

 SetLastUpdatedNil sets the value for LastUpdated to be an explicit nil

### UnsetLastUpdated
`func (o *SearchResultItem) UnsetLastUpdated()`

UnsetLastUpdated ensures that no value is present for LastUpdated, not even an explicit nil
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


