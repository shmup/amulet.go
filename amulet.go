package amulet

import "crypto/sha256"

func IsAmulet(text string) (bool, int) {
	if len(text) == 0 || len(text) > 64 {
		return false, 0
	}

	hash := sha256.Sum256([]byte(text))

	// each byte becomes 2 hex chars
	maxCount := 0
	currentCount := 0

	// process two hex chars per byte directly
	for _, b := range hash {
		// check high nibble
		if b>>4 == 0x8 {
			currentCount++
		} else {
			currentCount = 0
		}
		if currentCount > maxCount {
			maxCount = currentCount
		}

		// check low nibble
		if b&0xF == 0x8 {
			currentCount++
		} else {
			currentCount = 0
		}
		if currentCount > maxCount {
			maxCount = currentCount
		}
	}

	return maxCount >= 4, maxCount
}
