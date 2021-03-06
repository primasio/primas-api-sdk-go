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

	url := config.Gogal_Server + `/accounts/` + account_id + `/tokens/incentives`
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
