package time

import "time"

// GetTimestampInMillis returns the current time in milliseconds passed Epoch
func GetTimestampInMillis() int64 {
	return time.Now().UnixNano() / 1000000
}
