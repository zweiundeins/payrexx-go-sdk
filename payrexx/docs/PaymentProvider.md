# PaymentProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [default to 0]
**Name** | Pointer to **string** |  | [optional] 
**AvailableBalance** | Pointer to **interface{}** |  | [optional] 
**PaymentMethods** | Pointer to **[]string** |  | [optional] 
**ActivePaymentMethods** | Pointer to **[]string** |  | [optional] 

## Methods

### NewPaymentProvider

`func NewPaymentProvider() *PaymentProvider`

NewPaymentProvider instantiates a new PaymentProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentProviderWithDefaults

`func NewPaymentProviderWithDefaults() *PaymentProvider`

NewPaymentProviderWithDefaults instantiates a new PaymentProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PaymentProvider) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentProvider) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentProvider) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *PaymentProvider) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *PaymentProvider) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PaymentProvider) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PaymentProvider) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PaymentProvider) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAvailableBalance

`func (o *PaymentProvider) GetAvailableBalance() interface{}`

GetAvailableBalance returns the AvailableBalance field if non-nil, zero value otherwise.

### GetAvailableBalanceOk

`func (o *PaymentProvider) GetAvailableBalanceOk() (*interface{}, bool)`

GetAvailableBalanceOk returns a tuple with the AvailableBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableBalance

`func (o *PaymentProvider) SetAvailableBalance(v interface{})`

SetAvailableBalance sets AvailableBalance field to given value.

### HasAvailableBalance

`func (o *PaymentProvider) HasAvailableBalance() bool`

HasAvailableBalance returns a boolean if a field has been set.

### SetAvailableBalanceNil

`func (o *PaymentProvider) SetAvailableBalanceNil(b bool)`

 SetAvailableBalanceNil sets the value for AvailableBalance to be an explicit nil

### UnsetAvailableBalance
`func (o *PaymentProvider) UnsetAvailableBalance()`

UnsetAvailableBalance ensures that no value is present for AvailableBalance, not even an explicit nil
### GetPaymentMethods

`func (o *PaymentProvider) GetPaymentMethods() []string`

GetPaymentMethods returns the PaymentMethods field if non-nil, zero value otherwise.

### GetPaymentMethodsOk

`func (o *PaymentProvider) GetPaymentMethodsOk() (*[]string, bool)`

GetPaymentMethodsOk returns a tuple with the PaymentMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentMethods

`func (o *PaymentProvider) SetPaymentMethods(v []string)`

SetPaymentMethods sets PaymentMethods field to given value.

### HasPaymentMethods

`func (o *PaymentProvider) HasPaymentMethods() bool`

HasPaymentMethods returns a boolean if a field has been set.

### GetActivePaymentMethods

`func (o *PaymentProvider) GetActivePaymentMethods() []string`

GetActivePaymentMethods returns the ActivePaymentMethods field if non-nil, zero value otherwise.

### GetActivePaymentMethodsOk

`func (o *PaymentProvider) GetActivePaymentMethodsOk() (*[]string, bool)`

GetActivePaymentMethodsOk returns a tuple with the ActivePaymentMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivePaymentMethods

`func (o *PaymentProvider) SetActivePaymentMethods(v []string)`

SetActivePaymentMethods sets ActivePaymentMethods field to given value.

### HasActivePaymentMethods

`func (o *PaymentProvider) HasActivePaymentMethods() bool`

HasActivePaymentMethods returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


