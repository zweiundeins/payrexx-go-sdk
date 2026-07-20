# BillCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | Pointer to **string** |  | [optional] 
**Language** | **string** | The language ISO code by ISO 639-1. | 
**Currency** | **string** | Currency code (ISO 4217) | 
**DueAfterDays** | **int32** |  | 
**Recipient** | [**BillCreateRequestRecipient**](BillCreateRequestRecipient.md) |  | 
**ServicePeriod** | Pointer to [**BillCreateRequestServicePeriod**](BillCreateRequestServicePeriod.md) |  | [optional] 
**Discount** | Pointer to [**BillCreateRequestDiscount**](BillCreateRequestDiscount.md) |  | [optional] 
**CashDiscounts** | Pointer to [**[]BillCreateRequestCashDiscountsInner**](BillCreateRequestCashDiscountsInner.md) |  | [optional] 
**ShippingCost** | Pointer to **int32** | Amount in the smallest unit of the currency. | [optional] 
**ApplicationFee** | Pointer to **int32** | Amount in the smallest unit of the currency. | [optional] 
**Note** | Pointer to **string** |  | [optional] 
**Terms** | Pointer to **string** |  | [optional] 
**Positions** | [**[]BillCreateRequestPositionsInner**](BillCreateRequestPositionsInner.md) |  | 
**Attachments** | Pointer to [**[]BillCreateRequestAttachmentsInner**](BillCreateRequestAttachmentsInner.md) |  | [optional] 
**BankInformation** | Pointer to [**BillCreateRequestBankInformation**](BillCreateRequestBankInformation.md) |  | [optional] 
**PayoutDescriptor** | Pointer to **string** | This payout descriptor is added to the payout statement (Only for Payrexx Swiss Collecting payments and single transaction payouts). Max. 80 chars. | [optional] 
**Psp** | Pointer to **[]int32** |  | [optional] 
**Pm** | Pointer to **[]string** |  | [optional] 
**Reminders** | Pointer to [**[]BillCreateRequestRemindersInner**](BillCreateRequestRemindersInner.md) |  | [optional] 
**Reference** | Pointer to **string** |  | [optional] 
**Design** | Pointer to **string** |  | [optional] 
**Send** | Pointer to **bool** | Whether to send the invoice. | [optional] [default to false]
**AdditionalRecipients** | Pointer to **[]string** | Send invoice to additional recipients. Pass email addresses as Strings. | [optional] 
**Complete** | Pointer to **bool** | Whether to complete the invoice. | [optional] [default to false]

## Methods

### NewBillCreateRequest

`func NewBillCreateRequest(language string, currency string, dueAfterDays int32, recipient BillCreateRequestRecipient, positions []BillCreateRequestPositionsInner, ) *BillCreateRequest`

NewBillCreateRequest instantiates a new BillCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillCreateRequestWithDefaults

`func NewBillCreateRequestWithDefaults() *BillCreateRequest`

NewBillCreateRequestWithDefaults instantiates a new BillCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *BillCreateRequest) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *BillCreateRequest) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *BillCreateRequest) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *BillCreateRequest) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetLanguage

`func (o *BillCreateRequest) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *BillCreateRequest) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *BillCreateRequest) SetLanguage(v string)`

SetLanguage sets Language field to given value.


### GetCurrency

`func (o *BillCreateRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *BillCreateRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *BillCreateRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetDueAfterDays

`func (o *BillCreateRequest) GetDueAfterDays() int32`

GetDueAfterDays returns the DueAfterDays field if non-nil, zero value otherwise.

### GetDueAfterDaysOk

`func (o *BillCreateRequest) GetDueAfterDaysOk() (*int32, bool)`

GetDueAfterDaysOk returns a tuple with the DueAfterDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueAfterDays

`func (o *BillCreateRequest) SetDueAfterDays(v int32)`

SetDueAfterDays sets DueAfterDays field to given value.


### GetRecipient

`func (o *BillCreateRequest) GetRecipient() BillCreateRequestRecipient`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *BillCreateRequest) GetRecipientOk() (*BillCreateRequestRecipient, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *BillCreateRequest) SetRecipient(v BillCreateRequestRecipient)`

SetRecipient sets Recipient field to given value.


### GetServicePeriod

`func (o *BillCreateRequest) GetServicePeriod() BillCreateRequestServicePeriod`

GetServicePeriod returns the ServicePeriod field if non-nil, zero value otherwise.

### GetServicePeriodOk

`func (o *BillCreateRequest) GetServicePeriodOk() (*BillCreateRequestServicePeriod, bool)`

GetServicePeriodOk returns a tuple with the ServicePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePeriod

`func (o *BillCreateRequest) SetServicePeriod(v BillCreateRequestServicePeriod)`

SetServicePeriod sets ServicePeriod field to given value.

### HasServicePeriod

`func (o *BillCreateRequest) HasServicePeriod() bool`

