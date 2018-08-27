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
	"testing"

	"github.com/primasio/primas-api-sdk-go/core"
)

func TestGetGroupMemberWhitelist(t *testing.T) {
	group_id := "53687b42c97eb2dc38eac2483b7623d17ad4337dbec1cc54369cc8f26b52a71d"
	resultGetGroupMemberWhitelist, err := GetGroupMemberWhitelist(group_id, 0, 20, "")
	if err != nil {
		t.Errorf("GetGroupMemberWhitelist error:%v", err.Error())
		return
	}

	if resultGetGroupMemberWhitelist != nil {
		t.Logf("GetGroupMemberWhitelist response value:%v", resultGetGroupMemberWhitelist)
		if resultGetGroupMemberWhitelist.ResultCode != core.CONST_ResultCode_Success {
			t.Errorf("GetGroupMemberWhitelist response error:%v", resultGetGroupMemberWhitelist.ResultMsg)
			return
		}

		if resultGetGroupMemberWhitelist.Data != nil {
			t.Logf("GetGroupMemberWhitelist response data value:%#v", resultGetGroupMemberWhitelist.Data)
		} else {
			t.Logf("GetGroupMemberWhitelist response data value don't find ")
		}
	}
}
