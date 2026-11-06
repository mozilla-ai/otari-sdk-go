# FetchRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FetchToolName** | Pointer to **NullableString** | Configured fetch tool to run on, builtin_fetch included. Optional: the deployment&#39;s web_fetch_default_tool serves, else builtin_fetch. Ignored on POST /api/v1/fetch/{fetch_tool_name}. | [optional] 
**IncludeRaw** | Pointer to **bool** | Also return the provider&#39;s whole response, as raw. builtin_fetch calls no provider, so its raw holds only content_kind and redirect_count. | [optional] [default to false]
**MaxChars** | Pointer to **NullableInt32** | Cut the page text to this many characters. Unset, the tool&#39;s own max_chars option applies, else the provider&#39;s own limits. | [optional] 
**ProviderOptions** | Pointer to **map[string]interface{}** | Tags for cost attribution, recorded on the request&#39;s usage rows and filterable in the usage API: up to 16 string pairs, keys up to 64 characters and values up to 512. A null value is ignored. LiteLLM&#39;s nested &#x60;spend_logs_metadata&#x60; object is also read, and wins over a flat key of the same name; it is never forwarded to the provider. | [optional] 
**Url** | **string** | The page to fetch: one public http or https URL. | 
**User** | Pointer to **NullableString** | User ID for usage attribution | [optional] 

## Methods

### NewFetchRequest

`func NewFetchRequest(url string, ) *FetchRequest`

NewFetchRequest instantiates a new FetchRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFetchRequestWithDefaults

`func NewFetchRequestWithDefaults() *FetchRequest`

NewFetchRequestWithDefaults instantiates a new FetchRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFetchToolName

`func (o *FetchRequest) GetFetchToolName() string`

GetFetchToolName returns the FetchToolName field if non-nil, zero value otherwise.

### GetFetchToolNameOk

`func (o *FetchRequest) GetFetchToolNameOk() (*string, bool)`

GetFetchToolNameOk returns a tuple with the FetchToolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFetchToolName

`func (o *FetchRequest) SetFetchToolName(v string)`

SetFetchToolName sets FetchToolName field to given value.

### HasFetchToolName

`func (o *FetchRequest) HasFetchToolName() bool`

HasFetchToolName returns a boolean if a field has been set.

### SetFetchToolNameNil

`func (o *FetchRequest) SetFetchToolNameNil(b bool)`

 SetFetchToolNameNil sets the value for FetchToolName to be an explicit nil

### UnsetFetchToolName
`func (o *FetchRequest) UnsetFetchToolName()`

UnsetFetchToolName ensures that no value is present for FetchToolName, not even an explicit nil
### GetIncludeRaw

`func (o *FetchRequest) GetIncludeRaw() bool`

GetIncludeRaw returns the IncludeRaw field if non-nil, zero value otherwise.

### GetIncludeRawOk

`func (o *FetchRequest) GetIncludeRawOk() (*bool, bool)`

GetIncludeRawOk returns a tuple with the IncludeRaw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeRaw

`func (o *FetchRequest) SetIncludeRaw(v bool)`

SetIncludeRaw sets IncludeRaw field to given value.

### HasIncludeRaw

`func (o *FetchRequest) HasIncludeRaw() bool`

HasIncludeRaw returns a boolean if a field has been set.

### GetMaxChars

`func (o *FetchRequest) GetMaxChars() int32`

GetMaxChars returns the MaxChars field if non-nil, zero value otherwise.

### GetMaxCharsOk

`func (o *FetchRequest) GetMaxCharsOk() (*int32, bool)`

GetMaxCharsOk returns a tuple with the MaxChars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxChars

`func (o *FetchRequest) SetMaxChars(v int32)`

SetMaxChars sets MaxChars field to given value.

### HasMaxChars

`func (o *FetchRequest) HasMaxChars() bool`

HasMaxChars returns a boolean if a field has been set.

### SetMaxCharsNil

`func (o *FetchRequest) SetMaxCharsNil(b bool)`

 SetMaxCharsNil sets the value for MaxChars to be an explicit nil

### UnsetMaxChars
`func (o *FetchRequest) UnsetMaxChars()`

UnsetMaxChars ensures that no value is present for MaxChars, not even an explicit nil
### GetProviderOptions

`func (o *FetchRequest) GetProviderOptions() map[string]interface{}`

GetProviderOptions returns the ProviderOptions field if non-nil, zero value otherwise.

### GetProviderOptionsOk

`func (o *FetchRequest) GetProviderOptionsOk() (*map[string]interface{}, bool)`

GetProviderOptionsOk returns a tuple with the ProviderOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderOptions

`func (o *FetchRequest) SetProviderOptions(v map[string]interface{})`

SetProviderOptions sets ProviderOptions field to given value.

### HasProviderOptions

`func (o *FetchRequest) HasProviderOptions() bool`

HasProviderOptions returns a boolean if a field has been set.

### SetProviderOptionsNil

`func (o *FetchRequest) SetProviderOptionsNil(b bool)`

 SetProviderOptionsNil sets the value for ProviderOptions to be an explicit nil

### UnsetProviderOptions
`func (o *FetchRequest) UnsetProviderOptions()`

UnsetProviderOptions ensures that no value is present for ProviderOptions, not even an explicit nil
### GetUrl

`func (o *FetchRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *FetchRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *FetchRequest) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetUser

`func (o *FetchRequest) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *FetchRequest) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *FetchRequest) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *FetchRequest) HasUser() bool`

HasUser returns a boolean if a field has been set.

### SetUserNil

`func (o *FetchRequest) SetUserNil(b bool)`

 SetUserNil sets the value for User to be an explicit nil

### UnsetUser
`func (o *FetchRequest) UnsetUser()`

UnsetUser ensures that no value is present for User, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


