package data

import (
	"github.com/dtm-labs/rockscache"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
	"world/internal/conf"
)

// ProviderSet is data providers.
// Notice data层需要初始化的Mysql、Redis等资源直接放在这里即可，wire的时候会生成初始化代码
var ProviderSet = wire.NewSet(NewData, NewMysql, NewRedis, NewRocksCache, NewCountryRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	Mysql      *gorm.DB
	RedisCli   *redis.Client
	RocksCache *rockscache.Client // 缓存一致性的库
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, mysqlDB *gorm.DB, redisCli *redis.Client, rocksCli *rockscache.Client) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	// Notice 初始化Data后封装了Mysql与Redis客户端
	return &Data{
		Mysql:      mysqlDB,
		RedisCli:   redisCli,
		RocksCache: rocksCli,
	}, cleanup, nil
}

func NewMysql(c *conf.Data, logger log.Logger) *gorm.DB {
	// recover
	defer func() {
		if err := recover(); err != nil {
			log.NewHelper(logger).Errorw("kind", "mysql", "error", err)
		}
	}()
	// Notice pkg
	// mysql数据库连接
	db, err := gorm.Open(mysql.Open(c.Database.Source))
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	// 连接池参数设置
	sqlDB.SetMaxIdleConns(int(c.Database.MinIdleConns))
	sqlDB.SetMaxOpenConns(int(c.Database.MaxOpenConns))
	sqlDB.SetConnMaxLifetime(time.Hour * time.Duration(c.Database.ConMaxLeftTime))

	return db
}

func NewRedis(c *conf.Data, logger log.Logger) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: c.Redis.Addr,
		//Password:     c.Redis.Password,
		DB:           int(c.Redis.Db),
		PoolSize:     int(c.Redis.PoolSize),     // 连接池数量
		MinIdleConns: int(c.Redis.MinIdleConns), //好比最小连接数
		MaxRetries:   int(c.Redis.MaxRetries),   // 命令执行失败时，最多重试多少次，默认为0即不重试
	})
	//rdb.AddHook(redisotel.NewTracingHook(
	//	redisotel.WithAttributes(
	//		attribute.String("db.type", "redis"),
	//		attribute.String("db.ip", c.Redis.Addr),
	//		attribute.String("db.instance", c.Redis.Addr+"/"+convertor.ToString(c.Redis.Db)),
	//	),
	//))
	log.NewHelper(logger).Infow("kind", "redis", "status", "enable")
	return rdb
}

func NewRocksCache(c *conf.Data, logger log.Logger, rdb *redis.Client) *rockscache.Client {
	var dc = rockscache.NewClient(rdb, rockscache.NewDefaultOptions())
	log.NewHelper(logger).Infow("kind", "rocksCache", "status", "enable")
	return dc
}
