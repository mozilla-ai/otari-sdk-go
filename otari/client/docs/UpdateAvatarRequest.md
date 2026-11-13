# UpdateAvatarRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Face** | **NullableInt32** | The index of a preset face from the dashboard&#39;s avatar set, or null to go back to the face this identity&#39;s id gives it. | 

## Methods

### NewUpdateAvatarRequest

`func NewUpdateAvatarRequest(face NullableInt32, ) *UpdateAvatarRequest`

NewUpdateAvatarRequest instantiates a new UpdateAvatarRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateAvatarRequestWithDefaults

`func NewUpdateAvatarRequestWithDefaults() *UpdateAvatarRequest`

NewUpdateAvatarRequestWithDefaults instantiates a new UpdateAvatarRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFace

`func (o *UpdateAvatarRequest) GetFace() int32`

GetFace returns the Face field if non-nil, zero value otherwise.

### GetFaceOk

`func (o *UpdateAvatarRequest) GetFaceOk() (*int32, bool)`

GetFaceOk returns a tuple with the Face field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFace

`func (o *UpdateAvatarRequest) SetFace(v int32)`

SetFace sets Face field to given value.


### SetFaceNil

`func (o *UpdateAvatarRequest) SetFaceNil(b bool)`

 SetFaceNil sets the value for Face to be an explicit nil

### UnsetFace
`func (o *UpdateAvatarRequest) UnsetFace()`

UnsetFace ensures that no value is present for Face, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


