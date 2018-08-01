package account

import (
	"testing"

	"github.com/primasio/primas-api-sdk-go/core"
)

func TestGetAccountSharesList(t *testing.T) {
	account_id := "809a85f7ddf8ae5aaa49fe30be10e07e09156dc04166fab98bbd7bb42b2dc26c"
	resultAccountGroupListResponse, err := GetAccountSharesList(account_id, 0, 20)
	if err != nil {
		t.Errorf("GetAccountSharesList error:%v", err.Error())
		return
	}

	if resultAccountGroupListResponse != nil {
		t.Logf("GetAccountSharesList response value:%v", resultAccountGroupListResponse)
		if resultAccountGroupListResponse.ResultCode != core.CONST_ResultCode_Success {
			t.Errorf("GetAccountSharesList response error:%v", resultAccountGroupListResponse.ResultMsg)
			return
		}

		if resultAccountGroupListResponse.Data != nil {
			t.Logf("GetAccountSharesList response AccountTokensData value:%#v", resultAccountGroupListResponse.Data)
		} else {
			t.Logf("GetAccountSharesList response AccountTokensData value don't find ")
		}
	}
}

func TestGetSubAccountSharesList(t *testing.T) {
	account_id := "809a85f7ddf8ae5aaa49fe30be10e07e09156dc04166fab98bbd7bb42b2dc26c"
	sub_account_id := "a_0001"
	resultAccountGroupListResponse, err := GetSubAccountSharesList(account_id, sub_account_id, 0, 20)
	if err != nil {
		t.Errorf("GetSubAccountSharesList error:%v", err.Error())
		return
	}

	if resultAccountGroupListResponse != nil {
		t.Logf("GetSubAccountSharesList response value:%v", resultAccountGroupListResponse)
		if resultAccountGroupListResponse.ResultCode != core.CONST_ResultCode_Success {
			t.Errorf("GetSubAccountSharesList response error:%v", resultAccountGroupListResponse.ResultMsg)
			return
		}

		if resultAccountGroupListResponse.Data != nil {
			t.Logf("GetSubAccountSharesList response AccountTokensData value:%#v", resultAccountGroupListResponse.Data)
		} else {
			t.Logf("GetSubAccountSharesList response AccountTokensData value don't find ")
		}
	}
}
