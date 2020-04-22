package option

import (
	"github.com/go-redis/redis"
)

const (
	LevelDisable = 0 //禁止
	LevelModel   = 1 //只缓存模型
	LevelSearch  = 2 //查询缓存
	MinExpires   = 30
	MaxExpires   = 3600
)

type Option struct {
	Expires         int  //默认120秒，30-900
	Level           int  //默认LevelSearch，LevelDisable:关闭，LevelModel:模型缓存， LevelSearch:查询缓存
	AsyncWrite      bool //默认false， insert update delete 成功后是否异步更新缓存
	PenetrationSafe bool //默认false, 开启防穿透。
	RedisClient     *redis.Client
}

func (o *Option) Init() {
	if o.Level == 0 {
		o.Level = LevelSearch
	}
	if o.Expires == 0 {
		o.Expires = 120
	}

	if o.Expires < MinExpires || o.Expires > MaxExpires {
		o.Expires = 300
	}
}
