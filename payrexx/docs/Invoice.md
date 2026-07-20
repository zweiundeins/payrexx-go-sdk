# Invoice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**PaymentStatus** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Date** | Pointer to **string** |  | [optional] 
**Number** | Pointer to **int32** |  | [optional] [default to 0]
**Total** | Pointer to **int32** |  | [optional] [default to 0]
**Language** | Pointer to **string** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**DueAfterDays** | Pointer to **int32** |  | [optional] [default to 0]
**Recipient** | Pointer to [**InvoiceRecipient**](InvoiceRecipient.md) |  | [optional] 
**ServicePeriod** | Pointer to [**InvoiceServicePeriod**](InvoiceServicePeriod.md) |  | [optional] 
**Discount** | Pointer to [**InvoiceDiscount**](InvoiceDiscount.md) |  | [optional] 
**CashDiscounts** | Pointer to [**[]InvoiceCashDiscountsInner**](InvoiceCashDiscountsInner.md) |  | [optional] 
**ShippingCost** | Pointer to **int32** |  | [optional] [default to 0]
**ApplicationFee** | Pointer to **int32** |  | [optional] [default to 0]
**Note** | Pointer to **string** |  | [optional] 
**Terms** | Pointer to **string** |  | [optional] 
**Transactions** | Pointer to **[]string** |  | [optional] 
**Positions** | Pointer to [**[]InvoicePositionsInner**](InvoicePositionsInner.md) |  | [optional] 
**PaymentLink** | Pointer to **string** |  | [optional] 
**Attachments** | Pointer to **[]string** |  | [optional] 
**BankInformation** | Pointer to [**InvoiceBankInformation**](InvoiceBankInformation.md) |  | [optional] 
**PayoutDescriptor** | Pointer to **string** |  | [optional] 
**Psp** | Pointer to **[]string** |  | [optional] 
**Pm** | Pointer to **[]string** |  | [optional] 
**Mails** | Pointer to **[]string** |  | [optional] 
**Reference** | Pointer to **string** |  | [optional] 
**Design** | Pointer to **string** |  | [optional] 
**Reminders** | Pointer to [**[]InvoiceRemindersInner**](InvoiceRemindersInner.md) |  | [optional] 
**PurchaseOnInvoiceInformation** | Pointer to [**InvoicePurchaseOnInvoiceInformation**](InvoicePurchaseOnInvoiceInformation.md) |  | [optional] 

## Methods

### NewInvoice

`func NewInvoice() *Invoice`

NewInvoice instantiates a new Invoice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvoiceWithDefaults

`func NewInvoiceWithDefaults() *Invoice`

NewInvoiceWithDefaults instantiates a new Invoice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *Invoice) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Invoice) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Invoice) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Invoice) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPaymentStatus

`func (o *Invoice) GetPaymentStatus() string`

GetPaymentStatus returns the PaymentStatus field if non-nil, zero value otherwise.

### GetPaymentStatusOk

`func (o *Invoice) GetPaymentStatusOk() (*string, bool)`

GetPaymentStatusOk returns a tuple with the PaymentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentStatus

`func (o *Invoice) SetPaymentStatus(v string)`

SetPaymentStatus sets PaymentStatus field to given value.

### HasPaymentStatus

`func (o *Invoice) HasPaymentStatus() bool`

HasPaymentStatus returns a boolean if a field has been set.

### GetId

`func (o *Invoice) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Invoice) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Invoice) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Invoice) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDate

`func (o *Invoice) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *Invoice) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *Invoice) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *Invoice) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetNumber

`func (o *Invoice) GetNumber() int32`

GetNumber returns the Number field if non-nil, zero value otherwise.

### GetNumberOk

`func (o *Invoice) GetNumberOk() (*int32, bool)`

GetNumberOk returns a tuple with the Number field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumber

`func (o *Invoice) SetNumber(v int32)`

SetNumber sets Number field to given value.

### HasNumber

`func (o *Invoice) HasNumber() bool`

HasNumber returns a boolean if a field has been set.

### GetTotal

