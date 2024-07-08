package cache

// CommonCache
// @Description: 本地、远端缓存公共处理逻辑
type CommonCache[T any] struct {
	//
	//  loader
	//  @Description: 回源接口
	//
	loader Loader

	//
	//  statsHandler
	//  @Description: 统计组件
	//
	statsHandler *StatsHandler
}

func NewCommonCache[T any](loader Loader, config *Config) CommonCache[T] {
	return CommonCache[T]{
		loader:       loader,
		statsHandler: NewStatsHandler(config.statsDisable, config.statsHandler),
	}
}
