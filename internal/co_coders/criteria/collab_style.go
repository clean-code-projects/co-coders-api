package criteria


// NewCollabStyles ..
func NewCollabStyles(styles ...string) NamedCriteria {
	criteria := []Criterion{}
	for _, name := range styles {
		criteria = append(criteria, NewNamedCriterion(name))
	}
	return NewNamedCriteria(criteria...)
}
