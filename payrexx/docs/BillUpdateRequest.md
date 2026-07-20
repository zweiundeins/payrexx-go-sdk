# BillUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | Pointer to **string** |  | [optional] 
**Language** | Pointer to **string** | The language ISO code by ISO 639-1. | [optional] 
**Currency** | Pointer to **string** | Currency code (ISO 4217) | [optional] 
**DueAfterDays** | Pointer to **int32** |  | [optional] 
**Recipient** | Pointer to [**BillCreateRequestRecipient**](BillCreateRequestRecipient.md) |  | [optional] 
**ServicePeriod** | Pointer to [**BillCreateRequestServicePeriod**](BillCreateRequestServicePeriod.md) |  | [optional] 
**Discount** | Pointer to [**BillCreateRequestDiscount**](BillCreateRequestDiscount.md) |  | [optional] 
**CashDiscounts** | Pointer to [**[]BillCreateRequestCashDiscountsInner**](BillCreateRequestCashDiscountsInner.md) |  | [optional] 
**ShippingCost** | Pointer to **int32** | Amount in the smallest unit of the currency. | [optional] 
**ApplicationFee** | Pointer to **int32** | Amount in the smallest unit of the currency. | [optional] 
**Note** | Pointer to **string** |  | [optional] 
**Terms** | Pointer to **string** |  | [optional] 
**Positions** | Pointer to [**[]BillCreateRequestPositionsInner**](BillCreateRequestPositionsInner.md) |  | [optional] 
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

### NewBillUpdateRequest

`func NewBillUpdateRequest() *BillUpdateRequest`

NewBillUpdateRequest instantiates a new BillUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillUpdateRequestWithDefaults

`func NewBillUpdateRequestWithDefaults() *BillUpdateRequest`

NewBillUpdateRequestWithDefaults instantiates a new BillUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *BillUpdateRequest) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *BillUpdateRequest) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *BillUpdateRequest) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *BillUpdateRequest) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetLanguage

`func (o *BillUpdateRequest) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *BillUpdateRequest) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *BillUpdateRequest) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *BillUpdateRequest) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetCurrency

`func (o *BillUpdateRequest) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *BillUpdateRequest) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *BillUpdateRequest) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *BillUpdateRequest) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDueAfterDays

`func (o *BillUpdateRequest) GetDueAfterDays() int32`

GetDueAfterDays returns the DueAfterDays field if non-nil, zero value otherwise.

### GetDueAfterDaysOk

`func (o *BillUpdateRequest) GetDueAfterDaysOk() (*int32, bool)`

GetDueAfterDaysOk returns a tuple with the DueAfterDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueAfterDays

`func (o *BillUpdateRequest) SetDueAfterDays(v int32)`

SetDueAfterDays sets DueAfterDays field to given value.

### HasDueAfterDays

`func (o *BillUpdateRequest) HasDueAfterDays() bool`

HasDueAfterDays returns a boolean if a field has been set.

### GetRecipient

`func (o *BillUpdateRequest) GetRecipient() BillCreateRequestRecipient`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *BillUpdateRequest) GetRecipientOk() (*BillCreateRequestRecipient, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *BillUpdateRequest) SetRecipient(v BillCreateRequestRecipient)`

SetRecipient sets Recipient field to given value.

### HasRecipient

`func (o *BillUpdateRequest) HasRecipient() bool`

HasRecipient returns a boolean if a field has been set.

### GetServicePeriod

`func (o *BillUpdateRequest) GetServicePeriod() BillCreateRequestServicePeriod`

GetServicePeriod returns the ServicePeriod field if non-nil, zero value otherwise.

### GetServicePeriodOk

`func (o *BillUpdateRequest) GetServicePeriodOk() (*BillCreateRequestServicePeriod, bool)`

GetServicePeriodOk returns a tuple with the ServicePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePeriod

`func (o *BillUpdateRequest) SetServicePeriod(v BillCreateRequestServicePeriod)`

SetServicePeriod sets ServicePeriod field to given value.

### HasServicePeriod

`func (o *BillUpdateRequest) HasServicePeriod() bool`

HasServicePeriod returns a boolean if a field has been set.

### GetDiscount

`func (o *BillUpdateRequest) GetDiscount() BillCreateRequestDiscount`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *BillUpdateRequest) GetDiscountOk() (*BillCreateRequestDiscount, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *BillUpdateRequest) SetDiscount(v BillCreateRequestDiscount)`

SetDiscount sets Discount field to given value.

### HasDiscount

`func (o *BillUpdateRequest) HasDiscount() bool`

HasDiscount returns a boolean if a field has been set.

### GetCashDiscounts

`func (o *BillUpdateRequest) GetCashDiscounts() []BillCreateRequestCashDiscountsInner`

