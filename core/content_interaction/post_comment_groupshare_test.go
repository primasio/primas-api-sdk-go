package content_interaction

import (
	"testing"
	"time"

	"github.com/primasio/primas-api-sdk-go/core"
	"github.com/primasio/primas-api-sdk-go/core/tool"
)

func TestPostCommentOfGroupshare(t *testing.T) {
	account_id := "32fc4139f7d0347ca9ea70d30caad45a5d90fc23aaefacedf6bff2746e2073f3"
	share_id := "60e4388167b2cb963964409b87510c57f8ba5a39487a900efc29b677d62cf449"
	sub_account_id := ""
	sub_account_name := ""
	parent_comment_id := ""
	content := "nice"
	created := time.Now().Unix()
	signature, preObj, err := PostCommentOfGroupshare_SignatureStr(account_id, share_id, sub_account_id,
		sub_account_name, int(created), parent_comment_id, content)
	if err != nil {
		t.Errorf("TestPostCommentOfGroupshare error:%v", err.Error())
		return
	}
	if preObj == nil {
		t.Errorf("TestPostCommentOfGroupshare preObj object is nil")
		return
	}
	if signature == "" {
		t.Errorf("TestPostCommentOfGroupshare signature value is empty")
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

	postContent, err := PostCommentOfGroupshare(signValue, preObj)
	if err != nil {
		t.Errorf("TestPostCommentOfGroupshare error:%v", err.Error())
		return
	}

	if postContent != nil {
		t.Logf("TestPostCommentOfGroupshare response value:%v", postContent)
		if postContent.ResultCode != core.CONST_ResultCode_Success {
			t.Errorf("TestPostCommentOfGroupshare response error:%v", postContent.ResultMsg)
			return
		}
		if postContent.Data != nil {
			t.Logf("TestPostCommentOfGroupshare response value:%v", postContent.Data)
		} else {
			t.Logf("TestPostCommentOfGroupshare response value don't find ")
		}
	}
}
