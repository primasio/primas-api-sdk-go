package group

import (
	"errors"

	"github.com/primasio/primas-api-sdk-go/config"
	"github.com/primasio/primas-api-sdk-go/core/tool"
)

func GetGroupAvatarRawImage(group_id string) ([]byte, error) {
	if group_id == "" {
		return nil, errors.New("group_id is empty")
	}

	url := config.CONST_Server + `/groups/` + group_id + `/avatar/raw`

	reValues, err := tool.Http_Get(url, nil)
	if err != nil {
		return nil, err
	}

	return reValues, nil
}
