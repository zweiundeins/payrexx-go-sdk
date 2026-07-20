# GatewayCreateRequestFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | Pointer to [**GatewayCreateRequestFieldsTitle**](GatewayCreateRequestFieldsTitle.md) |  | [optional] 
**Forename** | Pointer to [**GatewayCreateRequestFieldsTitle**](GatewayCreateRequestFieldsTitle.md) |  | [optional] 
**Surname** | Pointer to [**GatewayCreateRequestFieldsTitle**](GatewayCreateRequestFieldsTitle.md) |  | [optional] 
**Company** | Pointer to [**GatewayCreateRequestFieldsTitle**](GatewayCreateRequestFieldsTitle.md) |  | [optional] 
**Street** | Pointer to [**GatewayCreateRequestFieldsTitle**](GatewayCreateRequestFieldsTitle.md) |  | [optional] 
**Postcode** | Pointer to [**GatewayCreateRequestFieldsTitle**](GatewayCreateRequestFieldsTitle.md) |  | [optional] 
**Place** | Pointer to [**GatewayCreateRequestFieldsTitle**](GatewayCreateRequestFieldsTitle.md) |  | [optional] 
**Country** | Pointer to [**GatewayCreateRequestFieldsTitle**](GatewayCreateRequestFieldsTitle.md) |  | [optional] 
**DeliveryTitle** | Pointer to [**GatewayCreateRequestFieldsTitle**](GatewayCreateRequestFieldsTitle.md) |  | [optional] 
**DeliveryForename** | Pointer to [**GatewayCreateRequestFieldsTitle**](GatewayCreateRequestFieldsTitle.md) |  | [optional] 
**DeliverySurname** | Pointer to [**GatewayCreateRequestFieldsTitle**](GatewayCreateRequestFieldsTitle.md) |  | [optional] 
**DeliveryCompany** | Pointer to [**GatewayCreateRequestFieldsTitle**](GatewayCreateRequestFieldsTitle.md) |  | [optional] 
**DeliveryStreet** | Pointer to [**GatewayCreateRequestFieldsTitle**](GatewayCreateRequestFieldsTitle.md) |  | [optional] 
**DeliveryPostcode** | Pointer to [**GatewayCreateRequestFieldsTitle**](GatewayCreateRequestFieldsTitle.md) |  | [optional] 
**DeliveryPlace** | Pointer to [**GatewayCreateRequestFieldsTitle**](GatewayCreateRequestFieldsTitle.md) |  | [optional] 
**DeliveryCountry** | Pointer to [**GatewayCreateRequestFieldsTitle**](GatewayCreateRequestFieldsTitle.md) |  | [optional] 
**Phone** | Pointer to [**GatewayCreateRequestFieldsTitle**](GatewayCreateRequestFieldsTitle.md) |  | [optional] 
**Email** | Pointer to [**GatewayCreateRequestFieldsTitle**](GatewayCreateRequestFieldsTitle.md) |  | [optional] 
**DateOfBirth** | Pointer to [**GatewayCreateRequestFieldsTitle**](GatewayCreateRequestFieldsTitle.md) |  | [optional] 
**Terms** | Pointer to **string** | The terms must be accepted if this field is defined in the request | [optional] 
**PrivacyPolicy** | Pointer to **string** | The privacy policy must be accepted if this field is defined in the request | [optional] 
**CustomField1** | Pointer to [**GatewayCreateRequestFieldsCustomField1**](GatewayCreateRequestFieldsCustomField1.md) |  | [optional] 
**CustomField2** | Pointer to [**GatewayCreateRequestFieldsCustomField2**](GatewayCreateRequestFieldsCustomField2.md) |  | [optional] 
**CustomField3** | Pointer to [**GatewayCreateRequestFieldsCustomField2**](GatewayCreateRequestFieldsCustomField2.md) |  | [optional] 
**CustomField4** | Pointer to [**GatewayCreateRequestFieldsCustomField2**](GatewayCreateRequestFieldsCustomField2.md) |  | [optional] 
**CustomField5** | Pointer to [**GatewayCreateRequestFieldsCustomField2**](GatewayCreateRequestFieldsCustomField2.md) |  | [optional] 

## Methods

### NewGatewayCreateRequestFields

`func NewGatewayCreateRequestFields() *GatewayCreateRequestFields`

NewGatewayCreateRequestFields instantiates a new GatewayCreateRequestFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayCreateRequestFieldsWithDefaults

`func NewGatewayCreateRequestFieldsWithDefaults() *GatewayCreateRequestFields`

