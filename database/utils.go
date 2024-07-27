package database

import (
	"regexp"
	"strings"
)

func Slugify(s string) string {
	// Convert to lowercase
	s = strings.ToLower(s)

	// Remove spaces and replace with dashes
	s = strings.ReplaceAll(s, " ", "-")

	// Remove non-alphanumeric characters
	re := regexp.MustCompile("[^a-z0-9]+")
	s = re.ReplaceAllString(s, "")

	// Remove leading and trailing dashes
	s = strings.Trim(s, "-")

	return s
}
