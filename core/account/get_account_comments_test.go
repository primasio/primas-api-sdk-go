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

func TestGetAccountComments(t *testing.T) {
	account_id := "809a85f7ddf8ae5aaa49fe30be10e07e09156dc04166fab98bbd7bb42b2dc26c"
	resultAccountCommentsResponse, err := GetAccountComments(account_id, 0, 20)
	if err != nil {
		t.Errorf("GetAccountComments error:%v", err.Error())
		return
	}

	if resultAccountCommentsResponse != nil {
		t.Logf("GetAccountComments response value:%v", resultAccountCommentsResponse)
		if resultAccountCommentsResponse.ResultCode != core.CONST_ResultCode_Success {
			t.Errorf("GetAccountComments response error:%v", resultAccountCommentsResponse.ResultMsg)
			return
		}

		if resultAccountCommentsResponse.Data != nil {
			t.Logf("GetAccountComments response AccountTokensData value:%#v", resultAccountCommentsResponse.Data)
		} else {
			t.Logf("GetAccountComments response AccountTokensData value don't find ")
		}
	}
}

func TestGetSubAccountComments(t *testing.T) {
	account_id := "809a85f7ddf8ae5aaa49fe30be10e07e09156dc04166fab98bbd7bb42b2dc26c"
	sub_account_id := "a_0001"
	resultAccountCommentsResponse, err := GetSubAccountComments(account_id, sub_account_id, 0, 20)
	if err != nil {
		t.Errorf("GetSubAccountComments error:%v", err.Error())
		return
	}

	if resultAccountCommentsResponse != nil {
		t.Logf("GetSubAccountComments response value:%v", resultAccountCommentsResponse)
		if resultAccountCommentsResponse.ResultCode != core.CONST_ResultCode_Success {
			t.Errorf("GetSubAccountComments response error:%v", resultAccountCommentsResponse.ResultMsg)
			return
		}

		if resultAccountCommentsResponse.Data != nil {
			t.Logf("GetSubAccountComments response AccountTokensData value:%#v", resultAccountCommentsResponse.Data)
		} else {
			t.Logf("GetSubAccountComments response AccountTokensData value don't find ")
		}
	}
}
