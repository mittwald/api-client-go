package pointer

// To is a quick&easy helper function to generate a pointer to a static value.
//
// This is required in some cases where optional values are modelled using
// pointers to primitive types (for example, when an "optional boolean" is
// modelled as `*bool`
func To[T any](val T) *T {
	return &val
}
