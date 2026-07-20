# GatewayCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int32** | Amount of payment in cents. | 
**VatRate** | Pointer to **float32** | VAT Rate Percentage | [optional] 
**Currency** | **string** | Currency of payment (ISO code). | 
**Sku** | Pointer to **string** | Product stock keeping unit | [optional] 
**Purpose** | Pointer to **string** | The purpose of the payment. | [optional] 
**SuccessRedirectUrl** | Pointer to **string** | URL to redirect to after successful payment. | [optional] 
**FailedRedirectUrl** | Pointer to **string** | URL to redirect to after failed payment. | [optional] 
**CancelRedirectUrl** | Pointer to **string** | URL to redirect to after manual cancellation by shopper. | [optional] 
**Basket** | Pointer to [**[]GatewayCreateRequestBasketInner**](GatewayCreateRequestBasketInner.md) | List of all products (incl. shipping costs). The sum of all product amounts must match the amount parameter above. Basket products contains name, description, quantity, amount (in cents) and vatRate (in Percent). | [optional] 
**Psp** | Pointer to **[]int32** | List of PSPs to provide for payment. If empty all available PSPs are provied. | [optional] 
**Pm** | Pointer to **[]string** | List of payment mean names to display | [optional] 
**PreAuthorization** | Pointer to **bool** | Whether charge payment manually at a later date (type authorization) | [optional] [default to false]
**Reservation** | Pointer to **bool** | Whether charge payment manually at a later date (type reservation) | [optional] [default to false]
**ReferenceId** | Pointer to **string** | An internal reference id used by your system. | [optional] 
**Fields** | Pointer to [**GatewayCreateRequestFields**](GatewayCreateRequestFields.md) |  | [optional] 
**Language** | Pointer to **string** | The language ISO code by ISO 639-1. | [optional] 
**ConcardisOrderId** | Pointer to **string** | Only available for Concardis PSP and if the custom ORDERID option is activated in PSP settings in Payrexx administration. This ORDERID will be transferred to the Payengine. | [optional] 
**SkipResultPage** | Pointer to **bool** | Skip result page and directly redirect to success or failed URL | [optional] [default to false]
**ChargeOnAuthorization** | Pointer to **bool** | preAuthorization needs to be \&quot;true\&quot;. This will charge the authorization during the first payment. | [optional] [default to false]
**Validity** | Pointer to **int32** | Gateway validity in minutes. | [optional] 
**SubscriptionState** | Pointer to **bool** | Defines whether the payment should be handled as subscription. | [optional] [default to false]
**SubscriptionInterval** | Pointer to **string** | Payment interval | [optional] 
**SubscriptionPeriod** | Pointer to **string** | Duration of the subscription | [optional] 
**SubscriptionCancellationInterval** | Pointer to **string** | Defines the period, in which a subscription can be canceled. | [optional] 
**ButtonText** | Pointer to **[]string** | Change the default button Text \&quot;Pay\&quot; to a custom String | [optional] 
**LookAndFeelProfile** | Pointer to **string** | UUID of the look and feel profile to use. | [optional] 
**SuccessMessage** | Pointer to **string** | Custom success message on result page. | [optional] 
**QrCodeSessionId** | Pointer to **string** | Holds the session ID of a scanned QR code. Only available and needed for Static QR Code with Twint. | [optional] 
**ApplicationFee** | Pointer to **int32** | Amount in the smallest unit of the currency. | [optional] 
**ReserveOnAuthorization** | Pointer to **bool** | preAuthorization needs to be \&quot;true\&quot;. This will create a reservation based on an authorization during the first payment. | [optional] 
**IsPriceExclusiveVat** | Pointer to **bool** | if set to  true the vat will be added on top of amount | [optional] 

## Methods

### NewGatewayCreateRequest

`func NewGatewayCreateRequest(amount int32, currency string, ) *GatewayCreateRequest`

NewGatewayCreateRequest instantiates a new GatewayCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayCreateRequestWithDefaults

`func NewGatewayCreateRequestWithDefaults() *GatewayCreateRequest`

NewGatewayCreateRequestWithDefaults instantiates a new GatewayCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *GatewayCreateRequest) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *GatewayCreateRequest) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *GatewayCreateRequest) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetVatRate

`func (o *GatewayCreateRequest) GetVatRate() float32`

GetVatRate returns the VatRate field if non-nil, zero value otherwise.

### GetVatRateOk

`func (o *GatewayCreateRequest) GetVatRateOk() (*float32, bool)`

GetVatRateOk returns a tuple with the VatRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatRate

`func (o *GatewayCreateRequest) SetVatRate(v float32)`

