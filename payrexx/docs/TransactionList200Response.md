# TransactionList200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**Data** | Pointer to [**[]Transaction**](Transaction.md) |  | [optional] 

## Methods

### NewTransactionList200Response

`func NewTransactionList200Response() *TransactionList200Response`

NewTransactionList200Response instantiates a new TransactionList200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionList200ResponseWithDefaults

`func NewTransactionList200ResponseWithDefaults() *TransactionList200Response`

NewTransactionList200ResponseWithDefaults instantiates a new TransactionList200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *TransactionList200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TransactionList200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TransactionList200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TransactionList200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetData

`func (o *TransactionList200Response) GetData() []Transaction`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TransactionList200Response) GetDataOk() (*[]Transaction, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TransactionList200Response) SetData(v []Transaction)`

SetData sets Data field to given value.

### HasData

`func (o *TransactionList200Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


