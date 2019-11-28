package services

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"server/handlers"
	"server/utils"
)

func RunHttpService() {
	r := mux.NewRouter()

	r.HandleFunc("/login", handlers.Login).Methods("POST")                    //登录
	r.HandleFunc("/register", handlers.Register).Methods("POST")              //注册
	r.HandleFunc("/parsejwt", handlers.ParseJwt)                                        //解析token
	r.HandleFunc("/sms", handlers.SendSmsCode).Methods("POST")                //发送短信

	r.HandleFunc("/user", handlers.UserHandler).Methods("GET", "POST")        //用户相关
	r.HandleFunc("/user/subject", handlers.UserSubjectHandler).Methods("GET") //用户帖子相关
	r.HandleFunc("/user/getSubjectCount", handlers.GetUserSubjectsCount)                //获取用户发帖总数

	r.HandleFunc("/uploadAvatar", handlers.UploadAvatar) //上传相关
	r.HandleFunc("/uploadImage", handlers.UploadImage)

	r.HandleFunc("/subject", handlers.SubjectHandler).Methods("GET", "POST")    //帖子相关
	r.HandleFunc("/deleteSub", handlers.DeleteSubject)                                    //删除帖子
	r.HandleFunc("/indexTotalPage", handlers.GetTotalPage)                                //获取首页帖子总页数
	r.HandleFunc("/count", handlers.CountHandler)                                         //统计相关
	r.HandleFunc("/getSubjectReview", handlers.GetSubjectReview).Methods("GET") //获取帖子的浏览量

	r.HandleFunc("/search", handlers.SearchHandler).Methods("GET")               //搜索
	r.HandleFunc("/searchTotalPage", handlers.GetSearchTotalPage).Methods("GET") //获取搜索结果的最大页数

	r.HandleFunc("/prise", handlers.PriseSubject).Methods("POST")                //点赞
	r.HandleFunc("/queryHasPrise", handlers.QueryHasPrise).Methods("POST")       //查询是否已经点赞
	r.HandleFunc("/queryPriseCount", handlers.GetPriseCount).Methods("GET")      //查询帖子点赞数量

	r.HandleFunc("/comment", handlers.CommentHandler).Methods("GET", "POST")     //评论
	r.HandleFunc("/commentCount", handlers.GetCommentCount).Methods("GET")       //评论数量

	//启动http服务
	if err := http.ListenAndServe(utils.Config.HostAddr, r); err != nil {
		log.Fatalf("http服务启动失败:%s", err.Error())
		return
	}
}
