# McpConnectionStarted

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthorizationUrl** | **string** | The MCP server&#39;s authorization page. It carries a single-use state that expires in 30 minutes | 

## Methods

### NewMcpConnectionStarted

`func NewMcpConnectionStarted(authorizationUrl string, ) *McpConnectionStarted`

NewMcpConnectionStarted instantiates a new McpConnectionStarted object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMcpConnectionStartedWithDefaults

`func NewMcpConnectionStartedWithDefaults() *McpConnectionStarted`

NewMcpConnectionStartedWithDefaults instantiates a new McpConnectionStarted object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthorizationUrl

`func (o *McpConnectionStarted) GetAuthorizationUrl() string`

GetAuthorizationUrl returns the AuthorizationUrl field if non-nil, zero value otherwise.

### GetAuthorizationUrlOk

`func (o *McpConnectionStarted) GetAuthorizationUrlOk() (*string, bool)`

GetAuthorizationUrlOk returns a tuple with the AuthorizationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizationUrl

`func (o *McpConnectionStarted) SetAuthorizationUrl(v string)`

SetAuthorizationUrl sets AuthorizationUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


