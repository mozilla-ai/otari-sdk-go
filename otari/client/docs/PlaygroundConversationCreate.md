# PlaygroundConversationCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Messages** | [**[]PlaygroundMessageCreate**](PlaygroundMessageCreate.md) |  | 
**Model** | **string** |  | 
**Title** | **string** |  | 
**WorkspaceId** | **string** |  | 

## Methods

### NewPlaygroundConversationCreate

`func NewPlaygroundConversationCreate(messages []PlaygroundMessageCreate, model string, title string, workspaceId string, ) *PlaygroundConversationCreate`

NewPlaygroundConversationCreate instantiates a new PlaygroundConversationCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlaygroundConversationCreateWithDefaults

`func NewPlaygroundConversationCreateWithDefaults() *PlaygroundConversationCreate`

NewPlaygroundConversationCreateWithDefaults instantiates a new PlaygroundConversationCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessages

`func (o *PlaygroundConversationCreate) GetMessages() []PlaygroundMessageCreate`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *PlaygroundConversationCreate) GetMessagesOk() (*[]PlaygroundMessageCreate, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *PlaygroundConversationCreate) SetMessages(v []PlaygroundMessageCreate)`

SetMessages sets Messages field to given value.


### GetModel

`func (o *PlaygroundConversationCreate) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *PlaygroundConversationCreate) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *PlaygroundConversationCreate) SetModel(v string)`

SetModel sets Model field to given value.


### GetTitle

`func (o *PlaygroundConversationCreate) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PlaygroundConversationCreate) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PlaygroundConversationCreate) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetWorkspaceId

`func (o *PlaygroundConversationCreate) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *PlaygroundConversationCreate) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *PlaygroundConversationCreate) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


