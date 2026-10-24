# TraceListPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasMore** | **bool** |  | 
**Items** | [**[]TraceSummaryPublic**](TraceSummaryPublic.md) |  | 

## Methods

### NewTraceListPublic

`func NewTraceListPublic(hasMore bool, items []TraceSummaryPublic, ) *TraceListPublic`

NewTraceListPublic instantiates a new TraceListPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTraceListPublicWithDefaults

`func NewTraceListPublicWithDefaults() *TraceListPublic`

NewTraceListPublicWithDefaults instantiates a new TraceListPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasMore

`func (o *TraceListPublic) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *TraceListPublic) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *TraceListPublic) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.


### GetItems

`func (o *TraceListPublic) GetItems() []TraceSummaryPublic`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *TraceListPublic) GetItemsOk() (*[]TraceSummaryPublic, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *TraceListPublic) SetItems(v []TraceSummaryPublic)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


