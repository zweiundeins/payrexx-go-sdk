# SubscriptionUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **string** | The new amount in cents which will be payed on the next payment interval. | [optional] 
**Currency** | Pointer to **string** | The ISO code of the currency of the payment. | [optional] 
**Purpose** | Pointer to **string** | The payment purpose. What is the payer paying for? | [optional] 
**VatRate** | Pointer to **string** | VAT Rate Percentage | [optional] 

## Methods

### NewSubscriptionUpdateRequest

`func NewSubscriptionUpdateRequest() *SubscriptionUpdateRequest`

NewSubscriptionUpdateRequest instantiates a new SubscriptionUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionUpdateRequestWithDefaults

`func NewSubscriptionUpdateRequestWithDefaults() *SubscriptionUpdateRequest`

NewSubscriptionUpdateRequestWithDefaults instantiates a new SubscriptionUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *SubscriptionUpdateRequest) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *SubscriptionUpdateRequest) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *SubscriptionUpdateRequest) SetAmount(v string)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *SubscriptionUpdateRequest) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCurrency

`func (o *SubscriptionUpdateRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *SubscriptionUpdateRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *SubscriptionUpdateRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *SubscriptionUpdateRequest) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetPurpose

`func (o *SubscriptionUpdateRequest) GetPurpose() string`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *SubscriptionUpdateRequest) GetPurposeOk() (*string, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *SubscriptionUpdateRequest) SetPurpose(v string)`

SetPurpose sets Purpose field to given value.

### HasPurpose

`func (o *SubscriptionUpdateRequest) HasPurpose() bool`

HasPurpose returns a boolean if a field has been set.

### GetVatRate

`func (o *SubscriptionUpdateRequest) GetVatRate() string`

GetVatRate returns the VatRate field if non-nil, zero value otherwise.

### GetVatRateOk

`func (o *SubscriptionUpdateRequest) GetVatRateOk() (*string, bool)`

GetVatRateOk returns a tuple with the VatRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatRate

`func (o *SubscriptionUpdateRequest) SetVatRate(v string)`

SetVatRate sets VatRate field to given value.

### HasVatRate

`func (o *SubscriptionUpdateRequest) HasVatRate() bool`

HasVatRate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


