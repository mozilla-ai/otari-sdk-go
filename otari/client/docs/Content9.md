# Content9

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorCode** | **string** |  | 
**Type** | **string** |  | 
**Content** | [**MRBetaDocumentBlock**](MRBetaDocumentBlock.md) |  | 
**RetrievedAt** | Pointer to **string** |  | [optional] 
**Url** | **string** |  | 

## Methods

### NewContent9

`func NewContent9(errorCode string, type_ string, content MRBetaDocumentBlock, url string, ) *Content9`

NewContent9 instantiates a new Content9 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContent9WithDefaults

`func NewContent9WithDefaults() *Content9`

NewContent9WithDefaults instantiates a new Content9 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorCode

`func (o *Content9) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *Content9) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *Content9) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.


### GetType

`func (o *Content9) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Content9) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Content9) SetType(v string)`

SetType sets Type field to given value.


### GetContent

`func (o *Content9) GetContent() MRBetaDocumentBlock`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *Content9) GetContentOk() (*MRBetaDocumentBlock, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *Content9) SetContent(v MRBetaDocumentBlock)`

SetContent sets Content field to given value.


### GetRetrievedAt

`func (o *Content9) GetRetrievedAt() string`

GetRetrievedAt returns the RetrievedAt field if non-nil, zero value otherwise.

### GetRetrievedAtOk

`func (o *Content9) GetRetrievedAtOk() (*string, bool)`

GetRetrievedAtOk returns a tuple with the RetrievedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetrievedAt

`func (o *Content9) SetRetrievedAt(v string)`

SetRetrievedAt sets RetrievedAt field to given value.

### HasRetrievedAt

`func (o *Content9) HasRetrievedAt() bool`

HasRetrievedAt returns a boolean if a field has been set.

### GetUrl

`func (o *Content9) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Content9) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Content9) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


