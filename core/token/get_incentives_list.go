package token

import (
	"encoding/json"
	"errors"

	"github.com/primasio/primas-api-sdk-go/config"
	"github.com/primasio/primas-api-sdk-go/core"
	"github.com/primasio/primas-api-sdk-go/core/tool"
	"github.com/shopspring/decimal"
)

type IncentiveByDate struct {
	Date      string          `json:"date"`      // Incentives date. Unix timestamp.
	Total     decimal.Decimal `json:"total"`     // Total incentives get for today.
	Originals decimal.Decimal `json:"originals"` // Incentives get from original.
	Likes     decimal.Decimal `json:"likes"`     // Incentives get from likes.
	Comments  decimal.Decimal `json:"comments"`  // Incentives get from comments.
	Shares    decimal.Decimal `json:"shares"`    // Incentives get from shares.
	Groups    decimal.Decimal `json:"groups"`    // Incentives get from group management.
}

type IncentivesListResponse struct {
	core.Response
	Data []IncentiveByDate `json:"data"`
}

func GetIncentivesList(account_id string, startDate, endDate, page, pageSize int) (*IncentivesListResponse, error) {
	if account_id == "" {
		return nil, errors.New("account_id is empty")
	}

	url := config.CONST_Server + `/accounts/` + account_id + `/tokens/incentives`
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

	response, err := tool.Http_Get(url, queryParams)
	if err != nil {
		return nil, err
	}

	var responseObj IncentivesListResponse
	err = json.Unmarshal(response, &responseObj)
	if err != nil {
		return nil, err
	}

	if responseObj.ResultCode != core.CONST_ResultCode_Success {
		return nil, errors.New(responseObj.ResultMsg)
	}

	return &responseObj, nil
}
