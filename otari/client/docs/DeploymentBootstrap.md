# DeploymentBootstrap

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataPlaneUrl** | **NullableString** | Where this deployment&#39;s inference traffic belongs, when it is not served here. The mirror of management_url: that one says where management lives when this deployment is not the control plane, this one says where the data plane is when this deployment is not it. Set only by a hosted control plane, which serves the dashboard but not inference (otari#822); null for standalone and hybrid, both of which serve inference at the address that reached this page. Not a human link target like management_url: it is the gateway&#39;s bare address, which the dashboard suffixes with /v1 to build its request snippets. So it must carry no /v1 path segment anywhere (a value ending in /v1, or a whole endpoint like /v1/chat/completions, renders that path twice) and no credential, since this response is unauthenticated. This gateway refuses both at startup; any deployment serving this contract should publish the same shape. Null on a hosted deployment means unconfigured, and the dashboard then shows no snippet rather than one naming this host. | 
**DeploymentType** | **string** | Which deployment serves this URL. &#39;standalone&#39; owns its own data and serves one tenant; &#39;hosted&#39; owns its own data and serves many (otari.ai, or any deployment run as a control plane), which is why its management surfaces are the per-organization ones; &#39;hybrid&#39; is a gateway attached to otari.ai, which is data-plane only and holds no management surface of its own. | 
**DocsUrl** | **NullableString** | Where this deployment&#39;s documentation lives, when it is not the operator guide bundled with the gateway. Set, the dashboard&#39;s Documentation links open it in a new tab; null, they go to the bundled guide at /#/docs, which stays served either way. A link target an operator configured, validated at startup as an absolute http(s) URL carrying no credential, since this response is unauthenticated. | 
**MailReady** | **bool** | Whether this deployment can deliver a message carrying a link back to itself (an invitation&#39;s accept link, and the verification and reset links to come), not merely whether a transport is configured: it also needs to know its own public URL to put in one. Lets the dashboard disable or hide a mail-dependent affordance instead of offering one that would fail at send time. Every message this control plane sends carries such a link, which is why this is one flag and not one per feature. False for a hybrid gateway, whose control plane is otari.ai and which sends no mail of its own. | 
**MaintenanceMode** | **bool** | Whether this deployment is refusing new dashboard sign-ins while an operator redeploys it. The sign-in screen says so rather than presenting a form whose only outcome is a 503. Sessions already issued keep working, and the management API and the data plane are unaffected. False for a hybrid gateway, which issues no session. | 
**ManagementUrl** | **NullableString** | Where the authoritative control plane lives when it is not this deployment. Set for a hybrid gateway so its landing page can link to otari.ai; null otherwise. | 
**OauthProviders** | **[]string** | OAuth providers this deployment can sign somebody in with, sorted, one entry per provider with a client ID, a client secret and a public_base_url to build a redirect URI from. The sign-in screen renders a button per entry and none at all when the list is empty, so a provider nobody configured is absent rather than offered and then refused. Additive to sign_in_methods rather than part of it: an OAuth sign-in coexists with whichever typed credential is current, the way a passkey does. Empty for a hybrid gateway, which issues no session. | 
**PasskeysReady** | **bool** | Whether this deployment can run a passkey ceremony at all: it has a relying-party ID (webauthn_rp_id, or derived from public_base_url) and an origin to serve one from. Distinct from &#39;passkey&#39; in sign_in_methods, which is narrower and answers whether a registered passkey could sign somebody in *right now*: an operator with none yet needs this one, or the page that registers the first would be hidden from them. False for a hybrid gateway, which issues no session of its own. | 
**PrivacyUrl** | **NullableString** | Where this deployment&#39;s privacy notice lives. Set, the account menu&#39;s Data &amp; Privacy row links to it; null, no address is configured and that row stays disabled, carrying the standing note that there is nothing to configure there yet. A link target an operator configured, validated at startup as an absolute http(s) URL carrying no credential, since this response is unauthenticated. | 
**SessionType** | **string** | The kind of session this deployment issues, not whether the caller holds one. &#39;local_operator&#39; is the standalone operator sign-in (see sign_in_methods for which credential it currently accepts), &#39;hosted_user&#39; an otari.ai account, and &#39;none&#39; a deployment that issues no management session at all. | 
**SignInMethods** | **[]string** | How POST /v1/auth/session may be authenticated right now, sorted. &#39;master_key&#39; is the first-boot credential and is offered until the operator identity has a password, which is what claiming the deployment means; &#39;password&#39; replaces it from then on, and the master key stays the credential for the management API. &#39;passkey&#39; appears alongside either one when this deployment is configured for WebAuthn and holds at least one passkey that its current relying-party ID can assert. Empty for a hybrid gateway, which issues no session. The login page renders from this rather than trying a credential to find out. | 
**Surfaces** | **[]string** | Management API groups this deployment serves, sorted, which is what its dashboard pages gate on. Named surfaces, not capabilities: capability is otari.ai&#39;s word for the entitlement (licensing) axis, and this is the deployment (topology) axis. Empty for a hybrid gateway. | 
**TermsUrl** | **NullableString** | Where this deployment&#39;s terms of service live. Set, the account menu carries a Terms of service row pointing at them; null, no address is configured and the menu carries no such row. A link target an operator configured, validated at startup as an absolute http(s) URL carrying no credential, since this response is unauthenticated. | 

## Methods

### NewDeploymentBootstrap

`func NewDeploymentBootstrap(dataPlaneUrl NullableString, deploymentType string, docsUrl NullableString, mailReady bool, maintenanceMode bool, managementUrl NullableString, oauthProviders []string, passkeysReady bool, privacyUrl NullableString, sessionType string, signInMethods []string, surfaces []string, termsUrl NullableString, ) *DeploymentBootstrap`

NewDeploymentBootstrap instantiates a new DeploymentBootstrap object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentBootstrapWithDefaults

`func NewDeploymentBootstrapWithDefaults() *DeploymentBootstrap`

NewDeploymentBootstrapWithDefaults instantiates a new DeploymentBootstrap object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataPlaneUrl

`func (o *DeploymentBootstrap) GetDataPlaneUrl() string`

GetDataPlaneUrl returns the DataPlaneUrl field if non-nil, zero value otherwise.

### GetDataPlaneUrlOk

`func (o *DeploymentBootstrap) GetDataPlaneUrlOk() (*string, bool)`

GetDataPlaneUrlOk returns a tuple with the DataPlaneUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataPlaneUrl

`func (o *DeploymentBootstrap) SetDataPlaneUrl(v string)`

SetDataPlaneUrl sets DataPlaneUrl field to given value.


### SetDataPlaneUrlNil

`func (o *DeploymentBootstrap) SetDataPlaneUrlNil(b bool)`

 SetDataPlaneUrlNil sets the value for DataPlaneUrl to be an explicit nil

### UnsetDataPlaneUrl
`func (o *DeploymentBootstrap) UnsetDataPlaneUrl()`

UnsetDataPlaneUrl ensures that no value is present for DataPlaneUrl, not even an explicit nil
### GetDeploymentType

`func (o *DeploymentBootstrap) GetDeploymentType() string`

GetDeploymentType returns the DeploymentType field if non-nil, zero value otherwise.

### GetDeploymentTypeOk

`func (o *DeploymentBootstrap) GetDeploymentTypeOk() (*string, bool)`

GetDeploymentTypeOk returns a tuple with the DeploymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentType

`func (o *DeploymentBootstrap) SetDeploymentType(v string)`

SetDeploymentType sets DeploymentType field to given value.


### GetDocsUrl

`func (o *DeploymentBootstrap) GetDocsUrl() string`

GetDocsUrl returns the DocsUrl field if non-nil, zero value otherwise.

### GetDocsUrlOk

`func (o *DeploymentBootstrap) GetDocsUrlOk() (*string, bool)`

GetDocsUrlOk returns a tuple with the DocsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocsUrl

`func (o *DeploymentBootstrap) SetDocsUrl(v string)`

SetDocsUrl sets DocsUrl field to given value.


### SetDocsUrlNil

`func (o *DeploymentBootstrap) SetDocsUrlNil(b bool)`

 SetDocsUrlNil sets the value for DocsUrl to be an explicit nil

### UnsetDocsUrl
`func (o *DeploymentBootstrap) UnsetDocsUrl()`

UnsetDocsUrl ensures that no value is present for DocsUrl, not even an explicit nil
### GetMailReady

`func (o *DeploymentBootstrap) GetMailReady() bool`

GetMailReady returns the MailReady field if non-nil, zero value otherwise.

### GetMailReadyOk

`func (o *DeploymentBootstrap) GetMailReadyOk() (*bool, bool)`

GetMailReadyOk returns a tuple with the MailReady field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMailReady

`func (o *DeploymentBootstrap) SetMailReady(v bool)`

SetMailReady sets MailReady field to given value.


### GetMaintenanceMode

`func (o *DeploymentBootstrap) GetMaintenanceMode() bool`

GetMaintenanceMode returns the MaintenanceMode field if non-nil, zero value otherwise.

### GetMaintenanceModeOk

`func (o *DeploymentBootstrap) GetMaintenanceModeOk() (*bool, bool)`

GetMaintenanceModeOk returns a tuple with the MaintenanceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceMode

`func (o *DeploymentBootstrap) SetMaintenanceMode(v bool)`

SetMaintenanceMode sets MaintenanceMode field to given value.


### GetManagementUrl

`func (o *DeploymentBootstrap) GetManagementUrl() string`

GetManagementUrl returns the ManagementUrl field if non-nil, zero value otherwise.

### GetManagementUrlOk

`func (o *DeploymentBootstrap) GetManagementUrlOk() (*string, bool)`

GetManagementUrlOk returns a tuple with the ManagementUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementUrl

`func (o *DeploymentBootstrap) SetManagementUrl(v string)`

SetManagementUrl sets ManagementUrl field to given value.


### SetManagementUrlNil

`func (o *DeploymentBootstrap) SetManagementUrlNil(b bool)`

 SetManagementUrlNil sets the value for ManagementUrl to be an explicit nil

### UnsetManagementUrl
`func (o *DeploymentBootstrap) UnsetManagementUrl()`

UnsetManagementUrl ensures that no value is present for ManagementUrl, not even an explicit nil
### GetOauthProviders

`func (o *DeploymentBootstrap) GetOauthProviders() []string`

GetOauthProviders returns the OauthProviders field if non-nil, zero value otherwise.

### GetOauthProvidersOk

`func (o *DeploymentBootstrap) GetOauthProvidersOk() (*[]string, bool)`

GetOauthProvidersOk returns a tuple with the OauthProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthProviders

`func (o *DeploymentBootstrap) SetOauthProviders(v []string)`

SetOauthProviders sets OauthProviders field to given value.


### GetPasskeysReady

`func (o *DeploymentBootstrap) GetPasskeysReady() bool`

GetPasskeysReady returns the PasskeysReady field if non-nil, zero value otherwise.

### GetPasskeysReadyOk

`func (o *DeploymentBootstrap) GetPasskeysReadyOk() (*bool, bool)`

GetPasskeysReadyOk returns a tuple with the PasskeysReady field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasskeysReady

`func (o *DeploymentBootstrap) SetPasskeysReady(v bool)`

SetPasskeysReady sets PasskeysReady field to given value.


### GetPrivacyUrl

`func (o *DeploymentBootstrap) GetPrivacyUrl() string`

GetPrivacyUrl returns the PrivacyUrl field if non-nil, zero value otherwise.

### GetPrivacyUrlOk

`func (o *DeploymentBootstrap) GetPrivacyUrlOk() (*string, bool)`

GetPrivacyUrlOk returns a tuple with the PrivacyUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyUrl

`func (o *DeploymentBootstrap) SetPrivacyUrl(v string)`

SetPrivacyUrl sets PrivacyUrl field to given value.


### SetPrivacyUrlNil

`func (o *DeploymentBootstrap) SetPrivacyUrlNil(b bool)`

 SetPrivacyUrlNil sets the value for PrivacyUrl to be an explicit nil

### UnsetPrivacyUrl
`func (o *DeploymentBootstrap) UnsetPrivacyUrl()`

UnsetPrivacyUrl ensures that no value is present for PrivacyUrl, not even an explicit nil
### GetSessionType

`func (o *DeploymentBootstrap) GetSessionType() string`

GetSessionType returns the SessionType field if non-nil, zero value otherwise.

### GetSessionTypeOk

`func (o *DeploymentBootstrap) GetSessionTypeOk() (*string, bool)`

GetSessionTypeOk returns a tuple with the SessionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionType

`func (o *DeploymentBootstrap) SetSessionType(v string)`

SetSessionType sets SessionType field to given value.


### GetSignInMethods

`func (o *DeploymentBootstrap) GetSignInMethods() []string`

GetSignInMethods returns the SignInMethods field if non-nil, zero value otherwise.

### GetSignInMethodsOk

`func (o *DeploymentBootstrap) GetSignInMethodsOk() (*[]string, bool)`

GetSignInMethodsOk returns a tuple with the SignInMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignInMethods

`func (o *DeploymentBootstrap) SetSignInMethods(v []string)`

SetSignInMethods sets SignInMethods field to given value.


### GetSurfaces

`func (o *DeploymentBootstrap) GetSurfaces() []string`

GetSurfaces returns the Surfaces field if non-nil, zero value otherwise.

### GetSurfacesOk

`func (o *DeploymentBootstrap) GetSurfacesOk() (*[]string, bool)`

GetSurfacesOk returns a tuple with the Surfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurfaces

`func (o *DeploymentBootstrap) SetSurfaces(v []string)`

SetSurfaces sets Surfaces field to given value.


### GetTermsUrl

`func (o *DeploymentBootstrap) GetTermsUrl() string`

GetTermsUrl returns the TermsUrl field if non-nil, zero value otherwise.

### GetTermsUrlOk

`func (o *DeploymentBootstrap) GetTermsUrlOk() (*string, bool)`

GetTermsUrlOk returns a tuple with the TermsUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsUrl

`func (o *DeploymentBootstrap) SetTermsUrl(v string)`

SetTermsUrl sets TermsUrl field to given value.


### SetTermsUrlNil

`func (o *DeploymentBootstrap) SetTermsUrlNil(b bool)`

 SetTermsUrlNil sets the value for TermsUrl to be an explicit nil

### UnsetTermsUrl
`func (o *DeploymentBootstrap) UnsetTermsUrl()`

UnsetTermsUrl ensures that no value is present for TermsUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


