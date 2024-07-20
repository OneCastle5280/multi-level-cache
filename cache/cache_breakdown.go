package cache

import "context"

// CacheBreakDownHandler
//
//	@Description: 缓存击穿处理器
type CacheBreakDownHandler interface {
	//
	// IsBreakDownKeys
	//  @Description: 判断是否为击穿节点
	//  @param value
	//  @return bool
	//
	IsBreakDownKeys(value []byte) bool

	//
	// HandleBreakDownKeys
	//  @Description: 处理击穿节点
	//  @param ctx
	//  @param breakDownKeys
	//  @return error
	//
	HandleBreakDownKeys(ctx context.Context, breakDownKeys []string) error
}
