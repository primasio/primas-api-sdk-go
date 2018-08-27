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

package group

import (
	"log"
	"testing"
	"time"

	"github.com/primasio/primas-api-sdk-go/core"
	"github.com/primasio/primas-api-sdk-go/core/tool"
)

func TestPostJoinGroup(t *testing.T) {
	src_id := `809a85f7ddf8ae5aaa49fe30be10e07e09156dc04166fab98bbd7bb42b2dc26c`
	dest_id := `53687b42c97eb2dc38eac2483b7623d17ad4337dbec1cc54369cc8f26b52a71d`
	account_id := `32fc4139f7d0347ca9ea70d30caad45a5d90fc23aaefacedf6bff2746e2073f3`
	sub_account_id := ``
	sub_account_name := ``
	created := int(time.Now().Unix())
	application_status := "pending"
	application_expire := int(time.Unix(9*60*60, 0).Unix())
	signature, preObj, err := PostJoinGroup_SignatureStr(src_id, dest_id, account_id, sub_account_id, sub_account_name, created,
		application_status, application_expire)
	if err != nil {
		t.Errorf("TestPostJoinGroup error:%v", err.Error())
		return
	}
	if preObj == nil {
		t.Errorf("TestPostJoinGroup preObj object is nil")
		return
	}
	if signature == "" {
		t.Errorf("TestPostJoinGroup signature value is empty")
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

	postJoinGroup, err := PostJoinGroup(signValue, preObj)
	if err != nil {
		t.Errorf("TestPostJoinGroup error:%v", err.Error())
		return
	}

	if postJoinGroup != nil {
		t.Logf("TestPostJoinGroup response value:%v", postJoinGroup)
		if postJoinGroup.ResultCode != core.CONST_ResultCode_Success {
			t.Errorf("TestPostJoinGroup response error:%v", postJoinGroup.ResultMsg)
			return
		}
		if postJoinGroup.Data != nil {
			t.Logf("TestPostJoinGroup response value:%v", postJoinGroup.Data)
		} else {
			t.Logf("TestPostJoinGroup response value don't find ")
		}
	}
}
