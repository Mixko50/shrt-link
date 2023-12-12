package text

import "regexp"

func IsAlphanumeric(s string) bool {
	regex := regexp.MustCompile(`^[a-zA-Z0-9-_]+$`)
	return regex.MatchString(s)
}
