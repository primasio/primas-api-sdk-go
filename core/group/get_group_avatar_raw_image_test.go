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
)

func TestGetGroupAvatarRawImage(t *testing.T) {
	group_id := "88da2092cd8230c6dbbab6b555e08b5b0eb1f7523055d0df9230399f7bbd858e"
	resultReponse, err := GetGroupAvatarRawImage(group_id)
	if err != nil {
		t.Errorf("GetGroupAvatarMetadata error:%v", err.Error())
		return
	}

	t.Logf("resultReponse:" + string(resultReponse))
}
