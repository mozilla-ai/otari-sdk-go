# SsoLookupResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Found** | **bool** | Whether an organization&#39;s own connection admits a sign-in at this address. | 
**OrganizationName** | Pointer to **NullableString** | The organization&#39;s own name, set only when found. | [optional] 
**Slug** | Pointer to **NullableString** | The organization&#39;s slug, to start the sign-in with, set only when found. | [optional] 

## Methods

### NewSsoLookupResponse

`func NewSsoLookupResponse(found bool, ) *SsoLookupResponse`

NewSsoLookupResponse instantiates a new SsoLookupResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSsoLookupResponseWithDefaults

`func NewSsoLookupResponseWithDefaults() *SsoLookupResponse`

NewSsoLookupResponseWithDefaults instantiates a new SsoLookupResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFound

`func (o *SsoLookupResponse) GetFound() bool`

GetFound returns the Found field if non-nil, zero value otherwise.

### GetFoundOk

`func (o *SsoLookupResponse) GetFoundOk() (*bool, bool)`

GetFoundOk returns a tuple with the Found field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFound

`func (o *SsoLookupResponse) SetFound(v bool)`

SetFound sets Found field to given value.


### GetOrganizationName

`func (o *SsoLookupResponse) GetOrganizationName() string`

GetOrganizationName returns the OrganizationName field if non-nil, zero value otherwise.

### GetOrganizationNameOk

`func (o *SsoLookupResponse) GetOrganizationNameOk() (*string, bool)`

GetOrganizationNameOk returns a tuple with the OrganizationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationName

`func (o *SsoLookupResponse) SetOrganizationName(v string)`

SetOrganizationName sets OrganizationName field to given value.

### HasOrganizationName

`func (o *SsoLookupResponse) HasOrganizationName() bool`

HasOrganizationName returns a boolean if a field has been set.

### SetOrganizationNameNil

`func (o *SsoLookupResponse) SetOrganizationNameNil(b bool)`

 SetOrganizationNameNil sets the value for OrganizationName to be an explicit nil

### UnsetOrganizationName
`func (o *SsoLookupResponse) UnsetOrganizationName()`

UnsetOrganizationName ensures that no value is present for OrganizationName, not even an explicit nil
### GetSlug

`func (o *SsoLookupResponse) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *SsoLookupResponse) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *SsoLookupResponse) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *SsoLookupResponse) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### SetSlugNil

`func (o *SsoLookupResponse) SetSlugNil(b bool)`

 SetSlugNil sets the value for Slug to be an explicit nil

### UnsetSlug
`func (o *SsoLookupResponse) UnsetSlug()`

UnsetSlug ensures that no value is present for Slug, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


