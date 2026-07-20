# TransactionListRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilterDatetimeUtcGreaterThan** | Pointer to **string** | Lower DateTime limit in the format of YYYY-MM-DD HH:MM:SS | [optional] 
**FilterDatetimeUtcLessThan** | Pointer to **string** | Upper DateTime limit in the format of YYYY-MM-DD HH:MM:SS | [optional] 
**FilterMyTransactionsOnly** | Pointer to **bool** | When set to 1, only transactions related to this API key used will be returned. | [optional] [default to false]
**OrderByTime** | Pointer to **string** | ASC (ascending) or DESC (descending) for ordering by time of the transaction | [optional] [default to "ASC"]
**Offset** | Pointer to **int32** | Row count to be used as offset | [optional] 
**Limit** | Pointer to **int32** | Max row count | [optional] 

## Methods

### NewTransactionListRequest

`func NewTransactionListRequest() *TransactionListRequest`

NewTransactionListRequest instantiates a new TransactionListRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionListRequestWithDefaults

`func NewTransactionListRequestWithDefaults() *TransactionListRequest`

NewTransactionListRequestWithDefaults instantiates a new TransactionListRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilterDatetimeUtcGreaterThan

`func (o *TransactionListRequest) GetFilterDatetimeUtcGreaterThan() string`

GetFilterDatetimeUtcGreaterThan returns the FilterDatetimeUtcGreaterThan field if non-nil, zero value otherwise.

### GetFilterDatetimeUtcGreaterThanOk

`func (o *TransactionListRequest) GetFilterDatetimeUtcGreaterThanOk() (*string, bool)`

GetFilterDatetimeUtcGreaterThanOk returns a tuple with the FilterDatetimeUtcGreaterThan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterDatetimeUtcGreaterThan

`func (o *TransactionListRequest) SetFilterDatetimeUtcGreaterThan(v string)`

SetFilterDatetimeUtcGreaterThan sets FilterDatetimeUtcGreaterThan field to given value.

### HasFilterDatetimeUtcGreaterThan

`func (o *TransactionListRequest) HasFilterDatetimeUtcGreaterThan() bool`

HasFilterDatetimeUtcGreaterThan returns a boolean if a field has been set.

### GetFilterDatetimeUtcLessThan

`func (o *TransactionListRequest) GetFilterDatetimeUtcLessThan() string`

GetFilterDatetimeUtcLessThan returns the FilterDatetimeUtcLessThan field if non-nil, zero value otherwise.

### GetFilterDatetimeUtcLessThanOk

`func (o *TransactionListRequest) GetFilterDatetimeUtcLessThanOk() (*string, bool)`

GetFilterDatetimeUtcLessThanOk returns a tuple with the FilterDatetimeUtcLessThan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterDatetimeUtcLessThan

`func (o *TransactionListRequest) SetFilterDatetimeUtcLessThan(v string)`

SetFilterDatetimeUtcLessThan sets FilterDatetimeUtcLessThan field to given value.

### HasFilterDatetimeUtcLessThan

`func (o *TransactionListRequest) HasFilterDatetimeUtcLessThan() bool`

HasFilterDatetimeUtcLessThan returns a boolean if a field has been set.

### GetFilterMyTransactionsOnly

`func (o *TransactionListRequest) GetFilterMyTransactionsOnly() bool`

GetFilterMyTransactionsOnly returns the FilterMyTransactionsOnly field if non-nil, zero value otherwise.

### GetFilterMyTransactionsOnlyOk

`func (o *TransactionListRequest) GetFilterMyTransactionsOnlyOk() (*bool, bool)`

GetFilterMyTransactionsOnlyOk returns a tuple with the FilterMyTransactionsOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterMyTransactionsOnly

`func (o *TransactionListRequest) SetFilterMyTransactionsOnly(v bool)`

SetFilterMyTransactionsOnly sets FilterMyTransactionsOnly field to given value.

### HasFilterMyTransactionsOnly

`func (o *TransactionListRequest) HasFilterMyTransactionsOnly() bool`

HasFilterMyTransactionsOnly returns a boolean if a field has been set.

### GetOrderByTime

`func (o *TransactionListRequest) GetOrderByTime() string`

GetOrderByTime returns the OrderByTime field if non-nil, zero value otherwise.

### GetOrderByTimeOk

`func (o *TransactionListRequest) GetOrderByTimeOk() (*string, bool)`

GetOrderByTimeOk returns a tuple with the OrderByTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderByTime

`func (o *TransactionListRequest) SetOrderByTime(v string)`

SetOrderByTime sets OrderByTime field to given value.

### HasOrderByTime

`func (o *TransactionListRequest) HasOrderByTime() bool`

HasOrderByTime returns a boolean if a field has been set.

### GetOffset

`func (o *TransactionListRequest) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *TransactionListRequest) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *TransactionListRequest) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *TransactionListRequest) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetLimit

`func (o *TransactionListRequest) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *TransactionListRequest) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *TransactionListRequest) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *TransactionListRequest) HasLimit() bool`

HasLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


