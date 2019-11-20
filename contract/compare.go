package contract

import "strings"

// CompareNoCap ...
func CompareNoCap(a, b string) int {
	a = strings.ToUpper(a)
	b = strings.ToUpper(b)
	return strings.Compare(a, b)
}
