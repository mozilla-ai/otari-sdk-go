# ModelObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | **int32** |  | 
**Id** | **string** |  | 
**Object** | Pointer to **string** |  | [optional] [default to "model"]
**OwnedBy** | **string** |  | 
**Pricing** | Pointer to [**NullableModelPricingInfo**](ModelPricingInfo.md) |  | [optional] 

## Methods

### NewModelObject

`func NewModelObject(created int32, id string, ownedBy string, ) *ModelObject`

NewModelObject instantiates a new ModelObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelObjectWithDefaults

`func NewModelObjectWithDefaults() *ModelObject`

NewModelObjectWithDefaults instantiates a new ModelObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *ModelObject) GetCreated() int32`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ModelObject) GetCreatedOk() (*int32, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ModelObject) SetCreated(v int32)`

SetCreated sets Created field to given value.


### GetId

`func (o *ModelObject) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelObject) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelObject) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *ModelObject) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ModelObject) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ModelObject) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *ModelObject) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetOwnedBy

`func (o *ModelObject) GetOwnedBy() string`

GetOwnedBy returns the OwnedBy field if non-nil, zero value otherwise.

### GetOwnedByOk

`func (o *ModelObject) GetOwnedByOk() (*string, bool)`

GetOwnedByOk returns a tuple with the OwnedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnedBy

`func (o *ModelObject) SetOwnedBy(v string)`

SetOwnedBy sets OwnedBy field to given value.


### GetPricing

`func (o *ModelObject) GetPricing() ModelPricingInfo`

GetPricing returns the Pricing field if non-nil, zero value otherwise.

### GetPricingOk

`func (o *ModelObject) GetPricingOk() (*ModelPricingInfo, bool)`

GetPricingOk returns a tuple with the Pricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricing

`func (o *ModelObject) SetPricing(v ModelPricingInfo)`

SetPricing sets Pricing field to given value.

### HasPricing

`func (o *ModelObject) HasPricing() bool`

HasPricing returns a boolean if a field has been set.

### SetPricingNil

`func (o *ModelObject) SetPricingNil(b bool)`

 SetPricingNil sets the value for Pricing to be an explicit nil

### UnsetPricing
`func (o *ModelObject) UnsetPricing()`

UnsetPricing ensures that no value is present for Pricing, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


