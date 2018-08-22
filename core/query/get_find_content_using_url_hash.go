package query

import (
	"encoding/json"
	"errors"

	"github.com/primasio/primas-api-sdk-go/config"
	"github.com/primasio/primas-api-sdk-go/core"
	"github.com/primasio/primas-api-sdk-go/core/dtcp/dtcpv1"
	"github.com/primasio/primas-api-sdk-go/core/tool"
)

type FindContentUsingUrlResponse struct {
	core.Response
	Data []dtcpv1.ContentGet `json:"data"`
}

func FindContentUsingUrlOrHash(urlParam, hashParam string) (*FindContentUsingUrlResponse, error) {
	if urlParam == "" && hashParam == "" {
		return nil, errors.New("url and hash are all empty")
	}

	queryParams := make(map[string]interface{}, 0)

	if urlParam != "" {
		queryParams["url"] = urlParam
	}
	if hashParam != "" {
		queryParams["hash"] = hashParam
	}

	url := config.CONST_Server + `/query/content`

	response, err := tool.Http_Get(url, queryParams)
	if err != nil {
		return nil, err
	}

	var responseObj FindContentUsingUrlResponse
	err = json.Unmarshal(response, &responseObj)
	if err != nil {
		return nil, err
	}

	if responseObj.ResultCode != core.CONST_ResultCode_Success {
		return nil, errors.New(responseObj.ResultMsg)
	}

	return &responseObj, nil
}
