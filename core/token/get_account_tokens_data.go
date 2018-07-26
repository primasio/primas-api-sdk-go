package token

import (
	"encoding/json"
	"errors"

	"github.com/primasio/primas-api-sdk-go/config"
	"github.com/primasio/primas-api-sdk-go/core"
	"github.com/primasio/primas-api-sdk-go/core/tool"
	"github.com/shopspring/decimal"
)

type AccountTokensData struct {
	Balance          decimal.Decimal `json:"balance"`            // Token balance.
	PreLockAll       decimal.Decimal `json:"pre_lock_all"`       // Total pre-locked amount.
	PreLockAvailable decimal.Decimal `json:"pre_lock_available"` // Remaining pre-locked amount.
	IncentivesAll    decimal.Decimal `json:"incentives_all"`     // Total incentives.
	IncentivesLocked decimal.Decimal `json:"incentives_locked"`  // Incentives locked on this node.
	IncentivesOnNode decimal.Decimal `json:"incentives_on_node"` // Amount in the node's incentives pool.
}

type AccountTokensResponse struct {
	core.Response
	Data *AccountTokensData `json:"data"`
}

func GetAccountTokensData(account_id string) (*AccountTokensResponse, error) {
	if account_id == "" {
		return nil, errors.New("account_id is empty")
	}

	url := config.CONST_Server + `/accounts/` + account_id + `/tokens`
	response, err := tool.Http_Get(url, nil)
	if err != nil {
		return nil, err
	}

	//log.Println("response:", string(response))

	var responseObj AccountTokensResponse
	err = json.Unmarshal(response, &responseObj)
	if err != nil {
		return nil, err
	}

	if responseObj.ResultCode != core.CONST_ResultCode_Success {
		return nil, errors.New(responseObj.ResultMsg)
	}

	return &responseObj, nil
}
