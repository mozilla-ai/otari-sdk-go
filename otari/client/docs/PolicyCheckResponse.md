# PolicyCheckResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Blocked** | **bool** |  | 
**PolicyId** | **string** |  | 
**Provenance** | Pointer to **string** |  | [optional] [default to "client_reported"]
**Results** | [**[]GateResultResponse**](GateResultResponse.md) |  | 
**SchemaVersion** | **string** |  | 

## Methods

### NewPolicyCheckResponse

`func NewPolicyCheckResponse(blocked bool, policyId string, results []GateResultResponse, schemaVersion string, ) *PolicyCheckResponse`

NewPolicyCheckResponse instantiates a new PolicyCheckResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyCheckResponseWithDefaults

`func NewPolicyCheckResponseWithDefaults() *PolicyCheckResponse`

NewPolicyCheckResponseWithDefaults instantiates a new PolicyCheckResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlocked

`func (o *PolicyCheckResponse) GetBlocked() bool`

GetBlocked returns the Blocked field if non-nil, zero value otherwise.

### GetBlockedOk

`func (o *PolicyCheckResponse) GetBlockedOk() (*bool, bool)`

GetBlockedOk returns a tuple with the Blocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocked

`func (o *PolicyCheckResponse) SetBlocked(v bool)`

SetBlocked sets Blocked field to given value.


### GetPolicyId

`func (o *PolicyCheckResponse) GetPolicyId() string`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *PolicyCheckResponse) GetPolicyIdOk() (*string, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *PolicyCheckResponse) SetPolicyId(v string)`

SetPolicyId sets PolicyId field to given value.


### GetProvenance

`func (o *PolicyCheckResponse) GetProvenance() string`

GetProvenance returns the Provenance field if non-nil, zero value otherwise.

### GetProvenanceOk

`func (o *PolicyCheckResponse) GetProvenanceOk() (*string, bool)`

GetProvenanceOk returns a tuple with the Provenance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvenance

`func (o *PolicyCheckResponse) SetProvenance(v string)`

SetProvenance sets Provenance field to given value.

### HasProvenance

`func (o *PolicyCheckResponse) HasProvenance() bool`

HasProvenance returns a boolean if a field has been set.

### GetResults

`func (o *PolicyCheckResponse) GetResults() []GateResultResponse`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PolicyCheckResponse) GetResultsOk() (*[]GateResultResponse, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PolicyCheckResponse) SetResults(v []GateResultResponse)`

SetResults sets Results field to given value.


### GetSchemaVersion

`func (o *PolicyCheckResponse) GetSchemaVersion() string`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *PolicyCheckResponse) GetSchemaVersionOk() (*string, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *PolicyCheckResponse) SetSchemaVersion(v string)`

SetSchemaVersion sets SchemaVersion field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


