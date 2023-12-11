// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type StageSendSource struct {
	Account *StageSendSourceAccount `json:"account,omitempty"`
	Payment *StageSendSourcePayment `json:"payment,omitempty"`
	Wallet  *StageSendSourceWallet  `json:"wallet,omitempty"`
}

func (o *StageSendSource) GetAccount() *StageSendSourceAccount {
	if o == nil {
		return nil
	}
	return o.Account
}

func (o *StageSendSource) GetPayment() *StageSendSourcePayment {
	if o == nil {
		return nil
	}
	return o.Payment
}

func (o *StageSendSource) GetWallet() *StageSendSourceWallet {
	if o == nil {
		return nil
	}
	return o.Wallet
}
