package content_interaction

import (
	"testing"

	"github.com/primasio/primas-api-sdk-go/core"
)

func TestGetLikesOfGroupshare(t *testing.T) {
	share_id := "7187ddc476076fce73201ba91d20600e46b8f2d18d828fa7438c2bbd536ba115"
	resultGetLikesOfGroupshare, err := GetLikesOfGroupshare(share_id, "", 0, 20)
	if err != nil {
		t.Errorf("GetLikesOfGroupshare error:%v", err.Error())
		return
	}

	if resultGetLikesOfGroupshare != nil {
		t.Logf("GetLikesOfGroupshare response value:%v", resultGetLikesOfGroupshare)
		if resultGetLikesOfGroupshare.ResultCode != core.CONST_ResultCode_Success {
			t.Errorf("GetLikesOfGroupshare response error:%v", resultGetLikesOfGroupshare.ResultMsg)
			return
		}

		if resultGetLikesOfGroupshare.Data != nil {
			t.Logf("GetLikesOfGroupshare response data value:%#v", resultGetLikesOfGroupshare.Data)
		} else {
			t.Logf("GetLikesOfGroupshare response data value don't find ")
		}
	}
}
