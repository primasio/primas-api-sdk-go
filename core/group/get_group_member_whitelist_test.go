package group

import (
	"testing"

	"github.com/primasio/primas-api-sdk-go/core"
)

func TestGetGroupMemberWhitelist(t *testing.T) {
	group_id := "53687b42c97eb2dc38eac2483b7623d17ad4337dbec1cc54369cc8f26b52a71d"
	resultGetGroupMemberWhitelist, err := GetGroupMemberWhitelist(group_id, 0, 20, "")
	if err != nil {
		t.Errorf("GetGroupMemberWhitelist error:%v", err.Error())
		return
	}

	if resultGetGroupMemberWhitelist != nil {
		t.Logf("GetGroupMemberWhitelist response value:%v", resultGetGroupMemberWhitelist)
		if resultGetGroupMemberWhitelist.ResultCode != core.CONST_ResultCode_Success {
			t.Errorf("GetGroupMemberWhitelist response error:%v", resultGetGroupMemberWhitelist.ResultMsg)
			return
		}

		if resultGetGroupMemberWhitelist.Data != nil {
			t.Logf("GetGroupMemberWhitelist response data value:%#v", resultGetGroupMemberWhitelist.Data)
		} else {
			t.Logf("GetGroupMemberWhitelist response data value don't find ")
		}
	}
}
