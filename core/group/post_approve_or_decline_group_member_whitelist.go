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

type PutAddGroupMemberWhiteResult struct {
	Dna string `json:"dna"` // Group member whitelist DNA.
}

type PutAddGroupMemberWhiteResponse struct {
	core.Response
	Data *PutAddGroupMemberWhiteResult `json:"data"`
}

func PostApproveOrDeclineGroupMemberWhitelist_SignatureStr(parent_dna string, updated int,
	account_id, sub_account_id string, application_status string) (string, *dtcpv1.GroupMemberWhitelistPut, error) {
	var putCreator dtcpv1.GroupMemberWhitePutCreator
	putCreator.AccountId = account_id
	putCreator.SubAccountId = sub_account_id

	var putExtra dtcpv1.GroupMemberWhitePutExtra
	putExtra.ApplicationStatus = application_status

	putGroupMemberWhite := dtcpv1.NewGroupMemberWhitelistPut()
	putGroupMemberWhite.ParentDna = parent_dna
	putGroupMemberWhite.Creator = &putCreator
	putGroupMemberWhite.Updated = updated
	putGroupMemberWhite.Extra = &putExtra

	err := putGroupMemberWhitelist_check(putGroupMemberWhite)
	if err != nil {
		return "", nil, err
	}

	sigSoure, err := tool.StructToSignature(putGroupMemberWhite)
	if err != nil {
		return "", nil, err
	}

	return sigSoure, putGroupMemberWhite, nil
}

func PostApproveOrDeclineGroupMemberWhitelist(group_id, whitelist_id, signature string,
	preObj *dtcpv1.GroupMemberWhitelistPut) (*PutAddGroupMemberWhiteResponse, error) {
	if signature == "" {
		return nil, errors.New("param signature is empty")
	}

	if group_id == "" {
		return nil, errors.New("param group_id is empty")
	}

	if whitelist_id == "" {
		return nil, errors.New("param whitelist_id is empty")
	}

	err := putGroupMemberWhitelist_check(preObj)
	if err != nil {
		return nil, err
	}

	preObj.Signature = signature

	url := config.CONST_Server + `/groups/` + group_id + `/whitelist/members/` + whitelist_id

	requestBody, err := json.Marshal(preObj)
	if err != nil {
		return nil, err
	}

	response, err := tool.Http_Put(url, string(requestBody))
	if err != nil {
		return nil, err
	}

	log.Printf("response:%v", string(response))

	var responseObj PutAddGroupMemberWhiteResponse
	err = json.Unmarshal(response, &responseObj)
	if err != nil {
		return nil, err
	}

	if responseObj.ResultCode != core.CONST_ResultCode_Success {
		return nil, errors.New(responseObj.ResultMsg)
	}

	return &responseObj, nil
}

func putGroupMemberWhitelist_check(preObj *dtcpv1.GroupMemberWhitelistPut) error {
	if preObj.Version != dtcpv1.CONST_DTCP_Version_v1 {
		return errors.New("parameter version error")
	}

	if preObj.Atype != dtcpv1.CONST_DTCP_Type_Relation {
		return errors.New("parameter type error")
	}

	if preObj.Tag != dtcpv1.CONST_DTCP_Tag_Group_member_whitelist {
		return errors.New("parameter tag error")
	}

	if preObj.Status != dtcpv1.CONST_DTCP_Status_Updated {
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

	if preObj.Extra == nil {
		return errors.New("parameter extra error")
	}

	if preObj.Extra.ApplicationStatus != dtcpv1.CONST_DTCP_ApplicationStatus_Approved &&
		preObj.Extra.ApplicationStatus != dtcpv1.CONST_DTCP_ApplicationStatus_Declined {
		return errors.New("parameter application_status error")
	}

	return nil
}