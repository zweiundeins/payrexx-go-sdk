# PayoutDestination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**Iban** | Pointer to **string** |  | [optional] 
**AccountHolder** | Pointer to **string** |  | [optional] 

## Methods

### NewPayoutDestination

`func NewPayoutDestination() *PayoutDestination`

NewPayoutDestination instantiates a new PayoutDestination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayoutDestinationWithDefaults

`func NewPayoutDestinationWithDefaults() *PayoutDestination`

NewPayoutDestinationWithDefaults instantiates a new PayoutDestination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PayoutDestination) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PayoutDestination) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PayoutDestination) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PayoutDestination) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIban

`func (o *PayoutDestination) GetIban() string`

GetIban returns the Iban field if non-nil, zero value otherwise.

### GetIbanOk

`func (o *PayoutDestination) GetIbanOk() (*string, bool)`

GetIbanOk returns a tuple with the Iban field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIban

`func (o *PayoutDestination) SetIban(v string)`

SetIban sets Iban field to given value.

### HasIban

`func (o *PayoutDestination) HasIban() bool`

HasIban returns a boolean if a field has been set.

### GetAccountHolder

`func (o *PayoutDestination) GetAccountHolder() string`

GetAccountHolder returns the AccountHolder field if non-nil, zero value otherwise.

### GetAccountHolderOk

`func (o *PayoutDestination) GetAccountHolderOk() (*string, bool)`

GetAccountHolderOk returns a tuple with the AccountHolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountHolder

`func (o *PayoutDestination) SetAccountHolder(v string)`

SetAccountHolder sets AccountHolder field to given value.

### HasAccountHolder

`func (o *PayoutDestination) HasAccountHolder() bool`

HasAccountHolder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


