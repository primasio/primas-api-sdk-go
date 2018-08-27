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
	CONST_Incetives_Withdrawal_Status_Pending = "pending"
	CONST_Incetives_Withdrawal_Status_Done    = "done"
)

type TransactionObj struct {
	Transaction_id     string `json:"transaction_id,omitempty"`      // Transaction hash.
	BlockNumber        int    `json:"block_number,omitempty"`        // Block number of this transaction.
	BlockConfirmations int    `json:"block_confirmations,omitempty"` // Block confirmation time.
	EstimatedTime      int    `json:"estimated_time"`                // Estimated confirmation time. Unix timestamp.
	ConfirmedTime      int    `json:"confirmed_time,omitempty"`      // Confirmation time. Unix timestamp.
}

type NodeAccountWithdraw struct {
	Id           string          `json:"id"`            // Withdrawal id.
	Created      int             `json:"created"`       // Withdrawal created time. Unix timestamp.
	Updated      int             `json:"updated"`       // Withdrawal updated time.
	Amount       decimal.Decimal `json:"amount"`        // Withdrawal amount.
	BalanceAfter decimal.Decimal `json:"balance_after"` // Balance after withdrawal.
	NodeFee      decimal.Decimal `json:"node_fee"`      // Node charged withdrawal fee.
	Status       string          `json:"status"`        // Withdrawal status. "pending", "done" or "cancelled".
	Transaction  *TransactionObj `json:"transaction"`   // Withdrawal transaction object
}

type IncentivesWithdrawalListResponse struct {
	core.Response
	Data []NodeAccountWithdraw `json:"data"`
}

func GetIncentivesWithdrawalList(account_id string, startDate, endDate, page, pageSize int, status string) (*IncentivesWithdrawalListResponse, error) {
	if account_id == "" {
		return nil, errors.New("account_id is empty")
	}

	if status != "" {
		if (status != CONST_Incetives_Withdrawal_Status_Pending) && (status != CONST_Incetives_Withdrawal_Status_Done) {
			return nil, errors.New("status is error")
		}
	}

	url := config.Gogal_Server + `/accounts/` + account_id + `/tokens/incentives/withdrawal`
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
	if status != "" {
		queryParams["status"] = status
	}

	response, err := tool.Http_Get(url, queryParams)
	if err != nil {
		return nil, err
	}

	var responseObj IncentivesWithdrawalListResponse
	err = json.Unmarshal(response, &responseObj)
	if err != nil {
		return nil, err
	}

	if responseObj.ResultCode != core.CONST_ResultCode_Success {
		return nil, errors.New(responseObj.ResultMsg)
	}

	return &responseObj, nil
}
