package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

const (
	ServerPort             = ":8000"
	ResultCodeSuccess      = "0"
	ResultCodeFail         = "9"
	ResultCodeLoginFail    = "1"
	ResultCodeNeedLogin    = "2"
	ResultCodeNoPermission = "3"
	ResultCodeNotFound     = "4"
)

type Result struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (r Result) responseJSON(context *gin.Context) {
	context.JSON(http.StatusOK, r)
}

func ResultSuccess(data interface{}) *Result {
	return &Result{Code: ResultCodeSuccess, Msg: "Success", Data: data}
}

func ResultFail(msg string) *Result {
	return &Result{Code: ResultCodeFail, Msg: msg, Data: nil}
}

func ResultFailWithCode(code string, msg string) *Result {
	return &Result{Code: code, Msg: msg, Data: nil}
}

func StartServer() {
	router := gin.Default()

	router.Any("/ping", Pong)

	RouteKu(router)

	if err := router.Run(ServerPort); err != nil {
		log.Fatalln(err)
	}
}

func Pong(context *gin.Context) {
	ResultSuccess("pong").responseJSON(context)
}

func main() {

	//ConnectToRedis()

	ConnectToDatabase()

	StartServer()
}
