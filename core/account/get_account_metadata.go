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

type GetAccountMetadata struct {
	Id            string                     `json:"id"`                 // Account id.
	Address       string                     `json:"address"`            // Account address.
	Title         string                     `json:"title"`              // Account name.
	Abstract      string                     `json:"abstract,omitempty"` // Description.
	Avatar        string                     `json:"avatar,omitempty"`   // An image id used for avatar.
	Creator       *GetAccountMetadataCreator `json:"creator"`            // Creator of the sub account.
	Created       int                        `json:"created"`            // Account creation time. Unix timestamp.
	Updated       int                        `json:"updated"`            // Account last updating time. Unix timestamp.
	Extra         *GetAccountMetadataExtra   `json:"extra,omitempty"`    // Extra metadata.
	Signature     string                     `json:"signature"`          // Metadata signature.
	Dna           string                     `json:"dna"`                // DNA of the account.
	Hp            int                        `json:"hp"`                 // available hp value
	TotalHp       int                        `json:"total_hp"`           // total hp value
	Credits       int                        `json:"credits"`            // Current credits.
	TransactionId string                     `json:"transaction_id"`     // Latest transaction id.
}

type GetAccountMetadataCreator struct {
	AccountId   string `json:"account_id"`   // Root account id.
	AccountName string `json:"account_name"` // Root account name.
}

type GetAccountMetadataExtra struct {
	Hash string `json:"hash"` // In the case of proof of existence of secret data. The hash can be filled in this field.
}

type GetAccountMetadataResponse struct {
	core.Response
	Data *GetAccountMetadata `json:"data"`
}

func GetAccountTokenMetadata(account_id string) (*GetAccountMetadataResponse, error) {
	if account_id == "" {
		return nil, errors.New("account_id is empty")
	}

	url := config.Gogal_Server + `/accounts/` + account_id + `/metadata`
	response, err := tool.Http_Get(url, nil)
	if err != nil {
		return nil, err
	}

	var responseObj GetAccountMetadataResponse
	err = json.Unmarshal(response, &responseObj)
	if err != nil {
		return nil, err
	}

	if responseObj.ResultCode != core.CONST_ResultCode_Success {
		return nil, errors.New(responseObj.ResultMsg)
	}

	return &responseObj, nil
}

func GetSubAccountTokenMetadata(account_id, sub_account_id string) (*GetAccountMetadataResponse, error) {
	if account_id == "" {
		return nil, errors.New("account_id is empty")
	}

	if sub_account_id == "" {
		return nil, errors.New("sub_account_id is empty")
	}

	url := config.Gogal_Server + `/accounts/` + account_id + `/sub/` + sub_account_id + `/metadata`
	response, err := tool.Http_Get(url, nil)
	if err != nil {
		return nil, err
	}

	var responseObj GetAccountMetadataResponse
	err = json.Unmarshal(response, &responseObj)
	if err != nil {
		return nil, err
	}

	if responseObj.ResultCode != core.CONST_ResultCode_Success {
		return nil, errors.New(responseObj.ResultMsg)
	}

	return &responseObj, nil
}
