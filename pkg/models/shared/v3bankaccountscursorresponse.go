// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type V3BankAccountsCursorResponseCursor struct {
	Data     []V3BankAccount `json:"data"`
	HasMore  bool            `json:"hasMore"`
	Next     *string         `json:"next,omitempty"`
	PageSize int64           `json:"pageSize"`
	Previous *string         `json:"previous,omitempty"`
}

func (o *V3BankAccountsCursorResponseCursor) GetData() []V3BankAccount {
	if o == nil {
		return []V3BankAccount{}
	}
	return o.Data
}

func (o *V3BankAccountsCursorResponseCursor) GetHasMore() bool {
	if o == nil {
		return false
	}
	return o.HasMore
}

func (o *V3BankAccountsCursorResponseCursor) GetNext() *string {
	if o == nil {
		return nil
	}
	return o.Next
}

func (o *V3BankAccountsCursorResponseCursor) GetPageSize() int64 {
	if o == nil {
		return 0
	}
	return o.PageSize
}

func (o *V3BankAccountsCursorResponseCursor) GetPrevious() *string {
	if o == nil {
		return nil
	}
	return o.Previous
}

type V3BankAccountsCursorResponse struct {
	Cursor V3BankAccountsCursorResponseCursor `json:"cursor"`
}

func (o *V3BankAccountsCursorResponse) GetCursor() V3BankAccountsCursorResponseCursor {
	if o == nil {
		return V3BankAccountsCursorResponseCursor{}
	}
	return o.Cursor
}
