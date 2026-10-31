# ContentShareListPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasMore** | Pointer to **bool** | More shares follow this page. | [optional] [default to false]
**Items** | [**[]ContentSharePublic**](ContentSharePublic.md) |  | 

## Methods

### NewContentShareListPublic

`func NewContentShareListPublic(items []ContentSharePublic, ) *ContentShareListPublic`

NewContentShareListPublic instantiates a new ContentShareListPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContentShareListPublicWithDefaults

`func NewContentShareListPublicWithDefaults() *ContentShareListPublic`

NewContentShareListPublicWithDefaults instantiates a new ContentShareListPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasMore

`func (o *ContentShareListPublic) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *ContentShareListPublic) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *ContentShareListPublic) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.

### HasHasMore

`func (o *ContentShareListPublic) HasHasMore() bool`

HasHasMore returns a boolean if a field has been set.

### GetItems

`func (o *ContentShareListPublic) GetItems() []ContentSharePublic`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ContentShareListPublic) GetItemsOk() (*[]ContentSharePublic, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ContentShareListPublic) SetItems(v []ContentSharePublic)`

SetItems sets Items field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


