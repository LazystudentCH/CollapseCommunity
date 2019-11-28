package handlers

import (
	"encoding/json"
	"errors"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"server/models"
	"server/utils"
	"strconv"
)


var HotSubjectChan = make(chan *models.AddCommentRequest,100)
/*--------------业务函数-------------*/
func SubjectHandler(w http.ResponseWriter, r *http.Request) {
	mdb := utils.GetMysqlMasterConn()
	sdb := utils.GetMysqlSlaveConn()
	rc := utils.GetRedisConn()
	defer rc.Close()

	if r.Method == "GET" {
		//获取该页的帖子,如果没有获取到详情,则默认获取全部
		subject,ok := GetSubjectDetail(r,sdb)
		if ok {
			if subject != nil {
				//在redis中记录浏览量
				utils.RecordSubjectReview(subject.ID,rc)
				utils.ReturnData(w,models.Success,"",subject)
			}else{
				utils.ReturnData(w,models.Fail,"获取帖子详情失败",subject)
			}
			return
		}

		result, err := GetIndexSubject(r,sdb)
		if err != nil {
			utils.ReturnData(w, models.Fail, "获取帖子失败", nil)
			log.Info("获取帖子失败:", err)
			return
		}

		utils.ReturnData(w, models.Success, "", result)
		return

	} else {
		//POST请求为发帖子,需要登录
		userInfo, ok := utils.ParseToken(r)
		if !ok {
			//Token过期
			utils.ReturnData(w, models.TokenExpire, "", nil)
			return
		}

		subject := &models.Subject{}
		err := utils.PostMethodData(r,subject)
		if err != nil{
			utils.ReturnData(w,models.Error,"参数非法",nil)
			return
		}
		//判断是否已经存在,如果存在,则更新

		created,err := UpdateSubjectOrCreate(subject,mdb,userInfo)
		if err != nil {
			utils.ReturnData(w, models.Fail, "发帖失败", nil)
			log.Info("发帖失败:", err)
			return
		}

		if created {
			utils.IndexCount("SubjectCount")
		}

		utils.ReturnData(w, models.Success, "", nil)
		return
	}
}

func CommentHandler(w http.ResponseWriter, r *http.Request)  {
	mdb := utils.GetMysqlMasterConn()
	sdb := utils.GetMysqlSlaveConn()
	if r.Method == "GET"{
		result,err := GetSubjectAllComments(r,sdb)
		if err != nil {
			utils.ReturnData(w,models.Error,"获取评论失败",nil)
			log.Info("获取评论失败:",err)
			return
		}

		utils.ReturnData(w,models.Success,"",result)
		return
	}else{
		user,ok := utils.ParseToken(r)
		if !ok {
			//Token过期
			utils.ReturnData(w,models.TokenExpire,"",nil)
			return
		}

		req,err := AddSubjectComment(r,user.ID,mdb)
		if err != nil {
			utils.ReturnData(w,models.Error,"添加评论失败",nil)
			log.Info("添加评论失败:",err)
			return
		}

		utils.IndexCount("CommentCount")
		utils.ReturnData(w,models.Success,"",nil)
		HotSubjectChan <- req
		return
	}


}

func GetTotalPage(w http.ResponseWriter, r *http.Request) {
	sdb := utils.GetMysqlSlaveConn()
	subArr := make([]models.Subject, 0)
	total := 0

	sdb.Find(&subArr).Count(&total)
	utils.ReturnData(w, models.Success, "", total)
	return
}

func GetSearchTotalPage(w http.ResponseWriter, r *http.Request)  {
	sdb := utils.GetMysqlSlaveConn()
	keyword := r.URL.Query()["keyWord"]
	subArr := make([]models.Subject, 0)
	total := 0

	sdb.Find(&subArr).Where("title like ?",keyword[0]+"%").Count(&total)
	utils.ReturnData(w, models.Success, "", total)
	return
}

