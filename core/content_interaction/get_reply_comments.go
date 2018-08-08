package content_interaction

import (
	"encoding/json"
	"errors"

	"github.com/primasio/primas-api-sdk-go/config"
	"github.com/primasio/primas-api-sdk-go/core"
	"github.com/primasio/primas-api-sdk-go/core/dtcp/dtcpv1"
	"github.com/primasio/primas-api-sdk-go/core/tool"
)

type GetReplyCommentsResponse struct {
	core.Response
	Data []dtcpv1.ContentCommentGet `json:"data"`
}

func GetReplyComments(comment_id string) (*GetReplyCommentsResponse, error) {
	if comment_id == "" {
		return nil, errors.New("share_id is empty")
	}

	queryParams := make(map[string]interface{}, 0)
	url := config.CONST_Server + `/comments/` + comment_id + `/reply`

	response, err := tool.Http_Get(url, queryParams)
	if err != nil {
		return nil, err
	}

	var responseObj GetReplyCommentsResponse
	err = json.Unmarshal(response, &responseObj)
	if err != nil {
		return nil, err
	}

	if responseObj.ResultCode != core.CONST_ResultCode_Success {
		return nil, errors.New(responseObj.ResultMsg)
	}

	return &responseObj, nil
}
