# \OrganizationPricingAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**OrganizationPricingCreateOrganizationPricing**](OrganizationPricingAPI.md#OrganizationPricingCreateOrganizationPricing) | **Post** /api/v1/organizations/me/pricing | Create Organization Pricing
[**OrganizationPricingDeleteOrganizationPricing**](OrganizationPricingAPI.md#OrganizationPricingDeleteOrganizationPricing) | **Delete** /api/v1/organizations/me/pricing/{pricing_id} | Delete Organization Pricing
[**OrganizationPricingListOrganizationPricing**](OrganizationPricingAPI.md#OrganizationPricingListOrganizationPricing) | **Get** /api/v1/organizations/me/pricing | List Organization Pricing
[**OrganizationPricingReplaceOrganizationPricing**](OrganizationPricingAPI.md#OrganizationPricingReplaceOrganizationPricing) | **Put** /api/v1/organizations/me/pricing/{pricing_id} | Replace Organization Pricing



## OrganizationPricingCreateOrganizationPricing

> OrganizationModelPricingPublic OrganizationPricingCreateOrganizationPricing(ctx).OrganizationModelPricingCreate(organizationModelPricingCreate).Execute()

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
	resp, r, err := apiClient.OrganizationPricingAPI.OrganizationPricingCreateOrganizationPricing(context.Background()).OrganizationModelPricingCreate(organizationModelPricingCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationPricingAPI.OrganizationPricingCreateOrganizationPricing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationPricingCreateOrganizationPricing`: OrganizationModelPricingPublic
	fmt.Fprintf(os.Stdout, "Response from `OrganizationPricingAPI.OrganizationPricingCreateOrganizationPricing`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationPricingCreateOrganizationPricingRequest struct via the builder pattern


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


## OrganizationPricingDeleteOrganizationPricing

> OrganizationPricingDeleteOrganizationPricing(ctx, pricingId).Execute()

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
	r, err := apiClient.OrganizationPricingAPI.OrganizationPricingDeleteOrganizationPricing(context.Background(), pricingId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationPricingAPI.OrganizationPricingDeleteOrganizationPricing``: %v\n", err)
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

Other parameters are passed through a pointer to a apiOrganizationPricingDeleteOrganizationPricingRequest struct via the builder pattern


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


## OrganizationPricingListOrganizationPricing

> OrganizationModelPricingsPublic OrganizationPricingListOrganizationPricing(ctx).Skip(skip).Limit(limit).Execute()

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
	resp, r, err := apiClient.OrganizationPricingAPI.OrganizationPricingListOrganizationPricing(context.Background()).Skip(skip).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationPricingAPI.OrganizationPricingListOrganizationPricing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationPricingListOrganizationPricing`: OrganizationModelPricingsPublic
	fmt.Fprintf(os.Stdout, "Response from `OrganizationPricingAPI.OrganizationPricingListOrganizationPricing`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationPricingListOrganizationPricingRequest struct via the builder pattern


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


## OrganizationPricingReplaceOrganizationPricing

> OrganizationModelPricingPublic OrganizationPricingReplaceOrganizationPricing(ctx, pricingId).OrganizationModelPricingUpdate(organizationModelPricingUpdate).Execute()

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
	resp, r, err := apiClient.OrganizationPricingAPI.OrganizationPricingReplaceOrganizationPricing(context.Background(), pricingId).OrganizationModelPricingUpdate(organizationModelPricingUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrganizationPricingAPI.OrganizationPricingReplaceOrganizationPricing``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `OrganizationPricingReplaceOrganizationPricing`: OrganizationModelPricingPublic
	fmt.Fprintf(os.Stdout, "Response from `OrganizationPricingAPI.OrganizationPricingReplaceOrganizationPricing`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pricingId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiOrganizationPricingReplaceOrganizationPricingRequest struct via the builder pattern


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

