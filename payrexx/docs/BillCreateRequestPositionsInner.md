# BillCreateRequestPositionsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Type** | **string** | Types supported: \&quot;piece\&quot;, \&quot;hour\&quot;, \&quot;day\&quot;, \&quot;flat\&quot;. | 
**Number** | Pointer to **float64** |  | [optional] 
**Price** | **int32** | Amount in the smallest unit of the currency. | 
**Discount** | Pointer to [**BillCreateRequestDiscount**](BillCreateRequestDiscount.md) |  | [optional] 
**Vat** | Pointer to **float64** | Vat in percent. | [optional] 

## Methods

### NewBillCreateRequestPositionsInner

`func NewBillCreateRequestPositionsInner(title string, type_ string, price int32, ) *BillCreateRequestPositionsInner`

NewBillCreateRequestPositionsInner instantiates a new BillCreateRequestPositionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillCreateRequestPositionsInnerWithDefaults

`func NewBillCreateRequestPositionsInnerWithDefaults() *BillCreateRequestPositionsInner`

NewBillCreateRequestPositionsInnerWithDefaults instantiates a new BillCreateRequestPositionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *BillCreateRequestPositionsInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *BillCreateRequestPositionsInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *BillCreateRequestPositionsInner) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *BillCreateRequestPositionsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BillCreateRequestPositionsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BillCreateRequestPositionsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BillCreateRequestPositionsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *BillCreateRequestPositionsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BillCreateRequestPositionsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BillCreateRequestPositionsInner) SetType(v string)`

SetType sets Type field to given value.


### GetNumber

`func (o *BillCreateRequestPositionsInner) GetNumber() float64`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *BillCreateRequestPositionsInner) GetNumberOk() (*float64, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *BillCreateRequestPositionsInner) SetNumber(v float64)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *BillCreateRequestPositionsInner) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetPrice

`func (o *BillCreateRequestPositionsInner) GetPrice() int32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *BillCreateRequestPositionsInner) GetPriceOk() (*int32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *BillCreateRequestPositionsInner) SetPrice(v int32)`

SetPrice sets Price field to given value.


### GetDiscount

`func (o *BillCreateRequestPositionsInner) GetDiscount() BillCreateRequestDiscount`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *BillCreateRequestPositionsInner) GetDiscountOk() (*BillCreateRequestDiscount, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *BillCreateRequestPositionsInner) SetDiscount(v BillCreateRequestDiscount)`

SetDiscount sets Discount field to given value.

### HasDiscount

`func (o *BillCreateRequestPositionsInner) HasDiscount() bool`

HasDiscount returns a boolean if a field has been set.

### GetVat

`func (o *BillCreateRequestPositionsInner) GetVat() float64`

GetVat returns the Vat field if non-nil, zero value otherwise.

### GetVatOk

`func (o *BillCreateRequestPositionsInner) GetVatOk() (*float64, bool)`

GetVatOk returns a tuple with the Vat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVat

`func (o *BillCreateRequestPositionsInner) SetVat(v float64)`

SetVat sets Vat field to given value.

### HasVat

`func (o *BillCreateRequestPositionsInner) HasVat() bool`

HasVat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


