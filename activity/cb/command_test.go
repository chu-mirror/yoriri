package cb

import (
	"testing"
)

func TestDamageParsing(t *testing.T) {
	tests := []struct{
		dmg string
		d int
	}{
		{"20.1M", 20100},
		{"40K", 40},
		{"1.11M", 1110},
	}
	for _, test := range tests {
		d_, _ := damageNumber(test.dmg)
		if d_ != test.d {
			t.Errorf("%s -> %d", test.dmg, d_)
		}
	}
}

