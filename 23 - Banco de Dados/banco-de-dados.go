package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3" // Driver SQLite
)

func main() {
	//urlConexao := "usuario:senha/database?charset=utf8&parseTime=true&loc=Local"
	//urlConexao := "devbook:AReallyStrongPasswd@tcp(localhost:5000)/devbook?charset=utf8&parseTime=true&loc=Local&tls=true&timeout=30s"
	urlConexao := "root:password@2@/devbook?charset=utf8&parseTime=true&loc=Local&tls=true&timeout=30s"
	//db, erro := sql.Open("mysql", urlConexao)
	//db, erro := sql.Open("sqlite3", "./sqlite.db")
	//db.SetConnMaxLifetime(time.Minute * 4)
	db, erro := sql.Open("mysql", urlConexao)
	if erro != nil {
		log.Fatal(erro)
	}

	_, erro = db.Exec("CREATE TABLE `users` (`id` INTEGER PRIMARY KEY AUTOINCREMENT, `name` VARCHAR(30))")
	if erro != nil {
		log.Fatal(erro)
	}

	defer db.Close()

	if erro = db.Ping(); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println("Connection is open!")

	lines, erro := db.Query("select * from users")
	if erro != nil {
		log.Fatal(erro)
	}
	defer lines.Close()

	fmt.Println(lines)

}
