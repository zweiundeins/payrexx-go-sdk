# BillCreateRequestDiscount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Amount** | **int32** | If the type is \&quot;amount\&quot;, this amount here needs to be in the smallest unit of the currency. | 
**Type** | **string** | either \&quot;percent\&quot; or \&quot;amount\&quot; | 

## Methods

### NewBillCreateRequestDiscount

`func NewBillCreateRequestDiscount(amount int32, type_ string, ) *BillCreateRequestDiscount`

NewBillCreateRequestDiscount instantiates a new BillCreateRequestDiscount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillCreateRequestDiscountWithDefaults

`func NewBillCreateRequestDiscountWithDefaults() *BillCreateRequestDiscount`

NewBillCreateRequestDiscountWithDefaults instantiates a new BillCreateRequestDiscount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmount

`func (o *BillCreateRequestDiscount) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *BillCreateRequestDiscount) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *BillCreateRequestDiscount) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetType

`func (o *BillCreateRequestDiscount) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BillCreateRequestDiscount) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BillCreateRequestDiscount) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


