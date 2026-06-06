# RerankRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Documents** | **[]string** | List of document strings to rerank | 
**MaxTokensPerDoc** | Pointer to **NullableInt32** | Per-document truncation limit | [optional] 
**Model** | **string** | Provider-prefixed model ID, e.g. &#39;cohere:rerank-v3.5&#39; | 
**Query** | **string** | The search query to rerank documents against | 
**TopN** | Pointer to **NullableInt32** | Maximum number of results to return | [optional] 
**User** | Pointer to **NullableString** | User ID for usage attribution | [optional] 

## Methods

### NewRerankRequest

`func NewRerankRequest(documents []string, model string, query string, ) *RerankRequest`

NewRerankRequest instantiates a new RerankRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRerankRequestWithDefaults

`func NewRerankRequestWithDefaults() *RerankRequest`

NewRerankRequestWithDefaults instantiates a new RerankRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocuments

`func (o *RerankRequest) GetDocuments() []string`

GetDocuments returns the Documents field if non-nil, zero value otherwise.

### GetDocumentsOk

`func (o *RerankRequest) GetDocumentsOk() (*[]string, bool)`

GetDocumentsOk returns a tuple with the Documents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocuments

`func (o *RerankRequest) SetDocuments(v []string)`

SetDocuments sets Documents field to given value.


### GetMaxTokensPerDoc

`func (o *RerankRequest) GetMaxTokensPerDoc() int32`

GetMaxTokensPerDoc returns the MaxTokensPerDoc field if non-nil, zero value otherwise.

### GetMaxTokensPerDocOk

`func (o *RerankRequest) GetMaxTokensPerDocOk() (*int32, bool)`

GetMaxTokensPerDocOk returns a tuple with the MaxTokensPerDoc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTokensPerDoc

`func (o *RerankRequest) SetMaxTokensPerDoc(v int32)`

SetMaxTokensPerDoc sets MaxTokensPerDoc field to given value.

### HasMaxTokensPerDoc

`func (o *RerankRequest) HasMaxTokensPerDoc() bool`

HasMaxTokensPerDoc returns a boolean if a field has been set.

### SetMaxTokensPerDocNil

`func (o *RerankRequest) SetMaxTokensPerDocNil(b bool)`

 SetMaxTokensPerDocNil sets the value for MaxTokensPerDoc to be an explicit nil

### UnsetMaxTokensPerDoc
`func (o *RerankRequest) UnsetMaxTokensPerDoc()`

UnsetMaxTokensPerDoc ensures that no value is present for MaxTokensPerDoc, not even an explicit nil
### GetModel

`func (o *RerankRequest) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *RerankRequest) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *RerankRequest) SetModel(v string)`

SetModel sets Model field to given value.


### GetQuery

`func (o *RerankRequest) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *RerankRequest) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *RerankRequest) SetQuery(v string)`

SetQuery sets Query field to given value.


### GetTopN

`func (o *RerankRequest) GetTopN() int32`

GetTopN returns the TopN field if non-nil, zero value otherwise.

### GetTopNOk

`func (o *RerankRequest) GetTopNOk() (*int32, bool)`

GetTopNOk returns a tuple with the TopN field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopN

`func (o *RerankRequest) SetTopN(v int32)`

SetTopN sets TopN field to given value.

### HasTopN

`func (o *RerankRequest) HasTopN() bool`

HasTopN returns a boolean if a field has been set.

### SetTopNNil

`func (o *RerankRequest) SetTopNNil(b bool)`

 SetTopNNil sets the value for TopN to be an explicit nil

### UnsetTopN
`func (o *RerankRequest) UnsetTopN()`

UnsetTopN ensures that no value is present for TopN, not even an explicit nil
### GetUser

`func (o *RerankRequest) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *RerankRequest) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *RerankRequest) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *RerankRequest) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *RerankRequest) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *RerankRequest) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


