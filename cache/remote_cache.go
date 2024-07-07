package cache

type (
	RemoteCache[T any] struct {
		//
		//  CommonCache[T]
		//  @Description: 公共处理逻辑
		//
		CommonCache[T]

		//
		//  cache
		//  @Description: 远程缓存 cache
		//
		cache Cache
	}
)

// NewRemoteCache
//
//	@Description: 创建远程缓存
//	@param c
//	@param loader
//	@param config
//	@return *RemoteCache[T]
func NewRemoteCache[T any](config *Config) *RemoteCache[T] {
	return &RemoteCache[T]{
		cache:       config.GetRemoteCache(),
		CommonCache: NewCommonCache[T](config),
	}
}