NewGatewayCreateRequestFieldsWithDefaults instantiates a new GatewayCreateRequestFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *GatewayCreateRequestFields) GetTitle() GatewayCreateRequestFieldsTitle`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *GatewayCreateRequestFields) GetTitleOk() (*GatewayCreateRequestFieldsTitle, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *GatewayCreateRequestFields) SetTitle(v GatewayCreateRequestFieldsTitle)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *GatewayCreateRequestFields) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetForename

`func (o *GatewayCreateRequestFields) GetForename() GatewayCreateRequestFieldsTitle`

GetForename returns the Forename field if non-nil, zero value otherwise.

### GetForenameOk

`func (o *GatewayCreateRequestFields) GetForenameOk() (*GatewayCreateRequestFieldsTitle, bool)`

GetForenameOk returns a tuple with the Forename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForename

`func (o *GatewayCreateRequestFields) SetForename(v GatewayCreateRequestFieldsTitle)`

SetForename sets Forename field to given value.

### HasForename

`func (o *GatewayCreateRequestFields) HasForename() bool`

HasForename returns a boolean if a field has been set.

### GetSurname

`func (o *GatewayCreateRequestFields) GetSurname() GatewayCreateRequestFieldsTitle`

GetSurname returns the Surname field if non-nil, zero value otherwise.

### GetSurnameOk

`func (o *GatewayCreateRequestFields) GetSurnameOk() (*GatewayCreateRequestFieldsTitle, bool)`

GetSurnameOk returns a tuple with the Surname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurname

`func (o *GatewayCreateRequestFields) SetSurname(v GatewayCreateRequestFieldsTitle)`

SetSurname sets Surname field to given value.

### HasSurname

`func (o *GatewayCreateRequestFields) HasSurname() bool`

HasSurname returns a boolean if a field has been set.

### GetCompany

`func (o *GatewayCreateRequestFields) GetCompany() GatewayCreateRequestFieldsTitle`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *GatewayCreateRequestFields) GetCompanyOk() (*GatewayCreateRequestFieldsTitle, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *GatewayCreateRequestFields) SetCompany(v GatewayCreateRequestFieldsTitle)`

SetCompany sets Company field to given value.

### HasCompany

`func (o *GatewayCreateRequestFields) HasCompany() bool`

HasCompany returns a boolean if a field has been set.

### GetStreet

`func (o *GatewayCreateRequestFields) GetStreet() GatewayCreateRequestFieldsTitle`

GetStreet returns the Street field if non-nil, zero value otherwise.

### GetStreetOk

`func (o *GatewayCreateRequestFields) GetStreetOk() (*GatewayCreateRequestFieldsTitle, bool)`

GetStreetOk returns a tuple with the Street field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet

`func (o *GatewayCreateRequestFields) SetStreet(v GatewayCreateRequestFieldsTitle)`

SetStreet sets Street field to given value.

### HasStreet

`func (o *GatewayCreateRequestFields) HasStreet() bool`

HasStreet returns a boolean if a field has been set.

### GetPostcode

`func (o *GatewayCreateRequestFields) GetPostcode() GatewayCreateRequestFieldsTitle`

GetPostcode returns the Postcode field if non-nil, zero value otherwise.

### GetPostcodeOk

`func (o *GatewayCreateRequestFields) GetPostcodeOk() (*GatewayCreateRequestFieldsTitle, bool)`

GetPostcodeOk returns a tuple with the Postcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostcode

`func (o *GatewayCreateRequestFields) SetPostcode(v GatewayCreateRequestFieldsTitle)`

SetPostcode sets Postcode field to given value.

### HasPostcode

`func (o *GatewayCreateRequestFields) HasPostcode() bool`

HasPostcode returns a boolean if a field has been set.

### GetPlace

`func (o *GatewayCreateRequestFields) GetPlace() GatewayCreateRequestFieldsTitle`

GetPlace returns the Place field if non-nil, zero value otherwise.

### GetPlaceOk

`func (o *GatewayCreateRequestFields) GetPlaceOk() (*GatewayCreateRequestFieldsTitle, bool)`

GetPlaceOk returns a tuple with the Place field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlace

`func (o *GatewayCreateRequestFields) SetPlace(v GatewayCreateRequestFieldsTitle)`

SetPlace sets Place field to given value.

### HasPlace

`func (o *GatewayCreateRequestFields) HasPlace() bool`

HasPlace returns a boolean if a field has been set.

### GetCountry

`func (o *GatewayCreateRequestFields) GetCountry() GatewayCreateRequestFieldsTitle`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *GatewayCreateRequestFields) GetCountryOk() (*GatewayCreateRequestFieldsTitle, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *GatewayCreateRequestFields) SetCountry(v GatewayCreateRequestFieldsTitle)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *GatewayCreateRequestFields) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetDeliveryTitle

