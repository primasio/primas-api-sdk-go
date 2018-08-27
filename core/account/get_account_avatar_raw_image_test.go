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
)

func TestGetAccountAvatarRawImage(t *testing.T) {
	account_id := "809a85f7ddf8ae5aaa49fe30be10e07e09156dc04166fab98bbd7bb42b2dc26c"
	resultReponse, err := GetAccountAvatarRawImage(account_id)
	if err != nil {
		t.Errorf("GetAccountAvatarRawImage error:%v", err.Error())
		return
	}

	t.Logf("resultReponse:" + string(resultReponse))
}

func TestGetSubAccountAvatarRawImage(t *testing.T) {
	account_id := "809a85f7ddf8ae5aaa49fe30be10e07e09156dc04166fab98bbd7bb42b2dc26c"
	sub_account_id := "a_0001"

	resultReponse, err := GetSubAccountAvatarRawImage(account_id, sub_account_id)
	if err != nil {
		t.Errorf("GetSubAccountAvatarRawImage error:%v", err.Error())
		return
	}

	t.Logf("resultReponse:" + string(resultReponse))
}