SetVatRate sets VatRate field to given value.

### HasVatRate

`func (o *GatewayCreateRequest) HasVatRate() bool`

HasVatRate returns a boolean if a field has been set.

### GetCurrency

`func (o *GatewayCreateRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *GatewayCreateRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *GatewayCreateRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetSku

`func (o *GatewayCreateRequest) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *GatewayCreateRequest) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *GatewayCreateRequest) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *GatewayCreateRequest) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetPurpose

`func (o *GatewayCreateRequest) GetPurpose() string`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *GatewayCreateRequest) GetPurposeOk() (*string, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *GatewayCreateRequest) SetPurpose(v string)`

SetPurpose sets Purpose field to given value.

### HasPurpose

`func (o *GatewayCreateRequest) HasPurpose() bool`

HasPurpose returns a boolean if a field has been set.

### GetSuccessRedirectUrl

`func (o *GatewayCreateRequest) GetSuccessRedirectUrl() string`

GetSuccessRedirectUrl returns the SuccessRedirectUrl field if non-nil, zero value otherwise.

### GetSuccessRedirectUrlOk

`func (o *GatewayCreateRequest) GetSuccessRedirectUrlOk() (*string, bool)`

GetSuccessRedirectUrlOk returns a tuple with the SuccessRedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessRedirectUrl

`func (o *GatewayCreateRequest) SetSuccessRedirectUrl(v string)`

SetSuccessRedirectUrl sets SuccessRedirectUrl field to given value.

### HasSuccessRedirectUrl

`func (o *GatewayCreateRequest) HasSuccessRedirectUrl() bool`

HasSuccessRedirectUrl returns a boolean if a field has been set.

### GetFailedRedirectUrl

`func (o *GatewayCreateRequest) GetFailedRedirectUrl() string`

GetFailedRedirectUrl returns the FailedRedirectUrl field if non-nil, zero value otherwise.

### GetFailedRedirectUrlOk

`func (o *GatewayCreateRequest) GetFailedRedirectUrlOk() (*string, bool)`

GetFailedRedirectUrlOk returns a tuple with the FailedRedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedRedirectUrl

`func (o *GatewayCreateRequest) SetFailedRedirectUrl(v string)`

SetFailedRedirectUrl sets FailedRedirectUrl field to given value.

### HasFailedRedirectUrl

`func (o *GatewayCreateRequest) HasFailedRedirectUrl() bool`

HasFailedRedirectUrl returns a boolean if a field has been set.

### GetCancelRedirectUrl

`func (o *GatewayCreateRequest) GetCancelRedirectUrl() string`

GetCancelRedirectUrl returns the CancelRedirectUrl field if non-nil, zero value otherwise.

### GetCancelRedirectUrlOk

`func (o *GatewayCreateRequest) GetCancelRedirectUrlOk() (*string, bool)`

GetCancelRedirectUrlOk returns a tuple with the CancelRedirectUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelRedirectUrl

`func (o *GatewayCreateRequest) SetCancelRedirectUrl(v string)`

SetCancelRedirectUrl sets CancelRedirectUrl field to given value.

### HasCancelRedirectUrl

`func (o *GatewayCreateRequest) HasCancelRedirectUrl() bool`

HasCancelRedirectUrl returns a boolean if a field has been set.

### GetBasket

`func (o *GatewayCreateRequest) GetBasket() []GatewayCreateRequestBasketInner`

GetBasket returns the Basket field if non-nil, zero value otherwise.

### GetBasketOk

`func (o *GatewayCreateRequest) GetBasketOk() (*[]GatewayCreateRequestBasketInner, bool)`

GetBasketOk returns a tuple with the Basket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasket

`func (o *GatewayCreateRequest) SetBasket(v []GatewayCreateRequestBasketInner)`

SetBasket sets Basket field to given value.

### HasBasket

`func (o *GatewayCreateRequest) HasBasket() bool`

HasBasket returns a boolean if a field has been set.

### GetPsp

`func (o *GatewayCreateRequest) GetPsp() []int32`

GetPsp returns the Psp field if non-nil, zero value otherwise.

### GetPspOk

`func (o *GatewayCreateRequest) GetPspOk() (*[]int32, bool)`

GetPspOk returns a tuple with the Psp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPsp

`func (o *GatewayCreateRequest) SetPsp(v []int32)`

SetPsp sets Psp field to given value.

### HasPsp

`func (o *GatewayCreateRequest) HasPsp() bool`

HasPsp returns a boolean if a field has been set.

### GetPm

`func (o *GatewayCreateRequest) GetPm() []string`

GetPm returns the Pm field if non-nil, zero value otherwise.

### GetPmOk

`func (o *GatewayCreateRequest) GetPmOk() (*[]string, bool)`

GetPmOk returns a tuple with the Pm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPm

`func (o *GatewayCreateRequest) SetPm(v []string)`

SetPm sets Pm field to given value.

### HasPm

`func (o *GatewayCreateRequest) HasPm() bool`

HasPm returns a boolean if a field has been set.

### GetPreAuthorization

`func (o *GatewayCreateRequest) GetPreAuthorization() bool`

GetPreAuthorization returns the PreAuthorization field if non-nil, zero value otherwise.

### GetPreAuthorizationOk

`func (o *GatewayCreateRequest) GetPreAuthorizationOk() (*bool, bool)`

GetPreAuthorizationOk returns a tuple with the PreAuthorization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreAuthorization

`func (o *GatewayCreateRequest) SetPreAuthorization(v bool)`

SetPreAuthorization sets PreAuthorization field to given value.

### HasPreAuthorization

`func (o *GatewayCreateRequest) HasPreAuthorization() bool`

HasPreAuthorization returns a boolean if a field has been set.

### GetReservation

`func (o *GatewayCreateRequest) GetReservation() bool`

GetReservation returns the Reservation field if non-nil, zero value otherwise.

### GetReservationOk

`func (o *GatewayCreateRequest) GetReservationOk() (*bool, bool)`

GetReservationOk returns a tuple with the Reservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservation

`func (o *GatewayCreateRequest) SetReservation(v bool)`

SetReservation sets Reservation field to given value.

### HasReservation

`func (o *GatewayCreateRequest) HasReservation() bool`

HasReservation returns a boolean if a field has been set.

### GetReferenceId

`func (o *GatewayCreateRequest) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *GatewayCreateRequest) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *GatewayCreateRequest) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *GatewayCreateRequest) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.

