package cache

type (

	// Config
	// @Description: 缓存自定义配置
	//
	Config struct {
		localExpire      int                   // local  缓存过期时间, 单位：秒
		localLimitSize   int                   // local  缓存大小，默认为 512 KB
		remoteExpire     int                   // remote 缓存过期时间, 单位：秒
		statsDisable     bool                  // 日志统计开关
		statsHandler     Statistics            // 自定义命中率统计
		breakDownHandler CacheBreakDownHandler // 自定义缓存穿透处理器
		serialization    Serialization         // 自定义序列化方式
		remoteCache      Cache                 // 自定义远程缓存
		localCache       Cache                 // 自定义本地缓存
		mode             Mode                  // 缓存模式
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

// GetLocalExpire 获取 local 缓存超时时间
func (c *Config) GetLocalExpire() int {
	if c.localExpire <= 0 {
		return defaultLocalExpire
	}
	return c.localExpire
}

// GetRemoteExpire 获取 remote 缓存超时时间
func (c *Config) GetRemoteExpire() int {
	if c.remoteExpire <= 0 {
		return defaultRemoteExpire
	}
	return c.remoteExpire
}

// GetLocalLimitSize 获取本地缓存大小，默认 512 KB
func (c *Config) GetLocalLimitSize() int {
	if c.localLimitSize < defaultLocalCacheLimitSize {
		return defaultLocalCacheLimitSize
	}
	return c.localLimitSize
}

// GetMode 获取缓存模式
func (c *Config) GetMode() Mode {
	return c.mode
}

// GetStatsDisable 获取日志统计功能开关
func (c *Config) GetStatsDisable() bool {
	return c.statsDisable
}

// getRemoteCache 获取自定义远程缓存
func (c *Config) getRemoteCache() Cache {
	return c.remoteCache
}

// getLocalCache 获取自定义本地缓存
func (c *Config) getLocalCache() Cache {
	return c.localCache
}

// GetSerialization 获取序列化组件
func (c *Config) GetSerialization() Serialization {
	return c.serialization
}
