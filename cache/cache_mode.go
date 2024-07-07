package cache

const (
	LOCAL      Mode = "LOCAL"      // 本地缓存
	REMOTE     Mode = "REMOTE"     // 远程缓存
	MULTILEVEL Mode = "MULTILEVEL" // 多级
)

// Mode 缓存模式
type Mode string
