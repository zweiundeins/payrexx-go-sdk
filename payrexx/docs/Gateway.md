# Gateway

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [default to 0]
**Status** | Pointer to **string** |  | [optional] 
**Hash** | Pointer to **string** |  | [optional] 
**ReferenceId** | Pointer to **string** |  | [optional] 
**Link** | Pointer to **string** |  | [optional] 
**Invoices** | Pointer to **[]string** |  | [optional] 
**PreAuthorization** | Pointer to **bool** | Corrected against the live API: documented as an integer; POST /Gateway/ returns false, and the request schema for the same field is already a boolean. | [optional] 
**Reservation** | Pointer to **bool** | Corrected against the live API: documented as an integer, but it is preAuthorization&#39;s twin and the request schema types it as a boolean. | [optional] 
**Fields** | Pointer to **interface{}** | Free-form: an object of per-field settings, or [] when none are set. Corrected against the live API: documented as an object; PHP serializes the empty case as [], so a real response is an object or an empty array depending on the gateway. | [optional] 
**Psp** | Pointer to **[]string** |  | [optional] 
**Pm** | Pointer to **[]string** |  | [optional] 
**Amount** | Pointer to **int32** |  | [optional] [default to 0]
**VatRate** | Pointer to **float32** |  | [optional] [default to 0]
**Currency** | Pointer to **string** |  | [optional] 
**Sku** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **int32** |  | [optional] [default to 0]
**ApplicationFee** | Pointer to **int32** |  | [optional] [default to 0]
**Language** | Pointer to **string** | Corrected against the live API: returned by POST and GET /Gateway/ but absent from both pages. | [optional] 
**RequestId** | Pointer to **int32** | Corrected against the live API: returned by POST /Gateway/ but absent from the page. | [optional] 

## Methods

### NewGateway

`func NewGateway() *Gateway`

NewGateway instantiates a new Gateway object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayWithDefaults

`func NewGatewayWithDefaults() *Gateway`

NewGatewayWithDefaults instantiates a new Gateway object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Gateway) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Gateway) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Gateway) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Gateway) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *Gateway) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Gateway) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Gateway) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Gateway) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetHash

`func (o *Gateway) GetHash() string`

GetHash returns the Hash field if non-nil, zero value otherwise.

### GetHashOk

`func (o *Gateway) GetHashOk() (*string, bool)`

GetHashOk returns a tuple with the Hash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHash

`func (o *Gateway) SetHash(v string)`

SetHash sets Hash field to given value.

### HasHash

`func (o *Gateway) HasHash() bool`

HasHash returns a boolean if a field has been set.

### GetReferenceId

`func (o *Gateway) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *Gateway) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *Gateway) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *Gateway) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.

### GetLink

`func (o *Gateway) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *Gateway) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *Gateway) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *Gateway) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetInvoices

`func (o *Gateway) GetInvoices() []string`

GetInvoices returns the Invoices field if non-nil, zero value otherwise.

### GetInvoicesOk

`func (o *Gateway) GetInvoicesOk() (*[]string, bool)`

GetInvoicesOk returns a tuple with the Invoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoices

`func (o *Gateway) SetInvoices(v []string)`

SetInvoices sets Invoices field to given value.

### HasInvoices

`func (o *Gateway) HasInvoices() bool`

HasInvoices returns a boolean if a field has been set.

### GetPreAuthorization

`func (o *Gateway) GetPreAuthorization() bool`

GetPreAuthorization returns the PreAuthorization field if non-nil, zero value otherwise.

### GetPreAuthorizationOk

`func (o *Gateway) GetPreAuthorizationOk() (*bool, bool)`

GetPreAuthorizationOk returns a tuple with the PreAuthorization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreAuthorization

`func (o *Gateway) SetPreAuthorization(v bool)`

SetPreAuthorization sets PreAuthorization field to given value.

### HasPreAuthorization

`func (o *Gateway) HasPreAuthorization() bool`

