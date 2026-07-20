# EcrPairing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PairingStatus** | Pointer to **string** |  | [optional] 
**CashierName** | Pointer to **string** |  | [optional] 
**Configuration** | Pointer to [**EcrPairingConfiguration**](EcrPairingConfiguration.md) |  | [optional] 

## Methods

### NewEcrPairing

`func NewEcrPairing() *EcrPairing`

NewEcrPairing instantiates a new EcrPairing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEcrPairingWithDefaults

`func NewEcrPairingWithDefaults() *EcrPairing`

NewEcrPairingWithDefaults instantiates a new EcrPairing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPairingStatus

`func (o *EcrPairing) GetPairingStatus() string`

GetPairingStatus returns the PairingStatus field if non-nil, zero value otherwise.

### GetPairingStatusOk

`func (o *EcrPairing) GetPairingStatusOk() (*string, bool)`

GetPairingStatusOk returns a tuple with the PairingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPairingStatus

`func (o *EcrPairing) SetPairingStatus(v string)`

SetPairingStatus sets PairingStatus field to given value.

### HasPairingStatus

`func (o *EcrPairing) HasPairingStatus() bool`

HasPairingStatus returns a boolean if a field has been set.

### GetCashierName

`func (o *EcrPairing) GetCashierName() string`

GetCashierName returns the CashierName field if non-nil, zero value otherwise.

### GetCashierNameOk

`func (o *EcrPairing) GetCashierNameOk() (*string, bool)`

GetCashierNameOk returns a tuple with the CashierName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashierName

`func (o *EcrPairing) SetCashierName(v string)`

SetCashierName sets CashierName field to given value.

### HasCashierName

`func (o *EcrPairing) HasCashierName() bool`

HasCashierName returns a boolean if a field has been set.

### GetConfiguration

`func (o *EcrPairing) GetConfiguration() EcrPairingConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *EcrPairing) GetConfigurationOk() (*EcrPairingConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *EcrPairing) SetConfiguration(v EcrPairingConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *EcrPairing) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


