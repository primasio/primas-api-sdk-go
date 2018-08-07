package content

import (
	"errors"
	"net/http"

	"github.com/primasio/primas-api-sdk-go/config"
	"github.com/primasio/primas-api-sdk-go/core/tool"
)

func GetRawContent(content_id string) (*http.Response, error) {
	if content_id == "" {
		return nil, errors.New("content_id is empty")
	}

	url := config.CONST_Server + `/content/` + content_id + `/raw`
	queryParams := make(map[string]interface{}, 0)

	response, err := tool.Http_Get_Direct(url, queryParams)
	if err != nil {
		return nil, err
	}

	return response, nil
}
