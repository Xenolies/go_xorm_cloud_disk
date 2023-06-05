package svc

import (
	"cloud_disk/core/internal/config"
	"cloud_disk/core/models"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/go-redis/redis/v8"
	"xorm.io/xorm"
)

type ServiceContext struct {
	Config config.Config
	Engine *xorm.Engine
	RDB    *redis.Client
	AliOSS *oss.Bucket
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Engine: models.Init(c.Mysql.Driver, c.Mysql.DataSource),
		RDB:    models.InitRDB(c.Redis.DataSource, c.Redis.Password, c.Redis.DB),
		AliOSS: models.InitAliOSS(c.AliOSS.Endpoint, c.AliOSS.AccessKey, c.AliOSS.SecretKey, c.AliOSS.BucketName),
	}
}
