package boss

// BossNo denote the boss.
type No int64

func IntToNo(n int64) No {
	return No(n - 1)
}

var (
	HP   [5][5]int // the matrix of bosses' HP from tier 1 to 5, boss 1 to 5
	TierLaps = [4]int{3, 7, 20, 8} // laps in tier 1 to 4
)

