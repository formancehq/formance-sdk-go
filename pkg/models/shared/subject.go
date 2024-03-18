// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/formancehq/formance-sdk-go/v2/v2/pkg/utils"
)

type SubjectType string

const (
	SubjectTypeAccount SubjectType = "ACCOUNT"
	SubjectTypeWallet  SubjectType = "WALLET"
)

type Subject struct {
	LedgerAccountSubject *LedgerAccountSubject
	WalletSubject        *WalletSubject

	Type SubjectType
}

func CreateSubjectAccount(account LedgerAccountSubject) Subject {
	typ := SubjectTypeAccount

	typStr := string(typ)
	account.Type = typStr

	return Subject{
		LedgerAccountSubject: &account,
		Type:                 typ,
	}
}

func CreateSubjectWallet(wallet WalletSubject) Subject {
	typ := SubjectTypeWallet

	typStr := string(typ)
	wallet.Type = typStr

	return Subject{
		WalletSubject: &wallet,
		Type:          typ,
	}
}

func (u *Subject) UnmarshalJSON(data []byte) error {

	type discriminator struct {
		Type string `json:"type"`
	}

	dis := new(discriminator)
	if err := json.Unmarshal(data, &dis); err != nil {
		return fmt.Errorf("could not unmarshal discriminator: %w", err)
	}

	switch dis.Type {
	case "ACCOUNT":
		ledgerAccountSubject := new(LedgerAccountSubject)
		if err := utils.UnmarshalJSON(data, &ledgerAccountSubject, "", true, true); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.LedgerAccountSubject = ledgerAccountSubject
		u.Type = SubjectTypeAccount
		return nil
	case "WALLET":
		walletSubject := new(WalletSubject)
		if err := utils.UnmarshalJSON(data, &walletSubject, "", true, true); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.WalletSubject = walletSubject
		u.Type = SubjectTypeWallet
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u Subject) MarshalJSON() ([]byte, error) {
	if u.LedgerAccountSubject != nil {
		return utils.MarshalJSON(u.LedgerAccountSubject, "", true)
	}

	if u.WalletSubject != nil {
		return utils.MarshalJSON(u.WalletSubject, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}
