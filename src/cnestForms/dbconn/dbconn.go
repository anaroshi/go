package dbconn

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func GetInfoStudent(rw http.ResponseWriter, r *http.Request) {
	fmt.Println("start")
	dbname := "leak"
	// sql.DB 객체 생성
	
	db, err := sql.Open("mysql", "sundor:anaroshi@tcp(localhost:3306)/cst")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(20)
	db.SetConnMaxLifetime(time.Minute * 5)

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	err = db.PingContext(ctx)
	if err != nil {
		log.Printf("Errors %s pinging DB", err)
	}
	log.Printf("Connected to DB %s successfully\n", dbname)

	//복수 Row를 갖는 SQL 쿼리
	var cnt int
	var mem_id int
	var mem_username string
	var mem_school string
	var mem_year string

	
	rows, err := db.Query("SELECT mem_id, mem_username, mem_school, mem_year From cst_members")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close() //반드시 닫는다 (지연하여 닫기)
	fmt.Fprintln(rw,"<table>")
	for rows.Next() {
		err := rows.Scan(&mem_id, &mem_username, &mem_school, &mem_year)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintln(rw, "<tr><td>",mem_id, mem_username, mem_school, mem_year,"</td></tr>")
	}

	rows, err = db.Query("SELECT count(*) cnt From cst_members")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close() //반드시 닫는다 (지연하여 닫기)

	for rows.Next() {
		err := rows.Scan(&cnt)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintf(rw, "총 SN은 %d개 입니다.", cnt)
		
	}
	fmt.Fprintln(rw,"</table>")
}
 