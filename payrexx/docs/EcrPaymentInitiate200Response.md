# EcrPaymentInitiate200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** |  | 
**Data** | Pointer to [**EcrPayment**](EcrPayment.md) |  | [optional] 
**Message** | Pointer to **string** | Contains error messages | [optional] 

## Methods

### NewEcrPaymentInitiate200Response

`func NewEcrPaymentInitiate200Response(status string, ) *EcrPaymentInitiate200Response`

NewEcrPaymentInitiate200Response instantiates a new EcrPaymentInitiate200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEcrPaymentInitiate200ResponseWithDefaults

`func NewEcrPaymentInitiate200ResponseWithDefaults() *EcrPaymentInitiate200Response`

NewEcrPaymentInitiate200ResponseWithDefaults instantiates a new EcrPaymentInitiate200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *EcrPaymentInitiate200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EcrPaymentInitiate200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EcrPaymentInitiate200Response) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetData

`func (o *EcrPaymentInitiate200Response) GetData() EcrPayment`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *EcrPaymentInitiate200Response) GetDataOk() (*EcrPayment, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *EcrPaymentInitiate200Response) SetData(v EcrPayment)`

SetData sets Data field to given value.

### HasData

`func (o *EcrPaymentInitiate200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMessage

`func (o *EcrPaymentInitiate200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *EcrPaymentInitiate200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *EcrPaymentInitiate200Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *EcrPaymentInitiate200Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


