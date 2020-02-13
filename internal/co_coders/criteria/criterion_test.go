package criteria

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNamedCriterionCreatesAString(t *testing.T){
	criterion := NewNamedCriterion("Name")
	assert.Equal(t, "NamedCriterion{name: Name}", criterion.String())
}


func TestNamedCriteriaReturnsEmptyIfTheMatchableIsNotANamedCriteria(t *testing.T){
	namedCriteria := NewNamedCriteria("One", "Two", "Three")
	result := namedCriteria.Match(intCriteria{1, 2, 3})
	assert.Equal(t, generateNamedMatches(), result)
}

type intCriteria []int

func (i intCriteria) Match(m Matchable) []Criterion {
	return []Criterion{}
}
