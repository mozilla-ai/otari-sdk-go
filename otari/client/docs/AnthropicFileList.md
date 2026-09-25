# AnthropicFileList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]AnthropicFileMetadata**](AnthropicFileMetadata.md) |  | 
**NextPage** | **NullableString** |  | 

## Methods

### NewAnthropicFileList

`func NewAnthropicFileList(data []AnthropicFileMetadata, nextPage NullableString, ) *AnthropicFileList`

NewAnthropicFileList instantiates a new AnthropicFileList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnthropicFileListWithDefaults

`func NewAnthropicFileListWithDefaults() *AnthropicFileList`

NewAnthropicFileListWithDefaults instantiates a new AnthropicFileList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AnthropicFileList) GetData() []AnthropicFileMetadata`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AnthropicFileList) GetDataOk() (*[]AnthropicFileMetadata, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AnthropicFileList) SetData(v []AnthropicFileMetadata)`

SetData sets Data field to given value.


### GetNextPage

`func (o *AnthropicFileList) GetNextPage() string`

GetNextPage returns the NextPage field if non-nil, zero value otherwise.

### GetNextPageOk

`func (o *AnthropicFileList) GetNextPageOk() (*string, bool)`

GetNextPageOk returns a tuple with the NextPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPage

`func (o *AnthropicFileList) SetNextPage(v string)`

SetNextPage sets NextPage field to given value.


### SetNextPageNil

`func (o *AnthropicFileList) SetNextPageNil(b bool)`

 SetNextPageNil sets the value for NextPage to be an explicit nil

### UnsetNextPage
`func (o *AnthropicFileList) UnsetNextPage()`

UnsetNextPage ensures that no value is present for NextPage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


