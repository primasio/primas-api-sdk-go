package account

import (
	"testing"

	"github.com/primasio/primas-api-sdk-go/core"
)

func TestGetAccountAvatarMetadata(t *testing.T) {
	account_id := "809a85f7ddf8ae5aaa49fe30be10e07e09156dc04166fab98bbd7bb42b2dc26c"
	resultAccountAvatarResponse, err := GetAccountAvatarMetadata(account_id)
	if err != nil {
		t.Errorf("GetAccountComments error:%v", err.Error())
		return
	}

	if resultAccountAvatarResponse != nil {
		t.Logf("GetAccountAvatarMetadata response value:%v", resultAccountAvatarResponse)
		if resultAccountAvatarResponse.ResultCode != core.CONST_ResultCode_Success {
			t.Errorf("GetAccountAvatarMetadata response error:%v", resultAccountAvatarResponse.ResultMsg)
			return
		}

		if resultAccountAvatarResponse.Data != nil {
			t.Logf("GetAccountAvatarMetadata response AccountTokensData value:%#v", resultAccountAvatarResponse.Data)
		} else {
			t.Logf("GetAccountAvatarMetadata response AccountTokensData value don't find ")
		}
	}
}

func TestGetSubAccountAvatarMetadata(t *testing.T) {
	account_id := "809a85f7ddf8ae5aaa49fe30be10e07e09156dc04166fab98bbd7bb42b2dc26c"
	sub_account_id := "a_0001"
	resultAccountAvatarResponse, err := GetSubAccountAvatarMetadata(account_id, sub_account_id)
	if err != nil {
		t.Errorf("GetSubAccountAvatarMetadata error:%v", err.Error())
		return
	}

	if resultAccountAvatarResponse != nil {
		t.Logf("GetSubAccountAvatarMetadata response value:%v", resultAccountAvatarResponse)
		if resultAccountAvatarResponse.ResultCode != core.CONST_ResultCode_Success {
			t.Errorf("GetSubAccountAvatarMetadata response error:%v", resultAccountAvatarResponse.ResultMsg)
			return
		}

		if resultAccountAvatarResponse.Data != nil {
			t.Logf("GetSubAccountAvatarMetadata response AccountTokensData value:%#v", resultAccountAvatarResponse.Data)
		} else {
			t.Logf("GetSubAccountAvatarMetadata response AccountTokensData value don't find ")
		}
	}
}
