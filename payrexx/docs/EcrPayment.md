# EcrPayment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentId** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** | PAYMENT_REQUESTED | IN_PROGRESS | SUCCESS | DECLINED | UNDERPAID | TERMINATED | REVERTED | EXPIRED | FAILED | UNKNOWN | [optional] 
**Type** | Pointer to **string** | CHARGE or REFUND | [optional] [default to "CHARGE"]
**Amount** | Pointer to **int32** | in cents | [optional] 
**TransactionDescriptor** | Pointer to **string** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**TipAmount** | Pointer to **int32** | in cents | [optional] 
**PaymentMethod** | Pointer to **string** |  | [optional] 
**CardBrand** | Pointer to **string** |  | [optional] 
**PaymentReference** | Pointer to **string** |  | [optional] 
**PrintSlip** | Pointer to **bool** |  | [optional] 
**Terminal** | Pointer to **string** | Terminal serial number | [optional] 
**ReversalStatus** | Pointer to **string** | Raw provider reversal status — present only once a void/refund reversal has been initiated | [optional] 
**ShopItems** | Pointer to [**[]EcrPaymentShopItemsInner**](EcrPaymentShopItemsInner.md) |  | [optional] 
**Slip** | Pointer to **[]string** |  | [optional] 

## Methods

### NewEcrPayment

`func NewEcrPayment() *EcrPayment`

NewEcrPayment instantiates a new EcrPayment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEcrPaymentWithDefaults

`func NewEcrPaymentWithDefaults() *EcrPayment`

NewEcrPaymentWithDefaults instantiates a new EcrPayment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentId

`func (o *EcrPayment) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *EcrPayment) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *EcrPayment) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.

### HasPaymentId

`func (o *EcrPayment) HasPaymentId() bool`

HasPaymentId returns a boolean if a field has been set.

### GetStatus

`func (o *EcrPayment) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EcrPayment) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EcrPayment) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *EcrPayment) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *EcrPayment) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EcrPayment) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EcrPayment) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EcrPayment) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAmount

`func (o *EcrPayment) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *EcrPayment) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *EcrPayment) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *EcrPayment) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetTransactionDescriptor

`func (o *EcrPayment) GetTransactionDescriptor() string`

GetTransactionDescriptor returns the TransactionDescriptor field if non-nil, zero value otherwise.

### GetTransactionDescriptorOk

`func (o *EcrPayment) GetTransactionDescriptorOk() (*string, bool)`

GetTransactionDescriptorOk returns a tuple with the TransactionDescriptor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionDescriptor

`func (o *EcrPayment) SetTransactionDescriptor(v string)`

SetTransactionDescriptor sets TransactionDescriptor field to given value.

### HasTransactionDescriptor

`func (o *EcrPayment) HasTransactionDescriptor() bool`

HasTransactionDescriptor returns a boolean if a field has been set.

### GetCurrency

`func (o *EcrPayment) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *EcrPayment) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *EcrPayment) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *EcrPayment) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetTipAmount

`func (o *EcrPayment) GetTipAmount() int32`

GetTipAmount returns the TipAmount field if non-nil, zero value otherwise.

### GetTipAmountOk

`func (o *EcrPayment) GetTipAmountOk() (*int32, bool)`

GetTipAmountOk returns a tuple with the TipAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTipAmount

`func (o *EcrPayment) SetTipAmount(v int32)`

SetTipAmount sets TipAmount field to given value.

### HasTipAmount

`func (o *EcrPayment) HasTipAmount() bool`

HasTipAmount returns a boolean if a field has been set.

### GetPaymentMethod

`func (o *EcrPayment) GetPaymentMethod() string`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *EcrPayment) GetPaymentMethodOk() (*string, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *EcrPayment) SetPaymentMethod(v string)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *EcrPayment) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.

### GetCardBrand

`func (o *EcrPayment) GetCardBrand() string`

GetCardBrand returns the CardBrand field if non-nil, zero value otherwise.

### GetCardBrandOk

