package services

import (
	"github.com/garyburd/redigo/redis"
	"github.com/jinzhu/gorm"
	"server/handlers"
	"server/models"
	"server/utils"
	"time"
)

func RunHotSubjectService()  {
	rc := utils.GetRedisConn()
	initSubjectList(rc)
	go getTopSubject()
	for {
		req := <- handlers.HotSubjectChan
		diff := time.Now().Sub(req.PubTime)
		if diff > time.Hour*72 {
			continue
		}
		addScore(req.Sid)
	}
}

func getTopSubject()  {
	t := time.NewTicker(time.Hour*24)
	for range t.C {
		mdb := utils.GetMysqlMasterConn()
		resetList(mdb)

		rc := utils.GetRedisConn()
		res ,_ := redis.StringMap(rc.Do("zrange","SubjectList","0","2","withscores"))
		for k := range res {
			mdb.Model(&models.Subject{}).Where("id = ?",k).Update("is_hot",true)
		}
		rc.Do("del","SubjectList")
		initSubjectList(rc)
		rc.Close()
	}
}

func resetList(db *gorm.DB)  {
	db.Model(&models.Subject{}).Where("is_hot = ?",true).Update("is_hot",false)
}

func addScore(sid uint)  {
	rc := utils.GetRedisConn()
	defer rc.Close()
	rc.Do("zincrby","SubjectList",1,sid)
}

func initSubjectList(rc redis.Conn)  {
	subjects := make([]models.Subject,0)
	sdb := utils.GetMysqlSlaveConn()
	sdb.Where("created_at between ? and ?",time.Now().Add(-1*time.Hour*72),time.Now()).Find(&subjects)
	for i := 0 ; i < len(subjects);i++{
		rc.Do("zincrby","SubjectList",0,subjects[i].ID)
	}
}