### GetFields

`func (o *GatewayCreateRequest) GetFields() GatewayCreateRequestFields`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *GatewayCreateRequest) GetFieldsOk() (*GatewayCreateRequestFields, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *GatewayCreateRequest) SetFields(v GatewayCreateRequestFields)`

SetFields sets Fields field to given value.

### HasFields

`func (o *GatewayCreateRequest) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetLanguage

`func (o *GatewayCreateRequest) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *GatewayCreateRequest) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *GatewayCreateRequest) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *GatewayCreateRequest) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetConcardisOrderId

`func (o *GatewayCreateRequest) GetConcardisOrderId() string`

GetConcardisOrderId returns the ConcardisOrderId field if non-nil, zero value otherwise.

### GetConcardisOrderIdOk

`func (o *GatewayCreateRequest) GetConcardisOrderIdOk() (*string, bool)`

GetConcardisOrderIdOk returns a tuple with the ConcardisOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConcardisOrderId

`func (o *GatewayCreateRequest) SetConcardisOrderId(v string)`

SetConcardisOrderId sets ConcardisOrderId field to given value.

### HasConcardisOrderId

`func (o *GatewayCreateRequest) HasConcardisOrderId() bool`

HasConcardisOrderId returns a boolean if a field has been set.

### GetSkipResultPage

`func (o *GatewayCreateRequest) GetSkipResultPage() bool`

GetSkipResultPage returns the SkipResultPage field if non-nil, zero value otherwise.

### GetSkipResultPageOk

`func (o *GatewayCreateRequest) GetSkipResultPageOk() (*bool, bool)`

GetSkipResultPageOk returns a tuple with the SkipResultPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipResultPage

`func (o *GatewayCreateRequest) SetSkipResultPage(v bool)`

SetSkipResultPage sets SkipResultPage field to given value.

### HasSkipResultPage

`func (o *GatewayCreateRequest) HasSkipResultPage() bool`

HasSkipResultPage returns a boolean if a field has been set.

### GetChargeOnAuthorization

`func (o *GatewayCreateRequest) GetChargeOnAuthorization() bool`

GetChargeOnAuthorization returns the ChargeOnAuthorization field if non-nil, zero value otherwise.

### GetChargeOnAuthorizationOk

`func (o *GatewayCreateRequest) GetChargeOnAuthorizationOk() (*bool, bool)`

GetChargeOnAuthorizationOk returns a tuple with the ChargeOnAuthorization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChargeOnAuthorization

`func (o *GatewayCreateRequest) SetChargeOnAuthorization(v bool)`

SetChargeOnAuthorization sets ChargeOnAuthorization field to given value.

### HasChargeOnAuthorization

`func (o *GatewayCreateRequest) HasChargeOnAuthorization() bool`

HasChargeOnAuthorization returns a boolean if a field has been set.

### GetValidity

`func (o *GatewayCreateRequest) GetValidity() int32`

GetValidity returns the Validity field if non-nil, zero value otherwise.

### GetValidityOk

`func (o *GatewayCreateRequest) GetValidityOk() (*int32, bool)`

GetValidityOk returns a tuple with the Validity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidity

`func (o *GatewayCreateRequest) SetValidity(v int32)`

SetValidity sets Validity field to given value.

### HasValidity

`func (o *GatewayCreateRequest) HasValidity() bool`

HasValidity returns a boolean if a field has been set.

### GetSubscriptionState

`func (o *GatewayCreateRequest) GetSubscriptionState() bool`

GetSubscriptionState returns the SubscriptionState field if non-nil, zero value otherwise.

### GetSubscriptionStateOk

`func (o *GatewayCreateRequest) GetSubscriptionStateOk() (*bool, bool)`

GetSubscriptionStateOk returns a tuple with the SubscriptionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionState

`func (o *GatewayCreateRequest) SetSubscriptionState(v bool)`

SetSubscriptionState sets SubscriptionState field to given value.

### HasSubscriptionState

`func (o *GatewayCreateRequest) HasSubscriptionState() bool`

HasSubscriptionState returns a boolean if a field has been set.

### GetSubscriptionInterval

`func (o *GatewayCreateRequest) GetSubscriptionInterval() string`

GetSubscriptionInterval returns the SubscriptionInterval field if non-nil, zero value otherwise.

### GetSubscriptionIntervalOk

`func (o *GatewayCreateRequest) GetSubscriptionIntervalOk() (*string, bool)`

GetSubscriptionIntervalOk returns a tuple with the SubscriptionInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionInterval

`func (o *GatewayCreateRequest) SetSubscriptionInterval(v string)`

SetSubscriptionInterval sets SubscriptionInterval field to given value.

### HasSubscriptionInterval

`func (o *GatewayCreateRequest) HasSubscriptionInterval() bool`

HasSubscriptionInterval returns a boolean if a field has been set.

### GetSubscriptionPeriod

`func (o *GatewayCreateRequest) GetSubscriptionPeriod() string`

GetSubscriptionPeriod returns the SubscriptionPeriod field if non-nil, zero value otherwise.

### GetSubscriptionPeriodOk

`func (o *GatewayCreateRequest) GetSubscriptionPeriodOk() (*string, bool)`

GetSubscriptionPeriodOk returns a tuple with the SubscriptionPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionPeriod

`func (o *GatewayCreateRequest) SetSubscriptionPeriod(v string)`

SetSubscriptionPeriod sets SubscriptionPeriod field to given value.

### HasSubscriptionPeriod

`func (o *GatewayCreateRequest) HasSubscriptionPeriod() bool`

HasSubscriptionPeriod returns a boolean if a field has been set.

### GetSubscriptionCancellationInterval

`func (o *GatewayCreateRequest) GetSubscriptionCancellationInterval() string`

GetSubscriptionCancellationInterval returns the SubscriptionCancellationInterval field if non-nil, zero value otherwise.

### GetSubscriptionCancellationIntervalOk

`func (o *GatewayCreateRequest) GetSubscriptionCancellationIntervalOk() (*string, bool)`

GetSubscriptionCancellationIntervalOk returns a tuple with the SubscriptionCancellationInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionCancellationInterval

`func (o *GatewayCreateRequest) SetSubscriptionCancellationInterval(v string)`

SetSubscriptionCancellationInterval sets SubscriptionCancellationInterval field to given value.

### HasSubscriptionCancellationInterval

`func (o *GatewayCreateRequest) HasSubscriptionCancellationInterval() bool`

HasSubscriptionCancellationInterval returns a boolean if a field has been set.

### GetButtonText

`func (o *GatewayCreateRequest) GetButtonText() []string`

GetButtonText returns the ButtonText field if non-nil, zero value otherwise.

### GetButtonTextOk

`func (o *GatewayCreateRequest) GetButtonTextOk() (*[]string, bool)`

GetButtonTextOk returns a tuple with the ButtonText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetButtonText

`func (o *GatewayCreateRequest) SetButtonText(v []string)`

SetButtonText sets ButtonText field to given value.

### HasButtonText

`func (o *GatewayCreateRequest) HasButtonText() bool`

HasButtonText returns a boolean if a field has been set.

### GetLookAndFeelProfile

`func (o *GatewayCreateRequest) GetLookAndFeelProfile() string`

GetLookAndFeelProfile returns the LookAndFeelProfile field if non-nil, zero value otherwise.

### GetLookAndFeelProfileOk

`func (o *GatewayCreateRequest) GetLookAndFeelProfileOk() (*string, bool)`

GetLookAndFeelProfileOk returns a tuple with the LookAndFeelProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLookAndFeelProfile

`func (o *GatewayCreateRequest) SetLookAndFeelProfile(v string)`

SetLookAndFeelProfile sets LookAndFeelProfile field to given value.

### HasLookAndFeelProfile

`func (o *GatewayCreateRequest) HasLookAndFeelProfile() bool`

HasLookAndFeelProfile returns a boolean if a field has been set.

### GetSuccessMessage

`func (o *GatewayCreateRequest) GetSuccessMessage() string`

GetSuccessMessage returns the SuccessMessage field if non-nil, zero value otherwise.

### GetSuccessMessageOk

`func (o *GatewayCreateRequest) GetSuccessMessageOk() (*string, bool)`

GetSuccessMessageOk returns a tuple with the SuccessMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccessMessage

`func (o *GatewayCreateRequest) SetSuccessMessage(v string)`

SetSuccessMessage sets SuccessMessage field to given value.

### HasSuccessMessage

`func (o *GatewayCreateRequest) HasSuccessMessage() bool`

HasSuccessMessage returns a boolean if a field has been set.

### GetQrCodeSessionId

`func (o *GatewayCreateRequest) GetQrCodeSessionId() string`

GetQrCodeSessionId returns the QrCodeSessionId field if non-nil, zero value otherwise.

### GetQrCodeSessionIdOk

`func (o *GatewayCreateRequest) GetQrCodeSessionIdOk() (*string, bool)`

GetQrCodeSessionIdOk returns a tuple with the QrCodeSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQrCodeSessionId

`func (o *GatewayCreateRequest) SetQrCodeSessionId(v string)`

SetQrCodeSessionId sets QrCodeSessionId field to given value.

### HasQrCodeSessionId

`func (o *GatewayCreateRequest) HasQrCodeSessionId() bool`

HasQrCodeSessionId returns a boolean if a field has been set.

### GetApplicationFee

`func (o *GatewayCreateRequest) GetApplicationFee() int32`

GetApplicationFee returns the ApplicationFee field if non-nil, zero value otherwise.

### GetApplicationFeeOk

`func (o *GatewayCreateRequest) GetApplicationFeeOk() (*int32, bool)`

GetApplicationFeeOk returns a tuple with the ApplicationFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationFee

`func (o *GatewayCreateRequest) SetApplicationFee(v int32)`

SetApplicationFee sets ApplicationFee field to given value.

### HasApplicationFee

`func (o *GatewayCreateRequest) HasApplicationFee() bool`

HasApplicationFee returns a boolean if a field has been set.

### GetReserveOnAuthorization

`func (o *GatewayCreateRequest) GetReserveOnAuthorization() bool`

GetReserveOnAuthorization returns the ReserveOnAuthorization field if non-nil, zero value otherwise.

### GetReserveOnAuthorizationOk

`func (o *GatewayCreateRequest) GetReserveOnAuthorizationOk() (*bool, bool)`

GetReserveOnAuthorizationOk returns a tuple with the ReserveOnAuthorization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReserveOnAuthorization

`func (o *GatewayCreateRequest) SetReserveOnAuthorization(v bool)`

SetReserveOnAuthorization sets ReserveOnAuthorization field to given value.

### HasReserveOnAuthorization

`func (o *GatewayCreateRequest) HasReserveOnAuthorization() bool`

HasReserveOnAuthorization returns a boolean if a field has been set.

### GetIsPriceExclusiveVat

`func (o *GatewayCreateRequest) GetIsPriceExclusiveVat() bool`

GetIsPriceExclusiveVat returns the IsPriceExclusiveVat field if non-nil, zero value otherwise.

### GetIsPriceExclusiveVatOk

`func (o *GatewayCreateRequest) GetIsPriceExclusiveVatOk() (*bool, bool)`

GetIsPriceExclusiveVatOk returns a tuple with the IsPriceExclusiveVat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPriceExclusiveVat

`func (o *GatewayCreateRequest) SetIsPriceExclusiveVat(v bool)`

SetIsPriceExclusiveVat sets IsPriceExclusiveVat field to given value.

### HasIsPriceExclusiveVat

`func (o *GatewayCreateRequest) HasIsPriceExclusiveVat() bool`

HasIsPriceExclusiveVat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


