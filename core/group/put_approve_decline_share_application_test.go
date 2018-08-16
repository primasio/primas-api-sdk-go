package group

import (
	"log"
	"testing"
	"time"

	"github.com/primasio/primas-api-sdk-go/core"
	"github.com/primasio/primas-api-sdk-go/core/tool"
)

func TestPutApproveOrDeclineShareApplication(t *testing.T) {
	share_id := `9b319ea45770a6f4896a75938448a1654ceeba6aa38cdd35f6788ffa9b3e6143`
	parent_dna := `1d334e6f275ed3d7f862c6f9b3d1a2d495de4db9741109d335414efda0782f97`
	updated := int(time.Now().Unix())
	account_id := `32fc4139f7d0347ca9ea70d30caad45a5d90fc23aaefacedf6bff2746e2073f3`
	sub_account_id := ``
	application_status := `approved`
	signature, preObj, err := PutApproveOrDeclineShareApplication_SignatureStr(parent_dna, updated, account_id, sub_account_id, application_status)
	if err != nil {
		t.Errorf("TestPutApproveOrDeclineShareApplication error:%v", err.Error())
		return
	}
	if preObj == nil {
		t.Errorf("TestPutApproveOrDeclineShareApplication preObj object is nil")
		return
	}
	if signature == "" {
		t.Errorf("TestPutApproveOrDeclineShareApplication signature value is empty")
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

	putShareGroupApp, err := PutApproveOrDeclineShareApplication(share_id, signValue, preObj)
	if err != nil {
		t.Errorf("TestPutApproveOrDeclineShareApplication error:%v", err.Error())
		return
	}

	if putShareGroupApp != nil {
		t.Logf("TestPutApproveOrDeclineShareApplication response value:%v", putShareGroupApp)
		if putShareGroupApp.ResultCode != core.CONST_ResultCode_Success {
			t.Errorf("TestPutApproveOrDeclineShareApplication response error:%v", putShareGroupApp.ResultMsg)
			return
		}
		if putShareGroupApp.Data != nil {
			t.Logf("TestPutApproveOrDeclineShareApplication response value:%v", putShareGroupApp.Data)
		} else {
			t.Logf("TestPutApproveOrDeclineShareApplication response value don't find ")
		}
	}
}
