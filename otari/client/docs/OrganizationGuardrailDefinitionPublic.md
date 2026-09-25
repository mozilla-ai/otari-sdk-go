# OrganizationGuardrailDefinitionPublic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BuildState** | **string** | Whether the worker that answered this request holds a guardrail built from this version of the definition. built means it does and the check runs; failed means that worker tried these exact arguments and could not build them, so every mandate pointing here is unevaluable; pending means it holds nothing for this version yet, which is the answer right after a write and on any worker that has not caught up within the refresh interval; disabled means the definition is switched off and nothing is built on purpose. It answers for one worker, so two reads can disagree while a write propagates. Why a build failed is never reported here: the reason is in the gateway&#39;s log | 
**CreateKwargs** | **map[string]interface{}** | The constructor arguments the catalog does not mark secret, as they were stored. Returned in clear: the secrets were taken out of this map by flag, and a form has to round-trip an endpoint or a project id | 
**CreateSecrets** | **map[string]string** | The secret constructor arguments this definition holds, each as ***. Sending one back unchanged keeps the stored value; sending a new one rotates it, and leaving one out clears it | 
**CreatedAt** | **string** |  | 
**Enabled** | **bool** |  | 
**GuardrailName** | **string** |  | 
**Id** | **string** |  | 
**Name** | **string** |  | 
**OrganizationId** | **string** |  | 
**SecretsDecryptable** | **bool** | False when the stored secrets cannot be read with the current OTARI_SECRET_KEY, in which case create_secrets is empty and the definition needs its credentials sent again | 
**UpdatedAt** | **string** |  | 

## Methods

### NewOrganizationGuardrailDefinitionPublic

`func NewOrganizationGuardrailDefinitionPublic(buildState string, createKwargs map[string]interface{}, createSecrets map[string]string, createdAt string, enabled bool, guardrailName string, id string, name string, organizationId string, secretsDecryptable bool, updatedAt string, ) *OrganizationGuardrailDefinitionPublic`

NewOrganizationGuardrailDefinitionPublic instantiates a new OrganizationGuardrailDefinitionPublic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrganizationGuardrailDefinitionPublicWithDefaults

`func NewOrganizationGuardrailDefinitionPublicWithDefaults() *OrganizationGuardrailDefinitionPublic`

NewOrganizationGuardrailDefinitionPublicWithDefaults instantiates a new OrganizationGuardrailDefinitionPublic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBuildState

`func (o *OrganizationGuardrailDefinitionPublic) GetBuildState() string`

GetBuildState returns the BuildState field if non-nil, zero value otherwise.

### GetBuildStateOk

`func (o *OrganizationGuardrailDefinitionPublic) GetBuildStateOk() (*string, bool)`

GetBuildStateOk returns a tuple with the BuildState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildState

`func (o *OrganizationGuardrailDefinitionPublic) SetBuildState(v string)`

SetBuildState sets BuildState field to given value.


### GetCreateKwargs

`func (o *OrganizationGuardrailDefinitionPublic) GetCreateKwargs() map[string]interface{}`

GetCreateKwargs returns the CreateKwargs field if non-nil, zero value otherwise.

### GetCreateKwargsOk

`func (o *OrganizationGuardrailDefinitionPublic) GetCreateKwargsOk() (*map[string]interface{}, bool)`

GetCreateKwargsOk returns a tuple with the CreateKwargs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateKwargs

`func (o *OrganizationGuardrailDefinitionPublic) SetCreateKwargs(v map[string]interface{})`

SetCreateKwargs sets CreateKwargs field to given value.


### GetCreateSecrets

`func (o *OrganizationGuardrailDefinitionPublic) GetCreateSecrets() map[string]string`

GetCreateSecrets returns the CreateSecrets field if non-nil, zero value otherwise.

### GetCreateSecretsOk

`func (o *OrganizationGuardrailDefinitionPublic) GetCreateSecretsOk() (*map[string]string, bool)`

GetCreateSecretsOk returns a tuple with the CreateSecrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateSecrets

`func (o *OrganizationGuardrailDefinitionPublic) SetCreateSecrets(v map[string]string)`

SetCreateSecrets sets CreateSecrets field to given value.


### GetCreatedAt

`func (o *OrganizationGuardrailDefinitionPublic) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OrganizationGuardrailDefinitionPublic) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OrganizationGuardrailDefinitionPublic) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetEnabled

`func (o *OrganizationGuardrailDefinitionPublic) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *OrganizationGuardrailDefinitionPublic) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *OrganizationGuardrailDefinitionPublic) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetGuardrailName

`func (o *OrganizationGuardrailDefinitionPublic) GetGuardrailName() string`

GetGuardrailName returns the GuardrailName field if non-nil, zero value otherwise.

### GetGuardrailNameOk

`func (o *OrganizationGuardrailDefinitionPublic) GetGuardrailNameOk() (*string, bool)`

GetGuardrailNameOk returns a tuple with the GuardrailName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuardrailName

`func (o *OrganizationGuardrailDefinitionPublic) SetGuardrailName(v string)`

SetGuardrailName sets GuardrailName field to given value.


### GetId

`func (o *OrganizationGuardrailDefinitionPublic) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrganizationGuardrailDefinitionPublic) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrganizationGuardrailDefinitionPublic) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *OrganizationGuardrailDefinitionPublic) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OrganizationGuardrailDefinitionPublic) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OrganizationGuardrailDefinitionPublic) SetName(v string)`

SetName sets Name field to given value.


### GetOrganizationId

`func (o *OrganizationGuardrailDefinitionPublic) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *OrganizationGuardrailDefinitionPublic) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *OrganizationGuardrailDefinitionPublic) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.


### GetSecretsDecryptable

`func (o *OrganizationGuardrailDefinitionPublic) GetSecretsDecryptable() bool`

GetSecretsDecryptable returns the SecretsDecryptable field if non-nil, zero value otherwise.

### GetSecretsDecryptableOk

`func (o *OrganizationGuardrailDefinitionPublic) GetSecretsDecryptableOk() (*bool, bool)`

GetSecretsDecryptableOk returns a tuple with the SecretsDecryptable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretsDecryptable

`func (o *OrganizationGuardrailDefinitionPublic) SetSecretsDecryptable(v bool)`

SetSecretsDecryptable sets SecretsDecryptable field to given value.


### GetUpdatedAt

`func (o *OrganizationGuardrailDefinitionPublic) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *OrganizationGuardrailDefinitionPublic) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *OrganizationGuardrailDefinitionPublic) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


