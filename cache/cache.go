package cache

import (
	"context"
	"time"
)

// Cache
// @Description:
type Cache interface {
	Set(ctx context.Context, key string, value any, expire time.Duration) error         // 设置缓存
	BatchSet(ctx context.Context, values map[string][]byte, expire time.Duration) error // 批量设置缓存

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
	//
	// IsNotFoundErr
	//  @Description: 判断 err 是否为记录不存在
	//  @param ctx
	//  @param err
	//  @return bool
	//
	IsNotFoundErr(ctx context.Context, err error) bool
}
