package token

import (
	"fmt"
	"testing"
	"time"

	"github.com/primasio/primas-api-sdk-go/core"
)

func TestGetAccountTokensData(t *testing.T) {
	fmt.Println(time.Now().UnixNano() / 1e6)

	account_id := "809a85f7ddf8ae5aaa49fe30be10e07e09156dc04166fab98bbd7bb42b2dc26c"
	resultAccountTokensResponse, err := GetAccountTokensData(account_id)
	if err != nil {
		t.Errorf("GetAccountTokensData error:%v", err.Error())
		return
	}

	if resultAccountTokensResponse != nil {
		t.Logf("GetAccountTokensData response value:%v", resultAccountTokensResponse)
		if resultAccountTokensResponse.ResultCode != core.CONST_ResultCode_Success {
			t.Errorf("GetAccountTokensData response error:%v", resultAccountTokensResponse.ResultMsg)
			return
		}
		if resultAccountTokensResponse.Data != nil {
			t.Logf("GetAccountTokensData response AccountTokensData value:%v", resultAccountTokensResponse.Data)
		} else {
			t.Logf("GetAccountTokensData response AccountTokensData value don't find ")
		}
	}

}
