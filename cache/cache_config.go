package cache

type (

	// Config
	// @Description: 缓存自定义配置
	//
	Config struct {
		localExpire    int  // local  缓存过期时间, 单位：秒
		localLimitSize int  // local  缓存大小，默认为 512 KB
		remoteExpire   int  // remote 缓存过期时间, 单位：秒
		mode           Mode // 缓存模式
	}

	ConfigOption func(option *Config)
)

// NewCacheConfig 创建缓存配置
func NewCacheConfig(options ...ConfigOption) *Config {

	config := &Config{
		localExpire:    defaultLocalExpire,
		remoteExpire:   defaultRemoteExpire,
		localLimitSize: defaultLocalCacheLimitSize,
		mode:           MULTILEVEL, // 默认为多级缓存模式
	}

	// 自定义 config
	for _, option := range options {
		option(config)
	}

	return config
}
