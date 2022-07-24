package boss

import (
	"encoding/json"
	"time"
)

// BossNo denote the boss.
type No int64

func IntToNo(n int64) No {
	return No(n - 1)
}

var (
	HP   [5][5]int // the matrix of bosses' HP from tier 1 to 5, boss 1 to 5
	TierLaps = [4]int{3, 7, 20, 8} // laps in tier 1 to 4
)

func init() {
	HP = [5][5]int {
		{6000, 8000, 10000, 12000, 15000},
		{6000, 8000, 10000, 12000, 15000},
		{12000, 14000, 17000, 19000, 22000},
		{22000, 23000, 27000, 29000, 31000},
		{105000, 110000, 125000, 140000, 150000},
	}
}

