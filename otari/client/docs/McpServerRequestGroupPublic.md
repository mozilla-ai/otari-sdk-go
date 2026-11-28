# McpServerRequestGroupPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CatalogName** | **string** |  | 
**CatalogServer** | [**NullableMcpCatalogServerPublic**](McpCatalogServerPublic.md) | The catalog&#39;s entry as it is now; null once the catalog no longer lists the server | 
**FirstRequestedAt** | **time.Time** |  | 
**RequesterCount** | **int32** | How many pending requests ask for this server in all | 
**Requesters** | [**[]McpServerRequesterPublic**](McpServerRequesterPublic.md) | Who asked, earliest first, up to 5. Past that, the requesters list of this server has the rest | 
**Title** | **string** |  | 
**Url** | **string** |  | 

## Methods

### NewMcpServerRequestGroupPublic

`func NewMcpServerRequestGroupPublic(catalogName string, catalogServer NullableMcpCatalogServerPublic, firstRequestedAt time.Time, requesterCount int32, requesters []McpServerRequesterPublic, title string, url string, ) *McpServerRequestGroupPublic`

NewMcpServerRequestGroupPublic instantiates a new McpServerRequestGroupPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMcpServerRequestGroupPublicWithDefaults

`func NewMcpServerRequestGroupPublicWithDefaults() *McpServerRequestGroupPublic`

NewMcpServerRequestGroupPublicWithDefaults instantiates a new McpServerRequestGroupPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCatalogName

`func (o *McpServerRequestGroupPublic) GetCatalogName() string`

GetCatalogName returns the CatalogName field if non-nil, zero value otherwise.

### GetCatalogNameOk

`func (o *McpServerRequestGroupPublic) GetCatalogNameOk() (*string, bool)`

GetCatalogNameOk returns a tuple with the CatalogName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogName

`func (o *McpServerRequestGroupPublic) SetCatalogName(v string)`

SetCatalogName sets CatalogName field to given value.


### GetCatalogServer

`func (o *McpServerRequestGroupPublic) GetCatalogServer() McpCatalogServerPublic`

GetCatalogServer returns the CatalogServer field if non-nil, zero value otherwise.

### GetCatalogServerOk

`func (o *McpServerRequestGroupPublic) GetCatalogServerOk() (*McpCatalogServerPublic, bool)`

GetCatalogServerOk returns a tuple with the CatalogServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogServer

`func (o *McpServerRequestGroupPublic) SetCatalogServer(v McpCatalogServerPublic)`

SetCatalogServer sets CatalogServer field to given value.


### SetCatalogServerNil

`func (o *McpServerRequestGroupPublic) SetCatalogServerNil(b bool)`

 SetCatalogServerNil sets the value for CatalogServer to be an explicit nil

### UnsetCatalogServer
`func (o *McpServerRequestGroupPublic) UnsetCatalogServer()`

UnsetCatalogServer ensures that no value is present for CatalogServer, not even an explicit nil
### GetFirstRequestedAt

`func (o *McpServerRequestGroupPublic) GetFirstRequestedAt() time.Time`

GetFirstRequestedAt returns the FirstRequestedAt field if non-nil, zero value otherwise.

### GetFirstRequestedAtOk

`func (o *McpServerRequestGroupPublic) GetFirstRequestedAtOk() (*time.Time, bool)`

GetFirstRequestedAtOk returns a tuple with the FirstRequestedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstRequestedAt

`func (o *McpServerRequestGroupPublic) SetFirstRequestedAt(v time.Time)`

SetFirstRequestedAt sets FirstRequestedAt field to given value.


### GetRequesterCount

`func (o *McpServerRequestGroupPublic) GetRequesterCount() int32`

GetRequesterCount returns the RequesterCount field if non-nil, zero value otherwise.

### GetRequesterCountOk

`func (o *McpServerRequestGroupPublic) GetRequesterCountOk() (*int32, bool)`

GetRequesterCountOk returns a tuple with the RequesterCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesterCount

`func (o *McpServerRequestGroupPublic) SetRequesterCount(v int32)`

SetRequesterCount sets RequesterCount field to given value.


### GetRequesters

`func (o *McpServerRequestGroupPublic) GetRequesters() []McpServerRequesterPublic`

GetRequesters returns the Requesters field if non-nil, zero value otherwise.

### GetRequestersOk

`func (o *McpServerRequestGroupPublic) GetRequestersOk() (*[]McpServerRequesterPublic, bool)`

GetRequestersOk returns a tuple with the Requesters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequesters

`func (o *McpServerRequestGroupPublic) SetRequesters(v []McpServerRequesterPublic)`

SetRequesters sets Requesters field to given value.


### GetTitle

`func (o *McpServerRequestGroupPublic) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *McpServerRequestGroupPublic) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *McpServerRequestGroupPublic) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetUrl

`func (o *McpServerRequestGroupPublic) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *McpServerRequestGroupPublic) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *McpServerRequestGroupPublic) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


