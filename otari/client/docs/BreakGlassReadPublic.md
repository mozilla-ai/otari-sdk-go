# BreakGlassReadPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Kind** | **string** | &#39;organization&#39; for an owner of the organization, &#39;deployment&#39; for the deployment&#39;s operator. | 
**OrganizationId** | **string** |  | 
**ReadAt** | **time.Time** |  | 
**ReaderEmail** | **string** | Their email address when they opened it. | 
**ReaderUserId** | **NullableString** | Who opened it; null once that account is gone. | 
**Reason** | **string** |  | 
**Reference** | **NullableString** |  | 
**SessionUserId** | **NullableString** | Whose session it is. | 
**TraceId** | **string** |  | 
**WorkspaceId** | **string** |  | 

## Methods

### NewBreakGlassReadPublic

`func NewBreakGlassReadPublic(id string, kind string, organizationId string, readAt time.Time, readerEmail string, readerUserId NullableString, reason string, reference NullableString, sessionUserId NullableString, traceId string, workspaceId string, ) *BreakGlassReadPublic`

NewBreakGlassReadPublic instantiates a new BreakGlassReadPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBreakGlassReadPublicWithDefaults

`func NewBreakGlassReadPublicWithDefaults() *BreakGlassReadPublic`

NewBreakGlassReadPublicWithDefaults instantiates a new BreakGlassReadPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BreakGlassReadPublic) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BreakGlassReadPublic) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BreakGlassReadPublic) SetId(v string)`

SetId sets Id field to given value.


### GetKind

`func (o *BreakGlassReadPublic) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *BreakGlassReadPublic) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *BreakGlassReadPublic) SetKind(v string)`

SetKind sets Kind field to given value.


### GetOrganizationId

`func (o *BreakGlassReadPublic) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *BreakGlassReadPublic) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *BreakGlassReadPublic) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.


### GetReadAt

`func (o *BreakGlassReadPublic) GetReadAt() time.Time`

GetReadAt returns the ReadAt field if non-nil, zero value otherwise.

### GetReadAtOk

`func (o *BreakGlassReadPublic) GetReadAtOk() (*time.Time, bool)`

GetReadAtOk returns a tuple with the ReadAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadAt

`func (o *BreakGlassReadPublic) SetReadAt(v time.Time)`

SetReadAt sets ReadAt field to given value.


### GetReaderEmail

`func (o *BreakGlassReadPublic) GetReaderEmail() string`

GetReaderEmail returns the ReaderEmail field if non-nil, zero value otherwise.

### GetReaderEmailOk

`func (o *BreakGlassReadPublic) GetReaderEmailOk() (*string, bool)`

GetReaderEmailOk returns a tuple with the ReaderEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReaderEmail

`func (o *BreakGlassReadPublic) SetReaderEmail(v string)`

SetReaderEmail sets ReaderEmail field to given value.


### GetReaderUserId

`func (o *BreakGlassReadPublic) GetReaderUserId() string`

GetReaderUserId returns the ReaderUserId field if non-nil, zero value otherwise.

### GetReaderUserIdOk

`func (o *BreakGlassReadPublic) GetReaderUserIdOk() (*string, bool)`

GetReaderUserIdOk returns a tuple with the ReaderUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReaderUserId

`func (o *BreakGlassReadPublic) SetReaderUserId(v string)`

SetReaderUserId sets ReaderUserId field to given value.


### SetReaderUserIdNil

`func (o *BreakGlassReadPublic) SetReaderUserIdNil(b bool)`

 SetReaderUserIdNil sets the value for ReaderUserId to be an explicit nil

### UnsetReaderUserId
`func (o *BreakGlassReadPublic) UnsetReaderUserId()`

UnsetReaderUserId ensures that no value is present for ReaderUserId, not even an explicit nil
### GetReason

`func (o *BreakGlassReadPublic) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *BreakGlassReadPublic) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *BreakGlassReadPublic) SetReason(v string)`

SetReason sets Reason field to given value.


### GetReference

`func (o *BreakGlassReadPublic) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *BreakGlassReadPublic) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *BreakGlassReadPublic) SetReference(v string)`

SetReference sets Reference field to given value.


### SetReferenceNil

`func (o *BreakGlassReadPublic) SetReferenceNil(b bool)`

 SetReferenceNil sets the value for Reference to be an explicit nil

### UnsetReference
`func (o *BreakGlassReadPublic) UnsetReference()`

UnsetReference ensures that no value is present for Reference, not even an explicit nil
### GetSessionUserId

`func (o *BreakGlassReadPublic) GetSessionUserId() string`

GetSessionUserId returns the SessionUserId field if non-nil, zero value otherwise.

### GetSessionUserIdOk

`func (o *BreakGlassReadPublic) GetSessionUserIdOk() (*string, bool)`

GetSessionUserIdOk returns a tuple with the SessionUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionUserId

`func (o *BreakGlassReadPublic) SetSessionUserId(v string)`

SetSessionUserId sets SessionUserId field to given value.


### SetSessionUserIdNil

`func (o *BreakGlassReadPublic) SetSessionUserIdNil(b bool)`

 SetSessionUserIdNil sets the value for SessionUserId to be an explicit nil

### UnsetSessionUserId
`func (o *BreakGlassReadPublic) UnsetSessionUserId()`

UnsetSessionUserId ensures that no value is present for SessionUserId, not even an explicit nil
### GetTraceId

`func (o *BreakGlassReadPublic) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *BreakGlassReadPublic) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *BreakGlassReadPublic) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.


### GetWorkspaceId

`func (o *BreakGlassReadPublic) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *BreakGlassReadPublic) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *BreakGlassReadPublic) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


