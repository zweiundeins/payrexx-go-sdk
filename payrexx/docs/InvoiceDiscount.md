# InvoiceDiscount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int32** |  | [optional] [default to 0]
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewInvoiceDiscount

`func NewInvoiceDiscount() *InvoiceDiscount`

NewInvoiceDiscount instantiates a new InvoiceDiscount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceDiscountWithDefaults

`func NewInvoiceDiscountWithDefaults() *InvoiceDiscount`

NewInvoiceDiscountWithDefaults instantiates a new InvoiceDiscount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *InvoiceDiscount) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *InvoiceDiscount) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *InvoiceDiscount) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *InvoiceDiscount) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetType

`func (o *InvoiceDiscount) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InvoiceDiscount) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InvoiceDiscount) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InvoiceDiscount) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


