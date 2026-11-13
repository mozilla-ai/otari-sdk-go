# AppDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Issuer** | **string** |  | 
**Name** | **string** |  | 
**RedirectUri** | **string** |  | 

## Methods

### NewAppDetails

`func NewAppDetails(issuer string, name string, redirectUri string, ) *AppDetails`

NewAppDetails instantiates a new AppDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppDetailsWithDefaults

`func NewAppDetailsWithDefaults() *AppDetails`

NewAppDetailsWithDefaults instantiates a new AppDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIssuer

`func (o *AppDetails) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *AppDetails) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *AppDetails) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.


### GetName

`func (o *AppDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AppDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AppDetails) SetName(v string)`

SetName sets Name field to given value.


### GetRedirectUri

`func (o *AppDetails) GetRedirectUri() string`

GetRedirectUri returns the RedirectUri field if non-nil, zero value otherwise.

### GetRedirectUriOk

`func (o *AppDetails) GetRedirectUriOk() (*string, bool)`

GetRedirectUriOk returns a tuple with the RedirectUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUri

`func (o *AppDetails) SetRedirectUri(v string)`

SetRedirectUri sets RedirectUri field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


