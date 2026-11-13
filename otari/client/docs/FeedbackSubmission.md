# FeedbackSubmission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** |  | 
**ShareEmail** | Pointer to **NullableString** | The signed-in account&#39;s email as the form showed it, to send with the message so the Otari team can reply. Refused when the caller has no email the team can reply to, or when it is not this one. | [optional] 

## Methods

### NewFeedbackSubmission

`func NewFeedbackSubmission(message string, ) *FeedbackSubmission`

NewFeedbackSubmission instantiates a new FeedbackSubmission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeedbackSubmissionWithDefaults

`func NewFeedbackSubmissionWithDefaults() *FeedbackSubmission`

NewFeedbackSubmissionWithDefaults instantiates a new FeedbackSubmission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *FeedbackSubmission) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *FeedbackSubmission) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *FeedbackSubmission) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetShareEmail

`func (o *FeedbackSubmission) GetShareEmail() string`

GetShareEmail returns the ShareEmail field if non-nil, zero value otherwise.

### GetShareEmailOk

`func (o *FeedbackSubmission) GetShareEmailOk() (*string, bool)`

GetShareEmailOk returns a tuple with the ShareEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareEmail

`func (o *FeedbackSubmission) SetShareEmail(v string)`

SetShareEmail sets ShareEmail field to given value.

### HasShareEmail

`func (o *FeedbackSubmission) HasShareEmail() bool`

HasShareEmail returns a boolean if a field has been set.

### SetShareEmailNil

`func (o *FeedbackSubmission) SetShareEmailNil(b bool)`

 SetShareEmailNil sets the value for ShareEmail to be an explicit nil

### UnsetShareEmail
`func (o *FeedbackSubmission) UnsetShareEmail()`

UnsetShareEmail ensures that no value is present for ShareEmail, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