HasServicePeriod returns a boolean if a field has been set.

### GetDiscount

`func (o *BillCreateRequest) GetDiscount() BillCreateRequestDiscount`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *BillCreateRequest) GetDiscountOk() (*BillCreateRequestDiscount, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *BillCreateRequest) SetDiscount(v BillCreateRequestDiscount)`

SetDiscount sets Discount field to given value.

### HasDiscount

`func (o *BillCreateRequest) HasDiscount() bool`

HasDiscount returns a boolean if a field has been set.

### GetCashDiscounts

`func (o *BillCreateRequest) GetCashDiscounts() []BillCreateRequestCashDiscountsInner`

GetCashDiscounts returns the CashDiscounts field if non-nil, zero value otherwise.

### GetCashDiscountsOk

`func (o *BillCreateRequest) GetCashDiscountsOk() (*[]BillCreateRequestCashDiscountsInner, bool)`

GetCashDiscountsOk returns a tuple with the CashDiscounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashDiscounts

`func (o *BillCreateRequest) SetCashDiscounts(v []BillCreateRequestCashDiscountsInner)`

SetCashDiscounts sets CashDiscounts field to given value.

### HasCashDiscounts

`func (o *BillCreateRequest) HasCashDiscounts() bool`

HasCashDiscounts returns a boolean if a field has been set.

### GetShippingCost

`func (o *BillCreateRequest) GetShippingCost() int32`

GetShippingCost returns the ShippingCost field if non-nil, zero value otherwise.

### GetShippingCostOk

`func (o *BillCreateRequest) GetShippingCostOk() (*int32, bool)`

GetShippingCostOk returns a tuple with the ShippingCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingCost

`func (o *BillCreateRequest) SetShippingCost(v int32)`

SetShippingCost sets ShippingCost field to given value.

### HasShippingCost

`func (o *BillCreateRequest) HasShippingCost() bool`

HasShippingCost returns a boolean if a field has been set.

### GetApplicationFee

`func (o *BillCreateRequest) GetApplicationFee() int32`

GetApplicationFee returns the ApplicationFee field if non-nil, zero value otherwise.

### GetApplicationFeeOk

`func (o *BillCreateRequest) GetApplicationFeeOk() (*int32, bool)`

GetApplicationFeeOk returns a tuple with the ApplicationFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationFee

`func (o *BillCreateRequest) SetApplicationFee(v int32)`

SetApplicationFee sets ApplicationFee field to given value.

### HasApplicationFee

`func (o *BillCreateRequest) HasApplicationFee() bool`

HasApplicationFee returns a boolean if a field has been set.

### GetNote

`func (o *BillCreateRequest) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *BillCreateRequest) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *BillCreateRequest) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *BillCreateRequest) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetTerms

`func (o *BillCreateRequest) GetTerms() string`

GetTerms returns the Terms field if non-nil, zero value otherwise.

### GetTermsOk

`func (o *BillCreateRequest) GetTermsOk() (*string, bool)`

GetTermsOk returns a tuple with the Terms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerms

`func (o *BillCreateRequest) SetTerms(v string)`

SetTerms sets Terms field to given value.

### HasTerms

`func (o *BillCreateRequest) HasTerms() bool`

HasTerms returns a boolean if a field has been set.

### GetPositions

`func (o *BillCreateRequest) GetPositions() []BillCreateRequestPositionsInner`

GetPositions returns the Positions field if non-nil, zero value otherwise.

### GetPositionsOk

`func (o *BillCreateRequest) GetPositionsOk() (*[]BillCreateRequestPositionsInner, bool)`

GetPositionsOk returns a tuple with the Positions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositions

`func (o *BillCreateRequest) SetPositions(v []BillCreateRequestPositionsInner)`

SetPositions sets Positions field to given value.


### GetAttachments

`func (o *BillCreateRequest) GetAttachments() []BillCreateRequestAttachmentsInner`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *BillCreateRequest) GetAttachmentsOk() (*[]BillCreateRequestAttachmentsInner, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *BillCreateRequest) SetAttachments(v []BillCreateRequestAttachmentsInner)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *BillCreateRequest) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetBankInformation

`func (o *BillCreateRequest) GetBankInformation() BillCreateRequestBankInformation`

GetBankInformation returns the BankInformation field if non-nil, zero value otherwise.

### GetBankInformationOk

`func (o *BillCreateRequest) GetBankInformationOk() (*BillCreateRequestBankInformation, bool)`

GetBankInformationOk returns a tuple with the BankInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankInformation

`func (o *BillCreateRequest) SetBankInformation(v BillCreateRequestBankInformation)`

SetBankInformation sets BankInformation field to given value.

### HasBankInformation

`func (o *BillCreateRequest) HasBankInformation() bool`

HasBankInformation returns a boolean if a field has been set.

### GetPayoutDescriptor

