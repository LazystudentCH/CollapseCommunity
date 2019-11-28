package models

const (
	Success = iota
	Error
	Fail
	TokenExpire
)


type Result struct {
	Code int
	Msg string
	Data interface{}
}
