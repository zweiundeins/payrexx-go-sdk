# Transaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [default to 0]
**Uuid** | Pointer to **string** |  | [optional] 
**ReferenceId** | Pointer to **string** |  | [optional] 
**Time** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Lang** | Pointer to **string** |  | [optional] 
**Psp** | Pointer to **string** |  | [optional] 
**Amount** | Pointer to **int32** |  | [optional] [default to 0]
**Contact** | Pointer to [**TransactionContact**](TransactionContact.md) |  | [optional] 
**PageUuid** | Pointer to **string** |  | [optional] 
**Payment** | Pointer to [**TransactionPayment**](TransactionPayment.md) |  | [optional] 
**PayoutUuid** | Pointer to **string** |  | [optional] 
**PspId** | Pointer to **int32** |  | [optional] [default to 0]
**Mode** | Pointer to **string** |  | [optional] 
**Invoice** | Pointer to [**TransactionInvoice**](TransactionInvoice.md) |  | [optional] 
**Refundable** | Pointer to **bool** |  | [optional] [default to true]
**PartiallyRefundable** | Pointer to **bool** |  | [optional] [default to true]
**VatRate** | Pointer to **string** |  | [optional] 

## Methods

### NewTransaction

`func NewTransaction() *Transaction`

NewTransaction instantiates a new Transaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionWithDefaults

`func NewTransactionWithDefaults() *Transaction`

NewTransactionWithDefaults instantiates a new Transaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Transaction) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Transaction) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Transaction) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Transaction) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *Transaction) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Transaction) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Transaction) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *Transaction) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetReferenceId

`func (o *Transaction) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *Transaction) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *Transaction) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *Transaction) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.

### GetTime

`func (o *Transaction) GetTime() string`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *Transaction) GetTimeOk() (*string, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *Transaction) SetTime(v string)`

SetTime sets Time field to given value.

### HasTime

`func (o *Transaction) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetStatus

`func (o *Transaction) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Transaction) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Transaction) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Transaction) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetLang

`func (o *Transaction) GetLang() string`

GetLang returns the Lang field if non-nil, zero value otherwise.

### GetLangOk

`func (o *Transaction) GetLangOk() (*string, bool)`

GetLangOk returns a tuple with the Lang field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLang

`func (o *Transaction) SetLang(v string)`

SetLang sets Lang field to given value.

### HasLang

`func (o *Transaction) HasLang() bool`

HasLang returns a boolean if a field has been set.

### GetPsp

`func (o *Transaction) GetPsp() string`

GetPsp returns the Psp field if non-nil, zero value otherwise.

### GetPspOk

`func (o *Transaction) GetPspOk() (*string, bool)`

GetPspOk returns a tuple with the Psp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPsp

`func (o *Transaction) SetPsp(v string)`

SetPsp sets Psp field to given value.

### HasPsp

`func (o *Transaction) HasPsp() bool`

HasPsp returns a boolean if a field has been set.

### GetAmount

`func (o *Transaction) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Transaction) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Transaction) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *Transaction) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetContact

`func (o *Transaction) GetContact() TransactionContact`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *Transaction) GetContactOk() (*TransactionContact, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *Transaction) SetContact(v TransactionContact)`

SetContact sets Contact field to given value.

### HasContact

