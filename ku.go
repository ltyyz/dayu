package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	PasswordKey = "<#--mimamiyv-->"
)

func RouteKu(router *gin.Engine) {
	ku := router.Group("/ku")
	{
		ku.GET("/test", KuTest)
		ku.GET("/test_add", KuTestAdd)
		ku.GET("/test_query", KuTestQuery)
		ku.GET("/test_update", KuTestUpdate)
		ku.GET("/test_delete", KuTestDelete)

		ku.POST("/user/login", KuUserLogin)
	}
}

type User struct {
	ParentModel
	Name     string `json:"name"`
	Password string `json:"password"`
	Status   string `gorm:"default:1" json:"status"`
}

func KuTest(context *gin.Context) {
	context.JSON(http.StatusOK, ResultSuccess(Md5("Golang")))
}

func KuTestAdd(c *gin.Context) {
	user := User{Name: "Jack", Password: "111111"}
	user.Id = "4"

	if result := db.Create(&user); result.Error != nil {
		c.JSON(http.StatusOK, ResultFail(result.Error.Error()))

	} else {
		c.JSON(http.StatusOK, ResultSuccess(user))
	}
}

func KuTestQuery(context *gin.Context) {
	var users = make([]User, 0)

	db.Where("status = ?", "1").
		Limit(DefaultPageSize).
		Offset((DefaultPage - 1) * DefaultPageSize).
		Find(&users)

	ResultSuccess(users).responseJSON(context)
}

func KuTestUpdate(context *gin.Context) {
	var user = User{}
	user.Id = "1"

	db.Model(&user).Update("name", "Jack2")
	context.JSON(http.StatusOK, ResultSuccess(db.First(&user).Value))
}

func KuTestDelete(context *gin.Context) {
	var user = User{}
	user.Id = "4"
	db.Delete(&user)
	context.JSON(http.StatusOK, ResultSuccess(user))
}

// Test Code --END--

func KuUserLogin(context *gin.Context) {
	username := context.PostForm("username")
	password := context.PostForm("password")

	var user User
	if err := db.Where("name = ?", username).Find(&user).Error; err != nil {
		ResultFailWithCode(ResultCodeLoginFail, "").responseJSON(context)

	} else {
		password = Md5(password + PasswordKey)
		if password == user.Password {
			ResultSuccess("").responseJSON(context)

		} else {
			ResultFailWithCode(ResultCodeLoginFail, "").responseJSON(context)
		}
	}
}
