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

package group

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/primasio/primas-api-sdk-go/config"
	"github.com/primasio/primas-api-sdk-go/core"
	"github.com/primasio/primas-api-sdk-go/core/dtcp/dtcpv1"
	"github.com/primasio/primas-api-sdk-go/core/tool"
)

type PostAddGroupMemberWhiteResult struct {
	Id  string `json:"id"`  // Group member whitelist id.
	Dna string `json:"dna"` // Group member whitelist DNA.
}

type PostAddGroupMemberWhiteResponse struct {
	core.Response
	Data *PostAddGroupMemberWhiteResult `json:"data"`
}

func PostAddGroupMemberWhitelist_SignatureStr(src_id, dest_id, account_id, sub_account_id, sub_account_name string, created int,
	application_status string) (string, *dtcpv1.GroupMemberWhitelistPost, error) {
	var newCreator dtcpv1.GroupMemberWhitePostCreator
	newCreator.AccountId = account_id
	newCreator.SubAccountId = sub_account_id
	newCreator.SubAccountName = sub_account_name

	var newExtra dtcpv1.GroupMemberWhitePostExtra
	newExtra.ApplicationStatus = application_status

	newGroupMemberWhite := dtcpv1.NewGroupMemberWhitelistPost()
	newGroupMemberWhite.SrcId = src_id
	newGroupMemberWhite.DestId = dest_id
	newGroupMemberWhite.Creator = &newCreator
	newGroupMemberWhite.Created = created
	newGroupMemberWhite.Extra = &newExtra

	err := postGroupMemberWhitelist_check(newGroupMemberWhite)
	if err != nil {
		return "", nil, err
	}

	sigSoure, err := tool.StructToSignature(newGroupMemberWhite)
	if err != nil {
		return "", nil, err
	}

	return sigSoure, newGroupMemberWhite, nil
}

func PostAddGroupMemberWhitelist(signature string, preObj *dtcpv1.GroupMemberWhitelistPost) (*PostAddGroupMemberWhiteResponse, error) {
	if signature == "" {
		return nil, errors.New("param signature is empty")
	}

	err := postGroupMemberWhitelist_check(preObj)
	if err != nil {
		return nil, err
	}

	preObj.Signature = signature

	url := config.Gogal_Server + `/groups/` + preObj.DestId + `/whitelist/members`

	requestBody, err := json.Marshal(preObj)
	if err != nil {
		return nil, err
	}

	response, err := tool.Http_Post(url, string(requestBody))
	if err != nil {
		return nil, err
	}

	log.Printf("response:%v", string(response))

	var responseObj PostAddGroupMemberWhiteResponse
	err = json.Unmarshal(response, &responseObj)
	if err != nil {
		return nil, err
	}

	if responseObj.ResultCode != core.CONST_ResultCode_Success {
		return nil, errors.New(responseObj.ResultMsg)
	}

	return &responseObj, nil
}

func postGroupMemberWhitelist_check(preObj *dtcpv1.GroupMemberWhitelistPost) error {
	if preObj.Version != dtcpv1.CONST_DTCP_Version_v1 {
		return errors.New("parameter version error")
	}

	if preObj.Atype != dtcpv1.CONST_DTCP_Type_Relation {
		return errors.New("parameter type error")
	}

	if preObj.Tag != dtcpv1.CONST_DTCP_Tag_Group_member_whitelist {
		return errors.New("parameter tag error")
	}

	if preObj.SrcId == "" {
		return errors.New("parameter src_id error")
	}

	if preObj.DestId == "" {
		return errors.New("parameter dest_id error")
	}

	if preObj.Creator == nil {
		return errors.New("parameter creator error")
	}

	if preObj.Creator.AccountId == "" {
		return errors.New("parameter account_id error")
	}

	if preObj.Created <= 0 {
		return errors.New("parameter created error")
	}

	if preObj.Extra == nil {
		return errors.New("parameter extra error")
	}

	if preObj.Status != dtcpv1.CONST_DTCP_Status_Created {
		return errors.New("parameter status error")
	}

	if preObj.Extra.ApplicationStatus != dtcpv1.CONST_DTCP_ApplicationStatus_Pending {
		return errors.New("parameter application_status error")
	}

	return nil
}
