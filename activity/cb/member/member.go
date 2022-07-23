package member

import (
	"github.com/chu-mirror/yoriri/activity/cb/boss"
	"github.com/chu-mirror/yoriri/activity/cb/hit"
	"github.com/chu-mirror/yoriri/activity/cb/state"
)

type stateHitting struct {
	hitting       bool
	full bool
	bossInHitting boss.No
}

var (
	Records []hit.Record
	memberState = make(map[string]*stateHitting)
	lock        [5]bool
)

const (
	HitSuccess = iota
	HitLockedFail
	HitInHittingFail
	HitNoHittingFail
	HitIllegalHpFail
)

func Hit(uid string, target boss.No, sync bool) int {
	_, ok := memberState[uid]
	if !ok {
		memberState[uid] = &stateHitting{
			hitting:       false,
			full: true,
			bossInHitting: 0,
		}
	}
	if memberState[uid].hitting {
		return HitInHittingFail
	}

	if lock[target] && !sync {
		return HitLockedFail
	}

	lock[target] = true

	s := memberState[uid]
	s.hitting = true
	s.bossInHitting = target
	return HitSuccess
}

func Dump(uid string, target boss.No, sync bool) int {
	if memberState[uid].hitting {
		return HitInHittingFail
	}
	if lock[target] && !sync {
		return HitLockedFail
	}
	lock[target] = true

	s := memberState[uid]
	s.hitting = true
	s.full = false
	s.bossInHitting = target
	return HitSuccess
}

func Done(uid string, dmg int) int {
	s, ok := memberState[uid]
	if !ok || !memberState[uid].hitting {
		return HitNoHittingFail
	}
	target := s.bossInHitting

	Records, ok = state.Done(Records, uid, dmg, target, s.full)
	if !ok {
		return HitIllegalHpFail
	}
	lock[target] = false
	s.hitting = false
	return HitSuccess
}

func Kill(uid string) int {
	s, ok := memberState[uid]
	if !ok || !memberState[uid].hitting {
		return HitNoHittingFail
	}
	target := s.bossInHitting
	Records = state.Kill(Records, uid, target)
	lock[target] = false
	s.hitting = false
	return HitSuccess
}