`func (o *GatewayCreateRequestFields) GetDeliveryTitle() GatewayCreateRequestFieldsTitle`

GetDeliveryTitle returns the DeliveryTitle field if non-nil, zero value otherwise.

### GetDeliveryTitleOk

`func (o *GatewayCreateRequestFields) GetDeliveryTitleOk() (*GatewayCreateRequestFieldsTitle, bool)`

GetDeliveryTitleOk returns a tuple with the DeliveryTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryTitle

`func (o *GatewayCreateRequestFields) SetDeliveryTitle(v GatewayCreateRequestFieldsTitle)`

SetDeliveryTitle sets DeliveryTitle field to given value.

### HasDeliveryTitle

`func (o *GatewayCreateRequestFields) HasDeliveryTitle() bool`

HasDeliveryTitle returns a boolean if a field has been set.

### GetDeliveryForename

`func (o *GatewayCreateRequestFields) GetDeliveryForename() GatewayCreateRequestFieldsTitle`

GetDeliveryForename returns the DeliveryForename field if non-nil, zero value otherwise.

### GetDeliveryForenameOk

`func (o *GatewayCreateRequestFields) GetDeliveryForenameOk() (*GatewayCreateRequestFieldsTitle, bool)`

GetDeliveryForenameOk returns a tuple with the DeliveryForename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryForename

`func (o *GatewayCreateRequestFields) SetDeliveryForename(v GatewayCreateRequestFieldsTitle)`

SetDeliveryForename sets DeliveryForename field to given value.

### HasDeliveryForename

`func (o *GatewayCreateRequestFields) HasDeliveryForename() bool`

HasDeliveryForename returns a boolean if a field has been set.

### GetDeliverySurname

`func (o *GatewayCreateRequestFields) GetDeliverySurname() GatewayCreateRequestFieldsTitle`

GetDeliverySurname returns the DeliverySurname field if non-nil, zero value otherwise.

### GetDeliverySurnameOk

`func (o *GatewayCreateRequestFields) GetDeliverySurnameOk() (*GatewayCreateRequestFieldsTitle, bool)`

GetDeliverySurnameOk returns a tuple with the DeliverySurname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliverySurname

`func (o *GatewayCreateRequestFields) SetDeliverySurname(v GatewayCreateRequestFieldsTitle)`

SetDeliverySurname sets DeliverySurname field to given value.

### HasDeliverySurname

`func (o *GatewayCreateRequestFields) HasDeliverySurname() bool`

HasDeliverySurname returns a boolean if a field has been set.

### GetDeliveryCompany

`func (o *GatewayCreateRequestFields) GetDeliveryCompany() GatewayCreateRequestFieldsTitle`

GetDeliveryCompany returns the DeliveryCompany field if non-nil, zero value otherwise.

### GetDeliveryCompanyOk

`func (o *GatewayCreateRequestFields) GetDeliveryCompanyOk() (*GatewayCreateRequestFieldsTitle, bool)`

GetDeliveryCompanyOk returns a tuple with the DeliveryCompany field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryCompany

`func (o *GatewayCreateRequestFields) SetDeliveryCompany(v GatewayCreateRequestFieldsTitle)`

SetDeliveryCompany sets DeliveryCompany field to given value.

### HasDeliveryCompany

`func (o *GatewayCreateRequestFields) HasDeliveryCompany() bool`

HasDeliveryCompany returns a boolean if a field has been set.

### GetDeliveryStreet

`func (o *GatewayCreateRequestFields) GetDeliveryStreet() GatewayCreateRequestFieldsTitle`

GetDeliveryStreet returns the DeliveryStreet field if non-nil, zero value otherwise.

### GetDeliveryStreetOk

`func (o *GatewayCreateRequestFields) GetDeliveryStreetOk() (*GatewayCreateRequestFieldsTitle, bool)`

GetDeliveryStreetOk returns a tuple with the DeliveryStreet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryStreet

`func (o *GatewayCreateRequestFields) SetDeliveryStreet(v GatewayCreateRequestFieldsTitle)`

SetDeliveryStreet sets DeliveryStreet field to given value.

### HasDeliveryStreet

`func (o *GatewayCreateRequestFields) HasDeliveryStreet() bool`

HasDeliveryStreet returns a boolean if a field has been set.

### GetDeliveryPostcode

`func (o *GatewayCreateRequestFields) GetDeliveryPostcode() GatewayCreateRequestFieldsTitle`

GetDeliveryPostcode returns the DeliveryPostcode field if non-nil, zero value otherwise.

