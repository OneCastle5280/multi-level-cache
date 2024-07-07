package cache

type RemoteCache[T any] struct {
	//
	//  commonCache[T]
	//  @Description: 公共处理逻辑
	//
	commonCache[T]

	//
	//  cache
	//  @Description: 远程缓存 cache
	//
	cache Cache[T]
}
