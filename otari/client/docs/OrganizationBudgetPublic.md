# OrganizationBudgetPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BudgetDurationSec** | **NullableInt32** |  | 
**BudgetId** | **string** |  | 
**CeilingCount** | **int32** |  | 
**CreatedAt** | **string** |  | 
**MaxBudget** | **NullableFloat32** |  | 
**Name** | **NullableString** |  | 
**OrganizationId** | **string** |  | 
**RequestLimit** | **NullableInt32** |  | 
**ResetAlignment** | **NullableString** |  | 
**TokenLimit** | **NullableInt32** |  | 
**UpdatedAt** | **string** |  | 

## Methods

### NewOrganizationBudgetPublic

`func NewOrganizationBudgetPublic(budgetDurationSec NullableInt32, budgetId string, ceilingCount int32, createdAt string, maxBudget NullableFloat32, name NullableString, organizationId string, requestLimit NullableInt32, resetAlignment NullableString, tokenLimit NullableInt32, updatedAt string, ) *OrganizationBudgetPublic`

NewOrganizationBudgetPublic instantiates a new OrganizationBudgetPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationBudgetPublicWithDefaults

`func NewOrganizationBudgetPublicWithDefaults() *OrganizationBudgetPublic`

NewOrganizationBudgetPublicWithDefaults instantiates a new OrganizationBudgetPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBudgetDurationSec

`func (o *OrganizationBudgetPublic) GetBudgetDurationSec() int32`

GetBudgetDurationSec returns the BudgetDurationSec field if non-nil, zero value otherwise.

### GetBudgetDurationSecOk

`func (o *OrganizationBudgetPublic) GetBudgetDurationSecOk() (*int32, bool)`

GetBudgetDurationSecOk returns a tuple with the BudgetDurationSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetDurationSec

`func (o *OrganizationBudgetPublic) SetBudgetDurationSec(v int32)`

SetBudgetDurationSec sets BudgetDurationSec field to given value.


### SetBudgetDurationSecNil

`func (o *OrganizationBudgetPublic) SetBudgetDurationSecNil(b bool)`

 SetBudgetDurationSecNil sets the value for BudgetDurationSec to be an explicit nil

### UnsetBudgetDurationSec
`func (o *OrganizationBudgetPublic) UnsetBudgetDurationSec()`

UnsetBudgetDurationSec ensures that no value is present for BudgetDurationSec, not even an explicit nil
### GetBudgetId

`func (o *OrganizationBudgetPublic) GetBudgetId() string`

GetBudgetId returns the BudgetId field if non-nil, zero value otherwise.

### GetBudgetIdOk

`func (o *OrganizationBudgetPublic) GetBudgetIdOk() (*string, bool)`

GetBudgetIdOk returns a tuple with the BudgetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetId

`func (o *OrganizationBudgetPublic) SetBudgetId(v string)`

SetBudgetId sets BudgetId field to given value.


### GetCeilingCount

`func (o *OrganizationBudgetPublic) GetCeilingCount() int32`

GetCeilingCount returns the CeilingCount field if non-nil, zero value otherwise.

### GetCeilingCountOk

`func (o *OrganizationBudgetPublic) GetCeilingCountOk() (*int32, bool)`

GetCeilingCountOk returns a tuple with the CeilingCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCeilingCount

`func (o *OrganizationBudgetPublic) SetCeilingCount(v int32)`

SetCeilingCount sets CeilingCount field to given value.


### GetCreatedAt

