package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func UserByAge(db *sql.DB, age int) (names []string, err error) {
	rows, err := db.Query(
		`select name from users where age = ?`,
		age)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		names = append(names, name)
	}
	return names, nil
}

func main() {
	var (
		age = flag.Int("age", 10, "user's age")
	)
	db, err := sql.Open("mysql", "")
	if err != nil {
		log.Fatalf("open db failed: %s", err)
	}
	users, err := UserByAge(db, *age)
	if err != nil {
		log.Fatalf("query failed: %s", err)
	}
	fmt.Printf("%v", users)
}
