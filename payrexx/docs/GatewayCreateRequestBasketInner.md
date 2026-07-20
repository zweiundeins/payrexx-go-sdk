# GatewayCreateRequestBasketInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **[]string** | Name of item | [default to {}]
**Description** | Pointer to **[]string** | Description of item | [optional] [default to {}]
**Quantity** | **int32** | Quantity of item | 
**Amount** | **int32** | Amount in cents | 
**VatRate** | Pointer to **float32** | VAT rate of item (overrides vatRate of Gateway, can be 0) | [optional] 
**Sku** | Pointer to **string** | Product stock keeping unit | [optional] 

## Methods

### NewGatewayCreateRequestBasketInner

`func NewGatewayCreateRequestBasketInner(name []string, quantity int32, amount int32, ) *GatewayCreateRequestBasketInner`

NewGatewayCreateRequestBasketInner instantiates a new GatewayCreateRequestBasketInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayCreateRequestBasketInnerWithDefaults

`func NewGatewayCreateRequestBasketInnerWithDefaults() *GatewayCreateRequestBasketInner`

NewGatewayCreateRequestBasketInnerWithDefaults instantiates a new GatewayCreateRequestBasketInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GatewayCreateRequestBasketInner) GetName() []string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GatewayCreateRequestBasketInner) GetNameOk() (*[]string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GatewayCreateRequestBasketInner) SetName(v []string)`

SetName sets Name field to given value.


### GetDescription

`func (o *GatewayCreateRequestBasketInner) GetDescription() []string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GatewayCreateRequestBasketInner) GetDescriptionOk() (*[]string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GatewayCreateRequestBasketInner) SetDescription(v []string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GatewayCreateRequestBasketInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetQuantity

`func (o *GatewayCreateRequestBasketInner) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *GatewayCreateRequestBasketInner) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *GatewayCreateRequestBasketInner) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.


### GetAmount

`func (o *GatewayCreateRequestBasketInner) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *GatewayCreateRequestBasketInner) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *GatewayCreateRequestBasketInner) SetAmount(v int32)`

SetAmount sets Amount field to given value.


### GetVatRate

`func (o *GatewayCreateRequestBasketInner) GetVatRate() float32`

GetVatRate returns the VatRate field if non-nil, zero value otherwise.

### GetVatRateOk

`func (o *GatewayCreateRequestBasketInner) GetVatRateOk() (*float32, bool)`

GetVatRateOk returns a tuple with the VatRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVatRate

`func (o *GatewayCreateRequestBasketInner) SetVatRate(v float32)`

SetVatRate sets VatRate field to given value.

### HasVatRate

`func (o *GatewayCreateRequestBasketInner) HasVatRate() bool`

HasVatRate returns a boolean if a field has been set.

### GetSku

`func (o *GatewayCreateRequestBasketInner) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *GatewayCreateRequestBasketInner) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *GatewayCreateRequestBasketInner) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *GatewayCreateRequestBasketInner) HasSku() bool`

HasSku returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