`func (o *Invoice) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *Invoice) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *Invoice) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *Invoice) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetLanguage

`func (o *Invoice) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *Invoice) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *Invoice) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *Invoice) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetCurrency

`func (o *Invoice) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Invoice) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Invoice) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *Invoice) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDueAfterDays

`func (o *Invoice) GetDueAfterDays() int32`

GetDueAfterDays returns the DueAfterDays field if non-nil, zero value otherwise.

### GetDueAfterDaysOk

`func (o *Invoice) GetDueAfterDaysOk() (*int32, bool)`

GetDueAfterDaysOk returns a tuple with the DueAfterDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDueAfterDays

`func (o *Invoice) SetDueAfterDays(v int32)`

SetDueAfterDays sets DueAfterDays field to given value.

### HasDueAfterDays

`func (o *Invoice) HasDueAfterDays() bool`

HasDueAfterDays returns a boolean if a field has been set.

### GetRecipient

`func (o *Invoice) GetRecipient() InvoiceRecipient`

GetRecipient returns the Recipient field if non-nil, zero value otherwise.

### GetRecipientOk

`func (o *Invoice) GetRecipientOk() (*InvoiceRecipient, bool)`

GetRecipientOk returns a tuple with the Recipient field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipient

`func (o *Invoice) SetRecipient(v InvoiceRecipient)`

SetRecipient sets Recipient field to given value.

### HasRecipient

`func (o *Invoice) HasRecipient() bool`

HasRecipient returns a boolean if a field has been set.

### GetServicePeriod

`func (o *Invoice) GetServicePeriod() InvoiceServicePeriod`

GetServicePeriod returns the ServicePeriod field if non-nil, zero value otherwise.

### GetServicePeriodOk

`func (o *Invoice) GetServicePeriodOk() (*InvoiceServicePeriod, bool)`

GetServicePeriodOk returns a tuple with the ServicePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePeriod

`func (o *Invoice) SetServicePeriod(v InvoiceServicePeriod)`

SetServicePeriod sets ServicePeriod field to given value.

### HasServicePeriod

`func (o *Invoice) HasServicePeriod() bool`

HasServicePeriod returns a boolean if a field has been set.

### GetDiscount

`func (o *Invoice) GetDiscount() InvoiceDiscount`

GetDiscount returns the Discount field if non-nil, zero value otherwise.

### GetDiscountOk

`func (o *Invoice) GetDiscountOk() (*InvoiceDiscount, bool)`

GetDiscountOk returns a tuple with the Discount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscount

`func (o *Invoice) SetDiscount(v InvoiceDiscount)`

SetDiscount sets Discount field to given value.

### HasDiscount

`func (o *Invoice) HasDiscount() bool`

HasDiscount returns a boolean if a field has been set.

### GetCashDiscounts

`func (o *Invoice) GetCashDiscounts() []InvoiceCashDiscountsInner`

GetCashDiscounts returns the CashDiscounts field if non-nil, zero value otherwise.

### GetCashDiscountsOk

`func (o *Invoice) GetCashDiscountsOk() (*[]InvoiceCashDiscountsInner, bool)`

GetCashDiscountsOk returns a tuple with the CashDiscounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashDiscounts

`func (o *Invoice) SetCashDiscounts(v []InvoiceCashDiscountsInner)`

SetCashDiscounts sets CashDiscounts field to given value.

### HasCashDiscounts

`func (o *Invoice) HasCashDiscounts() bool`

HasCashDiscounts returns a boolean if a field has been set.

### GetShippingCost

`func (o *Invoice) GetShippingCost() int32`

GetShippingCost returns the ShippingCost field if non-nil, zero value otherwise.

### GetShippingCostOk

`func (o *Invoice) GetShippingCostOk() (*int32, bool)`

GetShippingCostOk returns a tuple with the ShippingCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingCost

`func (o *Invoice) SetShippingCost(v int32)`

SetShippingCost sets ShippingCost field to given value.

### HasShippingCost

`func (o *Invoice) HasShippingCost() bool`

HasShippingCost returns a boolean if a field has been set.

### GetApplicationFee

`func (o *Invoice) GetApplicationFee() int32`

GetApplicationFee returns the ApplicationFee field if non-nil, zero value otherwise.

### GetApplicationFeeOk

`func (o *Invoice) GetApplicationFeeOk() (*int32, bool)`

GetApplicationFeeOk returns a tuple with the ApplicationFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationFee

`func (o *Invoice) SetApplicationFee(v int32)`

SetApplicationFee sets ApplicationFee field to given value.

### HasApplicationFee

`func (o *Invoice) HasApplicationFee() bool`

HasApplicationFee returns a boolean if a field has been set.

### GetNote

`func (o *Invoice) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *Invoice) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *Invoice) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *Invoice) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetTerms

