package main

import (
	"encoding/json"
	"net/http"
	"path"
	"strconv"
	"log"
	"github.com/gorilla/mux"
)

type AllPost struct {
	Posts []Post `json:"posts"`
}

type Post struct {
	Trial_category_id int    `json:"trial_category_id"`
	Trial_name        string `json:"trial_name"`
}

func main() {
	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(pageNotFound)
	r.HandleFunc("/api/trial-category/{trial_category_id}", handleSingleRequest)
	r.HandleFunc("/api/all-trial-category/", handleMultipleRequest)

	log.Fatal(http.ListenAndServe(":8080", r))
}

func pageNotFound(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "404 Page not found", 404)
	return
}

// main handler function
func handleSingleRequest(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		err = handleSingleGet(w, r)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func handleMultipleRequest(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		err = handleMultipleGet(w, r)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func handleSingleGet(w http.ResponseWriter, r *http.Request) (err error) {
	trial_category_id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	post, err := getTrialCategory(trial_category_id)
	if err != nil {
		return
	}
	output, err := json.MarshalIndent(&post, "", "\t\t")
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}

func handleMultipleGet(w http.ResponseWriter, r *http.Request) (err error) {
	post, err := getTrialCategoryList()
	if err != nil {
		return
	}
	output, err := json.Marshal(post)
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}
