package account

import (
	"testing"

	"github.com/primasio/primas-api-sdk-go/core"
)

func TestGetAccountSharesInSingleGroup(t *testing.T) {
	account_id := "809a85f7ddf8ae5aaa49fe30be10e07e09156dc04166fab98bbd7bb42b2dc26c"
	group_id := "550b81d5d34a7f211897f8ed04ae13c96c0093b0b16645a8d7ee36edbaa4db6b"
	resultAccountSharesInSingleGroup, err := GetAccountSharesInSingleGroup(account_id, group_id, 0, 20)
	if err != nil {
		t.Errorf("GetAccountSharesInSingleGroup error:%v", err.Error())
		return
	}

	if resultAccountSharesInSingleGroup != nil {
		t.Logf("GetAccountSharesInSingleGroup response value:%v", resultAccountSharesInSingleGroup)
		if resultAccountSharesInSingleGroup.ResultCode != core.CONST_ResultCode_Success {
			t.Errorf("GetAccountSharesInSingleGroup response error:%v", resultAccountSharesInSingleGroup.ResultMsg)
			return
		}

		if resultAccountSharesInSingleGroup.Data != nil {
			t.Logf("GetAccountSharesInSingleGroup response AccountTokensData value:%#v", resultAccountSharesInSingleGroup.Data)
		} else {
			t.Logf("GetAccountSharesInSingleGroup response AccountTokensData value don't find ")
		}
	}
}

func TestGetSubAccountSharesInSingleGroup(t *testing.T) {
	account_id := "809a85f7ddf8ae5aaa49fe30be10e07e09156dc04166fab98bbd7bb42b2dc26c"
	group_id := "550b81d5d34a7f211897f8ed04ae13c96c0093b0b16645a8d7ee36edbaa4db6b"
	sub_account_id := "a_0001"
	resultAccountSharesInSingleGroup, err := GetSubAccountSharesInSingleGroup(account_id, sub_account_id, group_id, 0, 20)
	if err != nil {
		t.Errorf("GetSubAccountSharesInSingleGroup error:%v", err.Error())
		return
	}

	if resultAccountSharesInSingleGroup != nil {
		t.Logf("GetSubAccountSharesInSingleGroup response value:%v", resultAccountSharesInSingleGroup)
		if resultAccountSharesInSingleGroup.ResultCode != core.CONST_ResultCode_Success {
			t.Errorf("GetSubAccountSharesInSingleGroup response error:%v", resultAccountSharesInSingleGroup.ResultMsg)
			return
		}

		if resultAccountSharesInSingleGroup.Data != nil {
			t.Logf("GetSubAccountSharesInSingleGroup response AccountTokensData value:%#v", resultAccountSharesInSingleGroup.Data)
		} else {
			t.Logf("GetSubAccountSharesInSingleGroup response AccountTokensData value don't find ")
		}
	}
}
