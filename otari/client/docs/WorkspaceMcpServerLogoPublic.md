# WorkspaceMcpServerLogoPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasLogoOverride** | **bool** | An admin uploaded this server&#39;s logo. | 
**LogoTile** | **string** | An admin&#39;s choice: a tile only where the logo would fade out, always, or never. | 
**LogoTone** | **NullableString** | Where the logo needs a tile to stay visible: &#x60;dark&#x60; ink on a dark background, &#x60;light&#x60; ink on a light background, &#x60;tile&#x60; on every background. Null when it reads on both, or the tile is &#x60;never&#x60;. | 
**LogoUrl** | **NullableString** | Where this gateway serves the server&#39;s logo: an admin&#39;s upload, else the catalog&#39;s logo for its URL, else one found on its host. Null when none was found; the dashboard shows the server&#39;s initial. | 
**ServerId** | **string** |  | 

## Methods

### NewWorkspaceMcpServerLogoPublic

`func NewWorkspaceMcpServerLogoPublic(hasLogoOverride bool, logoTile string, logoTone NullableString, logoUrl NullableString, serverId string, ) *WorkspaceMcpServerLogoPublic`

NewWorkspaceMcpServerLogoPublic instantiates a new WorkspaceMcpServerLogoPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkspaceMcpServerLogoPublicWithDefaults

`func NewWorkspaceMcpServerLogoPublicWithDefaults() *WorkspaceMcpServerLogoPublic`

NewWorkspaceMcpServerLogoPublicWithDefaults instantiates a new WorkspaceMcpServerLogoPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasLogoOverride

`func (o *WorkspaceMcpServerLogoPublic) GetHasLogoOverride() bool`

GetHasLogoOverride returns the HasLogoOverride field if non-nil, zero value otherwise.

### GetHasLogoOverrideOk

`func (o *WorkspaceMcpServerLogoPublic) GetHasLogoOverrideOk() (*bool, bool)`

GetHasLogoOverrideOk returns a tuple with the HasLogoOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasLogoOverride

`func (o *WorkspaceMcpServerLogoPublic) SetHasLogoOverride(v bool)`

SetHasLogoOverride sets HasLogoOverride field to given value.


### GetLogoTile

`func (o *WorkspaceMcpServerLogoPublic) GetLogoTile() string`

GetLogoTile returns the LogoTile field if non-nil, zero value otherwise.

### GetLogoTileOk

`func (o *WorkspaceMcpServerLogoPublic) GetLogoTileOk() (*string, bool)`

GetLogoTileOk returns a tuple with the LogoTile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoTile

`func (o *WorkspaceMcpServerLogoPublic) SetLogoTile(v string)`

SetLogoTile sets LogoTile field to given value.


### GetLogoTone

`func (o *WorkspaceMcpServerLogoPublic) GetLogoTone() string`

GetLogoTone returns the LogoTone field if non-nil, zero value otherwise.

### GetLogoToneOk

`func (o *WorkspaceMcpServerLogoPublic) GetLogoToneOk() (*string, bool)`

GetLogoToneOk returns a tuple with the LogoTone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoTone

`func (o *WorkspaceMcpServerLogoPublic) SetLogoTone(v string)`

SetLogoTone sets LogoTone field to given value.


### SetLogoToneNil

`func (o *WorkspaceMcpServerLogoPublic) SetLogoToneNil(b bool)`

 SetLogoToneNil sets the value for LogoTone to be an explicit nil

### UnsetLogoTone
`func (o *WorkspaceMcpServerLogoPublic) UnsetLogoTone()`

UnsetLogoTone ensures that no value is present for LogoTone, not even an explicit nil
### GetLogoUrl

`func (o *WorkspaceMcpServerLogoPublic) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *WorkspaceMcpServerLogoPublic) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *WorkspaceMcpServerLogoPublic) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.


### SetLogoUrlNil

`func (o *WorkspaceMcpServerLogoPublic) SetLogoUrlNil(b bool)`

 SetLogoUrlNil sets the value for LogoUrl to be an explicit nil

### UnsetLogoUrl
`func (o *WorkspaceMcpServerLogoPublic) UnsetLogoUrl()`

UnsetLogoUrl ensures that no value is present for LogoUrl, not even an explicit nil
### GetServerId

`func (o *WorkspaceMcpServerLogoPublic) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *WorkspaceMcpServerLogoPublic) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *WorkspaceMcpServerLogoPublic) SetServerId(v string)`

SetServerId sets ServerId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


