package util

// Keys
//
//	@Description: 获取 source 的 all Keys
//	@param source
//	@return []K
func Keys[K comparable, V any](source map[K]V) []K {
	return Flatten(source, func(target *[]K) {
		for k, _ := range source {
			*target = append(*target, k)
		}
	})
}

// Flatten
//
//	@Description: 将 source 展开
//	@param source
//	@param apply
func Flatten[K comparable, V any, E any, S ~[]E](source map[K]V, apply func(target *S)) S {
	if source == nil {
		return nil
	}
	if len(source) == 0 {
		return S{}
	}
	result := make(S, 0, len(source))
	apply(&result)
	return result
}