`func (o *Invoice) GetTerms() string`

GetTerms returns the Terms field if non-nil, zero value otherwise.

### GetTermsOk

`func (o *Invoice) GetTermsOk() (*string, bool)`

GetTermsOk returns a tuple with the Terms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerms

`func (o *Invoice) SetTerms(v string)`

SetTerms sets Terms field to given value.

### HasTerms

`func (o *Invoice) HasTerms() bool`

HasTerms returns a boolean if a field has been set.

### GetTransactions

`func (o *Invoice) GetTransactions() []string`

GetTransactions returns the Transactions field if non-nil, zero value otherwise.

### GetTransactionsOk

`func (o *Invoice) GetTransactionsOk() (*[]string, bool)`

GetTransactionsOk returns a tuple with the Transactions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactions

`func (o *Invoice) SetTransactions(v []string)`

SetTransactions sets Transactions field to given value.

### HasTransactions

`func (o *Invoice) HasTransactions() bool`

HasTransactions returns a boolean if a field has been set.

### GetPositions

`func (o *Invoice) GetPositions() []InvoicePositionsInner`

GetPositions returns the Positions field if non-nil, zero value otherwise.

### GetPositionsOk

`func (o *Invoice) GetPositionsOk() (*[]InvoicePositionsInner, bool)`

GetPositionsOk returns a tuple with the Positions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositions

`func (o *Invoice) SetPositions(v []InvoicePositionsInner)`

SetPositions sets Positions field to given value.

### HasPositions

`func (o *Invoice) HasPositions() bool`

HasPositions returns a boolean if a field has been set.

### GetPaymentLink

`func (o *Invoice) GetPaymentLink() string`

GetPaymentLink returns the PaymentLink field if non-nil, zero value otherwise.

### GetPaymentLinkOk

`func (o *Invoice) GetPaymentLinkOk() (*string, bool)`

GetPaymentLinkOk returns a tuple with the PaymentLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentLink

`func (o *Invoice) SetPaymentLink(v string)`

SetPaymentLink sets PaymentLink field to given value.

### HasPaymentLink

`func (o *Invoice) HasPaymentLink() bool`

HasPaymentLink returns a boolean if a field has been set.

### GetAttachments

`func (o *Invoice) GetAttachments() []string`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *Invoice) GetAttachmentsOk() (*[]string, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *Invoice) SetAttachments(v []string)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *Invoice) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### GetBankInformation

`func (o *Invoice) GetBankInformation() InvoiceBankInformation`

GetBankInformation returns the BankInformation field if non-nil, zero value otherwise.

### GetBankInformationOk

`func (o *Invoice) GetBankInformationOk() (*InvoiceBankInformation, bool)`

GetBankInformationOk returns a tuple with the BankInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankInformation

`func (o *Invoice) SetBankInformation(v InvoiceBankInformation)`

SetBankInformation sets BankInformation field to given value.

### HasBankInformation

`func (o *Invoice) HasBankInformation() bool`

HasBankInformation returns a boolean if a field has been set.

### GetPayoutDescriptor

`func (o *Invoice) GetPayoutDescriptor() string`

GetPayoutDescriptor returns the PayoutDescriptor field if non-nil, zero value otherwise.

### GetPayoutDescriptorOk

`func (o *Invoice) GetPayoutDescriptorOk() (*string, bool)`

GetPayoutDescriptorOk returns a tuple with the PayoutDescriptor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayoutDescriptor

