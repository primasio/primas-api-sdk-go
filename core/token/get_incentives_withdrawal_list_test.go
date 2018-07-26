package token

import (
	"testing"

	"github.com/primasio/primas-api-sdk-go/core"
)

func TestGetIncentivesWithdrawalList(t *testing.T) {
	account_id := "809a85f7ddf8ae5aaa49fe30be10e07e09156dc04166fab98bbd7bb42b2dc26c"
	resultIncentivesWithdrawalList, err := GetIncentivesWithdrawalList(account_id, 1527695509, 1597695509,
		0, 20, CONST_Incetives_Withdrawal_Status_Pending)
	if err != nil {
		t.Errorf("GetIncentivesWithdrawalList error:%v", err.Error())
		return
	}

	if resultIncentivesWithdrawalList != nil {
		t.Logf("GetIncentivesWithdrawalList response value:%v", resultIncentivesWithdrawalList)
		if resultIncentivesWithdrawalList.ResultCode != core.CONST_ResultCode_Success {
			t.Errorf("GetIncentivesWithdrawalList response error:%v", resultIncentivesWithdrawalList.ResultMsg)
			return
		}

		if resultIncentivesWithdrawalList.Data != nil {
			t.Logf("GetIncentivesWithdrawalList response AccountTokensData value:%v", resultIncentivesWithdrawalList.Data)
		} else {
			t.Logf("GetIncentivesWithdrawalList response AccountTokensData value don't find ")
		}
	}
}
