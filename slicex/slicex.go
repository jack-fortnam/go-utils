package slicex

import "golang.org/x/exp/constraints"

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}

func Map[T, V any](slice []T, fn func(T) V) []V {
	result := make([]V, len(slice))
	for i, t := range slice {
		result[i] = fn(t)
	}
	return result
}

func Unique[T comparable](slice []T) []T {
	seen := make(map[T]struct{}, len(slice))
	result := make([]T, 0, len(slice))

	for _, v := range slice {
		if _, ok := seen[v]; !ok {
			seen[v] = struct{}{}
			result = append(result, v)
		}
	}
	return result
}

func Filter[T any](slice []T, predicate func(T) bool) []T {
	result := make([]T, 0, len(slice))
	for _, x := range slice {
		if predicate(x) {
			result = append(result, x)
		}
	}
	return result
}

func Sum[T constraints.Integer | constraints.Float](slice []T) T {
	return Reduce(slice, 0, func(acc T, x T) T {
		return acc + x
	})
}

func Reduce[T, V any](slice []T, initial V, fn func(V, T) V) V {
	acc := initial
	for _, x := range slice {
		acc = fn(acc, x)
	}
	return acc
}

func Flatten[T any](nested [][]T) []T {
	var totalLen int
	for _, sub := range nested {
		totalLen += len(sub)
	}

	result := make([]T, 0, totalLen)
	for _, sub := range nested {
		result = append(result, sub...)
	}
	return result
}

func Chunk[T any](items []T, size int) [][]T {
	if size <= 0 {
		return nil
	}
	chunks := make([][]T, 0, (len(items)+size-1)/size)

	for size < len(items) {
		items, chunks = items[size:], append(chunks, items[0:size:size])
	}

	return append(chunks, items)
}
