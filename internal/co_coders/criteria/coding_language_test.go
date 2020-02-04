package criteria

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	C = NewCodingLanguage("C")
	CSharp = NewCodingLanguage("C#")
	Go = NewCodingLanguage("Go")
	Java = NewCodingLanguage("Java")
	JavaScript = NewCodingLanguage("JavaScript")
	Python = NewCodingLanguage("Python")	
)

func TestMatchWithNoCodingLanguagesReturnsEmpty(t *testing.T){
	codingLanguages := NewCodingLanguages(Java)
	noCodingLanguages := NewCodingLanguages()

	result := codingLanguages.Match(noCodingLanguages)

	assert.Equal(t, []Criterion{}, result)
}

func TestMatchWithOneMatchingCodingLanguageReturnsOneMatch(t *testing.T){
	codingLanguages := NewCodingLanguages(Python)
	matchingCodingLanguages := NewCodingLanguages(Python)

	result := codingLanguages.Match(matchingCodingLanguages)

	assert.Equal(t, []Criterion{Python}, result)
}

func TestMatchWithOneSomeMatchingCodingLanguagesReturnsOnlyMatches(t *testing.T){
	codingLanguages := NewCodingLanguages(CSharp, JavaScript, Go)
	matchingCodingLanguages := NewCodingLanguages(C, CSharp, JavaScript)

	result := codingLanguages.Match(matchingCodingLanguages)

	assert.Equal(t, []Criterion{CSharp, JavaScript}, result)
}
