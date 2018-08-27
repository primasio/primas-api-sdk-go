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

package token

import (
	"testing"

	"github.com/primasio/primas-api-sdk-go/core"
)

func TestGetTokenPreLockList(t *testing.T) {
	account_id := "809a85f7ddf8ae5aaa49fe30be10e07e09156dc04166fab98bbd7bb42b2dc26c"
	resultIncentivesListResponse, err := GetTokenPreLockList(account_id, 1527695509, 1597695509,
		0, 20, CONST_Token_PreLock_Type_Lock)
	if err != nil {
		t.Errorf("GetTokenPreLockList error:%v", err.Error())
		return
	}

	if resultIncentivesListResponse != nil {
		t.Logf("GetTokenPreLockList response value:%v", resultIncentivesListResponse)
		if resultIncentivesListResponse.ResultCode != core.CONST_ResultCode_Success {
			t.Errorf("GetTokenPreLockList response error:%v", resultIncentivesListResponse.ResultMsg)
			return
		}

		if resultIncentivesListResponse.Data != nil {
			t.Logf("GetTokenPreLockList response AccountTokensData value:%v", resultIncentivesListResponse.Data)
		} else {
			t.Logf("GetTokenPreLockList response AccountTokensData value don't find ")
		}
	}
}
