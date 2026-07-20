# SubscriptionListRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderByStartDate** | Pointer to **string** | ASC (ascending) or DESC (descending) for ordering by  start date of the transaction | [optional] [default to "ASC"]
**Limit** | Pointer to **int32** | Max row count | [optional] 
**Offset** | Pointer to **int32** | Row count to be used as offset | [optional] 

## Methods

### NewSubscriptionListRequest

`func NewSubscriptionListRequest() *SubscriptionListRequest`

NewSubscriptionListRequest instantiates a new SubscriptionListRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionListRequestWithDefaults

`func NewSubscriptionListRequestWithDefaults() *SubscriptionListRequest`

NewSubscriptionListRequestWithDefaults instantiates a new SubscriptionListRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderByStartDate

`func (o *SubscriptionListRequest) GetOrderByStartDate() string`

GetOrderByStartDate returns the OrderByStartDate field if non-nil, zero value otherwise.

### GetOrderByStartDateOk

`func (o *SubscriptionListRequest) GetOrderByStartDateOk() (*string, bool)`

GetOrderByStartDateOk returns a tuple with the OrderByStartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderByStartDate

`func (o *SubscriptionListRequest) SetOrderByStartDate(v string)`

SetOrderByStartDate sets OrderByStartDate field to given value.

### HasOrderByStartDate

`func (o *SubscriptionListRequest) HasOrderByStartDate() bool`

HasOrderByStartDate returns a boolean if a field has been set.

### GetLimit

`func (o *SubscriptionListRequest) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *SubscriptionListRequest) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *SubscriptionListRequest) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *SubscriptionListRequest) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetOffset

`func (o *SubscriptionListRequest) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *SubscriptionListRequest) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *SubscriptionListRequest) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *SubscriptionListRequest) HasOffset() bool`

HasOffset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


