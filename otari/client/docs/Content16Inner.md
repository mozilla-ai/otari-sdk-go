# Content16Inner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Signature** | **string** |  | 
**Thinking** | **string** |  | 
**Type** | **string** |  | 
**Citations** | Pointer to [**[]MRBetaTextBlockCitationsInner**](MRBetaTextBlockCitationsInner.md) |  | [optional] 
**Text** | **string** |  | 
**Data** | **string** |  | 
**Id** | **string** |  | 
**Caller** | Pointer to [**NullableCaller**](Caller.md) |  | [optional] 
**Input** | **map[string]interface{}** |  | 
**Name** | **string** |  | 
**Content** | **string** |  | 
**ToolUseId** | **string** |  | 
**FileId** | **string** |  | 
**ServerName** | **string** |  | 
**IsError** | **bool** |  | 

## Methods

### NewContent16Inner

`func NewContent16Inner(signature string, thinking string, type_ string, text string, data string, id string, input map[string]interface{}, name string, content string, toolUseId string, fileId string, serverName string, isError bool, ) *Content16Inner`

NewContent16Inner instantiates a new Content16Inner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContent16InnerWithDefaults

`func NewContent16InnerWithDefaults() *Content16Inner`

NewContent16InnerWithDefaults instantiates a new Content16Inner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSignature

`func (o *Content16Inner) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *Content16Inner) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *Content16Inner) SetSignature(v string)`

SetSignature sets Signature field to given value.


### GetThinking

`func (o *Content16Inner) GetThinking() string`

GetThinking returns the Thinking field if non-nil, zero value otherwise.

### GetThinkingOk

`func (o *Content16Inner) GetThinkingOk() (*string, bool)`

GetThinkingOk returns a tuple with the Thinking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThinking

`func (o *Content16Inner) SetThinking(v string)`

SetThinking sets Thinking field to given value.


### GetType

`func (o *Content16Inner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Content16Inner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Content16Inner) SetType(v string)`

SetType sets Type field to given value.


### GetCitations

`func (o *Content16Inner) GetCitations() []MRBetaTextBlockCitationsInner`

GetCitations returns the Citations field if non-nil, zero value otherwise.

### GetCitationsOk

`func (o *Content16Inner) GetCitationsOk() (*[]MRBetaTextBlockCitationsInner, bool)`

GetCitationsOk returns a tuple with the Citations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitations

`func (o *Content16Inner) SetCitations(v []MRBetaTextBlockCitationsInner)`

SetCitations sets Citations field to given value.

### HasCitations

`func (o *Content16Inner) HasCitations() bool`

HasCitations returns a boolean if a field has been set.

### GetText

`func (o *Content16Inner) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *Content16Inner) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *Content16Inner) SetText(v string)`

SetText sets Text field to given value.


### GetData

`func (o *Content16Inner) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Content16Inner) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Content16Inner) SetData(v string)`

SetData sets Data field to given value.


### GetId

`func (o *Content16Inner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Content16Inner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Content16Inner) SetId(v string)`

SetId sets Id field to given value.


### GetCaller

`func (o *Content16Inner) GetCaller() Caller`

GetCaller returns the Caller field if non-nil, zero value otherwise.

### GetCallerOk

`func (o *Content16Inner) GetCallerOk() (*Caller, bool)`

GetCallerOk returns a tuple with the Caller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaller

`func (o *Content16Inner) SetCaller(v Caller)`

SetCaller sets Caller field to given value.

### HasCaller

`func (o *Content16Inner) HasCaller() bool`

HasCaller returns a boolean if a field has been set.

### SetCallerNil

`func (o *Content16Inner) SetCallerNil(b bool)`

 SetCallerNil sets the value for Caller to be an explicit nil

### UnsetCaller
`func (o *Content16Inner) UnsetCaller()`

UnsetCaller ensures that no value is present for Caller, not even an explicit nil
### GetInput

`func (o *Content16Inner) GetInput() map[string]interface{}`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *Content16Inner) GetInputOk() (*map[string]interface{}, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *Content16Inner) SetInput(v map[string]interface{})`

SetInput sets Input field to given value.


### GetName

`func (o *Content16Inner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Content16Inner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Content16Inner) SetName(v string)`

SetName sets Name field to given value.


### GetContent

`func (o *Content16Inner) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *Content16Inner) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *Content16Inner) SetContent(v string)`

SetContent sets Content field to given value.


### GetToolUseId

`func (o *Content16Inner) GetToolUseId() string`

GetToolUseId returns the ToolUseId field if non-nil, zero value otherwise.

### GetToolUseIdOk

`func (o *Content16Inner) GetToolUseIdOk() (*string, bool)`

GetToolUseIdOk returns a tuple with the ToolUseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolUseId

`func (o *Content16Inner) SetToolUseId(v string)`

SetToolUseId sets ToolUseId field to given value.


### GetFileId

`func (o *Content16Inner) GetFileId() string`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *Content16Inner) GetFileIdOk() (*string, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *Content16Inner) SetFileId(v string)`

SetFileId sets FileId field to given value.


### GetServerName

`func (o *Content16Inner) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *Content16Inner) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *Content16Inner) SetServerName(v string)`

SetServerName sets ServerName field to given value.


### GetIsError

`func (o *Content16Inner) GetIsError() bool`

GetIsError returns the IsError field if non-nil, zero value otherwise.

### GetIsErrorOk

`func (o *Content16Inner) GetIsErrorOk() (*bool, bool)`

GetIsErrorOk returns a tuple with the IsError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsError

`func (o *Content16Inner) SetIsError(v bool)`

SetIsError sets IsError field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


