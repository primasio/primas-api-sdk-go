package account

import (
	"encoding/json"
	"errors"

	"github.com/primasio/primas-api-sdk-go/config"
	"github.com/primasio/primas-api-sdk-go/core"
	"github.com/primasio/primas-api-sdk-go/core/dtcp/dtcpv1"
	"github.com/primasio/primas-api-sdk-go/core/tool"
)

type AccountSharesInSingleGroup struct {
	core.Response
	Data []dtcpv1.Share `json:"data"`
}

func GetAccountSharesInSingleGroup(account_id, group_id string, page, pageSize int) (*AccountSharesInSingleGroup, error) {
	if account_id == "" {
		return nil, errors.New("account_id is empty")
	}
	if group_id == "" {
		return nil, errors.New("group_id is empty")
	}

	url := config.CONST_Server + `/accounts/` + account_id + `/groups/` + group_id + `/shares`
	queryParams := make(map[string]interface{}, 0)
	if page > 0 {
		queryParams["page"] = page
	}
	if pageSize > 0 {
		queryParams["page_size"] = pageSize
	}

	response, err := tool.Http_Get(url, queryParams)
	if err != nil {
		return nil, err
	}

	var responseObj AccountSharesInSingleGroup
	err = json.Unmarshal(response, &responseObj)
	if err != nil {
		return nil, err
	}

	if responseObj.ResultCode != core.CONST_ResultCode_Success {
		return nil, errors.New(responseObj.ResultMsg)
	}

	return &responseObj, nil
}

func GetSubAccountSharesInSingleGroup(account_id, sub_account_id, group_id string, page, pageSize int) (*AccountSharesInSingleGroup, error) {
	if account_id == "" {
		return nil, errors.New("account_id is empty")
	}

	if sub_account_id == "" {
		return nil, errors.New("sub_account_id is empty")
	}

	if group_id == "" {
		return nil, errors.New("group_id is empty")
	}

	url := config.CONST_Server + `/accounts/` + account_id + `/sub/` + sub_account_id + `/groups/` + group_id + `/shares`
	queryParams := make(map[string]interface{}, 0)
	if page > 0 {
		queryParams["page"] = page
	}
	if pageSize > 0 {
		queryParams["page_size"] = pageSize
	}

	response, err := tool.Http_Get(url, queryParams)
	if err != nil {
		return nil, err
	}

	var responseObj AccountSharesInSingleGroup
	err = json.Unmarshal(response, &responseObj)
	if err != nil {
		return nil, err
	}

	if responseObj.ResultCode != core.CONST_ResultCode_Success {
		return nil, errors.New(responseObj.ResultMsg)
	}

	return &responseObj, nil
}
