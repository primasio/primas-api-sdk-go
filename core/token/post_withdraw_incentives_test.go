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
	_ "github.com/primasio/primas-api-sdk-go/core/tool"
	"github.com/shopspring/decimal"
)

// todo not complete
func TestPostWithdrawIncentives(t *testing.T) {
	account_id := "809a85f7ddf8ae5aaa49fe30be10e07e09156dc04166fab98bbd7bb42b2dc26c"
	node_id := "58f47077984e5daa4d2ea46f2e689177a1655c1321544e69f851530a789e9fd7"
	created := 1532525161
	amount := decimal.New(123, 0)
	node_fee := decimal.New(123, 0)

	resultWithdrawIncetice, err := PostWithdrawIncentives(account_id, node_id, created, amount, node_fee)
	if err != nil {
		t.Errorf("PostWithdrawIncentives error:%v", err.Error())
		return
	}

	if resultWithdrawIncetice != nil {
		t.Logf("PostWithdrawIncentives response value:%v", resultWithdrawIncetice)
		if resultWithdrawIncetice.ResultCode != core.CONST_ResultCode_Success {
			t.Errorf("PostWithdrawIncentives response error:%v", resultWithdrawIncetice.ResultMsg)
			return
		}

		if resultWithdrawIncetice.Data != nil {
			t.Logf("PostWithdrawIncentives response AccountTokensData value:%v", resultWithdrawIncetice.Data)
		} else {
			t.Logf("PostWithdrawIncentives response AccountTokensData value don't find ")
		}
	}
}
