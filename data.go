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

// Get a single trial category
func getTrialCategory(trial_category_id int) (post Post, err error) {
	post = Post{}
	err = Db.QueryRow("SELECT trial_category_id,trial_name FROM ms_trial_category WHERE trial_category_id = $1", trial_category_id).Scan(&post.Trial_category_id, &post.Trial_name)

	return
}

func getTrialCategoryList() (post []Post, err error) {
	post = []Post{}

	rows, err := Db.Query("SELECT trial_category_id,trial_name FROM ms_trial_category LIMIT $1", 100)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		// Scan one customer record
		var temp = Post{}
		if err := rows.Scan(&temp.Trial_category_id, &temp.Trial_name); err != nil {
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
	statement := "INSERT INTO ms_trial_category (trial_name) VALUES ($1) RETURNING trial_category_id"
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	err = stmt.QueryRow(post.Trial_name).Scan(&post.Trial_category_id)
	return
}

// Update a post
func (post *Post) update() (err error) {
	_, err = Db.Exec("UPDATE ms_trial_category SET trial_name = $2 WHERE trial_category_id = $1", post.Trial_category_id, post.Trial_name)
	return
}

// Delete a post
func (post *Post) delete() (err error) {
	_, err = Db.Exec("DELETE FROM ms_trial_category WHERE trial_category_id = $1", post.Trial_category_id)
	return
}