`func (o *Invoice) SetPayoutDescriptor(v string)`

SetPayoutDescriptor sets PayoutDescriptor field to given value.

### HasPayoutDescriptor

`func (o *Invoice) HasPayoutDescriptor() bool`

HasPayoutDescriptor returns a boolean if a field has been set.

### GetPsp

`func (o *Invoice) GetPsp() []string`

GetPsp returns the Psp field if non-nil, zero value otherwise.

### GetPspOk

`func (o *Invoice) GetPspOk() (*[]string, bool)`

GetPspOk returns a tuple with the Psp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPsp

`func (o *Invoice) SetPsp(v []string)`

SetPsp sets Psp field to given value.

### HasPsp

`func (o *Invoice) HasPsp() bool`

HasPsp returns a boolean if a field has been set.

### GetPm

`func (o *Invoice) GetPm() []string`

GetPm returns the Pm field if non-nil, zero value otherwise.

### GetPmOk

`func (o *Invoice) GetPmOk() (*[]string, bool)`

GetPmOk returns a tuple with the Pm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPm

`func (o *Invoice) SetPm(v []string)`

SetPm sets Pm field to given value.

### HasPm

`func (o *Invoice) HasPm() bool`

HasPm returns a boolean if a field has been set.

### GetMails

`func (o *Invoice) GetMails() []string`

GetMails returns the Mails field if non-nil, zero value otherwise.

### GetMailsOk

`func (o *Invoice) GetMailsOk() (*[]string, bool)`

GetMailsOk returns a tuple with the Mails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMails

`func (o *Invoice) SetMails(v []string)`

SetMails sets Mails field to given value.

### HasMails

`func (o *Invoice) HasMails() bool`

HasMails returns a boolean if a field has been set.

### GetReference

`func (o *Invoice) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *Invoice) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *Invoice) SetReference(v string)`

SetReference sets Reference field to given value.

### HasReference

`func (o *Invoice) HasReference() bool`

HasReference returns a boolean if a field has been set.

### GetDesign

`func (o *Invoice) GetDesign() string`

GetDesign returns the Design field if non-nil, zero value otherwise.

### GetDesignOk

`func (o *Invoice) GetDesignOk() (*string, bool)`

GetDesignOk returns a tuple with the Design field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDesign

`func (o *Invoice) SetDesign(v string)`

SetDesign sets Design field to given value.

### HasDesign

`func (o *Invoice) HasDesign() bool`

HasDesign returns a boolean if a field has been set.

### GetReminders

`func (o *Invoice) GetReminders() []InvoiceRemindersInner`

GetReminders returns the Reminders field if non-nil, zero value otherwise.

### GetRemindersOk

`func (o *Invoice) GetRemindersOk() (*[]InvoiceRemindersInner, bool)`

GetRemindersOk returns a tuple with the Reminders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReminders

`func (o *Invoice) SetReminders(v []InvoiceRemindersInner)`

SetReminders sets Reminders field to given value.

### HasReminders

`func (o *Invoice) HasReminders() bool`

HasReminders returns a boolean if a field has been set.

### GetPurchaseOnInvoiceInformation

`func (o *Invoice) GetPurchaseOnInvoiceInformation() InvoicePurchaseOnInvoiceInformation`

GetPurchaseOnInvoiceInformation returns the PurchaseOnInvoiceInformation field if non-nil, zero value otherwise.

### GetPurchaseOnInvoiceInformationOk

`func (o *Invoice) GetPurchaseOnInvoiceInformationOk() (*InvoicePurchaseOnInvoiceInformation, bool)`

GetPurchaseOnInvoiceInformationOk returns a tuple with the PurchaseOnInvoiceInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseOnInvoiceInformation

`func (o *Invoice) SetPurchaseOnInvoiceInformation(v InvoicePurchaseOnInvoiceInformation)`

SetPurchaseOnInvoiceInformation sets PurchaseOnInvoiceInformation field to given value.

### HasPurchaseOnInvoiceInformation

`func (o *Invoice) HasPurchaseOnInvoiceInformation() bool`

HasPurchaseOnInvoiceInformation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


