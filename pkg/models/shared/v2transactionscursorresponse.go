// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type V2TransactionsCursorResponseCursor struct {
	Data     []V2ExpandedTransaction `json:"data"`
	HasMore  bool                    `json:"hasMore"`
	Next     *string                 `json:"next,omitempty"`
	PageSize int64                   `json:"pageSize"`
	Previous *string                 `json:"previous,omitempty"`
}

func (o *V2TransactionsCursorResponseCursor) GetData() []V2ExpandedTransaction {
	if o == nil {
		return []V2ExpandedTransaction{}
	}
	return o.Data
}

func (o *V2TransactionsCursorResponseCursor) GetHasMore() bool {
	if o == nil {
		return false
	}
	return o.HasMore
}

func (o *V2TransactionsCursorResponseCursor) GetNext() *string {
	if o == nil {
		return nil
	}
	return o.Next
}

func (o *V2TransactionsCursorResponseCursor) GetPageSize() int64 {
	if o == nil {
		return 0
	}
	return o.PageSize
}

func (o *V2TransactionsCursorResponseCursor) GetPrevious() *string {
	if o == nil {
		return nil
	}
	return o.Previous
}

type V2TransactionsCursorResponse struct {
	Cursor V2TransactionsCursorResponseCursor `json:"cursor"`
}

func (o *V2TransactionsCursorResponse) GetCursor() V2TransactionsCursorResponseCursor {
	if o == nil {
		return V2TransactionsCursorResponseCursor{}
	}
	return o.Cursor
}