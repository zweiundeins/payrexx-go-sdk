# EcrPaymentInitiateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int32** | in cents | 
**Currency** | **string** |  | 
**PaymentMethod** | Pointer to **string** |  | [optional] 
**PaymentReference** | Pointer to **string** |  | [optional] 
**PrintSlip** | Pointer to **bool** |  | [optional] 
**TipAmount** | Pointer to **string** | added to the total | [optional] 
**ShopItems** | Pointer to [**EcrPaymentInitiateRequestShopItems**](EcrPaymentInitiateRequestShopItems.md) |  | [optional] 

## Methods

### NewEcrPaymentInitiateRequest

`func NewEcrPaymentInitiateRequest(amount int32, currency string, ) *EcrPaymentInitiateRequest`

NewEcrPaymentInitiateRequest instantiates a new EcrPaymentInitiateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEcrPaymentInitiateRequestWithDefaults

`func NewEcrPaymentInitiateRequestWithDefaults() *EcrPaymentInitiateRequest`

NewEcrPaymentInitiateRequestWithDefaults instantiates a new EcrPaymentInitiateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *EcrPaymentInitiateRequest) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *EcrPaymentInitiateRequest) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *EcrPaymentInitiateRequest) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetCurrency

`func (o *EcrPaymentInitiateRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *EcrPaymentInitiateRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *EcrPaymentInitiateRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetPaymentMethod

`func (o *EcrPaymentInitiateRequest) GetPaymentMethod() string`

GetPaymentMethod returns the PaymentMethod field if non-nil, zero value otherwise.

### GetPaymentMethodOk

`func (o *EcrPaymentInitiateRequest) GetPaymentMethodOk() (*string, bool)`

GetPaymentMethodOk returns a tuple with the PaymentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethod

`func (o *EcrPaymentInitiateRequest) SetPaymentMethod(v string)`

SetPaymentMethod sets PaymentMethod field to given value.

### HasPaymentMethod

`func (o *EcrPaymentInitiateRequest) HasPaymentMethod() bool`

HasPaymentMethod returns a boolean if a field has been set.

### GetPaymentReference

`func (o *EcrPaymentInitiateRequest) GetPaymentReference() string`

GetPaymentReference returns the PaymentReference field if non-nil, zero value otherwise.

### GetPaymentReferenceOk

`func (o *EcrPaymentInitiateRequest) GetPaymentReferenceOk() (*string, bool)`

GetPaymentReferenceOk returns a tuple with the PaymentReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentReference

`func (o *EcrPaymentInitiateRequest) SetPaymentReference(v string)`

SetPaymentReference sets PaymentReference field to given value.

### HasPaymentReference

`func (o *EcrPaymentInitiateRequest) HasPaymentReference() bool`

HasPaymentReference returns a boolean if a field has been set.

### GetPrintSlip

`func (o *EcrPaymentInitiateRequest) GetPrintSlip() bool`

GetPrintSlip returns the PrintSlip field if non-nil, zero value otherwise.

### GetPrintSlipOk

`func (o *EcrPaymentInitiateRequest) GetPrintSlipOk() (*bool, bool)`

GetPrintSlipOk returns a tuple with the PrintSlip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrintSlip

`func (o *EcrPaymentInitiateRequest) SetPrintSlip(v bool)`

SetPrintSlip sets PrintSlip field to given value.

### HasPrintSlip

`func (o *EcrPaymentInitiateRequest) HasPrintSlip() bool`

HasPrintSlip returns a boolean if a field has been set.

### GetTipAmount

`func (o *EcrPaymentInitiateRequest) GetTipAmount() string`

GetTipAmount returns the TipAmount field if non-nil, zero value otherwise.

### GetTipAmountOk

`func (o *EcrPaymentInitiateRequest) GetTipAmountOk() (*string, bool)`

GetTipAmountOk returns a tuple with the TipAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTipAmount

`func (o *EcrPaymentInitiateRequest) SetTipAmount(v string)`

SetTipAmount sets TipAmount field to given value.

### HasTipAmount

`func (o *EcrPaymentInitiateRequest) HasTipAmount() bool`

HasTipAmount returns a boolean if a field has been set.

### GetShopItems

`func (o *EcrPaymentInitiateRequest) GetShopItems() EcrPaymentInitiateRequestShopItems`

GetShopItems returns the ShopItems field if non-nil, zero value otherwise.

### GetShopItemsOk

`func (o *EcrPaymentInitiateRequest) GetShopItemsOk() (*EcrPaymentInitiateRequestShopItems, bool)`

GetShopItemsOk returns a tuple with the ShopItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopItems

`func (o *EcrPaymentInitiateRequest) SetShopItems(v EcrPaymentInitiateRequestShopItems)`

SetShopItems sets ShopItems field to given value.

### HasShopItems

`func (o *EcrPaymentInitiateRequest) HasShopItems() bool`

HasShopItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


