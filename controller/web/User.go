package web

import (
	"log"
	"net/http"
	. "web/database"
	. "web/model"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Name string `form:"name"`
	Age  byte   `form:"age"`
	Sex  byte   `form:"sex"`
}

func UserIndex(context *gin.Context) {
	var person Person
	bind := context.ShouldBind(&person)
	if bind == nil {
		log.Println("====== Only Bind By Query String ======")
		log.Println(person.Name)
		log.Println(person.Age)
		log.Println(person.Sex)
	} else {
		log.Println("====== Bind Error ======")
		log.Println(bind)
	}
	//panic(http.ErrBodyNotAllowed)
	context.String(http.StatusOK, "User 222")
}

func UserName(context *gin.Context) {
	name := context.Param("name")
	var user User
	Db.First(&user, 1)
	context.JSON(http.StatusOK, gin.H{
		"name": name,
		"data": user,
	})
}

func UserUpdate(context *gin.Context) {
	message := context.PostForm("message")
	nick := context.DefaultPostForm("nick", "name")
	context.JSON(http.StatusOK, gin.H{
		"status":  "posted",
		"message": message,
		"nick":    nick,
	})
}
