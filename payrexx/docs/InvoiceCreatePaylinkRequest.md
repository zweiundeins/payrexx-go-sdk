# InvoiceCreatePaylinkRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** | This is the page title which will be shown on the payment page. | 
**Description** | **string** | This is a description which will be shown on the payment page. | 
**ReferenceId** | **string** | An internal reference id used by your system. | 
**Purpose** | **string** | The purpose of the payment. | 
**Amount** | **int32** | The amount of the payment in cents. | 
**VatRate** | Pointer to **float32** | VAT rate percentage | [optional] 
**Currency** | **string** | The currency of the payment. | 
**Psp** | Pointer to **[]int32** | The psp which should be used for the payment. (Can be an array of integers.) | [optional] 
**Pm** | Pointer to **[]string** | List of payment mean names to display | [optional] 
**Sku** | Pointer to **string** | Product stock keeping unit | [optional] 
**PreAuthorization** | Pointer to **bool** | Whether charge payment manually at a later date (type authorization). | [optional] [default to false]
**Reservation** | Pointer to **bool** | Whether charge payment manually at a later date (type reservation). | [optional] [default to false]
**Name** | Pointer to **string** | This is an internal name of the payment page. This name will be displayed to the administrator only. | [optional] 
**Fields** | Pointer to **[]string** | The contact data fields which should be displayed. See fields for more | [optional] 
**HideFields** | Pointer to **bool** | Hide the whole contact fields section on invoice page | [optional] [default to false]
**ConcardisOrderId** | Pointer to **string** | Only available for Concardis PSP and if the custom ORDERID option is activated in PSP settings in Payrexx administration. This ORDERID will be transferred to the Payengine. | [optional] 
**ButtonText** | Pointer to **string** | Custom pay button text. | [optional] 
**ExpirationDate** | Pointer to **string** | Expiration date for link. | [optional] 
**SuccessRedirectUrl** | Pointer to **string** | URL to redirect to after successful payment. | [optional] 
**FailedRedirectUrl** | Pointer to **string** | URL to redirect to after failed payment. | [optional] 
**SubscriptionState** | Pointer to **bool** | Defines whether the payment should be handled as subscription. | [optional] [default to false]
**SubscriptionInterval** | Pointer to **string** | Payment interval | [optional] 
**SubscriptionPeriod** | Pointer to **string** | Duration of the subscription | [optional] 
**SubscriptionCancellationInterval** | Pointer to **string** | Defines the period, in which a subscription can be canceled. | [optional] 
**Attachments** | Pointer to ***os.File** | Provide your customers file attachments. | [optional] 
**IsPriceExclusiveVat** | Pointer to **bool** | If set to true the Vat will be added on top of the price | [optional] 

## Methods

### NewInvoiceCreatePaylinkRequest

`func NewInvoiceCreatePaylinkRequest(title string, description string, referenceId string, purpose string, amount int32, currency string, ) *InvoiceCreatePaylinkRequest`

NewInvoiceCreatePaylinkRequest instantiates a new InvoiceCreatePaylinkRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceCreatePaylinkRequestWithDefaults

`func NewInvoiceCreatePaylinkRequestWithDefaults() *InvoiceCreatePaylinkRequest`

NewInvoiceCreatePaylinkRequestWithDefaults instantiates a new InvoiceCreatePaylinkRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *InvoiceCreatePaylinkRequest) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *InvoiceCreatePaylinkRequest) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *InvoiceCreatePaylinkRequest) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *InvoiceCreatePaylinkRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InvoiceCreatePaylinkRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InvoiceCreatePaylinkRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetReferenceId

`func (o *InvoiceCreatePaylinkRequest) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *InvoiceCreatePaylinkRequest) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *InvoiceCreatePaylinkRequest) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.


### GetPurpose

`func (o *InvoiceCreatePaylinkRequest) GetPurpose() string`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *InvoiceCreatePaylinkRequest) GetPurposeOk() (*string, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *InvoiceCreatePaylinkRequest) SetPurpose(v string)`

SetPurpose sets Purpose field to given value.


### GetAmount

