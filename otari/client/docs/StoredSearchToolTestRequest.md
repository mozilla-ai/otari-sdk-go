# StoredSearchToolTestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Query** | Pointer to **NullableString** | For a search instance: the query to run. | [optional] 
**Url** | Pointer to **NullableString** | For a fetch instance: the page to fetch. | [optional] 

## Methods

### NewStoredSearchToolTestRequest

`func NewStoredSearchToolTestRequest() *StoredSearchToolTestRequest`

NewStoredSearchToolTestRequest instantiates a new StoredSearchToolTestRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoredSearchToolTestRequestWithDefaults

`func NewStoredSearchToolTestRequestWithDefaults() *StoredSearchToolTestRequest`

NewStoredSearchToolTestRequestWithDefaults instantiates a new StoredSearchToolTestRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQuery

`func (o *StoredSearchToolTestRequest) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *StoredSearchToolTestRequest) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *StoredSearchToolTestRequest) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *StoredSearchToolTestRequest) HasQuery() bool`

HasQuery returns a boolean if a field has been set.

### SetQueryNil

`func (o *StoredSearchToolTestRequest) SetQueryNil(b bool)`

 SetQueryNil sets the value for Query to be an explicit nil

### UnsetQuery
`func (o *StoredSearchToolTestRequest) UnsetQuery()`

UnsetQuery ensures that no value is present for Query, not even an explicit nil
### GetUrl

`func (o *StoredSearchToolTestRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *StoredSearchToolTestRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *StoredSearchToolTestRequest) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *StoredSearchToolTestRequest) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *StoredSearchToolTestRequest) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *StoredSearchToolTestRequest) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