`func (o *BillCreateRequest) GetPayoutDescriptor() string`

GetPayoutDescriptor returns the PayoutDescriptor field if non-nil, zero value otherwise.

### GetPayoutDescriptorOk

`func (o *BillCreateRequest) GetPayoutDescriptorOk() (*string, bool)`

GetPayoutDescriptorOk returns a tuple with the PayoutDescriptor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayoutDescriptor

`func (o *BillCreateRequest) SetPayoutDescriptor(v string)`

SetPayoutDescriptor sets PayoutDescriptor field to given value.

### HasPayoutDescriptor

`func (o *BillCreateRequest) HasPayoutDescriptor() bool`

HasPayoutDescriptor returns a boolean if a field has been set.

### GetPsp

`func (o *BillCreateRequest) GetPsp() []int32`

GetPsp returns the Psp field if non-nil, zero value otherwise.

### GetPspOk

`func (o *BillCreateRequest) GetPspOk() (*[]int32, bool)`

GetPspOk returns a tuple with the Psp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPsp

`func (o *BillCreateRequest) SetPsp(v []int32)`

SetPsp sets Psp field to given value.

### HasPsp

`func (o *BillCreateRequest) HasPsp() bool`

HasPsp returns a boolean if a field has been set.

### GetPm

`func (o *BillCreateRequest) GetPm() []string`

GetPm returns the Pm field if non-nil, zero value otherwise.

### GetPmOk

`func (o *BillCreateRequest) GetPmOk() (*[]string, bool)`

GetPmOk returns a tuple with the Pm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPm

`func (o *BillCreateRequest) SetPm(v []string)`

SetPm sets Pm field to given value.

### HasPm

`func (o *BillCreateRequest) HasPm() bool`

HasPm returns a boolean if a field has been set.

### GetReminders

`func (o *BillCreateRequest) GetReminders() []BillCreateRequestRemindersInner`

GetReminders returns the Reminders field if non-nil, zero value otherwise.

### GetRemindersOk

`func (o *BillCreateRequest) GetRemindersOk() (*[]BillCreateRequestRemindersInner, bool)`

GetRemindersOk returns a tuple with the Reminders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReminders

`func (o *BillCreateRequest) SetReminders(v []BillCreateRequestRemindersInner)`

SetReminders sets Reminders field to given value.

### HasReminders

`func (o *BillCreateRequest) HasReminders() bool`

HasReminders returns a boolean if a field has been set.

### GetReference

`func (o *BillCreateRequest) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *BillCreateRequest) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *BillCreateRequest) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *BillCreateRequest) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetDesign

`func (o *BillCreateRequest) GetDesign() string`

GetDesign returns the Design field if non-nil, zero value otherwise.

### GetDesignOk

`func (o *BillCreateRequest) GetDesignOk() (*string, bool)`

GetDesignOk returns a tuple with the Design field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesign

`func (o *BillCreateRequest) SetDesign(v string)`

SetDesign sets Design field to given value.

### HasDesign

`func (o *BillCreateRequest) HasDesign() bool`

HasDesign returns a boolean if a field has been set.

### GetSend

`func (o *BillCreateRequest) GetSend() bool`

GetSend returns the Send field if non-nil, zero value otherwise.

### GetSendOk

`func (o *BillCreateRequest) GetSendOk() (*bool, bool)`

GetSendOk returns a tuple with the Send field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSend

`func (o *BillCreateRequest) SetSend(v bool)`

SetSend sets Send field to given value.

### HasSend

`func (o *BillCreateRequest) HasSend() bool`

HasSend returns a boolean if a field has been set.

### GetAdditionalRecipients

`func (o *BillCreateRequest) GetAdditionalRecipients() []string`

GetAdditionalRecipients returns the AdditionalRecipients field if non-nil, zero value otherwise.

### GetAdditionalRecipientsOk

`func (o *BillCreateRequest) GetAdditionalRecipientsOk() (*[]string, bool)`

GetAdditionalRecipientsOk returns a tuple with the AdditionalRecipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalRecipients

`func (o *BillCreateRequest) SetAdditionalRecipients(v []string)`

SetAdditionalRecipients sets AdditionalRecipients field to given value.

### HasAdditionalRecipients

`func (o *BillCreateRequest) HasAdditionalRecipients() bool`

HasAdditionalRecipients returns a boolean if a field has been set.

### GetComplete

`func (o *BillCreateRequest) GetComplete() bool`

GetComplete returns the Complete field if non-nil, zero value otherwise.

### GetCompleteOk

`func (o *BillCreateRequest) GetCompleteOk() (*bool, bool)`

GetCompleteOk returns a tuple with the Complete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplete

`func (o *BillCreateRequest) SetComplete(v bool)`

SetComplete sets Complete field to given value.

### HasComplete

`func (o *BillCreateRequest) HasComplete() bool`

HasComplete returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


