# ClientToolResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Tools** | [**[]MRBetaWebFetchURLSourceToolReference**](MRBetaWebFetchURLSourceToolReference.md) |  | 

## Methods

### NewClientToolResults

`func NewClientToolResults(type_ string, tools []MRBetaWebFetchURLSourceToolReference, ) *ClientToolResults`

NewClientToolResults instantiates a new ClientToolResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientToolResultsWithDefaults

`func NewClientToolResultsWithDefaults() *ClientToolResults`

NewClientToolResultsWithDefaults instantiates a new ClientToolResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ClientToolResults) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ClientToolResults) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ClientToolResults) SetType(v string)`

SetType sets Type field to given value.


### GetTools

`func (o *ClientToolResults) GetTools() []MRBetaWebFetchURLSourceToolReference`

GetTools returns the Tools field if non-nil, zero value otherwise.

### GetToolsOk

`func (o *ClientToolResults) GetToolsOk() (*[]MRBetaWebFetchURLSourceToolReference, bool)`

GetToolsOk returns a tuple with the Tools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTools

`func (o *ClientToolResults) SetTools(v []MRBetaWebFetchURLSourceToolReference)`

SetTools sets Tools field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


