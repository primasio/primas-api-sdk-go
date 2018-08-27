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

package account

import (
	"encoding/json"
	"errors"

	"github.com/primasio/primas-api-sdk-go/config"
	"github.com/primasio/primas-api-sdk-go/core"
	"github.com/primasio/primas-api-sdk-go/core/tool"
)

type AccountCredists struct {
	UserAccountId string `json:"user_account_id"`
	Score         int    `json:"score"`
}

type AccountCredistsResponse struct {
	core.Response
	Data *AccountCredists `json:"data"`
}

func GetAccountCreditsList(account_id string) (*AccountCredistsResponse, error) {
	if account_id == "" {
		return nil, errors.New("param account_id is empty")
	}

	url := config.Gogal_Server + `/accounts/` + account_id + `/credits`

	response, err := tool.Http_Get(url, nil)
	if err != nil {
		return nil, err
	}

	var responseObj AccountCredistsResponse
	err = json.Unmarshal(response, &responseObj)
	if err != nil {
		return nil, err
	}

	if responseObj.ResultCode != core.CONST_ResultCode_Success {
		return nil, errors.New(responseObj.ResultMsg)
	}

	return &responseObj, nil
}

type SubAccountCredists struct {
	UserAccountId string `json:"user_account_id"`
	SubAccountId  string `json:"sub_account_id"`
	Score         int    `json:"score"`
}

type SubAccountCredistsResponse struct {
	core.Response
	Data *SubAccountCredists `json:"data"`
}

func GetSubAccountCreditsList(account_id, sub_account_id string) (*SubAccountCredistsResponse, error) {
	if account_id == "" {
		return nil, errors.New("param account_id is empty")
	}

	if sub_account_id == "" {
		return nil, errors.New("param sub_account_id is empty")
	}

	url := config.Gogal_Server + `/accounts/` + account_id + `/sub/` + sub_account_id + `/credits`

	response, err := tool.Http_Get(url, nil)
	if err != nil {
		return nil, err
	}

	var responseObj SubAccountCredistsResponse
	err = json.Unmarshal(response, &responseObj)
	if err != nil {
		return nil, err
	}

	if responseObj.ResultCode != core.CONST_ResultCode_Success {
		return nil, errors.New(responseObj.ResultMsg)
	}

	return &responseObj, nil
}
