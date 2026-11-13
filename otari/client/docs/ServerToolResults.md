# ServerToolResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Tools** | [**[]MRBetaWebFetchURLSourceToolReference**](MRBetaWebFetchURLSourceToolReference.md) |  | 

## Methods

### NewServerToolResults

`func NewServerToolResults(type_ string, tools []MRBetaWebFetchURLSourceToolReference, ) *ServerToolResults`

NewServerToolResults instantiates a new ServerToolResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServerToolResultsWithDefaults

`func NewServerToolResultsWithDefaults() *ServerToolResults`

NewServerToolResultsWithDefaults instantiates a new ServerToolResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ServerToolResults) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ServerToolResults) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ServerToolResults) SetType(v string)`

SetType sets Type field to given value.


### GetTools

`func (o *ServerToolResults) GetTools() []MRBetaWebFetchURLSourceToolReference`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *ServerToolResults) GetToolsOk() (*[]MRBetaWebFetchURLSourceToolReference, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *ServerToolResults) SetTools(v []MRBetaWebFetchURLSourceToolReference)`

SetTools sets Tools field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


