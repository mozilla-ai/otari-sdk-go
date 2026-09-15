# CallToolResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to **map[string]interface{}** |  | [optional] 
**Content** | [**[]ContentInner**](ContentInner.md) |  | 
**IsError** | Pointer to **bool** |  | [optional] [default to false]
**StructuredContent** | Pointer to **map[string]interface{}** | Provider-native request fields used as defaults (e.g. exa&#39;s &#39;type&#39;, searxng&#39;s &#39;engines&#39;). | [optional] 

## Methods

### NewCallToolResult

`func NewCallToolResult(content []ContentInner, ) *CallToolResult`

NewCallToolResult instantiates a new CallToolResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallToolResultWithDefaults

`func NewCallToolResultWithDefaults() *CallToolResult`

NewCallToolResultWithDefaults instantiates a new CallToolResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *CallToolResult) GetMeta() map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *CallToolResult) GetMetaOk() (*map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *CallToolResult) SetMeta(v map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *CallToolResult) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### SetMetaNil

`func (o *CallToolResult) SetMetaNil(b bool)`

 SetMetaNil sets the value for Meta to be an explicit nil

### UnsetMeta
`func (o *CallToolResult) UnsetMeta()`

UnsetMeta ensures that no value is present for Meta, not even an explicit nil
### GetContent

`func (o *CallToolResult) GetContent() []ContentInner`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *CallToolResult) GetContentOk() (*[]ContentInner, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *CallToolResult) SetContent(v []ContentInner)`

SetContent sets Content field to given value.


### GetIsError

`func (o *CallToolResult) GetIsError() bool`

GetIsError returns the IsError field if non-nil, zero value otherwise.

### GetIsErrorOk

`func (o *CallToolResult) GetIsErrorOk() (*bool, bool)`

GetIsErrorOk returns a tuple with the IsError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsError

`func (o *CallToolResult) SetIsError(v bool)`

SetIsError sets IsError field to given value.

### HasIsError

`func (o *CallToolResult) HasIsError() bool`

HasIsError returns a boolean if a field has been set.

### GetStructuredContent

`func (o *CallToolResult) GetStructuredContent() map[string]interface{}`

GetStructuredContent returns the StructuredContent field if non-nil, zero value otherwise.

### GetStructuredContentOk

`func (o *CallToolResult) GetStructuredContentOk() (*map[string]interface{}, bool)`

GetStructuredContentOk returns a tuple with the StructuredContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStructuredContent

`func (o *CallToolResult) SetStructuredContent(v map[string]interface{})`

SetStructuredContent sets StructuredContent field to given value.

### HasStructuredContent

`func (o *CallToolResult) HasStructuredContent() bool`

HasStructuredContent returns a boolean if a field has been set.

### SetStructuredContentNil

`func (o *CallToolResult) SetStructuredContentNil(b bool)`

 SetStructuredContentNil sets the value for StructuredContent to be an explicit nil

### UnsetStructuredContent
`func (o *CallToolResult) UnsetStructuredContent()`

UnsetStructuredContent ensures that no value is present for StructuredContent, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


