package criteria

// CollabStyle ..
type CollabStyle int

// CollabStyle Constants ..
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

// Match ..
func (c CollabStyles) Match(stylesToMatch CollabStyles) (matched CollabStyles)  {
	matched = CollabStyles{}
	for _, style := range stylesToMatch {
		if c.HasCollabStyle(style) {
			matched = append(matched, style)
		}
	}
	return matched
}
