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

type Info struct {
	ID           uint   `json:"id"`
	CreatedAt    uint   `json:"created_at"`
	UserAddress  string `json:"user_address,omitempty"`
	Address      string `json:"address,omitempty"`
	Name         string `json:"name,omitempty"`
	Title        string `json:"title,omitempty"`
	Abstract     string `json:"abstract,omitempty"`
	Description  string `json:"description,omitempty"`
	License      string `json:"license,omitempty"`
	Extra        string `json:"extra,omitempty"`
	DNA          string `json:"dna,omitempty"`
	Status       string `json:"status,omitempty"`
	TxStatus     int    `json:"tx_status,omitempty"`
	MemberCount  uint   `json:"member_count,omitempty"`
	ArticleCount uint   `json:"article_count,omitempty"`
	LikeCount    uint   `json:"like_count,omitempty"`
	CommentCount uint   `json:"comment_count,omitempty"`
	ShareCount   uint   `json:"share_count,omitempty"`
	FilePath     string `json:"file_path,omitempty"`
}

type RawData struct {
	DataType     string             `json:"DataType"`
	ID           uint               `json:"Id,omitempty"`
	CreatedAt    uint               `json:"CreatedAt,omitempty"`
	UserAddress  string             `json:"UserAddress,omitempty"`
	Address      string             `json:"Address,omitempty"`
	Name         string             `json:"Name,omitempty"`
	Title        string             `json:"Title,omitempty"`
	Abstract     string             `json:"Abstract,omitempty"`
	Description  string             `json:"Description,omitempty"`
	License      string             `json:"License,omitempty"`
	Extra        string             `json:"Extra,omitempty"`
	DNA          string             `json:"DNA,omitempty"`
	Status       string             `json:"Status,omitempty"`
	TxStatus     int                `json:"TxStatus,omitempty"`
	MemberCount  uint               `json:"MemberCount,omitempty"`
	ArticleCount uint               `json:"ArticleCount,omitempty"`
	LikeCount    uint               `json:"LikeCount,omitempty"`
	CommentCount uint               `json:"CommentCount,omitempty"`
	ShareCount   uint               `json:"ShareCount,omitempty"`
	FilePath     string             `json:"FilePath,omitempty"`
	Highlight    []string           `json:"Highlight,omitempty"`
	Group        []dtcpv1.GroupGet  `json:"Group,omitempty"`
	Article      *dtcpv1.ContentGet `json:"Article,omitempty"`
	SmallImage   []string           `json:"SmallImgs,omitempty"`
}

type SearchResult struct {
	Total  int64 `json:"Total"`
	Offset int   `json:"Offset"`
	Data   []RawData
}

type GetAllResponse struct {
	core.Response
	Data *SearchResult `json:"data"`
}

func GetAll(page, page_size int, text, qtype, category string) (*GetAllResponse, error) {
	queryParams := make(map[string]interface{}, 0)

	if page >= 0 {
		queryParams["page"] = page
	}
	if page_size >= 0 {
		queryParams["page_size"] = page_size
	}
	if text != "" {
		queryParams["text"] = text
	}
	if qtype != "" {
		queryParams["type"] = qtype
	}
	if category != "" {
		queryParams["category"] = category
	}

	url := config.Gogal_Server + `/query`

	response, err := tool.Http_Get(url, queryParams)
	if err != nil {
		return nil, err
	}

	var responseObj GetAllResponse
	err = json.Unmarshal(response, &responseObj)
	if err != nil {
		return nil, err
	}

	if responseObj.ResultCode != core.CONST_ResultCode_Success {
		return nil, errors.New(responseObj.ResultMsg)
	}

	return &responseObj, nil
}
