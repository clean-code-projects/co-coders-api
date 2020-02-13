package criteria


// Criteria ..
type Criteria struct {
	CollabStyles NamedCriteria
}

// New ...
func New() Criteria {
	return Criteria{CollabStyles:NewCollabStyles()}
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
