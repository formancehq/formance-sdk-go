// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/formancehq/formance-sdk-go/v3/pkg/utils"
)

type V3BankingcircleConfig struct {
	AuthorizationEndpoint string  `json:"authorizationEndpoint"`
	Endpoint              string  `json:"endpoint"`
	Name                  string  `json:"name"`
	PageSize              *int64  `default:"25" json:"pageSize"`
	Password              string  `json:"password"`
	PollingPeriod         *string `default:"2m" json:"pollingPeriod"`
	UserCertificate       string  `json:"userCertificate"`
	UserCertificateKey    string  `json:"userCertificateKey"`
	Username              string  `json:"username"`
}

func (v V3BankingcircleConfig) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(v, "", false)
}

func (v *V3BankingcircleConfig) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &v, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *V3BankingcircleConfig) GetAuthorizationEndpoint() string {
	if o == nil {
		return ""
	}
	return o.AuthorizationEndpoint
}

func (o *V3BankingcircleConfig) GetEndpoint() string {
	if o == nil {
		return ""
	}
	return o.Endpoint
}

func (o *V3BankingcircleConfig) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *V3BankingcircleConfig) GetPageSize() *int64 {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *V3BankingcircleConfig) GetPassword() string {
	if o == nil {
		return ""
	}
	return o.Password
}

func (o *V3BankingcircleConfig) GetPollingPeriod() *string {
	if o == nil {
		return nil
	}
	return o.PollingPeriod
}

func (o *V3BankingcircleConfig) GetUserCertificate() string {
	if o == nil {
		return ""
	}
	return o.UserCertificate
}

func (o *V3BankingcircleConfig) GetUserCertificateKey() string {
	if o == nil {
		return ""
	}
	return o.UserCertificateKey
}

func (o *V3BankingcircleConfig) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}
