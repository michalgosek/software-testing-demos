package main

import (
	"fmt"
)

func main() {
	URL := ParseToURL("http://localhost:8081")
	cli := NewMessageClient(URL)

	message, err := cli.MessageWithID("10")
	if err != nil {
		fmt.Printf("Endpoint Error: %s\n", err)
	} else {
		fmt.Println(message)
	}
}
