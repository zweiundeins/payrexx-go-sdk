# TransactionInvoice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrencyAlpha3** | Pointer to **string** |  | [optional] 
**Products** | Pointer to [**[]TransactionInvoiceProductsInner**](TransactionInvoiceProductsInner.md) |  | [optional] 
**Discount** | Pointer to [**TransactionInvoiceDiscount**](TransactionInvoiceDiscount.md) |  | [optional] 
**ShippingAmount** | Pointer to **int32** |  | [optional] [default to 0]
**TotalAmount** | Pointer to **int32** |  | [optional] [default to 0]
**CustomFields** | Pointer to [**TransactionInvoiceCustomFields**](TransactionInvoiceCustomFields.md) |  | [optional] 

## Methods

### NewTransactionInvoice

`func NewTransactionInvoice() *TransactionInvoice`

NewTransactionInvoice instantiates a new TransactionInvoice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionInvoiceWithDefaults

`func NewTransactionInvoiceWithDefaults() *TransactionInvoice`

NewTransactionInvoiceWithDefaults instantiates a new TransactionInvoice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrencyAlpha3

`func (o *TransactionInvoice) GetCurrencyAlpha3() string`

GetCurrencyAlpha3 returns the CurrencyAlpha3 field if non-nil, zero value otherwise.

### GetCurrencyAlpha3Ok

`func (o *TransactionInvoice) GetCurrencyAlpha3Ok() (*string, bool)`

GetCurrencyAlpha3Ok returns a tuple with the CurrencyAlpha3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrencyAlpha3

`func (o *TransactionInvoice) SetCurrencyAlpha3(v string)`

SetCurrencyAlpha3 sets CurrencyAlpha3 field to given value.

### HasCurrencyAlpha3

`func (o *TransactionInvoice) HasCurrencyAlpha3() bool`

HasCurrencyAlpha3 returns a boolean if a field has been set.

### GetProducts

`func (o *TransactionInvoice) GetProducts() []TransactionInvoiceProductsInner`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *TransactionInvoice) GetProductsOk() (*[]TransactionInvoiceProductsInner, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *TransactionInvoice) SetProducts(v []TransactionInvoiceProductsInner)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *TransactionInvoice) HasProducts() bool`

HasProducts returns a boolean if a field has been set.

### GetDiscount

`func (o *TransactionInvoice) GetDiscount() TransactionInvoiceDiscount`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *TransactionInvoice) GetDiscountOk() (*TransactionInvoiceDiscount, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *TransactionInvoice) SetDiscount(v TransactionInvoiceDiscount)`

SetDiscount sets Discount field to given value.

### HasDiscount

`func (o *TransactionInvoice) HasDiscount() bool`

HasDiscount returns a boolean if a field has been set.

### GetShippingAmount

`func (o *TransactionInvoice) GetShippingAmount() int32`

GetShippingAmount returns the ShippingAmount field if non-nil, zero value otherwise.

### GetShippingAmountOk

`func (o *TransactionInvoice) GetShippingAmountOk() (*int32, bool)`

GetShippingAmountOk returns a tuple with the ShippingAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAmount

`func (o *TransactionInvoice) SetShippingAmount(v int32)`

SetShippingAmount sets ShippingAmount field to given value.

### HasShippingAmount

`func (o *TransactionInvoice) HasShippingAmount() bool`

HasShippingAmount returns a boolean if a field has been set.

### GetTotalAmount

`func (o *TransactionInvoice) GetTotalAmount() int32`

GetTotalAmount returns the TotalAmount field if non-nil, zero value otherwise.

### GetTotalAmountOk

`func (o *TransactionInvoice) GetTotalAmountOk() (*int32, bool)`

GetTotalAmountOk returns a tuple with the TotalAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalAmount

`func (o *TransactionInvoice) SetTotalAmount(v int32)`

SetTotalAmount sets TotalAmount field to given value.

### HasTotalAmount

`func (o *TransactionInvoice) HasTotalAmount() bool`

HasTotalAmount returns a boolean if a field has been set.

### GetCustomFields

`func (o *TransactionInvoice) GetCustomFields() TransactionInvoiceCustomFields`

GetCustomFields returns the CustomFields field if non-nil, zero value otherwise.

### GetCustomFieldsOk

`func (o *TransactionInvoice) GetCustomFieldsOk() (*TransactionInvoiceCustomFields, bool)`

GetCustomFieldsOk returns a tuple with the CustomFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFields

`func (o *TransactionInvoice) SetCustomFields(v TransactionInvoiceCustomFields)`

SetCustomFields sets CustomFields field to given value.

### HasCustomFields

`func (o *TransactionInvoice) HasCustomFields() bool`

HasCustomFields returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


