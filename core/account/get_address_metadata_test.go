package account

import (
	"testing"

	"github.com/primasio/primas-api-sdk-go/core"
)

func TestGetAddressMetadata(t *testing.T) {
	address := "0xA0970333D335266382960A320773A16EaaF8d2E2"
	addressMetadata, err := GetAddressMetadata(address)
	if err != nil {
		t.Errorf("GetAddressMetadata error:%v", err.Error())
		return
	}

	if addressMetadata != nil {
		t.Logf("GetAddressMetadata response value:%v", addressMetadata)
		if addressMetadata.ResultCode != core.CONST_ResultCode_Success {
			t.Errorf("GetAddressMetadata response error:%v", addressMetadata.ResultMsg)
			return
		}
		if addressMetadata.Data != nil {
			t.Logf("GetAddressMetadata response value:%v", addressMetadata.Data)
		} else {
			t.Logf("GetAddressMetadata response value don't find ")
		}
	}
}
