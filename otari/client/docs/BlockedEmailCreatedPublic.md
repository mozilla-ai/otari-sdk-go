# BlockedEmailCreatedPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlockedAccountCount** | **int32** | How many existing accounts the entry matches. Their dashboard sessions were ended. | 
**CreatedAt** | **time.Time** |  | 
**CreatedByUserId** | **NullableString** | The operator who added it, or null once removed. | 
**Id** | **string** |  | 
**Kind** | **string** | address or domain. A domain also blocks its subdomains. | 
**Pattern** | **string** | The normalized address or domain the entry blocks. | 
**Reason** | **NullableString** |  | 

## Methods

### NewBlockedEmailCreatedPublic

`func NewBlockedEmailCreatedPublic(blockedAccountCount int32, createdAt time.Time, createdByUserId NullableString, id string, kind string, pattern string, reason NullableString, ) *BlockedEmailCreatedPublic`

NewBlockedEmailCreatedPublic instantiates a new BlockedEmailCreatedPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlockedEmailCreatedPublicWithDefaults

`func NewBlockedEmailCreatedPublicWithDefaults() *BlockedEmailCreatedPublic`

NewBlockedEmailCreatedPublicWithDefaults instantiates a new BlockedEmailCreatedPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlockedAccountCount

`func (o *BlockedEmailCreatedPublic) GetBlockedAccountCount() int32`

GetBlockedAccountCount returns the BlockedAccountCount field if non-nil, zero value otherwise.

### GetBlockedAccountCountOk

`func (o *BlockedEmailCreatedPublic) GetBlockedAccountCountOk() (*int32, bool)`

GetBlockedAccountCountOk returns a tuple with the BlockedAccountCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlockedAccountCount

`func (o *BlockedEmailCreatedPublic) SetBlockedAccountCount(v int32)`

SetBlockedAccountCount sets BlockedAccountCount field to given value.


### GetCreatedAt

`func (o *BlockedEmailCreatedPublic) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BlockedEmailCreatedPublic) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BlockedEmailCreatedPublic) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedByUserId

`func (o *BlockedEmailCreatedPublic) GetCreatedByUserId() string`

GetCreatedByUserId returns the CreatedByUserId field if non-nil, zero value otherwise.

### GetCreatedByUserIdOk

`func (o *BlockedEmailCreatedPublic) GetCreatedByUserIdOk() (*string, bool)`

GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserId

`func (o *BlockedEmailCreatedPublic) SetCreatedByUserId(v string)`

SetCreatedByUserId sets CreatedByUserId field to given value.


### SetCreatedByUserIdNil

`func (o *BlockedEmailCreatedPublic) SetCreatedByUserIdNil(b bool)`

 SetCreatedByUserIdNil sets the value for CreatedByUserId to be an explicit nil

### UnsetCreatedByUserId
`func (o *BlockedEmailCreatedPublic) UnsetCreatedByUserId()`

UnsetCreatedByUserId ensures that no value is present for CreatedByUserId, not even an explicit nil
### GetId

`func (o *BlockedEmailCreatedPublic) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BlockedEmailCreatedPublic) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BlockedEmailCreatedPublic) SetId(v string)`

SetId sets Id field to given value.


### GetKind

`func (o *BlockedEmailCreatedPublic) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *BlockedEmailCreatedPublic) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *BlockedEmailCreatedPublic) SetKind(v string)`

SetKind sets Kind field to given value.


### GetPattern

`func (o *BlockedEmailCreatedPublic) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *BlockedEmailCreatedPublic) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *BlockedEmailCreatedPublic) SetPattern(v string)`

SetPattern sets Pattern field to given value.


### GetReason

`func (o *BlockedEmailCreatedPublic) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *BlockedEmailCreatedPublic) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *BlockedEmailCreatedPublic) SetReason(v string)`

SetReason sets Reason field to given value.


### SetReasonNil

`func (o *BlockedEmailCreatedPublic) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *BlockedEmailCreatedPublic) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


