# CountryDatabaseStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientCountryDatabaseSet** | **bool** | Whether &#x60;client_country_database&#x60; names a file. That file is read, and a stored one is not. | 
**Stored** | [**NullableStoredCountryDatabasePublic**](StoredCountryDatabasePublic.md) | The stored file, or null when none is stored. | 

## Methods

### NewCountryDatabaseStatus

`func NewCountryDatabaseStatus(clientCountryDatabaseSet bool, stored NullableStoredCountryDatabasePublic, ) *CountryDatabaseStatus`

NewCountryDatabaseStatus instantiates a new CountryDatabaseStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCountryDatabaseStatusWithDefaults

`func NewCountryDatabaseStatusWithDefaults() *CountryDatabaseStatus`

NewCountryDatabaseStatusWithDefaults instantiates a new CountryDatabaseStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientCountryDatabaseSet

`func (o *CountryDatabaseStatus) GetClientCountryDatabaseSet() bool`

GetClientCountryDatabaseSet returns the ClientCountryDatabaseSet field if non-nil, zero value otherwise.

### GetClientCountryDatabaseSetOk

`func (o *CountryDatabaseStatus) GetClientCountryDatabaseSetOk() (*bool, bool)`

GetClientCountryDatabaseSetOk returns a tuple with the ClientCountryDatabaseSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCountryDatabaseSet

`func (o *CountryDatabaseStatus) SetClientCountryDatabaseSet(v bool)`

SetClientCountryDatabaseSet sets ClientCountryDatabaseSet field to given value.


### GetStored

`func (o *CountryDatabaseStatus) GetStored() StoredCountryDatabasePublic`

GetStored returns the Stored field if non-nil, zero value otherwise.

### GetStoredOk

`func (o *CountryDatabaseStatus) GetStoredOk() (*StoredCountryDatabasePublic, bool)`

GetStoredOk returns a tuple with the Stored field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStored

`func (o *CountryDatabaseStatus) SetStored(v StoredCountryDatabasePublic)`

SetStored sets Stored field to given value.


### SetStoredNil

`func (o *CountryDatabaseStatus) SetStoredNil(b bool)`

 SetStoredNil sets the value for Stored to be an explicit nil

### UnsetStored
`func (o *CountryDatabaseStatus) UnsetStored()`

UnsetStored ensures that no value is present for Stored, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


