# AutomaticFallbackPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FellBack** | **int32** | Of &#x60;requests&#x60;, the successful ones a later offering served. | 
**Id** | **string** | The catalog ID a caller sends to get the fallback, e.g. &#x60;anthropic/claude-sonnet-5&#x60;. | 
**Name** | **string** |  | 
**Offerings** | [**[]FallbackOfferingPublic**](FallbackOfferingPublic.md) | What a request sent by catalog ID tries, first first. One entry means no backup. | 
**Pinned** | **int32** | Of &#x60;requests&#x60;, those that named an offering or a policy rather than the catalog ID, and so got no automatic fallback. | 
**PinnedProvider** | **NullableString** | The instance most pinned requests reached. | 
**Requests** | **int32** | Requests served in the window, each counted once however many attempts. | 

## Methods

### NewAutomaticFallbackPublic

`func NewAutomaticFallbackPublic(fellBack int32, id string, name string, offerings []FallbackOfferingPublic, pinned int32, pinnedProvider NullableString, requests int32, ) *AutomaticFallbackPublic`

NewAutomaticFallbackPublic instantiates a new AutomaticFallbackPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAutomaticFallbackPublicWithDefaults

`func NewAutomaticFallbackPublicWithDefaults() *AutomaticFallbackPublic`

NewAutomaticFallbackPublicWithDefaults instantiates a new AutomaticFallbackPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFellBack

`func (o *AutomaticFallbackPublic) GetFellBack() int32`

GetFellBack returns the FellBack field if non-nil, zero value otherwise.

### GetFellBackOk

`func (o *AutomaticFallbackPublic) GetFellBackOk() (*int32, bool)`

GetFellBackOk returns a tuple with the FellBack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFellBack

`func (o *AutomaticFallbackPublic) SetFellBack(v int32)`

SetFellBack sets FellBack field to given value.


### GetId

`func (o *AutomaticFallbackPublic) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AutomaticFallbackPublic) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AutomaticFallbackPublic) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *AutomaticFallbackPublic) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AutomaticFallbackPublic) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AutomaticFallbackPublic) SetName(v string)`

SetName sets Name field to given value.


### GetOfferings

`func (o *AutomaticFallbackPublic) GetOfferings() []FallbackOfferingPublic`

GetOfferings returns the Offerings field if non-nil, zero value otherwise.

### GetOfferingsOk

`func (o *AutomaticFallbackPublic) GetOfferingsOk() (*[]FallbackOfferingPublic, bool)`

GetOfferingsOk returns a tuple with the Offerings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferings

`func (o *AutomaticFallbackPublic) SetOfferings(v []FallbackOfferingPublic)`

SetOfferings sets Offerings field to given value.


### GetPinned

`func (o *AutomaticFallbackPublic) GetPinned() int32`

GetPinned returns the Pinned field if non-nil, zero value otherwise.

### GetPinnedOk

`func (o *AutomaticFallbackPublic) GetPinnedOk() (*int32, bool)`

GetPinnedOk returns a tuple with the Pinned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinned

`func (o *AutomaticFallbackPublic) SetPinned(v int32)`

SetPinned sets Pinned field to given value.


### GetPinnedProvider

`func (o *AutomaticFallbackPublic) GetPinnedProvider() string`

GetPinnedProvider returns the PinnedProvider field if non-nil, zero value otherwise.

### GetPinnedProviderOk

`func (o *AutomaticFallbackPublic) GetPinnedProviderOk() (*string, bool)`

GetPinnedProviderOk returns a tuple with the PinnedProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinnedProvider

`func (o *AutomaticFallbackPublic) SetPinnedProvider(v string)`

SetPinnedProvider sets PinnedProvider field to given value.


### SetPinnedProviderNil

`func (o *AutomaticFallbackPublic) SetPinnedProviderNil(b bool)`

 SetPinnedProviderNil sets the value for PinnedProvider to be an explicit nil

### UnsetPinnedProvider
`func (o *AutomaticFallbackPublic) UnsetPinnedProvider()`

UnsetPinnedProvider ensures that no value is present for PinnedProvider, not even an explicit nil
### GetRequests

`func (o *AutomaticFallbackPublic) GetRequests() int32`

GetRequests returns the Requests field if non-nil, zero value otherwise.

### GetRequestsOk

`func (o *AutomaticFallbackPublic) GetRequestsOk() (*int32, bool)`

GetRequestsOk returns a tuple with the Requests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequests

`func (o *AutomaticFallbackPublic) SetRequests(v int32)`

SetRequests sets Requests field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


