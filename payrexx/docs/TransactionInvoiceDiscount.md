# TransactionInvoiceDiscount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** |  | [optional] 
**Percentage** | Pointer to **int32** |  | [optional] [default to 0]
**Amount** | Pointer to **int32** |  | [optional] [default to 0]

## Methods

### NewTransactionInvoiceDiscount

`func NewTransactionInvoiceDiscount() *TransactionInvoiceDiscount`

NewTransactionInvoiceDiscount instantiates a new TransactionInvoiceDiscount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionInvoiceDiscountWithDefaults

`func NewTransactionInvoiceDiscountWithDefaults() *TransactionInvoiceDiscount`

NewTransactionInvoiceDiscountWithDefaults instantiates a new TransactionInvoiceDiscount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *TransactionInvoiceDiscount) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *TransactionInvoiceDiscount) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *TransactionInvoiceDiscount) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *TransactionInvoiceDiscount) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetPercentage

`func (o *TransactionInvoiceDiscount) GetPercentage() int32`

GetPercentage returns the Percentage field if non-nil, zero value otherwise.

### GetPercentageOk

`func (o *TransactionInvoiceDiscount) GetPercentageOk() (*int32, bool)`

GetPercentageOk returns a tuple with the Percentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentage

`func (o *TransactionInvoiceDiscount) SetPercentage(v int32)`

SetPercentage sets Percentage field to given value.

### HasPercentage

`func (o *TransactionInvoiceDiscount) HasPercentage() bool`

HasPercentage returns a boolean if a field has been set.

### GetAmount

`func (o *TransactionInvoiceDiscount) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TransactionInvoiceDiscount) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TransactionInvoiceDiscount) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *TransactionInvoiceDiscount) HasAmount() bool`

HasAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