GetCashDiscounts returns the CashDiscounts field if non-nil, zero value otherwise.

### GetCashDiscountsOk

`func (o *BillUpdateRequest) GetCashDiscountsOk() (*[]BillCreateRequestCashDiscountsInner, bool)`

GetCashDiscountsOk returns a tuple with the CashDiscounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashDiscounts

`func (o *BillUpdateRequest) SetCashDiscounts(v []BillCreateRequestCashDiscountsInner)`

SetCashDiscounts sets CashDiscounts field to given value.

### HasCashDiscounts

`func (o *BillUpdateRequest) HasCashDiscounts() bool`

HasCashDiscounts returns a boolean if a field has been set.

### GetShippingCost

`func (o *BillUpdateRequest) GetShippingCost() int32`

GetShippingCost returns the ShippingCost field if non-nil, zero value otherwise.

### GetShippingCostOk

`func (o *BillUpdateRequest) GetShippingCostOk() (*int32, bool)`

GetShippingCostOk returns a tuple with the ShippingCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingCost

`func (o *BillUpdateRequest) SetShippingCost(v int32)`

SetShippingCost sets ShippingCost field to given value.

### HasShippingCost

`func (o *BillUpdateRequest) HasShippingCost() bool`

HasShippingCost returns a boolean if a field has been set.

### GetApplicationFee

`func (o *BillUpdateRequest) GetApplicationFee() int32`

GetApplicationFee returns the ApplicationFee field if non-nil, zero value otherwise.

### GetApplicationFeeOk

`func (o *BillUpdateRequest) GetApplicationFeeOk() (*int32, bool)`

GetApplicationFeeOk returns a tuple with the ApplicationFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationFee

`func (o *BillUpdateRequest) SetApplicationFee(v int32)`

SetApplicationFee sets ApplicationFee field to given value.

### HasApplicationFee

`func (o *BillUpdateRequest) HasApplicationFee() bool`

HasApplicationFee returns a boolean if a field has been set.

### GetNote

`func (o *BillUpdateRequest) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *BillUpdateRequest) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *BillUpdateRequest) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *BillUpdateRequest) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetTerms

`func (o *BillUpdateRequest) GetTerms() string`

GetTerms returns the Terms field if non-nil, zero value otherwise.

### GetTermsOk

`func (o *BillUpdateRequest) GetTermsOk() (*string, bool)`

GetTermsOk returns a tuple with the Terms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerms

`func (o *BillUpdateRequest) SetTerms(v string)`

SetTerms sets Terms field to given value.

### HasTerms

`func (o *BillUpdateRequest) HasTerms() bool`

HasTerms returns a boolean if a field has been set.

### GetPositions

`func (o *BillUpdateRequest) GetPositions() []BillCreateRequestPositionsInner`

GetPositions returns the Positions field if non-nil, zero value otherwise.

### GetPositionsOk

`func (o *BillUpdateRequest) GetPositionsOk() (*[]BillCreateRequestPositionsInner, bool)`

GetPositionsOk returns a tuple with the Positions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositions

`func (o *BillUpdateRequest) SetPositions(v []BillCreateRequestPositionsInner)`

SetPositions sets Positions field to given value.

### HasPositions

`func (o *BillUpdateRequest) HasPositions() bool`

HasPositions returns a boolean if a field has been set.

### GetAttachments

`func (o *BillUpdateRequest) GetAttachments() []BillCreateRequestAttachmentsInner`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *BillUpdateRequest) GetAttachmentsOk() (*[]BillCreateRequestAttachmentsInner, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *BillUpdateRequest) SetAttachments(v []BillCreateRequestAttachmentsInner)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *BillUpdateRequest) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetBankInformation

`func (o *BillUpdateRequest) GetBankInformation() BillCreateRequestBankInformation`

GetBankInformation returns the BankInformation field if non-nil, zero value otherwise.

### GetBankInformationOk

`func (o *BillUpdateRequest) GetBankInformationOk() (*BillCreateRequestBankInformation, bool)`

GetBankInformationOk returns a tuple with the BankInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankInformation

`func (o *BillUpdateRequest) SetBankInformation(v BillCreateRequestBankInformation)`

SetBankInformation sets BankInformation field to given value.

### HasBankInformation

`func (o *BillUpdateRequest) HasBankInformation() bool`

HasBankInformation returns a boolean if a field has been set.

### GetPayoutDescriptor

`func (o *BillUpdateRequest) GetPayoutDescriptor() string`

GetPayoutDescriptor returns the PayoutDescriptor field if non-nil, zero value otherwise.

### GetPayoutDescriptorOk

`func (o *BillUpdateRequest) GetPayoutDescriptorOk() (*string, bool)`

GetPayoutDescriptorOk returns a tuple with the PayoutDescriptor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayoutDescriptor

`func (o *BillUpdateRequest) SetPayoutDescriptor(v string)`

SetPayoutDescriptor sets PayoutDescriptor field to given value.

### HasPayoutDescriptor

`func (o *BillUpdateRequest) HasPayoutDescriptor() bool`

HasPayoutDescriptor returns a boolean if a field has been set.

### GetPsp

`func (o *BillUpdateRequest) GetPsp() []int32`

GetPsp returns the Psp field if non-nil, zero value otherwise.

### GetPspOk

`func (o *BillUpdateRequest) GetPspOk() (*[]int32, bool)`

GetPspOk returns a tuple with the Psp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPsp

`func (o *BillUpdateRequest) SetPsp(v []int32)`

SetPsp sets Psp field to given value.

### HasPsp

`func (o *BillUpdateRequest) HasPsp() bool`

HasPsp returns a boolean if a field has been set.

### GetPm

`func (o *BillUpdateRequest) GetPm() []string`

GetPm returns the Pm field if non-nil, zero value otherwise.

### GetPmOk

`func (o *BillUpdateRequest) GetPmOk() (*[]string, bool)`

GetPmOk returns a tuple with the Pm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPm

`func (o *BillUpdateRequest) SetPm(v []string)`

SetPm sets Pm field to given value.

### HasPm

`func (o *BillUpdateRequest) HasPm() bool`

HasPm returns a boolean if a field has been set.

### GetReminders

`func (o *BillUpdateRequest) GetReminders() []BillCreateRequestRemindersInner`

GetReminders returns the Reminders field if non-nil, zero value otherwise.

### GetRemindersOk

`func (o *BillUpdateRequest) GetRemindersOk() (*[]BillCreateRequestRemindersInner, bool)`

GetRemindersOk returns a tuple with the Reminders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReminders

`func (o *BillUpdateRequest) SetReminders(v []BillCreateRequestRemindersInner)`

SetReminders sets Reminders field to given value.

### HasReminders

`func (o *BillUpdateRequest) HasReminders() bool`

HasReminders returns a boolean if a field has been set.

### GetReference

`func (o *BillUpdateRequest) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *BillUpdateRequest) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *BillUpdateRequest) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *BillUpdateRequest) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetDesign

