package controller

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"go-gin-demo/database"
	"go-gin-demo/model"
	"log"
	"net/http"
)

var db *sql.DB

func init() {
	log.Println(">>>> get database connection start <<<<")
	db = database.GetDataBase()
}

// localhost:9090/user/query?id=2&name=hello
func QueryParam(context *gin.Context) {
	println(">>>> query user by url params action start <<<<")
	id := context.Query("id")
	name := context.Request.URL.Query().Get("name")

	var u model.User
	context.Bind(&u)
	context.ShouldBind(&u)

	println(u.Username)
	rows := db.QueryRow("select username,address,age,mobile,sex from t_user where id = $1 and username = $2", id, name)

	var user model.User
	err := rows.Scan(&user.Username, &user.Address, &user.Age, &user.Mobile, &user.Sex)
	checkError(err)

	checkError(err)
	context.JSON(200, gin.H{
		"result": user,
	})

}

func QueryById(context *gin.Context) {
	println(">>>> get user by id and name action start <<<<")

	id := context.Param("id")
	name := context.Param("username")

	rows := db.QueryRow("select username,address,age,mobile,sex from t_user where id = ? and username = ?", id, name)

	var user model.User

	err := rows.Scan(&user.Username, &user.Address, &user.Age, &user.Mobile, &user.Sex)
	checkError(err)

	checkError(err)
	context.JSON(200, gin.H{
		"result": user,
	})
}

func InsertNewUser(context *gin.Context) {
	println(">>>> insert controller action start <<<<")
	var user model.User

	// 直接将结构体和提交的json参数作绑定
	err := context.ShouldBindJSON(&user)

	// 写入数据库
	res, err := db.Exec("insert into t_user (username,sex,address,mobile,age) values (?,?,?,?,?)",
		&user.Username, &user.Sex, &user.Address, &user.Mobile, &user.Age)
	var count int64
	count, err = res.RowsAffected()
	checkError(err)
	if count != 1 {
		context.JSON(200, gin.H{
			"success": false,
		})
	} else {
		context.JSON(200, gin.H{
			"success": true,
		})
	}
}

func PostForm(context *gin.Context) {
	println(">>>> bind form post params action start <<<<")
	var u model.User

	// 绑定参数到结构体
	context.Bind(&u)
	context.ShouldBind(&u)
	res, err := db.Exec("insert into t_user (username,sex,address,mobile,age) values ($1,$2,$3,$4,$5)",
		&u.Username, &u.Sex, &u.Address, &u.Mobile, &u.Age)
	var count int64

	count, err = res.RowsAffected()
	checkError(err)

	if count != 1 {
		context.JSON(200, gin.H{
			"success": false,
		})
	} else {
		context.Redirect(http.StatusMovedPermanently, "/file/view")
	}
}

func RenderForm(context *gin.Context) {
	println(">>>> render to html action start <<<<")

	context.Header("Content-Type", "text/html; charset=utf-8")
	context.HTML(200, "insertUser.html", gin.H{})
}

func checkError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
