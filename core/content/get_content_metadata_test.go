package content

import (
	"testing"

	"github.com/primasio/primas-api-sdk-go/core"
)

func TestGetContentMetadata(t *testing.T) {
	content_id := "0b0d27adf09b17e4511e210ddeec0cf3136cbc2d214d6d84a80fb37577957c08"
	resultGetContentMetadataResponse, err := GetContentMetadata(content_id)
	if err != nil {
		t.Errorf("GetContentMetadata error:%v", err.Error())
		return
	}

	if resultGetContentMetadataResponse != nil {
		t.Logf("GetContentMetadata response value:%v", resultGetContentMetadataResponse)
		if resultGetContentMetadataResponse.ResultCode != core.CONST_ResultCode_Success {
			t.Errorf("GetContentMetadata response error:%v", resultGetContentMetadataResponse.ResultMsg)
			return
		}

		if resultGetContentMetadataResponse.Data != nil {
			t.Logf("GetContentMetadata response AccountTokensData value:%#v", resultGetContentMetadataResponse.Data)
		} else {
			t.Logf("GetContentMetadata response AccountTokensData value don't find ")
		}
	}
}
