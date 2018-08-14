package bytes

import "log"

// CheckBounds --
func CheckBounds(index int32, length int32, limit int32) {
	if (index + length) > limit {
		log.Fatalf("ou of ounds. index: %d, length: %d, capacity: %d", index, length, limit)
	}
}
