# MessageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Container** | Pointer to [**NullableMRContainer**](MRContainer.md) |  | [optional] 
**Content** | [**[]Content9Inner**](Content9Inner.md) |  | 
**Model** | [**Model**](Model.md) |  | 
**Role** | **string** |  | 
**StopDetails** | Pointer to [**NullableMRRefusalStopDetails**](MRRefusalStopDetails.md) |  | [optional] 
**StopReason** | Pointer to **NullableString** |  | [optional] 
**StopSequence** | Pointer to **NullableString** | Filter models by provider name | [optional] 
**Type** | **string** |  | 
**Usage** | [**MRUsage**](MRUsage.md) |  | 

## Methods

### NewMessageResponse

`func NewMessageResponse(id string, content []Content9Inner, model Model, role string, type_ string, usage MRUsage, ) *MessageResponse`

NewMessageResponse instantiates a new MessageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageResponseWithDefaults

`func NewMessageResponseWithDefaults() *MessageResponse`

NewMessageResponseWithDefaults instantiates a new MessageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MessageResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MessageResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MessageResponse) SetId(v string)`

SetId sets Id field to given value.


### GetContainer

`func (o *MessageResponse) GetContainer() MRContainer`

GetContainer returns the Container field if non-nil, zero value otherwise.

### GetContainerOk

`func (o *MessageResponse) GetContainerOk() (*MRContainer, bool)`

GetContainerOk returns a tuple with the Container field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainer

`func (o *MessageResponse) SetContainer(v MRContainer)`

SetContainer sets Container field to given value.

### HasContainer

`func (o *MessageResponse) HasContainer() bool`

HasContainer returns a boolean if a field has been set.

### SetContainerNil

`func (o *MessageResponse) SetContainerNil(b bool)`

 SetContainerNil sets the value for Container to be an explicit nil

### UnsetContainer
`func (o *MessageResponse) UnsetContainer()`

UnsetContainer ensures that no value is present for Container, not even an explicit nil
### GetContent

`func (o *MessageResponse) GetContent() []Content9Inner`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *MessageResponse) GetContentOk() (*[]Content9Inner, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *MessageResponse) SetContent(v []Content9Inner)`

SetContent sets Content field to given value.


### GetModel

`func (o *MessageResponse) GetModel() Model`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *MessageResponse) GetModelOk() (*Model, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *MessageResponse) SetModel(v Model)`

SetModel sets Model field to given value.


### GetRole

`func (o *MessageResponse) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *MessageResponse) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *MessageResponse) SetRole(v string)`

SetRole sets Role field to given value.


### GetStopDetails

`func (o *MessageResponse) GetStopDetails() MRRefusalStopDetails`

GetStopDetails returns the StopDetails field if non-nil, zero value otherwise.

### GetStopDetailsOk

`func (o *MessageResponse) GetStopDetailsOk() (*MRRefusalStopDetails, bool)`

GetStopDetailsOk returns a tuple with the StopDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopDetails

`func (o *MessageResponse) SetStopDetails(v MRRefusalStopDetails)`

SetStopDetails sets StopDetails field to given value.

### HasStopDetails

`func (o *MessageResponse) HasStopDetails() bool`

HasStopDetails returns a boolean if a field has been set.

### SetStopDetailsNil

`func (o *MessageResponse) SetStopDetailsNil(b bool)`

 SetStopDetailsNil sets the value for StopDetails to be an explicit nil

### UnsetStopDetails
`func (o *MessageResponse) UnsetStopDetails()`

UnsetStopDetails ensures that no value is present for StopDetails, not even an explicit nil
### GetStopReason

`func (o *MessageResponse) GetStopReason() string`

GetStopReason returns the StopReason field if non-nil, zero value otherwise.

### GetStopReasonOk

`func (o *MessageResponse) GetStopReasonOk() (*string, bool)`

GetStopReasonOk returns a tuple with the StopReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopReason

`func (o *MessageResponse) SetStopReason(v string)`

SetStopReason sets StopReason field to given value.

### HasStopReason

`func (o *MessageResponse) HasStopReason() bool`

HasStopReason returns a boolean if a field has been set.

### SetStopReasonNil

`func (o *MessageResponse) SetStopReasonNil(b bool)`

 SetStopReasonNil sets the value for StopReason to be an explicit nil

### UnsetStopReason
`func (o *MessageResponse) UnsetStopReason()`

UnsetStopReason ensures that no value is present for StopReason, not even an explicit nil
### GetStopSequence

`func (o *MessageResponse) GetStopSequence() string`

GetStopSequence returns the StopSequence field if non-nil, zero value otherwise.

### GetStopSequenceOk

`func (o *MessageResponse) GetStopSequenceOk() (*string, bool)`

GetStopSequenceOk returns a tuple with the StopSequence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopSequence

`func (o *MessageResponse) SetStopSequence(v string)`

SetStopSequence sets StopSequence field to given value.

### HasStopSequence

`func (o *MessageResponse) HasStopSequence() bool`

HasStopSequence returns a boolean if a field has been set.

### SetStopSequenceNil

`func (o *MessageResponse) SetStopSequenceNil(b bool)`

 SetStopSequenceNil sets the value for StopSequence to be an explicit nil

### UnsetStopSequence
`func (o *MessageResponse) UnsetStopSequence()`

UnsetStopSequence ensures that no value is present for StopSequence, not even an explicit nil
### GetType

`func (o *MessageResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MessageResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MessageResponse) SetType(v string)`

SetType sets Type field to given value.


### GetUsage

`func (o *MessageResponse) GetUsage() MRUsage`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *MessageResponse) GetUsageOk() (*MRUsage, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *MessageResponse) SetUsage(v MRUsage)`

SetUsage sets Usage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


