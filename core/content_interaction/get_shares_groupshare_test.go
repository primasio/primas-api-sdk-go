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

func TestGetSharesOfGroupShare(t *testing.T) {
	share_id := "5c7375693afdddd71a4c283bb34e6503ab549fdab8a6a5392a77a12ada7f75c2"
	resultGetShareGroupshareResponse, err := GetSharesOfGroupShare(share_id, "", 0, 20)
	if err != nil {
		t.Errorf("GetContentMetadata error:%v", err.Error())
		return
	}

	if resultGetShareGroupshareResponse != nil {
		t.Logf("GetContentMetadata response value:%v", resultGetShareGroupshareResponse)
		if resultGetShareGroupshareResponse.ResultCode != core.CONST_ResultCode_Success {
			t.Errorf("GetContentMetadata response error:%v", resultGetShareGroupshareResponse.ResultMsg)
			return
		}

		if resultGetShareGroupshareResponse.Data != nil {
			t.Logf("GetContentMetadata response data value:%#v", resultGetShareGroupshareResponse.Data)
		} else {
			t.Logf("GetContentMetadata response data value don't find ")
		}
	}
}
