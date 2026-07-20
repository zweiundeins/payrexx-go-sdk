# EcrPayment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentId** | Pointer to **string** |  | [optional] 
**PaymentStatus** | Pointer to **string** |  | [optional] 
**Slip** | Pointer to **interface{}** | Free-form: object on the payment pages, array on the paymentMethods page. | [optional] 

## Methods

### NewEcrPayment

`func NewEcrPayment() *EcrPayment`

NewEcrPayment instantiates a new EcrPayment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEcrPaymentWithDefaults

`func NewEcrPaymentWithDefaults() *EcrPayment`

NewEcrPaymentWithDefaults instantiates a new EcrPayment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentId

`func (o *EcrPayment) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *EcrPayment) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *EcrPayment) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.

### HasPaymentId

`func (o *EcrPayment) HasPaymentId() bool`

HasPaymentId returns a boolean if a field has been set.

### GetPaymentStatus

`func (o *EcrPayment) GetPaymentStatus() string`

GetPaymentStatus returns the PaymentStatus field if non-nil, zero value otherwise.

### GetPaymentStatusOk

`func (o *EcrPayment) GetPaymentStatusOk() (*string, bool)`

GetPaymentStatusOk returns a tuple with the PaymentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentStatus

`func (o *EcrPayment) SetPaymentStatus(v string)`

SetPaymentStatus sets PaymentStatus field to given value.

### HasPaymentStatus

`func (o *EcrPayment) HasPaymentStatus() bool`

HasPaymentStatus returns a boolean if a field has been set.

### GetSlip

`func (o *EcrPayment) GetSlip() interface{}`

GetSlip returns the Slip field if non-nil, zero value otherwise.

### GetSlipOk

`func (o *EcrPayment) GetSlipOk() (*interface{}, bool)`

GetSlipOk returns a tuple with the Slip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlip

`func (o *EcrPayment) SetSlip(v interface{})`

SetSlip sets Slip field to given value.

### HasSlip

`func (o *EcrPayment) HasSlip() bool`

HasSlip returns a boolean if a field has been set.

### SetSlipNil

`func (o *EcrPayment) SetSlipNil(b bool)`

 SetSlipNil sets the value for Slip to be an explicit nil

### UnsetSlip
`func (o *EcrPayment) UnsetSlip()`

UnsetSlip ensures that no value is present for Slip, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


