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

package content_interaction

import (
	"encoding/json"
	"errors"

	"github.com/primasio/primas-api-sdk-go/config"
	"github.com/primasio/primas-api-sdk-go/core"
	"github.com/primasio/primas-api-sdk-go/core/dtcp/dtcpv1"
	"github.com/primasio/primas-api-sdk-go/core/tool"
)

type PostCommentGroupshareResult struct {
	Id  string `json:"id"`  // Comment id.
	Dna string `json:"dna"` // Comment dna.
}

type PostCommentGroupshareResponse struct {
	core.Response
	Data *PostCommentGroupshareResult `json:"data"`
}

func PostCommentOfGroupshare_SignatureStr(account_id, share_id, sub_account_id, sub_account_name string,
	hp, created int, parent_comment_id, content string) (string, *dtcpv1.ContentCommentPost, error) {
	newPostCommentCreator := dtcpv1.ContentCommentPostCreator{
		AccountId: account_id,
	}
	if sub_account_id != "" {
		newPostCommentCreator.SubAccountId = sub_account_id
	}
	if sub_account_name != "" {
		newPostCommentCreator.SubAccountName = sub_account_name
	}

	newPostCommentExtra := dtcpv1.ContentCommentPostExtra{}
	if content != "" {
		newPostCommentExtra.Content = content
		newPostCommentExtra.ContentHash = dtcpv1.HashValue([]byte(content))
	}
	if parent_comment_id != "" {
		newPostCommentExtra.ParentCommentId = parent_comment_id
	}

	newCommentGroupshare := dtcpv1.NewContentCommentPost()
	newCommentGroupshare.SrcId = account_id
	newCommentGroupshare.DestId = share_id
	newCommentGroupshare.Hp = hp
	newCommentGroupshare.Creator = &newPostCommentCreator
	newCommentGroupshare.Created = created
	newCommentGroupshare.Extra = &newPostCommentExtra

	err := postCommentOfGroupshare_check(newCommentGroupshare)
	if err != nil {
		return "", nil, err
	}

	sigSoure, err := tool.StructToSignature(newCommentGroupshare)
	if err != nil {
		return "", nil, err
	}

	return sigSoure, newCommentGroupshare, nil
}

func PostCommentOfGroupshare(signature string, preObj *dtcpv1.ContentCommentPost) (*PostCommentGroupshareResponse, error) {
	if signature == "" {
		return nil, errors.New("param signature is empty")
	}

	err := postCommentOfGroupshare_check(preObj)
	if err != nil {
		return nil, err
	}

	preObj.Signature = signature

	url := config.Gogal_Server + `/shares/` + preObj.DestId + `/comments`

	requestBody, err := json.Marshal(preObj)
	if err != nil {
		return nil, err
	}

	response, err := tool.Http_Post(url, string(requestBody))
	if err != nil {
		return nil, err
	}

	var responseObj PostCommentGroupshareResponse
	err = json.Unmarshal(response, &responseObj)
	if err != nil {
		return nil, err
	}

	if responseObj.ResultCode != core.CONST_ResultCode_Success {
		return nil, errors.New(responseObj.ResultMsg)
	}

	return &responseObj, nil
}

func postCommentOfGroupshare_check(value *dtcpv1.ContentCommentPost) error {
	if value.SrcId == "" {
		return errors.New("parameter account_id error")
	}

	if value.DestId == "" {
		return errors.New("parameter share_id error")
	}

	if value.Created <= 0 {
		return errors.New("parameter created error")
	}

	if value.Creator == nil {
		return errors.New("parameter creator is nil")
	}

	if value.Creator.AccountId == "" {
		return errors.New("parameter creator error")
	}

	if value.Extra == nil {
		return errors.New("parameter extra is nil")
	}

	if value.Extra.Content == "" {
		return errors.New("parameter content is empty")
	}

	if value.Extra.ContentHash == "" {
		return errors.New("parameter content_hash is empty")
	}

	return nil
}
