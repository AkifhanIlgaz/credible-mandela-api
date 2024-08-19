package crypto_test

import (
	"encoding/json"
	"testing"

	"github.com/AkifhanIlgaz/credible-mandela-api/models"
	cr "github.com/AkifhanIlgaz/credible-mandela-api/utils/crypto"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
)

func TestVerifySignature(t *testing.T) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		t.Fatalf("Failed to generate key pair: %v", err)
	}

	publicKey := privateKey.PublicKey
	address := crypto.PubkeyToAddress(publicKey).Hex()

	data := models.RegisterForm{
		Username:        "testing",
		Password:        "123456",
		ConfirmPassword: "123456",
		Address:         address,
	}

	msg, err := json.Marshal(&data)
	if err != nil {
		t.Error(err)
	}
	msgHash := accounts.TextHash(msg)

	sigBytes, err := crypto.Sign(msgHash, privateKey)
	if err != nil {
		t.Error(err)
	}
	signature := hexutil.Encode(sigBytes)

	err = cr.VerifySignature(msg, signature, address)
	assert.NoError(t, err, "The signature should be valid")

	invalidAddress := crypto.PubkeyToAddress(publicKey).String() + "1"
	err = cr.VerifySignature(msg, signature, invalidAddress)
	assert.Error(t, err, "The signature should be invalid for the incorrect address")

	invalidSignature := signature[:len(signature)-2] + "a0"
	err = cr.VerifySignature(msg, invalidSignature, address)
	assert.Error(t, err, "The signature should be invalid")
}
