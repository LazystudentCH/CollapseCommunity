package utils

import (
	"bytes"
	"context"
	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
)


func UploadToQiNiu(data []byte,filename string) error {
	putPolicy := storage.PutPolicy{
		Scope: "bengbengbeng",
	}
	mac := qbox.NewMac(Config.QiNiuAk, Config.QiNiuSk)
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuabei
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	dataLen := int64(len(data))
	putExtra := storage.PutExtra{
		Params: map[string]string{

		},
	}
	err := formUploader.Put(context.Background(), &ret, upToken, filename, bytes.NewReader(data), dataLen, &putExtra)
	if err != nil {
		return err
	}
	return nil
}
