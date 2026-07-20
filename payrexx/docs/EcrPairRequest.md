# EcrPairRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PairingCode** | **string** |  | 
**CashierName** | Pointer to **string** |  | [optional] [default to "terminalName"]

## Methods

### NewEcrPairRequest

`func NewEcrPairRequest(pairingCode string, ) *EcrPairRequest`

NewEcrPairRequest instantiates a new EcrPairRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEcrPairRequestWithDefaults

`func NewEcrPairRequestWithDefaults() *EcrPairRequest`

NewEcrPairRequestWithDefaults instantiates a new EcrPairRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPairingCode

`func (o *EcrPairRequest) GetPairingCode() string`

GetPairingCode returns the PairingCode field if non-nil, zero value otherwise.

### GetPairingCodeOk

`func (o *EcrPairRequest) GetPairingCodeOk() (*string, bool)`

GetPairingCodeOk returns a tuple with the PairingCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPairingCode

`func (o *EcrPairRequest) SetPairingCode(v string)`

SetPairingCode sets PairingCode field to given value.


### GetCashierName

`func (o *EcrPairRequest) GetCashierName() string`

GetCashierName returns the CashierName field if non-nil, zero value otherwise.

### GetCashierNameOk

`func (o *EcrPairRequest) GetCashierNameOk() (*string, bool)`

GetCashierNameOk returns a tuple with the CashierName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashierName

`func (o *EcrPairRequest) SetCashierName(v string)`

SetCashierName sets CashierName field to given value.

### HasCashierName

`func (o *EcrPairRequest) HasCashierName() bool`

HasCashierName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


