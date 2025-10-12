package utils

import (
	"strings"
)

func SpaceToBars(text string) string {
	return strings.ReplaceAll(text, " ", "-")
}
