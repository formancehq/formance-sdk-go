// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type ContractExpr struct {
}

type Contract struct {
	Account *string      `json:"account,omitempty"`
	Expr    ContractExpr `json:"expr"`
}

func (o *Contract) GetAccount() *string {
	if o == nil {
		return nil
	}
	return o.Account
}

func (o *Contract) GetExpr() ContractExpr {
	if o == nil {
		return ContractExpr{}
	}
	return o.Expr
}
