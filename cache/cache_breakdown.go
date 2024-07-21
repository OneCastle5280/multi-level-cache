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

type breakDownKeysHandleFunc func(ctx context.Context, breakDownKeys []string) error

// NewDefaultCacheBreakDownHandler
//
//	@Description: 创建默认缓存穿透处理器
//	@return DefaultCacheBreakDownHandler
func NewDefaultCacheBreakDownHandler() DefaultCacheBreakDownHandler {
	return DefaultCacheBreakDownHandler{}
}

type DefaultCacheBreakDownHandler struct{}

func (DefaultCacheBreakDownHandler) IsBreakDownKeys(ctx context.Context, value []byte) bool {
	return bytes.Equal(value, defaultCacheBreakdownNode)
}

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
