# GatewayCreate200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**Data** | Pointer to [**[]Gateway**](Gateway.md) |  | [optional] 

## Methods

### NewGatewayCreate200Response

`func NewGatewayCreate200Response() *GatewayCreate200Response`

NewGatewayCreate200Response instantiates a new GatewayCreate200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayCreate200ResponseWithDefaults

`func NewGatewayCreate200ResponseWithDefaults() *GatewayCreate200Response`

NewGatewayCreate200ResponseWithDefaults instantiates a new GatewayCreate200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *GatewayCreate200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GatewayCreate200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GatewayCreate200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GatewayCreate200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetData

`func (o *GatewayCreate200Response) GetData() []Gateway`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GatewayCreate200Response) GetDataOk() (*[]Gateway, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GatewayCreate200Response) SetData(v []Gateway)`

SetData sets Data field to given value.

### HasData

`func (o *GatewayCreate200Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


