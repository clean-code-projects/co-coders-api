package criteria


// NewCollabStyle ...
func NewCollabStyle(name string) NamedCriterion {
	return NewNamedCriterion(name)
}

// NewCollabStyles ..
func NewCollabStyles(styles ...Criterion) NamedCriteria {
	return NewNamedCriteria(styles...)
}
