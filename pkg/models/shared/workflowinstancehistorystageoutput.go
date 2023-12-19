// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type WorkflowInstanceHistoryStageOutput struct {
	CreateTransaction *ActivityCreateTransactionOutput `json:"CreateTransaction,omitempty"`
	DebitWallet       *ActivityDebitWalletOutput       `json:"DebitWallet,omitempty"`
	GetAccount        *ActivityGetAccountOutput        `json:"GetAccount,omitempty"`
	GetPayment        *ActivityGetPaymentOutput        `json:"GetPayment,omitempty"`
	GetWallet         *ActivityGetWalletOutput         `json:"GetWallet,omitempty"`
}

func (o *WorkflowInstanceHistoryStageOutput) GetCreateTransaction() *ActivityCreateTransactionOutput {
	if o == nil {
		return nil
	}
	return o.CreateTransaction
}

func (o *WorkflowInstanceHistoryStageOutput) GetDebitWallet() *ActivityDebitWalletOutput {
	if o == nil {
		return nil
	}
	return o.DebitWallet
}

func (o *WorkflowInstanceHistoryStageOutput) GetGetAccount() *ActivityGetAccountOutput {
	if o == nil {
		return nil
	}
	return o.GetAccount
}

func (o *WorkflowInstanceHistoryStageOutput) GetGetPayment() *ActivityGetPaymentOutput {
	if o == nil {
		return nil
	}
	return o.GetPayment
}

func (o *WorkflowInstanceHistoryStageOutput) GetGetWallet() *ActivityGetWalletOutput {
	if o == nil {
		return nil
	}
	return o.GetWallet
}
