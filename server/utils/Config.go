package utils

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
)

var Config *config

type config struct {
	HostAddr string `json:"host_addr"`

	MysqlMaster      string `json:"mysql_master"`
	MysqlSlave       string `json:"mysql_slave"`         //从节点配置
	MysqlMaxOpenConn int    `json:"mysql_max_open_conn"` //最大连接数
	MysqlMaxIdleConn int    `json:"mysql_max_idle_conn"` //最大空闲连接数

	Redis               string `json:"redis"`
	RedisMaxActiveConn  int    `json:"redis_max_active_conn"`
	RedisMaxIdleConn    int    `json:"redis_max_idle_conn"`
	RedisMaxIdleTimeout int    `json:"redis_max_idle_timeout"`

	SmsSignName     string `json:"sms_sign_name"` //发送短信模板签名
	TemplateCode    string `json:"template_code"` //短信模板
	AccessKeyId     string `json:"access_key_id"`
	AccessKeySecret string `json:"access_key_secret"`

	Issuer        string `json:"issuer"`
	JwtSecret     string `json:"jwt_secret"`
	JwtExpireTime int    `json:"jwt_expire_time"`

	QiNiuAk  string `json:"qi_niu_ak"`
	QiNiuSk  string `json:"qi_niu_sk"`
	QiNiuUrl string `json:"qi_niu_url"`

	PageLimit int `json:"page_limit"`
}

func LoadConfig() {
	cfg, err := ioutil.ReadFile("./config.json")
	if err != nil {
		log.Fatal("读取配置文件失败")
	}

	err = json.Unmarshal(cfg, &Config)
	if err != nil {
		log.Fatal("配置文件有误")
	}
	log.Info("读取配置文件成功")

	return
}
