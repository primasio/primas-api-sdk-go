package account

import (
	"testing"

	"github.com/primasio/primas-api-sdk-go/core"
)

func TestGetAccountGroupApplications(t *testing.T) {
	account_id := "809a85f7ddf8ae5aaa49fe30be10e07e09156dc04166fab98bbd7bb42b2dc26c"
	resultAccountGroupAppsResponse, err := GetAccountGroupApplications(account_id, 0, 20)
	if err != nil {
		t.Errorf("GetAccountGroupApplications error:%v", err.Error())
		return
	}

	if resultAccountGroupAppsResponse != nil {
		t.Logf("GetAccountGroupApplications response value:%v", resultAccountGroupAppsResponse)
		if resultAccountGroupAppsResponse.ResultCode != core.CONST_ResultCode_Success {
			t.Errorf("GetAccountGroupApplications response error:%v", resultAccountGroupAppsResponse.ResultMsg)
			return
		}

		if resultAccountGroupAppsResponse.Data != nil {
			t.Logf("GetAccountGroupApplications response AccountTokensData value:%#v", resultAccountGroupAppsResponse.Data)
		} else {
			t.Logf("GetAccountGroupApplications response AccountTokensData value don't find ")
		}
	}
}

func TestGetSubAccountGroupApplications(t *testing.T) {
	account_id := "809a85f7ddf8ae5aaa49fe30be10e07e09156dc04166fab98bbd7bb42b2dc26c"
	sub_account_id := "a_0001"
	resultAccountGroupAppsResponse, err := GetSubAccountGroupApplications(account_id, sub_account_id, 0, 20)
	if err != nil {
		t.Errorf("GetSubAccountGroupApplications error:%v", err.Error())
		return
	}

	if resultAccountGroupAppsResponse != nil {
		t.Logf("GetSubAccountGroupApplications response value:%v", resultAccountGroupAppsResponse)
		if resultAccountGroupAppsResponse.ResultCode != core.CONST_ResultCode_Success {
			t.Errorf("GetSubAccountGroupApplications response error:%v", resultAccountGroupAppsResponse.ResultMsg)
			return
		}

		if resultAccountGroupAppsResponse.Data != nil {
			t.Logf("GetSubAccountGroupApplications response AccountTokensData value:%#v", resultAccountGroupAppsResponse.Data)
		} else {
			t.Logf("GetSubAccountGroupApplications response AccountTokensData value don't find ")
		}
	}
}
