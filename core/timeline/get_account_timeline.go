package timeline

import (
	"encoding/json"
	"errors"

	"github.com/primasio/primas-api-sdk-go/config"
	"github.com/primasio/primas-api-sdk-go/core"
	"github.com/primasio/primas-api-sdk-go/core/dtcp/dtcpv1"
	"github.com/primasio/primas-api-sdk-go/core/tool"
)

type GetAccountTimelineResponse struct {
	core.Response
	Data []dtcpv1.Share `json:"data"`
}

func GetAccountTimeline(account_id string, page, pageSize int) (*GetAccountTimelineResponse, error) {
	if account_id == "" {
		return nil, errors.New("account_id is empty")
	}

	queryParams := make(map[string]interface{}, 0)
	if page > 0 {
		queryParams["page"] = page
	}
	if pageSize > 0 {
		queryParams["page_size"] = pageSize
	}

	url := config.CONST_Server + `/accounts/` + account_id + `/timeline`

	response, err := tool.Http_Get(url, queryParams)
	if err != nil {
		return nil, err
	}

	var responseObj GetAccountTimelineResponse
	err = json.Unmarshal(response, &responseObj)
	if err != nil {
		return nil, err
	}

	if responseObj.ResultCode != core.CONST_ResultCode_Success {
		return nil, errors.New(responseObj.ResultMsg)
	}

	return &responseObj, nil
}