HasPreAuthorization returns a boolean if a field has been set.

### GetReservation

`func (o *Gateway) GetReservation() bool`

GetReservation returns the Reservation field if non-nil, zero value otherwise.

### GetReservationOk

`func (o *Gateway) GetReservationOk() (*bool, bool)`

GetReservationOk returns a tuple with the Reservation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservation

`func (o *Gateway) SetReservation(v bool)`

SetReservation sets Reservation field to given value.

### HasReservation

`func (o *Gateway) HasReservation() bool`

HasReservation returns a boolean if a field has been set.

### GetFields

`func (o *Gateway) GetFields() interface{}`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *Gateway) GetFieldsOk() (*interface{}, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *Gateway) SetFields(v interface{})`

SetFields sets Fields field to given value.

### HasFields

`func (o *Gateway) HasFields() bool`

HasFields returns a boolean if a field has been set.

### SetFieldsNil

`func (o *Gateway) SetFieldsNil(b bool)`

 SetFieldsNil sets the value for Fields to be an explicit nil

### UnsetFields
`func (o *Gateway) UnsetFields()`

UnsetFields ensures that no value is present for Fields, not even an explicit nil
### GetPsp

`func (o *Gateway) GetPsp() []string`

GetPsp returns the Psp field if non-nil, zero value otherwise.

### GetPspOk

`func (o *Gateway) GetPspOk() (*[]string, bool)`

GetPspOk returns a tuple with the Psp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPsp

`func (o *Gateway) SetPsp(v []string)`

SetPsp sets Psp field to given value.

### HasPsp

`func (o *Gateway) HasPsp() bool`

HasPsp returns a boolean if a field has been set.

### GetPm

`func (o *Gateway) GetPm() []string`

GetPm returns the Pm field if non-nil, zero value otherwise.

### GetPmOk

`func (o *Gateway) GetPmOk() (*[]string, bool)`

GetPmOk returns a tuple with the Pm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPm

`func (o *Gateway) SetPm(v []string)`

SetPm sets Pm field to given value.

### HasPm

`func (o *Gateway) HasPm() bool`

HasPm returns a boolean if a field has been set.

### GetAmount

`func (o *Gateway) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Gateway) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Gateway) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *Gateway) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetVatRate

`func (o *Gateway) GetVatRate() float32`

GetVatRate returns the VatRate field if non-nil, zero value otherwise.

### GetVatRateOk

`func (o *Gateway) GetVatRateOk() (*float32, bool)`

GetVatRateOk returns a tuple with the VatRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatRate

`func (o *Gateway) SetVatRate(v float32)`

SetVatRate sets VatRate field to given value.

### HasVatRate

`func (o *Gateway) HasVatRate() bool`

HasVatRate returns a boolean if a field has been set.

### GetCurrency

`func (o *Gateway) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Gateway) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Gateway) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *Gateway) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetSku

`func (o *Gateway) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *Gateway) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *Gateway) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *Gateway) HasSku() bool`

HasSku returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Gateway) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Gateway) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Gateway) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Gateway) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetApplicationFee

`func (o *Gateway) GetApplicationFee() int32`

GetApplicationFee returns the ApplicationFee field if non-nil, zero value otherwise.

### GetApplicationFeeOk

`func (o *Gateway) GetApplicationFeeOk() (*int32, bool)`

GetApplicationFeeOk returns a tuple with the ApplicationFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationFee

`func (o *Gateway) SetApplicationFee(v int32)`

SetApplicationFee sets ApplicationFee field to given value.

### HasApplicationFee

`func (o *Gateway) HasApplicationFee() bool`

HasApplicationFee returns a boolean if a field has been set.

### GetLanguage

`func (o *Gateway) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *Gateway) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *Gateway) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *Gateway) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetRequestId

`func (o *Gateway) GetRequestId() int32`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *Gateway) GetRequestIdOk() (*int32, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *Gateway) SetRequestId(v int32)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *Gateway) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


