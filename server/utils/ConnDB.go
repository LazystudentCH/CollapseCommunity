package utils

import (
	"github.com/garyburd/redigo/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"regexp"
	"time"
)

var redisClient *redis.Pool
var mysqlMasterConn *gorm.DB
var mysqlSlaveConn map[string]*gorm.DB

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
	//主节点
	_db, err := gorm.Open("mysql", Config.MysqlMaster)
	if err != nil {
		return err
	}
	mysqlMasterConn = _db
	mysqlMasterConn.DB().SetMaxOpenConns(Config.MysqlMaxOpenConn)   //设置数据库连接池最大连接数
	mysqlMasterConn.DB().SetMaxIdleConns(Config.MysqlMaxIdleConn)    //连接池最大允许的空闲连接数

	//从节点
	mysqlSlaveConn = make(map[string]*gorm.DB,0)
	for _ , v := range Config.MysqlSlave {
		_db2, err := gorm.Open("mysql", v)
		if err != nil {
			log.Warn("数据库连接失败:",v)
			continue
		}

		_db2.DB().SetMaxOpenConns(Config.MysqlMaxOpenConn)   //设置数据库连接池最大连接数
		_db2.DB().SetMaxIdleConns(Config.MysqlMaxIdleConn)    //连接池最大允许的空闲连接数
		host := getDbHost(v)
		mysqlSlaveConn[host] = _db2
	}

	log.Info("连接Mysql数据库成功")
	return nil
}

func getDbHost(connStr string) string {
	r,_ := regexp.Compile("\\(.*?\\)")
	res := r.FindString(connStr)
	return res[1:len(res)-1]
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
	db := &gorm.DB{}
	for _,v:= range mysqlSlaveConn {
		db = v
		break
	}

	return db
}

func GetRedisConn() redis.Conn {
	rc := redisClient.Get()
	return rc
}