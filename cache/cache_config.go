package cache

type (

	// Config
	// @Description: 缓存自定义配置
	Config struct {
		localExpire           int                       // local  缓存过期时间, 单位：秒
		localLimitSize        int                       // local  缓存大小，默认为 512 KB
		remoteExpire          int                       // remote 缓存过期时间, 单位：秒
		statsDisable          bool                      // 日志统计开关
		statsHandler          StatisticsHandler         // 自定义命中率统计
		breakDownHandler      CacheBreakDownHandler     // 自定义缓存穿透处理器
		serialization         Serialization             // 自定义序列化方式
		remoteCache           Cache                     // 自定义远程缓存
		localCache            Cache                     // 自定义本地缓存
		mode                  Mode                      // 缓存模式
		batchDeleteLocalCache BatchDeleteLocalCacheFunc // 自定义本地缓存清理
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
		statsHandler:   NewStatsHandler(false, nil),
	}

	// 自定义 config
	for _, option := range options {
		option(config)
	}

	return config
}

// getLocalExpire 获取 local 缓存超时时间
func (c *Config) getLocalExpire() int {
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

// GetStatsHandler 获取日志统计组件
func (c *Config) GetStatsHandler() StatisticsHandler {
	return c.statsHandler
}

// getRemoteCache 获取自定义远程缓存
func (c *Config) getRemoteCache() Cache {
	return c.remoteCache
}

// GetLocalCache 获取自定义本地缓存
func (c *Config) GetLocalCache() Cache {
	return c.localCache
}

// GetSerialization 获取序列化组件
func (c *Config) GetSerialization() Serialization {
	return c.serialization
}

// GetBatchDeleteLocalCache 获取自定义缓存清理func
func (c *Config) GetBatchDeleteLocalCache() BatchDeleteLocalCacheFunc {
	return c.batchDeleteLocalCache
}

func WithLocalExpire(expire int) ConfigOption {
	return func(option *Config) {
		option.localExpire = expire
	}
}

func WithLocalLimitSize(limitSize int) ConfigOption {
	return func(option *Config) {
		option.localLimitSize = limitSize
	}
}

func WithRemoteExpire(expire int) ConfigOption {
	return func(option *Config) {
		option.remoteExpire = expire
	}
}

func WithStatsDisable(disable bool) ConfigOption {
	return func(option *Config) {
		option.statsDisable = disable
	}
}

func WithStatsHandler(handler StatisticsHandler) ConfigOption {
	return func(option *Config) {
		option.statsHandler = handler
	}
}

func WithBreakDownHandler(handler CacheBreakDownHandler) ConfigOption {
	return func(option *Config) {
		option.breakDownHandler = handler
	}
}

func WithSerialization(serialization Serialization) ConfigOption {
	return func(option *Config) {
		option.serialization = serialization
	}
}

func WithRemoteCache(remoteCache Cache) ConfigOption {
	return func(option *Config) {
		option.remoteCache = remoteCache
	}
}

func WithLocalCache(localCache Cache) ConfigOption {
	return func(option *Config) {
		option.localCache = localCache
	}
}

func WithMode(mode Mode) ConfigOption {
	return func(option *Config) {
		option.mode = mode
	}
}

func WithBatchDeleteLocalCache(deleteFunc BatchDeleteLocalCacheFunc) ConfigOption {
	return func(option *Config) {
		option.batchDeleteLocalCache = deleteFunc
	}
}
