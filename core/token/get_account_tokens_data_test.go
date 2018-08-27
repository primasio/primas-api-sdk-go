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
	"fmt"
	"testing"
	"time"

	"github.com/primasio/primas-api-sdk-go/core"
)

func TestGetAccountTokensData(t *testing.T) {
	fmt.Println(time.Now().UnixNano() / 1e6)

	account_id := "809a85f7ddf8ae5aaa49fe30be10e07e09156dc04166fab98bbd7bb42b2dc26c"
	resultAccountTokensResponse, err := GetAccountTokensData(account_id)
	if err != nil {
		t.Errorf("GetAccountTokensData error:%v", err.Error())
		return
	}

	if resultAccountTokensResponse != nil {
		t.Logf("GetAccountTokensData response value:%v", resultAccountTokensResponse)
		if resultAccountTokensResponse.ResultCode != core.CONST_ResultCode_Success {
			t.Errorf("GetAccountTokensData response error:%v", resultAccountTokensResponse.ResultMsg)
			return
		}
		if resultAccountTokensResponse.Data != nil {
			t.Logf("GetAccountTokensData response AccountTokensData value:%v", resultAccountTokensResponse.Data)
		} else {
			t.Logf("GetAccountTokensData response AccountTokensData value don't find ")
		}
	}

}
