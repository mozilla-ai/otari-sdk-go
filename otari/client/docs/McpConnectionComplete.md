# McpConnectionComplete

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | The authorization code from the redirect | 
**Iss** | Pointer to **NullableString** | The redirect&#39;s iss, the authorization server&#39;s own name, if it sent one | [optional] 
**State** | **string** | The state from the redirect | 

## Methods

### NewMcpConnectionComplete

`func NewMcpConnectionComplete(code string, state string, ) *McpConnectionComplete`

NewMcpConnectionComplete instantiates a new McpConnectionComplete object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMcpConnectionCompleteWithDefaults

`func NewMcpConnectionCompleteWithDefaults() *McpConnectionComplete`

NewMcpConnectionCompleteWithDefaults instantiates a new McpConnectionComplete object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *McpConnectionComplete) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *McpConnectionComplete) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *McpConnectionComplete) SetCode(v string)`

SetCode sets Code field to given value.


### GetIss

`func (o *McpConnectionComplete) GetIss() string`

GetIss returns the Iss field if non-nil, zero value otherwise.

### GetIssOk

`func (o *McpConnectionComplete) GetIssOk() (*string, bool)`

GetIssOk returns a tuple with the Iss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIss

`func (o *McpConnectionComplete) SetIss(v string)`

SetIss sets Iss field to given value.

### HasIss

`func (o *McpConnectionComplete) HasIss() bool`

HasIss returns a boolean if a field has been set.

### SetIssNil

`func (o *McpConnectionComplete) SetIssNil(b bool)`

 SetIssNil sets the value for Iss to be an explicit nil

### UnsetIss
`func (o *McpConnectionComplete) UnsetIss()`

UnsetIss ensures that no value is present for Iss, not even an explicit nil
### GetState

`func (o *McpConnectionComplete) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *McpConnectionComplete) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *McpConnectionComplete) SetState(v string)`

SetState sets State field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


