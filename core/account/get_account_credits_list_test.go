package account

import (
	"testing"

	"github.com/primasio/primas-api-sdk-go/core"
)

func TestGetAccountCreditsList(t *testing.T) {
	account_id := "809a85f7ddf8ae5aaa49fe30be10e07e09156dc04166fab98bbd7bb42b2dc26c"
	accountCredits, err := GetAccountCreditsList(account_id)
	if err != nil {
		t.Errorf("TestGetAccountCreditsList error:%v", err.Error())
		return
	}

	if accountCredits != nil {
		t.Logf("TestGetAccountCreditsList response value:%v", accountCredits)
		if accountCredits.ResultCode != core.CONST_ResultCode_Success {
			t.Errorf("TestGetAccountCreditsList response error:%v", accountCredits.ResultMsg)
			return
		}
		if accountCredits.Data != nil {
			t.Logf("TestGetAccountCreditsList response value:%v", accountCredits.Data)
		} else {
			t.Logf("TestGetAccountCreditsList response value don't find ")
		}
	}
}

func TestGetSubAccountCreditsList(t *testing.T) {
	account_id := "32fc4139f7d0347ca9ea70d30caad45a5d90fc23aaefacedf6bff2746e2073f3"
	sub_account_id := "a_0001"

	subAccountCredits, err := GetSubAccountCreditsList(account_id, sub_account_id)
	if err != nil {
		t.Errorf("TestGetSubAccountCreditsList error:%v", err.Error())
		return
	}

	if subAccountCredits != nil {
		t.Logf("TestGetSubAccountCreditsList response value:%v", subAccountCredits)
		if subAccountCredits.ResultCode != core.CONST_ResultCode_Success {
			t.Errorf("TestGetSubAccountCreditsList response error:%v", subAccountCredits.ResultMsg)
			return
		}
		if subAccountCredits.Data != nil {
			t.Logf("TestGetSubAccountCreditsList response value:%#v", subAccountCredits.Data)
		} else {
			t.Logf("TestGetSubAccountCreditsList response value don't find ")
		}
	}
}
