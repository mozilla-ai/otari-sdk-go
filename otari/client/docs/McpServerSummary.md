# McpServerSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**PurposeHint** | **NullableString** |  | 
**SignInRequired** | Pointer to **bool** | The server sends each user&#39;s own credential and this key&#39;s user has not signed in to it, so a call is refused with mcp_sign_in_required until they do | [optional] [default to false]
**SignInUrl** | Pointer to **NullableString** | With sign_in_required: a link that starts this key&#39;s user&#39;s own sign-in to the server, working for 30 minutes | [optional] 

## Methods

### NewMcpServerSummary

`func NewMcpServerSummary(id string, name string, purposeHint NullableString, ) *McpServerSummary`

NewMcpServerSummary instantiates a new McpServerSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMcpServerSummaryWithDefaults

`func NewMcpServerSummaryWithDefaults() *McpServerSummary`

NewMcpServerSummaryWithDefaults instantiates a new McpServerSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *McpServerSummary) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *McpServerSummary) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *McpServerSummary) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *McpServerSummary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *McpServerSummary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *McpServerSummary) SetName(v string)`

SetName sets Name field to given value.


### GetPurposeHint

`func (o *McpServerSummary) GetPurposeHint() string`

GetPurposeHint returns the PurposeHint field if non-nil, zero value otherwise.

### GetPurposeHintOk

`func (o *McpServerSummary) GetPurposeHintOk() (*string, bool)`

GetPurposeHintOk returns a tuple with the PurposeHint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurposeHint

`func (o *McpServerSummary) SetPurposeHint(v string)`

SetPurposeHint sets PurposeHint field to given value.


### SetPurposeHintNil

`func (o *McpServerSummary) SetPurposeHintNil(b bool)`

 SetPurposeHintNil sets the value for PurposeHint to be an explicit nil

### UnsetPurposeHint
`func (o *McpServerSummary) UnsetPurposeHint()`

UnsetPurposeHint ensures that no value is present for PurposeHint, not even an explicit nil
### GetSignInRequired

`func (o *McpServerSummary) GetSignInRequired() bool`

GetSignInRequired returns the SignInRequired field if non-nil, zero value otherwise.

### GetSignInRequiredOk

`func (o *McpServerSummary) GetSignInRequiredOk() (*bool, bool)`

GetSignInRequiredOk returns a tuple with the SignInRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignInRequired

`func (o *McpServerSummary) SetSignInRequired(v bool)`

SetSignInRequired sets SignInRequired field to given value.

### HasSignInRequired

`func (o *McpServerSummary) HasSignInRequired() bool`

HasSignInRequired returns a boolean if a field has been set.

### GetSignInUrl

`func (o *McpServerSummary) GetSignInUrl() string`

GetSignInUrl returns the SignInUrl field if non-nil, zero value otherwise.

### GetSignInUrlOk

`func (o *McpServerSummary) GetSignInUrlOk() (*string, bool)`

GetSignInUrlOk returns a tuple with the SignInUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignInUrl

`func (o *McpServerSummary) SetSignInUrl(v string)`

SetSignInUrl sets SignInUrl field to given value.

### HasSignInUrl

`func (o *McpServerSummary) HasSignInUrl() bool`

HasSignInUrl returns a boolean if a field has been set.

### SetSignInUrlNil

`func (o *McpServerSummary) SetSignInUrlNil(b bool)`

 SetSignInUrlNil sets the value for SignInUrl to be an explicit nil

### UnsetSignInUrl
`func (o *McpServerSummary) UnsetSignInUrl()`

UnsetSignInUrl ensures that no value is present for SignInUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


