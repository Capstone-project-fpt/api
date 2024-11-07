package util

import "github.com/thoas/go-funk"

func IsSameElementTwoArray[K comparable](a []K, b []K) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if !funk.Contains(b, a[i]) {
			return false
		}
	}

	return true
}