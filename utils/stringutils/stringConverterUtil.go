package stringutils

import (
	"strings"
)

func SpaceToBarsAndLower(text string) string {
	return strings.ToLower(strings.ReplaceAll(text, " ", "-"))
}

func BeforeFirstBars(text string) string {
	return strings.SplitN(text, "-", 2)[0]
}

func RemoveSpaceAndLower(text string) string {
	return strings.ToLower(strings.ReplaceAll(text, " ", ""))
}
