package arrays

import "errors"

func Filter[T any](arr []T, fn func(x T) bool) []T {
	res := make([]T, 0)
	for _, x := range arr {
		if fn(x) {
			res = append(res, x)
		}
	}

	return res
}

func Map[T, Y any](arr []T, fn func(x T) Y) []Y {
	res := make([]Y, 0, len(arr))
	for _, x := range arr {
		res = append(res, fn(x))
	}

	return res
}

func FirstOrDefault[T any](arr []T, defaultValue T) T {
	if len(arr) > 0 {
		return arr[0]
	}

	return defaultValue
}

func Some[T any](arr []T, fn func(x T) bool) bool {
	for _, x := range arr {
		if fn(x) {
			return true
		}
	}

	return false
}

func Every[T any](arr []T, fn func(x T) bool) bool {
	for _, x := range arr {
		if !fn(x) {
			return false
		}
	}

	return true
}

func Find[T any](arr []T, fn func(x T) bool) (T, int, error) {
	for i, x := range arr {
		if fn(x) {
			return x, i, nil
		}
	}

	var val T
	return val, -1, errors.New("Element not found")
}

// Sort - Sort the array maybe using a lambda?
