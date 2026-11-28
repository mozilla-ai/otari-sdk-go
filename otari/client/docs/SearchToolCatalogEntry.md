# SearchToolCatalogEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Default** | **bool** | True for the deployment&#39;s default, the tool a direct search that names none runs on unless the caller&#39;s organization brings its own search key. | 
**Id** | **string** | The tool&#39;s name, to send as &#39;search_tool_name&#39; or in /api/v1/search/{tool}. | 
**Label** | **string** | A display name for the tool: its provider&#39;s name, capitalized. | 
**Object** | Pointer to **string** |  | [optional] [default to "search_tool"]
**PricePerCall** | **NullableFloat32** | USD the caller&#39;s organization is charged for one search on the tool when the provider reports no cost of its own: the organization&#39;s override, else the deployment&#39;s price list. Null when no row prices it. | 
**Provider** | **string** | The search provider the tool runs on. | 

## Methods

### NewSearchToolCatalogEntry

`func NewSearchToolCatalogEntry(default_ bool, id string, label string, pricePerCall NullableFloat32, provider string, ) *SearchToolCatalogEntry`

NewSearchToolCatalogEntry instantiates a new SearchToolCatalogEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchToolCatalogEntryWithDefaults

`func NewSearchToolCatalogEntryWithDefaults() *SearchToolCatalogEntry`

NewSearchToolCatalogEntryWithDefaults instantiates a new SearchToolCatalogEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefault

`func (o *SearchToolCatalogEntry) GetDefault() bool`

GetDefault returns the Default field if non-nil, zero value otherwise.

### GetDefaultOk

`func (o *SearchToolCatalogEntry) GetDefaultOk() (*bool, bool)`

GetDefaultOk returns a tuple with the Default field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefault

`func (o *SearchToolCatalogEntry) SetDefault(v bool)`

SetDefault sets Default field to given value.


### GetId

`func (o *SearchToolCatalogEntry) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SearchToolCatalogEntry) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SearchToolCatalogEntry) SetId(v string)`

SetId sets Id field to given value.


### GetLabel

`func (o *SearchToolCatalogEntry) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *SearchToolCatalogEntry) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *SearchToolCatalogEntry) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetObject

`func (o *SearchToolCatalogEntry) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *SearchToolCatalogEntry) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *SearchToolCatalogEntry) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *SearchToolCatalogEntry) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetPricePerCall

`func (o *SearchToolCatalogEntry) GetPricePerCall() float32`

GetPricePerCall returns the PricePerCall field if non-nil, zero value otherwise.

### GetPricePerCallOk

`func (o *SearchToolCatalogEntry) GetPricePerCallOk() (*float32, bool)`

GetPricePerCallOk returns a tuple with the PricePerCall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricePerCall

`func (o *SearchToolCatalogEntry) SetPricePerCall(v float32)`

SetPricePerCall sets PricePerCall field to given value.


### SetPricePerCallNil

`func (o *SearchToolCatalogEntry) SetPricePerCallNil(b bool)`

 SetPricePerCallNil sets the value for PricePerCall to be an explicit nil

### UnsetPricePerCall
`func (o *SearchToolCatalogEntry) UnsetPricePerCall()`

UnsetPricePerCall ensures that no value is present for PricePerCall, not even an explicit nil
### GetProvider

`func (o *SearchToolCatalogEntry) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *SearchToolCatalogEntry) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *SearchToolCatalogEntry) SetProvider(v string)`

SetProvider sets Provider field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


