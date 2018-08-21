package content_interaction

import (
	"encoding/json"
	"errors"

	"github.com/primasio/primas-api-sdk-go/config"
	"github.com/primasio/primas-api-sdk-go/core"
	"github.com/primasio/primas-api-sdk-go/core/dtcp/dtcpv1"
	"github.com/primasio/primas-api-sdk-go/core/tool"
)

type PostLikeGroupshareResult struct {
	Id  string `json:"id"`  // Like id.
	Dna string `json:"dna"` // Like dna.
}

type PostLikeGroupshareResponse struct {
	core.Response
	Data *PostLikeGroupshareResult `json:"data"`
}

func PostLikeOfGroupshare_SignatureStr(account_id, share_id, sub_account_id, sub_account_name string,
	hp, created int) (string, *dtcpv1.ContentLikePost, error) {
	newPostLikeCreator := dtcpv1.ContentLikePostCreator{
		AccountId: account_id,
	}
	if sub_account_id != "" {
		newPostLikeCreator.SubAccountId = sub_account_id
	}
	if sub_account_name != "" {
		newPostLikeCreator.SubAccountName = sub_account_name
	}

	newLikeGroupshare := dtcpv1.NewContentLikePost()
	newLikeGroupshare.SrcId = account_id
	newLikeGroupshare.DestId = share_id
	newLikeGroupshare.Hp = hp
	newLikeGroupshare.Creator = &newPostLikeCreator
	newLikeGroupshare.Created = created

	err := postLikeOfGroupshare_check(newLikeGroupshare)
	if err != nil {
		return "", nil, err
	}

	sigSoure, err := tool.StructToSignature(newLikeGroupshare)
	if err != nil {
		return "", nil, err
	}

	return sigSoure, newLikeGroupshare, nil
}

func PostLikeOfGroupshare(signature string, preObj *dtcpv1.ContentLikePost) (*PostLikeGroupshareResponse, error) {
	if signature == "" {
		return nil, errors.New("param signature is empty")
	}

	err := postLikeOfGroupshare_check(preObj)
	if err != nil {
		return nil, err
	}

	preObj.Signature = signature

	url := config.CONST_Server + `/shares/` + preObj.DestId + `/likes`

	requestBody, err := json.Marshal(preObj)
	if err != nil {
		return nil, err
	}

	response, err := tool.Http_Post(url, string(requestBody))
	if err != nil {
		return nil, err
	}

	var responseObj PostLikeGroupshareResponse
	err = json.Unmarshal(response, &responseObj)
	if err != nil {
		return nil, err
	}

	if responseObj.ResultCode != core.CONST_ResultCode_Success {
		return nil, errors.New(responseObj.ResultMsg)
	}

	return &responseObj, nil
}

func postLikeOfGroupshare_check(value *dtcpv1.ContentLikePost) error {
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

	return nil
}
