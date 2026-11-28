# McpServerRequestDecline

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CatalogName** | **string** |  | 
**Reason** | Pointer to **NullableString** | Shown to each requester | [optional] 

## Methods

### NewMcpServerRequestDecline

`func NewMcpServerRequestDecline(catalogName string, ) *McpServerRequestDecline`

NewMcpServerRequestDecline instantiates a new McpServerRequestDecline object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMcpServerRequestDeclineWithDefaults

`func NewMcpServerRequestDeclineWithDefaults() *McpServerRequestDecline`

NewMcpServerRequestDeclineWithDefaults instantiates a new McpServerRequestDecline object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCatalogName

`func (o *McpServerRequestDecline) GetCatalogName() string`

GetCatalogName returns the CatalogName field if non-nil, zero value otherwise.

### GetCatalogNameOk

`func (o *McpServerRequestDecline) GetCatalogNameOk() (*string, bool)`

GetCatalogNameOk returns a tuple with the CatalogName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogName

`func (o *McpServerRequestDecline) SetCatalogName(v string)`

SetCatalogName sets CatalogName field to given value.


### GetReason

`func (o *McpServerRequestDecline) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *McpServerRequestDecline) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *McpServerRequestDecline) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *McpServerRequestDecline) HasReason() bool`

HasReason returns a boolean if a field has been set.

### SetReasonNil

`func (o *McpServerRequestDecline) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *McpServerRequestDecline) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


