package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"time"
)

func SendSms(phone string) (error,string) {
	client, err:= dysmsapi.NewClientWithAccessKey("cn-hangzhou", Config.AccessKeyId, Config.AccessKeySecret)

	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"

	request.PhoneNumbers = phone
	request.SignName = Config.SmsSignName
	request.TemplateCode = Config.TemplateCode

	verifyCode := getCode()
	code ,_:= json.Marshal(map[string]string{
		"code":verifyCode,
	})

	request.TemplateParam = string(code)
	response, err := client.SendSms(request)
	if err != nil || response.Code != "OK"{
		log.Info("发送短信失败,手机号码:",phone,"原因:",response.Message)
		return errors.New("response.Message"),""
	}

	log.Info("发送短信成功,手机号码:",phone)
	return nil,verifyCode
}

func getCode() string {
	return fmt.Sprintf("%04v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(9999))
}
