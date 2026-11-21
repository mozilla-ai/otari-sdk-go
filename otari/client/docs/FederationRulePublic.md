# FederationRulePublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKeyId** | **string** |  | 
**Audience** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**Email** | **NullableString** |  | 
**Id** | **string** |  | 
**Issuer** | **string** |  | 
**MemberManaged** | **bool** | Whether a member gave the rule to their own key: it then matches only in the token exchange. | 
**Subject** | **NullableString** |  | 

## Methods

### NewFederationRulePublic

`func NewFederationRulePublic(apiKeyId string, audience string, createdAt time.Time, email NullableString, id string, issuer string, memberManaged bool, subject NullableString, ) *FederationRulePublic`

NewFederationRulePublic instantiates a new FederationRulePublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFederationRulePublicWithDefaults

`func NewFederationRulePublicWithDefaults() *FederationRulePublic`

NewFederationRulePublicWithDefaults instantiates a new FederationRulePublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKeyId

`func (o *FederationRulePublic) GetApiKeyId() string`

GetApiKeyId returns the ApiKeyId field if non-nil, zero value otherwise.

### GetApiKeyIdOk

`func (o *FederationRulePublic) GetApiKeyIdOk() (*string, bool)`

GetApiKeyIdOk returns a tuple with the ApiKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKeyId

`func (o *FederationRulePublic) SetApiKeyId(v string)`

SetApiKeyId sets ApiKeyId field to given value.


### GetAudience

`func (o *FederationRulePublic) GetAudience() string`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *FederationRulePublic) GetAudienceOk() (*string, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *FederationRulePublic) SetAudience(v string)`

SetAudience sets Audience field to given value.


### GetCreatedAt

`func (o *FederationRulePublic) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FederationRulePublic) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FederationRulePublic) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetEmail

`func (o *FederationRulePublic) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *FederationRulePublic) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *FederationRulePublic) SetEmail(v string)`

SetEmail sets Email field to given value.


### SetEmailNil

`func (o *FederationRulePublic) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *FederationRulePublic) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetId

`func (o *FederationRulePublic) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FederationRulePublic) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FederationRulePublic) SetId(v string)`

SetId sets Id field to given value.


### GetIssuer

`func (o *FederationRulePublic) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *FederationRulePublic) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *FederationRulePublic) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.


### GetMemberManaged

`func (o *FederationRulePublic) GetMemberManaged() bool`

GetMemberManaged returns the MemberManaged field if non-nil, zero value otherwise.

### GetMemberManagedOk

`func (o *FederationRulePublic) GetMemberManagedOk() (*bool, bool)`

GetMemberManagedOk returns a tuple with the MemberManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberManaged

`func (o *FederationRulePublic) SetMemberManaged(v bool)`

SetMemberManaged sets MemberManaged field to given value.


### GetSubject

`func (o *FederationRulePublic) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *FederationRulePublic) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *FederationRulePublic) SetSubject(v string)`

SetSubject sets Subject field to given value.


### SetSubjectNil

`func (o *FederationRulePublic) SetSubjectNil(b bool)`

 SetSubjectNil sets the value for Subject to be an explicit nil

### UnsetSubject
`func (o *FederationRulePublic) UnsetSubject()`

UnsetSubject ensures that no value is present for Subject, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


