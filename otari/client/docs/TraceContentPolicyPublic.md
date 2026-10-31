# TraceContentPolicyPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentAvailable** | **bool** | Whether this deployment lets a workspace keep what its agent sessions were about. A workspace keeps it only once its managers turn it on as well. | 
**PlatformBreakGlass** | **bool** | Whether the deployment&#39;s operator may open any organization&#39;s session content without its owner&#39;s share, recorded with a reason and a reference. | 
**RetentionDays** | **int32** | Days after its last activity that a session, with what it was about and its shares, is deleted. | 
**SessionMaxAgeDays** | **int32** | Days after it started that a session is deleted, however active it still is. | 

## Methods

### NewTraceContentPolicyPublic

`func NewTraceContentPolicyPublic(contentAvailable bool, platformBreakGlass bool, retentionDays int32, sessionMaxAgeDays int32, ) *TraceContentPolicyPublic`

NewTraceContentPolicyPublic instantiates a new TraceContentPolicyPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTraceContentPolicyPublicWithDefaults

`func NewTraceContentPolicyPublicWithDefaults() *TraceContentPolicyPublic`

NewTraceContentPolicyPublicWithDefaults instantiates a new TraceContentPolicyPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentAvailable

`func (o *TraceContentPolicyPublic) GetContentAvailable() bool`

GetContentAvailable returns the ContentAvailable field if non-nil, zero value otherwise.

### GetContentAvailableOk

`func (o *TraceContentPolicyPublic) GetContentAvailableOk() (*bool, bool)`

GetContentAvailableOk returns a tuple with the ContentAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentAvailable

`func (o *TraceContentPolicyPublic) SetContentAvailable(v bool)`

SetContentAvailable sets ContentAvailable field to given value.


### GetPlatformBreakGlass

`func (o *TraceContentPolicyPublic) GetPlatformBreakGlass() bool`

GetPlatformBreakGlass returns the PlatformBreakGlass field if non-nil, zero value otherwise.

### GetPlatformBreakGlassOk

`func (o *TraceContentPolicyPublic) GetPlatformBreakGlassOk() (*bool, bool)`

GetPlatformBreakGlassOk returns a tuple with the PlatformBreakGlass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformBreakGlass

`func (o *TraceContentPolicyPublic) SetPlatformBreakGlass(v bool)`

SetPlatformBreakGlass sets PlatformBreakGlass field to given value.


### GetRetentionDays

`func (o *TraceContentPolicyPublic) GetRetentionDays() int32`

GetRetentionDays returns the RetentionDays field if non-nil, zero value otherwise.

### GetRetentionDaysOk

`func (o *TraceContentPolicyPublic) GetRetentionDaysOk() (*int32, bool)`

GetRetentionDaysOk returns a tuple with the RetentionDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionDays

`func (o *TraceContentPolicyPublic) SetRetentionDays(v int32)`

SetRetentionDays sets RetentionDays field to given value.


### GetSessionMaxAgeDays

`func (o *TraceContentPolicyPublic) GetSessionMaxAgeDays() int32`

GetSessionMaxAgeDays returns the SessionMaxAgeDays field if non-nil, zero value otherwise.

### GetSessionMaxAgeDaysOk

`func (o *TraceContentPolicyPublic) GetSessionMaxAgeDaysOk() (*int32, bool)`

GetSessionMaxAgeDaysOk returns a tuple with the SessionMaxAgeDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionMaxAgeDays

`func (o *TraceContentPolicyPublic) SetSessionMaxAgeDays(v int32)`

SetSessionMaxAgeDays sets SessionMaxAgeDays field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


