package models

import (
	"fmt"
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

// var Engine = Init()
//
// var RDB = InitRDB()

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

func InitRDB(Addr string, Password string, DB int) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     Addr,
		Password: Password,
		DB:       DB,
	})
}
