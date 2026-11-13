# SsoLookupRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | The address to resolve an organization from. | 

## Methods

### NewSsoLookupRequest

`func NewSsoLookupRequest(email string, ) *SsoLookupRequest`

NewSsoLookupRequest instantiates a new SsoLookupRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSsoLookupRequestWithDefaults

`func NewSsoLookupRequestWithDefaults() *SsoLookupRequest`

NewSsoLookupRequestWithDefaults instantiates a new SsoLookupRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *SsoLookupRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *SsoLookupRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *SsoLookupRequest) SetEmail(v string)`

SetEmail sets Email field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


