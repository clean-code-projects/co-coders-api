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
	collabStylesMatch := c.CollabStyles.Match(criteria.CollabStyles)
	return Matches{Score: collabStylesMatch.Score}
}

// Matches ..
type Matches struct {
	Score float64
}
