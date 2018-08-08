package content_interaction

import (
	"testing"

	"github.com/primasio/primas-api-sdk-go/core"
)

func TestGetCommentsOfGroupshare(t *testing.T) {
	share_id := "60e4388167b2cb963964409b87510c57f8ba5a39487a900efc29b677d62cf449"
	resultGetCommentsGroupshareResponse, err := GetCommentsOfGroupshare(share_id, "", 0, 20)
	if err != nil {
		t.Errorf("GetCommentsOfGroupshare error:%v", err.Error())
		return
	}

	if resultGetCommentsGroupshareResponse != nil {
		t.Logf("GetCommentsOfGroupshare response value:%v", resultGetCommentsGroupshareResponse)
		if resultGetCommentsGroupshareResponse.ResultCode != core.CONST_ResultCode_Success {
			t.Errorf("GetLikesOfGroupshare response error:%v", resultGetCommentsGroupshareResponse.ResultMsg)
			return
		}

		if resultGetCommentsGroupshareResponse.Data != nil {
			t.Logf("GetCommentsOfGroupshare response data value:%#v", resultGetCommentsGroupshareResponse.Data)
		} else {
			t.Logf("GetCommentsOfGroupshare response data value don't find ")
		}
	}
}
