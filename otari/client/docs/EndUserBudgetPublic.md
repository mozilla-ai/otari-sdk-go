# EndUserBudgetPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockedEndUserCount** | **int32** |  | 
**BudgetId** | **string** |  | 
**EndUserCount** | **int32** | End users of the key&#39;s user on this budget, blocked ones included | 
**IsDefault** | **bool** | Whether a new end user starts here when its request names no budget | 

## Methods

### NewEndUserBudgetPublic

`func NewEndUserBudgetPublic(blockedEndUserCount int32, budgetId string, endUserCount int32, isDefault bool, ) *EndUserBudgetPublic`

NewEndUserBudgetPublic instantiates a new EndUserBudgetPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndUserBudgetPublicWithDefaults

`func NewEndUserBudgetPublicWithDefaults() *EndUserBudgetPublic`

NewEndUserBudgetPublicWithDefaults instantiates a new EndUserBudgetPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockedEndUserCount

`func (o *EndUserBudgetPublic) GetBlockedEndUserCount() int32`

GetBlockedEndUserCount returns the BlockedEndUserCount field if non-nil, zero value otherwise.

### GetBlockedEndUserCountOk

`func (o *EndUserBudgetPublic) GetBlockedEndUserCountOk() (*int32, bool)`

GetBlockedEndUserCountOk returns a tuple with the BlockedEndUserCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedEndUserCount

`func (o *EndUserBudgetPublic) SetBlockedEndUserCount(v int32)`

SetBlockedEndUserCount sets BlockedEndUserCount field to given value.


### GetBudgetId

`func (o *EndUserBudgetPublic) GetBudgetId() string`

GetBudgetId returns the BudgetId field if non-nil, zero value otherwise.

### GetBudgetIdOk

`func (o *EndUserBudgetPublic) GetBudgetIdOk() (*string, bool)`

GetBudgetIdOk returns a tuple with the BudgetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetId

`func (o *EndUserBudgetPublic) SetBudgetId(v string)`

SetBudgetId sets BudgetId field to given value.


### GetEndUserCount

`func (o *EndUserBudgetPublic) GetEndUserCount() int32`

GetEndUserCount returns the EndUserCount field if non-nil, zero value otherwise.

### GetEndUserCountOk

`func (o *EndUserBudgetPublic) GetEndUserCountOk() (*int32, bool)`

GetEndUserCountOk returns a tuple with the EndUserCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndUserCount

`func (o *EndUserBudgetPublic) SetEndUserCount(v int32)`

SetEndUserCount sets EndUserCount field to given value.


### GetIsDefault

`func (o *EndUserBudgetPublic) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *EndUserBudgetPublic) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *EndUserBudgetPublic) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


