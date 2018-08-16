package group

import (
	"log"
	"testing"
	"time"

	"github.com/primasio/primas-api-sdk-go/core"
	"github.com/primasio/primas-api-sdk-go/core/tool"
)

func TestPostAddGroupMemberWhitelist(t *testing.T) {
	src_id := `32fc4139f7d0347ca9ea70d30caad45a5d90fc23aaefacedf6bff2746e2073f3`
	dest_id := `53687b42c97eb2dc38eac2483b7623d17ad4337dbec1cc54369cc8f26b52a71d`
	account_id := `32fc4139f7d0347ca9ea70d30caad45a5d90fc23aaefacedf6bff2746e2073f3`
	sub_account_id := ``
	sub_account_name := ``
	created := int(time.Now().Unix())
	application_status := "pending"

	signature, preObj, err := PostAddGroupMemberWhitelist_SignatureStr(src_id, dest_id, account_id, sub_account_id,
		sub_account_name, created, application_status)
	if err != nil {
		t.Errorf("TestPostAddGroupMemberWhitelist error:%v", err.Error())
		return
	}
	if preObj == nil {
		t.Errorf("TestPostAddGroupMemberWhitelist preObj object is nil")
		return
	}
	if signature == "" {
		t.Errorf("TestPostAddGroupMemberWhitelist signature value is empty")
		return
	}

	log.Println("signature:", signature)

	// mock Sign
	privateKey := tool.GetClientPrivateKey()
	signValue, err := tool.Sign([]byte(signature), privateKey)
	if err != nil {
		t.Errorf("Sign error %v:", err.Error())
		return
	}
	//

	postAddGroupWhite, err := PostAddGroupMemberWhitelist(signValue, preObj)
	if err != nil {
		t.Errorf("TestPostAddGroupMemberWhitelist error:%v", err.Error())
		return
	}

	if postAddGroupWhite != nil {
		t.Logf("TestPostAddGroupMemberWhitelist response value:%v", postAddGroupWhite)
		if postAddGroupWhite.ResultCode != core.CONST_ResultCode_Success {
			t.Errorf("TestPostAddGroupMemberWhitelist response error:%v", postAddGroupWhite.ResultMsg)
			return
		}
		if postAddGroupWhite.Data != nil {
			t.Logf("TestPostAddGroupMemberWhitelist response value:%v", postAddGroupWhite.Data)
		} else {
			t.Logf("TestPostAddGroupMemberWhitelist response value don't find ")
		}
	}
}
