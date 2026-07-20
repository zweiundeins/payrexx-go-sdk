# Paylink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [default to 0]
**Hash** | Pointer to **string** |  | [optional] 
**ReferenceId** | Pointer to **string** |  | [optional] 
**Link** | Pointer to **string** |  | [optional] 
**Invoices** | Pointer to **[]string** |  | [optional] 
**PreAuthorization** | Pointer to **int32** |  | [optional] [default to 0]
**Reservation** | Pointer to **int32** |  | [optional] [default to 0]
**Name** | Pointer to **string** |  | [optional] 
**Api** | Pointer to **bool** |  | [optional] [default to true]
**Fields** | Pointer to [**PaylinkFields**](PaylinkFields.md) |  | [optional] 
**Psp** | Pointer to **int32** |  | [optional] [default to 0]
**Pm** | Pointer to **[]string** |  | [optional] 
**Purpose** | Pointer to **string** |  | [optional] 
**Amount** | Pointer to **int32** |  | [optional] [default to 0]
**VatRate** | Pointer to **float32** |  | [optional] [default to 0]
**Currency** | Pointer to **string** |  | [optional] 
**Sku** | Pointer to **string** |  | [optional] 
**SubscriptionState** | Pointer to **bool** |  | [optional] [default to true]
**SubscriptionInterval** | Pointer to **string** |  | [optional] 
**SubscriptionPeriod** | Pointer to **string** |  | [optional] 
**SubscriptionPeriodMinAmount** | Pointer to **string** |  | [optional] 
**SubscriptionCancellationInterval** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **int32** |  | [optional] [default to 0]

## Methods

### NewPaylink

`func NewPaylink() *Paylink`

NewPaylink instantiates a new Paylink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaylinkWithDefaults

`func NewPaylinkWithDefaults() *Paylink`

NewPaylinkWithDefaults instantiates a new Paylink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Paylink) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Paylink) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Paylink) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Paylink) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHash

`func (o *Paylink) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *Paylink) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *Paylink) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *Paylink) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetReferenceId

`func (o *Paylink) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *Paylink) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *Paylink) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *Paylink) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.

### GetLink

`func (o *Paylink) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *Paylink) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *Paylink) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *Paylink) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetInvoices

`func (o *Paylink) GetInvoices() []string`

GetInvoices returns the Invoices field if non-nil, zero value otherwise.

### GetInvoicesOk

`func (o *Paylink) GetInvoicesOk() (*[]string, bool)`

GetInvoicesOk returns a tuple with the Invoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoices

`func (o *Paylink) SetInvoices(v []string)`

SetInvoices sets Invoices field to given value.

### HasInvoices

`func (o *Paylink) HasInvoices() bool`

HasInvoices returns a boolean if a field has been set.

### GetPreAuthorization

`func (o *Paylink) GetPreAuthorization() int32`

GetPreAuthorization returns the PreAuthorization field if non-nil, zero value otherwise.

### GetPreAuthorizationOk

`func (o *Paylink) GetPreAuthorizationOk() (*int32, bool)`

GetPreAuthorizationOk returns a tuple with the PreAuthorization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreAuthorization

`func (o *Paylink) SetPreAuthorization(v int32)`

SetPreAuthorization sets PreAuthorization field to given value.

### HasPreAuthorization

`func (o *Paylink) HasPreAuthorization() bool`

HasPreAuthorization returns a boolean if a field has been set.

### GetReservation

`func (o *Paylink) GetReservation() int32`

GetReservation returns the Reservation field if non-nil, zero value otherwise.

### GetReservationOk

`func (o *Paylink) GetReservationOk() (*int32, bool)`

GetReservationOk returns a tuple with the Reservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservation

`func (o *Paylink) SetReservation(v int32)`

SetReservation sets Reservation field to given value.

### HasReservation

`func (o *Paylink) HasReservation() bool`

HasReservation returns a boolean if a field has been set.

### GetName

`func (o *Paylink) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Paylink) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Paylink) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Paylink) HasName() bool`

HasName returns a boolean if a field has been set.

### GetApi

`func (o *Paylink) GetApi() bool`

GetApi returns the Api field if non-nil, zero value otherwise.

### GetApiOk

`func (o *Paylink) GetApiOk() (*bool, bool)`

GetApiOk returns a tuple with the Api field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApi

`func (o *Paylink) SetApi(v bool)`

SetApi sets Api field to given value.

### HasApi

`func (o *Paylink) HasApi() bool`

HasApi returns a boolean if a field has been set.

### GetFields

`func (o *Paylink) GetFields() PaylinkFields`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *Paylink) GetFieldsOk() (*PaylinkFields, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *Paylink) SetFields(v PaylinkFields)`

SetFields sets Fields field to given value.

### HasFields

`func (o *Paylink) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetPsp

`func (o *Paylink) GetPsp() int32`

GetPsp returns the Psp field if non-nil, zero value otherwise.

### GetPspOk

`func (o *Paylink) GetPspOk() (*int32, bool)`

GetPspOk returns a tuple with the Psp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPsp

`func (o *Paylink) SetPsp(v int32)`

SetPsp sets Psp field to given value.

### HasPsp

`func (o *Paylink) HasPsp() bool`

HasPsp returns a boolean if a field has been set.

### GetPm

`func (o *Paylink) GetPm() []string`

GetPm returns the Pm field if non-nil, zero value otherwise.

### GetPmOk

`func (o *Paylink) GetPmOk() (*[]string, bool)`

GetPmOk returns a tuple with the Pm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPm

`func (o *Paylink) SetPm(v []string)`

SetPm sets Pm field to given value.

### HasPm

`func (o *Paylink) HasPm() bool`

HasPm returns a boolean if a field has been set.

### GetPurpose

`func (o *Paylink) GetPurpose() string`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *Paylink) GetPurposeOk() (*string, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *Paylink) SetPurpose(v string)`

