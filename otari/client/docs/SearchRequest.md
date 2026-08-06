# SearchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Country** | Pointer to **NullableString** | Two-letter ISO country code to localize results to | [optional] 
**MaxResults** | Pointer to **NullableInt32** | Maximum number of results to return | [optional] 
**MaxTokensPerPage** | Pointer to **NullableInt32** | Approximate cap on the page content returned per result | [optional] 
**Query** | **string** | The search query | 
**SearchDomainFilter** | Pointer to **[]string** | Restrict results to these domains; prefix a domain with &#39;-&#39; to exclude it instead | [optional] 
**SearchToolName** | Pointer to **NullableString** | Configured search tool to run against. Optional when exactly one tool is configured, and ignored on POST /v1/search/{search_tool_name}. | [optional] 
**User** | Pointer to **NullableString** | User ID for usage attribution | [optional] 

## Methods

### NewSearchRequest

`func NewSearchRequest(query string, ) *SearchRequest`

NewSearchRequest instantiates a new SearchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchRequestWithDefaults

`func NewSearchRequestWithDefaults() *SearchRequest`

NewSearchRequestWithDefaults instantiates a new SearchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCountry

`func (o *SearchRequest) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *SearchRequest) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *SearchRequest) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *SearchRequest) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *SearchRequest) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *SearchRequest) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetMaxResults

`func (o *SearchRequest) GetMaxResults() int32`

GetMaxResults returns the MaxResults field if non-nil, zero value otherwise.

### GetMaxResultsOk

`func (o *SearchRequest) GetMaxResultsOk() (*int32, bool)`

GetMaxResultsOk returns a tuple with the MaxResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxResults

`func (o *SearchRequest) SetMaxResults(v int32)`

SetMaxResults sets MaxResults field to given value.

### HasMaxResults

`func (o *SearchRequest) HasMaxResults() bool`

HasMaxResults returns a boolean if a field has been set.

### SetMaxResultsNil

`func (o *SearchRequest) SetMaxResultsNil(b bool)`

 SetMaxResultsNil sets the value for MaxResults to be an explicit nil

### UnsetMaxResults
`func (o *SearchRequest) UnsetMaxResults()`

UnsetMaxResults ensures that no value is present for MaxResults, not even an explicit nil
### GetMaxTokensPerPage

`func (o *SearchRequest) GetMaxTokensPerPage() int32`

GetMaxTokensPerPage returns the MaxTokensPerPage field if non-nil, zero value otherwise.

### GetMaxTokensPerPageOk

`func (o *SearchRequest) GetMaxTokensPerPageOk() (*int32, bool)`

GetMaxTokensPerPageOk returns a tuple with the MaxTokensPerPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTokensPerPage

`func (o *SearchRequest) SetMaxTokensPerPage(v int32)`

SetMaxTokensPerPage sets MaxTokensPerPage field to given value.

### HasMaxTokensPerPage

`func (o *SearchRequest) HasMaxTokensPerPage() bool`

HasMaxTokensPerPage returns a boolean if a field has been set.

### SetMaxTokensPerPageNil

`func (o *SearchRequest) SetMaxTokensPerPageNil(b bool)`

 SetMaxTokensPerPageNil sets the value for MaxTokensPerPage to be an explicit nil

### UnsetMaxTokensPerPage
`func (o *SearchRequest) UnsetMaxTokensPerPage()`

UnsetMaxTokensPerPage ensures that no value is present for MaxTokensPerPage, not even an explicit nil
### GetQuery

`func (o *SearchRequest) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *SearchRequest) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *SearchRequest) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetSearchDomainFilter

`func (o *SearchRequest) GetSearchDomainFilter() []string`

GetSearchDomainFilter returns the SearchDomainFilter field if non-nil, zero value otherwise.

### GetSearchDomainFilterOk

`func (o *SearchRequest) GetSearchDomainFilterOk() (*[]string, bool)`

GetSearchDomainFilterOk returns a tuple with the SearchDomainFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchDomainFilter

`func (o *SearchRequest) SetSearchDomainFilter(v []string)`

SetSearchDomainFilter sets SearchDomainFilter field to given value.

### HasSearchDomainFilter

`func (o *SearchRequest) HasSearchDomainFilter() bool`

HasSearchDomainFilter returns a boolean if a field has been set.

### SetSearchDomainFilterNil

`func (o *SearchRequest) SetSearchDomainFilterNil(b bool)`

 SetSearchDomainFilterNil sets the value for SearchDomainFilter to be an explicit nil

### UnsetSearchDomainFilter
`func (o *SearchRequest) UnsetSearchDomainFilter()`

UnsetSearchDomainFilter ensures that no value is present for SearchDomainFilter, not even an explicit nil
### GetSearchToolName

`func (o *SearchRequest) GetSearchToolName() string`

GetSearchToolName returns the SearchToolName field if non-nil, zero value otherwise.

### GetSearchToolNameOk

`func (o *SearchRequest) GetSearchToolNameOk() (*string, bool)`

GetSearchToolNameOk returns a tuple with the SearchToolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchToolName

`func (o *SearchRequest) SetSearchToolName(v string)`

SetSearchToolName sets SearchToolName field to given value.

### HasSearchToolName

`func (o *SearchRequest) HasSearchToolName() bool`

HasSearchToolName returns a boolean if a field has been set.

### SetSearchToolNameNil

`func (o *SearchRequest) SetSearchToolNameNil(b bool)`

 SetSearchToolNameNil sets the value for SearchToolName to be an explicit nil

### UnsetSearchToolName
`func (o *SearchRequest) UnsetSearchToolName()`

UnsetSearchToolName ensures that no value is present for SearchToolName, not even an explicit nil
### GetUser

`func (o *SearchRequest) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *SearchRequest) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *SearchRequest) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *SearchRequest) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *SearchRequest) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *SearchRequest) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


