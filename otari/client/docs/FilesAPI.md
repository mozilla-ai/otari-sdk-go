# \FilesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFileV1FilesPost**](FilesAPI.md#CreateFileV1FilesPost) | **Post** /v1/files | Create File
[**DeleteFileV1FilesFileIdDelete**](FilesAPI.md#DeleteFileV1FilesFileIdDelete) | **Delete** /v1/files/{file_id} | Delete File
[**GetFileContentV1FilesFileIdContentGet**](FilesAPI.md#GetFileContentV1FilesFileIdContentGet) | **Get** /v1/files/{file_id}/content | Get File Content
[**GetFileV1FilesFileIdGet**](FilesAPI.md#GetFileV1FilesFileIdGet) | **Get** /v1/files/{file_id} | Get File
[**ListFilesV1FilesGet**](FilesAPI.md#ListFilesV1FilesGet) | **Get** /v1/files | List Files



## CreateFileV1FilesPost

> map[string]interface{} CreateFileV1FilesPost(ctx).File(file).Purpose(purpose).User(user).Execute()

Create File



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
	file := "file_example" // string | 
	purpose := "purpose_example" // string |  (optional) (default to "user_data")
	user := "user_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesAPI.CreateFileV1FilesPost(context.Background()).File(file).Purpose(purpose).User(user).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.CreateFileV1FilesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFileV1FilesPost`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FilesAPI.CreateFileV1FilesPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFileV1FilesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | **string** |  | 
 **purpose** | **string** |  | [default to &quot;user_data&quot;]
 **user** | **string** |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFileV1FilesFileIdDelete

> map[string]interface{} DeleteFileV1FilesFileIdDelete(ctx, fileId).User(user).Execute()

Delete File



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
	fileId := "fileId_example" // string | 
	user := "user_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesAPI.DeleteFileV1FilesFileIdDelete(context.Background(), fileId).User(user).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.DeleteFileV1FilesFileIdDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteFileV1FilesFileIdDelete`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FilesAPI.DeleteFileV1FilesFileIdDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFileV1FilesFileIdDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **user** | **string** |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFileContentV1FilesFileIdContentGet

> interface{} GetFileContentV1FilesFileIdContentGet(ctx, fileId).User(user).Execute()

Get File Content



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
	fileId := "fileId_example" // string | 
	user := "user_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesAPI.GetFileContentV1FilesFileIdContentGet(context.Background(), fileId).User(user).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.GetFileContentV1FilesFileIdContentGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFileContentV1FilesFileIdContentGet`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `FilesAPI.GetFileContentV1FilesFileIdContentGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFileContentV1FilesFileIdContentGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **user** | **string** |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFileV1FilesFileIdGet

> map[string]interface{} GetFileV1FilesFileIdGet(ctx, fileId).User(user).Execute()

Get File



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
	fileId := "fileId_example" // string | 
	user := "user_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesAPI.GetFileV1FilesFileIdGet(context.Background(), fileId).User(user).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.GetFileV1FilesFileIdGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFileV1FilesFileIdGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FilesAPI.GetFileV1FilesFileIdGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fileId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFileV1FilesFileIdGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **user** | **string** |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFilesV1FilesGet

> map[string]interface{} ListFilesV1FilesGet(ctx).User(user).Purpose(purpose).Execute()

List Files



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
	user := "user_example" // string |  (optional)
	purpose := "purpose_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilesAPI.ListFilesV1FilesGet(context.Background()).User(user).Purpose(purpose).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilesAPI.ListFilesV1FilesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListFilesV1FilesGet`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `FilesAPI.ListFilesV1FilesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListFilesV1FilesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **user** | **string** |  | 
 **purpose** | **string** |  | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

