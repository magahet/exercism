// Package bob pretends to be a teenager
package bob

import (
	s "strings"
)

// Hey responds like a teenager.
func Hey(remark string) string {
	remark = s.TrimSpace(remark)

	if remark == "" {
		return "Fine. Be that way!"
	}

	yelling := remark != s.ToLower(remark) && remark == s.ToUpper(remark)
	question := s.HasSuffix(remark, "?")

	switch {
	case question && !yelling:
		return "Sure."
	case yelling && !question:
		return "Whoa, chill out!"
	case yelling && question:
		return "Calm down, I know what I'm doing!"
	default:
		return "Whatever."
	}
}
