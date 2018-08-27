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

package group

import (
	"testing"

	"github.com/primasio/primas-api-sdk-go/core"
)

func TestGetGroupMetadata(t *testing.T) {
	group_id := "550b81d5d34a7f211897f8ed04ae13c96c0093b0b16645a8d7ee36edbaa4db6b"
	resultGetGroupMetadataResponse, err := GetGroupMetadata(group_id, "")
	if err != nil {
		t.Errorf("GetGroupMetadata error:%v", err.Error())
		return
	}

	if resultGetGroupMetadataResponse != nil {
		t.Logf("GetGroupMetadata response value:%v", resultGetGroupMetadataResponse)
		if resultGetGroupMetadataResponse.ResultCode != core.CONST_ResultCode_Success {
			t.Errorf("GetGroupMetadata response error:%v", resultGetGroupMetadataResponse.ResultMsg)
			return
		}

		if resultGetGroupMetadataResponse.Data != nil {
			t.Logf("GetGroupMetadata response data value:%#v", resultGetGroupMetadataResponse.Data)
		} else {
			t.Logf("GetGroupMetadata response data value don't find ")
		}
	}
}
