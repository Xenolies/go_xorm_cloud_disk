package models

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"xorm.io/xorm"
)

/**
* @Author: Xenolies
* @Date: 2023/5/17 16:21
* @Version: 1.0
 */

// Init 初始化 Xorm
func Init(driver string, dataSource string) *xorm.Engine {
	engine, err := xorm.NewEngine(driver, dataSource)
	if err != nil {
		fmt.Println("xorm.NewEngine ERROR")
		log.Printf("xorm.NewEngine Error: %v", err)
		return nil
	}
	fmt.Println("xorm.NewEngine Success")
	return engine
}

// InitRDB 初始化 Redis
func InitRDB(Addr string, Password string, DB int) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     Addr,
		Password: Password,
		DB:       DB,
	})
}

// InitAliOSS 初始化阿里云存储
func InitAliOSS(Endpoint string, AccessKey string, SecretKey string, BucketName string) *oss.Bucket {
	client, err := oss.New(Endpoint, AccessKey, SecretKey)
	if err != nil {
		log.Println(err)
		return nil
	}
	bucket, err := client.Bucket(BucketName)
	if err != nil {
		log.Println(err)
		return nil
	}
	return bucket
}
