package cache

import (
	"context"
	"sync"
)

type Initializer struct {
	sync.Mutex
	Initializer func(ctx context.Context) // 初始化组件
	Initialized bool                      // 是否已经初始化完成
}
