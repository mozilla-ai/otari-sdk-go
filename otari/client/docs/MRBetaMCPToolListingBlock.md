# MRBetaMCPToolListingBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**McpServerName** | **string** |  | 
**Tools** | [**[]MRBetaMCPTool**](MRBetaMCPTool.md) |  | 
**Type** | **string** |  | 

## Methods

### NewMRBetaMCPToolListingBlock

`func NewMRBetaMCPToolListingBlock(mcpServerName string, tools []MRBetaMCPTool, type_ string, ) *MRBetaMCPToolListingBlock`

NewMRBetaMCPToolListingBlock instantiates a new MRBetaMCPToolListingBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMRBetaMCPToolListingBlockWithDefaults

`func NewMRBetaMCPToolListingBlockWithDefaults() *MRBetaMCPToolListingBlock`

NewMRBetaMCPToolListingBlockWithDefaults instantiates a new MRBetaMCPToolListingBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMcpServerName

`func (o *MRBetaMCPToolListingBlock) GetMcpServerName() string`

GetMcpServerName returns the McpServerName field if non-nil, zero value otherwise.

### GetMcpServerNameOk

`func (o *MRBetaMCPToolListingBlock) GetMcpServerNameOk() (*string, bool)`

GetMcpServerNameOk returns a tuple with the McpServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcpServerName

`func (o *MRBetaMCPToolListingBlock) SetMcpServerName(v string)`

SetMcpServerName sets McpServerName field to given value.


### GetTools

`func (o *MRBetaMCPToolListingBlock) GetTools() []MRBetaMCPTool`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *MRBetaMCPToolListingBlock) GetToolsOk() (*[]MRBetaMCPTool, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *MRBetaMCPToolListingBlock) SetTools(v []MRBetaMCPTool)`

SetTools sets Tools field to given value.


### GetType

`func (o *MRBetaMCPToolListingBlock) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MRBetaMCPToolListingBlock) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MRBetaMCPToolListingBlock) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


