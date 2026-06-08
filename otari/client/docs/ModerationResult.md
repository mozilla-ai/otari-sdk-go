# ModerationResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Categories** | Pointer to **map[string]bool** |  | [optional] 
**CategoryAppliedInputTypes** | Pointer to **map[string][]string** |  | [optional] 
**CategoryScores** | Pointer to **map[string]float32** |  | [optional] 
**Flagged** | **bool** |  | 
**ProviderRaw** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewModerationResult

`func NewModerationResult(flagged bool, ) *ModerationResult`

NewModerationResult instantiates a new ModerationResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModerationResultWithDefaults

`func NewModerationResultWithDefaults() *ModerationResult`

NewModerationResultWithDefaults instantiates a new ModerationResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategories

`func (o *ModerationResult) GetCategories() map[string]bool`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *ModerationResult) GetCategoriesOk() (*map[string]bool, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *ModerationResult) SetCategories(v map[string]bool)`

SetCategories sets Categories field to given value.

### HasCategories

`func (o *ModerationResult) HasCategories() bool`

HasCategories returns a boolean if a field has been set.

### GetCategoryAppliedInputTypes

`func (o *ModerationResult) GetCategoryAppliedInputTypes() map[string][]string`

GetCategoryAppliedInputTypes returns the CategoryAppliedInputTypes field if non-nil, zero value otherwise.

### GetCategoryAppliedInputTypesOk

`func (o *ModerationResult) GetCategoryAppliedInputTypesOk() (*map[string][]string, bool)`

GetCategoryAppliedInputTypesOk returns a tuple with the CategoryAppliedInputTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryAppliedInputTypes

`func (o *ModerationResult) SetCategoryAppliedInputTypes(v map[string][]string)`

SetCategoryAppliedInputTypes sets CategoryAppliedInputTypes field to given value.

### HasCategoryAppliedInputTypes

`func (o *ModerationResult) HasCategoryAppliedInputTypes() bool`

HasCategoryAppliedInputTypes returns a boolean if a field has been set.

### SetCategoryAppliedInputTypesNil

`func (o *ModerationResult) SetCategoryAppliedInputTypesNil(b bool)`

 SetCategoryAppliedInputTypesNil sets the value for CategoryAppliedInputTypes to be an explicit nil

### UnsetCategoryAppliedInputTypes
`func (o *ModerationResult) UnsetCategoryAppliedInputTypes()`

UnsetCategoryAppliedInputTypes ensures that no value is present for CategoryAppliedInputTypes, not even an explicit nil
### GetCategoryScores

`func (o *ModerationResult) GetCategoryScores() map[string]float32`

GetCategoryScores returns the CategoryScores field if non-nil, zero value otherwise.

### GetCategoryScoresOk

`func (o *ModerationResult) GetCategoryScoresOk() (*map[string]float32, bool)`

GetCategoryScoresOk returns a tuple with the CategoryScores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryScores

`func (o *ModerationResult) SetCategoryScores(v map[string]float32)`

SetCategoryScores sets CategoryScores field to given value.

### HasCategoryScores

`func (o *ModerationResult) HasCategoryScores() bool`

HasCategoryScores returns a boolean if a field has been set.

### GetFlagged

`func (o *ModerationResult) GetFlagged() bool`

GetFlagged returns the Flagged field if non-nil, zero value otherwise.

### GetFlaggedOk

`func (o *ModerationResult) GetFlaggedOk() (*bool, bool)`

GetFlaggedOk returns a tuple with the Flagged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagged

`func (o *ModerationResult) SetFlagged(v bool)`

SetFlagged sets Flagged field to given value.


### GetProviderRaw

`func (o *ModerationResult) GetProviderRaw() map[string]interface{}`

GetProviderRaw returns the ProviderRaw field if non-nil, zero value otherwise.

### GetProviderRawOk

`func (o *ModerationResult) GetProviderRawOk() (*map[string]interface{}, bool)`

GetProviderRawOk returns a tuple with the ProviderRaw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderRaw

`func (o *ModerationResult) SetProviderRaw(v map[string]interface{})`

SetProviderRaw sets ProviderRaw field to given value.

### HasProviderRaw

`func (o *ModerationResult) HasProviderRaw() bool`

HasProviderRaw returns a boolean if a field has been set.

### SetProviderRawNil

`func (o *ModerationResult) SetProviderRawNil(b bool)`

 SetProviderRawNil sets the value for ProviderRaw to be an explicit nil

### UnsetProviderRaw
`func (o *ModerationResult) UnsetProviderRaw()`

UnsetProviderRaw ensures that no value is present for ProviderRaw, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


