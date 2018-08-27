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

func TestGetAccountTokenMetadata(t *testing.T) {
	account_id := "809a85f7ddf8ae5aaa49fe30be10e07e09156dc04166fab98bbd7bb42b2dc26c"
	accountMetadata, err := GetAccountTokenMetadata(account_id)
	if err != nil {
		t.Errorf("GetAccountTokenMetadata error:%v", err.Error())
		return
	}

	if accountMetadata != nil {
		t.Logf("GetAccountTokensData response value:%v", accountMetadata)
		if accountMetadata.ResultCode != core.CONST_ResultCode_Success {
			t.Errorf("GetAccountTokensData response error:%v", accountMetadata.ResultMsg)
			return
		}
		if accountMetadata.Data != nil {
			t.Logf("GetAccountTokenMetadata response value:%v", accountMetadata.Data)
		} else {
			t.Logf("GetAccountTokenMetadata response value don't find ")
		}
	}
}

func TestGetSubAccountTokenMetadata(t *testing.T) {
	account_id := "809a85f7ddf8ae5aaa49fe30be10e07e09156dc04166fab98bbd7bb42b2dc26c"
	sub_account_id := "a1ad1adf1"
	subAccountMetadata, err := GetSubAccountTokenMetadata(account_id, sub_account_id)
	if err != nil {
		t.Errorf("GetSubAccountTokenMetadata error:%v", err.Error())
		return
	}

	if subAccountMetadata != nil {
		t.Logf("GetSubAccountTokenMetadata response value:%v", subAccountMetadata)
		if subAccountMetadata.ResultCode != core.CONST_ResultCode_Success {
			t.Errorf("GetSubAccountTokenMetadata response error:%v", subAccountMetadata.ResultMsg)
			return
		}
		if subAccountMetadata.Data != nil {
			t.Logf("GetSubAccountTokenMetadata response value:%v", subAccountMetadata.Data)
		} else {
			t.Logf("GetSubAccountTokenMetadata response value don't find ")
		}
	}
}
