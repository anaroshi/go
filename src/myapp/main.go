package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 입력 받을 구조체
type Login struct {
	Id  string `json:"id" xml:"id" form:"id" query:"id"`
	Pwd string `json:"pwd" xml:"pwd" form:"pwd" query:"pwd"`
}

// 결과값 구조체
type LoginExport struct {
	Id     string `json:"id" xml:"id" form:"id" query:"id"`
	Pwd    string `json:"pwd" xml:"pwd" form:"pwd" query:"pwd"`
	Result string `json:"result" xml:"result" form:"result" query:"result"`
}

type sendData struct {
	idx     int
	name    string
	id      string
	pwd     string
	regdate time.Time
}

type User struct {
	Name    string `json:"name" xml:"name" form:"name" query:"name"`
	Id      string `json:"id" xml:"id" form:"id" query:"id"`
	Pwd     string `json:"pwd" xml:"pwd" form:"pwd" query:"pwd"`
	Regdate time.Time
}

func login(c echo.Context) error {
	u := new(Login)
	if err := c.Bind(u); err != nil {
		return err
	}

	oriId := "sundor"
	oriPwd := "kemon"

	r := &LoginExport{
		Id:  u.Id,
		Pwd: u.Pwd,
	}

	if u.Id == oriId && u.Pwd == oriPwd {
		r.Result = "성공"
	} else {
		r.Result = "실폐"
	}

	return c.JSON(http.StatusCreated, r)
	// or
	// return c.XML(http.StatusCreated, u)
}

func join(c echo.Context) error {
	dsn := "sundor:anaroshi@tcp(localhost:3306)/goTest?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	errChk(err)

	j := new(User)
	if err := c.Bind(j); err != nil {
		return err
	}

	if len(j.Name) == 0 {
		//return c.HTML(http.StatusOK, "<script>window.location.href='http://localhost:8080/public/join.html';alert('사용자명을 입력하세요.');</script>")
		return c.HTML(http.StatusOK, "<script>window.location.href='http://localhost:8080/join';alert('사용자명을 입력하세요.');</script>")
	}
	if len(j.Id) < 5 {
		return c.HTML(http.StatusOK, "<script>window.location.href='http://localhost:8080/join';alert('아이디를 입력하세요.');</script>")
	}
	if len(j.Pwd) < 8 {
		return c.HTML(http.StatusOK, "<script>window.location.href='http://localhost:8080/join';alert('비밀번호를 입력하세요.');</script>")
	}

	ur := &User{
		Id:      j.Id,
		Pwd:     j.Pwd,
		Name:    j.Name,
		Regdate: time.Now(),
	}

	//db.Create(&ur)
	if err := db.Create(&ur).Error; err != nil {
		return c.HTML(http.StatusOK, "<script>window.location.href='http://localhost:8080/join';alert('사용되어진 ID입니다.');</script>")
	}

	//return c.HTML(http.StatusOK, "<h1>joining us is successfuly done.</h1>")
	//return c.HTML(http.StatusOK, "<script>window.location.href='http://localhost:8080/public/complete.html';</script>")
	return c.HTML(http.StatusOK, "<script>window.location.href='http://localhost:8080/complete';</script>")
}

func errChk(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func showLogin(c echo.Context) error {
	fmt.Println("start")

	dbname := "mysql"
	//url := "scsol:scsol92595@tcp(ithingsware.com:5022)/mysql"
	url := "sundor:anaroshi@tcp(127.0.0.1:3306)/goTest"

	db, err := sql.Open(dbname, url)
	errChk(err)
	defer db.Close()

	db.SetMaxIdleConns(20)
	db.SetMaxIdleConns(20)
	db.SetConnMaxLifetime(time.Minute * 5)

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	err = db.PingContext(ctx)
	if err != nil {
		log.Printf("Error %s pinging DB", err)
	}
	log.Printf("Connected to DB %s successfully\n", dbname)

	sql := "select idx, name, id, pwd from user"
	fmt.Println(sql)
	rows, err := db.Query(sql)
	errChk(err)
	defer rows.Close()
	
	var s sendData

	i := 0
	for rows.Next() {
		fmt.Println("next : ", i)
		err := rows.Scan(&s.idx, &s.name, &s.id, &s.pwd)
		errChk(err)
		fmt.Println(s.idx, s.name, s.id, s.pwd)
		i++
	}

	return c.HTML(http.StatusOK, "<h1>loginDB</h1>")
}

func (User) TableName() string {
	return "user"
}

func dbconn(c echo.Context) error {
	fmt.Println("db conn")
	// dsn := "sundor:Akemon97_@tcp(gotest.c7bbvlsum0db.us-east-1.rds.amazonaws.com:3306)/gotest?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "sundor:anaroshi@tcp(localhost:3306)/goTest?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	errChk(err)
	//defer db.Close()
	fmt.Println(db)

	user := User{
		Name: "Jinzhu",
		Id:   "Jinzhu",
		Pwd:  "pwd",
	}
	// Insert(Create a Record)
	db.Create(&user) // pass pointer of data to Create

	// Update a single column
	db.Model(User{}).Where("name = ?", "jinzhu").Updates(User{Name: "hello", Pwd: "hi"})

	// Delete a Record
	db.Where("Name = ?", "jinzhu").Delete(&user)

	return c.HTML(http.StatusOK, "<h1>loginDB</h1>")
}

func main() {
	e := echo.New()
	e.Static("/", "assets") // frontEnd files are loaded
	e.File("/login", "assets/public/login.html")
	e.POST("/login", login)
	e.File("/join", "assets/public/join.html")
	e.POST("/join", join)
	e.File("/complete", "assets/public/complete.html")
	e.GET("/show", showLogin)
	e.GET("/gorm", dbconn)
	e.Logger.Fatal(e.Start(":4000"))
}
