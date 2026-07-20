# \EcrAPI

All URIs are relative to *https://api.payrexx.com/v1.16*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EcrPair**](EcrAPI.md#EcrPair) | **Post** /ecr/{serialNumber}/pair | 
[**EcrPairingRetrieve**](EcrAPI.md#EcrPairingRetrieve) | **Get** /ecr/{serialNumber}/pair | 
[**EcrPaymentCancel**](EcrAPI.md#EcrPaymentCancel) | **Post** /ecr/{serialNumber}/payment/{paymentId}/cancel | 
[**EcrPaymentInitiate**](EcrAPI.md#EcrPaymentInitiate) | **Post** /ecr/{serialNumber}/payment | 
[**EcrPaymentMethodList**](EcrAPI.md#EcrPaymentMethodList) | **Get** /ecr/{serialNumber}/paymentMethods | 
[**EcrPaymentRetrieve**](EcrAPI.md#EcrPaymentRetrieve) | **Get** /ecr/{serialNumber}/payment/{paymentId} | 
[**EcrPaymentVoid**](EcrAPI.md#EcrPaymentVoid) | **Post** /ecr/{serialNumber}/payment/{paymentId}/void | 
[**EcrUnpair**](EcrAPI.md#EcrUnpair) | **Delete** /ecr/{serialNumber}/pair | 



## EcrPair

> EcrPair200Response EcrPair(ctx, serialNumber).EcrPairRequest(ecrPairRequest).Execute()





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
	serialNumber := "serialNumber_example" // string | 
	ecrPairRequest := *openapiclient.NewEcrPairRequest("PairingCode_example") // EcrPairRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EcrAPI.EcrPair(context.Background(), serialNumber).EcrPairRequest(ecrPairRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EcrAPI.EcrPair``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EcrPair`: EcrPair200Response
	fmt.Fprintf(os.Stdout, "Response from `EcrAPI.EcrPair`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serialNumber** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEcrPairRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ecrPairRequest** | [**EcrPairRequest**](EcrPairRequest.md) |  | 

### Return type

[**EcrPair200Response**](EcrPair200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EcrPairingRetrieve

> EcrPairingRetrieve200Response EcrPairingRetrieve(ctx, serialNumber).Execute()





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
	serialNumber := "serialNumber_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EcrAPI.EcrPairingRetrieve(context.Background(), serialNumber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EcrAPI.EcrPairingRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EcrPairingRetrieve`: EcrPairingRetrieve200Response
	fmt.Fprintf(os.Stdout, "Response from `EcrAPI.EcrPairingRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serialNumber** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEcrPairingRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EcrPairingRetrieve200Response**](EcrPairingRetrieve200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EcrPaymentCancel

> EcrPaymentInitiate200Response EcrPaymentCancel(ctx, serialNumber, paymentId).Execute()





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
	serialNumber := "serialNumber_example" // string | 
	paymentId := "paymentId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EcrAPI.EcrPaymentCancel(context.Background(), serialNumber, paymentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EcrAPI.EcrPaymentCancel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EcrPaymentCancel`: EcrPaymentInitiate200Response
	fmt.Fprintf(os.Stdout, "Response from `EcrAPI.EcrPaymentCancel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serialNumber** | **string** |  | 
**paymentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEcrPaymentCancelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**EcrPaymentInitiate200Response**](EcrPaymentInitiate200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EcrPaymentInitiate

> EcrPaymentInitiate200Response EcrPaymentInitiate(ctx, serialNumber).EcrPaymentInitiateRequest(ecrPaymentInitiateRequest).Execute()





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
	serialNumber := "serialNumber_example" // string | 
	ecrPaymentInitiateRequest := *openapiclient.NewEcrPaymentInitiateRequest(int32(123), "Currency_example") // EcrPaymentInitiateRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EcrAPI.EcrPaymentInitiate(context.Background(), serialNumber).EcrPaymentInitiateRequest(ecrPaymentInitiateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EcrAPI.EcrPaymentInitiate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EcrPaymentInitiate`: EcrPaymentInitiate200Response
	fmt.Fprintf(os.Stdout, "Response from `EcrAPI.EcrPaymentInitiate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serialNumber** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEcrPaymentInitiateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ecrPaymentInitiateRequest** | [**EcrPaymentInitiateRequest**](EcrPaymentInitiateRequest.md) |  | 

### Return type

[**EcrPaymentInitiate200Response**](EcrPaymentInitiate200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EcrPaymentMethodList

> EcrPaymentMethodList200Response EcrPaymentMethodList(ctx, serialNumber).Execute()





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
	serialNumber := "serialNumber_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EcrAPI.EcrPaymentMethodList(context.Background(), serialNumber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EcrAPI.EcrPaymentMethodList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EcrPaymentMethodList`: EcrPaymentMethodList200Response
	fmt.Fprintf(os.Stdout, "Response from `EcrAPI.EcrPaymentMethodList`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serialNumber** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEcrPaymentMethodListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EcrPaymentMethodList200Response**](EcrPaymentMethodList200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EcrPaymentRetrieve

> EcrPaymentRetrieve200Response EcrPaymentRetrieve(ctx, serialNumber, paymentId).Execute()





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
	serialNumber := "serialNumber_example" // string | 
	paymentId := "paymentId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EcrAPI.EcrPaymentRetrieve(context.Background(), serialNumber, paymentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EcrAPI.EcrPaymentRetrieve``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EcrPaymentRetrieve`: EcrPaymentRetrieve200Response
	fmt.Fprintf(os.Stdout, "Response from `EcrAPI.EcrPaymentRetrieve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serialNumber** | **string** |  | 
**paymentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEcrPaymentRetrieveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**EcrPaymentRetrieve200Response**](EcrPaymentRetrieve200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EcrPaymentVoid

> EcrPaymentInitiate200Response EcrPaymentVoid(ctx, serialNumber, paymentId).Execute()





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
	serialNumber := "serialNumber_example" // string | 
	paymentId := "paymentId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EcrAPI.EcrPaymentVoid(context.Background(), serialNumber, paymentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EcrAPI.EcrPaymentVoid``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EcrPaymentVoid`: EcrPaymentInitiate200Response
	fmt.Fprintf(os.Stdout, "Response from `EcrAPI.EcrPaymentVoid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serialNumber** | **string** |  | 
**paymentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEcrPaymentVoidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**EcrPaymentInitiate200Response**](EcrPaymentInitiate200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EcrUnpair

> EcrUnpair200Response EcrUnpair(ctx, serialNumber).Execute()





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
	serialNumber := "serialNumber_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EcrAPI.EcrUnpair(context.Background(), serialNumber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EcrAPI.EcrUnpair``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EcrUnpair`: EcrUnpair200Response
	fmt.Fprintf(os.Stdout, "Response from `EcrAPI.EcrUnpair`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serialNumber** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiEcrUnpairRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EcrUnpair200Response**](EcrUnpair200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

