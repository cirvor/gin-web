package web

import (
	"log"
	. "web/database"
	. "web/model"
	"web/utils"

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
	utils.Success(context, "User 222")
}

func UserName(context *gin.Context) {
	name := context.Param("name")
	var user User
	Db.First(&user, 1)
	utils.Success(context, gin.H{
		"name": name,
		"user": user,
	})
}

func UserUpdate(context *gin.Context) {
	message := context.PostForm("message")
	nick := context.DefaultPostForm("nick", "name")
	utils.Success(context, gin.H{
		"nick":    nick,
		"message": message,
	})
}
