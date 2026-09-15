# WorkspaceCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** |  | [optional] 
**Name** | **string** |  | 

## Methods

### NewWorkspaceCreate

`func NewWorkspaceCreate(name string, ) *WorkspaceCreate`

NewWorkspaceCreate instantiates a new WorkspaceCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkspaceCreateWithDefaults

`func NewWorkspaceCreateWithDefaults() *WorkspaceCreate`

NewWorkspaceCreateWithDefaults instantiates a new WorkspaceCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *WorkspaceCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkspaceCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkspaceCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkspaceCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *WorkspaceCreate) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *WorkspaceCreate) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetName

`func (o *WorkspaceCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkspaceCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkspaceCreate) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


