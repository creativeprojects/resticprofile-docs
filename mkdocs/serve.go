package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func serveDirectory(root string) error {
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT)
	defer signal.Stop(done)

	rootFS, err := os.OpenRoot(root)
	if err != nil {
		return err
	}
	http.Handle("GET /", http.FileServerFS(rootFS.FS())) // registered in the default handler

	fmt.Printf("Listening on http://localhost:1313/\n")
	server := &http.Server{Addr: "localhost:1313", Handler: http.DefaultServeMux}

	go func() {
		<-done
		_ = server.Close()
	}()
	err = server.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	close(done)
	return nil
}
