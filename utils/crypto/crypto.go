package crypto

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/bcrypt"
)

func VerifySignature(message []byte, signature, from string) error {
	hashedMessage := accounts.TextHash(message)

	signatureBytes, err := hexutil.Decode(signature)
	if err != nil {
		return fmt.Errorf("error decoding signature: %v", err)
	}

	if len(signatureBytes) != crypto.SignatureLength {
		return fmt.Errorf("invalid signature length")
	}

	if signatureBytes[crypto.RecoveryIDOffset] == 27 || signatureBytes[crypto.RecoveryIDOffset] == 28 {
		signatureBytes[crypto.RecoveryIDOffset] -= 27 // Transform yellow paper V from 27/28 to 0/1
	}

	sigPublicKey, err := crypto.SigToPub(hashedMessage, signatureBytes)
	if err != nil {
		return fmt.Errorf("error recovering public key from signature: %w", err)
	}

	recoveredAddress := crypto.PubkeyToAddress(*sigPublicKey)
	if from != recoveredAddress.Hex() {
		return fmt.Errorf("signature verification failed: recovered address %s does not match provided address %s", recoveredAddress, from)
	}

	return nil
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("could not hash password: %w", err)
	}

	return string(hashedPassword), nil
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
