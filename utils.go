package investec

// EqualError compares two errors and returns true if they are equal.
func EqualError(a, b error) bool {
	if a == nil && b == nil {
		return true
	} else if a == nil || b == nil {
		// either a or b is nil
		return false
	}
	// the errors have the same message
	if a.Error() == b.Error() {
		return true
	}
	return false
}
