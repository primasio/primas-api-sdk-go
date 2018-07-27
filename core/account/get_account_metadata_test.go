package account

import (
	"testing"

	"github.com/primasio/primas-api-sdk-go/core"
)

func TestGetAccountTokenMetadata(t *testing.T) {
	account_id := "809a85f7ddf8ae5aaa49fe30be10e07e09156dc04166fab98bbd7bb42b2dc26c"
	accountMetadata, err := GetAccountTokenMetadata(account_id)
	if err != nil {
		t.Errorf("GetAccountTokenMetadata error:%v", err.Error())
		return
	}

	if accountMetadata != nil {
		t.Logf("GetAccountTokensData response value:%v", accountMetadata)
		if accountMetadata.ResultCode != core.CONST_ResultCode_Success {
			t.Errorf("GetAccountTokensData response error:%v", accountMetadata.ResultMsg)
			return
		}
		if accountMetadata.Data != nil {
			t.Logf("GetAccountTokenMetadata response value:%v", accountMetadata.Data)
		} else {
			t.Logf("GetAccountTokenMetadata response value don't find ")
		}
	}
}

func TestGetSubAccountTokenMetadata(t *testing.T) {
	account_id := "809a85f7ddf8ae5aaa49fe30be10e07e09156dc04166fab98bbd7bb42b2dc26c"
	sub_account_id := "a1ad1adf1"
	subAccountMetadata, err := GetSubAccountTokenMetadata(account_id, sub_account_id)
	if err != nil {
		t.Errorf("GetAccountTokenMetadata error:%v", err.Error())
		return
	}

	if subAccountMetadata != nil {
		t.Logf("GetAccountTokensData response value:%v", subAccountMetadata)
		if subAccountMetadata.ResultCode != core.CONST_ResultCode_Success {
			t.Errorf("GetAccountTokensData response error:%v", subAccountMetadata.ResultMsg)
			return
		}
		if subAccountMetadata.Data != nil {
			t.Logf("GetAccountTokenMetadata response value:%v", subAccountMetadata.Data)
		} else {
			t.Logf("GetAccountTokenMetadata response value don't find ")
		}
	}
}
