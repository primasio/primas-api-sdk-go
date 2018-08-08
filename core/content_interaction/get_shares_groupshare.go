package content_interaction

import (
	"encoding/json"
	"errors"

	"github.com/primasio/primas-api-sdk-go/config"
	"github.com/primasio/primas-api-sdk-go/core"
	"github.com/primasio/primas-api-sdk-go/core/dtcp/dtcpv1"
	"github.com/primasio/primas-api-sdk-go/core/tool"
)

type GetShareGroupshareResponse struct {
	core.Response
	Data []dtcpv1.Share `json:"data"`
}

func GetSharesOfGroupShare(share_id, account_id string, page, pageSize int) (*GetShareGroupshareResponse, error) {
	if share_id == "" {
		return nil, errors.New("share_id is empty")
	}

	queryParams := make(map[string]interface{}, 0)
	url := config.CONST_Server + `/shares/` + share_id + `/shares`

	if account_id != "" {
		queryParams["account_id"] = account_id
	}
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

	var responseObj GetShareGroupshareResponse
	err = json.Unmarshal(response, &responseObj)
	if err != nil {
		return nil, err
	}

	if responseObj.ResultCode != core.CONST_ResultCode_Success {
		return nil, errors.New(responseObj.ResultMsg)
	}

	return &responseObj, nil
}
