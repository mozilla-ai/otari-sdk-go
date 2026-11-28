# VideoGenerationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AspectRatio** | Pointer to **NullableString** |  | [optional] 
**FirstFrame** | Pointer to **NullableString** |  | [optional] 
**GenerateAudio** | Pointer to **NullableBool** |  | [optional] 
**LastFrame** | Pointer to **NullableString** |  | [optional] 
**Model** | **string** |  | 
**NegativePrompt** | Pointer to **NullableString** |  | [optional] 
**Prompt** | **string** |  | 
**ReferenceAudios** | Pointer to **[]string** |  | [optional] 
**ReferenceImages** | Pointer to **[]string** |  | [optional] 
**ReferenceVideos** | Pointer to **[]string** |  | [optional] 
**Resolution** | Pointer to **NullableString** |  | [optional] 
**Seconds** | Pointer to **NullableInt32** |  | [optional] 
**Seed** | Pointer to **NullableInt32** |  | [optional] 
**Size** | Pointer to **NullableString** |  | [optional] 
**User** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewVideoGenerationRequest

`func NewVideoGenerationRequest(model string, prompt string, ) *VideoGenerationRequest`

NewVideoGenerationRequest instantiates a new VideoGenerationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVideoGenerationRequestWithDefaults

`func NewVideoGenerationRequestWithDefaults() *VideoGenerationRequest`

NewVideoGenerationRequestWithDefaults instantiates a new VideoGenerationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAspectRatio

`func (o *VideoGenerationRequest) GetAspectRatio() string`

GetAspectRatio returns the AspectRatio field if non-nil, zero value otherwise.

### GetAspectRatioOk

`func (o *VideoGenerationRequest) GetAspectRatioOk() (*string, bool)`

GetAspectRatioOk returns a tuple with the AspectRatio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAspectRatio

`func (o *VideoGenerationRequest) SetAspectRatio(v string)`

SetAspectRatio sets AspectRatio field to given value.

### HasAspectRatio

`func (o *VideoGenerationRequest) HasAspectRatio() bool`

HasAspectRatio returns a boolean if a field has been set.

### SetAspectRatioNil

`func (o *VideoGenerationRequest) SetAspectRatioNil(b bool)`

 SetAspectRatioNil sets the value for AspectRatio to be an explicit nil

### UnsetAspectRatio
`func (o *VideoGenerationRequest) UnsetAspectRatio()`

UnsetAspectRatio ensures that no value is present for AspectRatio, not even an explicit nil
### GetFirstFrame

`func (o *VideoGenerationRequest) GetFirstFrame() string`

GetFirstFrame returns the FirstFrame field if non-nil, zero value otherwise.

### GetFirstFrameOk

`func (o *VideoGenerationRequest) GetFirstFrameOk() (*string, bool)`

GetFirstFrameOk returns a tuple with the FirstFrame field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstFrame

`func (o *VideoGenerationRequest) SetFirstFrame(v string)`

SetFirstFrame sets FirstFrame field to given value.

### HasFirstFrame

`func (o *VideoGenerationRequest) HasFirstFrame() bool`

HasFirstFrame returns a boolean if a field has been set.

### SetFirstFrameNil

`func (o *VideoGenerationRequest) SetFirstFrameNil(b bool)`

 SetFirstFrameNil sets the value for FirstFrame to be an explicit nil

### UnsetFirstFrame
`func (o *VideoGenerationRequest) UnsetFirstFrame()`

UnsetFirstFrame ensures that no value is present for FirstFrame, not even an explicit nil
### GetGenerateAudio

`func (o *VideoGenerationRequest) GetGenerateAudio() bool`

GetGenerateAudio returns the GenerateAudio field if non-nil, zero value otherwise.

### GetGenerateAudioOk

`func (o *VideoGenerationRequest) GetGenerateAudioOk() (*bool, bool)`

GetGenerateAudioOk returns a tuple with the GenerateAudio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerateAudio

`func (o *VideoGenerationRequest) SetGenerateAudio(v bool)`

SetGenerateAudio sets GenerateAudio field to given value.

### HasGenerateAudio

`func (o *VideoGenerationRequest) HasGenerateAudio() bool`

HasGenerateAudio returns a boolean if a field has been set.

### SetGenerateAudioNil

`func (o *VideoGenerationRequest) SetGenerateAudioNil(b bool)`

 SetGenerateAudioNil sets the value for GenerateAudio to be an explicit nil

### UnsetGenerateAudio
`func (o *VideoGenerationRequest) UnsetGenerateAudio()`

UnsetGenerateAudio ensures that no value is present for GenerateAudio, not even an explicit nil
### GetLastFrame

