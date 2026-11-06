# PasswordLinkPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpiresAt** | **time.Time** | When the link stops working. | 
**Link** | Pointer to **NullableString** | The link to pass on to the account&#39;s owner yourself, set only when mail_sent is false. Choosing a password through it records the address as vouched for rather than proven, since the mailbox never saw the link. | [optional] 
**MailSent** | **bool** | Whether the link was mailed to the account&#39;s address. False when this deployment cannot mail links, or the send failed. | 

## Methods

### NewPasswordLinkPublic

`func NewPasswordLinkPublic(expiresAt time.Time, mailSent bool, ) *PasswordLinkPublic`

NewPasswordLinkPublic instantiates a new PasswordLinkPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPasswordLinkPublicWithDefaults

`func NewPasswordLinkPublicWithDefaults() *PasswordLinkPublic`

NewPasswordLinkPublicWithDefaults instantiates a new PasswordLinkPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiresAt

`func (o *PasswordLinkPublic) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *PasswordLinkPublic) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *PasswordLinkPublic) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.


### GetLink

`func (o *PasswordLinkPublic) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *PasswordLinkPublic) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *PasswordLinkPublic) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *PasswordLinkPublic) HasLink() bool`

HasLink returns a boolean if a field has been set.

### SetLinkNil

`func (o *PasswordLinkPublic) SetLinkNil(b bool)`

 SetLinkNil sets the value for Link to be an explicit nil

### UnsetLink
`func (o *PasswordLinkPublic) UnsetLink()`

UnsetLink ensures that no value is present for Link, not even an explicit nil
### GetMailSent

`func (o *PasswordLinkPublic) GetMailSent() bool`

GetMailSent returns the MailSent field if non-nil, zero value otherwise.

### GetMailSentOk

`func (o *PasswordLinkPublic) GetMailSentOk() (*bool, bool)`

GetMailSentOk returns a tuple with the MailSent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailSent

`func (o *PasswordLinkPublic) SetMailSent(v bool)`

SetMailSent sets MailSent field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


