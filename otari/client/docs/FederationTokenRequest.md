# FederationTokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assertion** | Pointer to **NullableString** | The identity token. | [optional] 
**FederationRuleId** | Pointer to **NullableString** | The id of the federation rule the identity token must match. | [optional] 
**GrantType** | Pointer to **NullableString** | urn:ietf:params:oauth:grant-type:jwt-bearer | [optional] 
**OrganizationId** | Pointer to **NullableString** | When set, the UUID of the organization of the rule&#39;s key. | [optional] 
**ServiceAccountId** | Pointer to **NullableString** | When set, the id of the rule&#39;s key. | [optional] 
**WorkspaceId** | Pointer to **NullableString** | When set, the UUID of the rule&#39;s key&#39;s workspace. | [optional] 

## Methods

### NewFederationTokenRequest

`func NewFederationTokenRequest() *FederationTokenRequest`

NewFederationTokenRequest instantiates a new FederationTokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFederationTokenRequestWithDefaults

`func NewFederationTokenRequestWithDefaults() *FederationTokenRequest`

NewFederationTokenRequestWithDefaults instantiates a new FederationTokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssertion

`func (o *FederationTokenRequest) GetAssertion() string`

GetAssertion returns the Assertion field if non-nil, zero value otherwise.

### GetAssertionOk

`func (o *FederationTokenRequest) GetAssertionOk() (*string, bool)`

GetAssertionOk returns a tuple with the Assertion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssertion

`func (o *FederationTokenRequest) SetAssertion(v string)`

SetAssertion sets Assertion field to given value.

### HasAssertion

`func (o *FederationTokenRequest) HasAssertion() bool`

HasAssertion returns a boolean if a field has been set.

### SetAssertionNil

`func (o *FederationTokenRequest) SetAssertionNil(b bool)`

 SetAssertionNil sets the value for Assertion to be an explicit nil

### UnsetAssertion
`func (o *FederationTokenRequest) UnsetAssertion()`

UnsetAssertion ensures that no value is present for Assertion, not even an explicit nil
### GetFederationRuleId

`func (o *FederationTokenRequest) GetFederationRuleId() string`

GetFederationRuleId returns the FederationRuleId field if non-nil, zero value otherwise.

### GetFederationRuleIdOk

`func (o *FederationTokenRequest) GetFederationRuleIdOk() (*string, bool)`

GetFederationRuleIdOk returns a tuple with the FederationRuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederationRuleId

`func (o *FederationTokenRequest) SetFederationRuleId(v string)`

SetFederationRuleId sets FederationRuleId field to given value.

### HasFederationRuleId

`func (o *FederationTokenRequest) HasFederationRuleId() bool`

HasFederationRuleId returns a boolean if a field has been set.

### SetFederationRuleIdNil

`func (o *FederationTokenRequest) SetFederationRuleIdNil(b bool)`

 SetFederationRuleIdNil sets the value for FederationRuleId to be an explicit nil

### UnsetFederationRuleId
`func (o *FederationTokenRequest) UnsetFederationRuleId()`

UnsetFederationRuleId ensures that no value is present for FederationRuleId, not even an explicit nil
### GetGrantType

`func (o *FederationTokenRequest) GetGrantType() string`

GetGrantType returns the GrantType field if non-nil, zero value otherwise.

### GetGrantTypeOk

`func (o *FederationTokenRequest) GetGrantTypeOk() (*string, bool)`

GetGrantTypeOk returns a tuple with the GrantType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantType

`func (o *FederationTokenRequest) SetGrantType(v string)`

SetGrantType sets GrantType field to given value.

### HasGrantType

`func (o *FederationTokenRequest) HasGrantType() bool`

HasGrantType returns a boolean if a field has been set.

### SetGrantTypeNil

`func (o *FederationTokenRequest) SetGrantTypeNil(b bool)`

 SetGrantTypeNil sets the value for GrantType to be an explicit nil

### UnsetGrantType
`func (o *FederationTokenRequest) UnsetGrantType()`

UnsetGrantType ensures that no value is present for GrantType, not even an explicit nil
### GetOrganizationId

`func (o *FederationTokenRequest) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *FederationTokenRequest) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *FederationTokenRequest) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *FederationTokenRequest) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### SetOrganizationIdNil

`func (o *FederationTokenRequest) SetOrganizationIdNil(b bool)`

 SetOrganizationIdNil sets the value for OrganizationId to be an explicit nil

### UnsetOrganizationId
`func (o *FederationTokenRequest) UnsetOrganizationId()`

UnsetOrganizationId ensures that no value is present for OrganizationId, not even an explicit nil
### GetServiceAccountId

`func (o *FederationTokenRequest) GetServiceAccountId() string`

GetServiceAccountId returns the ServiceAccountId field if non-nil, zero value otherwise.

### GetServiceAccountIdOk

`func (o *FederationTokenRequest) GetServiceAccountIdOk() (*string, bool)`

GetServiceAccountIdOk returns a tuple with the ServiceAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccountId

`func (o *FederationTokenRequest) SetServiceAccountId(v string)`

SetServiceAccountId sets ServiceAccountId field to given value.

### HasServiceAccountId

`func (o *FederationTokenRequest) HasServiceAccountId() bool`

HasServiceAccountId returns a boolean if a field has been set.

### SetServiceAccountIdNil

`func (o *FederationTokenRequest) SetServiceAccountIdNil(b bool)`

 SetServiceAccountIdNil sets the value for ServiceAccountId to be an explicit nil

### UnsetServiceAccountId
`func (o *FederationTokenRequest) UnsetServiceAccountId()`

UnsetServiceAccountId ensures that no value is present for ServiceAccountId, not even an explicit nil
### GetWorkspaceId

`func (o *FederationTokenRequest) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *FederationTokenRequest) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *FederationTokenRequest) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.

### HasWorkspaceId

`func (o *FederationTokenRequest) HasWorkspaceId() bool`

HasWorkspaceId returns a boolean if a field has been set.

### SetWorkspaceIdNil

`func (o *FederationTokenRequest) SetWorkspaceIdNil(b bool)`

 SetWorkspaceIdNil sets the value for WorkspaceId to be an explicit nil

### UnsetWorkspaceId
`func (o *FederationTokenRequest) UnsetWorkspaceId()`

UnsetWorkspaceId ensures that no value is present for WorkspaceId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


