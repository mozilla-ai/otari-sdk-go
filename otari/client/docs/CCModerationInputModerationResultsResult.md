# CCModerationInputModerationResultsResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Categories** | **map[string]bool** |  | 
**CategoryAppliedInputTypes** | **map[string][]string** |  | 
**CategoryScores** | **map[string]float32** |  | 
**Flagged** | **bool** |  | 
**Model** | **string** |  | 
**Type** | **string** |  | 

## Methods

### NewCCModerationInputModerationResultsResult

`func NewCCModerationInputModerationResultsResult(categories map[string]bool, categoryAppliedInputTypes map[string][]string, categoryScores map[string]float32, flagged bool, model string, type_ string, ) *CCModerationInputModerationResultsResult`

NewCCModerationInputModerationResultsResult instantiates a new CCModerationInputModerationResultsResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCCModerationInputModerationResultsResultWithDefaults

`func NewCCModerationInputModerationResultsResultWithDefaults() *CCModerationInputModerationResultsResult`

NewCCModerationInputModerationResultsResultWithDefaults instantiates a new CCModerationInputModerationResultsResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategories

`func (o *CCModerationInputModerationResultsResult) GetCategories() map[string]bool`

GetCategories returns the Categories field if non-nil, zero value otherwise.

### GetCategoriesOk

`func (o *CCModerationInputModerationResultsResult) GetCategoriesOk() (*map[string]bool, bool)`

GetCategoriesOk returns a tuple with the Categories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategories

`func (o *CCModerationInputModerationResultsResult) SetCategories(v map[string]bool)`

SetCategories sets Categories field to given value.


### GetCategoryAppliedInputTypes

`func (o *CCModerationInputModerationResultsResult) GetCategoryAppliedInputTypes() map[string][]string`

GetCategoryAppliedInputTypes returns the CategoryAppliedInputTypes field if non-nil, zero value otherwise.

### GetCategoryAppliedInputTypesOk

`func (o *CCModerationInputModerationResultsResult) GetCategoryAppliedInputTypesOk() (*map[string][]string, bool)`

GetCategoryAppliedInputTypesOk returns a tuple with the CategoryAppliedInputTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryAppliedInputTypes

`func (o *CCModerationInputModerationResultsResult) SetCategoryAppliedInputTypes(v map[string][]string)`

SetCategoryAppliedInputTypes sets CategoryAppliedInputTypes field to given value.


### GetCategoryScores

`func (o *CCModerationInputModerationResultsResult) GetCategoryScores() map[string]float32`

GetCategoryScores returns the CategoryScores field if non-nil, zero value otherwise.

### GetCategoryScoresOk

`func (o *CCModerationInputModerationResultsResult) GetCategoryScoresOk() (*map[string]float32, bool)`

GetCategoryScoresOk returns a tuple with the CategoryScores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryScores

`func (o *CCModerationInputModerationResultsResult) SetCategoryScores(v map[string]float32)`

SetCategoryScores sets CategoryScores field to given value.


### GetFlagged

`func (o *CCModerationInputModerationResultsResult) GetFlagged() bool`

GetFlagged returns the Flagged field if non-nil, zero value otherwise.

### GetFlaggedOk

`func (o *CCModerationInputModerationResultsResult) GetFlaggedOk() (*bool, bool)`

GetFlaggedOk returns a tuple with the Flagged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlagged

`func (o *CCModerationInputModerationResultsResult) SetFlagged(v bool)`

SetFlagged sets Flagged field to given value.


### GetModel

`func (o *CCModerationInputModerationResultsResult) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *CCModerationInputModerationResultsResult) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *CCModerationInputModerationResultsResult) SetModel(v string)`

SetModel sets Model field to given value.


### GetType

`func (o *CCModerationInputModerationResultsResult) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CCModerationInputModerationResultsResult) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CCModerationInputModerationResultsResult) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


