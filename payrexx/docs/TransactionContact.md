# TransactionContact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [default to 0]
**Uuid** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** | Possible values are: mister, miss, diverse | [optional] 
**Firstname** | Pointer to **string** |  | [optional] 
**Lastname** | Pointer to **string** |  | [optional] 
**Company** | Pointer to **string** |  | [optional] 
**Street** | Pointer to **string** |  | [optional] 
**Zip** | Pointer to **string** |  | [optional] 
**Place** | Pointer to **string** |  | [optional] 
**Country** | Pointer to **string** |  | [optional] 
**CountryISO** | Pointer to **string** |  | [optional] 
**Phone** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**DateOfBirth** | Pointer to **string** |  | [optional] 
**DeliveryTitle** | Pointer to **string** | Possible values are: mister, miss, diverse | [optional] 
**DeliveryFirstname** | Pointer to **string** |  | [optional] 
**DeliveryLastname** | Pointer to **string** |  | [optional] 
**DeliveryCompany** | Pointer to **string** |  | [optional] 
**DeliveryStreet** | Pointer to **string** |  | [optional] 
**DeliveryZip** | Pointer to **string** |  | [optional] 
**DeliveryPlace** | Pointer to **string** |  | [optional] 
**DeliveryCountry** | Pointer to **string** |  | [optional] 
**DeliveryCountryISO** | Pointer to **string** |  | [optional] 

## Methods

### NewTransactionContact

`func NewTransactionContact() *TransactionContact`

NewTransactionContact instantiates a new TransactionContact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionContactWithDefaults

`func NewTransactionContactWithDefaults() *TransactionContact`

NewTransactionContactWithDefaults instantiates a new TransactionContact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TransactionContact) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TransactionContact) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TransactionContact) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *TransactionContact) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUuid

`func (o *TransactionContact) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *TransactionContact) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *TransactionContact) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *TransactionContact) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetTitle

`func (o *TransactionContact) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *TransactionContact) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *TransactionContact) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *TransactionContact) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetFirstname

`func (o *TransactionContact) GetFirstname() string`

GetFirstname returns the Firstname field if non-nil, zero value otherwise.

### GetFirstnameOk

`func (o *TransactionContact) GetFirstnameOk() (*string, bool)`

GetFirstnameOk returns a tuple with the Firstname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstname

`func (o *TransactionContact) SetFirstname(v string)`

SetFirstname sets Firstname field to given value.

### HasFirstname

`func (o *TransactionContact) HasFirstname() bool`

HasFirstname returns a boolean if a field has been set.

### GetLastname

`func (o *TransactionContact) GetLastname() string`

GetLastname returns the Lastname field if non-nil, zero value otherwise.

### GetLastnameOk

`func (o *TransactionContact) GetLastnameOk() (*string, bool)`

GetLastnameOk returns a tuple with the Lastname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastname

`func (o *TransactionContact) SetLastname(v string)`

SetLastname sets Lastname field to given value.

### HasLastname

`func (o *TransactionContact) HasLastname() bool`

HasLastname returns a boolean if a field has been set.

### GetCompany

`func (o *TransactionContact) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *TransactionContact) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *TransactionContact) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *TransactionContact) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetStreet

`func (o *TransactionContact) GetStreet() string`

GetStreet returns the Street field if non-nil, zero value otherwise.

### GetStreetOk

`func (o *TransactionContact) GetStreetOk() (*string, bool)`

GetStreetOk returns a tuple with the Street field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet

`func (o *TransactionContact) SetStreet(v string)`

SetStreet sets Street field to given value.

### HasStreet

`func (o *TransactionContact) HasStreet() bool`

HasStreet returns a boolean if a field has been set.

### GetZip

`func (o *TransactionContact) GetZip() string`

GetZip returns the Zip field if non-nil, zero value otherwise.

### GetZipOk

`func (o *TransactionContact) GetZipOk() (*string, bool)`

GetZipOk returns a tuple with the Zip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZip

`func (o *TransactionContact) SetZip(v string)`

SetZip sets Zip field to given value.

### HasZip

`func (o *TransactionContact) HasZip() bool`

HasZip returns a boolean if a field has been set.

### GetPlace

`func (o *TransactionContact) GetPlace() string`

GetPlace returns the Place field if non-nil, zero value otherwise.

### GetPlaceOk

`func (o *TransactionContact) GetPlaceOk() (*string, bool)`

GetPlaceOk returns a tuple with the Place field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlace

`func (o *TransactionContact) SetPlace(v string)`

SetPlace sets Place field to given value.

### HasPlace

`func (o *TransactionContact) HasPlace() bool`

HasPlace returns a boolean if a field has been set.

### GetCountry

`func (o *TransactionContact) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *TransactionContact) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *TransactionContact) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *TransactionContact) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCountryISO

