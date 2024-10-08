package helpers

import "time"

func CalculateTimeDifferenceMilliseconds(before, to time.Time) int {
	return int(to.Sub(before).Milliseconds())
}
