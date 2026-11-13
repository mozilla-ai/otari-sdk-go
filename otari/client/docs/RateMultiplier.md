# RateMultiplier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Factor** | [**Factor**](Factor.md) |  | 
**When** | **map[string]interface{}** | The condition, as models.dev writes it. | 

## Methods

### NewRateMultiplier

`func NewRateMultiplier(factor Factor, when map[string]interface{}, ) *RateMultiplier`

NewRateMultiplier instantiates a new RateMultiplier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRateMultiplierWithDefaults

`func NewRateMultiplierWithDefaults() *RateMultiplier`

NewRateMultiplierWithDefaults instantiates a new RateMultiplier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFactor

`func (o *RateMultiplier) GetFactor() Factor`

GetFactor returns the Factor field if non-nil, zero value otherwise.

### GetFactorOk

`func (o *RateMultiplier) GetFactorOk() (*Factor, bool)`

GetFactorOk returns a tuple with the Factor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactor

`func (o *RateMultiplier) SetFactor(v Factor)`

SetFactor sets Factor field to given value.


### GetWhen

`func (o *RateMultiplier) GetWhen() map[string]interface{}`

GetWhen returns the When field if non-nil, zero value otherwise.

### GetWhenOk

`func (o *RateMultiplier) GetWhenOk() (*map[string]interface{}, bool)`

GetWhenOk returns a tuple with the When field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhen

`func (o *RateMultiplier) SetWhen(v map[string]interface{})`

SetWhen sets When field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


