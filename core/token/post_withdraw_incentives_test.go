package token

import (
	"testing"

	"github.com/primasio/primas-api-sdk-go/core"
	_ "github.com/primasio/primas-api-sdk-go/core/tool"
	"github.com/shopspring/decimal"
)

// todo not complete
func TestPostWithdrawIncentives(t *testing.T) {
	account_id := "809a85f7ddf8ae5aaa49fe30be10e07e09156dc04166fab98bbd7bb42b2dc26c"
	node_id := "58f47077984e5daa4d2ea46f2e689177a1655c1321544e69f851530a789e9fd7"
	created := 1532525161
	amount := decimal.New(123, 0)
	node_fee := decimal.New(123, 0)

	resultWithdrawIncetice, err := PostWithdrawIncentives(account_id, node_id, created, amount, node_fee)
	if err != nil {
		t.Errorf("PostWithdrawIncentives error:%v", err.Error())
		return
	}

	if resultWithdrawIncetice != nil {
		t.Logf("PostWithdrawIncentives response value:%v", resultWithdrawIncetice)
		if resultWithdrawIncetice.ResultCode != core.CONST_ResultCode_Success {
			t.Errorf("PostWithdrawIncentives response error:%v", resultWithdrawIncetice.ResultMsg)
			return
		}

		if resultWithdrawIncetice.Data != nil {
			t.Logf("PostWithdrawIncentives response AccountTokensData value:%v", resultWithdrawIncetice.Data)
		} else {
			t.Logf("PostWithdrawIncentives response AccountTokensData value don't find ")
		}
	}
}
