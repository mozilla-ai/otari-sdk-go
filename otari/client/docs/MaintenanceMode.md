# MaintenanceMode

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | When true, POST /v1/auth/session refuses every credential with 503 so nobody starts a new dashboard session during a redeploy. Sessions already issued keep working, and the management API and the data plane are unaffected: a caller presenting the master key or an API key through the header is never frozen out. | 

## Methods

### NewMaintenanceMode

`func NewMaintenanceMode(enabled bool, ) *MaintenanceMode`

NewMaintenanceMode instantiates a new MaintenanceMode object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMaintenanceModeWithDefaults

`func NewMaintenanceModeWithDefaults() *MaintenanceMode`

NewMaintenanceModeWithDefaults instantiates a new MaintenanceMode object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *MaintenanceMode) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *MaintenanceMode) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *MaintenanceMode) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


