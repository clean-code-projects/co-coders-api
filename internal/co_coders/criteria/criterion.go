package criteria

import "fmt"

// Criterion ..
type Criterion interface {
	String() string
}

// Matchable ..
type Matchable interface {
	Match(m Matchable) []Criterion
}

// NamedCriterion ..
type NamedCriterion struct{
	name string
}

// NewNamedCriterion ..
func NewNamedCriterion(name string) NamedCriterion {
	return NamedCriterion{name: name}
}

// String ..
func (c NamedCriterion) String() string {
	return fmt.Sprintf("NamedCriterion{name: %s}", c.name)
}

// NamedCriteria ...
type NamedCriteria []Criterion 


// NewNamedCriteria ...
func NewNamedCriteria(names ...string) NamedCriteria {
	criteria := []Criterion{}
	for _, name := range names {
		criteria = append(criteria, NewNamedCriterion(name))
	}
	return append(NamedCriteria{}, criteria...)
}

func (c NamedCriteria) has(targetCriterion Criterion) bool {
	for _, criterion := range c {
		if criterion == targetCriterion {
			return true
		}
	}
	return false
}

// Match ..
func (c NamedCriteria) Match(targetCriteria Matchable) []Criterion {
	matches := []Criterion{}
	criteria, ok := targetCriteria.(NamedCriteria)

	if !ok {
		return matches
	}
	
	for _, criteria := range criteria {
		if c.has(criteria){
			matches = append(matches, criteria)
		}
	}
	
	return matches
}