`func (o *EcrPayment) GetCardBrandOk() (*string, bool)`

GetCardBrandOk returns a tuple with the CardBrand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardBrand

`func (o *EcrPayment) SetCardBrand(v string)`

SetCardBrand sets CardBrand field to given value.

### HasCardBrand

`func (o *EcrPayment) HasCardBrand() bool`

HasCardBrand returns a boolean if a field has been set.

### GetPaymentReference

`func (o *EcrPayment) GetPaymentReference() string`

GetPaymentReference returns the PaymentReference field if non-nil, zero value otherwise.

### GetPaymentReferenceOk

`func (o *EcrPayment) GetPaymentReferenceOk() (*string, bool)`

GetPaymentReferenceOk returns a tuple with the PaymentReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentReference

`func (o *EcrPayment) SetPaymentReference(v string)`

SetPaymentReference sets PaymentReference field to given value.

### HasPaymentReference

`func (o *EcrPayment) HasPaymentReference() bool`

HasPaymentReference returns a boolean if a field has been set.

### GetPrintSlip

`func (o *EcrPayment) GetPrintSlip() bool`

GetPrintSlip returns the PrintSlip field if non-nil, zero value otherwise.

### GetPrintSlipOk

`func (o *EcrPayment) GetPrintSlipOk() (*bool, bool)`

GetPrintSlipOk returns a tuple with the PrintSlip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrintSlip

`func (o *EcrPayment) SetPrintSlip(v bool)`

SetPrintSlip sets PrintSlip field to given value.

### HasPrintSlip

`func (o *EcrPayment) HasPrintSlip() bool`

HasPrintSlip returns a boolean if a field has been set.

### GetTerminal

`func (o *EcrPayment) GetTerminal() string`

GetTerminal returns the Terminal field if non-nil, zero value otherwise.

### GetTerminalOk

`func (o *EcrPayment) GetTerminalOk() (*string, bool)`

GetTerminalOk returns a tuple with the Terminal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerminal

`func (o *EcrPayment) SetTerminal(v string)`

SetTerminal sets Terminal field to given value.

### HasTerminal

`func (o *EcrPayment) HasTerminal() bool`

HasTerminal returns a boolean if a field has been set.

### GetReversalStatus

`func (o *EcrPayment) GetReversalStatus() string`

GetReversalStatus returns the ReversalStatus field if non-nil, zero value otherwise.

### GetReversalStatusOk

`func (o *EcrPayment) GetReversalStatusOk() (*string, bool)`

GetReversalStatusOk returns a tuple with the ReversalStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReversalStatus

`func (o *EcrPayment) SetReversalStatus(v string)`

SetReversalStatus sets ReversalStatus field to given value.

### HasReversalStatus

`func (o *EcrPayment) HasReversalStatus() bool`

HasReversalStatus returns a boolean if a field has been set.

### GetShopItems

`func (o *EcrPayment) GetShopItems() []EcrPaymentShopItemsInner`

GetShopItems returns the ShopItems field if non-nil, zero value otherwise.

### GetShopItemsOk

`func (o *EcrPayment) GetShopItemsOk() (*[]EcrPaymentShopItemsInner, bool)`

GetShopItemsOk returns a tuple with the ShopItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopItems

`func (o *EcrPayment) SetShopItems(v []EcrPaymentShopItemsInner)`

SetShopItems sets ShopItems field to given value.

### HasShopItems

`func (o *EcrPayment) HasShopItems() bool`

HasShopItems returns a boolean if a field has been set.

### GetSlip

`func (o *EcrPayment) GetSlip() []string`

GetSlip returns the Slip field if non-nil, zero value otherwise.

### GetSlipOk

`func (o *EcrPayment) GetSlipOk() (*[]string, bool)`

GetSlipOk returns a tuple with the Slip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlip

`func (o *EcrPayment) SetSlip(v []string)`

SetSlip sets Slip field to given value.

### HasSlip

`func (o *EcrPayment) HasSlip() bool`

HasSlip returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


