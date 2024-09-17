// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package formancesdkgo

type Payments struct {
	V1 *FormancePaymentsV1

	sdkConfiguration sdkConfiguration
}

func newPayments(sdkConfig sdkConfiguration) *Payments {
	return &Payments{
		sdkConfiguration: sdkConfig,
		V1:               newFormancePaymentsV1(sdkConfig),
	}
}
