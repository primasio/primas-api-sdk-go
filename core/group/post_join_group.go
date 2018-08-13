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

type PostJoinGroupResult struct {
	Id  string `json:"id"`  // Group member id.
	Dna string `json:"dna"` // Group member DNA.
}

type PostJoinGroupResponse struct {
	core.Response
	Data *PostJoinGroupResult `json:"data"`
}

func PostJoinGroup_SignatureStr(src_id, dest_id, account_id, sub_account_id, sub_account_name string, created int,
	application_status string, application_expire int) (string, *dtcpv1.GroupMemberPost, error) {
	var newCreator dtcpv1.GroupMemberPostCreator
	newCreator.AccountId = account_id
	newCreator.SubAccountId = sub_account_id
	newCreator.SubAccountName = sub_account_name

	var newExtra dtcpv1.GroupMemberPostExtra
	newExtra.ApplicationStatus = application_status
	newExtra.ApplicationExpire = application_expire

	newJoinGroup := dtcpv1.NewGroupMemberPost()
	newJoinGroup.SrcId = src_id
	newJoinGroup.DestId = dest_id
	newJoinGroup.Creator = &newCreator
	newJoinGroup.Created = created
	newJoinGroup.Extra = &newExtra

	err := postJoinGroup_check(newJoinGroup)
	if err != nil {
		return "", nil, err
	}

	sigSoure, err := tool.StructToSignature(newJoinGroup)
	if err != nil {
		return "", nil, err
	}

	return sigSoure, newJoinGroup, nil
}

func PostJoinGroup(signature string, preObj *dtcpv1.GroupMemberPost) (*PostJoinGroupResponse, error) {
	if signature == "" {
		return nil, errors.New("param signature is empty")
	}

	err := postJoinGroup_check(preObj)
	if err != nil {
		return nil, err
	}

	preObj.Signature = signature

	url := config.CONST_Server + `/groups/` + preObj.DestId + `/members`

	requestBody, err := json.Marshal(preObj)
	if err != nil {
		return nil, err
	}

	response, err := tool.Http_Post(url, string(requestBody))
	if err != nil {
		return nil, err
	}

	log.Printf("response:%v", string(response))

	var responseObj PostJoinGroupResponse
	err = json.Unmarshal(response, &responseObj)
	if err != nil {
		return nil, err
	}

	if responseObj.ResultCode != core.CONST_ResultCode_Success {
		return nil, errors.New(responseObj.ResultMsg)
	}

	return &responseObj, nil
}

func postJoinGroup_check(preObj *dtcpv1.GroupMemberPost) error {
	if preObj.Version != dtcpv1.CONST_DTCP_Version_v1 {
		return errors.New("parameter version error")
	}

	if preObj.Atype != dtcpv1.CONST_DTCP_Type_Relation {
		return errors.New("parameter type error")
	}

	if preObj.Tag != dtcpv1.CONST_DTCP_Tag_Group_member {
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
