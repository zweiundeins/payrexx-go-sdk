# BillDelete200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**Data** | Pointer to [**InvoiceDeletion**](InvoiceDeletion.md) |  | [optional] 
**Meta** | Pointer to [**BillCreate200ResponseMeta**](BillCreate200ResponseMeta.md) |  | [optional] 

## Methods

### NewBillDelete200Response

`func NewBillDelete200Response() *BillDelete200Response`

NewBillDelete200Response instantiates a new BillDelete200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillDelete200ResponseWithDefaults

`func NewBillDelete200ResponseWithDefaults() *BillDelete200Response`

NewBillDelete200ResponseWithDefaults instantiates a new BillDelete200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *BillDelete200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BillDelete200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BillDelete200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BillDelete200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetData

`func (o *BillDelete200Response) GetData() InvoiceDeletion`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BillDelete200Response) GetDataOk() (*InvoiceDeletion, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BillDelete200Response) SetData(v InvoiceDeletion)`

SetData sets Data field to given value.

### HasData

`func (o *BillDelete200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *BillDelete200Response) GetMeta() BillCreate200ResponseMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *BillDelete200Response) GetMetaOk() (*BillCreate200ResponseMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *BillDelete200Response) SetMeta(v BillCreate200ResponseMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *BillDelete200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


