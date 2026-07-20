# InvoicePositionsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Number** | Pointer to **int32** |  | [optional] [default to 0]
**Price** | Pointer to **int32** |  | [optional] [default to 0]
**Discount** | Pointer to [**InvoiceDiscount**](InvoiceDiscount.md) |  | [optional] 
**Vat** | Pointer to **int32** |  | [optional] [default to 0]
**Id** | Pointer to **string** |  | [optional] 

## Methods

### NewInvoicePositionsInner

`func NewInvoicePositionsInner() *InvoicePositionsInner`

NewInvoicePositionsInner instantiates a new InvoicePositionsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoicePositionsInnerWithDefaults

`func NewInvoicePositionsInnerWithDefaults() *InvoicePositionsInner`

NewInvoicePositionsInnerWithDefaults instantiates a new InvoicePositionsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *InvoicePositionsInner) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *InvoicePositionsInner) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *InvoicePositionsInner) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *InvoicePositionsInner) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetDescription

`func (o *InvoicePositionsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InvoicePositionsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InvoicePositionsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InvoicePositionsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *InvoicePositionsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InvoicePositionsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InvoicePositionsInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *InvoicePositionsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetNumber

`func (o *InvoicePositionsInner) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *InvoicePositionsInner) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *InvoicePositionsInner) SetNumber(v int32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *InvoicePositionsInner) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetPrice

`func (o *InvoicePositionsInner) GetPrice() int32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *InvoicePositionsInner) GetPriceOk() (*int32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *InvoicePositionsInner) SetPrice(v int32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *InvoicePositionsInner) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetDiscount

`func (o *InvoicePositionsInner) GetDiscount() InvoiceDiscount`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *InvoicePositionsInner) GetDiscountOk() (*InvoiceDiscount, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *InvoicePositionsInner) SetDiscount(v InvoiceDiscount)`

SetDiscount sets Discount field to given value.

### HasDiscount

`func (o *InvoicePositionsInner) HasDiscount() bool`

HasDiscount returns a boolean if a field has been set.

### GetVat

`func (o *InvoicePositionsInner) GetVat() int32`

GetVat returns the Vat field if non-nil, zero value otherwise.

### GetVatOk

`func (o *InvoicePositionsInner) GetVatOk() (*int32, bool)`

GetVatOk returns a tuple with the Vat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVat

`func (o *InvoicePositionsInner) SetVat(v int32)`

SetVat sets Vat field to given value.

### HasVat

`func (o *InvoicePositionsInner) HasVat() bool`

HasVat returns a boolean if a field has been set.

### GetId

`func (o *InvoicePositionsInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InvoicePositionsInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InvoicePositionsInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *InvoicePositionsInner) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


