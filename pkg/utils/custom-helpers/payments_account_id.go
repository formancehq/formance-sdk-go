// Package custom helpers provides utilities to ease usage of Formance API.
// This file is NOT auto-generated and can be freely modified.
package custom_helpers

import (
	"encoding/base64"
	"fmt"

	"github.com/gibson042/canonicaljson-go"
	"github.com/google/uuid"
)

// AccountID represents a Formance Payments account identifier.
type AccountID struct {
	Reference   string      `json:"Reference"`
	ConnectorID ConnectorID `json:"ConnectorID"`
}

// BuildAccountID creates an account ID string from a connector ID string and a reference.
// The reference is the account identifier from the connector's provider.
func BuildAccountID(connectorID string, reference string) (string, error) {
	cid, err := connectorIDFromString(connectorID)
	if err != nil {
		return "", fmt.Errorf("invalid connector ID: %w", err)
	}

	aid := AccountID{
		Reference:   reference,
		ConnectorID: cid,
	}

	return aid.string(), nil
}

func (aid *AccountID) string() string {
	if aid == nil || aid.Reference == "" {
		return ""
	}

	data, err := canonicaljson.Marshal(aid)
	if err != nil {
		panic(err)
	}

	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}

type ConnectorID struct {
	Reference uuid.UUID `json:"Reference"`
	Provider  string    `json:"Provider"`
}

func connectorIDFromString(value string) (ConnectorID, error) {
	data, err := base64.URLEncoding.WithPadding(base64.NoPadding).DecodeString(value)
	if err != nil {
		return ConnectorID{}, fmt.Errorf("failed to decode connector ID: %w", err)
	}

	var ret ConnectorID
	err = canonicaljson.Unmarshal(data, &ret)
	if err != nil {
		return ConnectorID{}, fmt.Errorf("failed to unmarshal connector ID: %w", err)
	}

	return ret, nil
}
