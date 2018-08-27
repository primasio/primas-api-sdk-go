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

package query

import (
	"testing"

	"github.com/primasio/primas-api-sdk-go/core"
)

func TestGetAll(t *testing.T) {
	page := 0
	page_size := 20
	text := "dtcp"
	qtype := "all"
	category := ""

	resultGetAll, err := GetAll(page, page_size, text, qtype, category)
	if err != nil {
		t.Errorf("GetAll error:%v", err.Error())
		return
	}

	if resultGetAll != nil {
		t.Logf("GetAll response value:%v", resultGetAll)
		if resultGetAll.ResultCode != core.CONST_ResultCode_Success {
			t.Errorf("GetAll response error:%v", resultGetAll.ResultMsg)
			return
		}

		if resultGetAll.Data != nil {
			t.Logf("GetAll response data value:%#v", resultGetAll.Data)
		} else {
			t.Logf("GetAll response data value don't find ")
		}
	}
}
