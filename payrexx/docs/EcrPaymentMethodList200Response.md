# EcrPaymentMethodList200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** |  | 
**Data** | Pointer to [**[]EcrPaymentMethod**](EcrPaymentMethod.md) |  | [optional] 

## Methods

### NewEcrPaymentMethodList200Response

`func NewEcrPaymentMethodList200Response(status string, ) *EcrPaymentMethodList200Response`

NewEcrPaymentMethodList200Response instantiates a new EcrPaymentMethodList200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEcrPaymentMethodList200ResponseWithDefaults

`func NewEcrPaymentMethodList200ResponseWithDefaults() *EcrPaymentMethodList200Response`

NewEcrPaymentMethodList200ResponseWithDefaults instantiates a new EcrPaymentMethodList200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *EcrPaymentMethodList200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EcrPaymentMethodList200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EcrPaymentMethodList200Response) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetData

`func (o *EcrPaymentMethodList200Response) GetData() []EcrPaymentMethod`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *EcrPaymentMethodList200Response) GetDataOk() (*[]EcrPaymentMethod, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *EcrPaymentMethodList200Response) SetData(v []EcrPaymentMethod)`

SetData sets Data field to given value.

### HasData

`func (o *EcrPaymentMethodList200Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


