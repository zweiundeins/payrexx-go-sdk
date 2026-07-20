# SubscriptionContact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [default to 0]
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

## Methods

### NewSubscriptionContact

`func NewSubscriptionContact() *SubscriptionContact`

NewSubscriptionContact instantiates a new SubscriptionContact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionContactWithDefaults

`func NewSubscriptionContactWithDefaults() *SubscriptionContact`

NewSubscriptionContactWithDefaults instantiates a new SubscriptionContact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SubscriptionContact) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SubscriptionContact) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SubscriptionContact) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *SubscriptionContact) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTitle

`func (o *SubscriptionContact) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *SubscriptionContact) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *SubscriptionContact) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *SubscriptionContact) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetFirstname

`func (o *SubscriptionContact) GetFirstname() string`

GetFirstname returns the Firstname field if non-nil, zero value otherwise.

### GetFirstnameOk

`func (o *SubscriptionContact) GetFirstnameOk() (*string, bool)`

GetFirstnameOk returns a tuple with the Firstname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstname

`func (o *SubscriptionContact) SetFirstname(v string)`

SetFirstname sets Firstname field to given value.

### HasFirstname

`func (o *SubscriptionContact) HasFirstname() bool`

HasFirstname returns a boolean if a field has been set.

### GetLastname

`func (o *SubscriptionContact) GetLastname() string`

GetLastname returns the Lastname field if non-nil, zero value otherwise.

### GetLastnameOk

`func (o *SubscriptionContact) GetLastnameOk() (*string, bool)`

GetLastnameOk returns a tuple with the Lastname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastname

`func (o *SubscriptionContact) SetLastname(v string)`

SetLastname sets Lastname field to given value.

### HasLastname

`func (o *SubscriptionContact) HasLastname() bool`

HasLastname returns a boolean if a field has been set.

### GetCompany

`func (o *SubscriptionContact) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *SubscriptionContact) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *SubscriptionContact) SetCompany(v string)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *SubscriptionContact) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetStreet

`func (o *SubscriptionContact) GetStreet() string`

GetStreet returns the Street field if non-nil, zero value otherwise.

### GetStreetOk

`func (o *SubscriptionContact) GetStreetOk() (*string, bool)`

GetStreetOk returns a tuple with the Street field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet

`func (o *SubscriptionContact) SetStreet(v string)`

SetStreet sets Street field to given value.

### HasStreet

`func (o *SubscriptionContact) HasStreet() bool`

HasStreet returns a boolean if a field has been set.

### GetZip

`func (o *SubscriptionContact) GetZip() string`

GetZip returns the Zip field if non-nil, zero value otherwise.

### GetZipOk

`func (o *SubscriptionContact) GetZipOk() (*string, bool)`

GetZipOk returns a tuple with the Zip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZip

`func (o *SubscriptionContact) SetZip(v string)`

SetZip sets Zip field to given value.

### HasZip

`func (o *SubscriptionContact) HasZip() bool`

HasZip returns a boolean if a field has been set.

### GetPlace

`func (o *SubscriptionContact) GetPlace() string`

GetPlace returns the Place field if non-nil, zero value otherwise.

### GetPlaceOk

`func (o *SubscriptionContact) GetPlaceOk() (*string, bool)`

GetPlaceOk returns a tuple with the Place field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlace

`func (o *SubscriptionContact) SetPlace(v string)`

SetPlace sets Place field to given value.

### HasPlace

`func (o *SubscriptionContact) HasPlace() bool`

HasPlace returns a boolean if a field has been set.

### GetCountry

`func (o *SubscriptionContact) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *SubscriptionContact) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *SubscriptionContact) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *SubscriptionContact) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCountryISO

`func (o *SubscriptionContact) GetCountryISO() string`

GetCountryISO returns the CountryISO field if non-nil, zero value otherwise.

### GetCountryISOOk

`func (o *SubscriptionContact) GetCountryISOOk() (*string, bool)`

GetCountryISOOk returns a tuple with the CountryISO field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryISO

`func (o *SubscriptionContact) SetCountryISO(v string)`

SetCountryISO sets CountryISO field to given value.

### HasCountryISO

`func (o *SubscriptionContact) HasCountryISO() bool`

HasCountryISO returns a boolean if a field has been set.

### GetPhone

`func (o *SubscriptionContact) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *SubscriptionContact) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *SubscriptionContact) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *SubscriptionContact) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetEmail

`func (o *SubscriptionContact) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *SubscriptionContact) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *SubscriptionContact) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *SubscriptionContact) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetDateOfBirth

`func (o *SubscriptionContact) GetDateOfBirth() string`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *SubscriptionContact) GetDateOfBirthOk() (*string, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *SubscriptionContact) SetDateOfBirth(v string)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *SubscriptionContact) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


