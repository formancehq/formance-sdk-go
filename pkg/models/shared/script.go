// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type Script struct {
	Metadata map[string]any `json:"metadata,omitempty"`
	Plain    string         `json:"plain"`
	// Reference to attach to the generated transaction
	Reference *string        `json:"reference,omitempty"`
	Vars      map[string]any `json:"vars,omitempty"`
}

func (o *Script) GetMetadata() map[string]any {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *Script) GetPlain() string {
	if o == nil {
		return ""
	}
	return o.Plain
}

func (o *Script) GetReference() *string {
	if o == nil {
		return nil
	}
	return o.Reference
}

func (o *Script) GetVars() map[string]any {
	if o == nil {
		return nil
	}
	return o.Vars
}
