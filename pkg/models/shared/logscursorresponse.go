// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package shared

type LogsCursorResponseCursor struct {
	Data     []Log   `json:"data"`
	HasMore  bool    `json:"hasMore"`
	Next     *string `json:"next,omitempty"`
	PageSize int64   `json:"pageSize"`
	Previous *string `json:"previous,omitempty"`
}

func (o *LogsCursorResponseCursor) GetData() []Log {
	if o == nil {
		return []Log{}
	}
	return o.Data
}

func (o *LogsCursorResponseCursor) GetHasMore() bool {
	if o == nil {
		return false
	}
	return o.HasMore
}

func (o *LogsCursorResponseCursor) GetNext() *string {
	if o == nil {
		return nil
	}
	return o.Next
}

func (o *LogsCursorResponseCursor) GetPageSize() int64 {
	if o == nil {
		return 0
	}
	return o.PageSize
}

func (o *LogsCursorResponseCursor) GetPrevious() *string {
	if o == nil {
		return nil
	}
	return o.Previous
}

type LogsCursorResponse struct {
	Cursor LogsCursorResponseCursor `json:"cursor"`
}

func (o *LogsCursorResponse) GetCursor() LogsCursorResponseCursor {
	if o == nil {
		return LogsCursorResponseCursor{}
	}
	return o.Cursor
}
