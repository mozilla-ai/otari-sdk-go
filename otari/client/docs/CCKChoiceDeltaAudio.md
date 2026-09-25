# CCKChoiceDeltaAudio

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | Filter to a single event type or metric name (e.g. &#39;tool_result&#39;, &#39;claude_code.commit.count&#39;) | [optional] 
**Data** | Pointer to **NullableString** |  | [optional] 
**Transcript** | Pointer to **NullableString** |  | [optional] 
**ExpiresAt** | Pointer to **NullableInt32** |  | [optional] 

## Methods

### NewCCKChoiceDeltaAudio

`func NewCCKChoiceDeltaAudio() *CCKChoiceDeltaAudio`

NewCCKChoiceDeltaAudio instantiates a new CCKChoiceDeltaAudio object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCCKChoiceDeltaAudioWithDefaults

`func NewCCKChoiceDeltaAudioWithDefaults() *CCKChoiceDeltaAudio`

NewCCKChoiceDeltaAudioWithDefaults instantiates a new CCKChoiceDeltaAudio object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CCKChoiceDeltaAudio) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CCKChoiceDeltaAudio) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CCKChoiceDeltaAudio) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CCKChoiceDeltaAudio) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *CCKChoiceDeltaAudio) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *CCKChoiceDeltaAudio) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetData

`func (o *CCKChoiceDeltaAudio) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CCKChoiceDeltaAudio) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CCKChoiceDeltaAudio) SetData(v string)`

SetData sets Data field to given value.

### HasData

`func (o *CCKChoiceDeltaAudio) HasData() bool`

HasData returns a boolean if a field has been set.

### SetDataNil

`func (o *CCKChoiceDeltaAudio) SetDataNil(b bool)`

 SetDataNil sets the value for Data to be an explicit nil

### UnsetData
`func (o *CCKChoiceDeltaAudio) UnsetData()`

UnsetData ensures that no value is present for Data, not even an explicit nil
### GetTranscript

`func (o *CCKChoiceDeltaAudio) GetTranscript() string`

GetTranscript returns the Transcript field if non-nil, zero value otherwise.

### GetTranscriptOk

`func (o *CCKChoiceDeltaAudio) GetTranscriptOk() (*string, bool)`

GetTranscriptOk returns a tuple with the Transcript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranscript

`func (o *CCKChoiceDeltaAudio) SetTranscript(v string)`

SetTranscript sets Transcript field to given value.

### HasTranscript

`func (o *CCKChoiceDeltaAudio) HasTranscript() bool`

HasTranscript returns a boolean if a field has been set.

### SetTranscriptNil

`func (o *CCKChoiceDeltaAudio) SetTranscriptNil(b bool)`

 SetTranscriptNil sets the value for Transcript to be an explicit nil

### UnsetTranscript
`func (o *CCKChoiceDeltaAudio) UnsetTranscript()`

UnsetTranscript ensures that no value is present for Transcript, not even an explicit nil
### GetExpiresAt

`func (o *CCKChoiceDeltaAudio) GetExpiresAt() int32`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *CCKChoiceDeltaAudio) GetExpiresAtOk() (*int32, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *CCKChoiceDeltaAudio) SetExpiresAt(v int32)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *CCKChoiceDeltaAudio) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### SetExpiresAtNil

`func (o *CCKChoiceDeltaAudio) SetExpiresAtNil(b bool)`

 SetExpiresAtNil sets the value for ExpiresAt to be an explicit nil

### UnsetExpiresAt
`func (o *CCKChoiceDeltaAudio) UnsetExpiresAt()`

UnsetExpiresAt ensures that no value is present for ExpiresAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


