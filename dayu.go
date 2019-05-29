package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

const (
	ServerPort        = ":8000"
	ResultCodeSuccess = "0"
	ResultCodeFail    = "9"
)

type Result struct {
	Code string      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func ResultSuccess(data interface{}) *Result {
	return &Result{Code: ResultCodeSuccess, Msg: "Success", Data: data}
}

func ResultFail(msg string) *Result {
	return &Result{Code: ResultCodeFail, Msg: msg, Data: nil}
}

func StartServer() {
	router := gin.Default()

	router.Any("/ping", Pong)

	RouteKu(router)

	err := router.Run(ServerPort)
	if err != nil {
		log.Fatalln(err)
	}
}

func Pong(context *gin.Context) {
	context.JSON(http.StatusOK, ResultSuccess("pong"))
}

func main() {

	//ConnectToRedis()

	ConnectToDatabase()

	StartServer()
}
