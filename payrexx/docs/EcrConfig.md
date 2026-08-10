# EcrConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**Tip** | Pointer to **bool** |  | [optional] 
**PrintReceipt** | Pointer to **bool** |  | [optional] 
**PaymentMethods** | Pointer to **[]string** |  | [optional] 

## Methods

### NewEcrConfig

`func NewEcrConfig() *EcrConfig`

NewEcrConfig instantiates a new EcrConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEcrConfigWithDefaults

`func NewEcrConfigWithDefaults() *EcrConfig`

NewEcrConfigWithDefaults instantiates a new EcrConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EcrConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EcrConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EcrConfig) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EcrConfig) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCurrency

`func (o *EcrConfig) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *EcrConfig) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *EcrConfig) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *EcrConfig) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetTip

`func (o *EcrConfig) GetTip() bool`

GetTip returns the Tip field if non-nil, zero value otherwise.

### GetTipOk

`func (o *EcrConfig) GetTipOk() (*bool, bool)`

GetTipOk returns a tuple with the Tip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTip

`func (o *EcrConfig) SetTip(v bool)`

SetTip sets Tip field to given value.

### HasTip

`func (o *EcrConfig) HasTip() bool`

HasTip returns a boolean if a field has been set.

### GetPrintReceipt

`func (o *EcrConfig) GetPrintReceipt() bool`

GetPrintReceipt returns the PrintReceipt field if non-nil, zero value otherwise.

### GetPrintReceiptOk

`func (o *EcrConfig) GetPrintReceiptOk() (*bool, bool)`

GetPrintReceiptOk returns a tuple with the PrintReceipt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrintReceipt

`func (o *EcrConfig) SetPrintReceipt(v bool)`

SetPrintReceipt sets PrintReceipt field to given value.

### HasPrintReceipt

`func (o *EcrConfig) HasPrintReceipt() bool`

HasPrintReceipt returns a boolean if a field has been set.

### GetPaymentMethods

`func (o *EcrConfig) GetPaymentMethods() []string`

GetPaymentMethods returns the PaymentMethods field if non-nil, zero value otherwise.

### GetPaymentMethodsOk

`func (o *EcrConfig) GetPaymentMethodsOk() (*[]string, bool)`

GetPaymentMethodsOk returns a tuple with the PaymentMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethods

`func (o *EcrConfig) SetPaymentMethods(v []string)`

SetPaymentMethods sets PaymentMethods field to given value.

### HasPaymentMethods

`func (o *EcrConfig) HasPaymentMethods() bool`

HasPaymentMethods returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


