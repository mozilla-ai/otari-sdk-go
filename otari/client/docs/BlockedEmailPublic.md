# BlockedEmailPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | **time.Time** |  | 
**CreatedByUserId** | **NullableString** | The operator who added it, or null once removed. | 
**Id** | **string** |  | 
**Kind** | **string** | address or domain. A domain also blocks its subdomains. | 
**Pattern** | **string** | The normalized address or domain the entry blocks. | 
**Reason** | **NullableString** |  | 

## Methods

### NewBlockedEmailPublic

`func NewBlockedEmailPublic(createdAt time.Time, createdByUserId NullableString, id string, kind string, pattern string, reason NullableString, ) *BlockedEmailPublic`

NewBlockedEmailPublic instantiates a new BlockedEmailPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlockedEmailPublicWithDefaults

`func NewBlockedEmailPublicWithDefaults() *BlockedEmailPublic`

NewBlockedEmailPublicWithDefaults instantiates a new BlockedEmailPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *BlockedEmailPublic) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BlockedEmailPublic) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BlockedEmailPublic) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetCreatedByUserId

`func (o *BlockedEmailPublic) GetCreatedByUserId() string`

GetCreatedByUserId returns the CreatedByUserId field if non-nil, zero value otherwise.

### GetCreatedByUserIdOk

`func (o *BlockedEmailPublic) GetCreatedByUserIdOk() (*string, bool)`

GetCreatedByUserIdOk returns a tuple with the CreatedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedByUserId

`func (o *BlockedEmailPublic) SetCreatedByUserId(v string)`

SetCreatedByUserId sets CreatedByUserId field to given value.


### SetCreatedByUserIdNil

`func (o *BlockedEmailPublic) SetCreatedByUserIdNil(b bool)`

 SetCreatedByUserIdNil sets the value for CreatedByUserId to be an explicit nil

### UnsetCreatedByUserId
`func (o *BlockedEmailPublic) UnsetCreatedByUserId()`

UnsetCreatedByUserId ensures that no value is present for CreatedByUserId, not even an explicit nil
### GetId

`func (o *BlockedEmailPublic) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BlockedEmailPublic) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BlockedEmailPublic) SetId(v string)`

SetId sets Id field to given value.


### GetKind

`func (o *BlockedEmailPublic) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *BlockedEmailPublic) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *BlockedEmailPublic) SetKind(v string)`

SetKind sets Kind field to given value.


### GetPattern

`func (o *BlockedEmailPublic) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *BlockedEmailPublic) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *BlockedEmailPublic) SetPattern(v string)`

SetPattern sets Pattern field to given value.


### GetReason

`func (o *BlockedEmailPublic) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *BlockedEmailPublic) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *BlockedEmailPublic) SetReason(v string)`

SetReason sets Reason field to given value.


### SetReasonNil

`func (o *BlockedEmailPublic) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *BlockedEmailPublic) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


