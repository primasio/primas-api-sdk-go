package account

import (
	"encoding/json"
	"errors"

	"github.com/primasio/primas-api-sdk-go/config"
	"github.com/primasio/primas-api-sdk-go/core"
	"github.com/primasio/primas-api-sdk-go/core/tool"
)

type GetAccountMetadata struct {
	Id            string                     `json:"id"`                 // Account id.
	Address       string                     `json:"address"`            // Account address.
	Title         string                     `json:"title"`              // Account name.
	Abstract      string                     `json:"abstract,omitempty"` // Description.
	Avatar        string                     `json:"avatar,omitempty"`   // An image id used for avatar.
	Creator       *GetAccountMetadataCreator `json:"creator"`            // Creator of the sub account.
	Created       int                        `json:"created"`            // Account creation time. Unix timestamp.
	Updated       int                        `json:"updated"`            // Account last updating time. Unix timestamp.
	Extra         *GetAccountMetadataExtra   `json:"extra,omitempty"`    // Extra metadata.
	Signature     string                     `json:"signature"`          // Metadata signature.
	Dna           string                     `json:"dna"`                // DNA of the account.
	Credits       int                        `json:"credits"`            // Current credits.
	TransactionId string                     `json:"transaction_id"`     // Latest transaction id.
}

type GetAccountMetadataCreator struct {
	AccountId   string `json:"account_id"`   // Root account id.
	AccountName string `json:"account_name"` // Root account name.
}

type GetAccountMetadataExtra struct {
	Hash string `json:"hash"` // In the case of proof of existence of secret data. The hash can be filled in this field.
}

type GetAccountMetadataResponse struct {
	core.Response
	Data *GetAccountMetadata `json:"data"`
}

func GetAccountTokenMetadata(account_id string) (*GetAccountMetadataResponse, error) {
	if account_id == "" {
		return nil, errors.New("account_id is empty")
	}

	url := config.CONST_Server + `/accounts/` + account_id + `/metadata`
	response, err := tool.Http_Get(url, nil)
	if err != nil {
		return nil, err
	}

	var responseObj GetAccountMetadataResponse
	err = json.Unmarshal(response, &responseObj)
	if err != nil {
		return nil, err
	}

	if responseObj.ResultCode != core.CONST_ResultCode_Success {
		return nil, errors.New(responseObj.ResultMsg)
	}

	return &responseObj, nil
}

func GetSubAccountTokenMetadata(account_id, sub_account_id string) (*GetAccountMetadataResponse, error) {
	if account_id == "" {
		return nil, errors.New("account_id is empty")
	}

	if sub_account_id == "" {
		return nil, errors.New("sub_account_id is empty")
	}

	url := config.CONST_Server + `/accounts/` + account_id + `/sub/` + sub_account_id + `/metadata`
	response, err := tool.Http_Get(url, nil)
	if err != nil {
		return nil, err
	}

	var responseObj GetAccountMetadataResponse
	err = json.Unmarshal(response, &responseObj)
	if err != nil {
		return nil, err
	}

	if responseObj.ResultCode != core.CONST_ResultCode_Success {
		return nil, errors.New(responseObj.ResultMsg)
	}

	return &responseObj, nil
}
