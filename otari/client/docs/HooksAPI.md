# \HooksAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**HooksCheckPolicy**](HooksAPI.md#HooksCheckPolicy) | **Post** /api/v1/hooks/check | Check Policy
[**HooksRecordGuardrailDefinition**](HooksAPI.md#HooksRecordGuardrailDefinition) | **Post** /api/v1/hooks/guardrails | Record Guardrail Definition



## HooksCheckPolicy

> PolicyCheckResponse HooksCheckPolicy(ctx).PolicyCheckRequest(policyCheckRequest).Execute()

Check Policy



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	policyCheckRequest := *openapiclient.NewPolicyCheckRequest("PolicyYaml_example") // PolicyCheckRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HooksAPI.HooksCheckPolicy(context.Background()).PolicyCheckRequest(policyCheckRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HooksAPI.HooksCheckPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HooksCheckPolicy`: PolicyCheckResponse
	fmt.Fprintf(os.Stdout, "Response from `HooksAPI.HooksCheckPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiHooksCheckPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **policyCheckRequest** | [**PolicyCheckRequest**](PolicyCheckRequest.md) |  | 

### Return type

[**PolicyCheckResponse**](PolicyCheckResponse.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## HooksRecordGuardrailDefinition

> GuardrailDefinitionResponse HooksRecordGuardrailDefinition(ctx).GuardrailDefinitionRequest(guardrailDefinitionRequest).Execute()

Record Guardrail Definition



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	guardrailDefinitionRequest := *openapiclient.NewGuardrailDefinitionRequest("Digest_example", []openapiclient.GateDefinitionRequest{*openapiclient.NewGateDefinitionRequest("Enforcement_example", "GateId_example", "GateType_example", "Message_example", "Origin_example", []string{"Runs_example"})}) // GuardrailDefinitionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.HooksAPI.HooksRecordGuardrailDefinition(context.Background()).GuardrailDefinitionRequest(guardrailDefinitionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `HooksAPI.HooksRecordGuardrailDefinition``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `HooksRecordGuardrailDefinition`: GuardrailDefinitionResponse
	fmt.Fprintf(os.Stdout, "Response from `HooksAPI.HooksRecordGuardrailDefinition`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiHooksRecordGuardrailDefinitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **guardrailDefinitionRequest** | [**GuardrailDefinitionRequest**](GuardrailDefinitionRequest.md) |  | 

### Return type

[**GuardrailDefinitionResponse**](GuardrailDefinitionResponse.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

