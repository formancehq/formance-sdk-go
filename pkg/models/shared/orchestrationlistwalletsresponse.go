// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type OrchestrationListWalletsResponseCursor struct {
	Data     []Wallet `json:"data"`
	HasMore  *bool    `json:"hasMore,omitempty"`
	Next     *string  `json:"next,omitempty"`
	PageSize int64    `json:"pageSize"`
	Previous *string  `json:"previous,omitempty"`
}

func (o *OrchestrationListWalletsResponseCursor) GetData() []Wallet {
	if o == nil {
		return []Wallet{}
	}
	return o.Data
}

func (o *OrchestrationListWalletsResponseCursor) GetHasMore() *bool {
	if o == nil {
		return nil
	}
	return o.HasMore
}

func (o *OrchestrationListWalletsResponseCursor) GetNext() *string {
	if o == nil {
		return nil
	}
	return o.Next
}

func (o *OrchestrationListWalletsResponseCursor) GetPageSize() int64 {
	if o == nil {
		return 0
	}
	return o.PageSize
}

func (o *OrchestrationListWalletsResponseCursor) GetPrevious() *string {
	if o == nil {
		return nil
	}
	return o.Previous
}

type OrchestrationListWalletsResponse struct {
	Cursor OrchestrationListWalletsResponseCursor `json:"cursor"`
}

func (o *OrchestrationListWalletsResponse) GetCursor() OrchestrationListWalletsResponseCursor {
	if o == nil {
		return OrchestrationListWalletsResponseCursor{}
	}
	return o.Cursor
}
