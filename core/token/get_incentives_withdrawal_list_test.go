package token

import (
	"testing"

	"github.com/primasio/primas-api-sdk-go/core"
)

func TestGetIncentivesWithdrawalList(t *testing.T) {
	account_id := "6adb0bc07b6ec4d992d5e0c051249f61024629a0bb5b264156463c788d9dc661"
	resultIncentivesWithdrawalList, err := GetIncentivesWithdrawalList(account_id, 0, 1534485897,
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
