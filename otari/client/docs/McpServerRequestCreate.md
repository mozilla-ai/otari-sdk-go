# McpServerRequestCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CatalogName** | **string** | The catalog server&#39;s &#x60;name&#x60;, such as &#x60;com.example/mcp&#x60; | 
**Note** | Pointer to **NullableString** | What the member needs the server for, shown to the admins | [optional] 

## Methods

### NewMcpServerRequestCreate

`func NewMcpServerRequestCreate(catalogName string, ) *McpServerRequestCreate`

NewMcpServerRequestCreate instantiates a new McpServerRequestCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMcpServerRequestCreateWithDefaults

`func NewMcpServerRequestCreateWithDefaults() *McpServerRequestCreate`

NewMcpServerRequestCreateWithDefaults instantiates a new McpServerRequestCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCatalogName

`func (o *McpServerRequestCreate) GetCatalogName() string`

GetCatalogName returns the CatalogName field if non-nil, zero value otherwise.

### GetCatalogNameOk

`func (o *McpServerRequestCreate) GetCatalogNameOk() (*string, bool)`

GetCatalogNameOk returns a tuple with the CatalogName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogName

`func (o *McpServerRequestCreate) SetCatalogName(v string)`

SetCatalogName sets CatalogName field to given value.


### GetNote

`func (o *McpServerRequestCreate) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *McpServerRequestCreate) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *McpServerRequestCreate) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *McpServerRequestCreate) HasNote() bool`

HasNote returns a boolean if a field has been set.

### SetNoteNil

`func (o *McpServerRequestCreate) SetNoteNil(b bool)`

 SetNoteNil sets the value for Note to be an explicit nil

### UnsetNote
`func (o *McpServerRequestCreate) UnsetNote()`

UnsetNote ensures that no value is present for Note, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


