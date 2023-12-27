// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type QueryRaw struct {
}

type Query struct {
	After    []string  `json:"after,omitempty"`
	Cursor   *string   `json:"cursor,omitempty"`
	Ledgers  []string  `json:"ledgers,omitempty"`
	PageSize *int64    `json:"pageSize,omitempty"`
	Policy   *string   `json:"policy,omitempty"`
	Raw      *QueryRaw `json:"raw,omitempty"`
	Sort     *string   `json:"sort,omitempty"`
	Target   *string   `json:"target,omitempty"`
	Terms    []string  `json:"terms,omitempty"`
}

func (o *Query) GetAfter() []string {
	if o == nil {
		return nil
	}
	return o.After
}

func (o *Query) GetCursor() *string {
	if o == nil {
		return nil
	}
	return o.Cursor
}

func (o *Query) GetLedgers() []string {
	if o == nil {
		return nil
	}
	return o.Ledgers
}

func (o *Query) GetPageSize() *int64 {
	if o == nil {
		return nil
	}
	return o.PageSize
}

func (o *Query) GetPolicy() *string {
	if o == nil {
		return nil
	}
	return o.Policy
}

func (o *Query) GetRaw() *QueryRaw {
	if o == nil {
		return nil
	}
	return o.Raw
}

func (o *Query) GetSort() *string {
	if o == nil {
		return nil
	}
	return o.Sort
}

func (o *Query) GetTarget() *string {
	if o == nil {
		return nil
	}
	return o.Target
}

func (o *Query) GetTerms() []string {
	if o == nil {
		return nil
	}
	return o.Terms
}
