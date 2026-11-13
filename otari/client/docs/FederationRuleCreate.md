# FederationRuleCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Audience** | **string** | A value the token&#39;s &#39;aud&#39; must contain. | 
**Email** | Pointer to **NullableString** | The token&#39;s verified &#39;email&#39;, when the rule names one. Compared without regard to case. | [optional] 
**Issuer** | **string** | The token&#39;s &#39;iss&#39;. On an operator&#39;s rule, one of the deployment&#39;s federation_issuers; on a member&#39;s own key, any https issuer, whose discovery document and signing keys are fetched only from a public host. | 
**Subject** | Pointer to **NullableString** | The token&#39;s &#39;sub&#39;, when the rule names one. | [optional] 

## Methods

### NewFederationRuleCreate

`func NewFederationRuleCreate(audience string, issuer string, ) *FederationRuleCreate`

NewFederationRuleCreate instantiates a new FederationRuleCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFederationRuleCreateWithDefaults

`func NewFederationRuleCreateWithDefaults() *FederationRuleCreate`

NewFederationRuleCreateWithDefaults instantiates a new FederationRuleCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudience

`func (o *FederationRuleCreate) GetAudience() string`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *FederationRuleCreate) GetAudienceOk() (*string, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *FederationRuleCreate) SetAudience(v string)`

SetAudience sets Audience field to given value.


### GetEmail

`func (o *FederationRuleCreate) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *FederationRuleCreate) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *FederationRuleCreate) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *FederationRuleCreate) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *FederationRuleCreate) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *FederationRuleCreate) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetIssuer

`func (o *FederationRuleCreate) GetIssuer() string`

GetIssuer returns the Issuer field if non-nil, zero value otherwise.

### GetIssuerOk

`func (o *FederationRuleCreate) GetIssuerOk() (*string, bool)`

GetIssuerOk returns a tuple with the Issuer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuer

`func (o *FederationRuleCreate) SetIssuer(v string)`

SetIssuer sets Issuer field to given value.


### GetSubject

`func (o *FederationRuleCreate) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *FederationRuleCreate) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *FederationRuleCreate) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *FederationRuleCreate) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### SetSubjectNil

`func (o *FederationRuleCreate) SetSubjectNil(b bool)`

 SetSubjectNil sets the value for Subject to be an explicit nil

### UnsetSubject
`func (o *FederationRuleCreate) UnsetSubject()`

UnsetSubject ensures that no value is present for Subject, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


