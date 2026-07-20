# SubscriptionCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserId** | **string** | The contact id you got from webhook. | 
**Psp** | **string** | The ID of the psp to use. | 
**Amount** | **string** | The amount of the payment to fire in cents. | 
**Currency** | **string** | The ISO code of the currency of the payment. | 
**Purpose** | **string** | The payment purpose. What is the payer paying for? | 
**PaymentInterval** | **string** | The payment interval as string. (see PHP documentation of date interval for format) | 
**Period** | **string** | The subscription&#39;s period as string. (see PHP documentation of date interval for format) | 
**CancellationInterval** | **string** | The interval of cancellation as string. (see PHP documentation of date interval for format) | 
**ReferenceId** | Pointer to **string** | The internal reference id. (Will be sent back with Webhook, can be used as identification) | [optional] 
**VatRate** | Pointer to **string** | VAT Rate Percentage | [optional] 

## Methods

### NewSubscriptionCreateRequest

`func NewSubscriptionCreateRequest(userId string, psp string, amount string, currency string, purpose string, paymentInterval string, period string, cancellationInterval string, ) *SubscriptionCreateRequest`

NewSubscriptionCreateRequest instantiates a new SubscriptionCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionCreateRequestWithDefaults

`func NewSubscriptionCreateRequestWithDefaults() *SubscriptionCreateRequest`

NewSubscriptionCreateRequestWithDefaults instantiates a new SubscriptionCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserId

`func (o *SubscriptionCreateRequest) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *SubscriptionCreateRequest) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *SubscriptionCreateRequest) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetPsp

`func (o *SubscriptionCreateRequest) GetPsp() string`

GetPsp returns the Psp field if non-nil, zero value otherwise.

### GetPspOk

`func (o *SubscriptionCreateRequest) GetPspOk() (*string, bool)`

GetPspOk returns a tuple with the Psp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPsp

`func (o *SubscriptionCreateRequest) SetPsp(v string)`

SetPsp sets Psp field to given value.


### GetAmount

`func (o *SubscriptionCreateRequest) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *SubscriptionCreateRequest) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *SubscriptionCreateRequest) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetCurrency

`func (o *SubscriptionCreateRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *SubscriptionCreateRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *SubscriptionCreateRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetPurpose

`func (o *SubscriptionCreateRequest) GetPurpose() string`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *SubscriptionCreateRequest) GetPurposeOk() (*string, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *SubscriptionCreateRequest) SetPurpose(v string)`

SetPurpose sets Purpose field to given value.


### GetPaymentInterval

`func (o *SubscriptionCreateRequest) GetPaymentInterval() string`

GetPaymentInterval returns the PaymentInterval field if non-nil, zero value otherwise.

### GetPaymentIntervalOk

`func (o *SubscriptionCreateRequest) GetPaymentIntervalOk() (*string, bool)`

GetPaymentIntervalOk returns a tuple with the PaymentInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentInterval

`func (o *SubscriptionCreateRequest) SetPaymentInterval(v string)`

SetPaymentInterval sets PaymentInterval field to given value.


### GetPeriod

`func (o *SubscriptionCreateRequest) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *SubscriptionCreateRequest) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *SubscriptionCreateRequest) SetPeriod(v string)`

SetPeriod sets Period field to given value.


### GetCancellationInterval

`func (o *SubscriptionCreateRequest) GetCancellationInterval() string`

GetCancellationInterval returns the CancellationInterval field if non-nil, zero value otherwise.

### GetCancellationIntervalOk

`func (o *SubscriptionCreateRequest) GetCancellationIntervalOk() (*string, bool)`

GetCancellationIntervalOk returns a tuple with the CancellationInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancellationInterval

`func (o *SubscriptionCreateRequest) SetCancellationInterval(v string)`

SetCancellationInterval sets CancellationInterval field to given value.


### GetReferenceId

`func (o *SubscriptionCreateRequest) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *SubscriptionCreateRequest) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *SubscriptionCreateRequest) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *SubscriptionCreateRequest) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.

### GetVatRate

`func (o *SubscriptionCreateRequest) GetVatRate() string`

GetVatRate returns the VatRate field if non-nil, zero value otherwise.

### GetVatRateOk

`func (o *SubscriptionCreateRequest) GetVatRateOk() (*string, bool)`

GetVatRateOk returns a tuple with the VatRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatRate

`func (o *SubscriptionCreateRequest) SetVatRate(v string)`

SetVatRate sets VatRate field to given value.

### HasVatRate

`func (o *SubscriptionCreateRequest) HasVatRate() bool`

HasVatRate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


