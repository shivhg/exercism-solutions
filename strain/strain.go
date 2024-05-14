package strain

// Implement the "Keep" and "Discard" function in this file.

// You will need typed parameters (aka "Generics") to solve this exercise.
// They are not part of the Exercism syllabus yet but you can learn about
// them here: https://go.dev/tour/generics/1

func Keep[T any](cp []T, filterFunc func(T) bool) []T {
	var result []T
	for _, val := range cp {
		if filterFunc(val) {
			result = append(result, val)
		}
	}
	return result
}

func Discard[T any](arr []T, filterFunc func(T) bool) []T {
	var result []T
	for _, val := range arr {
		if !filterFunc(val) {
			result = append(result, val)
		}
	}
	return result
}
