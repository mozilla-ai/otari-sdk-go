# MSGChatCompletionMessageCustomToolCallParam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Custom** | [**MSGCustom**](MSGCustom.md) |  | 
**Type** | **string** |  | 

## Methods

### NewMSGChatCompletionMessageCustomToolCallParam

`func NewMSGChatCompletionMessageCustomToolCallParam(id string, custom MSGCustom, type_ string, ) *MSGChatCompletionMessageCustomToolCallParam`

NewMSGChatCompletionMessageCustomToolCallParam instantiates a new MSGChatCompletionMessageCustomToolCallParam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMSGChatCompletionMessageCustomToolCallParamWithDefaults

`func NewMSGChatCompletionMessageCustomToolCallParamWithDefaults() *MSGChatCompletionMessageCustomToolCallParam`

NewMSGChatCompletionMessageCustomToolCallParamWithDefaults instantiates a new MSGChatCompletionMessageCustomToolCallParam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MSGChatCompletionMessageCustomToolCallParam) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MSGChatCompletionMessageCustomToolCallParam) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MSGChatCompletionMessageCustomToolCallParam) SetId(v string)`

SetId sets Id field to given value.


### GetCustom

`func (o *MSGChatCompletionMessageCustomToolCallParam) GetCustom() MSGCustom`

GetCustom returns the Custom field if non-nil, zero value otherwise.

### GetCustomOk

`func (o *MSGChatCompletionMessageCustomToolCallParam) GetCustomOk() (*MSGCustom, bool)`

GetCustomOk returns a tuple with the Custom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustom

`func (o *MSGChatCompletionMessageCustomToolCallParam) SetCustom(v MSGCustom)`

SetCustom sets Custom field to given value.


### GetType

`func (o *MSGChatCompletionMessageCustomToolCallParam) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MSGChatCompletionMessageCustomToolCallParam) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MSGChatCompletionMessageCustomToolCallParam) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


