package group

import (
	"log"
	"testing"
	"time"

	"github.com/primasio/primas-api-sdk-go/core"
	"github.com/primasio/primas-api-sdk-go/core/tool"
)

func TestPostGroup(t *testing.T) {
	title := `PST开发社区`
	account_id := `32fc4139f7d0347ca9ea70d30caad45a5d90fc23aaefacedf6bff2746e2073f3`
	sub_account_id := ``
	sub_account_name := ``
	avatar := `06354a3ee794cf210e0db0d6729710ee047227d679668de697431fdbd1232ffc`
	abstract := `PST开发开源社区`
	language := `zh`
	category := `PST`
	created := int(time.Now().Unix())
	allow_join := `all`
	allow_post := `all`
	signature, preObj, err := PostGroup_SignatureStr(title, account_id, sub_account_id, sub_account_name, avatar, abstract, language, category,
		created, allow_join, allow_post)
	if err != nil {
		t.Errorf("TestPostGroup error:%v", err.Error())
		return
	}
	if preObj == nil {
		t.Errorf("TestPostGroup preObj object is nil")
		return
	}
	if signature == "" {
		t.Errorf("TestPostGroup signature value is empty")
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

	postGroup, err := PostGroup(signValue, preObj)
	if err != nil {
		t.Errorf("TestPostGroup error:%v", err.Error())
		return
	}

	if postGroup != nil {
		t.Logf("TestPostGroup response value:%v", postGroup)
		if postGroup.ResultCode != core.CONST_ResultCode_Success {
			t.Errorf("TestPostGroup response error:%v", postGroup.ResultMsg)
			return
		}
		if postGroup.Data != nil {
			t.Logf("TestPostGroup response value:%v", postGroup.Data)
		} else {
			t.Logf("TestPostGroup response value don't find ")
		}
	}
}
