# ManagedTool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AcceptedTypes** | **[]string** | Every &#x60;tools[].type&#x60; this deployment currently routes to the tool. Always includes the canonical &#x60;otari_*&#x60; type; for web search it also includes the provider-named keywords when interception is enabled. | 
**Available** | **bool** | Whether this deployment has a backend configured for the tool. A request declaring an unavailable tool is rejected with 400. | 
**Description** | **string** | What the tool does, as the model is told. | 
**Example** | **map[string]interface{}** | A ready-to-use &#x60;tools[]&#x60; entry. | 
**Id** | **string** | The canonical tool type to put in &#x60;tools[]&#x60;. | 
**InputSchema** | **map[string]interface{}** | JSON Schema for the arguments the model supplies, as the model sees it. | 
**Object** | Pointer to **string** |  | [optional] [default to "tool"]

## Methods

### NewManagedTool

`func NewManagedTool(acceptedTypes []string, available bool, description string, example map[string]interface{}, id string, inputSchema map[string]interface{}, ) *ManagedTool`

NewManagedTool instantiates a new ManagedTool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewManagedToolWithDefaults

`func NewManagedToolWithDefaults() *ManagedTool`

NewManagedToolWithDefaults instantiates a new ManagedTool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAcceptedTypes

`func (o *ManagedTool) GetAcceptedTypes() []string`

GetAcceptedTypes returns the AcceptedTypes field if non-nil, zero value otherwise.

### GetAcceptedTypesOk

`func (o *ManagedTool) GetAcceptedTypesOk() (*[]string, bool)`

GetAcceptedTypesOk returns a tuple with the AcceptedTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedTypes

`func (o *ManagedTool) SetAcceptedTypes(v []string)`

SetAcceptedTypes sets AcceptedTypes field to given value.


### GetAvailable

`func (o *ManagedTool) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *ManagedTool) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *ManagedTool) SetAvailable(v bool)`

SetAvailable sets Available field to given value.


### GetDescription

`func (o *ManagedTool) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ManagedTool) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ManagedTool) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetExample

`func (o *ManagedTool) GetExample() map[string]interface{}`

GetExample returns the Example field if non-nil, zero value otherwise.

### GetExampleOk

`func (o *ManagedTool) GetExampleOk() (*map[string]interface{}, bool)`

GetExampleOk returns a tuple with the Example field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExample

`func (o *ManagedTool) SetExample(v map[string]interface{})`

SetExample sets Example field to given value.


### GetId

`func (o *ManagedTool) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ManagedTool) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ManagedTool) SetId(v string)`

SetId sets Id field to given value.


### GetInputSchema

`func (o *ManagedTool) GetInputSchema() map[string]interface{}`

GetInputSchema returns the InputSchema field if non-nil, zero value otherwise.

### GetInputSchemaOk

`func (o *ManagedTool) GetInputSchemaOk() (*map[string]interface{}, bool)`

GetInputSchemaOk returns a tuple with the InputSchema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputSchema

`func (o *ManagedTool) SetInputSchema(v map[string]interface{})`

SetInputSchema sets InputSchema field to given value.


### GetObject

`func (o *ManagedTool) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ManagedTool) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ManagedTool) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *ManagedTool) HasObject() bool`

HasObject returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


