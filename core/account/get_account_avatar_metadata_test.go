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

func TestGetAccountAvatarMetadata(t *testing.T) {
	account_id := "809a85f7ddf8ae5aaa49fe30be10e07e09156dc04166fab98bbd7bb42b2dc26c"
	resultAccountAvatarResponse, err := GetAccountAvatarMetadata(account_id)
	if err != nil {
		t.Errorf("GetAccountComments error:%v", err.Error())
		return
	}

	if resultAccountAvatarResponse != nil {
		t.Logf("GetAccountAvatarMetadata response value:%v", resultAccountAvatarResponse)
		if resultAccountAvatarResponse.ResultCode != core.CONST_ResultCode_Success {
			t.Errorf("GetAccountAvatarMetadata response error:%v", resultAccountAvatarResponse.ResultMsg)
			return
		}

		if resultAccountAvatarResponse.Data != nil {
			t.Logf("GetAccountAvatarMetadata response AccountTokensData value:%#v", resultAccountAvatarResponse.Data)
		} else {
			t.Logf("GetAccountAvatarMetadata response AccountTokensData value don't find ")
		}
	}
}

func TestGetSubAccountAvatarMetadata(t *testing.T) {
	account_id := "809a85f7ddf8ae5aaa49fe30be10e07e09156dc04166fab98bbd7bb42b2dc26c"
	sub_account_id := "a_0001"
	resultAccountAvatarResponse, err := GetSubAccountAvatarMetadata(account_id, sub_account_id)
	if err != nil {
		t.Errorf("GetSubAccountAvatarMetadata error:%v", err.Error())
		return
	}

	if resultAccountAvatarResponse != nil {
		t.Logf("GetSubAccountAvatarMetadata response value:%v", resultAccountAvatarResponse)
		if resultAccountAvatarResponse.ResultCode != core.CONST_ResultCode_Success {
			t.Errorf("GetSubAccountAvatarMetadata response error:%v", resultAccountAvatarResponse.ResultMsg)
			return
		}

		if resultAccountAvatarResponse.Data != nil {
			t.Logf("GetSubAccountAvatarMetadata response AccountTokensData value:%#v", resultAccountAvatarResponse.Data)
		} else {
			t.Logf("GetSubAccountAvatarMetadata response AccountTokensData value don't find ")
		}
	}
}
