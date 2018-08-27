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

package system

import (
	"testing"

	"github.com/primasio/primas-api-sdk-go/core"
)

func TestGetSystemParameters(t *testing.T) {
	resultGetSystemParameters, err := GetSystemParameters()
	if err != nil {
		t.Errorf("GetSystemParameters error:%v", err.Error())
		return
	}

	if resultGetSystemParameters != nil {
		t.Logf("GetSystemParameters response value:%v", resultGetSystemParameters)
		if resultGetSystemParameters.ResultCode != core.CONST_ResultCode_Success {
			t.Errorf("GetSystemParameters response error:%v", resultGetSystemParameters.ResultMsg)
			return
		}

		if resultGetSystemParameters.Data != nil {
			t.Logf("GetSystemParameters response data value:%#v", resultGetSystemParameters.Data)
		} else {
			t.Logf("GetSystemParameters response data value don't find ")
		}
	}

	t.Logf("LockAmountContent:%v", resultGetSystemParameters.Data.LockAmountContent)
	t.Logf("LockPeriodContent:%v", resultGetSystemParameters.Data.LockPeriodContent)
	t.Logf("LockAmountGroupJoin:%v", resultGetSystemParameters.Data.LockAmountGroupJoin)
	t.Logf("LockAmountGroupCreate:%v", resultGetSystemParameters.Data.LockAmountGroupCreate)
	t.Logf("ConsumeAmountReport:%v", resultGetSystemParameters.Data.ConsumeAmountReport)
}
