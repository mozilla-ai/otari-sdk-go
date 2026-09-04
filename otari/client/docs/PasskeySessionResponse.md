# PasskeySessionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveOrganizationId** | **string** | The organization that identity is acting in, which scopes every tenancy surface. | 
**ExpiresAt** | **time.Time** | When the session cookie stops being accepted. | 
**UserId** | **string** | The identity this session speaks for. | 

## Methods

### NewPasskeySessionResponse

`func NewPasskeySessionResponse(activeOrganizationId string, expiresAt time.Time, userId string, ) *PasskeySessionResponse`

NewPasskeySessionResponse instantiates a new PasskeySessionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPasskeySessionResponseWithDefaults

`func NewPasskeySessionResponseWithDefaults() *PasskeySessionResponse`

NewPasskeySessionResponseWithDefaults instantiates a new PasskeySessionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveOrganizationId

`func (o *PasskeySessionResponse) GetActiveOrganizationId() string`

GetActiveOrganizationId returns the ActiveOrganizationId field if non-nil, zero value otherwise.

### GetActiveOrganizationIdOk

`func (o *PasskeySessionResponse) GetActiveOrganizationIdOk() (*string, bool)`

GetActiveOrganizationIdOk returns a tuple with the ActiveOrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveOrganizationId

`func (o *PasskeySessionResponse) SetActiveOrganizationId(v string)`

SetActiveOrganizationId sets ActiveOrganizationId field to given value.


### GetExpiresAt

`func (o *PasskeySessionResponse) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *PasskeySessionResponse) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *PasskeySessionResponse) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.


### GetUserId

`func (o *PasskeySessionResponse) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *PasskeySessionResponse) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *PasskeySessionResponse) SetUserId(v string)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


