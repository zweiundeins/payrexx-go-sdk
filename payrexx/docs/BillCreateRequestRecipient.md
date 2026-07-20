# BillCreateRequestRecipient

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Company** | Pointer to **string** | **Required:** at least one of &#x60;company&#x60;, &#x60;firstName&#x60;, or &#x60;lastName&#x60; must be **provided** | [optional] 
**FirstName** | Pointer to **string** | **Required:** at least one of &#x60;company&#x60;, &#x60;firstName&#x60;, or &#x60;lastName&#x60; must be **provided** | [optional] 
**LastName** | Pointer to **string** | **Required:** at least one of &#x60;company&#x60;, &#x60;firstName&#x60;, or &#x60;lastName&#x60; must be **provided** | [optional] 
**Address** | Pointer to **string** |  | [optional] 
**Zip** | Pointer to **string** |  | [optional] 
**City** | Pointer to **string** |  | [optional] 
**Country** | Pointer to **string** | Country code (ISO 3166-1 alpha-2) | [optional] 
**Email** | **string** |  | 

## Methods

### NewBillCreateRequestRecipient

`func NewBillCreateRequestRecipient(email string, ) *BillCreateRequestRecipient`

NewBillCreateRequestRecipient instantiates a new BillCreateRequestRecipient object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillCreateRequestRecipientWithDefaults

`func NewBillCreateRequestRecipientWithDefaults() *BillCreateRequestRecipient`

NewBillCreateRequestRecipientWithDefaults instantiates a new BillCreateRequestRecipient object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompany

`func (o *BillCreateRequestRecipient) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *BillCreateRequestRecipient) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *BillCreateRequestRecipient) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *BillCreateRequestRecipient) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetFirstName

`func (o *BillCreateRequestRecipient) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *BillCreateRequestRecipient) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *BillCreateRequestRecipient) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *BillCreateRequestRecipient) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *BillCreateRequestRecipient) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *BillCreateRequestRecipient) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *BillCreateRequestRecipient) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *BillCreateRequestRecipient) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetAddress

`func (o *BillCreateRequestRecipient) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *BillCreateRequestRecipient) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *BillCreateRequestRecipient) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *BillCreateRequestRecipient) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetZip

`func (o *BillCreateRequestRecipient) GetZip() string`

GetZip returns the Zip field if non-nil, zero value otherwise.

### GetZipOk

`func (o *BillCreateRequestRecipient) GetZipOk() (*string, bool)`

GetZipOk returns a tuple with the Zip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZip

`func (o *BillCreateRequestRecipient) SetZip(v string)`

SetZip sets Zip field to given value.

### HasZip

`func (o *BillCreateRequestRecipient) HasZip() bool`

HasZip returns a boolean if a field has been set.

### GetCity

`func (o *BillCreateRequestRecipient) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *BillCreateRequestRecipient) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *BillCreateRequestRecipient) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *BillCreateRequestRecipient) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetCountry

`func (o *BillCreateRequestRecipient) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *BillCreateRequestRecipient) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *BillCreateRequestRecipient) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *BillCreateRequestRecipient) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetEmail

`func (o *BillCreateRequestRecipient) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *BillCreateRequestRecipient) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *BillCreateRequestRecipient) SetEmail(v string)`

SetEmail sets Email field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


