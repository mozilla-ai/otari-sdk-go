# SearchToolsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | [**[]ConfigSearchToolSchema**](ConfigSearchToolSchema.md) |  | 
**Stored** | [**[]StoredSearchToolSchema**](StoredSearchToolSchema.md) |  | 

## Methods

### NewSearchToolsResponse

`func NewSearchToolsResponse(config []ConfigSearchToolSchema, stored []StoredSearchToolSchema, ) *SearchToolsResponse`

NewSearchToolsResponse instantiates a new SearchToolsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchToolsResponseWithDefaults

`func NewSearchToolsResponseWithDefaults() *SearchToolsResponse`

NewSearchToolsResponseWithDefaults instantiates a new SearchToolsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *SearchToolsResponse) GetConfig() []ConfigSearchToolSchema`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *SearchToolsResponse) GetConfigOk() (*[]ConfigSearchToolSchema, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *SearchToolsResponse) SetConfig(v []ConfigSearchToolSchema)`

SetConfig sets Config field to given value.


### GetStored

`func (o *SearchToolsResponse) GetStored() []StoredSearchToolSchema`

GetStored returns the Stored field if non-nil, zero value otherwise.

### GetStoredOk

`func (o *SearchToolsResponse) GetStoredOk() (*[]StoredSearchToolSchema, bool)`

GetStoredOk returns a tuple with the Stored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStored

`func (o *SearchToolsResponse) SetStored(v []StoredSearchToolSchema)`

SetStored sets Stored field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


