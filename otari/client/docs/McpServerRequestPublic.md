# McpServerRequestPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CatalogName** | **string** |  | 
**CatalogServer** | [**NullableMcpCatalogServerPublic**](McpCatalogServerPublic.md) | The catalog&#39;s entry as it is now; null once the catalog no longer lists the server | 
**CreatedAt** | **time.Time** |  | 
**DecidedAt** | **NullableTime** |  | 
**DecisionReason** | **NullableString** | An admin&#39;s reason for declining, when they gave one | 
**Id** | **string** |  | 
**Note** | **NullableString** |  | 
**ServerId** | **NullableString** | The stored server an approval produced | 
**Status** | **string** |  | 
**Title** | **string** | The catalog&#39;s title for the server when it was requested | 
**Url** | **string** | The server&#39;s endpoint when it was requested | 
**WorkspaceId** | **string** |  | 

## Methods

### NewMcpServerRequestPublic

`func NewMcpServerRequestPublic(catalogName string, catalogServer NullableMcpCatalogServerPublic, createdAt time.Time, decidedAt NullableTime, decisionReason NullableString, id string, note NullableString, serverId NullableString, status string, title string, url string, workspaceId string, ) *McpServerRequestPublic`

NewMcpServerRequestPublic instantiates a new McpServerRequestPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMcpServerRequestPublicWithDefaults

`func NewMcpServerRequestPublicWithDefaults() *McpServerRequestPublic`

NewMcpServerRequestPublicWithDefaults instantiates a new McpServerRequestPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCatalogName

`func (o *McpServerRequestPublic) GetCatalogName() string`

GetCatalogName returns the CatalogName field if non-nil, zero value otherwise.

### GetCatalogNameOk

`func (o *McpServerRequestPublic) GetCatalogNameOk() (*string, bool)`

GetCatalogNameOk returns a tuple with the CatalogName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogName

`func (o *McpServerRequestPublic) SetCatalogName(v string)`

SetCatalogName sets CatalogName field to given value.


### GetCatalogServer

`func (o *McpServerRequestPublic) GetCatalogServer() McpCatalogServerPublic`

GetCatalogServer returns the CatalogServer field if non-nil, zero value otherwise.

### GetCatalogServerOk

`func (o *McpServerRequestPublic) GetCatalogServerOk() (*McpCatalogServerPublic, bool)`

GetCatalogServerOk returns a tuple with the CatalogServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogServer

`func (o *McpServerRequestPublic) SetCatalogServer(v McpCatalogServerPublic)`

SetCatalogServer sets CatalogServer field to given value.


### SetCatalogServerNil

`func (o *McpServerRequestPublic) SetCatalogServerNil(b bool)`

 SetCatalogServerNil sets the value for CatalogServer to be an explicit nil

### UnsetCatalogServer
`func (o *McpServerRequestPublic) UnsetCatalogServer()`

UnsetCatalogServer ensures that no value is present for CatalogServer, not even an explicit nil
### GetCreatedAt

`func (o *McpServerRequestPublic) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *McpServerRequestPublic) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *McpServerRequestPublic) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetDecidedAt

`func (o *McpServerRequestPublic) GetDecidedAt() time.Time`

GetDecidedAt returns the DecidedAt field if non-nil, zero value otherwise.

### GetDecidedAtOk

`func (o *McpServerRequestPublic) GetDecidedAtOk() (*time.Time, bool)`

GetDecidedAtOk returns a tuple with the DecidedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecidedAt

`func (o *McpServerRequestPublic) SetDecidedAt(v time.Time)`

SetDecidedAt sets DecidedAt field to given value.


### SetDecidedAtNil

`func (o *McpServerRequestPublic) SetDecidedAtNil(b bool)`

 SetDecidedAtNil sets the value for DecidedAt to be an explicit nil

### UnsetDecidedAt
`func (o *McpServerRequestPublic) UnsetDecidedAt()`

UnsetDecidedAt ensures that no value is present for DecidedAt, not even an explicit nil
### GetDecisionReason

`func (o *McpServerRequestPublic) GetDecisionReason() string`

GetDecisionReason returns the DecisionReason field if non-nil, zero value otherwise.

### GetDecisionReasonOk

`func (o *McpServerRequestPublic) GetDecisionReasonOk() (*string, bool)`

GetDecisionReasonOk returns a tuple with the DecisionReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecisionReason

`func (o *McpServerRequestPublic) SetDecisionReason(v string)`

SetDecisionReason sets DecisionReason field to given value.


### SetDecisionReasonNil

`func (o *McpServerRequestPublic) SetDecisionReasonNil(b bool)`

 SetDecisionReasonNil sets the value for DecisionReason to be an explicit nil

### UnsetDecisionReason
`func (o *McpServerRequestPublic) UnsetDecisionReason()`

UnsetDecisionReason ensures that no value is present for DecisionReason, not even an explicit nil
### GetId

`func (o *McpServerRequestPublic) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *McpServerRequestPublic) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *McpServerRequestPublic) SetId(v string)`

SetId sets Id field to given value.


### GetNote

`func (o *McpServerRequestPublic) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *McpServerRequestPublic) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *McpServerRequestPublic) SetNote(v string)`

SetNote sets Note field to given value.


### SetNoteNil

`func (o *McpServerRequestPublic) SetNoteNil(b bool)`

 SetNoteNil sets the value for Note to be an explicit nil

### UnsetNote
`func (o *McpServerRequestPublic) UnsetNote()`

UnsetNote ensures that no value is present for Note, not even an explicit nil
### GetServerId

`func (o *McpServerRequestPublic) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *McpServerRequestPublic) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *McpServerRequestPublic) SetServerId(v string)`

SetServerId sets ServerId field to given value.


### SetServerIdNil

`func (o *McpServerRequestPublic) SetServerIdNil(b bool)`

 SetServerIdNil sets the value for ServerId to be an explicit nil

### UnsetServerId
`func (o *McpServerRequestPublic) UnsetServerId()`

UnsetServerId ensures that no value is present for ServerId, not even an explicit nil
### GetStatus

`func (o *McpServerRequestPublic) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *McpServerRequestPublic) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *McpServerRequestPublic) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTitle

`func (o *McpServerRequestPublic) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *McpServerRequestPublic) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *McpServerRequestPublic) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetUrl

`func (o *McpServerRequestPublic) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *McpServerRequestPublic) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *McpServerRequestPublic) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetWorkspaceId

`func (o *McpServerRequestPublic) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *McpServerRequestPublic) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *McpServerRequestPublic) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


