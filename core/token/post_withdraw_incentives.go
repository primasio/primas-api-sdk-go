package token

import (
	"encoding/json"
	"errors"

	"github.com/primasio/primas-api-sdk-go/config"
	"github.com/primasio/primas-api-sdk-go/core"
	"github.com/primasio/primas-api-sdk-go/core/tool"
	"github.com/shopspring/decimal"
)

type AccountIdWithdrawalRequest struct {
	NodeId    string          `json:"node_id"`   // Node id.
	Created   int             `json:"created"`   // Withdrawal creation time. Unix timestamp.
	Amount    decimal.Decimal `json:"amount"`    // Withdraw amount value.
	NodeFee   decimal.Decimal `json:"node_fee"`  // Node charged withdrawal fee.
	Signature string          `json:"signature"` // Metadata signature.
}

type AccountIdWithdrawalId struct {
	Id string
}

type AccountIdWithdrawalResponse struct {
	core.Response
	Data *AccountIdWithdrawalId `json:"data"`
}

func PostWithdrawIncentives(account_id, node_id string, created int, amount decimal.Decimal, node_fee decimal.Decimal) (*AccountIdWithdrawalResponse, error) {
	if node_id == "" {
		return nil, errors.New("node_id is empty")
	}

	if created < 0 {
		return nil, errors.New("created less than zero")
	}

	if amount.Cmp(decimal.NewFromFloat(0)) <= 0 {
		return nil, errors.New("amount less than or equal to zero")
	}

	if node_fee.Cmp(decimal.NewFromFloat(0)) < 0 {
		return nil, errors.New("node_fee less than zero")
	}

	requestValue := AccountIdWithdrawalRequest{
		NodeId:  node_id,
		Created: created,
		Amount:  amount,
		NodeFee: node_fee,
	}

	sigSoure, err := tool.StructToSignature(requestValue)
	if err != nil {
		return nil, err
	}
	privateKey := tool.GetClientPrivateKey()

	signature, err := tool.Sign([]byte(sigSoure), privateKey)
	if err != nil {
		return nil, err
	}
	requestValue.Signature = signature

	url := config.CONST_Server + `/accounts/` + account_id + `/tokens/incentives/withdrawal`

	requestBody, err := json.Marshal(requestValue)
	if err != nil {
		return nil, err
	}

	response, err := tool.Http_Post(url, string(requestBody))
	if err != nil {
		return nil, err
	}

	//log.Println("response:", string(response))

	var responseObj AccountIdWithdrawalResponse
	err = json.Unmarshal(response, &responseObj)
	if err != nil {
		return nil, err
	}

	if responseObj.ResultCode != core.CONST_ResultCode_Success {
		return nil, errors.New(responseObj.ResultMsg)
	}

	return &responseObj, nil
}
