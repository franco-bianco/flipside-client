package flipside

import (
	"fmt"
	"strconv"
	"time"
)

func parseTimestamp(ts string) time.Time {
	t, err := time.Parse(time.RFC3339, ts)
	if err != nil {
		// Handle the error appropriately, maybe log it
		return time.Time{}
	}
	return t
}

func parseFloat64(v interface{}) (float64, error) {
	switch value := v.(type) {
	case float64:
		return value, nil
	case string:
		f, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return 0, fmt.Errorf("failed to parse float64: %w", err)
		}
		return f, nil
	default:
		return 0, fmt.Errorf("unsupported type for float64 conversion: %T", v)
	}
}
