package token

import (
	"testing"

	"github.com/primasio/primas-api-sdk-go/core"
)

func TestGetTokenPreLockList(t *testing.T) {
	account_id := "809a85f7ddf8ae5aaa49fe30be10e07e09156dc04166fab98bbd7bb42b2dc26c"
	resultIncentivesListResponse, err := GetTokenPreLockList(account_id, 1527695509, 1597695509,
		0, 20, CONST_Token_PreLock_Type_Lock)
	if err != nil {
		t.Errorf("GetTokenPreLockList error:%v", err.Error())
		return
	}

	if resultIncentivesListResponse != nil {
		t.Logf("GetTokenPreLockList response value:%v", resultIncentivesListResponse)
		if resultIncentivesListResponse.ResultCode != core.CONST_ResultCode_Success {
			t.Errorf("GetTokenPreLockList response error:%v", resultIncentivesListResponse.ResultMsg)
			return
		}

		if resultIncentivesListResponse.Data != nil {
			t.Logf("GetTokenPreLockList response AccountTokensData value:%v", resultIncentivesListResponse.Data)
		} else {
			t.Logf("GetTokenPreLockList response AccountTokensData value don't find ")
		}
	}
}
