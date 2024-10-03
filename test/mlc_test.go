package test

import (
	"context"
	"mlc/cache"
	"mlc/mlc"
	"testing"
)

const customCacheBreakDownValue = "custom cache break down handle"

type customCacheBreakDownHandler struct{}

func (customCacheBreakDownHandler) IsBreakDownKeys(ctx context.Context, value []byte) bool {
	if string(value) != customCacheBreakDownValue {
		return false
	}
	return true
}

func (customCacheBreakDownHandler) HandleBreakDownKeys(ctx context.Context, breakDownKeys []string) map[string][]byte {
	result := make(map[string][]byte, len(breakDownKeys))
	if len(breakDownKeys) == 0 {
		return result
	}
	for _, key := range breakDownKeys {
		result[key] = []byte(customCacheBreakDownValue)
	}
	return result
}

func TestNewDefaultMultiLevelCache(t *testing.T) {
	t.Run("test local cache", func(t *testing.T) {
		mlcTestLocal := mlc.NewDefaultMultiLevelCache[string](func(ctx context.Context, keys []string) (map[string][]byte, error) {
			result := make(map[string][]byte, len(keys))
			for _, key := range keys {
				value := "\"" + key + " value\""
				result[key] = []byte(value)
			}
			return result, nil
		}, "mlcTestLocal1", cache.WithMode(cache.LOCAL))

		ctx := context.Background()
		keys := []string{
			"mlcTestLocal-key-1",
			"mlcTestLocal-key-2",
			"mlcTestLocal-key-3",
		}
		values, err := mlcTestLocal.BatchGet(ctx, keys)
		if err != nil {
			t.Error("mlcTestLocal.BatchGet error")
		}
		for _, key := range keys {
			expect := key + " value"
			actual := values[key]
			if actual == nil || *actual != expect {
				t.Error("mlcTestLocal.BatchGet value err")
			}
		}
	})

	t.Run("test local cache breakDown", func(t *testing.T) {
		mlcTestLocal := mlc.NewDefaultMultiLevelCache[string](func(ctx context.Context, keys []string) (map[string][]byte, error) {
			result := make(map[string][]byte, len(keys))
			for _, key := range keys {
				if key == "mlcTestLocal-key-1" {
					continue
				}
				value := "\"" + key + " value\""
				result[key] = []byte(value)
			}
			return result, nil
		}, "mlcTestLocal2", cache.WithMode(cache.LOCAL))
		ctx := context.Background()
		keys := []string{
			"mlcTestLocal-key-1",
			"mlcTestLocal-key-2",
			"mlcTestLocal-key-3",
		}
		_, err := mlcTestLocal.BatchGet(ctx, keys)
		if err != nil {
			t.Error("mlcTestLocal.BatchGet error")
		}
		breakDownValues, err := mlcTestLocal.BatchGet(ctx, []string{keys[0]})
		if err != nil {
			t.Error("mlcTestLocal.BatchGet error")
		}
		// break down is not need return
		if len(breakDownValues) > 0 {
			t.Error("handle break down error")
		}
	})

	t.Run("test custom cacheBreakDownHandler", func(t *testing.T) {
		mlcTestLocal := mlc.NewDefaultMultiLevelCache[string](func(ctx context.Context, keys []string) (map[string][]byte, error) {
			result := make(map[string][]byte, len(keys))
			for _, key := range keys {
				if key == "mlcTestLocal-key-1" {
					continue
				}
				value := "\"" + key + " value\""
				result[key] = []byte(value)
			}
			return result, nil
		}, "mlcTestLocal2", cache.WithMode(cache.LOCAL), cache.WithBreakDownHandler(&customCacheBreakDownHandler{}))
		ctx := context.Background()
		keys := []string{
			"mlcTestLocal-key-1",
			"mlcTestLocal-key-2",
			"mlcTestLocal-key-3",
		}
		_, err := mlcTestLocal.BatchGet(ctx, keys)
		if err != nil {
			t.Error("mlcTestLocal.BatchGet error")
		}
		breakDownValues, err := mlcTestLocal.BatchGet(ctx, []string{keys[0]})
		if err != nil {
			t.Error("mlcTestLocal.BatchGet error")
		}
		// break down is not need return
		if len(breakDownValues) > 0 {
			t.Error("handle break down error")
		}
	})
}
