# ModerationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Model** | **string** |  | 
**Results** | [**[]ModerationResult**](ModerationResult.md) |  | 

## Methods

### NewModerationResponse

`func NewModerationResponse(id string, model string, results []ModerationResult, ) *ModerationResponse`

NewModerationResponse instantiates a new ModerationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModerationResponseWithDefaults

`func NewModerationResponseWithDefaults() *ModerationResponse`

NewModerationResponseWithDefaults instantiates a new ModerationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModerationResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModerationResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModerationResponse) SetId(v string)`

SetId sets Id field to given value.


### GetModel

`func (o *ModerationResponse) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ModerationResponse) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ModerationResponse) SetModel(v string)`

SetModel sets Model field to given value.


### GetResults

`func (o *ModerationResponse) GetResults() []ModerationResult`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *ModerationResponse) GetResultsOk() (*[]ModerationResult, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *ModerationResponse) SetResults(v []ModerationResult)`

SetResults sets Results field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