func GetCommentCount(w http.ResponseWriter, r *http.Request)  {
	rc := utils.GetRedisConn()
	defer rc.Close()

	sid , ok:= r.URL.Query()["sub_id"]
	if !ok || len(sid) == 0{
		utils.ReturnData(w,models.Error,"参数错误","")
		return
	}
	utils.ReturnData(w, models.Success, "", utils.GetSubjectComment(sid[0],rc))
	return
}

func GetSubjectReview(w http.ResponseWriter, r *http.Request)  {
	sid,ok:= r.URL.Query()["sub_id"]
	if !ok || len(sid) == 0{
		utils.ReturnData(w,models.Error,"参数错误","")
		return
	}
	utils.ReturnData(w,models.Success,"",utils.GetSubjectReview(sid[0]))
	return
}

func PriseSubject(w http.ResponseWriter,r *http.Request)  {
	rc := utils.GetRedisConn()
	defer rc.Close()

	user,ok := utils.ParseToken(r)
	if !ok {
		//Token过期
		utils.ReturnData(w,models.TokenExpire,"",nil)
		return
	}

	record := &models.PriseRecord{}
	err := utils.PostMethodData(r,record)
	if err != nil {
		utils.ReturnData(w,models.Error,err.Error(),nil)
		return
	}

	if record.HasPrise == false{
		utils.RecordSubjectPrise(user.ID,record.SubId,rc)
	}else{
		utils.CancelSubjectPrise(user.ID,record.SubId,rc)
	}

	utils.ReturnData(w,models.Success,"",nil)
	return
}

func QueryHasPrise(w http.ResponseWriter,r *http.Request)  {
	rc := utils.GetRedisConn()
	defer rc.Close()

	user,ok := utils.ParseToken(r)
	if !ok {
		//Token过期
		utils.ReturnData(w,models.TokenExpire,"",nil)
		return
	}

	record := &models.PriseRecord{}
	err := utils.PostMethodData(r,record)
	if err != nil {
		utils.ReturnData(w,models.Error,err.Error(),nil)
		return
	}

	exist := utils.QueryUserHasPrise(user.ID,record.SubId,rc)
	utils.ReturnData(w,models.Success,"",exist)
	return
}

func GetPriseCount(w http.ResponseWriter,r *http.Request)  {
	rc := utils.GetRedisConn()
	defer rc.Close()

	sid , ok:= r.URL.Query()["sub_id"]
	if !ok || len(sid) == 0{
		utils.ReturnData(w,models.Error,"参数错误","")
		return
	}
	utils.ReturnData(w, models.Success, "", utils.GetPriseCount(sid[0],rc))
	return
}

func SearchHandler(w http.ResponseWriter,r *http.Request)  {
	sdb := utils.GetMysqlSlaveConn()
	res, err := SearchSubject(r,sdb)
	if err != nil {
		utils.ReturnData(w,models.Error,err.Error(),nil)
		return
	}

	utils.ReturnData(w,models.Success,"",res)
	return

}

func DeleteSubject(w http.ResponseWriter,r *http.Request)  {
	_, ok := utils.ParseToken(r)
	if !ok {
		//Token过期
		utils.ReturnData(w, models.TokenExpire, "", nil)
		return
	}

	mdb := utils.GetMysqlMasterConn()
	err , sub_id := utils.GetMethodData(r,"sub_id")
	if err != nil {
		utils.ReturnData(w,models.Error,"参数有误",nil)
		return
	}

	err = mdb.Where("id = ?",sub_id[0]).Delete(models.Subject{}).Error
	if err != nil {
		utils.ReturnData(w,models.Error,"操作失败",nil)
		log.Info("删除失败:",err)
		return
	}

	utils.ReturnData(w,models.Success,"",nil)
	return
}

