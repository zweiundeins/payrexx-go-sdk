# Payout

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** |  | [optional] 
**Mode** | Pointer to **string** |  | [optional] 
**Object** | Pointer to **string** |  | [optional] 
**Amount** | Pointer to **int32** |  | [optional] [default to 0]
**TotalFees** | Pointer to **int32** |  | [optional] [default to 0]
**Currency** | Pointer to **string** |  | [optional] 
**Date** | Pointer to **string** |  | [optional] 
**Statement** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Destination** | Pointer to [**PayoutDestination**](PayoutDestination.md) |  | [optional] 
**Transfers** | Pointer to **[]string** |  | [optional] 
**Merchant** | Pointer to [**PayoutMerchant**](PayoutMerchant.md) |  | [optional] 

## Methods

### NewPayout

`func NewPayout() *Payout`

NewPayout instantiates a new Payout object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayoutWithDefaults

`func NewPayoutWithDefaults() *Payout`

NewPayoutWithDefaults instantiates a new Payout object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *Payout) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Payout) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Payout) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *Payout) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetMode

`func (o *Payout) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *Payout) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *Payout) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *Payout) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetObject

`func (o *Payout) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *Payout) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *Payout) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *Payout) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetAmount

`func (o *Payout) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *Payout) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *Payout) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *Payout) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetTotalFees

`func (o *Payout) GetTotalFees() int32`

GetTotalFees returns the TotalFees field if non-nil, zero value otherwise.

### GetTotalFeesOk

`func (o *Payout) GetTotalFeesOk() (*int32, bool)`

GetTotalFeesOk returns a tuple with the TotalFees field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalFees

`func (o *Payout) SetTotalFees(v int32)`

SetTotalFees sets TotalFees field to given value.

### HasTotalFees

`func (o *Payout) HasTotalFees() bool`

HasTotalFees returns a boolean if a field has been set.

### GetCurrency

`func (o *Payout) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Payout) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Payout) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *Payout) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDate

`func (o *Payout) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *Payout) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *Payout) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *Payout) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetStatement

`func (o *Payout) GetStatement() string`

GetStatement returns the Statement field if non-nil, zero value otherwise.

### GetStatementOk

`func (o *Payout) GetStatementOk() (*string, bool)`

GetStatementOk returns a tuple with the Statement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatement

`func (o *Payout) SetStatement(v string)`

SetStatement sets Statement field to given value.

### HasStatement

`func (o *Payout) HasStatement() bool`

HasStatement returns a boolean if a field has been set.

### GetStatus

`func (o *Payout) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Payout) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Payout) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Payout) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDestination

`func (o *Payout) GetDestination() PayoutDestination`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *Payout) GetDestinationOk() (*PayoutDestination, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *Payout) SetDestination(v PayoutDestination)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *Payout) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetTransfers

`func (o *Payout) GetTransfers() []string`

GetTransfers returns the Transfers field if non-nil, zero value otherwise.

### GetTransfersOk

`func (o *Payout) GetTransfersOk() (*[]string, bool)`

GetTransfersOk returns a tuple with the Transfers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransfers

`func (o *Payout) SetTransfers(v []string)`

SetTransfers sets Transfers field to given value.

### HasTransfers

`func (o *Payout) HasTransfers() bool`

HasTransfers returns a boolean if a field has been set.

### GetMerchant

`func (o *Payout) GetMerchant() PayoutMerchant`

GetMerchant returns the Merchant field if non-nil, zero value otherwise.

### GetMerchantOk

`func (o *Payout) GetMerchantOk() (*PayoutMerchant, bool)`

GetMerchantOk returns a tuple with the Merchant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchant

`func (o *Payout) SetMerchant(v PayoutMerchant)`

SetMerchant sets Merchant field to given value.

### HasMerchant

`func (o *Payout) HasMerchant() bool`

HasMerchant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


