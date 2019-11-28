package utils

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	log "github.com/sirupsen/logrus"
	"server/models"
	"strconv"
)

var countField = map[string]string{
	"MemberCount":"",
	"SubjectCount":"",
	"CommentCount":"",
}

var subComments = "subComments"
var subReview = "subReview"
var subPrise = "subPrise"
var userPrise = "userPrise"

func IndexCount(key string) {
	rc := GetRedisConn()
	defer rc.Close()
	_, err := rc.Do("incr", key)
	if err != nil {
		log.Info("计数失败:", err)
		return
	}
	return
}

func GetCount() map[string]string {
	rc := GetRedisConn()
	defer rc.Close()
	for key,_ := range countField {
		res, err := redis.String(rc.Do("get",key))
		if err != nil {
			log.Info("redis获取数据失败,key:",key)
			res = ""
		}
		countField[key] = res
	}
	return countField
}

func RecordHotSubject(sub *models.Subject)  {
	rc := GetRedisConn()
	defer rc.Close()
	data , _ := json.Marshal(sub)
	rc.Do("set",sub.ID,data)
	return
}

func RecordSubjectReview(sid uint, rc redis.Conn)  {
	_,err := rc.Do("hincrby",subReview,sid,1)
	if err != nil {
		log.Info("浏览量计数失败:",err)
		return
	}
	return
}

func RecordSubjectPrise(userId uint,sid uint,rc redis.Conn)  {

	uid := strconv.Itoa(int(userId))
	_,err := rc.Do("sadd",userPrise+":"+uid,sid)
	_,err2 := rc.Do("hincrby",subPrise,sid,1)
	if err != nil {
		log.Info("添加用户点赞失败",err)
		return
	}
	if err2 != nil {
		log.Info("统计点赞失败",err)
		return
	}

	return
}

func CancelSubjectPrise(userId uint,sid uint,rc redis.Conn)  {

	uid := strconv.Itoa(int(userId))
	_,err := rc.Do("srem",userPrise+":"+uid,sid)
	_,err2 := rc.Do("hincrby",subPrise,sid,-1)
	if err != nil {
		log.Info("减少用户点赞失败",err)
		return
	}
	if err2 != nil {
		fmt.Println(err2)
		log.Info("统计取消点赞失败",err)
		return
	}

	return
}

func RecordSubjectComment(sid uint)  {
	rc := GetRedisConn()
	defer rc.Close()
	_,err := rc.Do("hincrby",subComments,sid,1)
	if err != nil {
		log.Info("评论计数失败:",err)
		return
	}
	return
}

func QueryUserHasPrise(userId uint,sid uint,rc redis.Conn) bool {


	uid := strconv.Itoa(int(userId))
	res,err := rc.Do("sismember",userPrise+":"+uid,sid)
	if err != nil {
		log.Info("添加用户点赞失败",err)
		return false
	}

	return res == int64(1)
}

func GetPriseCount(sid string,rc redis.Conn) string {
	res,err := redis.String(rc.Do("hget",subPrise,sid))
	if err != nil {
		//log.Info("浏览量计数失败:",err)
		return "0"
	}
	return res
}

func GetSubjectComment(sid string,rc redis.Conn) string {
	res,err := redis.String(rc.Do("hget",subComments,sid))
	if err != nil {
		//log.Info("浏览量计数失败:",err)
		return "0"
	}
	return res
}

func GetSubjectReview(sid string) string {
	rc := GetRedisConn()
	defer rc.Close()
	res,err := redis.String(rc.Do("hget",subReview,sid))
	if err != nil {
		//log.Info("浏览量计数失败:",err)
		return "0"
	}
	return res
}

func GetHotSubject(sid uint) (*models.Subject,error) {
	rc := GetRedisConn()
	defer rc.Close()

	sub := &models.Subject{}
	res , err := redis.Bytes(rc.Do("get",sid))
	if  err != nil {
		return nil,err
	}
	json.Unmarshal(res,sub)
	return sub,nil
}