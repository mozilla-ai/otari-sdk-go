# MRBetaCompactionBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to **NullableString** |  | [optional] 
**EncryptedContent** | Pointer to **NullableString** | Filter to a single event type or metric name (e.g. &#39;tool_result&#39;, &#39;claude_code.commit.count&#39;) | [optional] 
**Type** | **string** |  | 

## Methods

### NewMRBetaCompactionBlock

`func NewMRBetaCompactionBlock(type_ string, ) *MRBetaCompactionBlock`

NewMRBetaCompactionBlock instantiates a new MRBetaCompactionBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMRBetaCompactionBlockWithDefaults

`func NewMRBetaCompactionBlockWithDefaults() *MRBetaCompactionBlock`

NewMRBetaCompactionBlockWithDefaults instantiates a new MRBetaCompactionBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *MRBetaCompactionBlock) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *MRBetaCompactionBlock) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *MRBetaCompactionBlock) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *MRBetaCompactionBlock) HasContent() bool`

HasContent returns a boolean if a field has been set.

### SetContentNil

`func (o *MRBetaCompactionBlock) SetContentNil(b bool)`

 SetContentNil sets the value for Content to be an explicit nil

### UnsetContent
`func (o *MRBetaCompactionBlock) UnsetContent()`

UnsetContent ensures that no value is present for Content, not even an explicit nil
### GetEncryptedContent

`func (o *MRBetaCompactionBlock) GetEncryptedContent() string`

GetEncryptedContent returns the EncryptedContent field if non-nil, zero value otherwise.

### GetEncryptedContentOk

`func (o *MRBetaCompactionBlock) GetEncryptedContentOk() (*string, bool)`

GetEncryptedContentOk returns a tuple with the EncryptedContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedContent

`func (o *MRBetaCompactionBlock) SetEncryptedContent(v string)`

SetEncryptedContent sets EncryptedContent field to given value.

### HasEncryptedContent

`func (o *MRBetaCompactionBlock) HasEncryptedContent() bool`

HasEncryptedContent returns a boolean if a field has been set.

### SetEncryptedContentNil

`func (o *MRBetaCompactionBlock) SetEncryptedContentNil(b bool)`

 SetEncryptedContentNil sets the value for EncryptedContent to be an explicit nil

### UnsetEncryptedContent
`func (o *MRBetaCompactionBlock) UnsetEncryptedContent()`

UnsetEncryptedContent ensures that no value is present for EncryptedContent, not even an explicit nil
### GetType

`func (o *MRBetaCompactionBlock) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MRBetaCompactionBlock) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MRBetaCompactionBlock) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


