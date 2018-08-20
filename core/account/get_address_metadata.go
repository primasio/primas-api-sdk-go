package account

import (
	"encoding/json"
	"errors"

	"github.com/primasio/primas-api-sdk-go/config"
	"github.com/primasio/primas-api-sdk-go/core"
	"github.com/primasio/primas-api-sdk-go/core/tool"
)

type AddressMetadata struct {
	Id            string                  `json:"id"`                 // Account id.
	Address       string                  `json:"address"`            // Account address.
	Title         string                  `json:"title"`              // Account name.
	Abstract      string                  `json:"abstract,omitempty"` // Description.
	Avatar        string                  `json:"avatar,omitempty"`   // An image id used for avatar.
	Creator       *AddressMetadataCreator `json:"creator"`            // Creator of the sub account.
	Created       int                     `json:"created"`            // Account creation time. Unix timestamp.
	Updated       int                     `json:"updated"`            // Account last updating time. Unix timestamp.
	Extra         *AddressMetadataExtra   `json:"extra,omitempty"`    // Extra metadata.
	Signature     string                  `json:"signature"`          // Metadata signature.
	Dna           string                  `json:"dna"`                // DNA of the account.
	Credits       int                     `json:"credits"`            // Current credits.
	TransactionId string                  `json:"transaction_id"`     // Latest transaction id.
}

type AddressMetadataCreator struct {
	AccountId   string `json:"account_id"`   // Root account id.
	AccountName string `json:"account_name"` // Root account name.
}

type AddressMetadataExtra struct {
	Hash string `json:"hash"` // In the case of proof of existence of secret data. The hash can be filled in this field.
}

type AddressMetadataResponse struct {
	core.Response
	Data *AddressMetadata `json:"data"`
}

func GetAddressMetadata(address string) (*AddressMetadataResponse, error) {
	if address == "" {
		return nil, errors.New("address is empty")
	}

	url := config.CONST_Server + `/main/accounts/` + address + `/metadata`
	response, err := tool.Http_Get(url, nil)
	if err != nil {
		return nil, err
	}

	var responseObj AddressMetadataResponse
	err = json.Unmarshal(response, &responseObj)
	if err != nil {
		return nil, err
	}

	if responseObj.ResultCode != core.CONST_ResultCode_Success {
		return nil, errors.New(responseObj.ResultMsg)
	}

	return &responseObj, nil
}
