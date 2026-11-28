# MailedLinkStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | [**MailedLinkStatus**](MailedLinkStatus.md) | valid, expired, or invalid. invalid covers a link that was never issued, one already used, and one whose account was deactivated. | 

## Methods

### NewMailedLinkStatusResponse

`func NewMailedLinkStatusResponse(status MailedLinkStatus, ) *MailedLinkStatusResponse`

NewMailedLinkStatusResponse instantiates a new MailedLinkStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMailedLinkStatusResponseWithDefaults

`func NewMailedLinkStatusResponseWithDefaults() *MailedLinkStatusResponse`

NewMailedLinkStatusResponseWithDefaults instantiates a new MailedLinkStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *MailedLinkStatusResponse) GetStatus() MailedLinkStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MailedLinkStatusResponse) GetStatusOk() (*MailedLinkStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MailedLinkStatusResponse) SetStatus(v MailedLinkStatus)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


