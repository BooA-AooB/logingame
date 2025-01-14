package main

import (
  	"gorm.io/driver/mysql" 
  	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
	//"github.com/golang-jwt/jwt/v5"
)

//構造体Userの宣言
type User struct{
	ID			uint	`json:"id" gorm:"primaryKey"`
	USERNAME	string	`json:"username"`
	PASSWORD	string	`json:"password"`
	CREDIT		uint	`json:"credit"`
}


var db *gorm.DB

func (User) TableName() string {
    return "usertable"
}

func main(){
	var err error
	dsn := "root:7gh342Fio-55aLNS@tcp(127.0.0.1:3306)/usertable?charset=utf8mb4&parseTime=True&loc=Local"
  	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}
	router:=gin.Default()
	router.POST("/login",login)
	router.GET("/credit",getcredit)
	router.POST("/register",register)
	router.Run(":8080")
}

//ログインを行う
//1.送信されたユーザー名とパスワードを確認
//2.一致すればJWTを発行
//3.メッセージの返答
func login(c*gin.Context){
	var user User
	//構造に合わなければエラー返す
	if err := c.BindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	//挿入
	result := db.Where("username = ? AND password = ?", user.USERNAME, user.PASSWORD).Take(&user) 
	//できなければエラー返す
	if result.Error != nil{
		c.JSON(500,gin.H{"error":result.Error.Error()})
		return
	}
	//成功した際に送信
	c.JSON(200,gin.H{"message":"ログインに成功しました"})
}

//所持しているクレジットを表示
//1.ログインしているかJWTで確認
//2.ログインしているならば所持しているクレジットを表示
//3.メッセージの返答
func getcredit(c*gin.Context){
	var credit User
	//構造に合わなければエラー返す
	if err := c.BindJSON(&credit); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	//クレジットを取得
	var wantuser User
	result := db.Where("username = ?", credit.USERNAME).First(&wantuser)
	//できなければエラー返す
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}
	c.JSON(200, gin.H{"credit": wantuser.CREDIT})

	//成功した際に送信
	c.JSON(200,gin.H{"message":"所持しているクレジットを表示します"})
}

//ユーザー登録
//1.送信されたユーザー名が登録されていないか確認
//2.登録されていなければユーザーとパスワードを登録する
//3.メッセージの返答
func register(c*gin.Context){

	//1.既に登録されているか
	var newname User
	readresult := db.First(&newname, "username = ?", newname.USERNAME)
	
	if readresult.Error != nil {
		c.JSON(400, gin.H{"error": readresult.Error.Error()})
		return
	}

	//構造に合わなければエラー返す
	if err := c.BindJSON(&newname); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	//2.登録

	newUser := User{
		USERNAME: newname.USERNAME,
		PASSWORD: newname.PASSWORD,
		CREDIT: 100,
	}

	result := db.Create(&newUser) 

	//できなければエラー返す
	if result.Error != nil{
		c.JSON(500,gin.H{"error":"登録に失敗しました"})
		return
	}
	//成功した際に送信
	c.JSON(200,gin.H{"message":"ユーザー登録に成功しました.記念として100クレジットをプレゼントします"})
}