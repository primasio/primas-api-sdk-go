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

func TestGetGroupAvatarMetadata(t *testing.T) {
	group_id := "88da2092cd8230c6dbbab6b555e08b5b0eb1f7523055d0df9230399f7bbd858e"
	resultGetGroupAvatarMetadata, err := GetGroupAvatarMetadata(group_id)
	if err != nil {
		t.Errorf("GetGroupAvatarMetadata error:%v", err.Error())
		return
	}

	if resultGetGroupAvatarMetadata != nil {
		t.Logf("GetGroupAvatarMetadata response value:%v", resultGetGroupAvatarMetadata)
		if resultGetGroupAvatarMetadata.ResultCode != core.CONST_ResultCode_Success {
			t.Errorf("GetGroupAvatarMetadata response error:%v", resultGetGroupAvatarMetadata.ResultMsg)
			return
		}

		if resultGetGroupAvatarMetadata.Data != nil {
			t.Logf("GetGroupAvatarMetadata response value:%#v", resultGetGroupAvatarMetadata.Data)
		} else {
			t.Logf("GetGroupAvatarMetadata response value don't find ")
		}
	}
}
