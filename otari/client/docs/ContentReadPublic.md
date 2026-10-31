# ContentReadPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReadAt** | **time.Time** |  | 
**ReaderEmail** | **string** | Their email address when they read it. | 
**ReaderUserId** | **NullableString** | Who read it; null once that account is gone. | 

## Methods

### NewContentReadPublic

`func NewContentReadPublic(readAt time.Time, readerEmail string, readerUserId NullableString, ) *ContentReadPublic`

NewContentReadPublic instantiates a new ContentReadPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentReadPublicWithDefaults

`func NewContentReadPublicWithDefaults() *ContentReadPublic`

NewContentReadPublicWithDefaults instantiates a new ContentReadPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReadAt

`func (o *ContentReadPublic) GetReadAt() time.Time`

GetReadAt returns the ReadAt field if non-nil, zero value otherwise.

### GetReadAtOk

`func (o *ContentReadPublic) GetReadAtOk() (*time.Time, bool)`

GetReadAtOk returns a tuple with the ReadAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadAt

`func (o *ContentReadPublic) SetReadAt(v time.Time)`

SetReadAt sets ReadAt field to given value.


### GetReaderEmail

`func (o *ContentReadPublic) GetReaderEmail() string`

GetReaderEmail returns the ReaderEmail field if non-nil, zero value otherwise.

### GetReaderEmailOk

`func (o *ContentReadPublic) GetReaderEmailOk() (*string, bool)`

GetReaderEmailOk returns a tuple with the ReaderEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReaderEmail

`func (o *ContentReadPublic) SetReaderEmail(v string)`

SetReaderEmail sets ReaderEmail field to given value.


### GetReaderUserId

`func (o *ContentReadPublic) GetReaderUserId() string`

GetReaderUserId returns the ReaderUserId field if non-nil, zero value otherwise.

### GetReaderUserIdOk

`func (o *ContentReadPublic) GetReaderUserIdOk() (*string, bool)`

GetReaderUserIdOk returns a tuple with the ReaderUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReaderUserId

`func (o *ContentReadPublic) SetReaderUserId(v string)`

SetReaderUserId sets ReaderUserId field to given value.


### SetReaderUserIdNil

`func (o *ContentReadPublic) SetReaderUserIdNil(b bool)`

 SetReaderUserIdNil sets the value for ReaderUserId to be an explicit nil

### UnsetReaderUserId
`func (o *ContentReadPublic) UnsetReaderUserId()`

UnsetReaderUserId ensures that no value is present for ReaderUserId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


