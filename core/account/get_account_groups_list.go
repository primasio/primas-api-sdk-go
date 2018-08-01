package account

import (
	"encoding/json"
	"errors"

	"github.com/primasio/primas-api-sdk-go/config"
	"github.com/primasio/primas-api-sdk-go/core"
	"github.com/primasio/primas-api-sdk-go/core/dtcp/dtcpv1"
	"github.com/primasio/primas-api-sdk-go/core/tool"
)

type AccountGroupListResponse struct {
	core.Response
	Data []dtcpv1.GroupGet `json:"data"`
}

func GetAccountGroupsList(account_id string, page, pageSize int) (*AccountGroupListResponse, error) {
	if account_id == "" {
		return nil, errors.New("account_id is empty")
	}

	url := config.CONST_Server + `/accounts/` + account_id + `/groups`
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

	var responseObj AccountGroupListResponse
	err = json.Unmarshal(response, &responseObj)
	if err != nil {
		return nil, err
	}

	if responseObj.ResultCode != core.CONST_ResultCode_Success {
		return nil, errors.New(responseObj.ResultMsg)
	}

	return &responseObj, nil
}

func GetSubAccountGroupsList(account_id, sub_account_id string, page, pageSize int) (*AccountGroupListResponse, error) {
	if account_id == "" {
		return nil, errors.New("account_id is empty")
	}

	if sub_account_id == "" {
		return nil, errors.New("sub_account_id is empty")
	}

	url := config.CONST_Server + `/accounts/` + account_id + `/sub/` + sub_account_id + `/groups`
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

	var responseObj AccountGroupListResponse
	err = json.Unmarshal(response, &responseObj)
	if err != nil {
		return nil, err
	}

	if responseObj.ResultCode != core.CONST_ResultCode_Success {
		return nil, errors.New(responseObj.ResultMsg)
	}

	return &responseObj, nil
}
