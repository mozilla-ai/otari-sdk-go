# MRBetaAdvisorRedactedResultBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EncryptedContent** | **string** |  | 
**StopReason** | Pointer to **NullableString** |  | [optional] 
**Type** | **string** |  | 

## Methods

### NewMRBetaAdvisorRedactedResultBlock

`func NewMRBetaAdvisorRedactedResultBlock(encryptedContent string, type_ string, ) *MRBetaAdvisorRedactedResultBlock`

NewMRBetaAdvisorRedactedResultBlock instantiates a new MRBetaAdvisorRedactedResultBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMRBetaAdvisorRedactedResultBlockWithDefaults

`func NewMRBetaAdvisorRedactedResultBlockWithDefaults() *MRBetaAdvisorRedactedResultBlock`

NewMRBetaAdvisorRedactedResultBlockWithDefaults instantiates a new MRBetaAdvisorRedactedResultBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEncryptedContent

`func (o *MRBetaAdvisorRedactedResultBlock) GetEncryptedContent() string`

GetEncryptedContent returns the EncryptedContent field if non-nil, zero value otherwise.

### GetEncryptedContentOk

`func (o *MRBetaAdvisorRedactedResultBlock) GetEncryptedContentOk() (*string, bool)`

GetEncryptedContentOk returns a tuple with the EncryptedContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedContent

`func (o *MRBetaAdvisorRedactedResultBlock) SetEncryptedContent(v string)`

SetEncryptedContent sets EncryptedContent field to given value.


### GetStopReason

`func (o *MRBetaAdvisorRedactedResultBlock) GetStopReason() string`

GetStopReason returns the StopReason field if non-nil, zero value otherwise.

### GetStopReasonOk

`func (o *MRBetaAdvisorRedactedResultBlock) GetStopReasonOk() (*string, bool)`

GetStopReasonOk returns a tuple with the StopReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopReason

`func (o *MRBetaAdvisorRedactedResultBlock) SetStopReason(v string)`

SetStopReason sets StopReason field to given value.

### HasStopReason

`func (o *MRBetaAdvisorRedactedResultBlock) HasStopReason() bool`

HasStopReason returns a boolean if a field has been set.

### SetStopReasonNil

`func (o *MRBetaAdvisorRedactedResultBlock) SetStopReasonNil(b bool)`

 SetStopReasonNil sets the value for StopReason to be an explicit nil

### UnsetStopReason
`func (o *MRBetaAdvisorRedactedResultBlock) UnsetStopReason()`

UnsetStopReason ensures that no value is present for StopReason, not even an explicit nil
### GetType

`func (o *MRBetaAdvisorRedactedResultBlock) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MRBetaAdvisorRedactedResultBlock) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MRBetaAdvisorRedactedResultBlock) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


