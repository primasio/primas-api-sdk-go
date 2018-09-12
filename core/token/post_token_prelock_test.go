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
	"github.com/shopspring/decimal"
)

func TestPostPreLockTokens(t *testing.T) {
	address := "0xd75407ad8cabeeebfed78c4f3794208b3339fbf4"
	account_id := "32fc4139f7d0347ca9ea70d30caad45a5d90fc23aaefacedf6bff2746e2073f3"
	amount := decimal.New(120, 18)

	resultPreLockTokenResponse, err := PostPreLockTokens(address, account_id, amount)
	if err != nil {
		t.Errorf("PostPreLockTokens error:%v", err.Error())
		return
	}

	if resultPreLockTokenResponse != nil {
		t.Logf("PostPreLockTokens response value:%v", resultPreLockTokenResponse)
		if resultPreLockTokenResponse.ResultCode != core.CONST_ResultCode_Success {
			t.Errorf("PostPreLockTokens response error:%v", resultPreLockTokenResponse.ResultMsg)
			return
		}

		if resultPreLockTokenResponse.Data != nil {
			t.Logf("PostPreLockTokens response data value:%v", resultPreLockTokenResponse.Data)
		} else {
			t.Logf("PostPreLockTokens response data value don't find ")
		}
	}
}
