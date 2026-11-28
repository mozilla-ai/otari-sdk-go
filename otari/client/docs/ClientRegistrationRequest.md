# ClientRegistrationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientName** | Pointer to **NullableString** | Filter to a single event type or metric name (e.g. &#39;tool_result&#39;, &#39;claude_code.commit.count&#39;) | [optional] 
**RedirectUris** | **[]string** |  | 

## Methods

### NewClientRegistrationRequest

`func NewClientRegistrationRequest(redirectUris []string, ) *ClientRegistrationRequest`

NewClientRegistrationRequest instantiates a new ClientRegistrationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientRegistrationRequestWithDefaults

`func NewClientRegistrationRequestWithDefaults() *ClientRegistrationRequest`

NewClientRegistrationRequestWithDefaults instantiates a new ClientRegistrationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientName

`func (o *ClientRegistrationRequest) GetClientName() string`

GetClientName returns the ClientName field if non-nil, zero value otherwise.

### GetClientNameOk

`func (o *ClientRegistrationRequest) GetClientNameOk() (*string, bool)`

GetClientNameOk returns a tuple with the ClientName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientName

`func (o *ClientRegistrationRequest) SetClientName(v string)`

SetClientName sets ClientName field to given value.

### HasClientName

`func (o *ClientRegistrationRequest) HasClientName() bool`

HasClientName returns a boolean if a field has been set.

### SetClientNameNil

`func (o *ClientRegistrationRequest) SetClientNameNil(b bool)`

 SetClientNameNil sets the value for ClientName to be an explicit nil

### UnsetClientName
`func (o *ClientRegistrationRequest) UnsetClientName()`

UnsetClientName ensures that no value is present for ClientName, not even an explicit nil
### GetRedirectUris

`func (o *ClientRegistrationRequest) GetRedirectUris() []string`

GetRedirectUris returns the RedirectUris field if non-nil, zero value otherwise.

### GetRedirectUrisOk

`func (o *ClientRegistrationRequest) GetRedirectUrisOk() (*[]string, bool)`

GetRedirectUrisOk returns a tuple with the RedirectUris field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedirectUris

`func (o *ClientRegistrationRequest) SetRedirectUris(v []string)`

SetRedirectUris sets RedirectUris field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


