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

package token

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"math/big"

	"github.com/primasio/go-ethereum/common"
	"github.com/primasio/go-ethereum/crypto"
	"github.com/primasio/primas-api-sdk-go/config"
	"github.com/primasio/primas-api-sdk-go/core"
	"github.com/primasio/primas-api-sdk-go/core/tool"
	"github.com/shopspring/decimal"
)

type PreLockTokensRequest struct {
	Transaction string `json:"transaction"`
}

type SignPreLock struct {
	Amount    decimal.Decimal `json:"amount"`    // Pre lock amount
	Nonce     string          `json:"nonce"`     // User operator nonce id
	Signature string          `json:"signature"` // User signature
}

type PreLockTokenId struct {
	Id string
}

type PreLockTokenResponse struct {
	core.Response
	Data *PreLockTokenId `json:"data"`
}

func PostPreLockTokens(user_address string, account_id string, amount decimal.Decimal, nonce string) (*PreLockTokenResponse, error) {
	if account_id == "" {
		return nil, errors.New("account_id is empty")
	}

	if amount.Cmp(decimal.Zero) <= 0 {
		return nil, errors.New("amount value less than zero")
	}

	if nonce == "" {
		return nil, errors.New("noce is empty")
	}

	signPreLock := SignPreLock{
		Amount:    amount,
		Nonce:     nonce,
		Signature: "",
	}

	tmpAmount := amount.Coefficient()
	tmpAmount = tmpAmount.Mul(tmpAmount, big.NewInt(1000000000000000000))

	amountBytes := tmpAmount.Bytes()
	result := common.LeftPadBytes(amountBytes, 32)

	msg1 := common.HexToAddress(user_address).Bytes()
	msg2 := []byte(nonce)
	msg3 := append(msg1[:], result...)
	msg4 := append(msg3, msg2...)
	msgBytes := crypto.Keccak256(msg4)
	sigBytes, _ := tool.GetClientKeystore().SignHash(*tool.GetClientAccount(), msgBytes)
	signMsg := hex.EncodeToString(sigBytes)
	signPreLock.Signature = signMsg

	transaction, err := json.Marshal(signPreLock)
	if err != nil {
		return nil, err
	}

	preLockTokensRequest := PreLockTokensRequest{
		Transaction: string(transaction),
	}

	url := config.Gogal_Server + `/accounts/` + account_id + `/tokens/pre_locks`

	requestBody, err := json.Marshal(preLockTokensRequest)
	if err != nil {
		return nil, err
	}

	response, err := tool.Http_Post(url, string(requestBody))
	if err != nil {
		return nil, err
	}

	//log.Println("response:", string(response))

	var responseObj PreLockTokenResponse
	err = json.Unmarshal(response, &responseObj)
	if err != nil {
		return nil, err
	}

	if responseObj.ResultCode != core.CONST_ResultCode_Success {
		return nil, errors.New(responseObj.ResultMsg)
	}

	return &responseObj, nil
}
