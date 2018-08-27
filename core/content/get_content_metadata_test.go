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

package content

import (
	"testing"

	"github.com/primasio/primas-api-sdk-go/core"
)

func TestGetContentMetadata(t *testing.T) {
	content_id := "0b0d27adf09b17e4511e210ddeec0cf3136cbc2d214d6d84a80fb37577957c08"
	resultGetContentMetadataResponse, err := GetContentMetadata(content_id)
	if err != nil {
		t.Errorf("GetContentMetadata error:%v", err.Error())
		return
	}

	if resultGetContentMetadataResponse != nil {
		t.Logf("GetContentMetadata response value:%v", resultGetContentMetadataResponse)
		if resultGetContentMetadataResponse.ResultCode != core.CONST_ResultCode_Success {
			t.Errorf("GetContentMetadata response error:%v", resultGetContentMetadataResponse.ResultMsg)
			return
		}

		if resultGetContentMetadataResponse.Data != nil {
			t.Logf("GetContentMetadata response AccountTokensData value:%#v", resultGetContentMetadataResponse.Data)
		} else {
			t.Logf("GetContentMetadata response AccountTokensData value don't find ")
		}
	}
}
