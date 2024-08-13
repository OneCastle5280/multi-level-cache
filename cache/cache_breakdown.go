package cache

import (
	"bytes"
	"context"
)

// 默认缓存击穿节点
var defaultCacheBreakdownNode = []byte("default break down")

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
	IsBreakDownKeys(ctx context.Context, value []byte) bool

	//
	// HandleBreakDownKeys
	//  @Description: 处理击穿节点
	//  @param ctx
	//  @param breakDownKeys
	//  @return map[string][]byte
	//
	HandleBreakDownKeys(ctx context.Context, breakDownKeys []string) map[string][]byte
}

// NewDefaultCacheBreakDownHandler
//
//	@Description: 创建默认缓存穿透处理器
//	@return DefaultCacheBreakDownHandler
func NewDefaultCacheBreakDownHandler() DefaultCacheBreakDownHandler {
	return DefaultCacheBreakDownHandler{}
}

type DefaultCacheBreakDownHandler struct{}

// IsBreakDownKeys
//
//	@Description:  判断是否为缓存穿透 keys
//	@receiver DefaultCacheBreakDownHandler
//	@param ctx
//	@param value
//	@return bool
func (DefaultCacheBreakDownHandler) IsBreakDownKeys(ctx context.Context, value []byte) bool {
	return bytes.Equal(value, defaultCacheBreakdownNode)
}

// HandleBreakDownKeys
//
//	@Description: 封装处理穿透 key
//	@receiver DefaultCacheBreakDownHandler
//	@param ctx
//	@param breakDownKeys
//	@return map[string][]byte
func (DefaultCacheBreakDownHandler) HandleBreakDownKeys(ctx context.Context, breakDownKeys []string) map[string][]byte {
	result := make(map[string][]byte, len(breakDownKeys))
	if len(breakDownKeys) == 0 {
		return result
	}
	for _, key := range breakDownKeys {
		result[key] = defaultCacheBreakdownNode
	}
	return result
}
