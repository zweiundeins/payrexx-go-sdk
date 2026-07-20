# Subscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [default to 0]
**Status** | Pointer to **string** |  | [optional] 
**Start** | Pointer to **string** |  | [optional] 
**End** | Pointer to **interface{}** |  | [optional] 
**ValidUntil** | Pointer to **string** |  | [optional] 
**Invoice** | Pointer to [**SubscriptionInvoice**](SubscriptionInvoice.md) |  | [optional] 
**Contact** | Pointer to [**SubscriptionContact**](SubscriptionContact.md) |  | [optional] 

## Methods

### NewSubscription

`func NewSubscription() *Subscription`

NewSubscription instantiates a new Subscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionWithDefaults

`func NewSubscriptionWithDefaults() *Subscription`

NewSubscriptionWithDefaults instantiates a new Subscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Subscription) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Subscription) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Subscription) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Subscription) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *Subscription) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Subscription) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Subscription) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Subscription) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStart

`func (o *Subscription) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *Subscription) GetStartOk() (*string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *Subscription) SetStart(v string)`

SetStart sets Start field to given value.

### HasStart

`func (o *Subscription) HasStart() bool`

HasStart returns a boolean if a field has been set.

### GetEnd

`func (o *Subscription) GetEnd() interface{}`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *Subscription) GetEndOk() (*interface{}, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *Subscription) SetEnd(v interface{})`

SetEnd sets End field to given value.

### HasEnd

`func (o *Subscription) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### SetEndNil

`func (o *Subscription) SetEndNil(b bool)`

 SetEndNil sets the value for End to be an explicit nil

### UnsetEnd
`func (o *Subscription) UnsetEnd()`

UnsetEnd ensures that no value is present for End, not even an explicit nil
### GetValidUntil

`func (o *Subscription) GetValidUntil() string`

GetValidUntil returns the ValidUntil field if non-nil, zero value otherwise.

### GetValidUntilOk

`func (o *Subscription) GetValidUntilOk() (*string, bool)`

GetValidUntilOk returns a tuple with the ValidUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidUntil

`func (o *Subscription) SetValidUntil(v string)`

SetValidUntil sets ValidUntil field to given value.

### HasValidUntil

`func (o *Subscription) HasValidUntil() bool`

HasValidUntil returns a boolean if a field has been set.

### GetInvoice

`func (o *Subscription) GetInvoice() SubscriptionInvoice`

GetInvoice returns the Invoice field if non-nil, zero value otherwise.

### GetInvoiceOk

`func (o *Subscription) GetInvoiceOk() (*SubscriptionInvoice, bool)`

GetInvoiceOk returns a tuple with the Invoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoice

`func (o *Subscription) SetInvoice(v SubscriptionInvoice)`

SetInvoice sets Invoice field to given value.

### HasInvoice

`func (o *Subscription) HasInvoice() bool`

HasInvoice returns a boolean if a field has been set.

### GetContact

`func (o *Subscription) GetContact() SubscriptionContact`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *Subscription) GetContactOk() (*SubscriptionContact, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *Subscription) SetContact(v SubscriptionContact)`

SetContact sets Contact field to given value.

### HasContact

`func (o *Subscription) HasContact() bool`

HasContact returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


