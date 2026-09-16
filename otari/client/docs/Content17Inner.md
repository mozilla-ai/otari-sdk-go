# Content17Inner

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
**ToolsetName** | Pointer to **string** |  | [optional] 
**Content** | **string** |  | 
**ToolUseId** | **string** |  | 
**FileId** | **string** |  | 
**ServerName** | **string** |  | 
**IsError** | **bool** |  | 
**EncryptedContent** | Pointer to **string** | Filter to a single event type or metric name (e.g. &#39;tool_result&#39;, &#39;claude_code.commit.count&#39;) | [optional] 
**From** | [**MRBetaFallbackInfo**](MRBetaFallbackInfo.md) |  | 
**To** | [**MRBetaFallbackInfo**](MRBetaFallbackInfo.md) |  | 
**Trigger** | [**MRBetaFallbackRefusalTrigger**](MRBetaFallbackRefusalTrigger.md) |  | 

## Methods

### NewContent17Inner

`func NewContent17Inner(signature string, thinking string, type_ string, text string, data string, id string, input map[string]interface{}, name string, content string, toolUseId string, fileId string, serverName string, isError bool, from MRBetaFallbackInfo, to MRBetaFallbackInfo, trigger MRBetaFallbackRefusalTrigger, ) *Content17Inner`

NewContent17Inner instantiates a new Content17Inner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContent17InnerWithDefaults

`func NewContent17InnerWithDefaults() *Content17Inner`

NewContent17InnerWithDefaults instantiates a new Content17Inner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSignature

`func (o *Content17Inner) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *Content17Inner) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *Content17Inner) SetSignature(v string)`

SetSignature sets Signature field to given value.


### GetThinking

`func (o *Content17Inner) GetThinking() string`

GetThinking returns the Thinking field if non-nil, zero value otherwise.

### GetThinkingOk

`func (o *Content17Inner) GetThinkingOk() (*string, bool)`

GetThinkingOk returns a tuple with the Thinking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThinking

`func (o *Content17Inner) SetThinking(v string)`

SetThinking sets Thinking field to given value.


### GetType

`func (o *Content17Inner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Content17Inner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Content17Inner) SetType(v string)`

SetType sets Type field to given value.


### GetCitations

`func (o *Content17Inner) GetCitations() []MRBetaTextBlockCitationsInner`

GetCitations returns the Citations field if non-nil, zero value otherwise.

### GetCitationsOk

`func (o *Content17Inner) GetCitationsOk() (*[]MRBetaTextBlockCitationsInner, bool)`

GetCitationsOk returns a tuple with the Citations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCitations

`func (o *Content17Inner) SetCitations(v []MRBetaTextBlockCitationsInner)`

SetCitations sets Citations field to given value.

### HasCitations

`func (o *Content17Inner) HasCitations() bool`

HasCitations returns a boolean if a field has been set.

### GetText

`func (o *Content17Inner) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *Content17Inner) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *Content17Inner) SetText(v string)`

SetText sets Text field to given value.


### GetData

`func (o *Content17Inner) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *Content17Inner) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *Content17Inner) SetData(v string)`

SetData sets Data field to given value.


### GetId

`func (o *Content17Inner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Content17Inner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Content17Inner) SetId(v string)`

SetId sets Id field to given value.


### GetCaller

`func (o *Content17Inner) GetCaller() Caller`

GetCaller returns the Caller field if non-nil, zero value otherwise.

### GetCallerOk

`func (o *Content17Inner) GetCallerOk() (*Caller, bool)`

GetCallerOk returns a tuple with the Caller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaller

`func (o *Content17Inner) SetCaller(v Caller)`

SetCaller sets Caller field to given value.

### HasCaller

`func (o *Content17Inner) HasCaller() bool`

HasCaller returns a boolean if a field has been set.

### SetCallerNil

`func (o *Content17Inner) SetCallerNil(b bool)`

 SetCallerNil sets the value for Caller to be an explicit nil

### UnsetCaller
`func (o *Content17Inner) UnsetCaller()`

UnsetCaller ensures that no value is present for Caller, not even an explicit nil
### GetInput

`func (o *Content17Inner) GetInput() map[string]interface{}`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *Content17Inner) GetInputOk() (*map[string]interface{}, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *Content17Inner) SetInput(v map[string]interface{})`

SetInput sets Input field to given value.


### GetName

`func (o *Content17Inner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Content17Inner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Content17Inner) SetName(v string)`

SetName sets Name field to given value.


### GetToolsetName

`func (o *Content17Inner) GetToolsetName() string`

GetToolsetName returns the ToolsetName field if non-nil, zero value otherwise.

### GetToolsetNameOk

`func (o *Content17Inner) GetToolsetNameOk() (*string, bool)`

GetToolsetNameOk returns a tuple with the ToolsetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolsetName

`func (o *Content17Inner) SetToolsetName(v string)`

SetToolsetName sets ToolsetName field to given value.

### HasToolsetName

`func (o *Content17Inner) HasToolsetName() bool`

HasToolsetName returns a boolean if a field has been set.

### GetContent

`func (o *Content17Inner) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *Content17Inner) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *Content17Inner) SetContent(v string)`

SetContent sets Content field to given value.


### GetToolUseId

`func (o *Content17Inner) GetToolUseId() string`

GetToolUseId returns the ToolUseId field if non-nil, zero value otherwise.

### GetToolUseIdOk

`func (o *Content17Inner) GetToolUseIdOk() (*string, bool)`

GetToolUseIdOk returns a tuple with the ToolUseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolUseId

`func (o *Content17Inner) SetToolUseId(v string)`

SetToolUseId sets ToolUseId field to given value.


### GetFileId

`func (o *Content17Inner) GetFileId() string`

GetFileId returns the FileId field if non-nil, zero value otherwise.

### GetFileIdOk

`func (o *Content17Inner) GetFileIdOk() (*string, bool)`

GetFileIdOk returns a tuple with the FileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileId

`func (o *Content17Inner) SetFileId(v string)`

SetFileId sets FileId field to given value.


### GetServerName

`func (o *Content17Inner) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *Content17Inner) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *Content17Inner) SetServerName(v string)`

SetServerName sets ServerName field to given value.


### GetIsError

`func (o *Content17Inner) GetIsError() bool`

GetIsError returns the IsError field if non-nil, zero value otherwise.

### GetIsErrorOk

`func (o *Content17Inner) GetIsErrorOk() (*bool, bool)`

GetIsErrorOk returns a tuple with the IsError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsError

`func (o *Content17Inner) SetIsError(v bool)`

SetIsError sets IsError field to given value.


### GetEncryptedContent

`func (o *Content17Inner) GetEncryptedContent() string`

GetEncryptedContent returns the EncryptedContent field if non-nil, zero value otherwise.

### GetEncryptedContentOk

`func (o *Content17Inner) GetEncryptedContentOk() (*string, bool)`

GetEncryptedContentOk returns a tuple with the EncryptedContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedContent

`func (o *Content17Inner) SetEncryptedContent(v string)`

SetEncryptedContent sets EncryptedContent field to given value.

### HasEncryptedContent

`func (o *Content17Inner) HasEncryptedContent() bool`

HasEncryptedContent returns a boolean if a field has been set.

### GetFrom

`func (o *Content17Inner) GetFrom() MRBetaFallbackInfo`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *Content17Inner) GetFromOk() (*MRBetaFallbackInfo, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *Content17Inner) SetFrom(v MRBetaFallbackInfo)`

SetFrom sets From field to given value.


### GetTo

`func (o *Content17Inner) GetTo() MRBetaFallbackInfo`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *Content17Inner) GetToOk() (*MRBetaFallbackInfo, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *Content17Inner) SetTo(v MRBetaFallbackInfo)`

SetTo sets To field to given value.


### GetTrigger

`func (o *Content17Inner) GetTrigger() MRBetaFallbackRefusalTrigger`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *Content17Inner) GetTriggerOk() (*MRBetaFallbackRefusalTrigger, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *Content17Inner) SetTrigger(v MRBetaFallbackRefusalTrigger)`

SetTrigger sets Trigger field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


