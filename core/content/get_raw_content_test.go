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
	"io/ioutil"
	"strconv"
	"testing"
)

func TestGetRawContent(t *testing.T) {
	//content_id := "0b0d27adf09b17e4511e210ddeec0cf3136cbc2d214d6d84a80fb37577957c08"
	content_id := "e8e565872c139fe06eb68c712aec000005475adc29d30ddcf8b0b132abc55eee"
	resultRawContent, err := GetRawContent(content_id)
	if err != nil {
		t.Errorf("GetRawContent error:%v", err.Error())
		return
	}

	if resultRawContent != nil {
		t.Logf("raw content:%v", resultRawContent.Body)

		defer resultRawContent.Body.Close()
		contents, err := ioutil.ReadAll(resultRawContent.Body)
		if err != nil {
			t.Errorf("ioutil.ReadAll %v", err.Error())
		}

		if resultRawContent.StatusCode != 200 {
			t.Errorf("response StatusCode error:%v" + strconv.Itoa(resultRawContent.StatusCode))
		}

		if len(contents) == 0 {
			t.Errorf("response body is empty")
		}

		t.Logf("raw body:%v", string(contents))
	}
}
