package split

import "strings"

func Split(str, seperator string) []string {
	return strings.Split(str, seperator)
}