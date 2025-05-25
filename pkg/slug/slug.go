package slug

import (
	"fmt"
	"strings"
	"unicode"
)

func GenerateUniqueSlug(title string, exists func(slug string) bool) string {
	// Step 1: Format the title into a base slug
	base := strings.ToLower(title)
	base = strings.ReplaceAll(base, " ", "-")
	base = strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) || unicode.IsDigit(r) || r == '-' {
			return r
		}
		return -1
	}, base)

	// Step 2: Add numeric suffix if slug already exists
	slug := base
	counter := 1
	for exists(slug) {
		slug = fmt.Sprintf("%s-%d", base, counter)
		counter++
	}

	return slug
}
