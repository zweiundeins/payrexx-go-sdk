# \GatewayAPI

All URIs are relative to *https://api.payrexx.com/v1.16*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GatewayCreate**](GatewayAPI.md#GatewayCreate) | **Post** /Gateway/ | Create a Gateway
[**GatewayDelete**](GatewayAPI.md#GatewayDelete) | **Delete** /Gateway/{id}/ | Delete a Gateway
[**GatewayRetrieve**](GatewayAPI.md#GatewayRetrieve) | **Get** /Gateway/{id}/ | Retrieve a Gateway



## GatewayCreate

> GatewayCreate200Response GatewayCreate(ctx).GatewayCreateRequest(gatewayCreateRequest).Execute()

Create a Gateway



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
	gatewayCreateRequest := *openapiclient.NewGatewayCreateRequest(int32(123), "Currency_example") // GatewayCreateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatewayAPI.GatewayCreate(context.Background()).GatewayCreateRequest(gatewayCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayAPI.GatewayCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GatewayCreate`: GatewayCreate200Response
	fmt.Fprintf(os.Stdout, "Response from `GatewayAPI.GatewayCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGatewayCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **gatewayCreateRequest** | [**GatewayCreateRequest**](GatewayCreateRequest.md) |  | 

### Return type

[**GatewayCreate200Response**](GatewayCreate200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GatewayDelete

> GatewayCreate200Response GatewayDelete(ctx, id).Execute()

Delete a Gateway



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
	id := int32(56) // int32 | The gateway id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatewayAPI.GatewayDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayAPI.GatewayDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GatewayDelete`: GatewayCreate200Response
	fmt.Fprintf(os.Stdout, "Response from `GatewayAPI.GatewayDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The gateway id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGatewayDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GatewayCreate200Response**](GatewayCreate200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GatewayRetrieve

> GatewayCreate200Response GatewayRetrieve(ctx, id).Execute()

Retrieve a Gateway



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
	id := int32(56) // int32 | ID of gateway payment to retrieve. (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatewayAPI.GatewayRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewayAPI.GatewayRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GatewayRetrieve`: GatewayCreate200Response
	fmt.Fprintf(os.Stdout, "Response from `GatewayAPI.GatewayRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID of gateway payment to retrieve. | [default to 0]

### Other Parameters

Other parameters are passed through a pointer to a apiGatewayRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GatewayCreate200Response**](GatewayCreate200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

