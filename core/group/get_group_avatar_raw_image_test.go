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
