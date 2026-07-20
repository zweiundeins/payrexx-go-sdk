# EcrPaymentMethod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Group** | Pointer to **string** |  | [optional] 
**OrderBy** | Pointer to **string** |  | [optional] 
**PaymentOptions** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewEcrPaymentMethod

`func NewEcrPaymentMethod() *EcrPaymentMethod`

NewEcrPaymentMethod instantiates a new EcrPaymentMethod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEcrPaymentMethodWithDefaults

`func NewEcrPaymentMethodWithDefaults() *EcrPaymentMethod`

NewEcrPaymentMethodWithDefaults instantiates a new EcrPaymentMethod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *EcrPaymentMethod) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EcrPaymentMethod) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EcrPaymentMethod) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EcrPaymentMethod) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *EcrPaymentMethod) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *EcrPaymentMethod) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *EcrPaymentMethod) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *EcrPaymentMethod) HasType() bool`

HasType returns a boolean if a field has been set.

### GetGroup

`func (o *EcrPaymentMethod) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *EcrPaymentMethod) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *EcrPaymentMethod) SetGroup(v string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *EcrPaymentMethod) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetOrderBy

`func (o *EcrPaymentMethod) GetOrderBy() string`

GetOrderBy returns the OrderBy field if non-nil, zero value otherwise.

### GetOrderByOk

`func (o *EcrPaymentMethod) GetOrderByOk() (*string, bool)`

GetOrderByOk returns a tuple with the OrderBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBy

`func (o *EcrPaymentMethod) SetOrderBy(v string)`

SetOrderBy sets OrderBy field to given value.

### HasOrderBy

`func (o *EcrPaymentMethod) HasOrderBy() bool`

HasOrderBy returns a boolean if a field has been set.

### GetPaymentOptions

`func (o *EcrPaymentMethod) GetPaymentOptions() []map[string]interface{}`

GetPaymentOptions returns the PaymentOptions field if non-nil, zero value otherwise.

### GetPaymentOptionsOk

`func (o *EcrPaymentMethod) GetPaymentOptionsOk() (*[]map[string]interface{}, bool)`

GetPaymentOptionsOk returns a tuple with the PaymentOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentOptions

`func (o *EcrPaymentMethod) SetPaymentOptions(v []map[string]interface{})`

SetPaymentOptions sets PaymentOptions field to given value.

### HasPaymentOptions

`func (o *EcrPaymentMethod) HasPaymentOptions() bool`

HasPaymentOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


