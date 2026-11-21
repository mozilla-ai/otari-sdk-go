# McpCatalogServerPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **NullableString** |  | 
**IconUrl** | **NullableString** | The https image the registry entry names as the server&#39;s icon, null when it names none. | 
**Id** | **string** |  | 
**LogoTone** | **NullableString** | Where the logo needs a tile to stay visible: &#x60;dark&#x60; ink on a dark background, &#x60;light&#x60; ink on a light background, &#x60;tile&#x60; on every background. Null when it reads on both. | 
**LogoUrl** | **NullableString** | Where this gateway serves the server&#39;s logo, which it fetched once and keeps. Null when it looked and found none. A URL without &#x60;v&#x60; has not been looked for yet, and asking for it starts the look. | 
**Name** | **string** | The source&#39;s own identifier, such as &#x60;com.example/mcp&#x60;. | 
**PublisherKind** | **string** | &#x60;vendor&#x60; when the server is its maker&#39;s own: the endpoint lives on the domain the registry verified for the publisher&#39;s namespace, the server is named for that domain and nothing else, and the domain is not a hosting platform&#39;s. A website or title the publisher writes never earns it, and neither does a GitHub namespace. A few makers&#39; servers the rule cannot recognize are listed by hand. &#x60;community&#x60; for every other, such as a gateway to another company&#39;s API. | 
**RequiresToken** | **bool** | The source says a call needs an Authorization header. False means nothing was declared, not that the server is open: one that signs users in with OAuth declares nothing. | 
**Source** | **string** | Where the entry came from. &#x60;registry&#x60; is the official MCP Registry; &#x60;curated&#x60; is a short list bundled with Otari of well-known servers the registry does not list. | 
**Title** | **string** |  | 
**Url** | **string** | The server&#39;s streamable HTTP endpoint. | 
**Version** | **NullableString** |  | 
**WebsiteUrl** | **NullableString** |  | 

## Methods

### NewMcpCatalogServerPublic

`func NewMcpCatalogServerPublic(description NullableString, iconUrl NullableString, id string, logoTone NullableString, logoUrl NullableString, name string, publisherKind string, requiresToken bool, source string, title string, url string, version NullableString, websiteUrl NullableString, ) *McpCatalogServerPublic`

NewMcpCatalogServerPublic instantiates a new McpCatalogServerPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMcpCatalogServerPublicWithDefaults

`func NewMcpCatalogServerPublicWithDefaults() *McpCatalogServerPublic`

NewMcpCatalogServerPublicWithDefaults instantiates a new McpCatalogServerPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *McpCatalogServerPublic) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *McpCatalogServerPublic) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *McpCatalogServerPublic) SetDescription(v string)`

SetDescription sets Description field to given value.


### SetDescriptionNil

`func (o *McpCatalogServerPublic) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *McpCatalogServerPublic) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetIconUrl

`func (o *McpCatalogServerPublic) GetIconUrl() string`

GetIconUrl returns the IconUrl field if non-nil, zero value otherwise.

### GetIconUrlOk

`func (o *McpCatalogServerPublic) GetIconUrlOk() (*string, bool)`

GetIconUrlOk returns a tuple with the IconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIconUrl

`func (o *McpCatalogServerPublic) SetIconUrl(v string)`

SetIconUrl sets IconUrl field to given value.


### SetIconUrlNil

`func (o *McpCatalogServerPublic) SetIconUrlNil(b bool)`

 SetIconUrlNil sets the value for IconUrl to be an explicit nil

### UnsetIconUrl
`func (o *McpCatalogServerPublic) UnsetIconUrl()`

UnsetIconUrl ensures that no value is present for IconUrl, not even an explicit nil
### GetId

`func (o *McpCatalogServerPublic) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *McpCatalogServerPublic) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *McpCatalogServerPublic) SetId(v string)`

SetId sets Id field to given value.


### GetLogoTone

`func (o *McpCatalogServerPublic) GetLogoTone() string`

GetLogoTone returns the LogoTone field if non-nil, zero value otherwise.

### GetLogoToneOk

`func (o *McpCatalogServerPublic) GetLogoToneOk() (*string, bool)`

GetLogoToneOk returns a tuple with the LogoTone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoTone

