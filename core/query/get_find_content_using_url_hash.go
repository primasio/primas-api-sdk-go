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

package query

import (
	"encoding/json"
	"errors"

	"github.com/primasio/primas-api-sdk-go/config"
	"github.com/primasio/primas-api-sdk-go/core"
	"github.com/primasio/primas-api-sdk-go/core/dtcp/dtcpv1"
	"github.com/primasio/primas-api-sdk-go/core/tool"
)

type FindContentUsingUrlResponse struct {
	core.Response
	Data []dtcpv1.ContentGet `json:"data"`
}

func FindContentUsingUrlOrHash(urlParam, hashParam string) (*FindContentUsingUrlResponse, error) {
	if urlParam == "" && hashParam == "" {
		return nil, errors.New("url and hash are all empty")
	}

	queryParams := make(map[string]interface{}, 0)

	if urlParam != "" {
		queryParams["url"] = urlParam
	}
	if hashParam != "" {
		queryParams["hash"] = hashParam
	}

	url := config.Gogal_Server + `/query/content`

	response, err := tool.Http_Get(url, queryParams)
	if err != nil {
		return nil, err
	}

	var responseObj FindContentUsingUrlResponse
	err = json.Unmarshal(response, &responseObj)
	if err != nil {
		return nil, err
	}

	if responseObj.ResultCode != core.CONST_ResultCode_Success {
		return nil, errors.New(responseObj.ResultMsg)
	}

	return &responseObj, nil
}
