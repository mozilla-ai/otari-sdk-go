# MRBetaUserLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**City** | Pointer to **NullableString** | Filter to a single event type or metric name (e.g. &#39;tool_result&#39;, &#39;claude_code.commit.count&#39;) | [optional] 
**Country** | Pointer to **NullableString** | Filter to a single event type or metric name (e.g. &#39;tool_result&#39;, &#39;claude_code.commit.count&#39;) | [optional] 
**Region** | Pointer to **NullableString** | Filter to a single event type or metric name (e.g. &#39;tool_result&#39;, &#39;claude_code.commit.count&#39;) | [optional] 
**Timezone** | Pointer to **NullableString** | Filter to a single event type or metric name (e.g. &#39;tool_result&#39;, &#39;claude_code.commit.count&#39;) | [optional] 

## Methods

### NewMRBetaUserLocation

`func NewMRBetaUserLocation(type_ string, ) *MRBetaUserLocation`

NewMRBetaUserLocation instantiates a new MRBetaUserLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMRBetaUserLocationWithDefaults

`func NewMRBetaUserLocationWithDefaults() *MRBetaUserLocation`

NewMRBetaUserLocationWithDefaults instantiates a new MRBetaUserLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *MRBetaUserLocation) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MRBetaUserLocation) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MRBetaUserLocation) SetType(v string)`

SetType sets Type field to given value.


### GetCity

`func (o *MRBetaUserLocation) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *MRBetaUserLocation) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *MRBetaUserLocation) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *MRBetaUserLocation) HasCity() bool`

HasCity returns a boolean if a field has been set.

### SetCityNil

`func (o *MRBetaUserLocation) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *MRBetaUserLocation) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetCountry

`func (o *MRBetaUserLocation) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *MRBetaUserLocation) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *MRBetaUserLocation) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *MRBetaUserLocation) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *MRBetaUserLocation) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *MRBetaUserLocation) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetRegion

`func (o *MRBetaUserLocation) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *MRBetaUserLocation) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *MRBetaUserLocation) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *MRBetaUserLocation) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *MRBetaUserLocation) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *MRBetaUserLocation) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetTimezone

`func (o *MRBetaUserLocation) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *MRBetaUserLocation) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *MRBetaUserLocation) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *MRBetaUserLocation) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### SetTimezoneNil

`func (o *MRBetaUserLocation) SetTimezoneNil(b bool)`

 SetTimezoneNil sets the value for Timezone to be an explicit nil

### UnsetTimezone
`func (o *MRBetaUserLocation) UnsetTimezone()`

UnsetTimezone ensures that no value is present for Timezone, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


