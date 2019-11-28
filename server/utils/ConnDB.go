package utils

import (
	"github.com/garyburd/redigo/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"time"
)

var redisClient *redis.Pool
var mysqlMasterConn *gorm.DB
var mysqlSlaveConn *gorm.DB

func ConnDb()  {
	err := connMysql()
	if err != nil {
		log.Fatal("连接Mysql数据库失败, error=" + err.Error())
		return
	}

	err = connRedis()
	if err != nil {
		log.Fatal("连接Redis数据库失败, error=" + err.Error())
		return
	}
	return
}

func connMysql() error {
	_db, err := gorm.Open("mysql", Config.MysqlMaster)
	_db2, err := gorm.Open("mysql", Config.MysqlSlave)
	if err != nil {

		return err
	}


	mysqlMasterConn = _db
	mysqlMasterConn.DB().SetMaxOpenConns(Config.MysqlMaxOpenConn)   //设置数据库连接池最大连接数
	mysqlMasterConn.DB().SetMaxIdleConns(Config.MysqlMaxIdleConn)    //连接池最大允许的空闲连接数

	mysqlSlaveConn = _db2
	mysqlSlaveConn.DB().SetMaxOpenConns(Config.MysqlMaxOpenConn)   //设置数据库连接池最大连接数
	mysqlSlaveConn.DB().SetMaxIdleConns(Config.MysqlMaxIdleConn)    //连接池最大允许的空闲连接数

	log.Info("连接Mysql数据库成功")
	return nil
}

func connRedis() error {
	redisClient = &redis.Pool{
		MaxIdle:     Config.RedisMaxIdleConn,
		MaxActive:   Config.RedisMaxActiveConn,
		IdleTimeout:  time.Duration(Config.RedisMaxIdleTimeout) * time.Second,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			con, err := redis.Dial("tcp", Config.Redis,
				redis.DialConnectTimeout(time.Duration(Config.RedisMaxIdleTimeout)*time.Second),
				redis.DialReadTimeout(time.Duration(Config.RedisMaxIdleTimeout)*time.Second),
				redis.DialWriteTimeout(time.Duration(Config.RedisMaxIdleTimeout)*time.Second))
			if err != nil {
				return nil, err
			}
			return con, nil
		},
	}

	rc := redisClient.Get()
	defer rc.Close()
	_ , err := rc.Do("PING")
	if err != nil {
		return err
	}
	log.Info("连接redis数据库成功")
	return nil
}

func GetMysqlMasterConn() *gorm.DB {
	return mysqlMasterConn
}

func GetMysqlSlaveConn() *gorm.DB {
	return mysqlSlaveConn
}

func GetRedisConn() redis.Conn {
	rc := redisClient.Get()
	return rc
}