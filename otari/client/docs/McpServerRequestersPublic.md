# McpServerRequestersPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Count** | **int32** | How many pending requests ask for this server in all | 
**Data** | [**[]McpServerRequesterPublic**](McpServerRequesterPublic.md) | Who asked, earliest first | 

## Methods

### NewMcpServerRequestersPublic

`func NewMcpServerRequestersPublic(count int32, data []McpServerRequesterPublic, ) *McpServerRequestersPublic`

NewMcpServerRequestersPublic instantiates a new McpServerRequestersPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMcpServerRequestersPublicWithDefaults

`func NewMcpServerRequestersPublicWithDefaults() *McpServerRequestersPublic`

NewMcpServerRequestersPublicWithDefaults instantiates a new McpServerRequestersPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCount

`func (o *McpServerRequestersPublic) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *McpServerRequestersPublic) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *McpServerRequestersPublic) SetCount(v int32)`

SetCount sets Count field to given value.


### GetData

`func (o *McpServerRequestersPublic) GetData() []McpServerRequesterPublic`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *McpServerRequestersPublic) GetDataOk() (*[]McpServerRequesterPublic, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *McpServerRequestersPublic) SetData(v []McpServerRequesterPublic)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


