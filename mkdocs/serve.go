package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
)

func serveDirectory(root string) error {
	rootFS, err := os.OpenRoot(root)
	if err != nil {
		return err
	}
	http.Handle("GET /", http.FileServerFS(rootFS.FS())) // registered in the default handler

	fmt.Printf("Listening on http://localhost:1313/\n")
	server := &http.Server{Addr: "localhost:1313", Handler: http.DefaultServeMux}
	err = server.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}
