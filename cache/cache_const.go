package cache

const (
	defaultLocalExpire         = 3 * 60     // 默认 local 缓存过期时间 3 分钟
	defaultRemoteExpire        = 30 * 60    // 默认 remote 缓存过期时间 30 分钟
	defaultLocalCacheLimitSize = 512 * 1024 // 默认本地缓存大小 512 KB
)
