# McpCatalogSyncStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | **NullableString** | Why the last sync was partial or failed. | 
**Fetched** | **int32** | Registry entries read. | 
**FinishedAt** | **NullableTime** |  | 
**IsStale** | **bool** | The catalog is empty, or no sync has read the whole registry into it in the last day. | 
**Kept** | **int32** | Entries stored: remote servers Otari can connect to. | 
**LastCompletedAt** | **NullableTime** | When a sync last read the whole registry. | 
**Mode** | **string** | The deployment&#39;s &#x60;mcp_catalog_sync&#x60; setting. | 
**NextRequestAt** | **NullableTime** | Before this time a workspace admin&#39;s sync request is refused with a 429, because a sync ended recently. Null when one may be asked for now. An operator is not held to it. | 
**NextSyncAt** | **NullableTime** | When the next scheduled sync is due, if the mode is scheduled. | 
**Outcome** | **NullableString** | How the last sync ended, null before the first. &#x60;partial&#x60; kept the pages it read and removed nothing; the next sync resumes where it stopped. | 
**Pages** | **int32** |  | 
**Removed** | **int32** | Servers a complete sync deleted because the registry no longer lists them. | 
**Requested** | **bool** | A sync was asked for and no worker has started it yet. | 
**Running** | **bool** |  | 
**ServerCount** | **int32** | Servers in the catalog now. | 
**StartedAt** | **NullableTime** |  | 

## Methods

### NewMcpCatalogSyncStatus

`func NewMcpCatalogSyncStatus(error_ NullableString, fetched int32, finishedAt NullableTime, isStale bool, kept int32, lastCompletedAt NullableTime, mode string, nextRequestAt NullableTime, nextSyncAt NullableTime, outcome NullableString, pages int32, removed int32, requested bool, running bool, serverCount int32, startedAt NullableTime, ) *McpCatalogSyncStatus`

NewMcpCatalogSyncStatus instantiates a new McpCatalogSyncStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMcpCatalogSyncStatusWithDefaults

`func NewMcpCatalogSyncStatusWithDefaults() *McpCatalogSyncStatus`

NewMcpCatalogSyncStatusWithDefaults instantiates a new McpCatalogSyncStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *McpCatalogSyncStatus) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *McpCatalogSyncStatus) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *McpCatalogSyncStatus) SetError(v string)`

SetError sets Error field to given value.


### SetErrorNil

`func (o *McpCatalogSyncStatus) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *McpCatalogSyncStatus) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetFetched

`func (o *McpCatalogSyncStatus) GetFetched() int32`

GetFetched returns the Fetched field if non-nil, zero value otherwise.

### GetFetchedOk

`func (o *McpCatalogSyncStatus) GetFetchedOk() (*int32, bool)`

GetFetchedOk returns a tuple with the Fetched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFetched

`func (o *McpCatalogSyncStatus) SetFetched(v int32)`

SetFetched sets Fetched field to given value.


### GetFinishedAt