SetPurpose sets Purpose field to given value.

### HasPurpose

`func (o *Paylink) HasPurpose() bool`

HasPurpose returns a boolean if a field has been set.

### GetAmount

`func (o *Paylink) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Paylink) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Paylink) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *Paylink) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetVatRate

`func (o *Paylink) GetVatRate() float32`

GetVatRate returns the VatRate field if non-nil, zero value otherwise.

### GetVatRateOk

`func (o *Paylink) GetVatRateOk() (*float32, bool)`

GetVatRateOk returns a tuple with the VatRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatRate

`func (o *Paylink) SetVatRate(v float32)`

SetVatRate sets VatRate field to given value.

### HasVatRate

`func (o *Paylink) HasVatRate() bool`

HasVatRate returns a boolean if a field has been set.

### GetCurrency

`func (o *Paylink) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Paylink) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Paylink) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *Paylink) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetSku

`func (o *Paylink) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *Paylink) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *Paylink) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *Paylink) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetSubscriptionState

`func (o *Paylink) GetSubscriptionState() bool`

GetSubscriptionState returns the SubscriptionState field if non-nil, zero value otherwise.

### GetSubscriptionStateOk

`func (o *Paylink) GetSubscriptionStateOk() (*bool, bool)`

GetSubscriptionStateOk returns a tuple with the SubscriptionState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionState

`func (o *Paylink) SetSubscriptionState(v bool)`

SetSubscriptionState sets SubscriptionState field to given value.

### HasSubscriptionState

`func (o *Paylink) HasSubscriptionState() bool`

HasSubscriptionState returns a boolean if a field has been set.

### GetSubscriptionInterval

`func (o *Paylink) GetSubscriptionInterval() string`

GetSubscriptionInterval returns the SubscriptionInterval field if non-nil, zero value otherwise.

### GetSubscriptionIntervalOk

`func (o *Paylink) GetSubscriptionIntervalOk() (*string, bool)`

GetSubscriptionIntervalOk returns a tuple with the SubscriptionInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionInterval

`func (o *Paylink) SetSubscriptionInterval(v string)`

SetSubscriptionInterval sets SubscriptionInterval field to given value.

### HasSubscriptionInterval

`func (o *Paylink) HasSubscriptionInterval() bool`

HasSubscriptionInterval returns a boolean if a field has been set.

### GetSubscriptionPeriod

`func (o *Paylink) GetSubscriptionPeriod() string`

GetSubscriptionPeriod returns the SubscriptionPeriod field if non-nil, zero value otherwise.

### GetSubscriptionPeriodOk

`func (o *Paylink) GetSubscriptionPeriodOk() (*string, bool)`

GetSubscriptionPeriodOk returns a tuple with the SubscriptionPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionPeriod

`func (o *Paylink) SetSubscriptionPeriod(v string)`

SetSubscriptionPeriod sets SubscriptionPeriod field to given value.

### HasSubscriptionPeriod

`func (o *Paylink) HasSubscriptionPeriod() bool`

HasSubscriptionPeriod returns a boolean if a field has been set.

### GetSubscriptionPeriodMinAmount

`func (o *Paylink) GetSubscriptionPeriodMinAmount() string`

GetSubscriptionPeriodMinAmount returns the SubscriptionPeriodMinAmount field if non-nil, zero value otherwise.

### GetSubscriptionPeriodMinAmountOk

`func (o *Paylink) GetSubscriptionPeriodMinAmountOk() (*string, bool)`

GetSubscriptionPeriodMinAmountOk returns a tuple with the SubscriptionPeriodMinAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionPeriodMinAmount

`func (o *Paylink) SetSubscriptionPeriodMinAmount(v string)`

SetSubscriptionPeriodMinAmount sets SubscriptionPeriodMinAmount field to given value.

### HasSubscriptionPeriodMinAmount

`func (o *Paylink) HasSubscriptionPeriodMinAmount() bool`

HasSubscriptionPeriodMinAmount returns a boolean if a field has been set.

### GetSubscriptionCancellationInterval

`func (o *Paylink) GetSubscriptionCancellationInterval() string`

GetSubscriptionCancellationInterval returns the SubscriptionCancellationInterval field if non-nil, zero value otherwise.

### GetSubscriptionCancellationIntervalOk

`func (o *Paylink) GetSubscriptionCancellationIntervalOk() (*string, bool)`

GetSubscriptionCancellationIntervalOk returns a tuple with the SubscriptionCancellationInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriptionCancellationInterval

`func (o *Paylink) SetSubscriptionCancellationInterval(v string)`

SetSubscriptionCancellationInterval sets SubscriptionCancellationInterval field to given value.

### HasSubscriptionCancellationInterval

`func (o *Paylink) HasSubscriptionCancellationInterval() bool`

HasSubscriptionCancellationInterval returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Paylink) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Paylink) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Paylink) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Paylink) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


