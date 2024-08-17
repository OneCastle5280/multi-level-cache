package cache

import (
	"context"
	"time"
)

// Cache
// @Description:
type Cache interface {
	//
	// BatchSet
	//  @Description: 批量设置缓存
	//  @param ctx
	//  @param values
	//  @param expire
	//  @return error
	//
	BatchSet(ctx context.Context, values map[string][]byte, expire time.Duration) error

	//
	// BatchGet
	//  @Description: 批量查询缓存
	//  @param ctx
	//  @param keys
	//  @return map[string][]byte	缓存中的值
	//  @return []string			缓存中没有查询到的值
	//  @return error				查询过程中出现的异常
	//
	BatchGet(ctx context.Context, keys []string) (map[string][]byte, []string, error)

	//
	// BatchDel
	//  @Description: 批量删除缓存
	//  @param ctx
	//  @param keys
	//  @return error
	//
	BatchDel(ctx context.Context, keys []string) error
}