`func (o *McpCatalogServerPublic) SetLogoTone(v string)`

SetLogoTone sets LogoTone field to given value.


### SetLogoToneNil

`func (o *McpCatalogServerPublic) SetLogoToneNil(b bool)`

 SetLogoToneNil sets the value for LogoTone to be an explicit nil

### UnsetLogoTone
`func (o *McpCatalogServerPublic) UnsetLogoTone()`

UnsetLogoTone ensures that no value is present for LogoTone, not even an explicit nil
### GetLogoUrl

`func (o *McpCatalogServerPublic) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *McpCatalogServerPublic) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *McpCatalogServerPublic) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.


### SetLogoUrlNil

`func (o *McpCatalogServerPublic) SetLogoUrlNil(b bool)`

 SetLogoUrlNil sets the value for LogoUrl to be an explicit nil

### UnsetLogoUrl
`func (o *McpCatalogServerPublic) UnsetLogoUrl()`

UnsetLogoUrl ensures that no value is present for LogoUrl, not even an explicit nil
### GetName

`func (o *McpCatalogServerPublic) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *McpCatalogServerPublic) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *McpCatalogServerPublic) SetName(v string)`

SetName sets Name field to given value.


### GetPublisherKind

`func (o *McpCatalogServerPublic) GetPublisherKind() string`

GetPublisherKind returns the PublisherKind field if non-nil, zero value otherwise.

### GetPublisherKindOk

`func (o *McpCatalogServerPublic) GetPublisherKindOk() (*string, bool)`

GetPublisherKindOk returns a tuple with the PublisherKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublisherKind

`func (o *McpCatalogServerPublic) SetPublisherKind(v string)`

SetPublisherKind sets PublisherKind field to given value.


### GetRequiresToken

`func (o *McpCatalogServerPublic) GetRequiresToken() bool`

GetRequiresToken returns the RequiresToken field if non-nil, zero value otherwise.

### GetRequiresTokenOk

`func (o *McpCatalogServerPublic) GetRequiresTokenOk() (*bool, bool)`

GetRequiresTokenOk returns a tuple with the RequiresToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiresToken

`func (o *McpCatalogServerPublic) SetRequiresToken(v bool)`

SetRequiresToken sets RequiresToken field to given value.


### GetSource

`func (o *McpCatalogServerPublic) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *McpCatalogServerPublic) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *McpCatalogServerPublic) SetSource(v string)`

SetSource sets Source field to given value.


### GetTitle

`func (o *McpCatalogServerPublic) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *McpCatalogServerPublic) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *McpCatalogServerPublic) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetUrl

`func (o *McpCatalogServerPublic) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *McpCatalogServerPublic) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *McpCatalogServerPublic) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetVersion

`func (o *McpCatalogServerPublic) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *McpCatalogServerPublic) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *McpCatalogServerPublic) SetVersion(v string)`

SetVersion sets Version field to given value.


### SetVersionNil

`func (o *McpCatalogServerPublic) SetVersionNil(b bool)`

 SetVersionNil sets the value for Version to be an explicit nil

### UnsetVersion
`func (o *McpCatalogServerPublic) UnsetVersion()`

UnsetVersion ensures that no value is present for Version, not even an explicit nil
### GetWebsiteUrl

`func (o *McpCatalogServerPublic) GetWebsiteUrl() string`

GetWebsiteUrl returns the WebsiteUrl field if non-nil, zero value otherwise.

### GetWebsiteUrlOk

`func (o *McpCatalogServerPublic) GetWebsiteUrlOk() (*string, bool)`

GetWebsiteUrlOk returns a tuple with the WebsiteUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsiteUrl

`func (o *McpCatalogServerPublic) SetWebsiteUrl(v string)`

SetWebsiteUrl sets WebsiteUrl field to given value.


### SetWebsiteUrlNil

`func (o *McpCatalogServerPublic) SetWebsiteUrlNil(b bool)`

 SetWebsiteUrlNil sets the value for WebsiteUrl to be an explicit nil

### UnsetWebsiteUrl
`func (o *McpCatalogServerPublic) UnsetWebsiteUrl()`

UnsetWebsiteUrl ensures that no value is present for WebsiteUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