`func (o *TransactionContact) GetCountryISO() string`

GetCountryISO returns the CountryISO field if non-nil, zero value otherwise.

### GetCountryISOOk

`func (o *TransactionContact) GetCountryISOOk() (*string, bool)`

GetCountryISOOk returns a tuple with the CountryISO field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryISO

`func (o *TransactionContact) SetCountryISO(v string)`

SetCountryISO sets CountryISO field to given value.

### HasCountryISO

`func (o *TransactionContact) HasCountryISO() bool`

HasCountryISO returns a boolean if a field has been set.

### GetPhone

`func (o *TransactionContact) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *TransactionContact) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *TransactionContact) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *TransactionContact) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetEmail

`func (o *TransactionContact) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *TransactionContact) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *TransactionContact) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *TransactionContact) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetDateOfBirth

`func (o *TransactionContact) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *TransactionContact) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *TransactionContact) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *TransactionContact) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### GetDeliveryTitle

`func (o *TransactionContact) GetDeliveryTitle() string`

GetDeliveryTitle returns the DeliveryTitle field if non-nil, zero value otherwise.

### GetDeliveryTitleOk

`func (o *TransactionContact) GetDeliveryTitleOk() (*string, bool)`

GetDeliveryTitleOk returns a tuple with the DeliveryTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryTitle

`func (o *TransactionContact) SetDeliveryTitle(v string)`

SetDeliveryTitle sets DeliveryTitle field to given value.

### HasDeliveryTitle

`func (o *TransactionContact) HasDeliveryTitle() bool`

HasDeliveryTitle returns a boolean if a field has been set.

### GetDeliveryFirstname

`func (o *TransactionContact) GetDeliveryFirstname() string`

GetDeliveryFirstname returns the DeliveryFirstname field if non-nil, zero value otherwise.

### GetDeliveryFirstnameOk

`func (o *TransactionContact) GetDeliveryFirstnameOk() (*string, bool)`

GetDeliveryFirstnameOk returns a tuple with the DeliveryFirstname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryFirstname

`func (o *TransactionContact) SetDeliveryFirstname(v string)`

SetDeliveryFirstname sets DeliveryFirstname field to given value.

### HasDeliveryFirstname

`func (o *TransactionContact) HasDeliveryFirstname() bool`

HasDeliveryFirstname returns a boolean if a field has been set.

### GetDeliveryLastname

`func (o *TransactionContact) GetDeliveryLastname() string`

GetDeliveryLastname returns the DeliveryLastname field if non-nil, zero value otherwise.

### GetDeliveryLastnameOk

`func (o *TransactionContact) GetDeliveryLastnameOk() (*string, bool)`

GetDeliveryLastnameOk returns a tuple with the DeliveryLastname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryLastname

`func (o *TransactionContact) SetDeliveryLastname(v string)`

SetDeliveryLastname sets DeliveryLastname field to given value.

### HasDeliveryLastname

`func (o *TransactionContact) HasDeliveryLastname() bool`

HasDeliveryLastname returns a boolean if a field has been set.

### GetDeliveryCompany

`func (o *TransactionContact) GetDeliveryCompany() string`

GetDeliveryCompany returns the DeliveryCompany field if non-nil, zero value otherwise.

### GetDeliveryCompanyOk

`func (o *TransactionContact) GetDeliveryCompanyOk() (*string, bool)`

GetDeliveryCompanyOk returns a tuple with the DeliveryCompany field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryCompany

`func (o *TransactionContact) SetDeliveryCompany(v string)`

SetDeliveryCompany sets DeliveryCompany field to given value.

### HasDeliveryCompany

`func (o *TransactionContact) HasDeliveryCompany() bool`

HasDeliveryCompany returns a boolean if a field has been set.

### GetDeliveryStreet

`func (o *TransactionContact) GetDeliveryStreet() string`

GetDeliveryStreet returns the DeliveryStreet field if non-nil, zero value otherwise.

### GetDeliveryStreetOk

`func (o *TransactionContact) GetDeliveryStreetOk() (*string, bool)`

GetDeliveryStreetOk returns a tuple with the DeliveryStreet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryStreet

`func (o *TransactionContact) SetDeliveryStreet(v string)`

SetDeliveryStreet sets DeliveryStreet field to given value.

### HasDeliveryStreet

`func (o *TransactionContact) HasDeliveryStreet() bool`

HasDeliveryStreet returns a boolean if a field has been set.

### GetDeliveryZip

`func (o *TransactionContact) GetDeliveryZip() string`

GetDeliveryZip returns the DeliveryZip field if non-nil, zero value otherwise.

### GetDeliveryZipOk

`func (o *TransactionContact) GetDeliveryZipOk() (*string, bool)`

GetDeliveryZipOk returns a tuple with the DeliveryZip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryZip

`func (o *TransactionContact) SetDeliveryZip(v string)`

SetDeliveryZip sets DeliveryZip field to given value.

### HasDeliveryZip

`func (o *TransactionContact) HasDeliveryZip() bool`

HasDeliveryZip returns a boolean if a field has been set.

### GetDeliveryPlace

`func (o *TransactionContact) GetDeliveryPlace() string`

GetDeliveryPlace returns the DeliveryPlace field if non-nil, zero value otherwise.

### GetDeliveryPlaceOk

`func (o *TransactionContact) GetDeliveryPlaceOk() (*string, bool)`

GetDeliveryPlaceOk returns a tuple with the DeliveryPlace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryPlace

`func (o *TransactionContact) SetDeliveryPlace(v string)`

SetDeliveryPlace sets DeliveryPlace field to given value.

### HasDeliveryPlace

`func (o *TransactionContact) HasDeliveryPlace() bool`

HasDeliveryPlace returns a boolean if a field has been set.

### GetDeliveryCountry

`func (o *TransactionContact) GetDeliveryCountry() string`

GetDeliveryCountry returns the DeliveryCountry field if non-nil, zero value otherwise.

### GetDeliveryCountryOk

`func (o *TransactionContact) GetDeliveryCountryOk() (*string, bool)`

GetDeliveryCountryOk returns a tuple with the DeliveryCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryCountry

`func (o *TransactionContact) SetDeliveryCountry(v string)`

SetDeliveryCountry sets DeliveryCountry field to given value.

### HasDeliveryCountry

`func (o *TransactionContact) HasDeliveryCountry() bool`

HasDeliveryCountry returns a boolean if a field has been set.

### GetDeliveryCountryISO

`func (o *TransactionContact) GetDeliveryCountryISO() string`

GetDeliveryCountryISO returns the DeliveryCountryISO field if non-nil, zero value otherwise.

### GetDeliveryCountryISOOk

`func (o *TransactionContact) GetDeliveryCountryISOOk() (*string, bool)`

GetDeliveryCountryISOOk returns a tuple with the DeliveryCountryISO field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryCountryISO

`func (o *TransactionContact) SetDeliveryCountryISO(v string)`

SetDeliveryCountryISO sets DeliveryCountryISO field to given value.

### HasDeliveryCountryISO

`func (o *TransactionContact) HasDeliveryCountryISO() bool`

HasDeliveryCountryISO returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


