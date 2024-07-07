package cache

type LocalCache[T any] struct {
	//
	//  commonCache[T]
	//  @Description: 公共处理逻辑
	//
	commonCache[T]

	//
	//  cache
	//  @Description: 本地缓存 cache
	//
	cache Cache[T]
}
