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

type QuitGroupMemberOutResult struct {
	Dna string `json:"dna"` // Group member DNA.
}

type QuitGroupMemberOutResponse struct {
	core.Response
	Data *QuitGroupMemberOutResult `json:"data"`
}

func DeleteQuitGroupOrKickMemberOut_SignatureStr(parent_dna string, updated int,
	account_id, sub_account_id string) (string, *dtcpv1.GroupMemberDelete, error) {
	var newGroupMemberDeleteCreator dtcpv1.GroupMemberDeleteCreator
	newGroupMemberDeleteCreator.AccountId = account_id
	newGroupMemberDeleteCreator.SubAccountId = sub_account_id

	newQuitGroupMember := dtcpv1.NewGroupMemberDelete()
	newQuitGroupMember.ParentDna = parent_dna
	newQuitGroupMember.Updated = updated
	newQuitGroupMember.Creator = &newGroupMemberDeleteCreator

	err := deleteQuitGroupOrKickMemberOut_check(newQuitGroupMember)
	if err != nil {
		return "", nil, err
	}

	sigSoure, err := tool.StructToSignature(newQuitGroupMember)
	if err != nil {
		return "", nil, err
	}

	return sigSoure, newQuitGroupMember, nil
}

func DeleteQuitGroupOrKickMemberOut(group_id, group_member_id, signature string,
	preObj *dtcpv1.GroupMemberDelete) (*QuitGroupMemberOutResponse, error) {
	if signature == "" {
		return nil, errors.New("param signature is empty")
	}

	if group_id == "" {
		return nil, errors.New("param group_id is empty")
	}

	if group_member_id == "" {
		return nil, errors.New("param group_member_id is empty")
	}

	err := deleteQuitGroupOrKickMemberOut_check(preObj)
	if err != nil {
		return nil, err
	}

	preObj.Signature = signature

	url := config.CONST_Server + `/groups/` + group_id + `/members/` + group_member_id

	requestBody, err := json.Marshal(preObj)
	if err != nil {
		return nil, err
	}

	response, err := tool.Http_Delete(url, string(requestBody))
	if err != nil {
		return nil, err
	}

	log.Printf("response:%v", string(response))

	var responseObj QuitGroupMemberOutResponse
	err = json.Unmarshal(response, &responseObj)
	if err != nil {
		return nil, err
	}

	if responseObj.ResultCode != core.CONST_ResultCode_Success {
		return nil, errors.New(responseObj.ResultMsg)
	}

	return &responseObj, nil
}

func deleteQuitGroupOrKickMemberOut_check(preObj *dtcpv1.GroupMemberDelete) error {
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
