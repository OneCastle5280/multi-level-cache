package cache

type (
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
