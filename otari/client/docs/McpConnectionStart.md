# McpConnectionStart

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReturnTo** | Pointer to **NullableString** | A dashboard path to land on once signed in, such as /tools/mcp-servers. Anything that is not a path inside the dashboard is refused | [optional] 

## Methods

### NewMcpConnectionStart

`func NewMcpConnectionStart() *McpConnectionStart`

NewMcpConnectionStart instantiates a new McpConnectionStart object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMcpConnectionStartWithDefaults

`func NewMcpConnectionStartWithDefaults() *McpConnectionStart`

NewMcpConnectionStartWithDefaults instantiates a new McpConnectionStart object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReturnTo

`func (o *McpConnectionStart) GetReturnTo() string`

GetReturnTo returns the ReturnTo field if non-nil, zero value otherwise.

### GetReturnToOk

`func (o *McpConnectionStart) GetReturnToOk() (*string, bool)`

GetReturnToOk returns a tuple with the ReturnTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnTo

`func (o *McpConnectionStart) SetReturnTo(v string)`

SetReturnTo sets ReturnTo field to given value.

### HasReturnTo

`func (o *McpConnectionStart) HasReturnTo() bool`

HasReturnTo returns a boolean if a field has been set.

### SetReturnToNil

`func (o *McpConnectionStart) SetReturnToNil(b bool)`

 SetReturnToNil sets the value for ReturnTo to be an explicit nil

### UnsetReturnTo
`func (o *McpConnectionStart) UnsetReturnTo()`

UnsetReturnTo ensures that no value is present for ReturnTo, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


