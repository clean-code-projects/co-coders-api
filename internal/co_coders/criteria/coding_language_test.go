package criteria

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMatchWithNoCodingLanguagesReturnsEmpty(t *testing.T){
	codingLanguages := CodingLanguages{Java}
	noCodingLanguages := CodingLanguages{}

	result := codingLanguages.Match(noCodingLanguages)

	assert.Equal(t, noCodingLanguages, result)
}

func TestMatchWithOneMatchingCodingLanguageReturnsOneMatch(t *testing.T){
	codingLanguages := CodingLanguages{Python}
	matchingCodingLanguages := CodingLanguages{Python}

	result := codingLanguages.Match(matchingCodingLanguages)

	assert.Equal(t, matchingCodingLanguages, result)
}

func TestMatchWithOneSomeMatchingCodingLanguagesReturnsOnlyMatches(t *testing.T){
	codingLanguages := CodingLanguages{CSharp, JavaScript, Go}
	matchingCodingLanguages := CodingLanguages{C, CSharp, JavaScript}

	result := codingLanguages.Match(matchingCodingLanguages)

	assert.Equal(t, CodingLanguages{CSharp, JavaScript}, result)
}
