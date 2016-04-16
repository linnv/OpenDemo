// Package main provides ...
package newDir

import (
	"fmt"
	"log"
	"os"
	"time"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func JustDemo() {
	println("<<<JustDemo start---------------------------")
	println("-----------------------------JustDemo end>>>")
	return

}

var (
	user      string
	pass      string
	prot      string
	addr      string
	dbname    string
	dsn       string
	netAddr   string
	available bool
	_db       *sql.DB
)

func init() {
	var err error
	_db, err = sql.Open("mysql", GetDNS())
	if err != nil {
		log.Fatalf("Error connecting: %s", err.Error())
	}
}

func GetDB() *sql.DB {
	if _db == nil {
		var err error
		_db, err = sql.Open("mysql", GetDNS())
		if err != nil {
			log.Fatalf("Error connecting: %s", err.Error())
		}
	}
	return _db
}

func GetDNS() string {
	if netAddr != "" {
		return fmt.Sprintf("%s:%s@%s/%s?timeout=30s&strict=true", user, pass, netAddr, dbname)
	}

	env := func(key, defaultValue string) string {
		if value := os.Getenv(key); value != "" {
			return value
		}
		return defaultValue
	}
	user = env("MYSQL_TEST_USER", "root")
	pass = env("MYSQL_TEST_PASS", "root")
	prot = env("MYSQL_TEST_PROT", "tcp")
	// addr = env("MYSQL_TEST_ADDR", "localhost:3306")
	addr = env("MYSQL_TEST_ADDR", "127.0.0.1:3306")
	dbname = env("MYSQL_TEST_DBNAME", "test")
	netAddr = fmt.Sprintf("%s(%s)", prot, addr)
	return fmt.Sprintf("%s:%s@%s/%s?timeout=30s&strict=true", user, pass, netAddr, dbname)
}

func QueryDemo() {
	println("//<<-------------------------QueryDemo start-----------")
	start := time.Now()

	// fmt.Printf("GetDNS(): %+v\n", GetDNS())
	// db, err := sql.Open("mysql", GetDNS())
	// if err != nil {
	// 	log.Fatalf("Error connecting: %s", err.Error())
	// }
	// defer db.Close()
	db := GetDB()
	query := "SELECT * FROM test_table"
	// rows, err := db.Query(query, args...)
	rows, err := db.Query(query)
	if err != nil {
		panic(err.Error())
		return
	}
	// fmt.Printf("rows: %+v\n", rows)
	for rows.Next() {
		var a, b, c string
		err = rows.Scan(&a, &b, &c)
		if err != nil {
			panic(err.Error())
			return
		}
		// fmt.Printf("a: %+v\n", a)
		// fmt.Printf("b: %+v\n", b)
		// fmt.Printf("c: %+v\n", c)
		// println("------ pause \n========================\n")
	}
	// return rows
	fmt.Printf("%v as %v millisecons\n", time.Since(start), time.Since(start).Nanoseconds()/1000000)
	println("//---------------------------QueryDemo end----------->>")
}

func checkErr(err error) {
	if err != nil {
		panic(err.Error())
		return
	}
}

func UpdateDemo() {
	println("//<<-------------------------UpdateDemo start-----------")
	start := time.Now()
	fmt.Printf("GetDNS(): %+v\n", GetDNS())
	db, err := sql.Open("mysql", GetDNS())
	if err != nil {
		log.Fatalf("Error connecting: %s", err.Error())
	}
	defer db.Close()
	//更新数据
	stmt, err := db.Prepare("update test_table set a=? where a=?")
	checkErr(err)

	res, err := stmt.Exec("JialinWu", "astaxie")
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)
	fmt.Printf("UpdateDemo costs  %d millisecons actually %v\n", time.Since(start).Nanoseconds()/1000000, time.Since(start))
	println("//---------------------------UpdateDemo end----------->>")
}

func InsertDemo() {
	println("//<<-------------------------InsertDemo start-----------")
	start := time.Now()

	// fmt.Printf("GetDNS(): %+v\n", GetDNS())
	// db, err := sql.Open("mysql", GetDNS())
	// if err != nil {
	// 	log.Fatalf("Error connecting: %s", err.Error())
	// }
	// defer db.Close()

	db := GetDB()
	// stmt, err := db.Prepare("INSERT test_table(a,b,c) values(?,?,?)")
	stmt, err := db.Prepare("INSERT test_table(a,b,c) values(?,?,?)")
	checkErr(err)
	res, err := stmt.Exec("jialinwu", "研发部门", "2016-03-22 11:43:15")
	checkErr(err)

	id, err := res.LastInsertId()
	fmt.Printf("id: %+v\n", id)
	checkErr(err)
	fmt.Printf("%v as %v millisecons\n", time.Since(start), time.Since(start).Nanoseconds()/1000000)
	println("//---------------------------InsertDemo end----------->>")
}