`func (o *BillUpdateRequest) GetDesign() string`

GetDesign returns the Design field if non-nil, zero value otherwise.

### GetDesignOk

`func (o *BillUpdateRequest) GetDesignOk() (*string, bool)`

GetDesignOk returns a tuple with the Design field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesign

`func (o *BillUpdateRequest) SetDesign(v string)`

SetDesign sets Design field to given value.

### HasDesign

`func (o *BillUpdateRequest) HasDesign() bool`

HasDesign returns a boolean if a field has been set.

### GetSend

`func (o *BillUpdateRequest) GetSend() bool`

GetSend returns the Send field if non-nil, zero value otherwise.

### GetSendOk

`func (o *BillUpdateRequest) GetSendOk() (*bool, bool)`

GetSendOk returns a tuple with the Send field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSend

`func (o *BillUpdateRequest) SetSend(v bool)`

SetSend sets Send field to given value.

### HasSend

`func (o *BillUpdateRequest) HasSend() bool`

HasSend returns a boolean if a field has been set.

### GetAdditionalRecipients

`func (o *BillUpdateRequest) GetAdditionalRecipients() []string`

GetAdditionalRecipients returns the AdditionalRecipients field if non-nil, zero value otherwise.

### GetAdditionalRecipientsOk

`func (o *BillUpdateRequest) GetAdditionalRecipientsOk() (*[]string, bool)`

GetAdditionalRecipientsOk returns a tuple with the AdditionalRecipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalRecipients

`func (o *BillUpdateRequest) SetAdditionalRecipients(v []string)`

SetAdditionalRecipients sets AdditionalRecipients field to given value.

### HasAdditionalRecipients

`func (o *BillUpdateRequest) HasAdditionalRecipients() bool`

HasAdditionalRecipients returns a boolean if a field has been set.

### GetComplete

`func (o *BillUpdateRequest) GetComplete() bool`

GetComplete returns the Complete field if non-nil, zero value otherwise.

### GetCompleteOk

`func (o *BillUpdateRequest) GetCompleteOk() (*bool, bool)`

GetCompleteOk returns a tuple with the Complete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplete

`func (o *BillUpdateRequest) SetComplete(v bool)`

SetComplete sets Complete field to given value.

### HasComplete

`func (o *BillUpdateRequest) HasComplete() bool`

HasComplete returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


