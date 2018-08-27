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

const (
	CONST_Token_PreLock_Type_Lock   = "lock"
	CONST_Token_PreLock_Type_unlock = "unlock"
)

type TokenPreLock struct {
	ID          int             `json:"id"`
	CreatedAt   int             `json:"created_at"`
	UserAddress string          `json:"user_address"`
	NodeAddress string          `json:"node_address"`
	Amount      decimal.Decimal `json:"amount"`
	NodeFee     decimal.Decimal `json:"node_fee"`
	LockType    int             `json:"lock_type"`
	OrderID     string          `json:"order_id"`
	TxStatus    int             `json:"tx_status"`
	TxHash      string          `json:"tx_hash"`
	AccountId   string          `json:"account_id"`
}

type TokenPreLockListResponse struct {
	core.Response
	Data []TokenPreLock `json:"data"`
}

func GetTokenPreLockList(account_id string, startDate, endDate, page, pageSize int, qtype string) (*TokenPreLockListResponse, error) {
	if account_id == "" {
		return nil, errors.New("account_id is empty")
	}

	if qtype != "" {
		if (qtype != CONST_Token_PreLock_Type_Lock) && (qtype != CONST_Token_PreLock_Type_unlock) {
			return nil, errors.New("type is error")
		}
	}

	url := config.Gogal_Server + `/accounts/` + account_id + `/tokens/pre_locks`
	queryParams := make(map[string]interface{}, 0)
	if startDate > 0 {
		queryParams["start_date"] = startDate
	}
	if endDate > 0 {
		queryParams["end_date"] = endDate
	}
	if page > 0 {
		queryParams["page"] = page
	}
	if pageSize > 0 {
		queryParams["page_size"] = pageSize
	}
	if qtype != "" {
		queryParams["type"] = qtype
	}

	response, err := tool.Http_Get(url, queryParams)
	if err != nil {
		return nil, err
	}

	//log.Println("response:", string(response))

	var responseObj TokenPreLockListResponse
	err = json.Unmarshal(response, &responseObj)
	if err != nil {
		return nil, err
	}

	if responseObj.ResultCode != core.CONST_ResultCode_Success {
		return nil, errors.New(responseObj.ResultMsg)
	}

	return &responseObj, nil
}
