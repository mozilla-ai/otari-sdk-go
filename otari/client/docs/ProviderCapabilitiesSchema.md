# ProviderCapabilitiesSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Audio** | **bool** |  | 
**Embeddings** | **bool** |  | 
**ImageGeneration** | **bool** |  | 
**ListModels** | **bool** |  | 
**Moderation** | **bool** |  | 
**Pdf** | **bool** |  | 
**Reasoning** | **bool** |  | 
**Rerank** | **bool** |  | 
**ResponsesApi** | **bool** |  | 
**Streaming** | **bool** |  | 
**Vision** | **bool** |  | 

## Methods

### NewProviderCapabilitiesSchema

`func NewProviderCapabilitiesSchema(audio bool, embeddings bool, imageGeneration bool, listModels bool, moderation bool, pdf bool, reasoning bool, rerank bool, responsesApi bool, streaming bool, vision bool, ) *ProviderCapabilitiesSchema`

NewProviderCapabilitiesSchema instantiates a new ProviderCapabilitiesSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderCapabilitiesSchemaWithDefaults

`func NewProviderCapabilitiesSchemaWithDefaults() *ProviderCapabilitiesSchema`

NewProviderCapabilitiesSchemaWithDefaults instantiates a new ProviderCapabilitiesSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAudio

`func (o *ProviderCapabilitiesSchema) GetAudio() bool`

GetAudio returns the Audio field if non-nil, zero value otherwise.

### GetAudioOk

`func (o *ProviderCapabilitiesSchema) GetAudioOk() (*bool, bool)`

GetAudioOk returns a tuple with the Audio field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudio

`func (o *ProviderCapabilitiesSchema) SetAudio(v bool)`

SetAudio sets Audio field to given value.


### GetEmbeddings

`func (o *ProviderCapabilitiesSchema) GetEmbeddings() bool`

GetEmbeddings returns the Embeddings field if non-nil, zero value otherwise.

### GetEmbeddingsOk

`func (o *ProviderCapabilitiesSchema) GetEmbeddingsOk() (*bool, bool)`

GetEmbeddingsOk returns a tuple with the Embeddings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbeddings

`func (o *ProviderCapabilitiesSchema) SetEmbeddings(v bool)`

SetEmbeddings sets Embeddings field to given value.


### GetImageGeneration

`func (o *ProviderCapabilitiesSchema) GetImageGeneration() bool`

GetImageGeneration returns the ImageGeneration field if non-nil, zero value otherwise.

### GetImageGenerationOk

`func (o *ProviderCapabilitiesSchema) GetImageGenerationOk() (*bool, bool)`

GetImageGenerationOk returns a tuple with the ImageGeneration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageGeneration

`func (o *ProviderCapabilitiesSchema) SetImageGeneration(v bool)`

SetImageGeneration sets ImageGeneration field to given value.


### GetListModels

`func (o *ProviderCapabilitiesSchema) GetListModels() bool`

GetListModels returns the ListModels field if non-nil, zero value otherwise.

### GetListModelsOk

`func (o *ProviderCapabilitiesSchema) GetListModelsOk() (*bool, bool)`

GetListModelsOk returns a tuple with the ListModels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListModels

`func (o *ProviderCapabilitiesSchema) SetListModels(v bool)`

SetListModels sets ListModels field to given value.


### GetModeration

`func (o *ProviderCapabilitiesSchema) GetModeration() bool`

GetModeration returns the Moderation field if non-nil, zero value otherwise.

### GetModerationOk

`func (o *ProviderCapabilitiesSchema) GetModerationOk() (*bool, bool)`

GetModerationOk returns a tuple with the Moderation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModeration

`func (o *ProviderCapabilitiesSchema) SetModeration(v bool)`

SetModeration sets Moderation field to given value.


### GetPdf

`func (o *ProviderCapabilitiesSchema) GetPdf() bool`

GetPdf returns the Pdf field if non-nil, zero value otherwise.

### GetPdfOk

`func (o *ProviderCapabilitiesSchema) GetPdfOk() (*bool, bool)`

GetPdfOk returns a tuple with the Pdf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPdf

`func (o *ProviderCapabilitiesSchema) SetPdf(v bool)`

SetPdf sets Pdf field to given value.


### GetReasoning

`func (o *ProviderCapabilitiesSchema) GetReasoning() bool`

GetReasoning returns the Reasoning field if non-nil, zero value otherwise.

### GetReasoningOk

`func (o *ProviderCapabilitiesSchema) GetReasoningOk() (*bool, bool)`

GetReasoningOk returns a tuple with the Reasoning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasoning

`func (o *ProviderCapabilitiesSchema) SetReasoning(v bool)`

SetReasoning sets Reasoning field to given value.


### GetRerank

`func (o *ProviderCapabilitiesSchema) GetRerank() bool`

GetRerank returns the Rerank field if non-nil, zero value otherwise.

### GetRerankOk

`func (o *ProviderCapabilitiesSchema) GetRerankOk() (*bool, bool)`

GetRerankOk returns a tuple with the Rerank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRerank

`func (o *ProviderCapabilitiesSchema) SetRerank(v bool)`

SetRerank sets Rerank field to given value.


### GetResponsesApi

`func (o *ProviderCapabilitiesSchema) GetResponsesApi() bool`

GetResponsesApi returns the ResponsesApi field if non-nil, zero value otherwise.

### GetResponsesApiOk

`func (o *ProviderCapabilitiesSchema) GetResponsesApiOk() (*bool, bool)`

GetResponsesApiOk returns a tuple with the ResponsesApi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponsesApi

`func (o *ProviderCapabilitiesSchema) SetResponsesApi(v bool)`

SetResponsesApi sets ResponsesApi field to given value.


### GetStreaming

`func (o *ProviderCapabilitiesSchema) GetStreaming() bool`

GetStreaming returns the Streaming field if non-nil, zero value otherwise.

### GetStreamingOk

`func (o *ProviderCapabilitiesSchema) GetStreamingOk() (*bool, bool)`

GetStreamingOk returns a tuple with the Streaming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreaming

`func (o *ProviderCapabilitiesSchema) SetStreaming(v bool)`

SetStreaming sets Streaming field to given value.


### GetVision

`func (o *ProviderCapabilitiesSchema) GetVision() bool`

GetVision returns the Vision field if non-nil, zero value otherwise.

### GetVisionOk

`func (o *ProviderCapabilitiesSchema) GetVisionOk() (*bool, bool)`

GetVisionOk returns a tuple with the Vision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVision

`func (o *ProviderCapabilitiesSchema) SetVision(v bool)`

SetVision sets Vision field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


