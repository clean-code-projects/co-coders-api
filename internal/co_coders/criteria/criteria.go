package criteria



// CollabStyle ..
type CollabStyle int

const (
	Team CollabStyle = iota + 1
	Pair
	Mob
)

// Criteria ..
type Criteria struct {
	CollabStyles []CollabStyle
}

type Match struct {
	Score float64
}

