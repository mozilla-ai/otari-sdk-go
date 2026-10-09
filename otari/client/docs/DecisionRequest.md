# DecisionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Images** | Pointer to **[]string** | Images for a vision decision model, as data URLs (data:image/...;base64,...). An extension llama-server supports; a provider that does not refuses the request. | [optional] 
**Model** | **string** | Decision provider and model, e.g. &#39;typesafe:jev-latest&#39; | 
**Questions** | [**map[string]QuestionsValue**](QuestionsValue.md) | Questions, keyed by answer name | 
**State** | [**State**](State.md) |  | 
**User** | Pointer to **NullableString** | User ID for usage attribution; not sent upstream | [optional] 

## Methods

### NewDecisionRequest

`func NewDecisionRequest(model string, questions map[string]QuestionsValue, state State, ) *DecisionRequest`

NewDecisionRequest instantiates a new DecisionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDecisionRequestWithDefaults

`func NewDecisionRequestWithDefaults() *DecisionRequest`

NewDecisionRequestWithDefaults instantiates a new DecisionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImages

`func (o *DecisionRequest) GetImages() []string`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *DecisionRequest) GetImagesOk() (*[]string, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *DecisionRequest) SetImages(v []string)`

SetImages sets Images field to given value.

### HasImages

`func (o *DecisionRequest) HasImages() bool`

HasImages returns a boolean if a field has been set.

### SetImagesNil

`func (o *DecisionRequest) SetImagesNil(b bool)`

 SetImagesNil sets the value for Images to be an explicit nil

### UnsetImages
`func (o *DecisionRequest) UnsetImages()`

UnsetImages ensures that no value is present for Images, not even an explicit nil
### GetModel

`func (o *DecisionRequest) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *DecisionRequest) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *DecisionRequest) SetModel(v string)`

SetModel sets Model field to given value.


### GetQuestions

`func (o *DecisionRequest) GetQuestions() map[string]QuestionsValue`

GetQuestions returns the Questions field if non-nil, zero value otherwise.

### GetQuestionsOk

`func (o *DecisionRequest) GetQuestionsOk() (*map[string]QuestionsValue, bool)`

GetQuestionsOk returns a tuple with the Questions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestions

`func (o *DecisionRequest) SetQuestions(v map[string]QuestionsValue)`

SetQuestions sets Questions field to given value.


### GetState

`func (o *DecisionRequest) GetState() State`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *DecisionRequest) GetStateOk() (*State, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *DecisionRequest) SetState(v State)`

SetState sets State field to given value.


### GetUser

`func (o *DecisionRequest) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *DecisionRequest) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *DecisionRequest) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *DecisionRequest) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *DecisionRequest) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *DecisionRequest) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


