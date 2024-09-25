// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type ScriptResponse struct {
	Details      *string      `json:"details,omitempty"`
	ErrorCode    *ErrorsEnum  `json:"errorCode,omitempty"`
	ErrorMessage *string      `json:"errorMessage,omitempty"`
	Transaction  *Transaction `json:"transaction,omitempty"`
}

func (o *ScriptResponse) GetDetails() *string {
	if o == nil {
		return nil
	}
	return o.Details
}

func (o *ScriptResponse) GetErrorCode() *ErrorsEnum {
	if o == nil {
		return nil
	}
	return o.ErrorCode
}

func (o *ScriptResponse) GetErrorMessage() *string {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
}

func (o *ScriptResponse) GetTransaction() *Transaction {
	if o == nil {
		return nil
	}
	return o.Transaction
}
