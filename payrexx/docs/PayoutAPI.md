# \PayoutAPI

All URIs are relative to *https://api.payrexx.com/v1.16*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PayoutList**](PayoutAPI.md#PayoutList) | **Get** /Payout/ | Retrieve Payouts
[**PayoutRetrieve**](PayoutAPI.md#PayoutRetrieve) | **Get** /Payout/{uuid}/ | Retrieve a Payout Without Details
[**PayoutRetrieveWithDetails**](PayoutAPI.md#PayoutRetrieveWithDetails) | **Get** /Payout/{uuid}/details | Retrieve a Payout With Details



## PayoutList

> PayoutList200Response PayoutList(ctx).PayoutListRequest(payoutListRequest).Execute()

Retrieve Payouts



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
	payoutListRequest := *openapiclient.NewPayoutListRequest() // PayoutListRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PayoutAPI.PayoutList(context.Background()).PayoutListRequest(payoutListRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PayoutAPI.PayoutList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PayoutList`: PayoutList200Response
	fmt.Fprintf(os.Stdout, "Response from `PayoutAPI.PayoutList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPayoutListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **payoutListRequest** | [**PayoutListRequest**](PayoutListRequest.md) |  | 

### Return type

[**PayoutList200Response**](PayoutList200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PayoutRetrieve

> PayoutRetrieve200Response PayoutRetrieve(ctx, uuid).Execute()

Retrieve a Payout Without Details



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
	uuid := "uuid_example" // string | Payout UUID (default to "8")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PayoutAPI.PayoutRetrieve(context.Background(), uuid).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PayoutAPI.PayoutRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PayoutRetrieve`: PayoutRetrieve200Response
	fmt.Fprintf(os.Stdout, "Response from `PayoutAPI.PayoutRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | Payout UUID | [default to &quot;8&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiPayoutRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PayoutRetrieve200Response**](PayoutRetrieve200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PayoutRetrieveWithDetails

> PayoutRetrieve200Response PayoutRetrieveWithDetails(ctx, uuid).Offset(offset).Limit(limit).Execute()

Retrieve a Payout With Details



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
	uuid := "uuid_example" // string | Payout UUID (default to "8")
	offset := int32(56) // int32 | Offset when using limit. (optional) (default to 0)
	limit := int32(56) // int32 | Limits the number of details to retrieve, recommended size is 100 elements. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PayoutAPI.PayoutRetrieveWithDetails(context.Background(), uuid).Offset(offset).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PayoutAPI.PayoutRetrieveWithDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PayoutRetrieveWithDetails`: PayoutRetrieve200Response
	fmt.Fprintf(os.Stdout, "Response from `PayoutAPI.PayoutRetrieveWithDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uuid** | **string** | Payout UUID | [default to &quot;8&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiPayoutRetrieveWithDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **offset** | **int32** | Offset when using limit. | [default to 0]
 **limit** | **int32** | Limits the number of details to retrieve, recommended size is 100 elements. | 

### Return type

[**PayoutRetrieve200Response**](PayoutRetrieve200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

