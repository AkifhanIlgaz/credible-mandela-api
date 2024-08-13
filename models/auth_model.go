package models

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

type RegisterForm struct {
	Username        string `json:"username" binding:"required"`
	Password        string `json:"password" binding:"required"`
	ConfirmPassword string `json:"confirmPassword" binding:"required"`
	Address         string `json:"address" binding:"required"`
}

type RegisterFormWithSignature struct {
	RegisterForm
	Signature string `json:"signature" binding:"required"`
}

func (rf RegisterFormWithSignature) Validate() error {
	if rf.Password != rf.ConfirmPassword {
		return fmt.Errorf("password does not match")
	}

	if !common.IsHexAddress(rf.Address) {
		return fmt.Errorf("%s is not valid ethereum address", rf.Address)
	}

	return nil
}
