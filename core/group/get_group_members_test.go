package group

import (
	"testing"

	"github.com/primasio/primas-api-sdk-go/core"
)

func TestGetGroupMembers(t *testing.T) {
	group_id := "e1c3c4c08be873447fd3e949d7e044c217dc585625b6eaa60f3233eb68fde30c"
	resultGetGroupMembersResponse, err := GetGroupMembers(group_id, 0, 20, "")
	if err != nil {
		t.Errorf("GetGroupMembers error:%v", err.Error())
		return
	}

	if resultGetGroupMembersResponse != nil {
		t.Logf("GetGroupMembers response value:%v", resultGetGroupMembersResponse)
		if resultGetGroupMembersResponse.ResultCode != core.CONST_ResultCode_Success {
			t.Errorf("GetGroupMembers response error:%v", resultGetGroupMembersResponse.ResultMsg)
			return
		}

		if resultGetGroupMembersResponse.Data != nil {
			t.Logf("GetGroupMembers response data value:%#v", resultGetGroupMembersResponse.Data)
		} else {
			t.Logf("GetGroupMembers response data value don't find ")
		}
	}
}
