package models

import (
	"encoding/json"
	"fmt"

	"github.com/AkifhanIlgaz/credible-mandela-api/utils/crypto"
	"github.com/ethereum/go-ethereum/common"
)

type RegisterForm struct {
	Username        string `json:"username" binding:"required,min=1"`
	Password        string `json:"password" binding:"required,min=6"`
	ConfirmPassword string `json:"confirmPassword" binding:"required,min=6"`
	Address         string `json:"address" binding:"required"`
}

type RegisterFormWithSignature struct {
	RegisterForm
	Signature string `json:"signature" binding:"required"`
}

func (rf RegisterFormWithSignature) Validate() error {
	if rf.Password != rf.ConfirmPassword {
		return fmt.Errorf("validate register form: password does not match")
	}

	if !common.IsHexAddress(rf.Address) {
		return fmt.Errorf("validate register form: %s is not valid ethereum address", rf.Address)
	}

	msg, err := json.Marshal(&rf.RegisterForm)
	if err != nil {
		return fmt.Errorf("validate register form: %w", err)
	}

	err = crypto.VerifySignature(msg, rf.Signature, rf.Address)
	if err != nil {

		return fmt.Errorf("validate register form: %w", err)
	}

	return nil
}

func (rf RegisterForm) ToUser() (User, error) {
	passwordHash, err := crypto.HashPassword(rf.Password)
	if err != nil {
		return User{}, fmt.Errorf("convert register form to user: %w", err)
	}

	return User{
		Username:     rf.Username,
		PasswordHash: passwordHash,
		Address:      rf.Address,
		Roles:        "",
	}, nil
}
