package models

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"zhu/config"
)

var db *gorm.DB
var rdb *redis.Client
var ctx = context.Background()

func Init() (err error) {

	err = NewMysql()
	if err != nil {

		return err
	}
	err = NewRedis()
	if err != nil {
		return err
	}
	return nil
}

func GetDB() *gorm.DB {
	return db
}

func NewMysql() (err error) {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	user := config.Conf.Value.GetString("database.user")
	pwd := config.Conf.Value.GetString("database.pwd")
	host := config.Conf.Value.GetString("database.host")
	port := config.Conf.Value.GetString("database.port")
	//dsn := "root:example@tcp(172.30.1.3:3306)/zhu?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/zhu?charset=utf8mb4&parseTime=True&loc=Local", user, pwd, host, port)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	return nil
}

func NewRedis() (err error) {
	host := config.Conf.Value.GetString("redis.host")
	port := config.Conf.Value.GetString("redis.port")
	rdb = redis.NewClient(&redis.Options{
		//Addr: "172.30.0.3:6379",
		Addr:     fmt.Sprintf("%s:%s", host, port),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err = rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	return nil
}
