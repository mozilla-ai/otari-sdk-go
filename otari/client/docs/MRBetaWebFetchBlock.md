# MRBetaWebFetchBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | [**MRBetaDocumentBlock**](MRBetaDocumentBlock.md) |  | 
**RetrievedAt** | Pointer to **NullableString** |  | [optional] 
**Type** | **string** |  | 
**Url** | **string** |  | 

## Methods

### NewMRBetaWebFetchBlock

`func NewMRBetaWebFetchBlock(content MRBetaDocumentBlock, type_ string, url string, ) *MRBetaWebFetchBlock`

NewMRBetaWebFetchBlock instantiates a new MRBetaWebFetchBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMRBetaWebFetchBlockWithDefaults

`func NewMRBetaWebFetchBlockWithDefaults() *MRBetaWebFetchBlock`

NewMRBetaWebFetchBlockWithDefaults instantiates a new MRBetaWebFetchBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *MRBetaWebFetchBlock) GetContent() MRBetaDocumentBlock`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *MRBetaWebFetchBlock) GetContentOk() (*MRBetaDocumentBlock, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *MRBetaWebFetchBlock) SetContent(v MRBetaDocumentBlock)`

SetContent sets Content field to given value.


### GetRetrievedAt

`func (o *MRBetaWebFetchBlock) GetRetrievedAt() string`

GetRetrievedAt returns the RetrievedAt field if non-nil, zero value otherwise.

### GetRetrievedAtOk

`func (o *MRBetaWebFetchBlock) GetRetrievedAtOk() (*string, bool)`

GetRetrievedAtOk returns a tuple with the RetrievedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetrievedAt

`func (o *MRBetaWebFetchBlock) SetRetrievedAt(v string)`

SetRetrievedAt sets RetrievedAt field to given value.

### HasRetrievedAt

`func (o *MRBetaWebFetchBlock) HasRetrievedAt() bool`

HasRetrievedAt returns a boolean if a field has been set.

### SetRetrievedAtNil

`func (o *MRBetaWebFetchBlock) SetRetrievedAtNil(b bool)`

 SetRetrievedAtNil sets the value for RetrievedAt to be an explicit nil

### UnsetRetrievedAt
`func (o *MRBetaWebFetchBlock) UnsetRetrievedAt()`

UnsetRetrievedAt ensures that no value is present for RetrievedAt, not even an explicit nil
### GetType

`func (o *MRBetaWebFetchBlock) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MRBetaWebFetchBlock) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MRBetaWebFetchBlock) SetType(v string)`

SetType sets Type field to given value.


### GetUrl

`func (o *MRBetaWebFetchBlock) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *MRBetaWebFetchBlock) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *MRBetaWebFetchBlock) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


