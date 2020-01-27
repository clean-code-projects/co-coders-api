package criteria

// CollabStyle ..
type CollabStyle int

// CollabStyle Constants
const (
	Team CollabStyle = iota + 1
	Pair
	Mob
)

// CollabStyles ..
type CollabStyles []CollabStyle

// HasCollabStyle ..
func (c CollabStyles) HasCollabStyle(collabStyle CollabStyle) bool {
	for _, style := range c {
		if style == collabStyle {
			return true
		}
	}
	return false
}

// CollabStylesMatch ..
type CollabStylesMatch struct {
	Matches CollabStyles
	Score   float64
}

// Match ..
func (c CollabStyles) Match(stylesToMatch CollabStyles) CollabStylesMatch {
	matches := c.collectMatchesFor(stylesToMatch)
	return CollabStylesMatch{
		Matches: matches,
		Score: c.calculateScore(matches),
	}
}

func (c CollabStyles) collectMatchesFor(stylesToMatch CollabStyles) CollabStyles {
	matches := CollabStyles{}
	for _, style := range stylesToMatch {
		if c.HasCollabStyle(style) {
			matches = append(matches, style)
		}
	}
	return matches
}

func (c CollabStyles) calculateScore(matchedStyles CollabStyles) float64 {
	return float64(len(matchedStyles)) / float64(len(c))
}
