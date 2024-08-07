// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package shared

type Security struct {
	Authorization string `security:"scheme,type=oauth2,name=Authorization"`
}

func (o *Security) GetAuthorization() string {
	if o == nil {
		return ""
	}
	return o.Authorization
}
