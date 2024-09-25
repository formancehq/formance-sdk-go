// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/formancehq/formance-sdk-go/v3/pkg/utils"
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
		if err := utils.UnmarshalJSON(data, &ledgerAccountSubject, "", true, false); err != nil {
			return fmt.Errorf("could not unmarshal `%s` into expected (Type == ACCOUNT) type LedgerAccountSubject within Subject: %w", string(data), err)
		}

		u.LedgerAccountSubject = ledgerAccountSubject
		u.Type = SubjectTypeAccount
		return nil
	case "WALLET":
		walletSubject := new(WalletSubject)
		if err := utils.UnmarshalJSON(data, &walletSubject, "", true, false); err != nil {
			return fmt.Errorf("could not unmarshal `%s` into expected (Type == WALLET) type WalletSubject within Subject: %w", string(data), err)
		}

		u.WalletSubject = walletSubject
		u.Type = SubjectTypeWallet
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for Subject", string(data))
}

func (u Subject) MarshalJSON() ([]byte, error) {
	if u.LedgerAccountSubject != nil {
		return utils.MarshalJSON(u.LedgerAccountSubject, "", true)
	}

	if u.WalletSubject != nil {
		return utils.MarshalJSON(u.WalletSubject, "", true)
	}

	return nil, errors.New("could not marshal union type Subject: all fields are null")
}
