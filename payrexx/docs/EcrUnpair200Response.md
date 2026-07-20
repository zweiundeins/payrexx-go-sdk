# EcrUnpair200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** |  | 
**Data** | Pointer to **[]string** | empty after success | [optional] 
**Message** | Pointer to **string** | Contains error messages | [optional] 

## Methods

### NewEcrUnpair200Response

`func NewEcrUnpair200Response(status string, ) *EcrUnpair200Response`

NewEcrUnpair200Response instantiates a new EcrUnpair200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEcrUnpair200ResponseWithDefaults

`func NewEcrUnpair200ResponseWithDefaults() *EcrUnpair200Response`

NewEcrUnpair200ResponseWithDefaults instantiates a new EcrUnpair200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *EcrUnpair200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *EcrUnpair200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *EcrUnpair200Response) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetData

`func (o *EcrUnpair200Response) GetData() []string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *EcrUnpair200Response) GetDataOk() (*[]string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *EcrUnpair200Response) SetData(v []string)`

SetData sets Data field to given value.

### HasData

`func (o *EcrUnpair200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMessage

`func (o *EcrUnpair200Response) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *EcrUnpair200Response) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *EcrUnpair200Response) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *EcrUnpair200Response) HasMessage() bool`

HasMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


