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
	Id  string `json:"id,omitempty"` // The id of the account. No id is returned for sub accounts.
	Dna string `json:"dna"`          // The DNA of the account.
}

type CreateAccountResponse struct {
	core.Response
	Data *CreateAccountResult `json:"data"`
}

// CreateAccount Signature string value
func CreateAccount_SignatureStr(name, abstract, avatar, account_id, sub_account_id string, created int,
	extra_hash string) (string, *dtcpv1.AccountPost, error) {
	var newCreator *dtcpv1.AccountPostCreator
	if account_id != "" {
		if sub_account_id == "" {
			return "", nil, errors.New("account_id and sub_account_id ")
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

	err := createAccount_check(locAccountPost)
	if err != nil {
		return "", nil, err
	}

	sigSoure, err := tool.StructToSignature(locAccountPost)
	if err != nil {
		return "", nil, err
	}

	return sigSoure, locAccountPost, nil
}

func CreateAccount(signature string, preObj *dtcpv1.AccountPost) (*CreateAccountResponse, error) {
	if signature == "" {
		return nil, errors.New("param signature is empty")
	}

	err := createAccount_check(preObj)
	if err != nil {
		return nil, err
	}

	preObj.Signature = signature

	url := config.CONST_Server + `/accounts`

	requestBody, err := json.Marshal(preObj)
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

func createAccount_check(preObj *dtcpv1.AccountPost) error {
	if preObj == nil {
		return errors.New("param preObj is nil")
	}

	if preObj.Created < 0 {
		return errors.New("created less than zero")
	}
	if preObj.Name == "" {
		return errors.New("name is empty")
	}

	if preObj.Version != dtcpv1.CONST_DTCP_Version_v1 {
		return errors.New("param version error")
	}

	if preObj.Atype != dtcpv1.CONST_DTCP_Type_Object {
		return errors.New("param type error")
	}

	if preObj.Tag != dtcpv1.CONST_DTCP_Tag_Account {
		return errors.New("param tag error")
	}

	if preObj.Created <= 0 {
		return errors.New("param created error")
	}

	if preObj.Address == "" {
		return errors.New("param address is empty")
	}

	if preObj.Status != dtcpv1.CONST_DTCP_Status_Created {
		return errors.New("param status error")
	}

	return nil
}
