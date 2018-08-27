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

func TestDeleteQuitGroupMemberWhitelist(t *testing.T) {
	parent_dna := ""
	updated := int(time.Now().Unix())
	account_id := "32fc4139f7d0347ca9ea70d30caad45a5d90fc23aaefacedf6bff2746e2073f3"
	sub_account_id := ""
	group_id := "53687b42c97eb2dc38eac2483b7623d17ad4337dbec1cc54369cc8f26b52a71d"
	whitelist_id := "4aedc51597bc24912bafdf3f9b7f104ccc1392d32001b184ffba1940ff319a8f"

	signature, preObj, err := DeleteQuitGroupMemberWhitelist_SignatureStr(parent_dna, updated,
		account_id, sub_account_id)
	if err != nil {
		t.Errorf("TestDeleteQuitGroupMemberWhitelist error:%v", err.Error())
		return
	}
	if preObj == nil {
		t.Errorf("TestDeleteQuitGroupMemberWhitelist preObj object is nil")
		return
	}
	if signature == "" {
		t.Errorf("TestDeleteQuitGroupMemberWhitelist signature value is empty")
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

	delAddGroupWhite, err := DeleteQuitGroupMemberWhitelist(group_id, whitelist_id, signValue, preObj)
	if err != nil {
		t.Errorf("TestDeleteQuitGroupMemberWhitelist error:%v", err.Error())
		return
	}

	if delAddGroupWhite != nil {
		t.Logf("TestPostApproveOrDeclineGroupMemberWhitelist response value:%v", delAddGroupWhite)
		if delAddGroupWhite.ResultCode != core.CONST_ResultCode_Success {
			t.Errorf("TestDeleteQuitGroupMemberWhitelist response error:%v", delAddGroupWhite.ResultMsg)
			return
		}
		if delAddGroupWhite.Data != nil {
			t.Logf("TestDeleteQuitGroupMemberWhitelist response value:%v", delAddGroupWhite.Data)
		} else {
			t.Logf("TestDeleteQuitGroupMemberWhitelist response value don't find ")
		}
	}
}
