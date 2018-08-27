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

func TestGetAccountLikes(t *testing.T) {
	account_id := "809a85f7ddf8ae5aaa49fe30be10e07e09156dc04166fab98bbd7bb42b2dc26c"
	resultAccountLikesResponse, err := GetAccountLikes(account_id, 0, 20)
	if err != nil {
		t.Errorf("GetAccountLikes error:%v", err.Error())
		return
	}

	if resultAccountLikesResponse != nil {
		t.Logf("GetAccountLikes response value:%v", resultAccountLikesResponse)
		if resultAccountLikesResponse.ResultCode != core.CONST_ResultCode_Success {
			t.Errorf("GetAccountLikes response error:%v", resultAccountLikesResponse.ResultMsg)
			return
		}

		if resultAccountLikesResponse.Data != nil {
			t.Logf("GetAccountLikes response AccountTokensData value:%#v", resultAccountLikesResponse.Data)
		} else {
			t.Logf("GetAccountLikes response AccountTokensData value don't find ")
		}
	}
}

func TestGetSubAccountLikes(t *testing.T) {
	account_id := "809a85f7ddf8ae5aaa49fe30be10e07e09156dc04166fab98bbd7bb42b2dc26c"
	sub_account_id := "a_0001"
	resultAccountLikesResponse, err := GetSubAccountLikes(account_id, sub_account_id, 0, 20)
	if err != nil {
		t.Errorf("GetSubAccountLikes error:%v", err.Error())
		return
	}

	if resultAccountLikesResponse != nil {
		t.Logf("GetSubAccountLikes response value:%v", resultAccountLikesResponse)
		if resultAccountLikesResponse.ResultCode != core.CONST_ResultCode_Success {
			t.Errorf("GetSubAccountLikes response error:%v", resultAccountLikesResponse.ResultMsg)
			return
		}

		if resultAccountLikesResponse.Data != nil {
			t.Logf("GetSubAccountLikes response AccountTokensData value:%#v", resultAccountLikesResponse.Data)
		} else {
			t.Logf("GetSubAccountLikes response AccountTokensData value don't find ")
		}
	}
}
