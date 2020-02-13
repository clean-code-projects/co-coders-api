package criteria


// NewCodingLanguages ..
func NewCodingLanguages(languages ...string) NamedCriteria {
	return NewNamedCriteria(languages...)
}
