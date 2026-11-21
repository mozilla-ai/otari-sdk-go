# OrganizationSsoPreflightMemberPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **NullableString** | The member&#39;s sign-in address, or null for an identity with none. | 
**FullName** | **NullableString** | The member&#39;s display name, or null when none is set. | 
**UserId** | **string** |  | 

## Methods

### NewOrganizationSsoPreflightMemberPublic

`func NewOrganizationSsoPreflightMemberPublic(email NullableString, fullName NullableString, userId string, ) *OrganizationSsoPreflightMemberPublic`

NewOrganizationSsoPreflightMemberPublic instantiates a new OrganizationSsoPreflightMemberPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationSsoPreflightMemberPublicWithDefaults

`func NewOrganizationSsoPreflightMemberPublicWithDefaults() *OrganizationSsoPreflightMemberPublic`

NewOrganizationSsoPreflightMemberPublicWithDefaults instantiates a new OrganizationSsoPreflightMemberPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *OrganizationSsoPreflightMemberPublic) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *OrganizationSsoPreflightMemberPublic) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *OrganizationSsoPreflightMemberPublic) SetEmail(v string)`

SetEmail sets Email field to given value.


### SetEmailNil

`func (o *OrganizationSsoPreflightMemberPublic) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *OrganizationSsoPreflightMemberPublic) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetFullName

`func (o *OrganizationSsoPreflightMemberPublic) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *OrganizationSsoPreflightMemberPublic) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *OrganizationSsoPreflightMemberPublic) SetFullName(v string)`

SetFullName sets FullName field to given value.


### SetFullNameNil

`func (o *OrganizationSsoPreflightMemberPublic) SetFullNameNil(b bool)`

 SetFullNameNil sets the value for FullName to be an explicit nil

### UnsetFullName
`func (o *OrganizationSsoPreflightMemberPublic) UnsetFullName()`

UnsetFullName ensures that no value is present for FullName, not even an explicit nil
### GetUserId

`func (o *OrganizationSsoPreflightMemberPublic) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *OrganizationSsoPreflightMemberPublic) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *OrganizationSsoPreflightMemberPublic) SetUserId(v string)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


