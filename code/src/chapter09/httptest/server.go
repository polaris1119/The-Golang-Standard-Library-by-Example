package main

import (
	"encoding/json"
	"net/http"
	"path"
	"strconv"
)

func main() {
	http.HandleFunc("/topic/", handleRequest)

	http.ListenAndServe(":2017", nil)
}

// main handler function
func handleRequest(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case http.MethodGet:
		err = handleGet(w, r)
	case http.MethodPost:
		err = handlePost(w, r)
	case http.MethodPut:
		err = handlePut(w, r)
	case http.MethodDelete:
		err = handleDelete(w, r)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// 获取一个帖子
// 如 GET /topic/1
func handleGet(w http.ResponseWriter, r *http.Request) error {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return err
	}
	topic, err := FindTopic(id)
	if err != nil {
		return err
	}
	output, err := json.MarshalIndent(&topic, "", "\t\t")
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return nil
}

// 增加一个帖子
// POST /topic/
func handlePost(w http.ResponseWriter, r *http.Request) (err error) {
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)
	var topic = new(Topic)
	err = json.Unmarshal(body, &topic)
	if err != nil {
		return
	}

	err = topic.Create()
	if err != nil {
		return
	}
	w.WriteHeader(http.StatusOK)
	return
}

// 更新一个帖子
// PUT /topic/1
func handlePut(w http.ResponseWriter, r *http.Request) error {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return err
	}
	topic, err := FindTopic(id)
	if err != nil {
		return err
	}
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)
	json.Unmarshal(body, topic)
	err = topic.Update()
	if err != nil {
		return err
	}
	w.WriteHeader(http.StatusOK)
	return nil
}

// 删除一个帖子
// DELETE /topic/1
func handleDelete(w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	topic, err := FindTopic(id)
	if err != nil {
		return
	}
	err = topic.Delete()
	if err != nil {
		return
	}
	w.WriteHeader(http.StatusOK)
	return
}
