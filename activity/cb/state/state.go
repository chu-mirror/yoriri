// Package state handle the state of bosses
package state

import (
	"github.com/chu-mirror/yoriri/activity/cb/hit"
	"github.com/chu-mirror/yoriri/activity/cb/boss"
)


// wholeDamage return the whole damage on a boss in current records
func wholeDamage(records []hit.Record, target boss.No) (w int) {
	for _, r := range records {
		if (r.Boss == target) {
			w += r.Damage
		}
	}
	return
}

// current returns current progress of target.
func current(records []hit.Record, target boss.No) (t, l, d int) {
	d = wholeDamage(records, target)

	var tn int
	for t, tn = range boss.TierLaps {
		if d < tn*boss.HP[t][target] {
			l = d / boss.HP[t][target]
			d -= l * boss.HP[t][target]
			return
		} else {
			d -= tn * boss.HP[t][target]
		}
	}
	t = 4
	l = d / boss.HP[t][target]
	d -= l * boss.HP[t][target]
	return
}

// Done records damage to target, return false if the damage is illegal.
// Whether illegal or not is depended on whether this damage will kill current target.
func Done(records []hit.Record, uid string, dmg int, target boss.No, full bool) ([]hit.Record, bool) {
	t, _, d := current(records, target)
	if d+dmg >= boss.HP[t][target] {
		return records, false
	}
	return hit.AppendRecord(records, uid, dmg, target, full), true
}

// Kill records damage to target in case current boss is killed.
func Kill(records []hit.Record, uid string, target boss.No) []hit.Record {
	t, _, d := current(records, target)
	return hit.AppendRecord(records, uid, boss.HP[t][target]-d, target, false)
}
