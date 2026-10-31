# \DecisionsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DecisionsCreateDecision**](DecisionsAPI.md#DecisionsCreateDecision) | **Post** /api/v1/decisions | Create Decision
[**DecisionsCreateSystemoneDecision**](DecisionsAPI.md#DecisionsCreateSystemoneDecision) | **Post** /api/v1/systemone | Create Systemone Decision



## DecisionsCreateDecision

> DecisionResponse DecisionsCreateDecision(ctx).DecisionRequest(decisionRequest).Execute()

Create Decision



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
	decisionRequest := *openapiclient.NewDecisionRequest("Model_example", map[string]QuestionsValue{"key": openapiclient.Questions_value{ChoiceQuestion: openapiclient.NewChoiceQuestion(map[string]CriteriaValue{"key": *openapiclient.NewCriteriaValue()}, *openapiclient.NewInstructions(), "Type_example")}}, *openapiclient.NewState()) // DecisionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DecisionsAPI.DecisionsCreateDecision(context.Background()).DecisionRequest(decisionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DecisionsAPI.DecisionsCreateDecision``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DecisionsCreateDecision`: DecisionResponse
	fmt.Fprintf(os.Stdout, "Response from `DecisionsAPI.DecisionsCreateDecision`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDecisionsCreateDecisionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **decisionRequest** | [**DecisionRequest**](DecisionRequest.md) |  | 

### Return type

[**DecisionResponse**](DecisionResponse.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DecisionsCreateSystemoneDecision

> DecisionResponse DecisionsCreateSystemoneDecision(ctx).DecisionRequest(decisionRequest).Execute()

Create Systemone Decision



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
	decisionRequest := *openapiclient.NewDecisionRequest("Model_example", map[string]QuestionsValue{"key": openapiclient.Questions_value{ChoiceQuestion: openapiclient.NewChoiceQuestion(map[string]CriteriaValue{"key": *openapiclient.NewCriteriaValue()}, *openapiclient.NewInstructions(), "Type_example")}}, *openapiclient.NewState()) // DecisionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DecisionsAPI.DecisionsCreateSystemoneDecision(context.Background()).DecisionRequest(decisionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DecisionsAPI.DecisionsCreateSystemoneDecision``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DecisionsCreateSystemoneDecision`: DecisionResponse
	fmt.Fprintf(os.Stdout, "Response from `DecisionsAPI.DecisionsCreateSystemoneDecision`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDecisionsCreateSystemoneDecisionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **decisionRequest** | [**DecisionRequest**](DecisionRequest.md) |  | 

### Return type

[**DecisionResponse**](DecisionResponse.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

