package hit

import (
	"time"

	"github.com/chu-mirror/yoriri/activity/cb/boss"
)

// A record of "a member do some damage on a boss at some time"
type Record struct {
	MemberId string
	Damage int
	Boss boss.No
	Full bool
	Time time.Time
}

func AppendRecord(records []Record, uid string, dmg int, target boss.No, full bool) []Record {
	loc := time.FixedZone("UTC+4", 4*60*60)
	return append(records, Record{uid, dmg, target, full, time.Now().In(loc)})
}
	

