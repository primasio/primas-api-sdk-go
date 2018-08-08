package content_interaction

import (
	"testing"

	"github.com/primasio/primas-api-sdk-go/core"
)

func TestGetShareMetadata(t *testing.T) {
	share_id := "5c7375693afdddd71a4c283bb34e6503ab549fdab8a6a5392a77a12ada7f75c2"
	resultGetShareMetadataResponse, err := GetShareMetadata(share_id, "")
	if err != nil {
		t.Errorf("GetContentMetadata error:%v", err.Error())
		return
	}

	if resultGetShareMetadataResponse != nil {
		t.Logf("GetContentMetadata response value:%v", resultGetShareMetadataResponse)
		if resultGetShareMetadataResponse.ResultCode != core.CONST_ResultCode_Success {
			t.Errorf("GetContentMetadata response error:%v", resultGetShareMetadataResponse.ResultMsg)
			return
		}

		if resultGetShareMetadataResponse.Data != nil {
			t.Logf("GetContentMetadata response data value:%#v", resultGetShareMetadataResponse.Data)
		} else {
			t.Logf("GetContentMetadata response data value don't find ")
		}
	}
}
