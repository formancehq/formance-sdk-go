// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"time"
)

type PostTransactionScriptVars struct {
}

type PostTransactionScript struct {
	Plain string                     `json:"plain"`
	Vars  *PostTransactionScriptVars `json:"vars,omitempty"`
}

// PostTransaction - The request body must contain at least one of the following objects:
//   - `postings`: suitable for simple transactions
//   - `script`: enabling more complex transactions with Numscript
type PostTransaction struct {
	Metadata  map[string]interface{} `json:"metadata,omitempty"`
	Postings  []Posting              `json:"postings,omitempty"`
	Reference *string                `json:"reference,omitempty"`
	Script    *PostTransactionScript `json:"script,omitempty"`
	Timestamp *time.Time             `json:"timestamp,omitempty"`
}