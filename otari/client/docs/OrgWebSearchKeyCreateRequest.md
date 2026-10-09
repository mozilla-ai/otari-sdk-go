# OrgWebSearchKeyCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiKey** | **string** |  | 
**Name** | **string** | How the organization tells this key apart from its others. | 
**Provider** | **string** | The search provider the key is for: tavily or brave. | 

## Methods

### NewOrgWebSearchKeyCreateRequest

`func NewOrgWebSearchKeyCreateRequest(apiKey string, name string, provider string, ) *OrgWebSearchKeyCreateRequest`

NewOrgWebSearchKeyCreateRequest instantiates a new OrgWebSearchKeyCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrgWebSearchKeyCreateRequestWithDefaults

`func NewOrgWebSearchKeyCreateRequestWithDefaults() *OrgWebSearchKeyCreateRequest`

NewOrgWebSearchKeyCreateRequestWithDefaults instantiates a new OrgWebSearchKeyCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiKey

`func (o *OrgWebSearchKeyCreateRequest) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *OrgWebSearchKeyCreateRequest) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *OrgWebSearchKeyCreateRequest) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.


### GetName

`func (o *OrgWebSearchKeyCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrgWebSearchKeyCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrgWebSearchKeyCreateRequest) SetName(v string)`

SetName sets Name field to given value.


### GetProvider

`func (o *OrgWebSearchKeyCreateRequest) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *OrgWebSearchKeyCreateRequest) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *OrgWebSearchKeyCreateRequest) SetProvider(v string)`

SetProvider sets Provider field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


