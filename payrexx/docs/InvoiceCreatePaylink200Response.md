# InvoiceCreatePaylink200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**Data** | Pointer to [**[]Paylink**](Paylink.md) |  | [optional] 

## Methods

### NewInvoiceCreatePaylink200Response

`func NewInvoiceCreatePaylink200Response() *InvoiceCreatePaylink200Response`

NewInvoiceCreatePaylink200Response instantiates a new InvoiceCreatePaylink200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceCreatePaylink200ResponseWithDefaults

`func NewInvoiceCreatePaylink200ResponseWithDefaults() *InvoiceCreatePaylink200Response`

NewInvoiceCreatePaylink200ResponseWithDefaults instantiates a new InvoiceCreatePaylink200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *InvoiceCreatePaylink200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *InvoiceCreatePaylink200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *InvoiceCreatePaylink200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *InvoiceCreatePaylink200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetData

`func (o *InvoiceCreatePaylink200Response) GetData() []Paylink`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *InvoiceCreatePaylink200Response) GetDataOk() (*[]Paylink, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *InvoiceCreatePaylink200Response) SetData(v []Paylink)`

SetData sets Data field to given value.

### HasData

`func (o *InvoiceCreatePaylink200Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


