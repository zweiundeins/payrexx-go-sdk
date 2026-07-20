# SubscriptionInvoice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Number** | Pointer to **string** |  | [optional] 
**Amount** | Pointer to **int32** |  | [optional] [default to 0]
**Currency** | Pointer to **string** |  | [optional] 
**ReferenceId** | Pointer to **string** |  | [optional] 
**PaymentRequestId** | Pointer to **interface{}** |  | [optional] 
**PaymentLink** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewSubscriptionInvoice

`func NewSubscriptionInvoice() *SubscriptionInvoice`

NewSubscriptionInvoice instantiates a new SubscriptionInvoice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionInvoiceWithDefaults

`func NewSubscriptionInvoiceWithDefaults() *SubscriptionInvoice`

NewSubscriptionInvoiceWithDefaults instantiates a new SubscriptionInvoice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNumber

`func (o *SubscriptionInvoice) GetNumber() string`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *SubscriptionInvoice) GetNumberOk() (*string, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *SubscriptionInvoice) SetNumber(v string)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *SubscriptionInvoice) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetAmount

`func (o *SubscriptionInvoice) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *SubscriptionInvoice) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *SubscriptionInvoice) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *SubscriptionInvoice) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCurrency

`func (o *SubscriptionInvoice) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *SubscriptionInvoice) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *SubscriptionInvoice) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *SubscriptionInvoice) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetReferenceId

`func (o *SubscriptionInvoice) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *SubscriptionInvoice) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *SubscriptionInvoice) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *SubscriptionInvoice) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.

### GetPaymentRequestId

`func (o *SubscriptionInvoice) GetPaymentRequestId() interface{}`

GetPaymentRequestId returns the PaymentRequestId field if non-nil, zero value otherwise.

### GetPaymentRequestIdOk

`func (o *SubscriptionInvoice) GetPaymentRequestIdOk() (*interface{}, bool)`

GetPaymentRequestIdOk returns a tuple with the PaymentRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentRequestId

`func (o *SubscriptionInvoice) SetPaymentRequestId(v interface{})`

SetPaymentRequestId sets PaymentRequestId field to given value.

### HasPaymentRequestId

`func (o *SubscriptionInvoice) HasPaymentRequestId() bool`

HasPaymentRequestId returns a boolean if a field has been set.

### SetPaymentRequestIdNil

`func (o *SubscriptionInvoice) SetPaymentRequestIdNil(b bool)`

 SetPaymentRequestIdNil sets the value for PaymentRequestId to be an explicit nil

### UnsetPaymentRequestId
`func (o *SubscriptionInvoice) UnsetPaymentRequestId()`

UnsetPaymentRequestId ensures that no value is present for PaymentRequestId, not even an explicit nil
### GetPaymentLink

`func (o *SubscriptionInvoice) GetPaymentLink() interface{}`

GetPaymentLink returns the PaymentLink field if non-nil, zero value otherwise.

### GetPaymentLinkOk

`func (o *SubscriptionInvoice) GetPaymentLinkOk() (*interface{}, bool)`

GetPaymentLinkOk returns a tuple with the PaymentLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentLink

`func (o *SubscriptionInvoice) SetPaymentLink(v interface{})`

SetPaymentLink sets PaymentLink field to given value.

### HasPaymentLink

`func (o *SubscriptionInvoice) HasPaymentLink() bool`

HasPaymentLink returns a boolean if a field has been set.

### SetPaymentLinkNil

`func (o *SubscriptionInvoice) SetPaymentLinkNil(b bool)`

 SetPaymentLinkNil sets the value for PaymentLink to be an explicit nil

### UnsetPaymentLink
`func (o *SubscriptionInvoice) UnsetPaymentLink()`

UnsetPaymentLink ensures that no value is present for PaymentLink, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


