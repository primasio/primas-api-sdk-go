package query

import (
	"testing"

	"github.com/primasio/primas-api-sdk-go/core"
)

func TestGetAll(t *testing.T) {
	page := 0
	page_size := 20
	text := "dtcp"
	qtype := "all"
	category := ""

	resultGetAll, err := GetAll(page, page_size, text, qtype, category)
	if err != nil {
		t.Errorf("GetAll error:%v", err.Error())
		return
	}

	if resultGetAll != nil {
		t.Logf("GetAll response value:%v", resultGetAll)
		if resultGetAll.ResultCode != core.CONST_ResultCode_Success {
			t.Errorf("GetAll response error:%v", resultGetAll.ResultMsg)
			return
		}

		if resultGetAll.Data != nil {
			t.Logf("GetAll response data value:%#v", resultGetAll.Data)
		} else {
			t.Logf("GetAll response data value don't find ")
		}
	}
}
