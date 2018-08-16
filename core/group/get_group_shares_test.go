package group

import (
	"testing"

	"github.com/primasio/primas-api-sdk-go/core"
)

func TestGetGroupShares(t *testing.T) {
	group_id := "e1c3c4c08be873447fd3e949d7e044c217dc585625b6eaa60f3233eb68fde30c"
	resultGetGroupSharesResponse, err := GetGroupShares(group_id, 0, 20, "")
	if err != nil {
		t.Errorf("GetGroupShares error:%v", err.Error())
		return
	}

	if resultGetGroupSharesResponse != nil {
		t.Logf("GetGroupShares response value:%v", resultGetGroupSharesResponse)
		if resultGetGroupSharesResponse.ResultCode != core.CONST_ResultCode_Success {
			t.Errorf("GetGroupShares response error:%v", resultGetGroupSharesResponse.ResultMsg)
			return
		}

		if resultGetGroupSharesResponse.Data != nil {
			t.Logf("GetGroupShares response data value:%#v", resultGetGroupSharesResponse.Data)
		} else {
			t.Logf("GetGroupShares response data value don't find ")
		}
	}
}
