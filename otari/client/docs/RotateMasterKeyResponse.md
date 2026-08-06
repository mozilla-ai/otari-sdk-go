# RotateMasterKeyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MasterKey** | **string** | The new plaintext master key. Store it now; it is never returned again. | 

## Methods

### NewRotateMasterKeyResponse

`func NewRotateMasterKeyResponse(masterKey string, ) *RotateMasterKeyResponse`

NewRotateMasterKeyResponse instantiates a new RotateMasterKeyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRotateMasterKeyResponseWithDefaults

`func NewRotateMasterKeyResponseWithDefaults() *RotateMasterKeyResponse`

NewRotateMasterKeyResponseWithDefaults instantiates a new RotateMasterKeyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMasterKey

`func (o *RotateMasterKeyResponse) GetMasterKey() string`

GetMasterKey returns the MasterKey field if non-nil, zero value otherwise.

### GetMasterKeyOk

`func (o *RotateMasterKeyResponse) GetMasterKeyOk() (*string, bool)`

GetMasterKeyOk returns a tuple with the MasterKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterKey

`func (o *RotateMasterKeyResponse) SetMasterKey(v string)`

SetMasterKey sets MasterKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


