# ContentSharePublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | **bool** | Whether it still lets its audience read. | 
**Audience** | **string** | Who may read: the organization&#39;s owners and admins, or the deployment&#39;s operator. | 
**CreatedAt** | **time.Time** |  | 
**Id** | **string** |  | 
**Note** | **NullableString** | Why it was shared, in its owner&#39;s words. | 
**ReadCount** | Pointer to **int32** | How many reads it has had in all. | [optional] [default to 0]
**Reads** | Pointer to [**[]ContentReadPublic**](ContentReadPublic.md) | Its newest reads, newest first; GET .../shares/{share_id}/reads pages through all of them. A reader&#39;s reads are recorded at most once an hour. Only its owner sees them. | [optional] 
**RevokedAt** | **NullableTime** | When it ended: its owner ended it, or lost the workspace. | 

## Methods

### NewContentSharePublic

`func NewContentSharePublic(active bool, audience string, createdAt time.Time, id string, note NullableString, revokedAt NullableTime, ) *ContentSharePublic`

NewContentSharePublic instantiates a new ContentSharePublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentSharePublicWithDefaults

`func NewContentSharePublicWithDefaults() *ContentSharePublic`

NewContentSharePublicWithDefaults instantiates a new ContentSharePublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *ContentSharePublic) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ContentSharePublic) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ContentSharePublic) SetActive(v bool)`

SetActive sets Active field to given value.


### GetAudience

`func (o *ContentSharePublic) GetAudience() string`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *ContentSharePublic) GetAudienceOk() (*string, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *ContentSharePublic) SetAudience(v string)`

SetAudience sets Audience field to given value.


### GetCreatedAt

`func (o *ContentSharePublic) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ContentSharePublic) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ContentSharePublic) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetId

`func (o *ContentSharePublic) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ContentSharePublic) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ContentSharePublic) SetId(v string)`

SetId sets Id field to given value.


### GetNote

`func (o *ContentSharePublic) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *ContentSharePublic) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *ContentSharePublic) SetNote(v string)`

SetNote sets Note field to given value.


### SetNoteNil

`func (o *ContentSharePublic) SetNoteNil(b bool)`

 SetNoteNil sets the value for Note to be an explicit nil

### UnsetNote
`func (o *ContentSharePublic) UnsetNote()`

UnsetNote ensures that no value is present for Note, not even an explicit nil
### GetReadCount

`func (o *ContentSharePublic) GetReadCount() int32`

GetReadCount returns the ReadCount field if non-nil, zero value otherwise.

### GetReadCountOk

`func (o *ContentSharePublic) GetReadCountOk() (*int32, bool)`

GetReadCountOk returns a tuple with the ReadCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadCount

`func (o *ContentSharePublic) SetReadCount(v int32)`

SetReadCount sets ReadCount field to given value.

### HasReadCount

`func (o *ContentSharePublic) HasReadCount() bool`

HasReadCount returns a boolean if a field has been set.

### GetReads

`func (o *ContentSharePublic) GetReads() []ContentReadPublic`

GetReads returns the Reads field if non-nil, zero value otherwise.

### GetReadsOk

`func (o *ContentSharePublic) GetReadsOk() (*[]ContentReadPublic, bool)`

GetReadsOk returns a tuple with the Reads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReads

`func (o *ContentSharePublic) SetReads(v []ContentReadPublic)`

SetReads sets Reads field to given value.

### HasReads

`func (o *ContentSharePublic) HasReads() bool`

HasReads returns a boolean if a field has been set.

### GetRevokedAt

`func (o *ContentSharePublic) GetRevokedAt() time.Time`

GetRevokedAt returns the RevokedAt field if non-nil, zero value otherwise.

### GetRevokedAtOk

`func (o *ContentSharePublic) GetRevokedAtOk() (*time.Time, bool)`

GetRevokedAtOk returns a tuple with the RevokedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokedAt

`func (o *ContentSharePublic) SetRevokedAt(v time.Time)`

SetRevokedAt sets RevokedAt field to given value.


### SetRevokedAtNil

`func (o *ContentSharePublic) SetRevokedAtNil(b bool)`

 SetRevokedAtNil sets the value for RevokedAt to be an explicit nil

### UnsetRevokedAt
`func (o *ContentSharePublic) UnsetRevokedAt()`

UnsetRevokedAt ensures that no value is present for RevokedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


