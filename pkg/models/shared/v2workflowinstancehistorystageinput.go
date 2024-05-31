// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type V2WorkflowInstanceHistoryStageInput struct {
	AddAccountMetadata *V2ActivityAddAccountMetadata `json:"AddAccountMetadata,omitempty"`
	ConfirmHold        *V2ActivityConfirmHold        `json:"ConfirmHold,omitempty"`
	CreateTransaction  *V2ActivityCreateTransaction  `json:"CreateTransaction,omitempty"`
	CreditWallet       *V2ActivityCreditWallet       `json:"CreditWallet,omitempty"`
	DebitWallet        *V2ActivityDebitWallet        `json:"DebitWallet,omitempty"`
	GetAccount         *V2ActivityGetAccount         `json:"GetAccount,omitempty"`
	GetPayment         *V2ActivityGetPayment         `json:"GetPayment,omitempty"`
	GetWallet          *V2ActivityGetWallet          `json:"GetWallet,omitempty"`
	ListWallets        *V2ActivityListWallets        `json:"ListWallets,omitempty"`
	StripeTransfer     *V2ActivityStripeTransfer     `json:"StripeTransfer,omitempty"`
	VoidHold           *V2ActivityVoidHold           `json:"VoidHold,omitempty"`
}

func (o *V2WorkflowInstanceHistoryStageInput) GetAddAccountMetadata() *V2ActivityAddAccountMetadata {
	if o == nil {
		return nil
	}
	return o.AddAccountMetadata
}

func (o *V2WorkflowInstanceHistoryStageInput) GetConfirmHold() *V2ActivityConfirmHold {
	if o == nil {
		return nil
	}
	return o.ConfirmHold
}

func (o *V2WorkflowInstanceHistoryStageInput) GetCreateTransaction() *V2ActivityCreateTransaction {
	if o == nil {
		return nil
	}
	return o.CreateTransaction
}

func (o *V2WorkflowInstanceHistoryStageInput) GetCreditWallet() *V2ActivityCreditWallet {
	if o == nil {
		return nil
	}
	return o.CreditWallet
}

func (o *V2WorkflowInstanceHistoryStageInput) GetDebitWallet() *V2ActivityDebitWallet {
	if o == nil {
		return nil
	}
	return o.DebitWallet
}

func (o *V2WorkflowInstanceHistoryStageInput) GetGetAccount() *V2ActivityGetAccount {
	if o == nil {
		return nil
	}
	return o.GetAccount
}

func (o *V2WorkflowInstanceHistoryStageInput) GetGetPayment() *V2ActivityGetPayment {
	if o == nil {
		return nil
	}
	return o.GetPayment
}

func (o *V2WorkflowInstanceHistoryStageInput) GetGetWallet() *V2ActivityGetWallet {
	if o == nil {
		return nil
	}
	return o.GetWallet
}

func (o *V2WorkflowInstanceHistoryStageInput) GetListWallets() *V2ActivityListWallets {
	if o == nil {
		return nil
	}
	return o.ListWallets
}

func (o *V2WorkflowInstanceHistoryStageInput) GetStripeTransfer() *V2ActivityStripeTransfer {
	if o == nil {
		return nil
	}
	return o.StripeTransfer
}

func (o *V2WorkflowInstanceHistoryStageInput) GetVoidHold() *V2ActivityVoidHold {
	if o == nil {
		return nil
	}
	return o.VoidHold
}
