# PayoutListRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Limit** | Pointer to **int32** | Max row count | [optional] 
**Offset** | Pointer to **int32** | Row count to be used as offset | [optional] 
**OrderByDate** | Pointer to **string** | ASC (ascending) or DESC (descending) for ordering by date of the payout | [optional] [default to "Defaults to ASC"]

## Methods

### NewPayoutListRequest

`func NewPayoutListRequest() *PayoutListRequest`

NewPayoutListRequest instantiates a new PayoutListRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayoutListRequestWithDefaults

`func NewPayoutListRequestWithDefaults() *PayoutListRequest`

NewPayoutListRequestWithDefaults instantiates a new PayoutListRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLimit

`func (o *PayoutListRequest) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *PayoutListRequest) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *PayoutListRequest) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *PayoutListRequest) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *PayoutListRequest) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *PayoutListRequest) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *PayoutListRequest) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *PayoutListRequest) HasOffset() bool`

HasOffset returns a boolean if a field has been set.

### GetOrderByDate

`func (o *PayoutListRequest) GetOrderByDate() string`

GetOrderByDate returns the OrderByDate field if non-nil, zero value otherwise.

### GetOrderByDateOk

`func (o *PayoutListRequest) GetOrderByDateOk() (*string, bool)`

GetOrderByDateOk returns a tuple with the OrderByDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderByDate

`func (o *PayoutListRequest) SetOrderByDate(v string)`

SetOrderByDate sets OrderByDate field to given value.

### HasOrderByDate

`func (o *PayoutListRequest) HasOrderByDate() bool`

HasOrderByDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


