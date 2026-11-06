# PricingRefreshPreviewResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddedCount** | **int32** |  | 
**ChangedCount** | **int32** |  | 
**Changes** | [**[]PricingRefreshChangeResponse**](PricingRefreshChangeResponse.md) |  | 
**ChangesTruncated** | **bool** |  | 
**Digest** | Pointer to **string** | Identity of the pending snapshot, to act on exactly this one: send it as &#x60;digest&#x60; to &#x60;POST /hosted-providers/catalog-update&#x60; to apply it, or to &#x60;POST /pricing/refresh/reject&#x60; to discard it. | [optional] [default to ""]
**FetchedAt** | **time.Time** |  | 
**NeedsReview** | Pointer to **bool** | True when the update is implausibly large; the scheduled &#x60;auto&#x60; policy leaves it pending. | [optional] [default to false]
**ProtectedModelCount** | **int32** |  | 
**RemovedCount** | **int32** |  | 
**ReviewReason** | Pointer to **NullableString** | Why the update needs review. | [optional] 

## Methods

### NewPricingRefreshPreviewResponse

`func NewPricingRefreshPreviewResponse(addedCount int32, changedCount int32, changes []PricingRefreshChangeResponse, changesTruncated bool, fetchedAt time.Time, protectedModelCount int32, removedCount int32, ) *PricingRefreshPreviewResponse`

NewPricingRefreshPreviewResponse instantiates a new PricingRefreshPreviewResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPricingRefreshPreviewResponseWithDefaults

`func NewPricingRefreshPreviewResponseWithDefaults() *PricingRefreshPreviewResponse`

NewPricingRefreshPreviewResponseWithDefaults instantiates a new PricingRefreshPreviewResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddedCount

`func (o *PricingRefreshPreviewResponse) GetAddedCount() int32`

GetAddedCount returns the AddedCount field if non-nil, zero value otherwise.

### GetAddedCountOk

`func (o *PricingRefreshPreviewResponse) GetAddedCountOk() (*int32, bool)`

GetAddedCountOk returns a tuple with the AddedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedCount

`func (o *PricingRefreshPreviewResponse) SetAddedCount(v int32)`

SetAddedCount sets AddedCount field to given value.


### GetChangedCount

`func (o *PricingRefreshPreviewResponse) GetChangedCount() int32`

GetChangedCount returns the ChangedCount field if non-nil, zero value otherwise.

### GetChangedCountOk

`func (o *PricingRefreshPreviewResponse) GetChangedCountOk() (*int32, bool)`

GetChangedCountOk returns a tuple with the ChangedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangedCount

`func (o *PricingRefreshPreviewResponse) SetChangedCount(v int32)`

SetChangedCount sets ChangedCount field to given value.


### GetChanges

`func (o *PricingRefreshPreviewResponse) GetChanges() []PricingRefreshChangeResponse`

GetChanges returns the Changes field if non-nil, zero value otherwise.

### GetChangesOk

`func (o *PricingRefreshPreviewResponse) GetChangesOk() (*[]PricingRefreshChangeResponse, bool)`

GetChangesOk returns a tuple with the Changes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChanges

`func (o *PricingRefreshPreviewResponse) SetChanges(v []PricingRefreshChangeResponse)`

SetChanges sets Changes field to given value.


### GetChangesTruncated

`func (o *PricingRefreshPreviewResponse) GetChangesTruncated() bool`

GetChangesTruncated returns the ChangesTruncated field if non-nil, zero value otherwise.

### GetChangesTruncatedOk

`func (o *PricingRefreshPreviewResponse) GetChangesTruncatedOk() (*bool, bool)`

GetChangesTruncatedOk returns a tuple with the ChangesTruncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangesTruncated

`func (o *PricingRefreshPreviewResponse) SetChangesTruncated(v bool)`

SetChangesTruncated sets ChangesTruncated field to given value.


### GetDigest

`func (o *PricingRefreshPreviewResponse) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *PricingRefreshPreviewResponse) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *PricingRefreshPreviewResponse) SetDigest(v string)`

SetDigest sets Digest field to given value.

### HasDigest

`func (o *PricingRefreshPreviewResponse) HasDigest() bool`

HasDigest returns a boolean if a field has been set.

### GetFetchedAt

`func (o *PricingRefreshPreviewResponse) GetFetchedAt() time.Time`

GetFetchedAt returns the FetchedAt field if non-nil, zero value otherwise.

### GetFetchedAtOk

`func (o *PricingRefreshPreviewResponse) GetFetchedAtOk() (*time.Time, bool)`

GetFetchedAtOk returns a tuple with the FetchedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFetchedAt

`func (o *PricingRefreshPreviewResponse) SetFetchedAt(v time.Time)`

SetFetchedAt sets FetchedAt field to given value.


### GetNeedsReview

`func (o *PricingRefreshPreviewResponse) GetNeedsReview() bool`

GetNeedsReview returns the NeedsReview field if non-nil, zero value otherwise.

### GetNeedsReviewOk

`func (o *PricingRefreshPreviewResponse) GetNeedsReviewOk() (*bool, bool)`

GetNeedsReviewOk returns a tuple with the NeedsReview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedsReview

`func (o *PricingRefreshPreviewResponse) SetNeedsReview(v bool)`

SetNeedsReview sets NeedsReview field to given value.

### HasNeedsReview

`func (o *PricingRefreshPreviewResponse) HasNeedsReview() bool`

HasNeedsReview returns a boolean if a field has been set.

### GetProtectedModelCount

`func (o *PricingRefreshPreviewResponse) GetProtectedModelCount() int32`

GetProtectedModelCount returns the ProtectedModelCount field if non-nil, zero value otherwise.

### GetProtectedModelCountOk

`func (o *PricingRefreshPreviewResponse) GetProtectedModelCountOk() (*int32, bool)`

GetProtectedModelCountOk returns a tuple with the ProtectedModelCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtectedModelCount

`func (o *PricingRefreshPreviewResponse) SetProtectedModelCount(v int32)`

SetProtectedModelCount sets ProtectedModelCount field to given value.


### GetRemovedCount

`func (o *PricingRefreshPreviewResponse) GetRemovedCount() int32`

GetRemovedCount returns the RemovedCount field if non-nil, zero value otherwise.

### GetRemovedCountOk

`func (o *PricingRefreshPreviewResponse) GetRemovedCountOk() (*int32, bool)`

GetRemovedCountOk returns a tuple with the RemovedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovedCount

`func (o *PricingRefreshPreviewResponse) SetRemovedCount(v int32)`

SetRemovedCount sets RemovedCount field to given value.


### GetReviewReason

`func (o *PricingRefreshPreviewResponse) GetReviewReason() string`

GetReviewReason returns the ReviewReason field if non-nil, zero value otherwise.

### GetReviewReasonOk

`func (o *PricingRefreshPreviewResponse) GetReviewReasonOk() (*string, bool)`

GetReviewReasonOk returns a tuple with the ReviewReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewReason

`func (o *PricingRefreshPreviewResponse) SetReviewReason(v string)`

SetReviewReason sets ReviewReason field to given value.

### HasReviewReason

`func (o *PricingRefreshPreviewResponse) HasReviewReason() bool`

HasReviewReason returns a boolean if a field has been set.

### SetReviewReasonNil

`func (o *PricingRefreshPreviewResponse) SetReviewReasonNil(b bool)`

 SetReviewReasonNil sets the value for ReviewReason to be an explicit nil

### UnsetReviewReason
`func (o *PricingRefreshPreviewResponse) UnsetReviewReason()`

UnsetReviewReason ensures that no value is present for ReviewReason, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


