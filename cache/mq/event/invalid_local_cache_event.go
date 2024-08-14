package event

type InvalidLocalCacheEvent struct {
	// cache union key
	CacheUnionKey string `json:"cacheUnionKey"`
	// local cache key
	CacheKey string `json:"cacheKey"`
}
