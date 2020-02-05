package criteria


// NewCodingLanguages ..
func NewCodingLanguages(languages ...string) NamedCriteria {
	criteria := []Criterion{}
	for _, name := range languages {
		criteria = append(criteria, NewNamedCriterion(name))
	}
	return NewNamedCriteria(criteria...)
}
