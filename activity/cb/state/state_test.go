package state

import "testing"

// the data is from CB june 2022
func init() {
	setHP(0, [5]int{6000, 8000, 10000, 12000, 15000});
	setHP(1, [5]int{6000, 8000, 10000, 12000, 15000});
	setHP(2, [5]int{12000, 14000, 17000, 19000, 22000});
	setHP(3, [5]int{22000, 23000, 27000, 29000, 31000});
	setHP(4, [5]int{104000, 110000, 125000, 140000, 150000});
}

func TestCurrent(t *testing.T) {
	tests := []struct {
		curDMG int
		l, t, d int
	}{
		{0, 0, 0, 0},
		{6000, 1, 0, 6000},
		{18000, 3, 1, 0},
	}

	for _, test := range tests {
		wholeDMG[0] = test.curDMG
		l, tier, d := current(0)
		if l != test.l || tier != test.t || d != test.d {
			t.Errorf("%v, but %d %d %d", test, l, tier, d)
		}
	}
}

func TestDone(t *testing.T) {
	tests := []struct {
		initial [5]int
		boss BossNo
		dmg int
		ok bool
		final [5]int
	}{
		// tier1
		{[5]int{0, 0, 0, 0, 0}, 0, 6000, false, [5]int{0, 0, 0, 0, 0}},
		{[5]int{0, 2000, 0, 0, 0}, 1, 4000, true, [5]int{0, 6000, 0, 0, 0}},
		{[5]int{0, 2000, 0, 0, 0}, 1, 7000, false, [5]int{0, 2000, 0, 0, 0}},
		// tier2
		{[5]int{18000, 0, 0, 0, 0}, 0, 3000, true, [5]int{21000, 0, 0, 0, 0}},
		{[5]int{18000, 0, 0, 0, 0}, 0, 6000, false, [5]int{18000, 0, 0, 0, 0}},
		// tier3
		{[5]int{60000, 0, 0, 0, 0}, 0, 6000, true, [5]int{66000, 0, 0, 0, 0}},
		{[5]int{66000, 0, 0, 0, 0}, 0, 6000, false, [5]int{66000, 0, 0, 0, 0}},
		// tier4
		{[5]int{300000, 0, 0, 0, 0}, 0, 11000, true, [5]int{311000, 0, 0, 0, 0}},
		{[5]int{311000, 0, 0, 0, 0}, 0, 11000, false, [5]int{311000, 0, 0, 0, 0}},
		// tier5
		{[5]int{476000, 0, 0, 0, 0}, 0, 52000, true, [5]int{528000, 0, 0, 0, 0}},
		{[5]int{528000, 0, 0, 0, 0}, 0, 52000, false, [5]int{528000, 0, 0, 0, 0}},
	}
	for _, test := range tests {
		wholeDMG = test.initial
		ok := Done(test.boss, test.dmg)
		if ok != test.ok || wholeDMG != test.final {
			t.Errorf("Done(%d, %d): %v -> %v, ok = %v", test.boss, test.dmg, test.initial, wholeDMG, ok)
		}
	}
}

func TestKill(t *testing.T) {
	wholeDMG = [5]int{0, 0, 0, 0, 0}

	final := []int{6000, 6000, 24000, 72000, 322000, 580000}
	for i, dmg := range []int{0, 2000, 18000, 60000, 300000, 528000} {
		wholeDMG[0] = dmg
		Kill(0)
		if wholeDMG[0] != final[i] {
			t.Errorf("Kill(0): %v -> %v", dmg, wholeDMG[0])
		}
	}
}
