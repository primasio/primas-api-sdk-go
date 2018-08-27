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

	"github.com/primasio/primas-api-sdk-go/config"
	"github.com/primasio/primas-api-sdk-go/core"
	"github.com/primasio/primas-api-sdk-go/core/dtcp/dtcpv1"
	"github.com/primasio/primas-api-sdk-go/core/tool"
)

type MemberApplication struct {
	Dna string `json:"dna"` // Group member DNA.
}

type PutMemberApplicationResponse struct {
	core.Response
	Data *MemberApplication `json:"data"`
}

func PutMemberApplication_SignatureStr(parent_dna string, updated int, account_id, sub_account_id,
	application_status string) (string, *dtcpv1.GroupMemberApplicationPut, error) {
	var newCreator dtcpv1.GroupMemberAppPutCreator
	newCreator.AccountId = account_id
	newCreator.SubAccountId = sub_account_id

	var newExtra dtcpv1.GroupMemberAppPutExtra
	newExtra.ApplicationStatus = application_status

	newGroupMemberApp := dtcpv1.NewGroupMemberApplicationPut()
	newGroupMemberApp.ParentDna = parent_dna
	newGroupMemberApp.Updated = updated
	newGroupMemberApp.Creator = &newCreator
	newGroupMemberApp.Extra = &newExtra

	err := putGroupMemberApplication_check(newGroupMemberApp)
	if err != nil {
		return "", nil, err
	}

	sigSoure, err := tool.StructToSignature(newGroupMemberApp)
	if err != nil {
		return "", nil, err
	}

	return sigSoure, newGroupMemberApp, nil
}

func PutMemberApplication(group_id, group_member_id, signature string,
	preObj *dtcpv1.GroupMemberApplicationPut) (*PutMemberApplicationResponse, error) {
	if signature == "" {
		return nil, errors.New("param signature is empty")
	}

	if group_id == "" {
		return nil, errors.New("param group_id is empty")
	}

	if group_member_id == "" {
		return nil, errors.New("param group_member_id is empty")
	}

	err := putGroupMemberApplication_check(preObj)
	if err != nil {
		return nil, err
	}

	preObj.Signature = signature

	url := config.Gogal_Server + `/groups/` + group_id + `/members/` + group_member_id

	requestBody, err := json.Marshal(preObj)
	if err != nil {
		return nil, err
	}

	response, err := tool.Http_Put(url, string(requestBody))
	if err != nil {
		return nil, err
	}

	var responseObj PutMemberApplicationResponse
	err = json.Unmarshal(response, &responseObj)
	if err != nil {
		return nil, err
	}

	if responseObj.ResultCode != core.CONST_ResultCode_Success {
		return nil, errors.New(responseObj.ResultMsg)
	}

	return &responseObj, nil
}

func putGroupMemberApplication_check(preObj *dtcpv1.GroupMemberApplicationPut) error {
	if preObj.Version != dtcpv1.CONST_DTCP_Version_v1 {
		return errors.New("parameter version error")
	}

	if preObj.Atype != dtcpv1.CONST_DTCP_Type_Relation {
		return errors.New("parameter type error")
	}

	if preObj.Tag != dtcpv1.CONST_DTCP_Tag_Group_member {
		return errors.New("parameter tag error")
	}

	if preObj.ParentDna == "" {
		return errors.New("parameter parent_dna error")
	}

	if preObj.Creator == nil {
		return errors.New("parameter creator error")
	}

	if preObj.Creator.AccountId == "" {
		return errors.New("parameter account_id error")
	}

	if preObj.Updated <= 0 {
		return errors.New("parameter updated error")
	}

	if preObj.Extra == nil {
		return errors.New("parameter extra error")
	}

	if preObj.Status != dtcpv1.CONST_DTCP_Status_Updated {
		return errors.New("parameter status error")
	}

	if preObj.Extra.ApplicationStatus != dtcpv1.CONST_DTCP_ApplicationStatus_Approved &&
		preObj.Extra.ApplicationStatus != dtcpv1.CONST_DTCP_ApplicationStatus_Declined {
		return errors.New("parameter application_status error")
	}

	return nil
}
