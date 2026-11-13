# BlockedEmailCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reason** | Pointer to **NullableString** | A note for other operators on why the entry exists. Never shown to the blocked person. | [optional] 
**Value** | **string** | An address (person@example.com) or a domain (example.com, also accepted as @example.com or *.example.com). A domain also blocks its subdomains. | 

## Methods

### NewBlockedEmailCreateRequest

`func NewBlockedEmailCreateRequest(value string, ) *BlockedEmailCreateRequest`

NewBlockedEmailCreateRequest instantiates a new BlockedEmailCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlockedEmailCreateRequestWithDefaults

`func NewBlockedEmailCreateRequestWithDefaults() *BlockedEmailCreateRequest`

NewBlockedEmailCreateRequestWithDefaults instantiates a new BlockedEmailCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReason

`func (o *BlockedEmailCreateRequest) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *BlockedEmailCreateRequest) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *BlockedEmailCreateRequest) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *BlockedEmailCreateRequest) HasReason() bool`

HasReason returns a boolean if a field has been set.

### SetReasonNil

`func (o *BlockedEmailCreateRequest) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *BlockedEmailCreateRequest) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil
### GetValue

`func (o *BlockedEmailCreateRequest) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *BlockedEmailCreateRequest) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *BlockedEmailCreateRequest) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


