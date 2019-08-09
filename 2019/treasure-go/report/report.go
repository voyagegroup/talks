package main

import (
	"database/sql"
	"flag"
	"time"
)

type Report struct {
	Title string
	Body  string
}

func (r *Report) Insert(db *sql.DB) (sql.Result, error) {
	stmt, err := db.Prepare(`insert into (title, body) values(?, ?)`)
	if err != nil {
		return err
	}
	defer stmt.Close()
	return stmt.Exec(r.Title, r.Body)
}

func Select(db *sql.DB) ([]Report, error) {
	rows, err := db.Query(`select * from reports`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	reports := make([]Report{})
	for rows.Next() {
		var title string
		var body string
		if err := rows.Scan(&title, &body); err != nil {
			return nil, err
		}
		r := Report{Title: title, Body: body}
		reports = append(reports, r)
	}
	return reports, nil
}

func main() {
	var (
		title = flag.String("title", "kuke", "cool title")
		body  = flag.String("body", "treasure", "cool body")
	)
	flag.Parse()
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatalf("open failed: %s", err)
	}
	r := &Report{*title, *body}
	if _, err := r.Insert(db); err != nil {
		log.Fatalf("insert failed: %s", err)
	}

	result, err := Select(db)
	if err != nil {
		log.Fatalf("select failed: %s", err)
	}
	fmt.Printf("%v", result)
}
