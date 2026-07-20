# BillCreate200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**Data** | Pointer to [**Invoice**](Invoice.md) |  | [optional] 
**Meta** | Pointer to [**BillCreate200ResponseMeta**](BillCreate200ResponseMeta.md) |  | [optional] 

## Methods

### NewBillCreate200Response

`func NewBillCreate200Response() *BillCreate200Response`

NewBillCreate200Response instantiates a new BillCreate200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillCreate200ResponseWithDefaults

`func NewBillCreate200ResponseWithDefaults() *BillCreate200Response`

NewBillCreate200ResponseWithDefaults instantiates a new BillCreate200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *BillCreate200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BillCreate200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BillCreate200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BillCreate200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetData

`func (o *BillCreate200Response) GetData() Invoice`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BillCreate200Response) GetDataOk() (*Invoice, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BillCreate200Response) SetData(v Invoice)`

SetData sets Data field to given value.

### HasData

`func (o *BillCreate200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *BillCreate200Response) GetMeta() BillCreate200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *BillCreate200Response) GetMetaOk() (*BillCreate200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *BillCreate200Response) SetMeta(v BillCreate200ResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *BillCreate200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


