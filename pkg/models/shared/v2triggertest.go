// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type Filter struct {
	Error *string `json:"error,omitempty"`
	Match *bool   `json:"match,omitempty"`
}

func (o *Filter) GetError() *string {
	if o == nil {
		return nil
	}
	return o.Error
}

func (o *Filter) GetMatch() *bool {
	if o == nil {
		return nil
	}
	return o.Match
}

type Variables struct {
	Error *string `json:"error,omitempty"`
	Value *string `json:"value,omitempty"`
}

func (o *Variables) GetError() *string {
	if o == nil {
		return nil
	}
	return o.Error
}

func (o *Variables) GetValue() *string {
	if o == nil {
		return nil
	}
	return o.Value
}

type V2TriggerTest struct {
	Filter    *Filter              `json:"filter,omitempty"`
	Variables map[string]Variables `json:"variables,omitempty"`
}

func (o *V2TriggerTest) GetFilter() *Filter {
	if o == nil {
		return nil
	}
	return o.Filter
}

func (o *V2TriggerTest) GetVariables() map[string]Variables {
	if o == nil {
		return nil
	}
	return o.Variables
}