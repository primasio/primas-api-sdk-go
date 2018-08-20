package system

import (
	"testing"

	"github.com/primasio/primas-api-sdk-go/core"
)

func TestGetSystemParameters(t *testing.T) {
	resultGetSystemParameters, err := GetSystemParameters()
	if err != nil {
		t.Errorf("GetSystemParameters error:%v", err.Error())
		return
	}

	if resultGetSystemParameters != nil {
		t.Logf("GetSystemParameters response value:%v", resultGetSystemParameters)
		if resultGetSystemParameters.ResultCode != core.CONST_ResultCode_Success {
			t.Errorf("GetSystemParameters response error:%v", resultGetSystemParameters.ResultMsg)
			return
		}

		if resultGetSystemParameters.Data != nil {
			t.Logf("GetSystemParameters response data value:%#v", resultGetSystemParameters.Data)
		} else {
			t.Logf("GetSystemParameters response data value don't find ")
		}
	}

	t.Logf("LockAmountContent:%v", resultGetSystemParameters.Data.LockAmountContent)
	t.Logf("LockPeriodContent:%v", resultGetSystemParameters.Data.LockPeriodContent)
	t.Logf("LockAmountGroupJoin:%v", resultGetSystemParameters.Data.LockAmountGroupJoin)
	t.Logf("LockAmountGroupCreate:%v", resultGetSystemParameters.Data.LockAmountGroupCreate)
	t.Logf("ConsumeAmountReport:%v", resultGetSystemParameters.Data.ConsumeAmountReport)
}