`func (o *Transaction) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetPageUuid

`func (o *Transaction) GetPageUuid() string`

GetPageUuid returns the PageUuid field if non-nil, zero value otherwise.

### GetPageUuidOk

`func (o *Transaction) GetPageUuidOk() (*string, bool)`

GetPageUuidOk returns a tuple with the PageUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageUuid

`func (o *Transaction) SetPageUuid(v string)`

SetPageUuid sets PageUuid field to given value.

### HasPageUuid

`func (o *Transaction) HasPageUuid() bool`

HasPageUuid returns a boolean if a field has been set.

### GetPayment

`func (o *Transaction) GetPayment() TransactionPayment`

GetPayment returns the Payment field if non-nil, zero value otherwise.

### GetPaymentOk

`func (o *Transaction) GetPaymentOk() (*TransactionPayment, bool)`

GetPaymentOk returns a tuple with the Payment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayment

`func (o *Transaction) SetPayment(v TransactionPayment)`

SetPayment sets Payment field to given value.

### HasPayment

`func (o *Transaction) HasPayment() bool`

HasPayment returns a boolean if a field has been set.

### GetPayoutUuid

`func (o *Transaction) GetPayoutUuid() string`

GetPayoutUuid returns the PayoutUuid field if non-nil, zero value otherwise.

### GetPayoutUuidOk

`func (o *Transaction) GetPayoutUuidOk() (*string, bool)`

GetPayoutUuidOk returns a tuple with the PayoutUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayoutUuid

`func (o *Transaction) SetPayoutUuid(v string)`

SetPayoutUuid sets PayoutUuid field to given value.

### HasPayoutUuid

`func (o *Transaction) HasPayoutUuid() bool`

HasPayoutUuid returns a boolean if a field has been set.

### GetPspId

`func (o *Transaction) GetPspId() int32`

GetPspId returns the PspId field if non-nil, zero value otherwise.

### GetPspIdOk

`func (o *Transaction) GetPspIdOk() (*int32, bool)`

GetPspIdOk returns a tuple with the PspId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPspId

`func (o *Transaction) SetPspId(v int32)`

SetPspId sets PspId field to given value.

### HasPspId

`func (o *Transaction) HasPspId() bool`

HasPspId returns a boolean if a field has been set.

### GetMode

`func (o *Transaction) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *Transaction) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *Transaction) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *Transaction) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetInvoice

`func (o *Transaction) GetInvoice() TransactionInvoice`

GetInvoice returns the Invoice field if non-nil, zero value otherwise.

### GetInvoiceOk

`func (o *Transaction) GetInvoiceOk() (*TransactionInvoice, bool)`

GetInvoiceOk returns a tuple with the Invoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoice

`func (o *Transaction) SetInvoice(v TransactionInvoice)`

SetInvoice sets Invoice field to given value.

### HasInvoice

`func (o *Transaction) HasInvoice() bool`

HasInvoice returns a boolean if a field has been set.

### GetRefundable

`func (o *Transaction) GetRefundable() bool`

GetRefundable returns the Refundable field if non-nil, zero value otherwise.

### GetRefundableOk

`func (o *Transaction) GetRefundableOk() (*bool, bool)`

GetRefundableOk returns a tuple with the Refundable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundable

`func (o *Transaction) SetRefundable(v bool)`

SetRefundable sets Refundable field to given value.

### HasRefundable

`func (o *Transaction) HasRefundable() bool`

HasRefundable returns a boolean if a field has been set.

### GetPartiallyRefundable

`func (o *Transaction) GetPartiallyRefundable() bool`

GetPartiallyRefundable returns the PartiallyRefundable field if non-nil, zero value otherwise.

### GetPartiallyRefundableOk

`func (o *Transaction) GetPartiallyRefundableOk() (*bool, bool)`

GetPartiallyRefundableOk returns a tuple with the PartiallyRefundable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPartiallyRefundable

`func (o *Transaction) SetPartiallyRefundable(v bool)`

SetPartiallyRefundable sets PartiallyRefundable field to given value.

### HasPartiallyRefundable

`func (o *Transaction) HasPartiallyRefundable() bool`

HasPartiallyRefundable returns a boolean if a field has been set.

### GetVatRate

`func (o *Transaction) GetVatRate() string`

GetVatRate returns the VatRate field if non-nil, zero value otherwise.

### GetVatRateOk

`func (o *Transaction) GetVatRateOk() (*string, bool)`

GetVatRateOk returns a tuple with the VatRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatRate

`func (o *Transaction) SetVatRate(v string)`

SetVatRate sets VatRate field to given value.

### HasVatRate

`func (o *Transaction) HasVatRate() bool`

HasVatRate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


