package group

import (
	"testing"

	"github.com/primasio/primas-api-sdk-go/core"
)

func TestGetGroupMetadata(t *testing.T) {
	group_id := "550b81d5d34a7f211897f8ed04ae13c96c0093b0b16645a8d7ee36edbaa4db6b"
	resultGetGroupMetadataResponse, err := GetGroupMetadata(group_id, "")
	if err != nil {
		t.Errorf("GetGroupMetadata error:%v", err.Error())
		return
	}

	if resultGetGroupMetadataResponse != nil {
		t.Logf("GetGroupMetadata response value:%v", resultGetGroupMetadataResponse)
		if resultGetGroupMetadataResponse.ResultCode != core.CONST_ResultCode_Success {
			t.Errorf("GetGroupMetadata response error:%v", resultGetGroupMetadataResponse.ResultMsg)
			return
		}

		if resultGetGroupMetadataResponse.Data != nil {
			t.Logf("GetGroupMetadata response data value:%#v", resultGetGroupMetadataResponse.Data)
		} else {
			t.Logf("GetGroupMetadata response data value don't find ")
		}
	}
}
