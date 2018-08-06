package account

import (
	"testing"
)

func TestGetAccountAvatarRawImage(t *testing.T) {
	account_id := "809a85f7ddf8ae5aaa49fe30be10e07e09156dc04166fab98bbd7bb42b2dc26c"
	resultReponse, err := GetAccountAvatarRawImage(account_id)
	if err != nil {
		t.Errorf("GetAccountAvatarRawImage error:%v", err.Error())
		return
	}

	t.Logf("resultReponse:" + string(resultReponse))
}

func TestGetSubAccountAvatarRawImage(t *testing.T) {
	account_id := "809a85f7ddf8ae5aaa49fe30be10e07e09156dc04166fab98bbd7bb42b2dc26c"
	sub_account_id := "a_0001"

	resultReponse, err := GetSubAccountAvatarRawImage(account_id, sub_account_id)
	if err != nil {
		t.Errorf("GetSubAccountAvatarRawImage error:%v", err.Error())
		return
	}

	t.Logf("resultReponse:" + string(resultReponse))
}
