# PaymentProviderList200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**Data** | Pointer to [**[]PaymentProvider**](PaymentProvider.md) |  | [optional] 

## Methods

### NewPaymentProviderList200Response

`func NewPaymentProviderList200Response() *PaymentProviderList200Response`

NewPaymentProviderList200Response instantiates a new PaymentProviderList200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentProviderList200ResponseWithDefaults

`func NewPaymentProviderList200ResponseWithDefaults() *PaymentProviderList200Response`

NewPaymentProviderList200ResponseWithDefaults instantiates a new PaymentProviderList200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *PaymentProviderList200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaymentProviderList200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaymentProviderList200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PaymentProviderList200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetData

`func (o *PaymentProviderList200Response) GetData() []PaymentProvider`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PaymentProviderList200Response) GetDataOk() (*[]PaymentProvider, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PaymentProviderList200Response) SetData(v []PaymentProvider)`

SetData sets Data field to given value.

### HasData

`func (o *PaymentProviderList200Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


