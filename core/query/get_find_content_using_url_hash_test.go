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

func TestFindContentUsingUrlOrHash(t *testing.T) {
	urlParam := "https://staging.primas.io/raw/76caf1af5befd80e6030eefccaf78fc19bd28b2ba895697bedf166196d1f5933"
	hashParam := "160dc52a9244e9e54084219ce40c9b62b19e05de0bc77792ca48f616e8aa18f9"

	resultFindContentUsingUrlOrHash, err := FindContentUsingUrlOrHash(urlParam, hashParam)
	if err != nil {
		t.Errorf("FindContentUsingUrlOrHash error:%v", err.Error())
		return
	}

	if resultFindContentUsingUrlOrHash != nil {
		t.Logf("FindContentUsingUrlOrHash response value:%v", resultFindContentUsingUrlOrHash)
		if resultFindContentUsingUrlOrHash.ResultCode != core.CONST_ResultCode_Success {
			t.Errorf("FindContentUsingUrlOrHash response error:%v", resultFindContentUsingUrlOrHash.ResultMsg)
			return
		}

		if resultFindContentUsingUrlOrHash.Data != nil {
			t.Logf("FindContentUsingUrlOrHash response data value:%#v", resultFindContentUsingUrlOrHash.Data)
		} else {
			t.Logf("FindContentUsingUrlOrHash response data value don't find ")
		}
	}
}
