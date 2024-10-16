package main

import (
	"log"
	"net/http"
	"os"
	"path"
)

func main() {
	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("./build"))
	mux.Handle("/", fs)

	// mux.HandleFunc("GET /data/{link}", FileServer)

	http.ListenAndServe(":5000", mux)
}

func FileServer(w http.ResponseWriter, r *http.Request) {
	url := r.PathValue("link")

	fullPath := path.Join("./build", url)

	data, err := os.ReadFile(fullPath)
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		log.Println(err)
		return
	}
	w.Header().Set("Content-Type", http.DetectContentType(data))

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
