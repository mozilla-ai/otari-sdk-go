# CreateSessionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MasterKey** | **string** | The gateway master key; verified once and never stored by the browser. | 

## Methods

### NewCreateSessionRequest

`func NewCreateSessionRequest(masterKey string, ) *CreateSessionRequest`

NewCreateSessionRequest instantiates a new CreateSessionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSessionRequestWithDefaults

`func NewCreateSessionRequestWithDefaults() *CreateSessionRequest`

NewCreateSessionRequestWithDefaults instantiates a new CreateSessionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMasterKey

`func (o *CreateSessionRequest) GetMasterKey() string`

GetMasterKey returns the MasterKey field if non-nil, zero value otherwise.

### GetMasterKeyOk

`func (o *CreateSessionRequest) GetMasterKeyOk() (*string, bool)`

GetMasterKeyOk returns a tuple with the MasterKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterKey

`func (o *CreateSessionRequest) SetMasterKey(v string)`

SetMasterKey sets MasterKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


