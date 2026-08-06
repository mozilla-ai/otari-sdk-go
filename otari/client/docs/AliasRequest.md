# AliasRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Display name callers use as the model, e.g. &#39;fast-model&#39;. | 
**Target** | **string** | Selector the alias resolves to, as &#39;provider:model&#39; or &#39;instance:model&#39;. | 
**UserId** | Pointer to **NullableString** | User this alias belongs to. Omit for a global alias every caller sees. A user-scoped alias resolves only for that user and shadows a global one of the same name. | [optional] 

## Methods

### NewAliasRequest

`func NewAliasRequest(name string, target string, ) *AliasRequest`

NewAliasRequest instantiates a new AliasRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAliasRequestWithDefaults

`func NewAliasRequestWithDefaults() *AliasRequest`

NewAliasRequestWithDefaults instantiates a new AliasRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AliasRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AliasRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AliasRequest) SetName(v string)`

SetName sets Name field to given value.


### GetTarget

`func (o *AliasRequest) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *AliasRequest) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *AliasRequest) SetTarget(v string)`

SetTarget sets Target field to given value.


### GetUserId

`func (o *AliasRequest) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *AliasRequest) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *AliasRequest) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *AliasRequest) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *AliasRequest) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *AliasRequest) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


