package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"xorm.io/xorm"
)

/**
* @Author: Xenolies
* @Date: 2023/5/17 16:21
* @Version: 1.0
 */

var Engine = Init()

func Init() *xorm.Engine {
	engine, err := xorm.NewEngine("mysql", "root:123123@/cloud_disk?charset=utf8")
	if err != nil {
		fmt.Println("xorm.NewEngine ERROR")
		log.Printf("xorm.NewEngine Error: %v", err)
		return nil
	}
	fmt.Println("xorm.NewEngine Success")
	return engine
}
