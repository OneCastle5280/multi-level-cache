package cache

import "sync/atomic"

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
		StatsHit()
		//
		// StatsMiss
		//  @Description: 统计不命中次数
		//
		StatsMiss()
		//
		// StatsLocalHit
		//  @Description: 统计本地缓存命中次数
		//
		StatsLocalHit()
		//
		// StatsLocalMiss
		//  @Description: 统计本地缓存不命中次数
		//
		StatsLocalMiss()
		//
		// StatsRemoteHit
		//  @Description: 统计远程缓存命中次数
		//
		StatsRemoteHit()
		//
		// StatsRemoteMiss
		//  @Description: 统计远程缓存不命中次数
		//
		StatsRemoteMiss()
		//
		// StatsQueryTotal
		//  @Description: 统计缓存访问总次数
		//
		StatsQueryTotal()
		//
		// StatsQueryFail
		//  @Description: 统计查询失败的次数
		//  @param err
		//
		StatsQueryFail(err error)
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

func (d *DefaultStatsHandler) StatsHit() {
	atomic.AddUint64(&d.HitTotal, 1)
}

func (d *DefaultStatsHandler) StatsMiss() {
	atomic.AddUint64(&d.MissTotal, 1)
}

func (d *DefaultStatsHandler) StatsLocalHit() {
	atomic.AddUint64(&d.LocalHit, 1)
}

func (d *DefaultStatsHandler) StatsLocalMiss() {
	atomic.AddUint64(&d.LocalMiss, 1)
}

func (d *DefaultStatsHandler) StatsRemoteHit() {
	atomic.AddUint64(&d.RemoteHit, 1)
}

func (d *DefaultStatsHandler) StatsRemoteMiss() {
	atomic.AddUint64(&d.RemoteMiss, 1)
}

func (d *DefaultStatsHandler) StatsQueryTotal() {
	atomic.AddUint64(&d.QueryTotal, 1)
}

func (d *DefaultStatsHandler) StatsQueryFail(err error) {
	atomic.AddUint64(&d.QueryFail, 1)
}
