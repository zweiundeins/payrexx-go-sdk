# PayoutMerchant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**SiteTitle** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to [**PayoutMerchantOwner**](PayoutMerchantOwner.md) |  | [optional] 

## Methods

### NewPayoutMerchant

`func NewPayoutMerchant() *PayoutMerchant`

NewPayoutMerchant instantiates a new PayoutMerchant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayoutMerchantWithDefaults

`func NewPayoutMerchantWithDefaults() *PayoutMerchant`

NewPayoutMerchantWithDefaults instantiates a new PayoutMerchant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PayoutMerchant) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PayoutMerchant) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PayoutMerchant) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PayoutMerchant) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSiteTitle

`func (o *PayoutMerchant) GetSiteTitle() string`

GetSiteTitle returns the SiteTitle field if non-nil, zero value otherwise.

### GetSiteTitleOk

`func (o *PayoutMerchant) GetSiteTitleOk() (*string, bool)`

GetSiteTitleOk returns a tuple with the SiteTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiteTitle

`func (o *PayoutMerchant) SetSiteTitle(v string)`

SetSiteTitle sets SiteTitle field to given value.

### HasSiteTitle

`func (o *PayoutMerchant) HasSiteTitle() bool`

HasSiteTitle returns a boolean if a field has been set.

### GetOwner

`func (o *PayoutMerchant) GetOwner() PayoutMerchantOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *PayoutMerchant) GetOwnerOk() (*PayoutMerchantOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *PayoutMerchant) SetOwner(v PayoutMerchantOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *PayoutMerchant) HasOwner() bool`

HasOwner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


