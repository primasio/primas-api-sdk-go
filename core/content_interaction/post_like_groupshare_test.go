package content_interaction

import (
	"testing"
	"time"

	"github.com/primasio/primas-api-sdk-go/core"
	"github.com/primasio/primas-api-sdk-go/core/tool"
)

func TestPostLikeOfGroupshare(t *testing.T) {
	account_id := "32fc4139f7d0347ca9ea70d30caad45a5d90fc23aaefacedf6bff2746e2073f3"
	share_id := "60e4388167b2cb963964409b87510c57f8ba5a39487a900efc29b677d62cf449"
	sub_account_id := ""
	sub_account_name := ""
	created := time.Now().Unix()
	signature, preObj, err := PostLikeOfGroupshare_SignatureStr(account_id, share_id, sub_account_id, sub_account_name, int(created))
	if err != nil {
		t.Errorf("TestPostLikeOfGroupshare error:%v", err.Error())
		return
	}
	if preObj == nil {
		t.Errorf("TestPostLikeOfGroupshare preObj object is nil")
		return
	}
	if signature == "" {
		t.Errorf("TestPostLikeOfGroupshare signature value is empty")
		return
	}

	// mock Sign
	privateKey := tool.GetClientPrivateKey()
	signValue, err := tool.Sign([]byte(signature), privateKey)
	if err != nil {
		t.Errorf("Sign error %v:", err.Error())
		return
	}
	//

	postContent, err := PostLikeOfGroupshare(signValue, preObj)
	if err != nil {
		t.Errorf("TestPostLikeOfGroupshare error:%v", err.Error())
		return
	}

	if postContent != nil {
		t.Logf("TestPostContent_Image response value:%v", postContent)
		if postContent.ResultCode != core.CONST_ResultCode_Success {
			t.Errorf("TestPostLikeOfGroupshare response error:%v", postContent.ResultMsg)
			return
		}
		if postContent.Data != nil {
			t.Logf("TestPostLikeOfGroupshare response value:%v", postContent.Data)
		} else {
			t.Logf("TestPostLikeOfGroupshare response value don't find ")
		}
	}
}