`func (o *InvoiceCreatePaylinkRequest) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *InvoiceCreatePaylinkRequest) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *InvoiceCreatePaylinkRequest) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetVatRate

`func (o *InvoiceCreatePaylinkRequest) GetVatRate() float32`

GetVatRate returns the VatRate field if non-nil, zero value otherwise.

### GetVatRateOk

`func (o *InvoiceCreatePaylinkRequest) GetVatRateOk() (*float32, bool)`

GetVatRateOk returns a tuple with the VatRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatRate

`func (o *InvoiceCreatePaylinkRequest) SetVatRate(v float32)`

SetVatRate sets VatRate field to given value.

### HasVatRate

`func (o *InvoiceCreatePaylinkRequest) HasVatRate() bool`

HasVatRate returns a boolean if a field has been set.

### GetCurrency

`func (o *InvoiceCreatePaylinkRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *InvoiceCreatePaylinkRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *InvoiceCreatePaylinkRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetPsp

`func (o *InvoiceCreatePaylinkRequest) GetPsp() []int32`

GetPsp returns the Psp field if non-nil, zero value otherwise.

### GetPspOk

`func (o *InvoiceCreatePaylinkRequest) GetPspOk() (*[]int32, bool)`

GetPspOk returns a tuple with the Psp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPsp

`func (o *InvoiceCreatePaylinkRequest) SetPsp(v []int32)`

SetPsp sets Psp field to given value.

### HasPsp

`func (o *InvoiceCreatePaylinkRequest) HasPsp() bool`

HasPsp returns a boolean if a field has been set.

### GetPm

`func (o *InvoiceCreatePaylinkRequest) GetPm() []string`

GetPm returns the Pm field if non-nil, zero value otherwise.

### GetPmOk

`func (o *InvoiceCreatePaylinkRequest) GetPmOk() (*[]string, bool)`

GetPmOk returns a tuple with the Pm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPm

`func (o *InvoiceCreatePaylinkRequest) SetPm(v []string)`

SetPm sets Pm field to given value.

### HasPm

`func (o *InvoiceCreatePaylinkRequest) HasPm() bool`

HasPm returns a boolean if a field has been set.

### GetSku

`func (o *InvoiceCreatePaylinkRequest) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *InvoiceCreatePaylinkRequest) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *InvoiceCreatePaylinkRequest) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *InvoiceCreatePaylinkRequest) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetPreAuthorization

`func (o *InvoiceCreatePaylinkRequest) GetPreAuthorization() bool`

GetPreAuthorization returns the PreAuthorization field if non-nil, zero value otherwise.

### GetPreAuthorizationOk

`func (o *InvoiceCreatePaylinkRequest) GetPreAuthorizationOk() (*bool, bool)`

GetPreAuthorizationOk returns a tuple with the PreAuthorization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreAuthorization

`func (o *InvoiceCreatePaylinkRequest) SetPreAuthorization(v bool)`

SetPreAuthorization sets PreAuthorization field to given value.

### HasPreAuthorization

`func (o *InvoiceCreatePaylinkRequest) HasPreAuthorization() bool`

HasPreAuthorization returns a boolean if a field has been set.

### GetReservation

`func (o *InvoiceCreatePaylinkRequest) GetReservation() bool`

GetReservation returns the Reservation field if non-nil, zero value otherwise.

### GetReservationOk

`func (o *InvoiceCreatePaylinkRequest) GetReservationOk() (*bool, bool)`

GetReservationOk returns a tuple with the Reservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservation

`func (o *InvoiceCreatePaylinkRequest) SetReservation(v bool)`

SetReservation sets Reservation field to given value.

### HasReservation

`func (o *InvoiceCreatePaylinkRequest) HasReservation() bool`

HasReservation returns a boolean if a field has been set.

### GetName

`func (o *InvoiceCreatePaylinkRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InvoiceCreatePaylinkRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InvoiceCreatePaylinkRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InvoiceCreatePaylinkRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFields

`func (o *InvoiceCreatePaylinkRequest) GetFields() []string`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *InvoiceCreatePaylinkRequest) GetFieldsOk() (*[]string, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *InvoiceCreatePaylinkRequest) SetFields(v []string)`

SetFields sets Fields field to given value.

### HasFields

`func (o *InvoiceCreatePaylinkRequest) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetHideFields

`func (o *InvoiceCreatePaylinkRequest) GetHideFields() bool`

GetHideFields returns the HideFields field if non-nil, zero value otherwise.

### GetHideFieldsOk

`func (o *InvoiceCreatePaylinkRequest) GetHideFieldsOk() (*bool, bool)`

GetHideFieldsOk returns a tuple with the HideFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideFields

`func (o *InvoiceCreatePaylinkRequest) SetHideFields(v bool)`

SetHideFields sets HideFields field to given value.

### HasHideFields

`func (o *InvoiceCreatePaylinkRequest) HasHideFields() bool`

HasHideFields returns a boolean if a field has been set.

### GetConcardisOrderId

`func (o *InvoiceCreatePaylinkRequest) GetConcardisOrderId() string`

GetConcardisOrderId returns the ConcardisOrderId field if non-nil, zero value otherwise.

### GetConcardisOrderIdOk

`func (o *InvoiceCreatePaylinkRequest) GetConcardisOrderIdOk() (*string, bool)`

GetConcardisOrderIdOk returns a tuple with the ConcardisOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcardisOrderId

`func (o *InvoiceCreatePaylinkRequest) SetConcardisOrderId(v string)`

SetConcardisOrderId sets ConcardisOrderId field to given value.

### HasConcardisOrderId

`func (o *InvoiceCreatePaylinkRequest) HasConcardisOrderId() bool`

HasConcardisOrderId returns a boolean if a field has been set.

### GetButtonText

`func (o *InvoiceCreatePaylinkRequest) GetButtonText() string`

GetButtonText returns the ButtonText field if non-nil, zero value otherwise.

### GetButtonTextOk

`func (o *InvoiceCreatePaylinkRequest) GetButtonTextOk() (*string, bool)`

GetButtonTextOk returns a tuple with the ButtonText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtonText

`func (o *InvoiceCreatePaylinkRequest) SetButtonText(v string)`

SetButtonText sets ButtonText field to given value.

### HasButtonText

`func (o *InvoiceCreatePaylinkRequest) HasButtonText() bool`

HasButtonText returns a boolean if a field has been set.

### GetExpirationDate

`func (o *InvoiceCreatePaylinkRequest) GetExpirationDate() string`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *InvoiceCreatePaylinkRequest) GetExpirationDateOk() (*string, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *InvoiceCreatePaylinkRequest) SetExpirationDate(v string)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *InvoiceCreatePaylinkRequest) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetSuccessRedirectUrl

