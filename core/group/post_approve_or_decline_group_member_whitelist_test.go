package group

import (
	"log"
	"testing"
	"time"

	"github.com/primasio/primas-api-sdk-go/core"
	"github.com/primasio/primas-api-sdk-go/core/tool"
)

func TestPostApproveOrDeclineGroupMemberWhitelist(t *testing.T) {
	parent_dna := ""
	updated := int(time.Now().Unix())
	account_id := "32fc4139f7d0347ca9ea70d30caad45a5d90fc23aaefacedf6bff2746e2073f3"
	sub_account_id := ""
	application_status := "approved"
	group_id := "53687b42c97eb2dc38eac2483b7623d17ad4337dbec1cc54369cc8f26b52a71d"
	whitelist_id := "4aedc51597bc24912bafdf3f9b7f104ccc1392d32001b184ffba1940ff319a8f"

	signature, preObj, err := PostApproveOrDeclineGroupMemberWhitelist_SignatureStr(parent_dna, updated,
		account_id, sub_account_id, application_status)
	if err != nil {
		t.Errorf("TestPostApproveOrDeclineGroupMemberWhitelist error:%v", err.Error())
		return
	}
	if preObj == nil {
		t.Errorf("TestPostApproveOrDeclineGroupMemberWhitelist preObj object is nil")
		return
	}
	if signature == "" {
		t.Errorf("TestPostApproveOrDeclineGroupMemberWhitelist signature value is empty")
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

	putAddGroupWhite, err := PostApproveOrDeclineGroupMemberWhitelist(group_id, whitelist_id, signValue, preObj)
	if err != nil {
		t.Errorf("TestPostApproveOrDeclineGroupMemberWhitelist error:%v", err.Error())
		return
	}

	if putAddGroupWhite != nil {
		t.Logf("TestPostApproveOrDeclineGroupMemberWhitelist response value:%v", putAddGroupWhite)
		if putAddGroupWhite.ResultCode != core.CONST_ResultCode_Success {
			t.Errorf("TestPostApproveOrDeclineGroupMemberWhitelist response error:%v", putAddGroupWhite.ResultMsg)
			return
		}
		if putAddGroupWhite.Data != nil {
			t.Logf("TestPostApproveOrDeclineGroupMemberWhitelist response value:%v", putAddGroupWhite.Data)
		} else {
			t.Logf("TestPostApproveOrDeclineGroupMemberWhitelist response value don't find ")
		}
	}
}
