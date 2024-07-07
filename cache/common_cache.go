package cache

// commonCache
// @Description: 本地、远端缓存公共处理逻辑
type commonCache[T any] struct {
	//
	//  loader
	//  @Description: 回源接口
	//
	loader Loader[T]

	//
	//  statsHandler
	//  @Description: 统计组件
	//
	statsHandler StatsHandler
}
