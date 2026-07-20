# BillCreateRequestRemindersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Days** | **int32** |  | 
**Fee** | **int32** |  | 
**Friendly** | **bool** |  | 
**DelayType** | Pointer to **string** | Possible values are: daysAfterInvoiceDate, daysBeforeDueDate, daysAfterDueDate | [optional] [default to "daysAfterInvoiceDate"]

## Methods

### NewBillCreateRequestRemindersInner

`func NewBillCreateRequestRemindersInner(days int32, fee int32, friendly bool, ) *BillCreateRequestRemindersInner`

NewBillCreateRequestRemindersInner instantiates a new BillCreateRequestRemindersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillCreateRequestRemindersInnerWithDefaults

`func NewBillCreateRequestRemindersInnerWithDefaults() *BillCreateRequestRemindersInner`

NewBillCreateRequestRemindersInnerWithDefaults instantiates a new BillCreateRequestRemindersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDays

`func (o *BillCreateRequestRemindersInner) GetDays() int32`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *BillCreateRequestRemindersInner) GetDaysOk() (*int32, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *BillCreateRequestRemindersInner) SetDays(v int32)`

SetDays sets Days field to given value.


### GetFee

`func (o *BillCreateRequestRemindersInner) GetFee() int32`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *BillCreateRequestRemindersInner) GetFeeOk() (*int32, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *BillCreateRequestRemindersInner) SetFee(v int32)`

SetFee sets Fee field to given value.


### GetFriendly

`func (o *BillCreateRequestRemindersInner) GetFriendly() bool`

GetFriendly returns the Friendly field if non-nil, zero value otherwise.

### GetFriendlyOk

`func (o *BillCreateRequestRemindersInner) GetFriendlyOk() (*bool, bool)`

GetFriendlyOk returns a tuple with the Friendly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFriendly

`func (o *BillCreateRequestRemindersInner) SetFriendly(v bool)`

SetFriendly sets Friendly field to given value.


### GetDelayType

`func (o *BillCreateRequestRemindersInner) GetDelayType() string`

GetDelayType returns the DelayType field if non-nil, zero value otherwise.

### GetDelayTypeOk

`func (o *BillCreateRequestRemindersInner) GetDelayTypeOk() (*string, bool)`

GetDelayTypeOk returns a tuple with the DelayType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayType

`func (o *BillCreateRequestRemindersInner) SetDelayType(v string)`

SetDelayType sets DelayType field to given value.

### HasDelayType

`func (o *BillCreateRequestRemindersInner) HasDelayType() bool`

HasDelayType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