/*----------------逻辑函数------------------*/
//获取帖子所有的评论
func GetSubjectAllComments(r *http.Request,db *gorm.DB) ([]models.SubjectComment,error) {
	result := make([]models.SubjectComment, utils.Config.PageLimit)
	p, ok := r.URL.Query()["page"]
	sid,ok2 := r.URL.Query()["sub_id"]
	if !ok || !ok2{
		return nil, errors.New("缺少必要参数")
	}

	page, _ := strconv.Atoi(p[0])
	subId,_ := strconv.Atoi(sid[0])

	err := db.Limit(utils.Config.PageLimit).Offset(utils.Config.PageLimit * (page - 1)).Where("sub_id = ?",subId).Order("created_at desc").Find(&result).Error
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(result); i++ {
		user := &models.User{}
		db.Model(result[i]).Related(user)
		user.PassWord = ""
		result[i].User = *user
	}

	return result, nil
}

//搜索
func SearchSubject(r *http.Request,db *gorm.DB)  ([]models.Subject, error)  {
	p, ok := r.URL.Query()["page"]
	if !ok || len(p) == 0{
		return nil, errors.New("获取的数据不存在")
	}

	keyword , ok2 := r.URL.Query()["keyWord"]
	if !ok2 || len(keyword) == 0{
		return nil, errors.New("获取的数据不存在")
	}

	page, _ := strconv.Atoi(p[0])

	result := make([]models.Subject,0)
	err := db.Where("title like ?",keyword[0]+"%").Limit(utils.Config.PageLimit).Offset(utils.Config.PageLimit * (page - 1)).Order("created_at desc").Find(&result).Error
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(result); i++ {
		user := &models.User{}
		db.Model(result[i]).Related(&user)
		result[i].User = *user
	}

	return result, nil
}

//为帖子增加评论
func AddSubjectComment(r *http.Request,userId uint,db *gorm.DB) (*models.AddCommentRequest,error) {
	comment := &models.SubjectComment{}
	req := &models.AddCommentRequest{}
	data , err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil,err
	}

	err = json.Unmarshal(data, comment)
	if err != nil {
		return nil,err
	}

	comment.UserId = userId
	err = db.Create(comment).Error
	if err != nil {
		return nil,err
	}

	req.Sid = comment.SubId
	sub := models.Subject{}
	db.Where("id = ?",comment.SubId).Find(&sub)
	req.PubTime = sub.CreatedAt
	utils.RecordSubjectComment(comment.SubId)
	return req,nil
}

//获取帖子详情
func GetSubjectDetail(r *http.Request,db *gorm.DB) (*models.Subject,bool) {
	sid , ok := r.URL.Query()["sub_id"]
	if !ok {
		return nil,false
	}

	subId ,_ := strconv.Atoi(sid[0])
	sub, err := utils.GetHotSubject(uint(subId))
	if err == nil {
		return sub,true
	}

	subject := &models.Subject{}
	err = db.Where("id = ?",sid).First(&subject).Error
	if err != nil {
		log.Info("获取帖子详情失败")
		return nil,true
	}

	user := models.User{}
	db.Model(subject).Related(&user)
	subject.User = user
	user.PassWord = ""

	if subject.IsHot == true {
		utils.RecordHotSubject(subject)
	}
	return subject,true
}

//首页展示帖子
func GetIndexSubject(r *http.Request,db *gorm.DB) ([]models.Subject, error) {
	result := make([]models.Subject, utils.Config.PageLimit)

	p, ok := r.URL.Query()["page"]
	page, _ := strconv.Atoi(p[0])

	if !ok {
		return nil, errors.New("获取的数据不存在")
	}

	err := db.Limit(utils.Config.PageLimit).Offset(utils.Config.PageLimit * (page - 1)).Order("created_at desc").Find(&result).Error
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(result); i++ {
		user := &models.User{}
		db.Model(result[i]).Related(&user)
		result[i].User = *user
	}

	return result, nil
}

//发帖
func UpdateSubjectOrCreate(subject *models.Subject,db *gorm.DB, userInfo *models.User) (bool,error) {
	subject.UserId = userInfo.ID
	if subject.ID == 0 {
		//不存在
		err := db.Create(subject).Error
		if err != nil {
			return true,err
		}
		return true,nil
	}else{
		err := db.Model(subject).Where("id = ? ",subject.ID).Update(subject).Error
		if err != nil {
			return false,err
		}
		return false,nil
	}
}