`func (o *OrganizationBudgetPublic) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OrganizationBudgetPublic) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OrganizationBudgetPublic) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetMaxBudget

`func (o *OrganizationBudgetPublic) GetMaxBudget() float32`

GetMaxBudget returns the MaxBudget field if non-nil, zero value otherwise.

### GetMaxBudgetOk

`func (o *OrganizationBudgetPublic) GetMaxBudgetOk() (*float32, bool)`

GetMaxBudgetOk returns a tuple with the MaxBudget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxBudget

`func (o *OrganizationBudgetPublic) SetMaxBudget(v float32)`

SetMaxBudget sets MaxBudget field to given value.


### SetMaxBudgetNil

`func (o *OrganizationBudgetPublic) SetMaxBudgetNil(b bool)`

 SetMaxBudgetNil sets the value for MaxBudget to be an explicit nil

### UnsetMaxBudget
`func (o *OrganizationBudgetPublic) UnsetMaxBudget()`

UnsetMaxBudget ensures that no value is present for MaxBudget, not even an explicit nil
### GetName

`func (o *OrganizationBudgetPublic) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrganizationBudgetPublic) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrganizationBudgetPublic) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *OrganizationBudgetPublic) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *OrganizationBudgetPublic) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetOrganizationId

`func (o *OrganizationBudgetPublic) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *OrganizationBudgetPublic) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *OrganizationBudgetPublic) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.


### GetRequestLimit

`func (o *OrganizationBudgetPublic) GetRequestLimit() int32`

GetRequestLimit returns the RequestLimit field if non-nil, zero value otherwise.

### GetRequestLimitOk

`func (o *OrganizationBudgetPublic) GetRequestLimitOk() (*int32, bool)`

GetRequestLimitOk returns a tuple with the RequestLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestLimit

`func (o *OrganizationBudgetPublic) SetRequestLimit(v int32)`

SetRequestLimit sets RequestLimit field to given value.


### SetRequestLimitNil

`func (o *OrganizationBudgetPublic) SetRequestLimitNil(b bool)`

 SetRequestLimitNil sets the value for RequestLimit to be an explicit nil

### UnsetRequestLimit
`func (o *OrganizationBudgetPublic) UnsetRequestLimit()`

UnsetRequestLimit ensures that no value is present for RequestLimit, not even an explicit nil
### GetResetAlignment

`func (o *OrganizationBudgetPublic) GetResetAlignment() string`

GetResetAlignment returns the ResetAlignment field if non-nil, zero value otherwise.

### GetResetAlignmentOk

`func (o *OrganizationBudgetPublic) GetResetAlignmentOk() (*string, bool)`

GetResetAlignmentOk returns a tuple with the ResetAlignment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResetAlignment

`func (o *OrganizationBudgetPublic) SetResetAlignment(v string)`

SetResetAlignment sets ResetAlignment field to given value.


### SetResetAlignmentNil

`func (o *OrganizationBudgetPublic) SetResetAlignmentNil(b bool)`

 SetResetAlignmentNil sets the value for ResetAlignment to be an explicit nil

### UnsetResetAlignment
`func (o *OrganizationBudgetPublic) UnsetResetAlignment()`

UnsetResetAlignment ensures that no value is present for ResetAlignment, not even an explicit nil
### GetTokenLimit

`func (o *OrganizationBudgetPublic) GetTokenLimit() int32`

GetTokenLimit returns the TokenLimit field if non-nil, zero value otherwise.

### GetTokenLimitOk

`func (o *OrganizationBudgetPublic) GetTokenLimitOk() (*int32, bool)`

GetTokenLimitOk returns a tuple with the TokenLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenLimit

`func (o *OrganizationBudgetPublic) SetTokenLimit(v int32)`

SetTokenLimit sets TokenLimit field to given value.


### SetTokenLimitNil

`func (o *OrganizationBudgetPublic) SetTokenLimitNil(b bool)`

 SetTokenLimitNil sets the value for TokenLimit to be an explicit nil

### UnsetTokenLimit
`func (o *OrganizationBudgetPublic) UnsetTokenLimit()`

UnsetTokenLimit ensures that no value is present for TokenLimit, not even an explicit nil
### GetUpdatedAt

`func (o *OrganizationBudgetPublic) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *OrganizationBudgetPublic) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *OrganizationBudgetPublic) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


