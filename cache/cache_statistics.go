package cache

import (
	"context"
	"sync/atomic"
)

type (

	// StatsHandler
	// @Description: 缓存组件上报
	StatsHandler struct {
		//  统计开关
		disable bool
		//  统计上报组件
		handler *StatisticsHandler
	}

	// DefaultStatsHandler
	// @Description: 默认命中率统计组件
	DefaultStatsHandler struct {
		Name       string
		HitTotal   uint64
		MissTotal  uint64
		LocalHit   uint64
		LocalMiss  uint64
		RemoteHit  uint64
		RemoteMiss uint64
		QueryTotal uint64
		QueryFail  uint64
	}

	// StatisticsHandler
	// @Description: 缓存命中率统计组件
	StatisticsHandler interface {
		//
		// StatsHit
		//  @Description: 统计命中次数
		//
		StatsHit(ctx context.Context, count int64)
		//
		// StatsMiss
		//  @Description: 统计不命中次数
		//
		StatsMiss(ctx context.Context, count int64)
		//
		// StatsLocalHit
		//  @Description: 统计本地缓存命中次数
		//
		StatsLocalHit(ctx context.Context, count int64)
		//
		// StatsLocalMiss
		//  @Description: 统计本地缓存不命中次数
		//
		StatsLocalMiss(ctx context.Context, count int64)
		//
		// StatsRemoteHit
		//  @Description: 统计远程缓存命中次数
		//
		StatsRemoteHit(ctx context.Context, count int64)
		//
		// StatsRemoteMiss
		//  @Description: 统计远程缓存不命中次数
		//
		StatsRemoteMiss(ctx context.Context, count int64)
		//
		// StatsQueryTotal
		//  @Description: 统计缓存访问总次数
		//
		StatsQueryTotal(ctx context.Context, count int64)
		//
		// StatsQueryFail
		//  @Description: 统计查询失败的次数
		//  @param err
		//
		StatsQueryFail(ctx context.Context, count int64, err error)
	}
)

// NewStatsHandler
//
//	@Description: 创建缓存命中率统计组件
//	@param disable
//	@param handler
//	@return *StatsHandler
func NewStatsHandler(disable bool, handler StatisticsHandler) *StatsHandler {
	if disable {
		// 不开启命中率统计功能
		return nil
	}

	if handler == nil {
		handler = &DefaultStatsHandler{}
	}
	return &StatsHandler{
		disable: disable,
		handler: &handler,
	}
}

func (d *DefaultStatsHandler) StatsHit(ctx context.Context, count int64) {
	atomic.AddUint64(&d.HitTotal, uint64(count))
}

func (d *DefaultStatsHandler) StatsMiss(ctx context.Context, count int64) {
	atomic.AddUint64(&d.MissTotal, uint64(count))
}

func (d *DefaultStatsHandler) StatsLocalHit(ctx context.Context, count int64) {
	atomic.AddUint64(&d.LocalHit, uint64(count))
}

func (d *DefaultStatsHandler) StatsLocalMiss(ctx context.Context, count int64) {
	atomic.AddUint64(&d.LocalMiss, uint64(count))
}

func (d *DefaultStatsHandler) StatsRemoteHit(ctx context.Context, count int64) {
	atomic.AddUint64(&d.RemoteHit, uint64(count))
}

func (d *DefaultStatsHandler) StatsRemoteMiss(ctx context.Context, count int64) {
	atomic.AddUint64(&d.RemoteMiss, uint64(count))
}

func (d *DefaultStatsHandler) StatsQueryTotal(ctx context.Context, count int64) {
	atomic.AddUint64(&d.QueryTotal, 1)
}

func (d *DefaultStatsHandler) StatsQueryFail(ctx context.Context, count int64, err error) {
	atomic.AddUint64(&d.QueryFail, 1)
}

func (s *StatsHandler) StatsHit(ctx context.Context, count int64) {
	if s.disable {
		return
	}
	handler := *s.handler
	handler.StatsHit(nil, count)
}

func (s *StatsHandler) StatsMiss(ctx context.Context, count int64) {
	if s.disable {
		return
	}
	handler := *s.handler
	handler.StatsMiss(nil, count)
}

func (s *StatsHandler) StatsLocalHit(ctx context.Context, count int64) {
	if s.disable {
		return
	}
	handler := *s.handler
	handler.StatsLocalHit(nil, count)
}

func (s *StatsHandler) StatsLocalMiss(ctx context.Context, count int64) {
	if s.disable {
		return
	}
	handler := *s.handler
	handler.StatsLocalMiss(nil, count)
}

func (s *StatsHandler) StatsRemoteHit(ctx context.Context, count int64) {
	if s.disable {
		return
	}
	handler := *s.handler
	handler.StatsRemoteHit(nil, count)
}

func (s *StatsHandler) StatsRemoteMiss(ctx context.Context, count int64) {
	if s.disable {
		return
	}
	handler := *s.handler
	handler.StatsRemoteMiss(nil, count)
}

func (s *StatsHandler) StatsQueryTotal(ctx context.Context, count int64) {
	if s.disable {
		return
	}
	handler := *s.handler
	handler.StatsQueryTotal(ctx, count)
}

func (s *StatsHandler) StatsQueryFail(ctx context.Context, count int64, err error) {
	if s.disable {
		return
	}
	handler := *s.handler
	handler.StatsQueryFail(ctx, count, err)
}
