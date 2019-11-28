package utils

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

func PostMethodData(r *http.Request,v interface{}) (error) {
	data , err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, v)
	if err != nil {
		return err
	}

	return nil
}

func GetMethodData(r *http.Request,key string) (error,[]string) {
	data,ok := r.URL.Query()[key]
	if !ok || len(data) == 0 {
		return errors.New("参数有误"),nil
	}
	return nil,data
}
