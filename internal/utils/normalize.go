package utils

import (
	"regexp"
	"strconv"
	"strings"
)

var uuidRegex = regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`)

// NormalizeLogPath converts dynamic path segments (IDs, UUIDs) to standardized :id placeholders.
// Example: "/users/123" -> "/users/:id", "/posts/abc-123-def-456-ghi" -> "/posts/:id"
func NormalizeLogPath(path string) string {
	if path == "" {
		return path
	}

	segments := strings.Split(path, "/")
	for i, segment := range segments {
		if segment == "" {
			continue
		}

		// Rule A: Numerical check - if segment is all digits, replace with :id
		if isNumeric(segment) {
			segments[i] = ":id"
			continue
		}

		// Rule B: UUID check - if segment matches UUID pattern, replace with :id
		if uuidRegex.MatchString(strings.ToLower(segment)) {
			segments[i] = ":id"
			continue
		}
	}

	return strings.Join(segments, "/")
}

// isNumeric checks if a string contains only digits.
func isNumeric(s string) bool {
	if s == "" {
		return false
	}
	_, err := strconv.ParseInt(s, 10, 64)
	return err == nil
}
