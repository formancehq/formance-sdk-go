// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package shared

type ActivityDebitWalletOutput struct {
	Data Hold `json:"data"`
}

func (o *ActivityDebitWalletOutput) GetData() Hold {
	if o == nil {
		return Hold{}
	}
	return o.Data
}
