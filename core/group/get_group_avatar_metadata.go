package group

import (
	"encoding/json"
	"errors"

	"github.com/primasio/primas-api-sdk-go/config"
	"github.com/primasio/primas-api-sdk-go/core"
	"github.com/primasio/primas-api-sdk-go/core/dtcp/dtcpv1"
	"github.com/primasio/primas-api-sdk-go/core/tool"
)

type GetGroupContentResponse struct {
	core.Response
	Data *dtcpv1.ContentGet `json:"data"`
}

func GetGroupAvatarMetadata(group_id string) (*GetGroupContentResponse, error) {
	if group_id == "" {
		return nil, errors.New("content_id is empty")
	}

	url := config.CONST_Server + `/groups/` + group_id + `/avatar`
	queryParams := make(map[string]interface{}, 0)

	response, err := tool.Http_Get(url, queryParams)
	if err != nil {
		return nil, err
	}

	var responseObj GetGroupContentResponse
	err = json.Unmarshal(response, &responseObj)
	if err != nil {
		return nil, err
	}

	if responseObj.ResultCode != core.CONST_ResultCode_Success {
		return nil, errors.New(responseObj.ResultMsg)
	}

	return &responseObj, nil
}
