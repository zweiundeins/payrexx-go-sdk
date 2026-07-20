# \SubscriptionAPI

All URIs are relative to *https://api.payrexx.com/v1.16*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SubscriptionCancel**](SubscriptionAPI.md#SubscriptionCancel) | **Delete** /Subscription/{id}/ | Cancel a Subscription
[**SubscriptionCreate**](SubscriptionAPI.md#SubscriptionCreate) | **Post** /Subscription/ | Create a New Subscription
[**SubscriptionList**](SubscriptionAPI.md#SubscriptionList) | **Get** /Subscription/ | Retrieve Subscriptions 
[**SubscriptionRetrieve**](SubscriptionAPI.md#SubscriptionRetrieve) | **Get** /Subscription/{id}/ | 
[**SubscriptionUpdate**](SubscriptionAPI.md#SubscriptionUpdate) | **Put** /Subscription/{id}/ | Update a Subscription



## SubscriptionCancel

> SubscriptionList200Response SubscriptionCancel(ctx, id).Execute()

Cancel a Subscription



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
	id := int32(56) // int32 | The subscription id.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionAPI.SubscriptionCancel(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionAPI.SubscriptionCancel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionCancel`: SubscriptionList200Response
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionAPI.SubscriptionCancel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The subscription id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionCancelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SubscriptionList200Response**](SubscriptionList200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionCreate

> SubscriptionList200Response SubscriptionCreate(ctx).SubscriptionCreateRequest(subscriptionCreateRequest).Execute()

Create a New Subscription



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
	subscriptionCreateRequest := *openapiclient.NewSubscriptionCreateRequest("UserId_example", "Psp_example", "Amount_example", "Currency_example", "Purpose_example", "PaymentInterval_example", "Period_example", "CancellationInterval_example") // SubscriptionCreateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionAPI.SubscriptionCreate(context.Background()).SubscriptionCreateRequest(subscriptionCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionAPI.SubscriptionCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionCreate`: SubscriptionList200Response
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionAPI.SubscriptionCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscriptionCreateRequest** | [**SubscriptionCreateRequest**](SubscriptionCreateRequest.md) |  | 

### Return type

[**SubscriptionList200Response**](SubscriptionList200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionList

> SubscriptionList200Response SubscriptionList(ctx).SubscriptionListRequest(subscriptionListRequest).Execute()

Retrieve Subscriptions 



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
	subscriptionListRequest := *openapiclient.NewSubscriptionListRequest() // SubscriptionListRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionAPI.SubscriptionList(context.Background()).SubscriptionListRequest(subscriptionListRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionAPI.SubscriptionList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionList`: SubscriptionList200Response
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionAPI.SubscriptionList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscriptionListRequest** | [**SubscriptionListRequest**](SubscriptionListRequest.md) |  | 

### Return type

[**SubscriptionList200Response**](SubscriptionList200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionRetrieve

> SubscriptionList200Response SubscriptionRetrieve(ctx, id).Execute()





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
	id := int32(56) // int32 | The id of the subscription to retrieve.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionAPI.SubscriptionRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionAPI.SubscriptionRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionRetrieve`: SubscriptionList200Response
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionAPI.SubscriptionRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The id of the subscription to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SubscriptionList200Response**](SubscriptionList200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubscriptionUpdate

> SubscriptionList200Response SubscriptionUpdate(ctx, id).SubscriptionUpdateRequest(subscriptionUpdateRequest).Execute()

Update a Subscription



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
	id := int32(56) // int32 | The id of the subscription to retrieve.
	subscriptionUpdateRequest := *openapiclient.NewSubscriptionUpdateRequest() // SubscriptionUpdateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SubscriptionAPI.SubscriptionUpdate(context.Background(), id).SubscriptionUpdateRequest(subscriptionUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionAPI.SubscriptionUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubscriptionUpdate`: SubscriptionList200Response
	fmt.Fprintf(os.Stdout, "Response from `SubscriptionAPI.SubscriptionUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | The id of the subscription to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscriptionUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **subscriptionUpdateRequest** | [**SubscriptionUpdateRequest**](SubscriptionUpdateRequest.md) |  | 

### Return type

[**SubscriptionList200Response**](SubscriptionList200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

