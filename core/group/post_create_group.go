package group

import (
	"encoding/json"
	"errors"

	"github.com/primasio/primas-api-sdk-go/config"
	"github.com/primasio/primas-api-sdk-go/core"
	"github.com/primasio/primas-api-sdk-go/core/dtcp/dtcpv1"
	"github.com/primasio/primas-api-sdk-go/core/tool"
)

type PostGroupResult struct {
	Id  string `json:"id"`  // The id of the group.
	Dna string `json:"dna"` // The DNA of the group.
}

type PostGroupResponse struct {
	core.Response
	Data *PostGroupResult `json:"data"`
}

func PostGroup_SignatureStr(title, account_id, sub_account_id, sub_account_name, avatar, abstract, language, category string,
	created int, allow_join, allow_post string) (string, *dtcpv1.GroupPost, error) {
	var newCreator dtcpv1.GroupPostCreator
	newCreator.AccountId = account_id
	newCreator.SubAccountId = sub_account_id
	newCreator.SubAccountName = sub_account_name

	newExtra := dtcpv1.GroupPostExtra{
		AllowJoin: allow_join,
		AllowPost: allow_post,
	}

	newCreateGroup := dtcpv1.NewGroupPost()
	newCreateGroup.Title = title
	newCreateGroup.Creator = &newCreator
	newCreateGroup.Avatar = avatar
	newCreateGroup.Abstract = abstract
	newCreateGroup.Language = language
	newCreateGroup.Category = category
	newCreateGroup.Created = created
	newCreateGroup.Extra = &newExtra

	err := postGroup_check(newCreateGroup)
	if err != nil {
		return "", nil, err
	}

	sigSoure, err := tool.StructToSignature(newCreateGroup)
	if err != nil {
		return "", nil, err
	}

	return sigSoure, newCreateGroup, nil
}

func PostGroup(signature string, preObj *dtcpv1.GroupPost) (*PostGroupResponse, error) {
	if signature == "" {
		return nil, errors.New("param signature is empty")
	}

	err := postGroup_check(preObj)
	if err != nil {
		return nil, err
	}

	preObj.Signature = signature

	url := config.CONST_Server + `/groups`

	requestBody, err := json.Marshal(preObj)
	if err != nil {
		return nil, err
	}

	response, err := tool.Http_Post(url, string(requestBody))
	if err != nil {
		return nil, err
	}

	var responseObj PostGroupResponse
	err = json.Unmarshal(response, &responseObj)
	if err != nil {
		return nil, err
	}

	if responseObj.ResultCode != core.CONST_ResultCode_Success {
		return nil, errors.New(responseObj.ResultMsg)
	}

	return &responseObj, nil
}

func postGroup_check(preObj *dtcpv1.GroupPost) error {
	if preObj.Version != dtcpv1.CONST_DTCP_Version_v1 {
		return errors.New("parameter version error")
	}

	if preObj.Atype != dtcpv1.CONST_DTCP_Type_Object {
		return errors.New("parameter type error")
	}

	if preObj.Tag != dtcpv1.CONST_DTCP_Tag_Group {
		return errors.New("parameter tag error")
	}

	if preObj.Title == "" {
		return errors.New("parameter title error")
	}

	if preObj.Creator == nil {
		return errors.New("parameter creator error")
	}

	if preObj.Creator.AccountId == "" {
		return errors.New("parameter account_id error")
	}

	if preObj.Avatar == "" {
		return errors.New("parameter avatar error")
	}

	if preObj.Abstract == "" {
		return errors.New("parameter abstract error")
	}

	if preObj.Language == "" {
		return errors.New("parameter language error")
	}

	if preObj.Category == "" {
		return errors.New("parameter category error")
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

	if preObj.Extra.AllowPost == "" {
		return errors.New("parameter allow_post error")
	}

	if preObj.Extra.AllowJoin == "" {
		return errors.New("parameter allow_join error")
	}

	return nil
}
