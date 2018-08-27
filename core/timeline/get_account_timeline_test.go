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

package timeline

import (
	"testing"

	"github.com/primasio/primas-api-sdk-go/core"
)

func TestGetAccountTimeline(t *testing.T) {
	account_id := "2def1e68569f09439f2b5ad2db5ef15196411fc8d16829e02e78af20c6fd8d6c"
	resultGetAccountTimeline, err := GetAccountTimeline(account_id, 0, 20)
	if err != nil {
		t.Errorf("GetAccountTimeline error:%v", err.Error())
		return
	}

	if resultGetAccountTimeline != nil {
		t.Logf("GetAccountTimeline response value:%v", resultGetAccountTimeline)
		if resultGetAccountTimeline.ResultCode != core.CONST_ResultCode_Success {
			t.Errorf("GetAccountTimeline response error:%v", resultGetAccountTimeline.ResultMsg)
			return
		}

		if resultGetAccountTimeline.Data != nil {
			t.Logf("GetAccountTimeline response data value:%#v", resultGetAccountTimeline.Data)
		} else {
			t.Logf("GetAccountTimeline response data value don't find ")
		}
	}
}
