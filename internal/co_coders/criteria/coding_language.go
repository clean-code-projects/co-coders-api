package criteria

// CodingLanguage ..
type CodingLanguage int

// CodingLanguage constants
const (
	AssemblyLanguage CodingLanguage = iota
	Bash
	C 
	Clojure
	CPP
	CSharp
	Dart
	Erlang
	FSharp
	Go
	Haskell
	Java
	JavaScript
	Kotlin
	Lisp
	ObjectiveC
	ObjectPascal
	Perl
	PHP
	Python
	R
	Ruby
	Rust
	Scala
	Swift
	SQL
	TypeScript
	VisualBasic
)

// CodingLanguages ..
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

// Criterion ..
func (c CodingLanguages) Match(targetLanguages CodingLanguages) CodingLanguages {
	matches := CodingLanguages{}
	for _, language := range targetLanguages {
		if c.HasCodingLanguage(language){
			matches = append(matches, language)
		}
	}
	return matches
}
