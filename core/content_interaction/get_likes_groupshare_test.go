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

func TestGetLikesOfGroupshare(t *testing.T) {
	share_id := "7187ddc476076fce73201ba91d20600e46b8f2d18d828fa7438c2bbd536ba115"
	resultGetLikesOfGroupshare, err := GetLikesOfGroupshare(share_id, "", 0, 20)
	if err != nil {
		t.Errorf("GetLikesOfGroupshare error:%v", err.Error())
		return
	}

	if resultGetLikesOfGroupshare != nil {
		t.Logf("GetLikesOfGroupshare response value:%v", resultGetLikesOfGroupshare)
		if resultGetLikesOfGroupshare.ResultCode != core.CONST_ResultCode_Success {
			t.Errorf("GetLikesOfGroupshare response error:%v", resultGetLikesOfGroupshare.ResultMsg)
			return
		}

		if resultGetLikesOfGroupshare.Data != nil {
			t.Logf("GetLikesOfGroupshare response data value:%#v", resultGetLikesOfGroupshare.Data)
		} else {
			t.Logf("GetLikesOfGroupshare response data value don't find ")
		}
	}
}