### GetDeliveryPostcodeOk

`func (o *GatewayCreateRequestFields) GetDeliveryPostcodeOk() (*GatewayCreateRequestFieldsTitle, bool)`

GetDeliveryPostcodeOk returns a tuple with the DeliveryPostcode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryPostcode

`func (o *GatewayCreateRequestFields) SetDeliveryPostcode(v GatewayCreateRequestFieldsTitle)`

SetDeliveryPostcode sets DeliveryPostcode field to given value.

### HasDeliveryPostcode

`func (o *GatewayCreateRequestFields) HasDeliveryPostcode() bool`

HasDeliveryPostcode returns a boolean if a field has been set.

### GetDeliveryPlace

`func (o *GatewayCreateRequestFields) GetDeliveryPlace() GatewayCreateRequestFieldsTitle`

GetDeliveryPlace returns the DeliveryPlace field if non-nil, zero value otherwise.

### GetDeliveryPlaceOk

`func (o *GatewayCreateRequestFields) GetDeliveryPlaceOk() (*GatewayCreateRequestFieldsTitle, bool)`

GetDeliveryPlaceOk returns a tuple with the DeliveryPlace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryPlace

`func (o *GatewayCreateRequestFields) SetDeliveryPlace(v GatewayCreateRequestFieldsTitle)`

SetDeliveryPlace sets DeliveryPlace field to given value.

### HasDeliveryPlace

`func (o *GatewayCreateRequestFields) HasDeliveryPlace() bool`

HasDeliveryPlace returns a boolean if a field has been set.

### GetDeliveryCountry

`func (o *GatewayCreateRequestFields) GetDeliveryCountry() GatewayCreateRequestFieldsTitle`

GetDeliveryCountry returns the DeliveryCountry field if non-nil, zero value otherwise.

### GetDeliveryCountryOk

`func (o *GatewayCreateRequestFields) GetDeliveryCountryOk() (*GatewayCreateRequestFieldsTitle, bool)`

GetDeliveryCountryOk returns a tuple with the DeliveryCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryCountry

`func (o *GatewayCreateRequestFields) SetDeliveryCountry(v GatewayCreateRequestFieldsTitle)`

SetDeliveryCountry sets DeliveryCountry field to given value.

### HasDeliveryCountry

`func (o *GatewayCreateRequestFields) HasDeliveryCountry() bool`

HasDeliveryCountry returns a boolean if a field has been set.

### GetPhone

`func (o *GatewayCreateRequestFields) GetPhone() GatewayCreateRequestFieldsTitle`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *GatewayCreateRequestFields) GetPhoneOk() (*GatewayCreateRequestFieldsTitle, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *GatewayCreateRequestFields) SetPhone(v GatewayCreateRequestFieldsTitle)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *GatewayCreateRequestFields) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetEmail

`func (o *GatewayCreateRequestFields) GetEmail() GatewayCreateRequestFieldsTitle`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *GatewayCreateRequestFields) GetEmailOk() (*GatewayCreateRequestFieldsTitle, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *GatewayCreateRequestFields) SetEmail(v GatewayCreateRequestFieldsTitle)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *GatewayCreateRequestFields) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetDateOfBirth

`func (o *GatewayCreateRequestFields) GetDateOfBirth() GatewayCreateRequestFieldsTitle`

GetDateOfBirth returns the DateOfBirth field if non-nil, zero value otherwise.

### GetDateOfBirthOk

`func (o *GatewayCreateRequestFields) GetDateOfBirthOk() (*GatewayCreateRequestFieldsTitle, bool)`

GetDateOfBirthOk returns a tuple with the DateOfBirth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfBirth

`func (o *GatewayCreateRequestFields) SetDateOfBirth(v GatewayCreateRequestFieldsTitle)`

SetDateOfBirth sets DateOfBirth field to given value.

### HasDateOfBirth

`func (o *GatewayCreateRequestFields) HasDateOfBirth() bool`

HasDateOfBirth returns a boolean if a field has been set.

### GetTerms

`func (o *GatewayCreateRequestFields) GetTerms() string`

GetTerms returns the Terms field if non-nil, zero value otherwise.

### GetTermsOk

`func (o *GatewayCreateRequestFields) GetTermsOk() (*string, bool)`

GetTermsOk returns a tuple with the Terms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerms

`func (o *GatewayCreateRequestFields) SetTerms(v string)`

SetTerms sets Terms field to given value.

### HasTerms

`func (o *GatewayCreateRequestFields) HasTerms() bool`

HasTerms returns a boolean if a field has been set.

### GetPrivacyPolicy

`func (o *GatewayCreateRequestFields) GetPrivacyPolicy() string`

