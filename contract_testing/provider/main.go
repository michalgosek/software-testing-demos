package main

import (
	"log"
	"net/http"
)

/*
	https://github.com/pact-foundation/homebrew-pact-ruby-standalone
*/

func main() {
	cache := NewMessageCache()
	cache.InsertMessages(
		Message{
			ID:      "1",
			Author:  "John Doe",
			Content: "Example 1",
		},
		Message{
			ID:      "2",
			Author:  "Jane Doe",
			Content: "Example content 1",
		},
	)

	const addr = "localhost:8081"
	srv := http.Server{
		Addr:    addr,
		Handler: NewMux(cache),
	}

	defer func() {
		err := srv.Close()
		if err != nil {
			panic(err)
		}
	}()

	log.Printf("Messages service API listens and serves on: (%s)", addr)
	err := srv.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
