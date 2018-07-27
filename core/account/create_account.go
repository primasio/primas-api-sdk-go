package account

import (
	"encoding/json"
	"errors"

	"github.com/primasio/primas-api-sdk-go/config"
	"github.com/primasio/primas-api-sdk-go/core"
	"github.com/primasio/primas-api-sdk-go/core/dtcp/dtcpv1"
	"github.com/primasio/primas-api-sdk-go/core/tool"
)

type CreateAccountResult struct {
	Id  string `json:"id"`  // The id of the account. No id is returned for sub accounts.
	Dna string `json:"dna"` // The DNA of the account.
}

type CreateAccountResponse struct {
	core.Response
	Data *CreateAccountResult `json:"data"`
}

func CreateAccount(name, abstract, avatar, account_id, sub_account_id string, created int,
	extra_hash string) (*CreateAccountResponse, error) {
	if created < 0 {
		return nil, errors.New("created less than zero")
	}
	if name == "" {
		return nil, errors.New("name is empty")
	}

	var newCreator *dtcpv1.AccountPostCreator
	if account_id != "" {
		if sub_account_id == "" {
			return nil, errors.New("account_id and sub_account_id ")
		} else {
			newCreator = &dtcpv1.AccountPostCreator{
				AccountId:    account_id,
				SubAccountId: sub_account_id,
			}
		}
	}

	var newExtra *dtcpv1.AccountPostExtra
	if extra_hash != "" {
		newExtra = &dtcpv1.AccountPostExtra{
			Hash: extra_hash,
		}
	}

	locAccountPost := dtcpv1.NewAccountPost()
	locAccountPost.Name = name
	locAccountPost.Abstract = abstract
	locAccountPost.Avatar = avatar
	if newCreator != nil {
		locAccountPost.Creator = newCreator
	}
	locAccountPost.Created = created
	if newCreator != nil {
		locAccountPost.Extra = newExtra
	}

	sigSoure, err := tool.StructToSignature(locAccountPost)
	if err != nil {
		return nil, err
	}
	privateKey := tool.GetClientPrivateKey()

	signature, err := tool.Sign([]byte(sigSoure), privateKey)
	if err != nil {
		return nil, err
	}
	locAccountPost.Signature = signature

	url := config.CONST_Server + `/accounts`

	requestBody, err := json.Marshal(locAccountPost)
	if err != nil {
		return nil, err
	}

	response, err := tool.Http_Post(url, string(requestBody))
	if err != nil {
		return nil, err
	}

	var responseObj CreateAccountResponse
	err = json.Unmarshal(response, &responseObj)
	if err != nil {
		return nil, err
	}

	if responseObj.ResultCode != core.CONST_ResultCode_Success {
		return nil, errors.New(responseObj.ResultMsg)
	}

	return &responseObj, nil
}
