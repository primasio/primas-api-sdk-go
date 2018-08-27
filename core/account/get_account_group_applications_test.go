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

func TestGetAccountGroupApplications(t *testing.T) {
	account_id := "809a85f7ddf8ae5aaa49fe30be10e07e09156dc04166fab98bbd7bb42b2dc26c"
	resultAccountGroupAppsResponse, err := GetAccountGroupApplications(account_id, 0, 20)
	if err != nil {
		t.Errorf("GetAccountGroupApplications error:%v", err.Error())
		return
	}

	if resultAccountGroupAppsResponse != nil {
		t.Logf("GetAccountGroupApplications response value:%v", resultAccountGroupAppsResponse)
		if resultAccountGroupAppsResponse.ResultCode != core.CONST_ResultCode_Success {
			t.Errorf("GetAccountGroupApplications response error:%v", resultAccountGroupAppsResponse.ResultMsg)
			return
		}

		if resultAccountGroupAppsResponse.Data != nil {
			t.Logf("GetAccountGroupApplications response AccountTokensData value:%#v", resultAccountGroupAppsResponse.Data)
		} else {
			t.Logf("GetAccountGroupApplications response AccountTokensData value don't find ")
		}
	}
}

func TestGetSubAccountGroupApplications(t *testing.T) {
	account_id := "809a85f7ddf8ae5aaa49fe30be10e07e09156dc04166fab98bbd7bb42b2dc26c"
	sub_account_id := "a_0001"
	resultAccountGroupAppsResponse, err := GetSubAccountGroupApplications(account_id, sub_account_id, 0, 20)
	if err != nil {
		t.Errorf("GetSubAccountGroupApplications error:%v", err.Error())
		return
	}

	if resultAccountGroupAppsResponse != nil {
		t.Logf("GetSubAccountGroupApplications response value:%v", resultAccountGroupAppsResponse)
		if resultAccountGroupAppsResponse.ResultCode != core.CONST_ResultCode_Success {
			t.Errorf("GetSubAccountGroupApplications response error:%v", resultAccountGroupAppsResponse.ResultMsg)
			return
		}

		if resultAccountGroupAppsResponse.Data != nil {
			t.Logf("GetSubAccountGroupApplications response AccountTokensData value:%#v", resultAccountGroupAppsResponse.Data)
		} else {
			t.Logf("GetSubAccountGroupApplications response AccountTokensData value don't find ")
		}
	}
}
