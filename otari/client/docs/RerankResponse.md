# RerankResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | Filter models by provider name | [optional] 
**Results** | [**[]RRRerankResult**](RRRerankResult.md) | Results sorted by relevance_score descending | 
**Meta** | Pointer to [**NullableRRRerankMeta**](RRRerankMeta.md) |  | [optional] 
**Usage** | Pointer to [**NullableRRRerankUsage**](RRRerankUsage.md) |  | [optional] 

## Methods

### NewRerankResponse

`func NewRerankResponse(results []RRRerankResult, ) *RerankResponse`

NewRerankResponse instantiates a new RerankResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRerankResponseWithDefaults

`func NewRerankResponseWithDefaults() *RerankResponse`

NewRerankResponseWithDefaults instantiates a new RerankResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RerankResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RerankResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RerankResponse) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RerankResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *RerankResponse) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *RerankResponse) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetResults

`func (o *RerankResponse) GetResults() []RRRerankResult`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *RerankResponse) GetResultsOk() (*[]RRRerankResult, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *RerankResponse) SetResults(v []RRRerankResult)`

SetResults sets Results field to given value.


### GetMeta

`func (o *RerankResponse) GetMeta() RRRerankMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *RerankResponse) GetMetaOk() (*RRRerankMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *RerankResponse) SetMeta(v RRRerankMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *RerankResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### SetMetaNil

`func (o *RerankResponse) SetMetaNil(b bool)`

 SetMetaNil sets the value for Meta to be an explicit nil

### UnsetMeta
`func (o *RerankResponse) UnsetMeta()`

UnsetMeta ensures that no value is present for Meta, not even an explicit nil
### GetUsage

`func (o *RerankResponse) GetUsage() RRRerankUsage`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *RerankResponse) GetUsageOk() (*RRRerankUsage, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *RerankResponse) SetUsage(v RRRerankUsage)`

SetUsage sets Usage field to given value.

### HasUsage

`func (o *RerankResponse) HasUsage() bool`

HasUsage returns a boolean if a field has been set.

### SetUsageNil

`func (o *RerankResponse) SetUsageNil(b bool)`

 SetUsageNil sets the value for Usage to be an explicit nil

### UnsetUsage
`func (o *RerankResponse) UnsetUsage()`

UnsetUsage ensures that no value is present for Usage, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


