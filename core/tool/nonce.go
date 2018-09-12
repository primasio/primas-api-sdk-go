package tool

import (
	"strings"

	"github.com/pborman/uuid"
)

func GetNonce() string {
	return strings.Replace(uuid.New(), "-", "", -1)
}
