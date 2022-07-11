package hit

import(
	"github.com/chu-mirror/yoriri/activity/cb/state"
	"github.com/chu-mirror/yoriri/activity"
	"github.com/bwmarrin/discordgo"
)

type stateHitting struct {
	hitted int
	hitting bool
	bossInHitting state.BossNo
}

var (
	memberState = make(map[string]*stateHitting)
	lock [5]bool
)

const (
	HitSuccess = iota
	HitLockedFail
	HitInHittingFail
	HitNoHittingFail
	HitIllegalHpFail
)

func Hit(uid string, boss state.BossNo, sync bool) int {
	if lock[boss] && !sync {
		return HitLockedFail
	}
	_, ok := memberState[uid]
	if !ok {
		memberState[uid] = &stateHitting{
			hitted: 0,
			hitting: false,
			bossInHitting: 0,
		}
	}
	if memberState[uid].hitting {
		return HitInHittingFail
	}
	lock[boss] = true

	s := memberState[uid]
	s.hitted++
	s.hitting = true
	s.bossInHitting = boss
	return HitSuccess
}

func Dump(uid string, boss state.BossNo, sync bool) int {
	if lock[boss] && !sync {
		return HitLockedFail
	}
	if memberState[uid].hitting {
		return HitInHittingFail
	}
	lock[boss] = true

	s := memberState[uid]
	s.hitting = true
	s.bossInHitting = boss
	return HitSuccess
}

func Done(uid string, dmg int) int {
	if !memberState[uid].hitting {
		return HitNoHittingFail
	}
	boss := memberState[uid].bossInHitting
	ok := state.Done(boss, dmg)
	if !ok {
		return HitIllegalHpFail
	}
	lock[boss] = false
	s := memberState[uid]
	s.hitting = false 
	return HitSuccess
}

func Kill(uid string) int {
	if !memberState[uid].hitting {
		return HitNoHittingFail
	}
	boss := memberState[uid].bossInHitting
	state.Kill(boss)
	lock[boss] = false
	s := memberState[uid]
	s.hitting = false 
	return HitSuccess
}

func init() {
	activity.RegisterInitialization(func(s *discordgo.Session) error {
		return nil
	})
}
