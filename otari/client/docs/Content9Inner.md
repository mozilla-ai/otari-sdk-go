# Content9Inner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Citations** | Pointer to [**[]MRTextBlockCitationsInner**](MRTextBlockCitationsInner.md) |  | [optional] 
**Text** | **string** |  | 
**Type** | **string** |  | 
**Signature** | **string** |  | 
**Thinking** | **string** |  | 
**Data** | **string** |  | 
**Id** | **string** |  | 
**Caller** | Pointer to [**NullableCaller**](Caller.md) |  | [optional] 
**Input** | **map[string]interface{}** |  | 
**Name** | **string** |  | 
**Content** | [**Content6**](Content6.md) |  | 
**ToolUseId** | **string** |  | 
**FileId** | **string** |  | 

## Methods

### NewContent9Inner

`func NewContent9Inner(text string, type_ string, signature string, thinking string, data string, id string, input map[string]interface{}, name string, content Content6, toolUseId string, fileId string, ) *Content9Inner`

NewContent9Inner instantiates a new Content9Inner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContent9InnerWithDefaults

`func NewContent9InnerWithDefaults() *Content9Inner`

NewContent9InnerWithDefaults instantiates a new Content9Inner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCitations

`func (o *Content9Inner) GetCitations() []MRTextBlockCitationsInner`

GetCitations returns the Citations field if non-nil, zero value otherwise.

### GetCitationsOk

`func (o *Content9Inner) GetCitationsOk() (*[]MRTextBlockCitationsInner, bool)`

GetCitationsOk returns a tuple with the Citations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitations

`func (o *Content9Inner) SetCitations(v []MRTextBlockCitationsInner)`

SetCitations sets Citations field to given value.

### HasCitations

`func (o *Content9Inner) HasCitations() bool`

HasCitations returns a boolean if a field has been set.

### GetText

`func (o *Content9Inner) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *Content9Inner) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *Content9Inner) SetText(v string)`

SetText sets Text field to given value.


### GetType

`func (o *Content9Inner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Content9Inner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Content9Inner) SetType(v string)`

SetType sets Type field to given value.


### GetSignature

`func (o *Content9Inner) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *Content9Inner) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *Content9Inner) SetSignature(v string)`

SetSignature sets Signature field to given value.


### GetThinking

`func (o *Content9Inner) GetThinking() string`

GetThinking returns the Thinking field if non-nil, zero value otherwise.

### GetThinkingOk

`func (o *Content9Inner) GetThinkingOk() (*string, bool)`

GetThinkingOk returns a tuple with the Thinking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThinking

`func (o *Content9Inner) SetThinking(v string)`

SetThinking sets Thinking field to given value.


### GetData

`func (o *Content9Inner) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Content9Inner) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Content9Inner) SetData(v string)`

SetData sets Data field to given value.


### GetId

`func (o *Content9Inner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Content9Inner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Content9Inner) SetId(v string)`

SetId sets Id field to given value.


### GetCaller

`func (o *Content9Inner) GetCaller() Caller`

GetCaller returns the Caller field if non-nil, zero value otherwise.

### GetCallerOk

`func (o *Content9Inner) GetCallerOk() (*Caller, bool)`

GetCallerOk returns a tuple with the Caller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaller

`func (o *Content9Inner) SetCaller(v Caller)`

SetCaller sets Caller field to given value.

### HasCaller

`func (o *Content9Inner) HasCaller() bool`

HasCaller returns a boolean if a field has been set.

### SetCallerNil

`func (o *Content9Inner) SetCallerNil(b bool)`

 SetCallerNil sets the value for Caller to be an explicit nil

### UnsetCaller
`func (o *Content9Inner) UnsetCaller()`

UnsetCaller ensures that no value is present for Caller, not even an explicit nil
### GetInput

`func (o *Content9Inner) GetInput() map[string]interface{}`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *Content9Inner) GetInputOk() (*map[string]interface{}, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *Content9Inner) SetInput(v map[string]interface{})`

SetInput sets Input field to given value.


### GetName

`func (o *Content9Inner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Content9Inner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Content9Inner) SetName(v string)`

SetName sets Name field to given value.


### GetContent

`func (o *Content9Inner) GetContent() Content6`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *Content9Inner) GetContentOk() (*Content6, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *Content9Inner) SetContent(v Content6)`

SetContent sets Content field to given value.


### GetToolUseId

`func (o *Content9Inner) GetToolUseId() string`

GetToolUseId returns the ToolUseId field if non-nil, zero value otherwise.

### GetToolUseIdOk

`func (o *Content9Inner) GetToolUseIdOk() (*string, bool)`

GetToolUseIdOk returns a tuple with the ToolUseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolUseId

`func (o *Content9Inner) SetToolUseId(v string)`

SetToolUseId sets ToolUseId field to given value.


### GetFileId

`func (o *Content9Inner) GetFileId() string`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *Content9Inner) GetFileIdOk() (*string, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *Content9Inner) SetFileId(v string)`

SetFileId sets FileId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


