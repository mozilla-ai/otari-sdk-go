# HostedCatalogUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cleanup** | Pointer to **bool** | Also remove the prices of models no hosted provider serves, with their history. | [optional] [default to false]
**Digest** | Pointer to **NullableString** | The &#x60;digest&#x60; of the previewed prices, to apply them; 409 when the pending prices differ. Without it, no price changes. | [optional] 

## Methods

### NewHostedCatalogUpdateRequest

`func NewHostedCatalogUpdateRequest() *HostedCatalogUpdateRequest`

NewHostedCatalogUpdateRequest instantiates a new HostedCatalogUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostedCatalogUpdateRequestWithDefaults

`func NewHostedCatalogUpdateRequestWithDefaults() *HostedCatalogUpdateRequest`

NewHostedCatalogUpdateRequestWithDefaults instantiates a new HostedCatalogUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCleanup

`func (o *HostedCatalogUpdateRequest) GetCleanup() bool`

GetCleanup returns the Cleanup field if non-nil, zero value otherwise.

### GetCleanupOk

`func (o *HostedCatalogUpdateRequest) GetCleanupOk() (*bool, bool)`

GetCleanupOk returns a tuple with the Cleanup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCleanup

`func (o *HostedCatalogUpdateRequest) SetCleanup(v bool)`

SetCleanup sets Cleanup field to given value.

### HasCleanup

`func (o *HostedCatalogUpdateRequest) HasCleanup() bool`

HasCleanup returns a boolean if a field has been set.

### GetDigest

`func (o *HostedCatalogUpdateRequest) GetDigest() string`

GetDigest returns the Digest field if non-nil, zero value otherwise.

### GetDigestOk

`func (o *HostedCatalogUpdateRequest) GetDigestOk() (*string, bool)`

GetDigestOk returns a tuple with the Digest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDigest

`func (o *HostedCatalogUpdateRequest) SetDigest(v string)`

SetDigest sets Digest field to given value.

### HasDigest

`func (o *HostedCatalogUpdateRequest) HasDigest() bool`

HasDigest returns a boolean if a field has been set.

### SetDigestNil

`func (o *HostedCatalogUpdateRequest) SetDigestNil(b bool)`

 SetDigestNil sets the value for Digest to be an explicit nil

### UnsetDigest
`func (o *HostedCatalogUpdateRequest) UnsetDigest()`

UnsetDigest ensures that no value is present for Digest, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


