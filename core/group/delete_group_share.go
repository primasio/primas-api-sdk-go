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

type DeleteGroupShareResult struct {
	Dna string `json:"dna"` // Share DNA.
}

type DeleteGroupShareResponse struct {
	core.Response
	Data *DeleteGroupShareResult `json:"data"`
}

func DeleteGroupShare_SignatureStr(parent_dna string, updated int,
	account_id, sub_account_id string) (string, *dtcpv1.GroupShareDelete, error) {
	var newCreator dtcpv1.GroupShareDeleteCreator
	newCreator.AccountId = account_id
	newCreator.SubAccountId = sub_account_id

	newGroupShareDelete := dtcpv1.NewGroupShareDelete()
	newGroupShareDelete.ParentDna = parent_dna
	newGroupShareDelete.Updated = updated
	newGroupShareDelete.Creator = &newCreator

	err := deleteGroupShare_check(newGroupShareDelete)
	if err != nil {
		return "", nil, err
	}

	sigSoure, err := tool.StructToSignature(newGroupShareDelete)
	if err != nil {
		return "", nil, err
	}

	return sigSoure, newGroupShareDelete, nil
}

func DeleteGroupShare(share_id, signature string, preObj *dtcpv1.GroupShareDelete) (*DeleteGroupShareResponse, error) {
	if signature == "" {
		return nil, errors.New("param signature is empty")
	}

	if share_id == "" {
		return nil, errors.New("param share_id is empty")
	}

	err := deleteGroupShare_check(preObj)
	if err != nil {
		return nil, err
	}

	preObj.Signature = signature

	url := config.Gogal_Server + `/shares/` + share_id

	requestBody, err := json.Marshal(preObj)
	if err != nil {
		return nil, err
	}

	response, err := tool.Http_Delete(url, string(requestBody))
	if err != nil {
		return nil, err
	}

	log.Printf("response:%v", string(response))

	var responseObj DeleteGroupShareResponse
	err = json.Unmarshal(response, &responseObj)
	if err != nil {
		return nil, err
	}

	if responseObj.ResultCode != core.CONST_ResultCode_Success {
		return nil, errors.New(responseObj.ResultMsg)
	}

	return &responseObj, nil
}

func deleteGroupShare_check(preObj *dtcpv1.GroupShareDelete) error {
	if preObj.Version != dtcpv1.CONST_DTCP_Version_v1 {
		return errors.New("parameter version error")
	}

	if preObj.Atype != dtcpv1.CONST_DTCP_Type_Relation {
		return errors.New("parameter type error")
	}

	if preObj.Tag != dtcpv1.CONST_DTCP_Tag_Group_share {
		return errors.New("parameter tag error")
	}

	if preObj.ParentDna == "" {
		return errors.New("parameter parent_dna error")
	}

	if preObj.Status != dtcpv1.CONST_DTCP_Status_Deleted {
		return errors.New("parameter status error")
	}

	if preObj.Updated <= 0 {
		return errors.New("parameter updated error")
	}

	if preObj.Creator == nil {
		return errors.New("parameter creator error")
	}

	if preObj.Creator.AccountId == "" {
		return errors.New("parameter account_id error")
	}

	return nil
}
