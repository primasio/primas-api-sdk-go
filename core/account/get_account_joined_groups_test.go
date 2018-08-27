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

package account

import (
	"testing"

	"github.com/primasio/primas-api-sdk-go/core"
)

func TestGetAccountJoinedGroups(t *testing.T) {
	account_id := "809a85f7ddf8ae5aaa49fe30be10e07e09156dc04166fab98bbd7bb42b2dc26c"
	resultGetAccountJoinedGroups, err := GetAccountJoinedGroups(account_id, 0, 20)
	if err != nil {
		t.Errorf("GetAccountJoinedGroups error:%v", err.Error())
		return
	}

	if resultGetAccountJoinedGroups != nil {
		t.Logf("GetAccountJoinedGroups response value:%v", resultGetAccountJoinedGroups)
		if resultGetAccountJoinedGroups.ResultCode != core.CONST_ResultCode_Success {
			t.Errorf("GetAccountJoinedGroups response error:%v", resultGetAccountJoinedGroups.ResultMsg)
			return
		}

		if resultGetAccountJoinedGroups.Data != nil {
			t.Logf("GetAccountJoinedGroups response AccountTokensData value:%#v", resultGetAccountJoinedGroups.Data)
		} else {
			t.Logf("GetAccountJoinedGroups response AccountTokensData value don't find ")
		}
	}
}

func TestGetSubAccountJoinedGroups(t *testing.T) {
	account_id := "809a85f7ddf8ae5aaa49fe30be10e07e09156dc04166fab98bbd7bb42b2dc26c"
	sub_account_id := "a_0001"
	resultGetSubAccountJoinedGroups, err := GetSubAccountJoinedGroups(account_id, sub_account_id, 0, 20)
	if err != nil {
		t.Errorf("GetSubAccountJoinedGroups error:%v", err.Error())
		return
	}

	if resultGetSubAccountJoinedGroups != nil {
		t.Logf("GetSubAccountJoinedGroups response value:%v", resultGetSubAccountJoinedGroups)
		if resultGetSubAccountJoinedGroups.ResultCode != core.CONST_ResultCode_Success {
			t.Errorf("GetSubAccountJoinedGroups response error:%v", resultGetSubAccountJoinedGroups.ResultMsg)
			return
		}

		if resultGetSubAccountJoinedGroups.Data != nil {
			t.Logf("GetSubAccountJoinedGroups response AccountTokensData value:%#v", resultGetSubAccountJoinedGroups.Data)
		} else {
			t.Logf("GetSubAccountJoinedGroups response AccountTokensData value don't find ")
		}
	}
}
