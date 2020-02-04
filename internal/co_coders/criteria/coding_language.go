package criteria


// NewCodingLanguage .. 
func NewCodingLanguage(name string) NamedCriterion {
	return NewNamedCriterion(name)
} 

// NewCodingLanguages ..
func NewCodingLanguages(languages ...Criterion) NamedCriteria {
	return NewNamedCriteria(languages...)
}