`func (o *VideoGenerationRequest) GetLastFrame() string`

GetLastFrame returns the LastFrame field if non-nil, zero value otherwise.

### GetLastFrameOk

`func (o *VideoGenerationRequest) GetLastFrameOk() (*string, bool)`

GetLastFrameOk returns a tuple with the LastFrame field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFrame

`func (o *VideoGenerationRequest) SetLastFrame(v string)`

SetLastFrame sets LastFrame field to given value.

### HasLastFrame

`func (o *VideoGenerationRequest) HasLastFrame() bool`

HasLastFrame returns a boolean if a field has been set.

### SetLastFrameNil

`func (o *VideoGenerationRequest) SetLastFrameNil(b bool)`

 SetLastFrameNil sets the value for LastFrame to be an explicit nil

### UnsetLastFrame
`func (o *VideoGenerationRequest) UnsetLastFrame()`

UnsetLastFrame ensures that no value is present for LastFrame, not even an explicit nil
### GetModel

`func (o *VideoGenerationRequest) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *VideoGenerationRequest) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *VideoGenerationRequest) SetModel(v string)`

SetModel sets Model field to given value.


### GetNegativePrompt

`func (o *VideoGenerationRequest) GetNegativePrompt() string`

GetNegativePrompt returns the NegativePrompt field if non-nil, zero value otherwise.

### GetNegativePromptOk

`func (o *VideoGenerationRequest) GetNegativePromptOk() (*string, bool)`

GetNegativePromptOk returns a tuple with the NegativePrompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNegativePrompt

`func (o *VideoGenerationRequest) SetNegativePrompt(v string)`

SetNegativePrompt sets NegativePrompt field to given value.

### HasNegativePrompt

`func (o *VideoGenerationRequest) HasNegativePrompt() bool`

HasNegativePrompt returns a boolean if a field has been set.

### SetNegativePromptNil

`func (o *VideoGenerationRequest) SetNegativePromptNil(b bool)`

 SetNegativePromptNil sets the value for NegativePrompt to be an explicit nil

### UnsetNegativePrompt
`func (o *VideoGenerationRequest) UnsetNegativePrompt()`

UnsetNegativePrompt ensures that no value is present for NegativePrompt, not even an explicit nil
### GetPrompt

`func (o *VideoGenerationRequest) GetPrompt() string`

GetPrompt returns the Prompt field if non-nil, zero value otherwise.

### GetPromptOk

`func (o *VideoGenerationRequest) GetPromptOk() (*string, bool)`

GetPromptOk returns a tuple with the Prompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompt

`func (o *VideoGenerationRequest) SetPrompt(v string)`

SetPrompt sets Prompt field to given value.


### GetReferenceAudios

`func (o *VideoGenerationRequest) GetReferenceAudios() []string`

GetReferenceAudios returns the ReferenceAudios field if non-nil, zero value otherwise.

### GetReferenceAudiosOk

`func (o *VideoGenerationRequest) GetReferenceAudiosOk() (*[]string, bool)`

GetReferenceAudiosOk returns a tuple with the ReferenceAudios field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceAudios

`func (o *VideoGenerationRequest) SetReferenceAudios(v []string)`

SetReferenceAudios sets ReferenceAudios field to given value.

### HasReferenceAudios

`func (o *VideoGenerationRequest) HasReferenceAudios() bool`

HasReferenceAudios returns a boolean if a field has been set.

### SetReferenceAudiosNil

`func (o *VideoGenerationRequest) SetReferenceAudiosNil(b bool)`

 SetReferenceAudiosNil sets the value for ReferenceAudios to be an explicit nil

### UnsetReferenceAudios
`func (o *VideoGenerationRequest) UnsetReferenceAudios()`

UnsetReferenceAudios ensures that no value is present for ReferenceAudios, not even an explicit nil
### GetReferenceImages

`func (o *VideoGenerationRequest) GetReferenceImages() []string`

GetReferenceImages returns the ReferenceImages field if non-nil, zero value otherwise.

### GetReferenceImagesOk

`func (o *VideoGenerationRequest) GetReferenceImagesOk() (*[]string, bool)`

GetReferenceImagesOk returns a tuple with the ReferenceImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceImages

`func (o *VideoGenerationRequest) SetReferenceImages(v []string)`

SetReferenceImages sets ReferenceImages field to given value.

### HasReferenceImages

`func (o *VideoGenerationRequest) HasReferenceImages() bool`

HasReferenceImages returns a boolean if a field has been set.

### SetReferenceImagesNil

