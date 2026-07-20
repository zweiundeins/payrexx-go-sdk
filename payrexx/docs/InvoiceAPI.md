# \InvoiceAPI

All URIs are relative to *https://api.payrexx.com/v1.16*

Method | HTTP request | Description
------------- | ------------- | -------------
[**InvoiceCreatePaylink**](InvoiceAPI.md#InvoiceCreatePaylink) | **Post** /Invoice/ | Create a Paylink



## InvoiceCreatePaylink

> InvoiceCreatePaylink200Response InvoiceCreatePaylink(ctx).InvoiceCreatePaylinkRequest(invoiceCreatePaylinkRequest).Execute()

Create a Paylink



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
	invoiceCreatePaylinkRequest := *openapiclient.NewInvoiceCreatePaylinkRequest("Title_example", "Description_example", "ReferenceId_example", "Purpose_example", int32(123), "Currency_example") // InvoiceCreatePaylinkRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.InvoiceAPI.InvoiceCreatePaylink(context.Background()).InvoiceCreatePaylinkRequest(invoiceCreatePaylinkRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `InvoiceAPI.InvoiceCreatePaylink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `InvoiceCreatePaylink`: InvoiceCreatePaylink200Response
	fmt.Fprintf(os.Stdout, "Response from `InvoiceAPI.InvoiceCreatePaylink`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInvoiceCreatePaylinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **invoiceCreatePaylinkRequest** | [**InvoiceCreatePaylinkRequest**](InvoiceCreatePaylinkRequest.md) |  | 

### Return type

[**InvoiceCreatePaylink200Response**](InvoiceCreatePaylink200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

