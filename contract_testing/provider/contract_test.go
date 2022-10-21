package main

import (
	"fmt"
	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
	"github.com/pact-foundation/pact-go/utils"
	"github.com/stretchr/testify/assert"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"testing"
)

func TestContractMessageServiceShouldReturnResponseWithIDField(t *testing.T) {
	assertions := assert.New(t)

	// given:
	pact := &dsl.Pact{
		Host:     "localhost",
		Provider: "MESSAGE_SERVICE",
		LogLevel: "INFO",
	}

	pact.Setup(true)
	defer pact.Teardown()

	cache := NewMessageCache()
	cache.InsertMessages(Message{
		ID:      "10",
		Author:  "John Doe",
		Content: "Example content",
	})

	addr := generateProviderAddr()
	go startProviderAPI(cache, addr)

	// when:
	_, err := pact.VerifyProvider(t, types.VerifyRequest{
		ProviderBaseURL: fmt.Sprintf("http://%s", addr),
		Tags:            []string{"master"},
		ProviderVersion: "1.0.0",
		PactURLs:        []string{filepath.FromSlash(fmt.Sprintf("%s/consumer/pacts/message_service_client-message_service.json", getParentDirectory()))},
	})

	// then:
	assertions.Nil(err)
}

func getParentDirectory() string {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	p := filepath.Dir(dir)
	return p
}

func generateProviderAddr() string {
	port, err := utils.GetFreePort()
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("localhost:%d", port)
}

func startProviderAPI(c *MessageCache, addr string) {
	srv := http.Server{
		Addr:    addr,
		Handler: NewMux(c),
	}
	defer func() {
		err := srv.Close()
		if err != nil {
			panic(err)
		}
	}()

	log.Printf("Messages service API listens and serves on: %s", srv.Addr)
	err := srv.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
