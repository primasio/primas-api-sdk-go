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
	"encoding/json"
	"errors"

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

func PostPreLockTokens(account_id string, amount decimal.Decimal, nonce string) (*PreLockTokenResponse, error) {
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

	sigSoure, err := tool.StructToSignature(signPreLock)
	if err != nil {
		return nil, err
	}
	privateKey := tool.GetClientPrivateKey()

	signature, err := tool.Sign([]byte(sigSoure), privateKey)
	if err != nil {
		return nil, err
	}
	signPreLock.Signature = signature

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
