package account

import (
	"errors"

	"github.com/primasio/primas-api-sdk-go/config"
	"github.com/primasio/primas-api-sdk-go/core/tool"
)

func GetAccountAvatarRawImage(account_id string) ([]byte, error) {
	if account_id == "" {
		return nil, errors.New("account_id is empty")
	}

	url := config.CONST_Server + `/accounts/` + account_id + `/avatar/raw`

	reValues, err := tool.Http_Get(url, nil)
	if err != nil {
		return nil, err
	}

	return reValues, nil
}

func GetSubAccountAvatarRawImage(account_id, sub_account_id string) ([]byte, error) {
	if account_id == "" {
		return nil, errors.New("account_id is empty")
	}

	if sub_account_id == "" {
		return nil, errors.New("sub_account_id is empty")
	}

	url := config.CONST_Server + `/accounts/` + account_id + `/sub/` + sub_account_id + `/avatar/raw`

	reValues, err := tool.Http_Get(url, nil)
	if err != nil {
		return nil, err
	}

	return reValues, nil
}
