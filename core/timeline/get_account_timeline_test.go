package timeline

import (
	"testing"

	"github.com/primasio/primas-api-sdk-go/core"
)

func TestGetAccountTimeline(t *testing.T) {
	account_id := "2def1e68569f09439f2b5ad2db5ef15196411fc8d16829e02e78af20c6fd8d6c"
	resultGetAccountTimeline, err := GetAccountTimeline(account_id, 0, 20)
	if err != nil {
		t.Errorf("GetAccountTimeline error:%v", err.Error())
		return
	}

	if resultGetAccountTimeline != nil {
		t.Logf("GetAccountTimeline response value:%v", resultGetAccountTimeline)
		if resultGetAccountTimeline.ResultCode != core.CONST_ResultCode_Success {
			t.Errorf("GetAccountTimeline response error:%v", resultGetAccountTimeline.ResultMsg)
			return
		}

		if resultGetAccountTimeline.Data != nil {
			t.Logf("GetAccountTimeline response data value:%#v", resultGetAccountTimeline.Data)
		} else {
			t.Logf("GetAccountTimeline response data value don't find ")
		}
	}
}
