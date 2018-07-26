package tool

import (
	"fmt"
	"strconv"

	"github.com/shopspring/decimal"
)

func ConvertToString(v interface{}) string {
	if v == nil {
		return ""
	}
	switch val := v.(type) {
	case string:
		return v.(string)
	case bool:
		return strconv.FormatBool(v.(bool))
	case int:
		return strconv.Itoa(v.(int))
	case int8, int16, int32, int64:
		return strconv.FormatInt(v.(int64), 10)
	case float32:
		return strconv.FormatFloat(v.(float64), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(v.(float64), 'f', -1, 64)
	case []byte:
		return string(v.([]byte))
	case uint8, uint16, uint32, uint64:
		return strconv.FormatUint(v.(uint64), 10)
	case uint:
		return strconv.FormatUint(uint64(val), 10)
	case decimal.Decimal:
		return val.String()
	default:
		return fmt.Sprintf("%s", v)
	}
}