GetPrivacyPolicy returns the PrivacyPolicy field if non-nil, zero value otherwise.

### GetPrivacyPolicyOk

`func (o *GatewayCreateRequestFields) GetPrivacyPolicyOk() (*string, bool)`

GetPrivacyPolicyOk returns a tuple with the PrivacyPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivacyPolicy

`func (o *GatewayCreateRequestFields) SetPrivacyPolicy(v string)`

SetPrivacyPolicy sets PrivacyPolicy field to given value.

### HasPrivacyPolicy

`func (o *GatewayCreateRequestFields) HasPrivacyPolicy() bool`

HasPrivacyPolicy returns a boolean if a field has been set.

### GetCustomField1

`func (o *GatewayCreateRequestFields) GetCustomField1() GatewayCreateRequestFieldsCustomField1`

GetCustomField1 returns the CustomField1 field if non-nil, zero value otherwise.

### GetCustomField1Ok

`func (o *GatewayCreateRequestFields) GetCustomField1Ok() (*GatewayCreateRequestFieldsCustomField1, bool)`

GetCustomField1Ok returns a tuple with the CustomField1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomField1

`func (o *GatewayCreateRequestFields) SetCustomField1(v GatewayCreateRequestFieldsCustomField1)`

SetCustomField1 sets CustomField1 field to given value.

### HasCustomField1

`func (o *GatewayCreateRequestFields) HasCustomField1() bool`

HasCustomField1 returns a boolean if a field has been set.

### GetCustomField2

`func (o *GatewayCreateRequestFields) GetCustomField2() GatewayCreateRequestFieldsCustomField2`

GetCustomField2 returns the CustomField2 field if non-nil, zero value otherwise.

### GetCustomField2Ok

`func (o *GatewayCreateRequestFields) GetCustomField2Ok() (*GatewayCreateRequestFieldsCustomField2, bool)`

GetCustomField2Ok returns a tuple with the CustomField2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomField2

`func (o *GatewayCreateRequestFields) SetCustomField2(v GatewayCreateRequestFieldsCustomField2)`

SetCustomField2 sets CustomField2 field to given value.

### HasCustomField2

`func (o *GatewayCreateRequestFields) HasCustomField2() bool`

HasCustomField2 returns a boolean if a field has been set.

### GetCustomField3

`func (o *GatewayCreateRequestFields) GetCustomField3() GatewayCreateRequestFieldsCustomField2`

GetCustomField3 returns the CustomField3 field if non-nil, zero value otherwise.

### GetCustomField3Ok

`func (o *GatewayCreateRequestFields) GetCustomField3Ok() (*GatewayCreateRequestFieldsCustomField2, bool)`

GetCustomField3Ok returns a tuple with the CustomField3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomField3

`func (o *GatewayCreateRequestFields) SetCustomField3(v GatewayCreateRequestFieldsCustomField2)`

SetCustomField3 sets CustomField3 field to given value.

### HasCustomField3

`func (o *GatewayCreateRequestFields) HasCustomField3() bool`

HasCustomField3 returns a boolean if a field has been set.

### GetCustomField4

`func (o *GatewayCreateRequestFields) GetCustomField4() GatewayCreateRequestFieldsCustomField2`

GetCustomField4 returns the CustomField4 field if non-nil, zero value otherwise.

### GetCustomField4Ok

`func (o *GatewayCreateRequestFields) GetCustomField4Ok() (*GatewayCreateRequestFieldsCustomField2, bool)`

GetCustomField4Ok returns a tuple with the CustomField4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomField4

`func (o *GatewayCreateRequestFields) SetCustomField4(v GatewayCreateRequestFieldsCustomField2)`

SetCustomField4 sets CustomField4 field to given value.

### HasCustomField4

`func (o *GatewayCreateRequestFields) HasCustomField4() bool`

HasCustomField4 returns a boolean if a field has been set.

### GetCustomField5

`func (o *GatewayCreateRequestFields) GetCustomField5() GatewayCreateRequestFieldsCustomField2`

GetCustomField5 returns the CustomField5 field if non-nil, zero value otherwise.

### GetCustomField5Ok

`func (o *GatewayCreateRequestFields) GetCustomField5Ok() (*GatewayCreateRequestFieldsCustomField2, bool)`

GetCustomField5Ok returns a tuple with the CustomField5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomField5

`func (o *GatewayCreateRequestFields) SetCustomField5(v GatewayCreateRequestFieldsCustomField2)`

SetCustomField5 sets CustomField5 field to given value.

### HasCustomField5

`func (o *GatewayCreateRequestFields) HasCustomField5() bool`

HasCustomField5 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


