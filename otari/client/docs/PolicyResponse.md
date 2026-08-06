# PolicyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **NullableString** |  | [optional] 
**IsDynamic** | Pointer to **bool** |  | [optional] [default to false]
**Name** | **string** |  | 
**Source** | **string** |  | 
**Spec** | **map[string]interface{}** |  | 
**UpdatedAt** | Pointer to **NullableString** |  | [optional] 
**UserId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewPolicyResponse

`func NewPolicyResponse(name string, source string, spec map[string]interface{}, ) *PolicyResponse`

NewPolicyResponse instantiates a new PolicyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyResponseWithDefaults

`func NewPolicyResponseWithDefaults() *PolicyResponse`

NewPolicyResponseWithDefaults instantiates a new PolicyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *PolicyResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PolicyResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PolicyResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PolicyResponse) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *PolicyResponse) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *PolicyResponse) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil
### GetIsDynamic

`func (o *PolicyResponse) GetIsDynamic() bool`

GetIsDynamic returns the IsDynamic field if non-nil, zero value otherwise.

### GetIsDynamicOk

`func (o *PolicyResponse) GetIsDynamicOk() (*bool, bool)`

GetIsDynamicOk returns a tuple with the IsDynamic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDynamic

`func (o *PolicyResponse) SetIsDynamic(v bool)`

SetIsDynamic sets IsDynamic field to given value.

### HasIsDynamic

`func (o *PolicyResponse) HasIsDynamic() bool`

HasIsDynamic returns a boolean if a field has been set.

### GetName

`func (o *PolicyResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PolicyResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PolicyResponse) SetName(v string)`

SetName sets Name field to given value.


### GetSource

`func (o *PolicyResponse) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *PolicyResponse) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *PolicyResponse) SetSource(v string)`

SetSource sets Source field to given value.


### GetSpec

`func (o *PolicyResponse) GetSpec() map[string]interface{}`

GetSpec returns the Spec field if non-nil, zero value otherwise.

### GetSpecOk

`func (o *PolicyResponse) GetSpecOk() (*map[string]interface{}, bool)`

GetSpecOk returns a tuple with the Spec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpec

`func (o *PolicyResponse) SetSpec(v map[string]interface{})`

SetSpec sets Spec field to given value.


### GetUpdatedAt

`func (o *PolicyResponse) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PolicyResponse) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PolicyResponse) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PolicyResponse) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *PolicyResponse) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *PolicyResponse) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil
### GetUserId

`func (o *PolicyResponse) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *PolicyResponse) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *PolicyResponse) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *PolicyResponse) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *PolicyResponse) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *PolicyResponse) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


