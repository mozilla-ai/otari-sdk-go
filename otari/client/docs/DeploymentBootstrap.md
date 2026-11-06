# DeploymentBootstrap

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocsUrl** | **NullableString** | Where this deployment&#39;s documentation lives, when it is not docs.otari.ai. Set, the dashboard&#39;s Documentation links open it in a new tab; null, they open https://docs.otari.ai/. A link target an operator configured, validated at startup as an absolute http(s) URL carrying no credential, since this response is unauthenticated. | 
**FederationMemberRules** | **bool** | Whether an organization&#39;s owners and admins may give their own API keys federation rules (the federation_member_rules setting), so the dashboard offers workload identity on their keys. | 
**FeedbackEnabled** | **bool** | Whether this deployment accepts deliberate feedback submissions. | 
**MailReady** | **bool** | Whether this deployment can deliver a message carrying a link back to itself (an invitation&#39;s accept link, and the verification and reset links to come), not merely whether a transport is configured: it also needs to know its own public URL to put in one. Lets the dashboard disable or hide a mail-dependent affordance instead of offering one that would fail at send time. Every message this deployment sends carries such a link, which is why this is one flag and not one per feature. | 
**MaintenanceMode** | **bool** | Whether this deployment is refusing new dashboard sign-ins while an operator redeploys it. The sign-in screen says so rather than presenting a form whose only outcome is a 503. Sessions already issued keep working, and the API is unaffected. | 
**OauthProviders** | **[]string** | OAuth providers this deployment can sign somebody in with, sorted, one entry per provider with a client ID, a client secret and a public_base_url to build a redirect URI from. The sign-in screen renders a button per entry and none at all when the list is empty, so a provider nobody configured is absent rather than offered and then refused. Additive to sign_in_methods rather than part of it: an OAuth sign-in coexists with whichever typed credential is current, the way a passkey does. | 
**OpenSignup** | **bool** | Whether POST /api/v1/auth/signup creates an account for an address nobody has added yet, each with an organization of its own, or only lets an address an admin already put on the roster set its password. The signup page reads as registration or as claiming an invitation accordingly, and the sign-in screen links to it with the wording that matches. | 
**PasskeysEnabled** | **bool** | Whether this deployment offers passkeys at all (the passkeys_enabled setting). False means the passkey routes are not mounted and the dashboard hides its passkey page, which is a different answer from passkeys_ready: a deployment that merely has not configured a relying party is ready for nobody yet but still shows the page that says so. | 
**PasskeysReady** | **bool** | Whether this deployment can run a passkey ceremony at all: it has a relying-party ID (webauthn_rp_id, or derived from public_base_url) and an origin to serve one from. Distinct from &#39;passkey&#39; in sign_in_methods, which is narrower and answers whether a registered passkey could sign somebody in *right now*: an operator with none yet needs this one, or the page that registers the first would be hidden from them. | 
**PrivacyUrl** | **NullableString** | Where this deployment&#39;s privacy notice lives. Set, the account menu&#39;s Data &amp; Privacy row links to it; null, no address is configured and that row stays disabled, carrying the standing note that there is nothing to configure there yet. A link target an operator configured, validated at startup as an absolute http(s) URL carrying no credential, since this response is unauthenticated. | 
**PublicCatalog** | Pointer to **bool** | Whether the model catalog is served to a visitor with no session: GET /api/v1/catalog/models answers anonymously and the dashboard renders Models ahead of sign-in. | [optional] [default to false]
**SignInMethods** | **[]string** | How POST /api/v1/auth/session may be authenticated right now, sorted. &#39;master_key&#39; is the first-boot credential and is offered until the operator identity has a password, which is what claiming the deployment means; past that it stays the credential for the management API but is no longer a dashboard login. &#39;password&#39; is offered while any active identity holds one, which is not the same question and not always the later half of it: a member can hold a password on a deployment whose operator never claimed it, so both typed credentials can appear together. &#39;passkey&#39; appears alongside either when this deployment is configured for WebAuthn and holds at least one passkey that its current relying-party ID can assert. The login page renders from this rather than trying a credential to find out. | 
**SiteUrl** | **NullableString** | Where this deployment&#39;s public website lives. Set, the logo on the pages a visitor reaches without an account links to it; null, it links to the public catalog where public_catalog is true, and is not a link otherwise. A link target an operator configured, validated at startup as an absolute http(s) URL carrying no credential, since this response is unauthenticated. | 
**Surfaces** | **[]string** | Management API groups this deployment serves, sorted, which is what its dashboard pages gate on: a fixed set plus one for each enabled feature. Named surfaces, not capabilities: capability is otari.ai&#39;s word for the entitlement (licensing) axis. | 
**TermsUrl** | **NullableString** | Where this deployment&#39;s terms of service live. Set, the account menu carries a Terms of service row pointing at them; null, no address is configured and the menu carries no such row. A link target an operator configured, validated at startup as an absolute http(s) URL carrying no credential, since this response is unauthenticated. | 

## Methods

### NewDeploymentBootstrap

`func NewDeploymentBootstrap(docsUrl NullableString, federationMemberRules bool, feedbackEnabled bool, mailReady bool, maintenanceMode bool, oauthProviders []string, openSignup bool, passkeysEnabled bool, passkeysReady bool, privacyUrl NullableString, signInMethods []string, siteUrl NullableString, surfaces []string, termsUrl NullableString, ) *DeploymentBootstrap`

NewDeploymentBootstrap instantiates a new DeploymentBootstrap object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentBootstrapWithDefaults

`func NewDeploymentBootstrapWithDefaults() *DeploymentBootstrap`

NewDeploymentBootstrapWithDefaults instantiates a new DeploymentBootstrap object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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
### GetFederationMemberRules

`func (o *DeploymentBootstrap) GetFederationMemberRules() bool`

GetFederationMemberRules returns the FederationMemberRules field if non-nil, zero value otherwise.

### GetFederationMemberRulesOk

`func (o *DeploymentBootstrap) GetFederationMemberRulesOk() (*bool, bool)`

GetFederationMemberRulesOk returns a tuple with the FederationMemberRules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFederationMemberRules

`func (o *DeploymentBootstrap) SetFederationMemberRules(v bool)`

SetFederationMemberRules sets FederationMemberRules field to given value.


### GetFeedbackEnabled

`func (o *DeploymentBootstrap) GetFeedbackEnabled() bool`

GetFeedbackEnabled returns the FeedbackEnabled field if non-nil, zero value otherwise.

### GetFeedbackEnabledOk

`func (o *DeploymentBootstrap) GetFeedbackEnabledOk() (*bool, bool)`

GetFeedbackEnabledOk returns a tuple with the FeedbackEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeedbackEnabled

`func (o *DeploymentBootstrap) SetFeedbackEnabled(v bool)`

SetFeedbackEnabled sets FeedbackEnabled field to given value.


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


### GetOpenSignup

`func (o *DeploymentBootstrap) GetOpenSignup() bool`

GetOpenSignup returns the OpenSignup field if non-nil, zero value otherwise.

### GetOpenSignupOk

`func (o *DeploymentBootstrap) GetOpenSignupOk() (*bool, bool)`

GetOpenSignupOk returns a tuple with the OpenSignup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenSignup

`func (o *DeploymentBootstrap) SetOpenSignup(v bool)`

SetOpenSignup sets OpenSignup field to given value.


### GetPasskeysEnabled

`func (o *DeploymentBootstrap) GetPasskeysEnabled() bool`

GetPasskeysEnabled returns the PasskeysEnabled field if non-nil, zero value otherwise.

### GetPasskeysEnabledOk

`func (o *DeploymentBootstrap) GetPasskeysEnabledOk() (*bool, bool)`

GetPasskeysEnabledOk returns a tuple with the PasskeysEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasskeysEnabled

`func (o *DeploymentBootstrap) SetPasskeysEnabled(v bool)`

SetPasskeysEnabled sets PasskeysEnabled field to given value.


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
### GetPublicCatalog

`func (o *DeploymentBootstrap) GetPublicCatalog() bool`

GetPublicCatalog returns the PublicCatalog field if non-nil, zero value otherwise.

### GetPublicCatalogOk

`func (o *DeploymentBootstrap) GetPublicCatalogOk() (*bool, bool)`

GetPublicCatalogOk returns a tuple with the PublicCatalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicCatalog

`func (o *DeploymentBootstrap) SetPublicCatalog(v bool)`

SetPublicCatalog sets PublicCatalog field to given value.

### HasPublicCatalog

`func (o *DeploymentBootstrap) HasPublicCatalog() bool`

HasPublicCatalog returns a boolean if a field has been set.

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


### GetSiteUrl

`func (o *DeploymentBootstrap) GetSiteUrl() string`

GetSiteUrl returns the SiteUrl field if non-nil, zero value otherwise.

### GetSiteUrlOk

`func (o *DeploymentBootstrap) GetSiteUrlOk() (*string, bool)`

GetSiteUrlOk returns a tuple with the SiteUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteUrl

`func (o *DeploymentBootstrap) SetSiteUrl(v string)`

SetSiteUrl sets SiteUrl field to given value.


### SetSiteUrlNil

`func (o *DeploymentBootstrap) SetSiteUrlNil(b bool)`

 SetSiteUrlNil sets the value for SiteUrl to be an explicit nil

### UnsetSiteUrl
`func (o *DeploymentBootstrap) UnsetSiteUrl()`

UnsetSiteUrl ensures that no value is present for SiteUrl, not even an explicit nil
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


