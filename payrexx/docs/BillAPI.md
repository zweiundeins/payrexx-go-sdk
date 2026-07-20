# \BillAPI

All URIs are relative to *https://api.payrexx.com/v1.16*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BillCreate**](BillAPI.md#BillCreate) | **Post** /Bill/ | Create an Invoice
[**BillDelete**](BillAPI.md#BillDelete) | **Delete** /Bill/{id}/ | Delete an Invoice
[**BillList**](BillAPI.md#BillList) | **Get** /Bill/ | List all Invoices
[**BillRetrieve**](BillAPI.md#BillRetrieve) | **Get** /Bill/{id}/ | Retrieve an Invoice
[**BillUpdate**](BillAPI.md#BillUpdate) | **Patch** /Bill/{id}/ | Update an Invoice



## BillCreate

> BillCreate200Response BillCreate(ctx).BillCreateRequest(billCreateRequest).Execute()

Create an Invoice



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
	billCreateRequest := *openapiclient.NewBillCreateRequest("Language_example", "Currency_example", int32(123), *openapiclient.NewBillCreateRequestRecipient("Email_example"), []openapiclient.BillCreateRequestPositionsInner{*openapiclient.NewBillCreateRequestPositionsInner("Title_example", "Type_example", int32(123))}) // BillCreateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillAPI.BillCreate(context.Background()).BillCreateRequest(billCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillAPI.BillCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BillCreate`: BillCreate200Response
	fmt.Fprintf(os.Stdout, "Response from `BillAPI.BillCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBillCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **billCreateRequest** | [**BillCreateRequest**](BillCreateRequest.md) |  | 

### Return type

[**BillCreate200Response**](BillCreate200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillDelete

> BillDelete200Response BillDelete(ctx, id).Execute()

Delete an Invoice



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
	id := "id_example" // string | ID of invoice to retrieve.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillAPI.BillDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillAPI.BillDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BillDelete`: BillDelete200Response
	fmt.Fprintf(os.Stdout, "Response from `BillAPI.BillDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of invoice to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBillDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BillDelete200Response**](BillDelete200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillList

> BillList200Response BillList(ctx).Execute()

List all Invoices



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
	resp, r, err := apiClient.BillAPI.BillList(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillAPI.BillList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BillList`: BillList200Response
	fmt.Fprintf(os.Stdout, "Response from `BillAPI.BillList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiBillListRequest struct via the builder pattern


### Return type

[**BillList200Response**](BillList200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillRetrieve

> BillRetrieve200Response BillRetrieve(ctx, id).Execute()

Retrieve an Invoice



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
	id := "id_example" // string | ID of invoice to retrieve.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillAPI.BillRetrieve(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillAPI.BillRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BillRetrieve`: BillRetrieve200Response
	fmt.Fprintf(os.Stdout, "Response from `BillAPI.BillRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of invoice to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBillRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BillRetrieve200Response**](BillRetrieve200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BillUpdate

> BillCreate200Response BillUpdate(ctx, id).BillUpdateRequest(billUpdateRequest).Execute()

Update an Invoice



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
	id := "id_example" // string | ID of invoice to retrieve.
	billUpdateRequest := *openapiclient.NewBillUpdateRequest() // BillUpdateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BillAPI.BillUpdate(context.Background(), id).BillUpdateRequest(billUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BillAPI.BillUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BillUpdate`: BillCreate200Response
	fmt.Fprintf(os.Stdout, "Response from `BillAPI.BillUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | ID of invoice to retrieve. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBillUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **billUpdateRequest** | [**BillUpdateRequest**](BillUpdateRequest.md) |  | 

### Return type

[**BillCreate200Response**](BillCreate200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

