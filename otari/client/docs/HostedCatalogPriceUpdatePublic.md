# HostedCatalogPriceUpdatePublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddedCount** | **int32** |  | 
**ChangedCount** | **int32** |  | 
**Changes** | [**[]HostedCatalogPriceChangePublic**](HostedCatalogPriceChangePublic.md) |  | 
**ChangesTruncated** | **bool** |  | 
**Digest** | **string** |  | 
**FetchedAt** | **time.Time** |  | 
**NeedsReview** | **bool** |  | 
**RemovedCount** | **int32** |  | 
**ReviewReason** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewHostedCatalogPriceUpdatePublic

`func NewHostedCatalogPriceUpdatePublic(addedCount int32, changedCount int32, changes []HostedCatalogPriceChangePublic, changesTruncated bool, digest string, fetchedAt time.Time, needsReview bool, removedCount int32, ) *HostedCatalogPriceUpdatePublic`

NewHostedCatalogPriceUpdatePublic instantiates a new HostedCatalogPriceUpdatePublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostedCatalogPriceUpdatePublicWithDefaults

`func NewHostedCatalogPriceUpdatePublicWithDefaults() *HostedCatalogPriceUpdatePublic`

NewHostedCatalogPriceUpdatePublicWithDefaults instantiates a new HostedCatalogPriceUpdatePublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddedCount

`func (o *HostedCatalogPriceUpdatePublic) GetAddedCount() int32`

GetAddedCount returns the AddedCount field if non-nil, zero value otherwise.

### GetAddedCountOk

`func (o *HostedCatalogPriceUpdatePublic) GetAddedCountOk() (*int32, bool)`

GetAddedCountOk returns a tuple with the AddedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedCount

`func (o *HostedCatalogPriceUpdatePublic) SetAddedCount(v int32)`

SetAddedCount sets AddedCount field to given value.


### GetChangedCount

`func (o *HostedCatalogPriceUpdatePublic) GetChangedCount() int32`

GetChangedCount returns the ChangedCount field if non-nil, zero value otherwise.

### GetChangedCountOk

`func (o *HostedCatalogPriceUpdatePublic) GetChangedCountOk() (*int32, bool)`

GetChangedCountOk returns a tuple with the ChangedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangedCount

`func (o *HostedCatalogPriceUpdatePublic) SetChangedCount(v int32)`

SetChangedCount sets ChangedCount field to given value.


### GetChanges

`func (o *HostedCatalogPriceUpdatePublic) GetChanges() []HostedCatalogPriceChangePublic`

GetChanges returns the Changes field if non-nil, zero value otherwise.

### GetChangesOk

`func (o *HostedCatalogPriceUpdatePublic) GetChangesOk() (*[]HostedCatalogPriceChangePublic, bool)`

GetChangesOk returns a tuple with the Changes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChanges

`func (o *HostedCatalogPriceUpdatePublic) SetChanges(v []HostedCatalogPriceChangePublic)`

SetChanges sets Changes field to given value.


### GetChangesTruncated

`func (o *HostedCatalogPriceUpdatePublic) GetChangesTruncated() bool`

GetChangesTruncated returns the ChangesTruncated field if non-nil, zero value otherwise.

### GetChangesTruncatedOk

`func (o *HostedCatalogPriceUpdatePublic) GetChangesTruncatedOk() (*bool, bool)`

GetChangesTruncatedOk returns a tuple with the ChangesTruncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangesTruncated

`func (o *HostedCatalogPriceUpdatePublic) SetChangesTruncated(v bool)`

SetChangesTruncated sets ChangesTruncated field to given value.


### GetDigest

`func (o *HostedCatalogPriceUpdatePublic) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *HostedCatalogPriceUpdatePublic) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *HostedCatalogPriceUpdatePublic) SetDigest(v string)`

SetDigest sets Digest field to given value.


### GetFetchedAt

`func (o *HostedCatalogPriceUpdatePublic) GetFetchedAt() time.Time`

GetFetchedAt returns the FetchedAt field if non-nil, zero value otherwise.

### GetFetchedAtOk

`func (o *HostedCatalogPriceUpdatePublic) GetFetchedAtOk() (*time.Time, bool)`

GetFetchedAtOk returns a tuple with the FetchedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFetchedAt

`func (o *HostedCatalogPriceUpdatePublic) SetFetchedAt(v time.Time)`

SetFetchedAt sets FetchedAt field to given value.


### GetNeedsReview

`func (o *HostedCatalogPriceUpdatePublic) GetNeedsReview() bool`

GetNeedsReview returns the NeedsReview field if non-nil, zero value otherwise.

### GetNeedsReviewOk

`func (o *HostedCatalogPriceUpdatePublic) GetNeedsReviewOk() (*bool, bool)`

GetNeedsReviewOk returns a tuple with the NeedsReview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedsReview

`func (o *HostedCatalogPriceUpdatePublic) SetNeedsReview(v bool)`

SetNeedsReview sets NeedsReview field to given value.


### GetRemovedCount

`func (o *HostedCatalogPriceUpdatePublic) GetRemovedCount() int32`

GetRemovedCount returns the RemovedCount field if non-nil, zero value otherwise.

### GetRemovedCountOk

`func (o *HostedCatalogPriceUpdatePublic) GetRemovedCountOk() (*int32, bool)`

GetRemovedCountOk returns a tuple with the RemovedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovedCount

`func (o *HostedCatalogPriceUpdatePublic) SetRemovedCount(v int32)`

SetRemovedCount sets RemovedCount field to given value.


### GetReviewReason

`func (o *HostedCatalogPriceUpdatePublic) GetReviewReason() string`

GetReviewReason returns the ReviewReason field if non-nil, zero value otherwise.

### GetReviewReasonOk

`func (o *HostedCatalogPriceUpdatePublic) GetReviewReasonOk() (*string, bool)`

GetReviewReasonOk returns a tuple with the ReviewReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewReason

`func (o *HostedCatalogPriceUpdatePublic) SetReviewReason(v string)`

SetReviewReason sets ReviewReason field to given value.

### HasReviewReason

`func (o *HostedCatalogPriceUpdatePublic) HasReviewReason() bool`

HasReviewReason returns a boolean if a field has been set.

### SetReviewReasonNil

`func (o *HostedCatalogPriceUpdatePublic) SetReviewReasonNil(b bool)`

 SetReviewReasonNil sets the value for ReviewReason to be an explicit nil

### UnsetReviewReason
`func (o *HostedCatalogPriceUpdatePublic) UnsetReviewReason()`

UnsetReviewReason ensures that no value is present for ReviewReason, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


