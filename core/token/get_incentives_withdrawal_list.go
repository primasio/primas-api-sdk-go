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

type NodeAccountWithdraw struct {
	Id                string          `json:"id"`                 // Withdrawal id.
	Created           int             `json:"created"`            // Withdrawal created time. Unix timestamp.
	Updated           int             `json:"updated"`            // Withdrawal updated time.
	Amount            decimal.Decimal `json:"amount"`             // Withdrawal amount.
	BalanceAfter      decimal.Decimal `json:"balance_after"`      // Balance after withdrawal.
	NodeFee           decimal.Decimal `json:"node_fee"`           // Node charged withdrawal fee.
	Status            string          `json:"status"`             // Withdrawal status. "pending", "done" or "cancelled".
	TransactionHash   string          `json:"transaction_hash"`   // Withdrawal transaction hash.
	TransactionStatus string          `json:"transaction_status"` // Withdrawal transaction status. "pending", "done" or "failed".
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

	url := config.CONST_Server + `/accounts/` + account_id + `/tokens/incentives/withdrawal`
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
