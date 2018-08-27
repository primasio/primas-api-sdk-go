/*
 * Copyright 2018 Primas Lab Foundation
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *       http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

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
