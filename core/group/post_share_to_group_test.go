package group

import (
	"log"
	"testing"
	"time"

	"github.com/primasio/primas-api-sdk-go/core"
	"github.com/primasio/primas-api-sdk-go/core/tool"
)

func TestPostShareToGroup(t *testing.T) {
	src_id := `9d3babdfb0a4053ebaefc4570040b90207e6a903026d20c228c46d1e52e64775`
	dest_id := `53687b42c97eb2dc38eac2483b7623d17ad4337dbec1cc54369cc8f26b52a71d`
	account_id := `32fc4139f7d0347ca9ea70d30caad45a5d90fc23aaefacedf6bff2746e2073f3`
	sub_account_id := ``
	sub_account_name := ``
	created := int(time.Now().Unix())
	share_id := "" // "631807f6f03a6799164683081836f4073c94af89148f8d4b42d1582dda5ae843"
	application_status := "pending"
	application_expire := int(time.Unix(9*60*60, 0).Unix())
	signature, preObj, err := PostShareToGroup_SignatureStr(src_id, dest_id, account_id, sub_account_id, sub_account_name, created,
		share_id, application_status, application_expire)
	if err != nil {
		t.Errorf("TestPostShareToGroup error:%v", err.Error())
		return
	}
	if preObj == nil {
		t.Errorf("TestPostShareToGroup preObj object is nil")
		return
	}
	if signature == "" {
		t.Errorf("TestPostShareToGroup signature value is empty")
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

	postJoinGroup, err := PostShareToGroup(signValue, preObj)
	if err != nil {
		t.Errorf("TestPostShareToGroup error:%v", err.Error())
		return
	}

	if postJoinGroup != nil {
		t.Logf("TestPostShareToGroup response value:%v", postJoinGroup)
		if postJoinGroup.ResultCode != core.CONST_ResultCode_Success {
			t.Errorf("TestPostShareToGroup response error:%v", postJoinGroup.ResultMsg)
			return
		}
		if postJoinGroup.Data != nil {
			t.Logf("TestPostShareToGroup response value:%v", postJoinGroup.Data)
		} else {
			t.Logf("TestPostShareToGroup response value don't find ")
		}
	}
}
