package math

import "golang.org/x/exp/constraints"

func IntAbs[T constraints.Integer](src T) T {
	if src < 0 {
		return -src
	} else {
		return src
	}
}
