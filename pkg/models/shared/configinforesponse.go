// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type ConfigInfoResponse struct {
	Data ConfigInfo `json:"data"`
}

func (o *ConfigInfoResponse) GetData() ConfigInfo {
	if o == nil {
		return ConfigInfo{}
	}
	return o.Data
}
