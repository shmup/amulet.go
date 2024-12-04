package amulet

import (
	"crypto/sha256"
)

func IsAmulet(text string) (bool, int) {
	if len(text) == 0 || len(text) > 64 {
		return false, 0 // Changed to return 0 for invalid cases
	}

	hash := sha256.Sum256([]byte(text))
	maxCount := 0
	count := 0

	for _, b := range hash {
		// check high nibble
		if b>>4 == 0x8 {
			count++
		} else {
			count = 0
		}
		if count > maxCount {
			maxCount = count
		}

		// check low nibble
		if b&0xF == 0x8 {
			count++
		} else {
			count = 0
		}
		if count > maxCount {
			maxCount = count
		}
	}

	if maxCount >= 4 {
		return true, maxCount - 3
	}
	return false, 0
}
