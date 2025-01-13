package main

import (
  	"gorm.io/driver/sqlite" // Sqlite driver based on CGO
  	// "github.com/glebarez/sqlite" // Pure go SQLite driver, checkout https://github.com/glebarez/sqlite for details
  	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

//構造体User_tableの宣言
type User_table struct{
	USERNAME	string	`json:"username"`
	PASSWORD	string	`json:"password"`
}

var db *gorm.DB

func main(){
	//POSTを受け取る
	router:=gin.Default()
	router.POST("/somePost",posting)
	router.GET("/someGet",getting)
	router.PATCH("/somePatch",patching)
	router.Run(":8080")
}

func posting(c*gin.Context){
	var err error
	// github.com/mattn/go-sqlite3
	//SQLiteを開く
	db, err := gorm.Open(sqlite.Open("usertable.db"), &gorm.Config{})
	//つながらないとエラー返す
	if err != nil {
		panic("failed to connect to database")
	}
	var img User_table
	//構造に合わなければエラー返す
	if err := c.BindJSON(&img); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	//挿入
	result := db.Create(&img) // pass pointer of data to Create
	//できなければエラー返す
	if result.Error != nil{
		c.JSON(500,gin.H{"error":result.Error.Error()})
		return
	}
	//成功した際に送信
	c.JSON(200,gin.H{"message":"ログインに成功しました"})
}

func getting(c*gin.Context){
	var err error
	// github.com/mattn/go-sqlite3
	//SQLiteを開く
	db, err := gorm.Open(sqlite.Open("usertable.db"), &gorm.Config{})
	//つながらないとエラー返す
	if err != nil {
		panic("failed to connect to database")
	}
	var img User_table
	//構造に合わなければエラー返す
	if err := c.BindJSON(&img); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	//挿入
	result := db.Create(&img) // pass pointer of data to Create
	//できなければエラー返す
	if result.Error != nil{
		c.JSON(500,gin.H{"error":result.Error.Error()})
		return
	}
	//成功した際に送信
	c.JSON(200,gin.H{"message":"所持しているクレジットを表示します"})
}

//ユーザー登録
//1.ユーザーが既に登録されているか.
//2.1を満たさなければユーザーを登録する
//3.メッセージの返答
func patching(c*gin.Context){
	var err error
	// github.com/mattn/go-sqlite3
	//SQLiteを開く
	db, err := gorm.Open(sqlite.Open("usertable.db"), &gorm.Config{})
	//つながらないとエラー返す
	if err != nil {
		panic("failed to connect to database")
	}

	//1.既に登録されているか
	var newname USERNAME
	readresult := db.First(&newname, "username = ?", usertable.USERNAME)
	
	if readresult.Error := nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	//構造に合わなければエラー返す
	if err := c.BindJSON(&newname); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	//2.登録
	createresult := db.Create(&newname) // pass pointer of data to Create
	//できなければエラー返す
	if createresult.Error != nil{
		c.JSON(500,gin.H{"error":登録に失敗しました})
		return
	}
	//成功した際に送信
	c.JSON(200,gin.H{"message":"ユーザー登録に成功しました."})
}