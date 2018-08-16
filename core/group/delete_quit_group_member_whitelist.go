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

type DelAddGroupMemberWhiteResult struct {
	Dna string `json:"dna"` // Group member whitelist DNA.
}

type DelAddGroupMemberWhiteResponse struct {
	core.Response
	Data *DelAddGroupMemberWhiteResult `json:"data"`
}

func DeleteQuitGroupMemberWhitelist_SignatureStr(parent_dna string, updated int,
	account_id, sub_account_id string) (string, *dtcpv1.GroupMemberWhitelistDelete, error) {

	var delCreator dtcpv1.GroupMemberWhitelistDeleteCreator
	delCreator.AccountId = account_id
	delCreator.SubAccountId = sub_account_id

	delGroupMemberWhite := dtcpv1.NewGroupMemberWhitelistDelete()
	delGroupMemberWhite.ParentDna = parent_dna
	delGroupMemberWhite.Creator = &delCreator
	delGroupMemberWhite.Updated = updated

	err := delGroupMemberWhitelist_check(delGroupMemberWhite)
	if err != nil {
		return "", nil, err
	}

	sigSoure, err := tool.StructToSignature(delGroupMemberWhite)
	if err != nil {
		return "", nil, err
	}

	return sigSoure, delGroupMemberWhite, nil
}

func DeleteQuitGroupMemberWhitelist(group_id, whitelist_id, signature string,
	preObj *dtcpv1.GroupMemberWhitelistDelete) (*DelAddGroupMemberWhiteResponse, error) {
	if signature == "" {
		return nil, errors.New("param signature is empty")
	}

	if group_id == "" {
		return nil, errors.New("param group_id is empty")
	}

	if whitelist_id == "" {
		return nil, errors.New("param whitelist_id is empty")
	}

	err := delGroupMemberWhitelist_check(preObj)
	if err != nil {
		return nil, err
	}

	preObj.Signature = signature

	url := config.CONST_Server + `/groups/` + group_id + `/whitelist/members/` + whitelist_id

	requestBody, err := json.Marshal(preObj)
	if err != nil {
		return nil, err
	}

	response, err := tool.Http_Delete(url, string(requestBody))
	if err != nil {
		return nil, err
	}

	log.Printf("response:%v", string(response))

	var responseObj DelAddGroupMemberWhiteResponse
	err = json.Unmarshal(response, &responseObj)
	if err != nil {
		return nil, err
	}

	if responseObj.ResultCode != core.CONST_ResultCode_Success {
		return nil, errors.New(responseObj.ResultMsg)
	}

	return &responseObj, nil
}

func delGroupMemberWhitelist_check(preObj *dtcpv1.GroupMemberWhitelistDelete) error {
	if preObj.Version != dtcpv1.CONST_DTCP_Version_v1 {
		return errors.New("parameter version error")
	}

	if preObj.Atype != dtcpv1.CONST_DTCP_Type_Relation {
		return errors.New("parameter type error")
	}

	if preObj.Tag != dtcpv1.CONST_DTCP_Tag_Group_member_whitelist {
		return errors.New("parameter tag error")
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
