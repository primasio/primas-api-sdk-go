package account

import (
	"testing"

	"github.com/primasio/primas-api-sdk-go/core"
)

func TestGetAccountLikes(t *testing.T) {
	account_id := "809a85f7ddf8ae5aaa49fe30be10e07e09156dc04166fab98bbd7bb42b2dc26c"
	resultAccountLikesResponse, err := GetAccountLikes(account_id, 0, 20)
	if err != nil {
		t.Errorf("GetAccountLikes error:%v", err.Error())
		return
	}

	if resultAccountLikesResponse != nil {
		t.Logf("GetAccountLikes response value:%v", resultAccountLikesResponse)
		if resultAccountLikesResponse.ResultCode != core.CONST_ResultCode_Success {
			t.Errorf("GetAccountLikes response error:%v", resultAccountLikesResponse.ResultMsg)
			return
		}

		if resultAccountLikesResponse.Data != nil {
			t.Logf("GetAccountLikes response AccountTokensData value:%#v", resultAccountLikesResponse.Data)
		} else {
			t.Logf("GetAccountLikes response AccountTokensData value don't find ")
		}
	}
}

func TestGetSubAccountLikes(t *testing.T) {
	account_id := "809a85f7ddf8ae5aaa49fe30be10e07e09156dc04166fab98bbd7bb42b2dc26c"
	sub_account_id := "a_0001"
	resultAccountLikesResponse, err := GetSubAccountLikes(account_id, sub_account_id, 0, 20)
	if err != nil {
		t.Errorf("GetSubAccountLikes error:%v", err.Error())
		return
	}

	if resultAccountLikesResponse != nil {
		t.Logf("GetSubAccountLikes response value:%v", resultAccountLikesResponse)
		if resultAccountLikesResponse.ResultCode != core.CONST_ResultCode_Success {
			t.Errorf("GetSubAccountLikes response error:%v", resultAccountLikesResponse.ResultMsg)
			return
		}

		if resultAccountLikesResponse.Data != nil {
			t.Logf("GetSubAccountLikes response AccountTokensData value:%#v", resultAccountLikesResponse.Data)
		} else {
			t.Logf("GetSubAccountLikes response AccountTokensData value don't find ")
		}
	}
}
