# \TransactionAPI

All URIs are relative to *https://api.payrexx.com/v1.16*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TransactionCancel**](TransactionAPI.md#TransactionCancel) | **Patch** /Transaction/{id}/cancel | Cancel a Waiting Transaction
[**TransactionList**](TransactionAPI.md#TransactionList) | **Get** /Transaction/ | Retrieve Transactions
[**TransactionRetrieve**](TransactionAPI.md#TransactionRetrieve) | **Get** /Transaction/{id}/ | Retrieve a Transaction
[**TransactionUpdateContactDetails**](TransactionAPI.md#TransactionUpdateContactDetails) | **Put** /Transaction/{id}/ | Update Pre-Authorization/Tokenization Contact Details
[**TransactionUpdateTokenization**](TransactionAPI.md#TransactionUpdateTokenization) | **Patch** /Transaction/{id}/updateTokenization | Update a Tokenization



## TransactionCancel

> TransactionList200Response TransactionCancel(ctx, id).Execute()

Cancel a Waiting Transaction



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
	id := int32(56) // int32 | ID of transaction to cancel.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionAPI.TransactionCancel(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionAPI.TransactionCancel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TransactionCancel`: TransactionList200Response
	fmt.Fprintf(os.Stdout, "Response from `TransactionAPI.TransactionCancel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID of transaction to cancel. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTransactionCancelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TransactionList200Response**](TransactionList200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TransactionList

> TransactionList200Response TransactionList(ctx).TransactionListRequest(transactionListRequest).Execute()

Retrieve Transactions



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
	transactionListRequest := *openapiclient.NewTransactionListRequest() // TransactionListRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionAPI.TransactionList(context.Background()).TransactionListRequest(transactionListRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionAPI.TransactionList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TransactionList`: TransactionList200Response
	fmt.Fprintf(os.Stdout, "Response from `TransactionAPI.TransactionList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTransactionListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transactionListRequest** | [**TransactionListRequest**](TransactionListRequest.md) |  | 

### Return type

[**TransactionList200Response**](TransactionList200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TransactionRetrieve

> TransactionRetrieve(ctx, id).Execute()

Retrieve a Transaction



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
	id := int32(56) // int32 | ID of transaction to retrieve. (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.TransactionAPI.TransactionRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionAPI.TransactionRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID of transaction to retrieve. | [default to 0]

### Other Parameters

Other parameters are passed through a pointer to a apiTransactionRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TransactionUpdateContactDetails

> TransactionList200Response TransactionUpdateContactDetails(ctx, id).TransactionUpdateContactDetailsRequest(transactionUpdateContactDetailsRequest).Execute()

Update Pre-Authorization/Tokenization Contact Details



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
	id := int32(56) // int32 | ID of the Pre-Authorization / Tokenization
	transactionUpdateContactDetailsRequest := *openapiclient.NewTransactionUpdateContactDetailsRequest(*openapiclient.NewTransactionUpdateContactDetailsRequestFields()) // TransactionUpdateContactDetailsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionAPI.TransactionUpdateContactDetails(context.Background(), id).TransactionUpdateContactDetailsRequest(transactionUpdateContactDetailsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionAPI.TransactionUpdateContactDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TransactionUpdateContactDetails`: TransactionList200Response
	fmt.Fprintf(os.Stdout, "Response from `TransactionAPI.TransactionUpdateContactDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID of the Pre-Authorization / Tokenization | 

### Other Parameters

Other parameters are passed through a pointer to a apiTransactionUpdateContactDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transactionUpdateContactDetailsRequest** | [**TransactionUpdateContactDetailsRequest**](TransactionUpdateContactDetailsRequest.md) |  | 

### Return type

[**TransactionList200Response**](TransactionList200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TransactionUpdateTokenization

> TransactionList200Response TransactionUpdateTokenization(ctx, id).TransactionUpdateTokenizationRequest(transactionUpdateTokenizationRequest).Execute()

Update a Tokenization



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
	id := int32(56) // int32 | ID of origin Tokenization
	transactionUpdateTokenizationRequest := *openapiclient.NewTransactionUpdateTokenizationRequest(int32(123)) // TransactionUpdateTokenizationRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TransactionAPI.TransactionUpdateTokenization(context.Background(), id).TransactionUpdateTokenizationRequest(transactionUpdateTokenizationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TransactionAPI.TransactionUpdateTokenization``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TransactionUpdateTokenization`: TransactionList200Response
	fmt.Fprintf(os.Stdout, "Response from `TransactionAPI.TransactionUpdateTokenization`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int32** | ID of origin Tokenization | 

### Other Parameters

Other parameters are passed through a pointer to a apiTransactionUpdateTokenizationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **transactionUpdateTokenizationRequest** | [**TransactionUpdateTokenizationRequest**](TransactionUpdateTokenizationRequest.md) |  | 

### Return type

[**TransactionList200Response**](TransactionList200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

