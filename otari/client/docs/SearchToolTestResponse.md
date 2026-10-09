# SearchToolTestResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Characters** | Pointer to **NullableInt32** | For a fetch that worked: how many characters of page text came back. | [optional] 
**Error** | Pointer to **NullableString** | When not ok, the error&#39;s tag: timeout, network, http_error, invalid_response, or the provider&#39;s own. | [optional] 
**Hits** | Pointer to **NullableInt32** | For a search that worked: how many hits came back. | [optional] 
**Ok** | **bool** | Whether the provider answered the call without an error. | 

## Methods

### NewSearchToolTestResponse

`func NewSearchToolTestResponse(ok bool, ) *SearchToolTestResponse`

NewSearchToolTestResponse instantiates a new SearchToolTestResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchToolTestResponseWithDefaults

`func NewSearchToolTestResponseWithDefaults() *SearchToolTestResponse`

NewSearchToolTestResponseWithDefaults instantiates a new SearchToolTestResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCharacters

`func (o *SearchToolTestResponse) GetCharacters() int32`

GetCharacters returns the Characters field if non-nil, zero value otherwise.

### GetCharactersOk

`func (o *SearchToolTestResponse) GetCharactersOk() (*int32, bool)`

GetCharactersOk returns a tuple with the Characters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharacters

`func (o *SearchToolTestResponse) SetCharacters(v int32)`

SetCharacters sets Characters field to given value.

### HasCharacters

`func (o *SearchToolTestResponse) HasCharacters() bool`

HasCharacters returns a boolean if a field has been set.

### SetCharactersNil

`func (o *SearchToolTestResponse) SetCharactersNil(b bool)`

 SetCharactersNil sets the value for Characters to be an explicit nil

### UnsetCharacters
`func (o *SearchToolTestResponse) UnsetCharacters()`

UnsetCharacters ensures that no value is present for Characters, not even an explicit nil
### GetError

`func (o *SearchToolTestResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *SearchToolTestResponse) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *SearchToolTestResponse) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *SearchToolTestResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### SetErrorNil

`func (o *SearchToolTestResponse) SetErrorNil(b bool)`

 SetErrorNil sets the value for Error to be an explicit nil

### UnsetError
`func (o *SearchToolTestResponse) UnsetError()`

UnsetError ensures that no value is present for Error, not even an explicit nil
### GetHits

`func (o *SearchToolTestResponse) GetHits() int32`

GetHits returns the Hits field if non-nil, zero value otherwise.

### GetHitsOk

`func (o *SearchToolTestResponse) GetHitsOk() (*int32, bool)`

GetHitsOk returns a tuple with the Hits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHits

`func (o *SearchToolTestResponse) SetHits(v int32)`

SetHits sets Hits field to given value.

### HasHits

`func (o *SearchToolTestResponse) HasHits() bool`

HasHits returns a boolean if a field has been set.

### SetHitsNil

`func (o *SearchToolTestResponse) SetHitsNil(b bool)`

 SetHitsNil sets the value for Hits to be an explicit nil

### UnsetHits
`func (o *SearchToolTestResponse) UnsetHits()`

UnsetHits ensures that no value is present for Hits, not even an explicit nil
### GetOk

`func (o *SearchToolTestResponse) GetOk() bool`

GetOk returns the Ok field if non-nil, zero value otherwise.

### GetOkOk

`func (o *SearchToolTestResponse) GetOkOk() (*bool, bool)`

GetOkOk returns a tuple with the Ok field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOk

`func (o *SearchToolTestResponse) SetOk(v bool)`

SetOk sets Ok field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


