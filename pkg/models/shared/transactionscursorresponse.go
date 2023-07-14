// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type TransactionsCursorResponseCursor struct {
	Data     []Transaction `json:"data"`
	HasMore  bool          `json:"hasMore"`
	Next     *string       `json:"next,omitempty"`
	PageSize int64         `json:"pageSize"`
	Previous *string       `json:"previous,omitempty"`
}

// TransactionsCursorResponse - OK
type TransactionsCursorResponse struct {
	Cursor TransactionsCursorResponseCursor `json:"cursor"`
}