package content_interaction

import (
	"testing"

	"github.com/primasio/primas-api-sdk-go/core"
)

func TestGetReplyComments(t *testing.T) {
	comment_id := "02b6218aeae6e6f72b25a1a0302c8edce7514a48b1211a0ce85bc01ad4de0f10"
	resultGetReplyComments, err := GetReplyComments(comment_id)
	if err != nil {
		t.Errorf("GetReplyComments error:%v", err.Error())
		return
	}

	if resultGetReplyComments != nil {
		t.Logf("GetReplyComments response value:%v", resultGetReplyComments)
		if resultGetReplyComments.ResultCode != core.CONST_ResultCode_Success {
			t.Errorf("GetReplyComments response error:%v", resultGetReplyComments.ResultMsg)
			return
		}

		if resultGetReplyComments.Data != nil {
			t.Logf("GetReplyComments response data value:%#v", resultGetReplyComments.Data)
		} else {
			t.Logf("GetReplyComments response data value don't find ")
		}
	}
}
