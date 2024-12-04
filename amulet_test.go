package amulet

import (
	"testing"

	"github.com/shmup/amulet.go/fast"
)

func TestIsAmulet(t *testing.T) {
	tests := []struct {
		input    string
		isAmulet bool
		count    int
	}{
		{"DON'T WORRY.", true, 1},
		{"If you can't write poems,\nwrite me", true, 2},
		{"This is definitely not an amulet!", false, 0},
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

	for _, tt := range tests {
		isAmulet, count := fast.IsAmulet(tt.input)
		if isAmulet != tt.isAmulet {
			t.Errorf("fast.IsAmulet(%q) = %v, want %v", tt.input, isAmulet, tt.isAmulet)
		}
		if count != tt.count {
			t.Errorf("fast.IsAmulet(%q) count = %d, want %d", tt.input, count, tt.count)
		}
	}
}
