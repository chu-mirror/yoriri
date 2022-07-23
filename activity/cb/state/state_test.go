package state

import (
	"testing"
	"time"
	"github.com/chu-mirror/yoriri/activity/cb/hit"
	"github.com/chu-mirror/yoriri/activity/cb/boss"
)

// the data is from CB june 2022
func init() {
	boss.HP = [5][5]int {
		{6000, 8000, 10000, 12000, 15000},
		{6000, 8000, 10000, 12000, 15000},
		{12000, 14000, 17000, 19000, 22000},
		{22000, 23000, 27000, 29000, 31000},
		{104000, 110000, 125000, 140000, 150000},
	}
}

func TestCurrent(t *testing.T) {
	tests := []struct {
		curDMG  int
		t, l, d int
	}{
		{0, 0, 0, 0},
		{6000, 0, 1, 0},
		{18000, 1, 0, 0},
		{20000, 1, 0, 2000},
		{24000, 1, 1, 0},
	}

	for _, test := range tests {
		records := []hit.Record{
			{"chu", test.curDMG, 0, true, time.Now()},
		}
		tier, l, d := current(records, 0)
		if l != test.l || tier != test.t || d != test.d {
			t.Errorf("%v, but %d %d %d", test, tier, l, d)
		}
	}
}

func TestDone(t *testing.T) {
	tests := []struct {
		initial int
		dmg     int
		ok      bool
		final   int
	}{
		// tier1
		{0, 6000, false, 0},
		// tier2
		{18000, 3000, true, 21000},
		{18000, 6000, false, 18000},
		// tier3
		{60000, 6000, true, 66000},
		{66000, 6000, false, 66000},
		// tier4
		{300000, 11000, true, 311000},
		{311000, 11000, false, 311000},
		// tier5
		{476000, 52000, true, 528000},
		{528000, 52000, false, 528000},
	}
	for _, test := range tests {
		records := []hit.Record {
			{"chu", test.initial, 0, true, time.Now()},
		}
		records, ok := Done(records, "chu", test.dmg, 0, true)
		if ok != test.ok || wholeDamage(records, 0) != test.final {
			t.Errorf("Done %d: %v -> %v, ok = %v", test.dmg, test.initial, wholeDamage(records, 0), ok)
		}
	}
}

func TestKill(t *testing.T) {
	tests := []struct {
		initial, final int
	}{
		{0, 6000},
		{2000, 6000},
		{18000, 24000},
		{60000, 72000},
		{300000, 322000},
		{528000, 580000},
	}
	for _, test := range tests {
		records := []hit.Record {
			{"chu", test.initial, 0, true, time.Now()},
		}
		records = Kill(records, "chu", 0)
		if wholeDamage(records, 0) != test.final {
			t.Errorf("Kill: %v -> %v", test.initial, wholeDamage(records, 0))
		}
	}
}
