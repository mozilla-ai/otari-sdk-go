# \OrganizationPricingAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrganizationPricingV1OrganizationsMePricingPost**](OrganizationPricingAPI.md#CreateOrganizationPricingV1OrganizationsMePricingPost) | **Post** /v1/organizations/me/pricing | Create Organization Pricing
[**DeleteOrganizationPricingV1OrganizationsMePricingPricingIdDelete**](OrganizationPricingAPI.md#DeleteOrganizationPricingV1OrganizationsMePricingPricingIdDelete) | **Delete** /v1/organizations/me/pricing/{pricing_id} | Delete Organization Pricing
[**ListOrganizationPricingV1OrganizationsMePricingGet**](OrganizationPricingAPI.md#ListOrganizationPricingV1OrganizationsMePricingGet) | **Get** /v1/organizations/me/pricing | List Organization Pricing
[**ReplaceOrganizationPricingV1OrganizationsMePricingPricingIdPut**](OrganizationPricingAPI.md#ReplaceOrganizationPricingV1OrganizationsMePricingPricingIdPut) | **Put** /v1/organizations/me/pricing/{pricing_id} | Replace Organization Pricing



## CreateOrganizationPricingV1OrganizationsMePricingPost

> OrganizationModelPricingPublic CreateOrganizationPricingV1OrganizationsMePricingPost(ctx).OrganizationModelPricingCreate(organizationModelPricingCreate).Execute()

Create Organization Pricing



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
	organizationModelPricingCreate := *openapiclient.NewOrganizationModelPricingCreate(float32(123), "ModelKey_example", float32(123)) // OrganizationModelPricingCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationPricingAPI.CreateOrganizationPricingV1OrganizationsMePricingPost(context.Background()).OrganizationModelPricingCreate(organizationModelPricingCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationPricingAPI.CreateOrganizationPricingV1OrganizationsMePricingPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrganizationPricingV1OrganizationsMePricingPost`: OrganizationModelPricingPublic
	fmt.Fprintf(os.Stdout, "Response from `OrganizationPricingAPI.CreateOrganizationPricingV1OrganizationsMePricingPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrganizationPricingV1OrganizationsMePricingPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **organizationModelPricingCreate** | [**OrganizationModelPricingCreate**](OrganizationModelPricingCreate.md) |  | 

### Return type

[**OrganizationModelPricingPublic**](OrganizationModelPricingPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrganizationPricingV1OrganizationsMePricingPricingIdDelete

> DeleteOrganizationPricingV1OrganizationsMePricingPricingIdDelete(ctx, pricingId).Execute()

Delete Organization Pricing



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
	pricingId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrganizationPricingAPI.DeleteOrganizationPricingV1OrganizationsMePricingPricingIdDelete(context.Background(), pricingId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationPricingAPI.DeleteOrganizationPricingV1OrganizationsMePricingPricingIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pricingId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOrganizationPricingV1OrganizationsMePricingPricingIdDeleteRequest struct via the builder pattern


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


## ListOrganizationPricingV1OrganizationsMePricingGet

> OrganizationModelPricingsPublic ListOrganizationPricingV1OrganizationsMePricingGet(ctx).Skip(skip).Limit(limit).Execute()

List Organization Pricing



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
	skip := int32(56) // int32 | Number of records to skip (optional) (default to 0)
	limit := int32(56) // int32 | Maximum number of records to return (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationPricingAPI.ListOrganizationPricingV1OrganizationsMePricingGet(context.Background()).Skip(skip).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationPricingAPI.ListOrganizationPricingV1OrganizationsMePricingGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListOrganizationPricingV1OrganizationsMePricingGet`: OrganizationModelPricingsPublic
	fmt.Fprintf(os.Stdout, "Response from `OrganizationPricingAPI.ListOrganizationPricingV1OrganizationsMePricingGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListOrganizationPricingV1OrganizationsMePricingGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **skip** | **int32** | Number of records to skip | [default to 0]
 **limit** | **int32** | Maximum number of records to return | [default to 100]

### Return type

[**OrganizationModelPricingsPublic**](OrganizationModelPricingsPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceOrganizationPricingV1OrganizationsMePricingPricingIdPut

> OrganizationModelPricingPublic ReplaceOrganizationPricingV1OrganizationsMePricingPricingIdPut(ctx, pricingId).OrganizationModelPricingUpdate(organizationModelPricingUpdate).Execute()

Replace Organization Pricing



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	pricingId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
	organizationModelPricingUpdate := *openapiclient.NewOrganizationModelPricingUpdate(time.Now(), float32(123), float32(123)) // OrganizationModelPricingUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrganizationPricingAPI.ReplaceOrganizationPricingV1OrganizationsMePricingPricingIdPut(context.Background(), pricingId).OrganizationModelPricingUpdate(organizationModelPricingUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationPricingAPI.ReplaceOrganizationPricingV1OrganizationsMePricingPricingIdPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplaceOrganizationPricingV1OrganizationsMePricingPricingIdPut`: OrganizationModelPricingPublic
	fmt.Fprintf(os.Stdout, "Response from `OrganizationPricingAPI.ReplaceOrganizationPricingV1OrganizationsMePricingPricingIdPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pricingId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceOrganizationPricingV1OrganizationsMePricingPricingIdPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **organizationModelPricingUpdate** | [**OrganizationModelPricingUpdate**](OrganizationModelPricingUpdate.md) |  | 

### Return type

[**OrganizationModelPricingPublic**](OrganizationModelPricingPublic.md)

### Authorization

[XApiKeyAuth](../README.md#XApiKeyAuth), [ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

