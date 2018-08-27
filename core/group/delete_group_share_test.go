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

func TestDeleteGroupShare(t *testing.T) {
	parent_dna := "04a4694c3825e3fee7508e3ce44b3dd16110ac82b3089517ac7cd5dc59e66103"
	updated := int(time.Now().Unix())
	account_id := "32fc4139f7d0347ca9ea70d30caad45a5d90fc23aaefacedf6bff2746e2073f3"
	sub_account_id := ""
	share_id := `6519eeec38cd8fc8ed915b8d2ababfa5c0179c16a642bedbb9e4eb236b76ca5a`

	signature, preObj, err := DeleteGroupShare_SignatureStr(parent_dna, updated, account_id, sub_account_id)
	if err != nil {
		t.Errorf("TestDeleteGroupShare error:%v", err.Error())
		return
	}
	if preObj == nil {
		t.Errorf("TestDeleteGroupShare preObj object is nil")
		return
	}
	if signature == "" {
		t.Errorf("TestDeleteGroupShare signature value is empty")
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

	deleteGroupShare, err := DeleteGroupShare(share_id, signValue, preObj)
	if err != nil {
		t.Errorf("TestDeleteGroupShare error:%v", err.Error())
		return
	}

	if deleteGroupShare != nil {
		t.Logf("TestDeleteGroupShare response value:%v", deleteGroupShare)
		if deleteGroupShare.ResultCode != core.CONST_ResultCode_Success {
			t.Errorf("TestDeleteGroupShare response error:%v", deleteGroupShare.ResultMsg)
			return
		}
		if deleteGroupShare.Data != nil {
			t.Logf("TestDeleteGroupShare response value:%v", deleteGroupShare.Data)
		} else {
			t.Logf("TestDeleteGroupShare response value don't find ")
		}
	}
}
