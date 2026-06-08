# MRWebFetchBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | [**MRDocumentBlock**](MRDocumentBlock.md) |  | 
**RetrievedAt** | Pointer to **NullableString** | Filter models by provider name | [optional] 
**Type** | **string** |  | 
**Url** | **string** |  | 

## Methods

### NewMRWebFetchBlock

`func NewMRWebFetchBlock(content MRDocumentBlock, type_ string, url string, ) *MRWebFetchBlock`

NewMRWebFetchBlock instantiates a new MRWebFetchBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMRWebFetchBlockWithDefaults

`func NewMRWebFetchBlockWithDefaults() *MRWebFetchBlock`

NewMRWebFetchBlockWithDefaults instantiates a new MRWebFetchBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *MRWebFetchBlock) GetContent() MRDocumentBlock`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *MRWebFetchBlock) GetContentOk() (*MRDocumentBlock, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *MRWebFetchBlock) SetContent(v MRDocumentBlock)`

SetContent sets Content field to given value.


### GetRetrievedAt

`func (o *MRWebFetchBlock) GetRetrievedAt() string`

GetRetrievedAt returns the RetrievedAt field if non-nil, zero value otherwise.

### GetRetrievedAtOk

`func (o *MRWebFetchBlock) GetRetrievedAtOk() (*string, bool)`

GetRetrievedAtOk returns a tuple with the RetrievedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetrievedAt

`func (o *MRWebFetchBlock) SetRetrievedAt(v string)`

SetRetrievedAt sets RetrievedAt field to given value.

### HasRetrievedAt

`func (o *MRWebFetchBlock) HasRetrievedAt() bool`

HasRetrievedAt returns a boolean if a field has been set.

### SetRetrievedAtNil

`func (o *MRWebFetchBlock) SetRetrievedAtNil(b bool)`

 SetRetrievedAtNil sets the value for RetrievedAt to be an explicit nil

### UnsetRetrievedAt
`func (o *MRWebFetchBlock) UnsetRetrievedAt()`

UnsetRetrievedAt ensures that no value is present for RetrievedAt, not even an explicit nil
### GetType

`func (o *MRWebFetchBlock) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MRWebFetchBlock) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MRWebFetchBlock) SetType(v string)`

SetType sets Type field to given value.


### GetUrl

`func (o *MRWebFetchBlock) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *MRWebFetchBlock) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *MRWebFetchBlock) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


