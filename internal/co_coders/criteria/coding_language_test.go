package criteria

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMatchWithNoCodingLanguagesReturnsEmpty(t *testing.T){
	codingLanguages := NewCodingLanguages("Java")
	noCodingLanguages := NewCodingLanguages()

	result := codingLanguages.Match(noCodingLanguages)

	assert.Equal(t, generateNamedMatches(), result)
}

func TestMatchWithOneMatchingCodingLanguageReturnsOneMatch(t *testing.T){
	codingLanguages := NewCodingLanguages("Python")
	matchingCodingLanguages := NewCodingLanguages("Python")

	result := codingLanguages.Match(matchingCodingLanguages)

	assert.Equal(t, generateNamedMatches("Python"), result)
}

func TestMatchWithOneSomeMatchingCodingLanguagesReturnsOnlyMatches(t *testing.T){
	codingLanguages := NewCodingLanguages("C#", "JavaScript", "Go")
	matchingCodingLanguages := NewCodingLanguages("C", "C#", "JavaScript")

	result := codingLanguages.Match(matchingCodingLanguages)

	assert.Equal(t, generateNamedMatches("C#", "JavaScript"), result)
}


func generateNamedMatches(names ...string) []Criterion {
	criteria := []Criterion{}
	for _, name := range names {
		criteria = append(criteria, NewNamedCriterion(name))
	}	
	return criteria
}