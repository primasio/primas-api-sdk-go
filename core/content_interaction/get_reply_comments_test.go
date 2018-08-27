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

package content_interaction

import (
	"testing"

	"github.com/primasio/primas-api-sdk-go/core"
)

func TestGetReplyComments(t *testing.T) {
	comment_id := "02b6218aeae6e6f72b25a1a0302c8edce7514a48b1211a0ce85bc01ad4de0f10"
	resultGetReplyComments, err := GetReplyComments(comment_id)
	if err != nil {
		t.Errorf("GetReplyComments error:%v", err.Error())
		return
	}

	if resultGetReplyComments != nil {
		t.Logf("GetReplyComments response value:%v", resultGetReplyComments)
		if resultGetReplyComments.ResultCode != core.CONST_ResultCode_Success {
			t.Errorf("GetReplyComments response error:%v", resultGetReplyComments.ResultMsg)
			return
		}

		if resultGetReplyComments.Data != nil {
			t.Logf("GetReplyComments response data value:%#v", resultGetReplyComments.Data)
		} else {
			t.Logf("GetReplyComments response data value don't find ")
		}
	}
}
