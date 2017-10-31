package main

import (
	"database/sql"
	_ "github.com/lib/pq"
)

var Db *sql.DB

// connect to the Db
func init() {
	var err error
	Db, err = sql.Open("postgres", "user=postgres dbname=trail-backend password=root sslmode=disable")
	if err != nil {
		panic(err)
	}
}

// Get a single trail category
func getTrailCategory(trail_category_id int) (post Post, err error) {
	post = Post{}
	err = Db.QueryRow("SELECT trail_category_id,trail_name FROM ms_trail_category WHERE trail_category_id = $1", trail_category_id).Scan(&post.Trail_category_id, &post.Trail_name)

	return
}

func getTrailCategoryList() (post []Post, err error) {
	post = []Post{}

	rows, err := Db.Query("SELECT trail_category_id,trail_name FROM ms_trail_category LIMIT $1", 100)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		// Scan one customer record
		var temp = Post{}
		if err := rows.Scan(&temp.Trail_category_id, &temp.Trail_name); err != nil {
			// handle error
		}
		post = append(post, temp)
	}
	if rows.Err() != nil {
		// handle error
	}
	return
}

// Create a new post
func (post *Post) create() (err error) {
	statement := "INSERT INTO ms_trail_category (trail_name) VALUES ($1) RETURNING trail_category_id"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(post.Trail_name).Scan(&post.Trail_category_id)
	return
}

// Update a post
func (post *Post) update() (err error) {
	_, err = Db.Exec("UPDATE ms_trail_category SET trail_name = $2 WHERE trail_category_id = $1", post.Trail_category_id, post.Trail_name)
	return
}

// Delete a post
func (post *Post) delete() (err error) {
	_, err = Db.Exec("DELETE FROM ms_trail_category WHERE trail_category_id = $1", post.Trail_category_id)
	return
}
