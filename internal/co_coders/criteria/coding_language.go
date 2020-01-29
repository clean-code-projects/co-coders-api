package criteria

// CodingLanguage ..
type CodingLanguage int

// CodingLanguage constants
const (
	C CodingLanguage = iota
	CSharp
	Go
	Java
	JavaScript
	Python
)

// CodingLanguage
type CodingLanguages []CodingLanguage

// HasCodingLanguage ..
func (c CodingLanguages) HasCodingLanguage(targetLanguage CodingLanguage) bool {
	for _, language := range c {
		if language == targetLanguage {
			return true
		}
	}
	return false
}

// Match ..
func (c CodingLanguages) Match(targetLanguages CodingLanguages) CodingLanguages {
	matches := CodingLanguages{}
	for _, language := range targetLanguages {
		if c.HasCodingLanguage(language){
			matches = append(matches, language)
		}
	}
	return matches
}
