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

package account

import (
	"testing"

	"github.com/primasio/primas-api-sdk-go/core"
)

func TestGetAddressMetadata(t *testing.T) {
	address := "0xA0970333D335266382960A320773A16EaaF8d2E2"
	addressMetadata, err := GetAddressMetadata(address)
	if err != nil {
		t.Errorf("GetAddressMetadata error:%v", err.Error())
		return
	}

	if addressMetadata != nil {
		t.Logf("GetAddressMetadata response value:%v", addressMetadata)
		if addressMetadata.ResultCode != core.CONST_ResultCode_Success {
			t.Errorf("GetAddressMetadata response error:%v", addressMetadata.ResultMsg)
			return
		}
		if addressMetadata.Data != nil {
			t.Logf("GetAddressMetadata response value:%v", addressMetadata.Data)
		} else {
			t.Logf("GetAddressMetadata response value don't find ")
		}
	}
}
