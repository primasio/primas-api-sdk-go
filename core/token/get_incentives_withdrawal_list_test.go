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

func TestGetIncentivesWithdrawalList(t *testing.T) {
	account_id := "6adb0bc07b6ec4d992d5e0c051249f61024629a0bb5b264156463c788d9dc661"
	resultIncentivesWithdrawalList, err := GetIncentivesWithdrawalList(account_id, 0, 1534485897,
		0, 20, CONST_Incetives_Withdrawal_Status_Pending)
	if err != nil {
		t.Errorf("GetIncentivesWithdrawalList error:%v", err.Error())
		return
	}

	if resultIncentivesWithdrawalList != nil {
		t.Logf("GetIncentivesWithdrawalList response value:%v", resultIncentivesWithdrawalList)
		if resultIncentivesWithdrawalList.ResultCode != core.CONST_ResultCode_Success {
			t.Errorf("GetIncentivesWithdrawalList response error:%v", resultIncentivesWithdrawalList.ResultMsg)
			return
		}

		if resultIncentivesWithdrawalList.Data != nil {
			t.Logf("GetIncentivesWithdrawalList response AccountTokensData value:%v", resultIncentivesWithdrawalList.Data)
		} else {
			t.Logf("GetIncentivesWithdrawalList response AccountTokensData value don't find ")
		}
	}
}
