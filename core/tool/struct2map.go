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
	"bytes"
	"encoding/hex"
	"encoding/json"
	"strings"

	"github.com/primasio/go-ethereum/crypto"
	"github.com/primasio/go-ethereum/crypto/secp256k1"
	"github.com/shopspring/decimal"
)

const (
	loc_fieldnamej_signature = "signature"
)

func init() {
	decimal.MarshalJSONWithoutQuotes = true
}

func StructToSignature(obj interface{}) (string, error) {
	beforeValue, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}

	midResultMap := make(map[string]interface{}, 0)

	jen := json.NewDecoder(bytes.NewBuffer(beforeValue))
	jen.UseNumber()
	jen.Decode(&midResultMap)

	delete(midResultMap, loc_fieldnamej_signature)

	eValue, err := jsonMarshal(midResultMap)
	if err != nil {
		return "", err
	}

	return eValue, nil
}

func StructToSignatureByExclude(obj interface{}, excludes ...string) (string, error) {
	beforeValue, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}

	midResultMap := make(map[string]interface{}, 0)

	jen := json.NewDecoder(bytes.NewBuffer(beforeValue))
	jen.UseNumber()
	jen.Decode(&midResultMap)

	delete(midResultMap, loc_fieldnamej_signature)

	for _, item := range excludes {
		delete(midResultMap, item)
	}

	eValue, err := jsonMarshal(midResultMap)
	if err != nil {
		return "", err
	}

	return string(eValue), nil
}

// The default behavior is to escape &, <, and > to \u0026, \u003c, and \u003e
// to avoid certain safety problems that can arise when embedding JSON in HTML.
func jsonMarshal(value map[string]interface{}) (string, error) {
	buffer := &bytes.Buffer{}

	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(value)
	if err != nil {
		return "", err
	}

	result := strings.Trim(buffer.String(), "\n")

	return result, nil
}

func RecoverPublickeyByEth(msg, sig string) (string, error) {
	msgBytes := crypto.Keccak256([]byte(msg))

	sigBytes, err := hex.DecodeString(sig)
	if err != nil {
		return "", err
	}

	publicKey, err := secp256k1.RecoverPubkey(msgBytes, sigBytes)
	if err != nil {
		return "", err
	}

	srcPublicKey := crypto.ToECDSAPub(publicKey)
	srcAddress := crypto.PubkeyToAddress(*srcPublicKey)

	return srcAddress.Hex(), nil
}
