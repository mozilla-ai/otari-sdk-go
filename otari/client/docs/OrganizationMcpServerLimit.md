# OrganizationMcpServerLimit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tools** | **[]string** | Tool patterns (&#39;*&#39;, a name, &#39;prefix*&#39;, &#39;@read&#39;) the workspace&#39;s callers may reach at most; [] turns the server off for the workspace, and its admins cannot turn it back on | 

## Methods

### NewOrganizationMcpServerLimit

`func NewOrganizationMcpServerLimit(tools []string, ) *OrganizationMcpServerLimit`

NewOrganizationMcpServerLimit instantiates a new OrganizationMcpServerLimit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationMcpServerLimitWithDefaults

`func NewOrganizationMcpServerLimitWithDefaults() *OrganizationMcpServerLimit`

NewOrganizationMcpServerLimitWithDefaults instantiates a new OrganizationMcpServerLimit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTools

`func (o *OrganizationMcpServerLimit) GetTools() []string`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *OrganizationMcpServerLimit) GetToolsOk() (*[]string, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *OrganizationMcpServerLimit) SetTools(v []string)`

SetTools sets Tools field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


