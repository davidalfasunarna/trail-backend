package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"path"
	"strconv"
)

type AllPost struct {
	Posts []Post `json:"posts"`
}

type Post struct {
	Trail_category_id int    `json:"trail_category_id"`
	Trail_name        string `json:"trail_name"`
}

func main() {
	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(pageNotFound)
	r.HandleFunc("/api/trail-category/", handleSingleRequest)
	r.HandleFunc("/api/trail-category/{trail_category_id}", handleSingleRequest)
	r.HandleFunc("/api/all-trail-category/", handleMultipleRequest)

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
	case "POST":
		err = handlePost(w, r)
	case "PUT":
		err = handlePut(w, r)
	case "DELETE":
		err = handleDelete(w, r)
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
	trail_category_id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	post, err := getTrailCategory(trail_category_id)
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
	post, err := getTrailCategoryList()
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

func handlePost(w http.ResponseWriter, r *http.Request) (err error) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	var post Post
	json.Unmarshal(body, &post)
	err = post.create()
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}

func handlePut(w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	post, err := getTrailCategory(id)
	if err != nil {
		return
	}
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	json.Unmarshal(body, &post)
	err = post.update()
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}

func handleDelete(w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	post, err := getTrailCategory(id)
	if err != nil {
		return
	}
	err = post.delete()
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}
