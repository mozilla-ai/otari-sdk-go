# MRWebSearchResultBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EncryptedContent** | **string** |  | 
**PageAge** | Pointer to **NullableString** | Filter models by provider name | [optional] 
**Title** | **string** |  | 
**Type** | **string** |  | 
**Url** | **string** |  | 

## Methods

### NewMRWebSearchResultBlock

`func NewMRWebSearchResultBlock(encryptedContent string, title string, type_ string, url string, ) *MRWebSearchResultBlock`

NewMRWebSearchResultBlock instantiates a new MRWebSearchResultBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMRWebSearchResultBlockWithDefaults

`func NewMRWebSearchResultBlockWithDefaults() *MRWebSearchResultBlock`

NewMRWebSearchResultBlockWithDefaults instantiates a new MRWebSearchResultBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEncryptedContent

`func (o *MRWebSearchResultBlock) GetEncryptedContent() string`

GetEncryptedContent returns the EncryptedContent field if non-nil, zero value otherwise.

### GetEncryptedContentOk

`func (o *MRWebSearchResultBlock) GetEncryptedContentOk() (*string, bool)`

GetEncryptedContentOk returns a tuple with the EncryptedContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptedContent

`func (o *MRWebSearchResultBlock) SetEncryptedContent(v string)`

SetEncryptedContent sets EncryptedContent field to given value.


### GetPageAge

`func (o *MRWebSearchResultBlock) GetPageAge() string`

GetPageAge returns the PageAge field if non-nil, zero value otherwise.

### GetPageAgeOk

`func (o *MRWebSearchResultBlock) GetPageAgeOk() (*string, bool)`

GetPageAgeOk returns a tuple with the PageAge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageAge

`func (o *MRWebSearchResultBlock) SetPageAge(v string)`

SetPageAge sets PageAge field to given value.

### HasPageAge

`func (o *MRWebSearchResultBlock) HasPageAge() bool`

HasPageAge returns a boolean if a field has been set.

### SetPageAgeNil

`func (o *MRWebSearchResultBlock) SetPageAgeNil(b bool)`

 SetPageAgeNil sets the value for PageAge to be an explicit nil

### UnsetPageAge
`func (o *MRWebSearchResultBlock) UnsetPageAge()`

UnsetPageAge ensures that no value is present for PageAge, not even an explicit nil
### GetTitle

`func (o *MRWebSearchResultBlock) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *MRWebSearchResultBlock) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *MRWebSearchResultBlock) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetType

`func (o *MRWebSearchResultBlock) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MRWebSearchResultBlock) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MRWebSearchResultBlock) SetType(v string)`

SetType sets Type field to given value.


### GetUrl

`func (o *MRWebSearchResultBlock) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *MRWebSearchResultBlock) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *MRWebSearchResultBlock) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


