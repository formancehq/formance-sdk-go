// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type V3PaymentInitiationRelatedPaymentsCursorResponseCursor struct {
	Data     []V3Payment `json:"data"`
	HasMore  bool        `json:"hasMore"`
	Next     *string     `json:"next,omitempty"`
	PageSize int64       `json:"pageSize"`
	Previous *string     `json:"previous,omitempty"`
}

func (o *V3PaymentInitiationRelatedPaymentsCursorResponseCursor) GetData() []V3Payment {
	if o == nil {
		return []V3Payment{}
	}
	return o.Data
}

func (o *V3PaymentInitiationRelatedPaymentsCursorResponseCursor) GetHasMore() bool {
	if o == nil {
		return false
	}
	return o.HasMore
}

func (o *V3PaymentInitiationRelatedPaymentsCursorResponseCursor) GetNext() *string {
	if o == nil {
		return nil
	}
	return o.Next
}

func (o *V3PaymentInitiationRelatedPaymentsCursorResponseCursor) GetPageSize() int64 {
	if o == nil {
		return 0
	}
	return o.PageSize
}

func (o *V3PaymentInitiationRelatedPaymentsCursorResponseCursor) GetPrevious() *string {
	if o == nil {
		return nil
	}
	return o.Previous
}

type V3PaymentInitiationRelatedPaymentsCursorResponse struct {
	Cursor V3PaymentInitiationRelatedPaymentsCursorResponseCursor `json:"cursor"`
}

func (o *V3PaymentInitiationRelatedPaymentsCursorResponse) GetCursor() V3PaymentInitiationRelatedPaymentsCursorResponseCursor {
	if o == nil {
		return V3PaymentInitiationRelatedPaymentsCursorResponseCursor{}
	}
	return o.Cursor
}
