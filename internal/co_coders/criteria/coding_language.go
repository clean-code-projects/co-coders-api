package criteria

import "fmt"

// CodingLanguage ..
type CodingLanguage struct {
	name string
}

// String ..
func (c CodingLanguage) String() string {
	return fmt.Sprintf("CodingLanguage{name: %s}", c.name)
}

// NewCodingLanguage .. 
func NewCodingLanguage(name string) CodingLanguage {
	return CodingLanguage{name: name}
} 

// CodingLanguages ..
type CodingLanguages []CodingLanguage

// NewCodingLanguages ..
func NewCodingLanguages(languages ...CodingLanguage) CodingLanguages {
	return append(CodingLanguages{}, languages...)
}

func (c CodingLanguages) has(targetLanguage CodingLanguage) bool {
	for _, language := range c {
		if language == targetLanguage {
			return true
		}
	}
	return false
}

// Match ..
func (c CodingLanguages) Match(targetLanguages Matchable) []Criterion {
	matches := []Criterion{}

	languages, ok := targetLanguages.(CodingLanguages)
	if !ok {
		return matches
	}

	for _, language := range languages {
		if c.has(language){
			matches = append(matches, language)
		}
	}
	return matches
}
