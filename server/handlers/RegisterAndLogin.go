package handlers

import (
	"crypto/md5"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"server/models"
	"server/utils"
)

//登录
func Login(w http.ResponseWriter, r *http.Request) {

	user := &models.User{}
	err := utils.PostMethodData(r,user)
	if err != nil {
		utils.ReturnData(w,models.Error,err.Error(),nil)
		return
	}

	sdb := utils.GetMysqlSlaveConn()
	userQuery := &models.User{}
	pwd := secretePwd(user.PassWord)
	err = sdb.Where("phone = ?",user.Phone).Find(&userQuery).Error
	if err != nil || userQuery.PassWord != pwd{
		utils.ReturnData(w, models.Fail, "用户名或者密码错误", nil)
		return
	}

	token , err := utils.CreateToken(userQuery)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		utils.ReturnData(w, models.Error, "服务器错误", nil)
	}
	utils.ReturnData(w,models.Success,"",token)
	return
}

//注册
func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		utils.ReturnData(w, models.Error, "不允许GET方法", nil)
		return
	}

	//解析参数
	user := &models.User{}
	err := utils.PostMethodData(r,user)
	if err != nil {
		utils.ReturnData(w,models.Error,err.Error(),nil)
		return
	}

	sdb := utils.GetMysqlSlaveConn()
	//验证数据
	if ok, msg := validateData(user, sdb, user.VerifyCode); !ok {
		utils.ReturnData(w, models.Fail, msg, nil)
		return
	}

	mdb := utils.GetMysqlMasterConn()
	//写入数据库
	if mdb.Create(user).Error != nil {
		utils.ReturnData(w, models.Fail, "注册失败", nil)
		return
	}

	//redis中记录总数
	utils.IndexCount("MemberCount")
	utils.ReturnData(w, models.Success, "", nil)
	return
}

//发送短信验证码
func SendSmsCode(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		utils.ReturnData(w, models.Error, "不允许GET方法", nil)
		return
	}

	data, err := ioutil.ReadAll(r.Body)
	phone := string(data)
	if err != nil {
		utils.ReturnData(w, models.Error, "数据非法", nil)
		log.Info("获取电话号码失败:", err)
		return
	}

	err, verifyCode := utils.SendSms(phone)
	if err != nil {
		utils.ReturnData(w, models.Error, "发送验证码失败,请稍后再试", nil)
		log.Info("发送短信失败:", err)
		return
	}

	//保存验证码至redis
	saveVerifyCodeToRedis(phone, verifyCode)
	utils.ReturnData(w, models.Success, "发送验证码成功", nil)
	return
}








//保存验证码至redis,过期时间为十分钟
func saveVerifyCodeToRedis(phone string, code string) {
	rc := utils.GetRedisConn()
	defer rc.Close()
	rc.Do("set", phone, code, "EX", 600)
	return
}

//从redis中获取验证码
func getVerifyCode(phone string, code string) bool {
	rc := utils.GetRedisConn()
	defer rc.Close()
	res, err := redis.String(rc.Do("get", phone))
	if err != nil || res != code {
		return false
	}
	return true
}

//加密密码
func secretePwd(pwd string) string {
	data := []byte(pwd)
	has := md5.Sum(data)
	return fmt.Sprintf("%x", has)
}

//验证数据
func validateData(user *models.User, db *gorm.DB, verifyCode string) (bool, string) {
	if hasRegister(user.Phone, user, db) == true {
		return false, "手机号码已经被注册"
	}

	if getVerifyCode(user.Phone, verifyCode) == false {
		return false, "验证码已经过期或者输入错误"
	}
	user.PassWord = secretePwd(user.PassWord)
	return true, ""
}

//判断手机号码是否已经注册
func hasRegister(phone string, user *models.User, db *gorm.DB) bool {
	if db.Where("phone = ?", phone).Find(user).Error == nil {
		//找到记录
		return true
	}
	return false
}

