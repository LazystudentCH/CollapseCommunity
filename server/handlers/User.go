package handlers

import (
	"encoding/json"
	"errors"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"server/models"
	"server/utils"
	"strconv"
)

func UserHandler(w http.ResponseWriter, r *http.Request)  {
	_,ok := utils.ParseToken(r)
	if !ok {
		//Token过期
		utils.ReturnData(w,models.TokenExpire,"",nil)
		return
	}
	if r.Method == "GET"{
		//GET方法为获取用户数据
		userInfo,err := GetUserInfo(r)
		if err != nil {
			utils.ReturnData(w,models.Fail,"获取用户信息失败",nil)
			return
		}

		utils.ReturnData(w,models.Success,"",userInfo)
		return
	}else{
		//POST方法为更新用户数据
		newToken,err := ModifyUserInfo(r)
		if err != nil {
			utils.ReturnData(w,models.Fail,"更新用户信息失败",nil)
			log.Info("更新用户信息失败:",err)
			return
		}
		utils.ReturnData(w,models.Success,"",newToken)
		return
	}

}

func UploadAvatar(w http.ResponseWriter, r *http.Request)  {
	_,ok := utils.ParseToken(r)
	if !ok {
		//Token过期
		utils.ReturnData(w,models.TokenExpire,"",nil)
		return
	}

	file,fileHead,err := r.FormFile("avatar")
	if err != nil {
		utils.ReturnData(w,models.Error,"数据非法",nil)
		return
	}

	defer file.Close()
	data := make([]byte,fileHead.Size)
	file.Read(data)
	err = utils.UploadToQiNiu(data,fileHead.Filename)
	if err != nil {
		utils.ReturnData(w,models.Fail,"上传头像失败",nil)
		return
	}

	utils.ReturnData(w,models.Success,"",utils.Config.QiNiuUrl+fileHead.Filename)
	return
}

func UserSubjectHandler(w http.ResponseWriter, r *http.Request)  {
	user,ok := utils.ParseToken(r)
	if !ok {
		//Token过期
		utils.ReturnData(w,models.TokenExpire,"",nil)
		return
	}

	result , err := GetAllUserSubjects(r,user)
	if err != nil {
		utils.ReturnData(w,models.Error,"获取数据失败",nil)
		log.Info("获取用户发帖记录失败:",err)
		return
	}

	utils.ReturnData(w,models.Success,"",result)
	return

}

func GetUserSubjectsCount(w http.ResponseWriter, r *http.Request)  {
	user,ok := utils.ParseToken(r)
	if !ok {
		//Token过期
		utils.ReturnData(w,models.TokenExpire,"",nil)
		return
	}

	sdb := utils.GetMysqlSlaveConn()
	subArr := make([]models.Subject, 0)
	total := 0

	sdb.Find(&subArr).Where("user_id = ?",user.ID).Count(&total)
	utils.ReturnData(w, models.Success, "", total)
	return
}

func ParseJwt(w http.ResponseWriter, r *http.Request)  {
	user,ok := utils.ParseToken(r)
	if !ok {
		utils.ReturnData(w,models.Error,"参数有误",nil)
		return
	}
	utils.ReturnData(w,models.Success,"",user)
	return
}

func UploadImage(w http.ResponseWriter, r *http.Request)  {
	_,ok := utils.ParseToken(r)
	if !ok {
		//Token过期
		utils.ReturnData(w,models.TokenExpire,"",nil)
		return
	}

	file,fileHead,err := r.FormFile("img")
	if err != nil {
		utils.ReturnData(w,models.Error,"数据非法",nil)
		return
	}

	defer file.Close()
	data := make([]byte,fileHead.Size)
	file.Read(data)

	err = utils.UploadToQiNiu(data,fileHead.Filename)
	if err != nil {
		utils.ReturnData(w,models.Fail,"上传图片失败",nil)
		log.Info("上传图片失败:",err)
		return
	}

	utils.ReturnData(w,models.Success,"",utils.Config.QiNiuUrl+fileHead.Filename)
	return
}

func GetAllUserSubjects(r *http.Request,user *models.User) ([]models.Subject, error) {
	result := make([]models.Subject, utils.Config.PageLimit)
	p, ok := r.URL.Query()["page"]
	if !ok {
		return nil, errors.New("获取的数据不存在")
	}
	page, _ := strconv.Atoi(p[0])

	sdb := utils.GetMysqlSlaveConn()
	err := sdb.Limit(utils.Config.PageLimit).Offset(utils.Config.PageLimit * (page - 1)).
		Where("user_id = ?",user.ID).Order("created_at desc").Find(&result).Error
	if err != nil {
		return nil, err
	}

	return result, nil
}

func GetUserInfo(r *http.Request) (*models.User,error) {
	phone , ok := r.URL.Query()["phone"]
	if len(phone) == 0 || !ok {
		return nil,nil
	}

	user := &models.User{}
	sdb := utils.GetMysqlSlaveConn()
	err := sdb.Where("phone = ?",phone).Find(&user).Error
	if err != nil {
		return nil,err
	}

	user.PassWord = ""
	return user,nil
}

func ModifyUserInfo(r *http.Request) (string,error) {
	user := &models.User{}
	data ,err := ioutil.ReadAll(r.Body)
	if err != nil {
		return "",err
	}

	err = json.Unmarshal(data, user)
	if err != nil {
		return "",err
	}

	mdb := utils.GetMysqlMasterConn()
	err = mdb.Model(user).Updates(user).Error
	if err != nil {
		return "",err
	}
	newToken,_ := utils.CreateToken(user)
	return newToken,nil
}
