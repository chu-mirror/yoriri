// Package state save and change the state of bosses.
package state

import (
//	"encoding/json"
)

// BossNo denote the boss.
type BossNo int64

func IntToBossNo(n int64) BossNo{
	return BossNo(n-1)
}

var (
	bossHP [5][5]int // the matrix of bosses' HP from tier 1 to 5, boss 1 to 5
	wholeDMG [5]int // the whole damage to boss 1 to 5 done by members

	tierLaps = [4]int{3, 7, 20, 8} // laps in tier 1 to 4
)

func setHP(boss BossNo, hp [5]int) {
	bossHP[boss] = hp
}

// current returns current progress of boss.
func current(boss BossNo) (l, t, dmg int) {
	dmg = wholeDMG[boss]
	for t = range tierLaps {
		if dmg < tierLaps[t]*bossHP[t][boss] {
			l += dmg / bossHP[t][boss]
			return
		} else {
			dmg -= tierLaps[t]*bossHP[t][boss]
			l += tierLaps[t]
		}
	}
	l += dmg / bossHP[4][boss]
	t = 4
	return
}

// Done records damage to boss, return true if the damage is legal,
// false if illegal.  Whether legal or not is depended by whether this
// damage will kill current boss.
func Done(boss BossNo, dmg int) bool {
	_, t, d := current(boss)
	if (d+dmg)/bossHP[t][boss] != d/bossHP[t][boss] {
		return false
	}
	wholeDMG[boss] += dmg
	return true
}

// Kill records damage to boss in case current boss is killed.
func Kill(boss BossNo) {
	_, t, dmg := current(boss)
	wholeDMG[boss] += (dmg/bossHP[t][boss]+1)*bossHP[t][boss] - dmg
}

