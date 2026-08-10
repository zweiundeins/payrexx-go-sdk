# EcrPaymentRefundRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | Pointer to **int32** | In cents. Optional — omit or send 0 for a full refund. Newland only accepts the full transaction amount; a smaller value is rejected. | [optional] 

## Methods

### NewEcrPaymentRefundRequest

`func NewEcrPaymentRefundRequest() *EcrPaymentRefundRequest`

NewEcrPaymentRefundRequest instantiates a new EcrPaymentRefundRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEcrPaymentRefundRequestWithDefaults

`func NewEcrPaymentRefundRequestWithDefaults() *EcrPaymentRefundRequest`

NewEcrPaymentRefundRequestWithDefaults instantiates a new EcrPaymentRefundRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *EcrPaymentRefundRequest) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *EcrPaymentRefundRequest) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *EcrPaymentRefundRequest) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *EcrPaymentRefundRequest) HasAmount() bool`

HasAmount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


