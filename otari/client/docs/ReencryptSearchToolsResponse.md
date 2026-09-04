# ReencryptSearchToolsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reencrypted** | **int32** | Number of stored search-tool keys re-encrypted. | 
**Unreadable** | **int32** | Number of encrypted keys left untouched because they could not be decrypted. | 

## Methods

### NewReencryptSearchToolsResponse

`func NewReencryptSearchToolsResponse(reencrypted int32, unreadable int32, ) *ReencryptSearchToolsResponse`

NewReencryptSearchToolsResponse instantiates a new ReencryptSearchToolsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReencryptSearchToolsResponseWithDefaults

`func NewReencryptSearchToolsResponseWithDefaults() *ReencryptSearchToolsResponse`

NewReencryptSearchToolsResponseWithDefaults instantiates a new ReencryptSearchToolsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReencrypted

`func (o *ReencryptSearchToolsResponse) GetReencrypted() int32`

GetReencrypted returns the Reencrypted field if non-nil, zero value otherwise.

### GetReencryptedOk

`func (o *ReencryptSearchToolsResponse) GetReencryptedOk() (*int32, bool)`

GetReencryptedOk returns a tuple with the Reencrypted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReencrypted

`func (o *ReencryptSearchToolsResponse) SetReencrypted(v int32)`

SetReencrypted sets Reencrypted field to given value.


### GetUnreadable

`func (o *ReencryptSearchToolsResponse) GetUnreadable() int32`

GetUnreadable returns the Unreadable field if non-nil, zero value otherwise.

### GetUnreadableOk

`func (o *ReencryptSearchToolsResponse) GetUnreadableOk() (*int32, bool)`

GetUnreadableOk returns a tuple with the Unreadable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnreadable

`func (o *ReencryptSearchToolsResponse) SetUnreadable(v int32)`

SetUnreadable sets Unreadable field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


