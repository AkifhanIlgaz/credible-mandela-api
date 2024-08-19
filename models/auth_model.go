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
	Signature       string `json:"signature" binding:"required"`
}

func (form RegisterForm) Validate() error {
	if form.Password != form.ConfirmPassword {
		return fmt.Errorf("validate register form: password does not match")
	}

	if !common.IsHexAddress(form.Address) {
		return fmt.Errorf("validate register form: %s is not valid ethereum address", form.Address)
	}

	msg, err := json.Marshal(form)
	if err != nil {
		return fmt.Errorf("validate register form: %w", err)
	}

	err = crypto.VerifySignature(msg, form.Signature, form.Address)
	if err != nil {
		return fmt.Errorf("validate register form: %w", err)
	}

	return nil
}

func (form RegisterForm) ToUser() (User, error) {
	passwordHash, err := crypto.HashPassword(form.Password)
	if err != nil {
		return User{}, fmt.Errorf("convert register form to user: %w", err)
	}

	return User{
		Username:     form.Username,
		PasswordHash: passwordHash,
		Address:      form.Address,
		Roles:        "",
	}, nil
}

type RegisterResponse struct {
	Uid          string `json:"uid"`
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type LoginForm struct {
	Username string `json:"username" binding:"required,min=1"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginResponse struct {
	Uid          string `json:"uid"`
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
