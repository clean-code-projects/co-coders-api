package criteria


// Criteria ..
type Criteria struct {
	CollabStyles CollabStyles
}

// New ...
func New() Criteria {
	return Criteria{CollabStyles:CollabStyles{}}
}

// Match ..
func (c Criteria) Match(criteria Criteria) Matches {
	collabStyleMatches := c.CollabStyles.Match(criteria.CollabStyles)
	onCollabStylesScore := float64(len(collabStyleMatches)) / float64(len(c.CollabStyles))
	return Matches{Score: onCollabStylesScore}
}

// Matches ..
type Matches struct {
	Score float64
}

// Criterion ..
type Criterion interface {
	String() string
}

// Matchable ..
type Matchable interface {
	Match(m Matchable) []Criterion
}

