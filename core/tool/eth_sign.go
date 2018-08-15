package tool

import (
	"crypto/ecdsa"
	"encoding/hex"

	"github.com/primasio/go-ethereum/crypto"
	"github.com/primasio/go-ethereum/crypto/secp256k1"
)

func Sign(data []byte, privateKey string) (string, error) {
	locPrivateKey, err := crypto.HexToECDSA(privateKey)

	sigValue := crypto.Keccak256(data)
	sigBytes, err := crypto.Sign(sigValue, locPrivateKey)

	return hex.EncodeToString(sigBytes), err
}

func Verify(data []byte, signature, public string) bool {
	msgBytes := crypto.Keccak256(data)

	sigBytes, err := hex.DecodeString(signature)
	if err != nil {
		return false
	}

	publicKey, err := secp256k1.RecoverPubkey(msgBytes, sigBytes)
	if err != nil {
		return false
	}

	locPublicKey := crypto.ToECDSAPub(publicKey)
	locAddress := crypto.PubkeyToAddress(*locPublicKey)

	return locAddress.Hex() == public
}

func SignByEth(data []byte, privateKey *ecdsa.PrivateKey) (string, error) {
	sigStr := crypto.Keccak256(data)
	sigBytes, err := crypto.Sign(sigStr, privateKey)

	return hex.EncodeToString(sigBytes), err
}
