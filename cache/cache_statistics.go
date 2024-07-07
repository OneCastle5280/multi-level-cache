package cache

type (

	// StatsHandler
	// @Description: 缓存组件上报
	StatsHandler struct {
		//  统计开关
		disable bool
		//  统计上报组件
		handler *Statistics
	}

	// DefaultStatsHandler
	// @Description: 默认命中率统计组件
	DefaultStatsHandler struct {
	}

	// Statistics
	// @Description: 缓存命中率统计模块
	Statistics interface {
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
		// StatsQuery
		//  @Description: 统计缓存访问总次数
		//
		StatsQuery()
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
func NewStatsHandler(disable bool, handler Statistics) *StatsHandler {
	if disable {
		// 不开启命中率统计功能
		return nil
	}

	if handler == nil {
		handler = DefaultStatsHandler{}
	}
	return &StatsHandler{
		disable: disable,
		handler: &handler,
	}
}

func (DefaultStatsHandler) StatsHit() {
	//TODO implement me
	panic("implement me")
}

func (DefaultStatsHandler) StatsMiss() {
	//TODO implement me
	panic("implement me")
}

func (DefaultStatsHandler) StatsLocalHit() {
	//TODO implement me
	panic("implement me")
}

func (DefaultStatsHandler) StatsLocalMiss() {
	//TODO implement me
	panic("implement me")
}

func (DefaultStatsHandler) StatsRemoteHit() {
	//TODO implement me
	panic("implement me")
}

func (DefaultStatsHandler) StatsRemoteMiss() {
	//TODO implement me
	panic("implement me")
}

func (DefaultStatsHandler) StatsQuery() {
	//TODO implement me
	panic("implement me")
}

func (DefaultStatsHandler) StatsQueryFail(err error) {
	//TODO implement me
	panic("implement me")
}
