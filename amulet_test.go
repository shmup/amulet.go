package amulet

import "testing"

func TestIsAmulet(t *testing.T) {
	tests := []struct {
		input    string
		isAmulet bool
		count    int
	}{
		{"DON'T WORRY.", true, 4},
		{"If you can't write poems,\nwrite me", true, 5},
		{"This is definitely not an amulet!", false, 1},
		{"" + string(make([]byte, 65)), false, 0},
	}

	for _, tt := range tests {
		isAmulet, count := IsAmulet(tt.input)
		if isAmulet != tt.isAmulet {
			t.Errorf("IsAmulet(%q) = %v, want %v", tt.input, isAmulet, tt.isAmulet)
		}
		if count != tt.count {
			t.Errorf("IsAmulet(%q) count = %d, want %d", tt.input, count, tt.count)
		}
	}
}
