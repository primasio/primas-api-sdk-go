package system

import (
	"encoding/json"
	"errors"

	"github.com/primasio/primas-api-sdk-go/config"
	"github.com/primasio/primas-api-sdk-go/core"
	"github.com/primasio/primas-api-sdk-go/core/dtcp/dtcpv1"
	"github.com/primasio/primas-api-sdk-go/core/tool"
)

type GetSystemParametersResponse struct {
	core.Response
	Data *dtcpv1.SystemParametersGet `json:"data"`
}

func GetSystemParameters() (*GetSystemParametersResponse, error) {
	url := config.CONST_Server + `/system`

	queryParams := make(map[string]interface{}, 0)
	response, err := tool.Http_Get(url, queryParams)
	if err != nil {
		return nil, err
	}

	var responseObj GetSystemParametersResponse
	err = json.Unmarshal(response, &responseObj)
	if err != nil {
		return nil, err
	}

	if responseObj.ResultCode != core.CONST_ResultCode_Success {
		return nil, errors.New(responseObj.ResultMsg)
	}

	return &responseObj, nil
}
