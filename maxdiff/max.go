package maxdiff

import "fmt"

// MaxDiff returns the biggest difference for any pair of numbers for any given array
func MaxDiff(a []int32) int32 {

	var bigDiff int32

	fmt.Printf("array length is: %d\n", len(a))

	for i := 1; i < len(a); i++ {
		// fmt.Printf("i - index %d and value %d\n", i, a[i])
		for j := 0; j < i; j++ {
			// fmt.Printf("-------j - index %d, value %d, difference %d\n", j, a[j], a[i]-a[j])
			if a[i]-a[j] > bigDiff {
				bigDiff = a[i] - a[j]
			}
		}
	}

	return bigDiff
}