`func (o *InvoiceCreatePaylinkRequest) GetSuccessRedirectUrl() string`

GetSuccessRedirectUrl returns the SuccessRedirectUrl field if non-nil, zero value otherwise.

### GetSuccessRedirectUrlOk

`func (o *InvoiceCreatePaylinkRequest) GetSuccessRedirectUrlOk() (*string, bool)`

GetSuccessRedirectUrlOk returns a tuple with the SuccessRedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessRedirectUrl

`func (o *InvoiceCreatePaylinkRequest) SetSuccessRedirectUrl(v string)`

SetSuccessRedirectUrl sets SuccessRedirectUrl field to given value.

### HasSuccessRedirectUrl

`func (o *InvoiceCreatePaylinkRequest) HasSuccessRedirectUrl() bool`

HasSuccessRedirectUrl returns a boolean if a field has been set.

### GetFailedRedirectUrl

`func (o *InvoiceCreatePaylinkRequest) GetFailedRedirectUrl() string`

GetFailedRedirectUrl returns the FailedRedirectUrl field if non-nil, zero value otherwise.

### GetFailedRedirectUrlOk

`func (o *InvoiceCreatePaylinkRequest) GetFailedRedirectUrlOk() (*string, bool)`

GetFailedRedirectUrlOk returns a tuple with the FailedRedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedRedirectUrl

`func (o *InvoiceCreatePaylinkRequest) SetFailedRedirectUrl(v string)`

SetFailedRedirectUrl sets FailedRedirectUrl field to given value.

### HasFailedRedirectUrl

`func (o *InvoiceCreatePaylinkRequest) HasFailedRedirectUrl() bool`

HasFailedRedirectUrl returns a boolean if a field has been set.

### GetSubscriptionState

`func (o *InvoiceCreatePaylinkRequest) GetSubscriptionState() bool`

GetSubscriptionState returns the SubscriptionState field if non-nil, zero value otherwise.

### GetSubscriptionStateOk

`func (o *InvoiceCreatePaylinkRequest) GetSubscriptionStateOk() (*bool, bool)`

