# \PlaygroundAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PlaygroundDeletePlaygroundComparison**](PlaygroundAPI.md#PlaygroundDeletePlaygroundComparison) | **Delete** /api/v1/playground/comparisons/{comparison_id} | Delete Playground Comparison
[**PlaygroundDeletePlaygroundConversation**](PlaygroundAPI.md#PlaygroundDeletePlaygroundConversation) | **Delete** /api/v1/playground/conversations/{conversation_id} | Delete Playground Conversation
[**PlaygroundListPlaygroundComparisons**](PlaygroundAPI.md#PlaygroundListPlaygroundComparisons) | **Get** /api/v1/playground/comparisons | List Playground Comparisons
[**PlaygroundListPlaygroundConversations**](PlaygroundAPI.md#PlaygroundListPlaygroundConversations) | **Get** /api/v1/playground/conversations | List Playground Conversations
[**PlaygroundPlaygroundChatCompletions**](PlaygroundAPI.md#PlaygroundPlaygroundChatCompletions) | **Post** /api/v1/playground/chat/completions | Playground Chat Completions
[**PlaygroundReadPlaygroundConsent**](PlaygroundAPI.md#PlaygroundReadPlaygroundConsent) | **Get** /api/v1/playground/consent | Read Playground Consent
[**PlaygroundReadPlaygroundConversationMessages**](PlaygroundAPI.md#PlaygroundReadPlaygroundConversationMessages) | **Get** /api/v1/playground/conversations/{conversation_id}/messages | Read Playground Conversation Messages
[**PlaygroundReadPlaygroundFavoriteModels**](PlaygroundAPI.md#PlaygroundReadPlaygroundFavoriteModels) | **Get** /api/v1/playground/favorite-models | Read Playground Favorite Models
[**PlaygroundReadPlaygroundTools**](PlaygroundAPI.md#PlaygroundReadPlaygroundTools) | **Get** /api/v1/playground/tools | Read Playground Tools
[**PlaygroundReplacePlaygroundFavoriteModels**](PlaygroundAPI.md#PlaygroundReplacePlaygroundFavoriteModels) | **Put** /api/v1/playground/favorite-models | Replace Playground Favorite Models
[**PlaygroundSavePlaygroundComparison**](PlaygroundAPI.md#PlaygroundSavePlaygroundComparison) | **Post** /api/v1/playground/comparisons | Save Playground Comparison
[**PlaygroundSavePlaygroundConversation**](PlaygroundAPI.md#PlaygroundSavePlaygroundConversation) | **Post** /api/v1/playground/conversations | Save Playground Conversation
[**PlaygroundUpdatePlaygroundConsent**](PlaygroundAPI.md#PlaygroundUpdatePlaygroundConsent) | **Put** /api/v1/playground/consent | Update Playground Consent



## PlaygroundDeletePlaygroundComparison

> PlaygroundDeletePlaygroundComparison(ctx, comparisonId).Execute()

Delete Playground Comparison



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
	comparisonId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PlaygroundAPI.PlaygroundDeletePlaygroundComparison(context.Background(), comparisonId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlaygroundAPI.PlaygroundDeletePlaygroundComparison``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**comparisonId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlaygroundDeletePlaygroundComparisonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlaygroundDeletePlaygroundConversation

> PlaygroundDeletePlaygroundConversation(ctx, conversationId).Execute()

Delete Playground Conversation



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
	conversationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PlaygroundAPI.PlaygroundDeletePlaygroundConversation(context.Background(), conversationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlaygroundAPI.PlaygroundDeletePlaygroundConversation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conversationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlaygroundDeletePlaygroundConversationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlaygroundListPlaygroundComparisons

> PlaygroundComparisonsPublic PlaygroundListPlaygroundComparisons(ctx).WorkspaceId(workspaceId).Execute()

List Playground Comparisons



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
	workspaceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Workspace to act in. Defaults to the caller's organization's default workspace. A workspace the caller is not a member of answers 404, as a nonexistent one does. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlaygroundAPI.PlaygroundListPlaygroundComparisons(context.Background()).WorkspaceId(workspaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlaygroundAPI.PlaygroundListPlaygroundComparisons``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlaygroundListPlaygroundComparisons`: PlaygroundComparisonsPublic
	fmt.Fprintf(os.Stdout, "Response from `PlaygroundAPI.PlaygroundListPlaygroundComparisons`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlaygroundListPlaygroundComparisonsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workspaceId** | **string** | Workspace to act in. Defaults to the caller&#39;s organization&#39;s default workspace. A workspace the caller is not a member of answers 404, as a nonexistent one does. | 

### Return type

[**PlaygroundComparisonsPublic**](PlaygroundComparisonsPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlaygroundListPlaygroundConversations

> PlaygroundConversationsPublic PlaygroundListPlaygroundConversations(ctx).WorkspaceId(workspaceId).Execute()

List Playground Conversations



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
	workspaceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Workspace to act in. Defaults to the caller's organization's default workspace. A workspace the caller is not a member of answers 404, as a nonexistent one does. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlaygroundAPI.PlaygroundListPlaygroundConversations(context.Background()).WorkspaceId(workspaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlaygroundAPI.PlaygroundListPlaygroundConversations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlaygroundListPlaygroundConversations`: PlaygroundConversationsPublic
	fmt.Fprintf(os.Stdout, "Response from `PlaygroundAPI.PlaygroundListPlaygroundConversations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlaygroundListPlaygroundConversationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workspaceId** | **string** | Workspace to act in. Defaults to the caller&#39;s organization&#39;s default workspace. A workspace the caller is not a member of answers 404, as a nonexistent one does. | 

### Return type

[**PlaygroundConversationsPublic**](PlaygroundConversationsPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlaygroundPlaygroundChatCompletions

> interface{} PlaygroundPlaygroundChatCompletions(ctx).ChatCompletionRequest(chatCompletionRequest).WorkspaceId(workspaceId).Execute()

Playground Chat Completions



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
	chatCompletionRequest := *openapiclient.NewChatCompletionRequest([]openapiclient.ChatMessageInput{*openapiclient.NewChatMessageInput("Content_example", "Role_example", "Name_example", "ToolCallId_example")}, "Model_example") // ChatCompletionRequest | 
	workspaceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Workspace to act in. Defaults to the caller's organization's default workspace. A workspace the caller is not a member of answers 404, as a nonexistent one does. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlaygroundAPI.PlaygroundPlaygroundChatCompletions(context.Background()).ChatCompletionRequest(chatCompletionRequest).WorkspaceId(workspaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlaygroundAPI.PlaygroundPlaygroundChatCompletions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlaygroundPlaygroundChatCompletions`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `PlaygroundAPI.PlaygroundPlaygroundChatCompletions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlaygroundPlaygroundChatCompletionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **chatCompletionRequest** | [**ChatCompletionRequest**](ChatCompletionRequest.md) |  | 
 **workspaceId** | **string** | Workspace to act in. Defaults to the caller&#39;s organization&#39;s default workspace. A workspace the caller is not a member of answers 404, as a nonexistent one does. | 

### Return type

**interface{}**

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlaygroundReadPlaygroundConsent

> PlaygroundConsentPublic PlaygroundReadPlaygroundConsent(ctx).Execute()

Read Playground Consent



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlaygroundAPI.PlaygroundReadPlaygroundConsent(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlaygroundAPI.PlaygroundReadPlaygroundConsent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlaygroundReadPlaygroundConsent`: PlaygroundConsentPublic
	fmt.Fprintf(os.Stdout, "Response from `PlaygroundAPI.PlaygroundReadPlaygroundConsent`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPlaygroundReadPlaygroundConsentRequest struct via the builder pattern


### Return type

[**PlaygroundConsentPublic**](PlaygroundConsentPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlaygroundReadPlaygroundConversationMessages

> PlaygroundMessagesPublic PlaygroundReadPlaygroundConversationMessages(ctx, conversationId).Execute()

Read Playground Conversation Messages



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
	conversationId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlaygroundAPI.PlaygroundReadPlaygroundConversationMessages(context.Background(), conversationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlaygroundAPI.PlaygroundReadPlaygroundConversationMessages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlaygroundReadPlaygroundConversationMessages`: PlaygroundMessagesPublic
	fmt.Fprintf(os.Stdout, "Response from `PlaygroundAPI.PlaygroundReadPlaygroundConversationMessages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conversationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlaygroundReadPlaygroundConversationMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PlaygroundMessagesPublic**](PlaygroundMessagesPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlaygroundReadPlaygroundFavoriteModels

> PlaygroundFavoriteModelsPublic PlaygroundReadPlaygroundFavoriteModels(ctx).WorkspaceId(workspaceId).Execute()

Read Playground Favorite Models



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
	workspaceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Workspace to act in. Defaults to the caller's organization's default workspace. A workspace the caller is not a member of answers 404, as a nonexistent one does. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlaygroundAPI.PlaygroundReadPlaygroundFavoriteModels(context.Background()).WorkspaceId(workspaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlaygroundAPI.PlaygroundReadPlaygroundFavoriteModels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlaygroundReadPlaygroundFavoriteModels`: PlaygroundFavoriteModelsPublic
	fmt.Fprintf(os.Stdout, "Response from `PlaygroundAPI.PlaygroundReadPlaygroundFavoriteModels`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlaygroundReadPlaygroundFavoriteModelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workspaceId** | **string** | Workspace to act in. Defaults to the caller&#39;s organization&#39;s default workspace. A workspace the caller is not a member of answers 404, as a nonexistent one does. | 

### Return type

[**PlaygroundFavoriteModelsPublic**](PlaygroundFavoriteModelsPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlaygroundReadPlaygroundTools

> PlaygroundToolsResponse PlaygroundReadPlaygroundTools(ctx).WorkspaceId(workspaceId).Execute()

Read Playground Tools



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
	workspaceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Workspace to act in. Defaults to the caller's organization's default workspace. A workspace the caller is not a member of answers 404, as a nonexistent one does. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlaygroundAPI.PlaygroundReadPlaygroundTools(context.Background()).WorkspaceId(workspaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlaygroundAPI.PlaygroundReadPlaygroundTools``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlaygroundReadPlaygroundTools`: PlaygroundToolsResponse
	fmt.Fprintf(os.Stdout, "Response from `PlaygroundAPI.PlaygroundReadPlaygroundTools`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlaygroundReadPlaygroundToolsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workspaceId** | **string** | Workspace to act in. Defaults to the caller&#39;s organization&#39;s default workspace. A workspace the caller is not a member of answers 404, as a nonexistent one does. | 

### Return type

[**PlaygroundToolsResponse**](PlaygroundToolsResponse.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlaygroundReplacePlaygroundFavoriteModels

> PlaygroundFavoriteModelsPublic PlaygroundReplacePlaygroundFavoriteModels(ctx).PlaygroundFavoriteModelsUpdate(playgroundFavoriteModelsUpdate).WorkspaceId(workspaceId).Execute()

Replace Playground Favorite Models



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
	playgroundFavoriteModelsUpdate := *openapiclient.NewPlaygroundFavoriteModelsUpdate([]string{"ModelKeys_example"}) // PlaygroundFavoriteModelsUpdate | 
	workspaceId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Workspace to act in. Defaults to the caller's organization's default workspace. A workspace the caller is not a member of answers 404, as a nonexistent one does. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlaygroundAPI.PlaygroundReplacePlaygroundFavoriteModels(context.Background()).PlaygroundFavoriteModelsUpdate(playgroundFavoriteModelsUpdate).WorkspaceId(workspaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlaygroundAPI.PlaygroundReplacePlaygroundFavoriteModels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlaygroundReplacePlaygroundFavoriteModels`: PlaygroundFavoriteModelsPublic
	fmt.Fprintf(os.Stdout, "Response from `PlaygroundAPI.PlaygroundReplacePlaygroundFavoriteModels`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlaygroundReplacePlaygroundFavoriteModelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **playgroundFavoriteModelsUpdate** | [**PlaygroundFavoriteModelsUpdate**](PlaygroundFavoriteModelsUpdate.md) |  | 
 **workspaceId** | **string** | Workspace to act in. Defaults to the caller&#39;s organization&#39;s default workspace. A workspace the caller is not a member of answers 404, as a nonexistent one does. | 

### Return type

[**PlaygroundFavoriteModelsPublic**](PlaygroundFavoriteModelsPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlaygroundSavePlaygroundComparison

> PlaygroundComparisonSummary PlaygroundSavePlaygroundComparison(ctx).PlaygroundComparisonCreate(playgroundComparisonCreate).Execute()

Save Playground Comparison



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
	playgroundComparisonCreate := *openapiclient.NewPlaygroundComparisonCreate("ModelA_example", "ModelAAnswer_example", "ModelB_example", "ModelBAnswer_example", "Preference_example", "UserQuestion_example", "WorkspaceId_example") // PlaygroundComparisonCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlaygroundAPI.PlaygroundSavePlaygroundComparison(context.Background()).PlaygroundComparisonCreate(playgroundComparisonCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlaygroundAPI.PlaygroundSavePlaygroundComparison``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlaygroundSavePlaygroundComparison`: PlaygroundComparisonSummary
	fmt.Fprintf(os.Stdout, "Response from `PlaygroundAPI.PlaygroundSavePlaygroundComparison`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlaygroundSavePlaygroundComparisonRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **playgroundComparisonCreate** | [**PlaygroundComparisonCreate**](PlaygroundComparisonCreate.md) |  | 

### Return type

[**PlaygroundComparisonSummary**](PlaygroundComparisonSummary.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlaygroundSavePlaygroundConversation

> PlaygroundConversationSummary PlaygroundSavePlaygroundConversation(ctx).PlaygroundConversationCreate(playgroundConversationCreate).Execute()

Save Playground Conversation



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
	playgroundConversationCreate := *openapiclient.NewPlaygroundConversationCreate([]openapiclient.PlaygroundMessageCreate{*openapiclient.NewPlaygroundMessageCreate("Content_example", "Role_example")}, "Model_example", "Title_example", "WorkspaceId_example") // PlaygroundConversationCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlaygroundAPI.PlaygroundSavePlaygroundConversation(context.Background()).PlaygroundConversationCreate(playgroundConversationCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlaygroundAPI.PlaygroundSavePlaygroundConversation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlaygroundSavePlaygroundConversation`: PlaygroundConversationSummary
	fmt.Fprintf(os.Stdout, "Response from `PlaygroundAPI.PlaygroundSavePlaygroundConversation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlaygroundSavePlaygroundConversationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **playgroundConversationCreate** | [**PlaygroundConversationCreate**](PlaygroundConversationCreate.md) |  | 

### Return type

[**PlaygroundConversationSummary**](PlaygroundConversationSummary.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlaygroundUpdatePlaygroundConsent

> PlaygroundConsentPublic PlaygroundUpdatePlaygroundConsent(ctx).PlaygroundConsentUpdate(playgroundConsentUpdate).Execute()

Update Playground Consent



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
	playgroundConsentUpdate := *openapiclient.NewPlaygroundConsentUpdate() // PlaygroundConsentUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlaygroundAPI.PlaygroundUpdatePlaygroundConsent(context.Background()).PlaygroundConsentUpdate(playgroundConsentUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlaygroundAPI.PlaygroundUpdatePlaygroundConsent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PlaygroundUpdatePlaygroundConsent`: PlaygroundConsentPublic
	fmt.Fprintf(os.Stdout, "Response from `PlaygroundAPI.PlaygroundUpdatePlaygroundConsent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPlaygroundUpdatePlaygroundConsentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **playgroundConsentUpdate** | [**PlaygroundConsentUpdate**](PlaygroundConsentUpdate.md) |  | 

### Return type

[**PlaygroundConsentPublic**](PlaygroundConsentPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

