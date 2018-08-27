/*
 * Copyright 2018 Primas Lab Foundation
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *       http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package content

import (
	"encoding/hex"
	"io/ioutil"
	"log"
	"testing"
	"time"

	"github.com/primasio/primas-api-sdk-go/core"
	"github.com/primasio/primas-api-sdk-go/core/dtcp/dtcpv1"
	"github.com/primasio/primas-api-sdk-go/core/tool"
)

func TestPostContent_Article(t *testing.T) {
	title := "test"
	account_id := "32fc4139f7d0347ca9ea70d30caad45a5d90fc23aaefacedf6bff2746e2073f3"
	sub_account_id := ""
	sub_account_name := ""
	abstract := "test abstract"
	language := "zh"
	category := "区块链"
	created := int(time.Now().Unix())
	content := "showtime"

	signature, preObj, err := PostContent_SignatureStr(dtcpv1.CONST_DTCP_Tag_Article,
		title, account_id, sub_account_id, sub_account_name, abstract, language, category,
		created, content, nil)
	if err != nil {
		t.Errorf("TestPostContent_Article error:%v", err.Error())
		return
	}
	if preObj == nil {
		t.Errorf("TestPostContent_Article preObj object is nil")
		return
	}
	if signature == "" {
		t.Errorf("TestPostContent_Article signature value is empty")
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

	postContent, err := PostContent_Aritcle(signValue, preObj)
	if err != nil {
		t.Errorf("TestPostContent_Article error:%v", err.Error())
		return
	}

	if postContent != nil {
		t.Logf("PostContent response value:%v", postContent)
		if postContent.ResultCode != core.CONST_ResultCode_Success {
			t.Errorf("TestPostContent_Article response error:%v", postContent.ResultMsg)
			return
		}
		if postContent.Data != nil {
			t.Logf("TestPostContent_Article response value:%v", postContent.Data)
		} else {
			t.Logf("TestPostContent_Article response value don't find ")
		}
	}
}

func TestPostContent_ImageUrlencoded(t *testing.T) {
	imageFilePath := "/Users/kevinchen/Downloads/jianpan.jpg"
	contentBody, err := ioutil.ReadFile(imageFilePath)
	if err != nil {
		t.Errorf("TestPostContent_Image error:%v", err.Error())
		return
	}

	title := "test"
	account_id := "32fc4139f7d0347ca9ea70d30caad45a5d90fc23aaefacedf6bff2746e2073f3"
	sub_account_id := ""
	sub_account_name := ""
	abstract := "test abstract"
	language := "zh"
	category := "区块链"
	created := int(time.Now().Unix())
	content := string(contentBody)

	signature, preObj, err := PostContent_SignatureStr(dtcpv1.CONST_DTCP_Tag_Image,
		title, account_id, sub_account_id, sub_account_name, abstract, language, category,
		created, content, nil)
	if err != nil {
		t.Errorf("TestPostContent_Image error:%v", err.Error())
		return
	}
	if preObj == nil {
		t.Errorf("TestPostContent_Image preObj object is nil")
		return
	}
	if signature == "" {
		t.Errorf("TestPostContent_Image signature value is empty")
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

	postContent, err := PostContent_ImageUrlencoded(signValue, preObj)
	if err != nil {
		t.Errorf("TestPostContent_Image error:%v", err.Error())
		return
	}

	if postContent != nil {
		t.Logf("TestPostContent_Image response value:%v", postContent)
		if postContent.ResultCode != core.CONST_ResultCode_Success {
			t.Errorf("PostContent response error:%v", postContent.ResultMsg)
			return
		}
		if postContent.Data != nil {
			t.Logf("TestPostContent_Image response value:%v", postContent.Data)
		} else {
			t.Logf("TestPostContent_Image response value don't find ")
		}
	}
}

func TestPostContent_Image(t *testing.T) {
	imageFilePath := "/Users/kevinchen/Desktop/test.png"
	contentBody, err := ioutil.ReadFile(imageFilePath)
	if err != nil {
		t.Errorf("TestPostContent_Image error:%v", err.Error())
		return
	}

	log.Printf("img hex:%v", hex.EncodeToString(contentBody))

	title := "content test"
	account_id := "32fc4139f7d0347ca9ea70d30caad45a5d90fc23aaefacedf6bff2746e2073f3"
	sub_account_id := ""
	sub_account_name := ""
	abstract := "abstract"
	language := "en-US"
	category := "test"
	created := 1534161371 //int(time.Now().Unix())
	content := string(contentBody)

	signature, preObj, err := PostContent_SignatureStr(dtcpv1.CONST_DTCP_Tag_Image,
		title, account_id, sub_account_id, sub_account_name, abstract, language, category,
		created, content, nil)
	if err != nil {
		t.Errorf("TestPostContent_Image error:%v", err.Error())
		return
	}
	if preObj == nil {
		t.Errorf("TestPostContent_Image preObj object is nil")
		return
	}
	if signature == "" {
		t.Errorf("TestPostContent_Image signature value is empty")
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

	log.Println("signValue:%v", signValue)
}

func TestPostContent_ImageMultipartForm(t *testing.T) {
	imageFilePath := "/Users/kevinchen/Downloads/jianpan.jpg"
	contentBody, err := ioutil.ReadFile(imageFilePath)
	if err != nil {
		t.Errorf("TestPostContent_Image error:%v", err.Error())
		return
	}

	title := "test"
	account_id := "32fc4139f7d0347ca9ea70d30caad45a5d90fc23aaefacedf6bff2746e2073f3"
	sub_account_id := ""
	sub_account_name := ""
	abstract := "test abstract"
	language := "zh"
	category := "区块链"
	created := int(time.Now().Unix())
	content := string(contentBody)

	signature, preObj, err := PostContent_SignatureStr(dtcpv1.CONST_DTCP_Tag_Image,
		title, account_id, sub_account_id, sub_account_name, abstract, language, category,
		created, content, nil)
	if err != nil {
		t.Errorf("TestPostContent_Image error:%v", err.Error())
		return
	}
	if preObj == nil {
		t.Errorf("TestPostContent_Image preObj object is nil")
		return
	}
	if signature == "" {
		t.Errorf("TestPostContent_Image signature value is empty")
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

	postContent, err := PostContent_ImageMultipartForm(signValue, preObj, imageFilePath)
	if err != nil {
		t.Errorf("TestPostContent_Image error:%v", err.Error())
		return
	}

	if postContent != nil {
		t.Logf("TestPostContent_Image response value:%v", postContent)
		if postContent.ResultCode != core.CONST_ResultCode_Success {
			t.Errorf("PostContent response error:%v", postContent.ResultMsg)
			return
		}
		if postContent.Data != nil {
			t.Logf("TestPostContent_Image response value:%v", postContent.Data)
		} else {
			t.Logf("TestPostContent_Image response value don't find ")
		}
	}
}
