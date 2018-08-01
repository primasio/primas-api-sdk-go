package account

import (
	"testing"

	"github.com/primasio/primas-api-sdk-go/core"
)

func TestGetAccountContentList(t *testing.T) {
	account_id := "809a85f7ddf8ae5aaa49fe30be10e07e09156dc04166fab98bbd7bb42b2dc26c"
	resultAccountGroupListResponse, err := GetAccountContentList(account_id, 0, 20)
	if err != nil {
		t.Errorf("GetAccountContentList error:%v", err.Error())
		return
	}

	if resultAccountGroupListResponse != nil {
		t.Logf("GetAccountContentList response value:%v", resultAccountGroupListResponse)
		if resultAccountGroupListResponse.ResultCode != core.CONST_ResultCode_Success {
			t.Errorf("GetAccountContentList response error:%v", resultAccountGroupListResponse.ResultMsg)
			return
		}

		if resultAccountGroupListResponse.Data != nil {
			t.Logf("GetAccountContentList response AccountTokensData value:%#v", resultAccountGroupListResponse.Data)
		} else {
			t.Logf("GetAccountContentList response AccountTokensData value don't find ")
		}
	}
}
