package regex

import "regexp"

const ALPHA_NUM_REGEX = "^[a-zA-Z0-9]+$"

func IsAlphaNumeric(s string) bool {
	var valid = regexp.MustCompile(ALPHA_NUM_REGEX)
	return valid.MatchString(s)
}