`func (o *VideoGenerationRequest) SetReferenceImagesNil(b bool)`

 SetReferenceImagesNil sets the value for ReferenceImages to be an explicit nil

### UnsetReferenceImages
`func (o *VideoGenerationRequest) UnsetReferenceImages()`

UnsetReferenceImages ensures that no value is present for ReferenceImages, not even an explicit nil
### GetReferenceVideos

`func (o *VideoGenerationRequest) GetReferenceVideos() []string`

GetReferenceVideos returns the ReferenceVideos field if non-nil, zero value otherwise.

### GetReferenceVideosOk

`func (o *VideoGenerationRequest) GetReferenceVideosOk() (*[]string, bool)`

GetReferenceVideosOk returns a tuple with the ReferenceVideos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceVideos

`func (o *VideoGenerationRequest) SetReferenceVideos(v []string)`

SetReferenceVideos sets ReferenceVideos field to given value.

### HasReferenceVideos

`func (o *VideoGenerationRequest) HasReferenceVideos() bool`

HasReferenceVideos returns a boolean if a field has been set.

### SetReferenceVideosNil

`func (o *VideoGenerationRequest) SetReferenceVideosNil(b bool)`

 SetReferenceVideosNil sets the value for ReferenceVideos to be an explicit nil

### UnsetReferenceVideos
`func (o *VideoGenerationRequest) UnsetReferenceVideos()`

UnsetReferenceVideos ensures that no value is present for ReferenceVideos, not even an explicit nil
### GetResolution

`func (o *VideoGenerationRequest) GetResolution() string`

GetResolution returns the Resolution field if non-nil, zero value otherwise.

### GetResolutionOk

`func (o *VideoGenerationRequest) GetResolutionOk() (*string, bool)`

GetResolutionOk returns a tuple with the Resolution field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolution

`func (o *VideoGenerationRequest) SetResolution(v string)`

SetResolution sets Resolution field to given value.

### HasResolution

`func (o *VideoGenerationRequest) HasResolution() bool`

HasResolution returns a boolean if a field has been set.

### SetResolutionNil

`func (o *VideoGenerationRequest) SetResolutionNil(b bool)`

 SetResolutionNil sets the value for Resolution to be an explicit nil

### UnsetResolution
`func (o *VideoGenerationRequest) UnsetResolution()`

UnsetResolution ensures that no value is present for Resolution, not even an explicit nil
### GetSeconds

`func (o *VideoGenerationRequest) GetSeconds() int32`

GetSeconds returns the Seconds field if non-nil, zero value otherwise.

### GetSecondsOk

`func (o *VideoGenerationRequest) GetSecondsOk() (*int32, bool)`

GetSecondsOk returns a tuple with the Seconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeconds

`func (o *VideoGenerationRequest) SetSeconds(v int32)`

SetSeconds sets Seconds field to given value.

### HasSeconds

`func (o *VideoGenerationRequest) HasSeconds() bool`

HasSeconds returns a boolean if a field has been set.

### SetSecondsNil

`func (o *VideoGenerationRequest) SetSecondsNil(b bool)`

 SetSecondsNil sets the value for Seconds to be an explicit nil

### UnsetSeconds
`func (o *VideoGenerationRequest) UnsetSeconds()`

UnsetSeconds ensures that no value is present for Seconds, not even an explicit nil
### GetSeed

`func (o *VideoGenerationRequest) GetSeed() int32`

GetSeed returns the Seed field if non-nil, zero value otherwise.

### GetSeedOk

`func (o *VideoGenerationRequest) GetSeedOk() (*int32, bool)`

GetSeedOk returns a tuple with the Seed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeed

`func (o *VideoGenerationRequest) SetSeed(v int32)`

SetSeed sets Seed field to given value.

### HasSeed

`func (o *VideoGenerationRequest) HasSeed() bool`

HasSeed returns a boolean if a field has been set.

### SetSeedNil

`func (o *VideoGenerationRequest) SetSeedNil(b bool)`

 SetSeedNil sets the value for Seed to be an explicit nil

### UnsetSeed
`func (o *VideoGenerationRequest) UnsetSeed()`

UnsetSeed ensures that no value is present for Seed, not even an explicit nil
### GetSize

`func (o *VideoGenerationRequest) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *VideoGenerationRequest) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *VideoGenerationRequest) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *VideoGenerationRequest) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSizeNil

`func (o *VideoGenerationRequest) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *VideoGenerationRequest) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil
### GetUser

`func (o *VideoGenerationRequest) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *VideoGenerationRequest) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *VideoGenerationRequest) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *VideoGenerationRequest) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *VideoGenerationRequest) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *VideoGenerationRequest) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


