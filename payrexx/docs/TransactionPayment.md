# TransactionPayment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Brand** | Pointer to **string** |  | [optional] 
**Wallet** | Pointer to **interface{}** |  | [optional] 
**PurchaseOnInvoiceInformation** | Pointer to [**TransactionPaymentPurchaseOnInvoiceInformation**](TransactionPaymentPurchaseOnInvoiceInformation.md) |  | [optional] 

## Methods

### NewTransactionPayment

`func NewTransactionPayment() *TransactionPayment`

NewTransactionPayment instantiates a new TransactionPayment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionPaymentWithDefaults

`func NewTransactionPaymentWithDefaults() *TransactionPayment`

NewTransactionPaymentWithDefaults instantiates a new TransactionPayment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrand

`func (o *TransactionPayment) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *TransactionPayment) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *TransactionPayment) SetBrand(v string)`

SetBrand sets Brand field to given value.

### HasBrand

`func (o *TransactionPayment) HasBrand() bool`

HasBrand returns a boolean if a field has been set.

### GetWallet

`func (o *TransactionPayment) GetWallet() interface{}`

GetWallet returns the Wallet field if non-nil, zero value otherwise.

### GetWalletOk

`func (o *TransactionPayment) GetWalletOk() (*interface{}, bool)`

GetWalletOk returns a tuple with the Wallet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWallet

`func (o *TransactionPayment) SetWallet(v interface{})`

SetWallet sets Wallet field to given value.

### HasWallet

`func (o *TransactionPayment) HasWallet() bool`

HasWallet returns a boolean if a field has been set.

### SetWalletNil

`func (o *TransactionPayment) SetWalletNil(b bool)`

 SetWalletNil sets the value for Wallet to be an explicit nil

### UnsetWallet
`func (o *TransactionPayment) UnsetWallet()`

UnsetWallet ensures that no value is present for Wallet, not even an explicit nil
### GetPurchaseOnInvoiceInformation

`func (o *TransactionPayment) GetPurchaseOnInvoiceInformation() TransactionPaymentPurchaseOnInvoiceInformation`

GetPurchaseOnInvoiceInformation returns the PurchaseOnInvoiceInformation field if non-nil, zero value otherwise.

### GetPurchaseOnInvoiceInformationOk

`func (o *TransactionPayment) GetPurchaseOnInvoiceInformationOk() (*TransactionPaymentPurchaseOnInvoiceInformation, bool)`

GetPurchaseOnInvoiceInformationOk returns a tuple with the PurchaseOnInvoiceInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseOnInvoiceInformation

`func (o *TransactionPayment) SetPurchaseOnInvoiceInformation(v TransactionPaymentPurchaseOnInvoiceInformation)`

SetPurchaseOnInvoiceInformation sets PurchaseOnInvoiceInformation field to given value.

### HasPurchaseOnInvoiceInformation

`func (o *TransactionPayment) HasPurchaseOnInvoiceInformation() bool`

HasPurchaseOnInvoiceInformation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