GetSubscriptionStateOk returns a tuple with the SubscriptionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionState

`func (o *InvoiceCreatePaylinkRequest) SetSubscriptionState(v bool)`

SetSubscriptionState sets SubscriptionState field to given value.

### HasSubscriptionState

`func (o *InvoiceCreatePaylinkRequest) HasSubscriptionState() bool`

HasSubscriptionState returns a boolean if a field has been set.

### GetSubscriptionInterval

`func (o *InvoiceCreatePaylinkRequest) GetSubscriptionInterval() string`

GetSubscriptionInterval returns the SubscriptionInterval field if non-nil, zero value otherwise.

### GetSubscriptionIntervalOk

`func (o *InvoiceCreatePaylinkRequest) GetSubscriptionIntervalOk() (*string, bool)`

GetSubscriptionIntervalOk returns a tuple with the SubscriptionInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionInterval

`func (o *InvoiceCreatePaylinkRequest) SetSubscriptionInterval(v string)`

SetSubscriptionInterval sets SubscriptionInterval field to given value.

### HasSubscriptionInterval

`func (o *InvoiceCreatePaylinkRequest) HasSubscriptionInterval() bool`

HasSubscriptionInterval returns a boolean if a field has been set.

### GetSubscriptionPeriod

`func (o *InvoiceCreatePaylinkRequest) GetSubscriptionPeriod() string`

GetSubscriptionPeriod returns the SubscriptionPeriod field if non-nil, zero value otherwise.

### GetSubscriptionPeriodOk

`func (o *InvoiceCreatePaylinkRequest) GetSubscriptionPeriodOk() (*string, bool)`

GetSubscriptionPeriodOk returns a tuple with the SubscriptionPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionPeriod

`func (o *InvoiceCreatePaylinkRequest) SetSubscriptionPeriod(v string)`

SetSubscriptionPeriod sets SubscriptionPeriod field to given value.

### HasSubscriptionPeriod

`func (o *InvoiceCreatePaylinkRequest) HasSubscriptionPeriod() bool`

HasSubscriptionPeriod returns a boolean if a field has been set.

### GetSubscriptionCancellationInterval

`func (o *InvoiceCreatePaylinkRequest) GetSubscriptionCancellationInterval() string`

GetSubscriptionCancellationInterval returns the SubscriptionCancellationInterval field if non-nil, zero value otherwise.

### GetSubscriptionCancellationIntervalOk

`func (o *InvoiceCreatePaylinkRequest) GetSubscriptionCancellationIntervalOk() (*string, bool)`

GetSubscriptionCancellationIntervalOk returns a tuple with the SubscriptionCancellationInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionCancellationInterval

`func (o *InvoiceCreatePaylinkRequest) SetSubscriptionCancellationInterval(v string)`

SetSubscriptionCancellationInterval sets SubscriptionCancellationInterval field to given value.

### HasSubscriptionCancellationInterval

`func (o *InvoiceCreatePaylinkRequest) HasSubscriptionCancellationInterval() bool`

HasSubscriptionCancellationInterval returns a boolean if a field has been set.

### GetAttachments

`func (o *InvoiceCreatePaylinkRequest) GetAttachments() *os.File`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *InvoiceCreatePaylinkRequest) GetAttachmentsOk() (**os.File, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *InvoiceCreatePaylinkRequest) SetAttachments(v *os.File)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *InvoiceCreatePaylinkRequest) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetIsPriceExclusiveVat

`func (o *InvoiceCreatePaylinkRequest) GetIsPriceExclusiveVat() bool`

GetIsPriceExclusiveVat returns the IsPriceExclusiveVat field if non-nil, zero value otherwise.

### GetIsPriceExclusiveVatOk

`func (o *InvoiceCreatePaylinkRequest) GetIsPriceExclusiveVatOk() (*bool, bool)`

GetIsPriceExclusiveVatOk returns a tuple with the IsPriceExclusiveVat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPriceExclusiveVat

`func (o *InvoiceCreatePaylinkRequest) SetIsPriceExclusiveVat(v bool)`

SetIsPriceExclusiveVat sets IsPriceExclusiveVat field to given value.

### HasIsPriceExclusiveVat

`func (o *InvoiceCreatePaylinkRequest) HasIsPriceExclusiveVat() bool`

HasIsPriceExclusiveVat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


