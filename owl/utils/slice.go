package utils

type unique interface {
	int | int64 | uint | uint64 | string
}

// 去除 slice 中重复的元素

func Unique[T unique](t []T) []T {
	seen := make(map[T]bool)
	var result []T

	for _, str := range t {
		if _, exists := seen[str]; !exists {
			seen[str] = true
			result = append(result, str)
		}
	}

	return result
}
