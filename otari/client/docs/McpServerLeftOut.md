# McpServerLeftOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Detail** | **string** | The reason in a sentence, for a person reading the response | 
**Id** | **string** |  | 
**Name** | **string** |  | 
**Reason** | **string** | &#x60;key_has_no_user&#x60;: the server signs each person in with their own account, and this key belongs to no person, so there is nobody to sign in as | 

## Methods

### NewMcpServerLeftOut

`func NewMcpServerLeftOut(detail string, id string, name string, reason string, ) *McpServerLeftOut`

NewMcpServerLeftOut instantiates a new McpServerLeftOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMcpServerLeftOutWithDefaults

`func NewMcpServerLeftOutWithDefaults() *McpServerLeftOut`

NewMcpServerLeftOutWithDefaults instantiates a new McpServerLeftOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetail

`func (o *McpServerLeftOut) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *McpServerLeftOut) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *McpServerLeftOut) SetDetail(v string)`

SetDetail sets Detail field to given value.


### GetId

`func (o *McpServerLeftOut) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *McpServerLeftOut) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *McpServerLeftOut) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *McpServerLeftOut) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *McpServerLeftOut) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *McpServerLeftOut) SetName(v string)`

SetName sets Name field to given value.


### GetReason

`func (o *McpServerLeftOut) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *McpServerLeftOut) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *McpServerLeftOut) SetReason(v string)`

SetReason sets Reason field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


