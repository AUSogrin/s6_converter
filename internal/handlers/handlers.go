package handlers

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func HTMLHandler(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile("index.html")
	if err != nil {
		log.Printf("Failed to read HTML: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	if _, err := w.Write(data); err != nil {
		log.Printf("Failed to write response: %v", err)
	}
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusInternalServerError)
		return
	}

	text := r.FormValue("text")
	var filename string

	if text == "" {

		file, header, err := r.FormFile("myFile")
		if err != nil {
			http.Error(w, "Need either text or file", http.StatusInternalServerError)
			return
		}
		defer file.Close()

		data, err := io.ReadAll(file)
		if err != nil {
			http.Error(w, "Can't read file", http.StatusInternalServerError)
			return
		}
		text = string(data)
		filename = header.Filename

		result, err := service.Converter(text)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if filename != "" {
			ext := filepath.Ext(filename)
			saveName := "result_" + time.Now().Format("02.01.2006_15.04.05") + ext
			os.WriteFile(saveName, []byte(result), 0644)
		}

		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(result))
	}
}
