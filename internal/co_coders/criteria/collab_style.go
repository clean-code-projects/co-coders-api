package criteria

import "fmt"

// CollabStyle ..
type CollabStyle struct {
	name string 
}

// String ..
func (c CollabStyle) String() string {
	return fmt.Sprintf("CollabStyle{name: %s}", c.name)
}

// NewCollabStyle ...
func NewCollabStyle(name string) CollabStyle {
	return CollabStyle{name: name}
}

// CollabStyles ..
type CollabStyles []CollabStyle

// NewCollabStyles ..
func NewCollabStyles(styles ...CollabStyle) CollabStyles {
	return append(CollabStyles{}, styles...)
}

// HasCollabStyle ..
func (c CollabStyles) has(collabStyle CollabStyle) bool {
	for _, style := range c {
		if style == collabStyle {
			return true
		}
	}
	return false
}

// Match ..
func (c CollabStyles) Match(targetStyles Matchable) (matches []Criterion)  {
	matches = []Criterion{}
	styles, ok := targetStyles.(CollabStyles)

	if !ok {
		return matches
	}

	for _, style := range styles {
		if c.has(style) {
			matches = append(matches, style)
		}
	}
	return matches
}
