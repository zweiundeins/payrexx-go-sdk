# \PaymentProviderAPI

All URIs are relative to *https://api.payrexx.com/v1.16*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PaymentProviderList**](PaymentProviderAPI.md#PaymentProviderList) | **Get** /PaymentProvider/ | Retrieve all Payment Providers



## PaymentProviderList

> PaymentProviderList200Response PaymentProviderList(ctx).Execute()

Retrieve all Payment Providers



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
	resp, r, err := apiClient.PaymentProviderAPI.PaymentProviderList(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentProviderAPI.PaymentProviderList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PaymentProviderList`: PaymentProviderList200Response
	fmt.Fprintf(os.Stdout, "Response from `PaymentProviderAPI.PaymentProviderList`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPaymentProviderListRequest struct via the builder pattern


### Return type

[**PaymentProviderList200Response**](PaymentProviderList200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

