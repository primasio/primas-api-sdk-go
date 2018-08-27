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

func TestPutMemberApplication(t *testing.T) {
	group_id := `53687b42c97eb2dc38eac2483b7623d17ad4337dbec1cc54369cc8f26b52a71d`
	group_member_id := `378fa5bf5779eef71445f8b289db30aafa8c7bfdcfa843e2c05d838a9b044a68`
	parent_dna := `1d334e6f275ed3d7f862c6f9b3d1a2d495de4db9741109d335414efda0782f97`
	updated := int(time.Now().Unix())
	account_id := `32fc4139f7d0347ca9ea70d30caad45a5d90fc23aaefacedf6bff2746e2073f3`
	sub_account_id := ``
	application_status := `approved`
	signature, preObj, err := PutMemberApplication_SignatureStr(parent_dna, updated, account_id, sub_account_id, application_status)
	if err != nil {
		t.Errorf("TestPutMemberApplication error:%v", err.Error())
		return
	}
	if preObj == nil {
		t.Errorf("TestPutMemberApplication preObj object is nil")
		return
	}
	if signature == "" {
		t.Errorf("TestPutMemberApplication signature value is empty")
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

	putMemberApp, err := PutMemberApplication(group_id, group_member_id, signValue, preObj)
	if err != nil {
		t.Errorf("TestPutMemberApplication error:%v", err.Error())
		return
	}

	if putMemberApp != nil {
		t.Logf("TestPutMemberApplication response value:%v", putMemberApp)
		if putMemberApp.ResultCode != core.CONST_ResultCode_Success {
			t.Errorf("TestPutMemberApplication response error:%v", putMemberApp.ResultMsg)
			return
		}
		if putMemberApp.Data != nil {
			t.Logf("TestPutMemberApplication response value:%v", putMemberApp.Data)
		} else {
			t.Logf("TestPutMemberApplication response value don't find ")
		}
	}
}
