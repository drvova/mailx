package service

import "strings"

// isUniqueConstraintError checks if the error is a unique constraint violation.
func isUniqueConstraintError(err error) bool {
	if err == nil {
		return false
	}
	return strings.Contains(err.Error(), "UNIQUE constraint failed") ||
		strings.Contains(err.Error(), "Duplicate entry") ||
		strings.Contains(err.Error(), "1062")
}
