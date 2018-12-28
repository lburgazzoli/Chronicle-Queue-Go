package bytes

import "log"

// CheckBounds --
func CheckBounds(index int64, length int64, limit int64) {
	if (index + length) > limit {
		log.Fatalf("ou of bounds. index: %d, length: %d, capacity: %d", index, length, limit)
	}
}
