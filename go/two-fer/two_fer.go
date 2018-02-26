// Package twofer does something amazing.
package twofer

import (
	"fmt"
)

// ShareWith returns a phrase.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