`func (o *McpCatalogSyncStatus) GetFinishedAt() time.Time`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *McpCatalogSyncStatus) GetFinishedAtOk() (*time.Time, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *McpCatalogSyncStatus) SetFinishedAt(v time.Time)`

SetFinishedAt sets FinishedAt field to given value.


### SetFinishedAtNil

`func (o *McpCatalogSyncStatus) SetFinishedAtNil(b bool)`

 SetFinishedAtNil sets the value for FinishedAt to be an explicit nil

### UnsetFinishedAt
`func (o *McpCatalogSyncStatus) UnsetFinishedAt()`

UnsetFinishedAt ensures that no value is present for FinishedAt, not even an explicit nil
### GetIsStale

`func (o *McpCatalogSyncStatus) GetIsStale() bool`

GetIsStale returns the IsStale field if non-nil, zero value otherwise.

### GetIsStaleOk

`func (o *McpCatalogSyncStatus) GetIsStaleOk() (*bool, bool)`

GetIsStaleOk returns a tuple with the IsStale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsStale

`func (o *McpCatalogSyncStatus) SetIsStale(v bool)`

SetIsStale sets IsStale field to given value.


### GetKept

`func (o *McpCatalogSyncStatus) GetKept() int32`

GetKept returns the Kept field if non-nil, zero value otherwise.

### GetKeptOk

`func (o *McpCatalogSyncStatus) GetKeptOk() (*int32, bool)`

GetKeptOk returns a tuple with the Kept field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKept

`func (o *McpCatalogSyncStatus) SetKept(v int32)`

SetKept sets Kept field to given value.


### GetLastCompletedAt

`func (o *McpCatalogSyncStatus) GetLastCompletedAt() time.Time`

GetLastCompletedAt returns the LastCompletedAt field if non-nil, zero value otherwise.

### GetLastCompletedAtOk

`func (o *McpCatalogSyncStatus) GetLastCompletedAtOk() (*time.Time, bool)`

GetLastCompletedAtOk returns a tuple with the LastCompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCompletedAt

`func (o *McpCatalogSyncStatus) SetLastCompletedAt(v time.Time)`

SetLastCompletedAt sets LastCompletedAt field to given value.


### SetLastCompletedAtNil

`func (o *McpCatalogSyncStatus) SetLastCompletedAtNil(b bool)`

 SetLastCompletedAtNil sets the value for LastCompletedAt to be an explicit nil

### UnsetLastCompletedAt
`func (o *McpCatalogSyncStatus) UnsetLastCompletedAt()`

UnsetLastCompletedAt ensures that no value is present for LastCompletedAt, not even an explicit nil
### GetMode

`func (o *McpCatalogSyncStatus) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *McpCatalogSyncStatus) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *McpCatalogSyncStatus) SetMode(v string)`

SetMode sets Mode field to given value.


### GetNextRequestAt

`func (o *McpCatalogSyncStatus) GetNextRequestAt() time.Time`

GetNextRequestAt returns the NextRequestAt field if non-nil, zero value otherwise.

### GetNextRequestAtOk

`func (o *McpCatalogSyncStatus) GetNextRequestAtOk() (*time.Time, bool)`

GetNextRequestAtOk returns a tuple with the NextRequestAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextRequestAt

`func (o *McpCatalogSyncStatus) SetNextRequestAt(v time.Time)`

SetNextRequestAt sets NextRequestAt field to given value.


### SetNextRequestAtNil

`func (o *McpCatalogSyncStatus) SetNextRequestAtNil(b bool)`

 SetNextRequestAtNil sets the value for NextRequestAt to be an explicit nil

### UnsetNextRequestAt
`func (o *McpCatalogSyncStatus) UnsetNextRequestAt()`

UnsetNextRequestAt ensures that no value is present for NextRequestAt, not even an explicit nil
### GetNextSyncAt

`func (o *McpCatalogSyncStatus) GetNextSyncAt() time.Time`

GetNextSyncAt returns the NextSyncAt field if non-nil, zero value otherwise.

### GetNextSyncAtOk

`func (o *McpCatalogSyncStatus) GetNextSyncAtOk() (*time.Time, bool)`

GetNextSyncAtOk returns a tuple with the NextSyncAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextSyncAt

`func (o *McpCatalogSyncStatus) SetNextSyncAt(v time.Time)`

SetNextSyncAt sets NextSyncAt field to given value.


### SetNextSyncAtNil

`func (o *McpCatalogSyncStatus) SetNextSyncAtNil(b bool)`

 SetNextSyncAtNil sets the value for NextSyncAt to be an explicit nil

### UnsetNextSyncAt
`func (o *McpCatalogSyncStatus) UnsetNextSyncAt()`

UnsetNextSyncAt ensures that no value is present for NextSyncAt, not even an explicit nil
### GetOutcome

`func (o *McpCatalogSyncStatus) GetOutcome() string`

GetOutcome returns the Outcome field if non-nil, zero value otherwise.

### GetOutcomeOk

`func (o *McpCatalogSyncStatus) GetOutcomeOk() (*string, bool)`

GetOutcomeOk returns a tuple with the Outcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcome

`func (o *McpCatalogSyncStatus) SetOutcome(v string)`

SetOutcome sets Outcome field to given value.


### SetOutcomeNil

`func (o *McpCatalogSyncStatus) SetOutcomeNil(b bool)`

 SetOutcomeNil sets the value for Outcome to be an explicit nil

### UnsetOutcome
`func (o *McpCatalogSyncStatus) UnsetOutcome()`

UnsetOutcome ensures that no value is present for Outcome, not even an explicit nil
### GetPages

`func (o *McpCatalogSyncStatus) GetPages() int32`

GetPages returns the Pages field if non-nil, zero value otherwise.

### GetPagesOk

`func (o *McpCatalogSyncStatus) GetPagesOk() (*int32, bool)`

GetPagesOk returns a tuple with the Pages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPages

`func (o *McpCatalogSyncStatus) SetPages(v int32)`

SetPages sets Pages field to given value.


### GetRemoved

`func (o *McpCatalogSyncStatus) GetRemoved() int32`

GetRemoved returns the Removed field if non-nil, zero value otherwise.

### GetRemovedOk

`func (o *McpCatalogSyncStatus) GetRemovedOk() (*int32, bool)`

GetRemovedOk returns a tuple with the Removed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoved

`func (o *McpCatalogSyncStatus) SetRemoved(v int32)`

SetRemoved sets Removed field to given value.


### GetRequested

`func (o *McpCatalogSyncStatus) GetRequested() bool`

GetRequested returns the Requested field if non-nil, zero value otherwise.

### GetRequestedOk

`func (o *McpCatalogSyncStatus) GetRequestedOk() (*bool, bool)`

GetRequestedOk returns a tuple with the Requested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequested

`func (o *McpCatalogSyncStatus) SetRequested(v bool)`

SetRequested sets Requested field to given value.


### GetRunning

`func (o *McpCatalogSyncStatus) GetRunning() bool`

GetRunning returns the Running field if non-nil, zero value otherwise.

### GetRunningOk

`func (o *McpCatalogSyncStatus) GetRunningOk() (*bool, bool)`

GetRunningOk returns a tuple with the Running field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunning

`func (o *McpCatalogSyncStatus) SetRunning(v bool)`

SetRunning sets Running field to given value.


### GetServerCount

`func (o *McpCatalogSyncStatus) GetServerCount() int32`

GetServerCount returns the ServerCount field if non-nil, zero value otherwise.

### GetServerCountOk

`func (o *McpCatalogSyncStatus) GetServerCountOk() (*int32, bool)`

GetServerCountOk returns a tuple with the ServerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerCount

`func (o *McpCatalogSyncStatus) SetServerCount(v int32)`

SetServerCount sets ServerCount field to given value.


### GetStartedAt

`func (o *McpCatalogSyncStatus) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *McpCatalogSyncStatus) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *McpCatalogSyncStatus) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.


### SetStartedAtNil

`func (o *McpCatalogSyncStatus) SetStartedAtNil(b bool)`

 SetStartedAtNil sets the value for StartedAt to be an explicit nil

### UnsetStartedAt
`func (o *McpCatalogSyncStatus) UnsetStartedAt()`

UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


