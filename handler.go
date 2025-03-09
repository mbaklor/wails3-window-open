package main

import (
	"embed"
	"io"
	"net/http"
)

func AssetHandler(assets embed.FS) http.Handler {
	h := http.NewServeMux()

	h.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		f, err := assets.Open("frontend/dist/index.html")
		if err != nil {
			w.WriteHeader(400)
			return
		}

		io.Copy(w, f)
	}))
	h.Handle("/settings", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		f, err := assets.Open("frontend/dist/settings.html")
		if err != nil {
			w.WriteHeader(400)
			return
		}

		io.Copy(w, f)
	}))

	return h
}
