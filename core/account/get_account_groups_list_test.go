package account

import (
	"testing"

	"github.com/primasio/primas-api-sdk-go/core"
)

func TestGetAccountGroupsList(t *testing.T) {
	account_id := "809a85f7ddf8ae5aaa49fe30be10e07e09156dc04166fab98bbd7bb42b2dc26c"
	resultAccountGroupListResponse, err := GetAccountGroupsList(account_id, 0, 20)
	if err != nil {
		t.Errorf("GetAccountGroupsList error:%v", err.Error())
		return
	}

	if resultAccountGroupListResponse != nil {
		t.Logf("GetAccountGroupsList response value:%v", resultAccountGroupListResponse)
		if resultAccountGroupListResponse.ResultCode != core.CONST_ResultCode_Success {
			t.Errorf("GetIncentivesList response error:%v", resultAccountGroupListResponse.ResultMsg)
			return
		}

		if resultAccountGroupListResponse.Data != nil {
			t.Logf("GetAccountGroupsList response AccountTokensData value:%#v", resultAccountGroupListResponse.Data)
		} else {
			t.Logf("GetAccountGroupsList response AccountTokensData value don't find ")
		}
	}
}

func TestGetSubAccountGroupsList(t *testing.T) {
	account_id := "809a85f7ddf8ae5aaa49fe30be10e07e09156dc04166fab98bbd7bb42b2dc26c"
	sub_account_id := "a_0001"
	resultAccountGroupListResponse, err := GetSubAccountGroupsList(account_id, sub_account_id, 0, 20)
	if err != nil {
		t.Errorf("GetSubAccountGroupsList error:%v", err.Error())
		return
	}

	if resultAccountGroupListResponse != nil {
		t.Logf("GetSubAccountGroupsList response value:%v", resultAccountGroupListResponse)
		if resultAccountGroupListResponse.ResultCode != core.CONST_ResultCode_Success {
			t.Errorf("GetSubAccountGroupsList response error:%v", resultAccountGroupListResponse.ResultMsg)
			return
		}

		if resultAccountGroupListResponse.Data != nil {
			t.Logf("GetSubAccountGroupsList response AccountTokensData value:%#v", resultAccountGroupListResponse.Data)
		} else {
			t.Logf("GetSubAccountGroupsList response AccountTokensData value don't find ")
		}
	}
}
