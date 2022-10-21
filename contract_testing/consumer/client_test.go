package main

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMessageServiceShouldReturnMessageWithID10(t *testing.T) {
	assertions := assert.New(t)

	// given:
	const messageID = "10"
	expectedEndpoint := fmt.Sprintf("/messages/%s", messageID)
	expectedMessage := Message{
		ID:      messageID,
		Author:  "John Doe",
		Content: "Example content",
	}

	testSrv := newTestHTTPServer(assertions, expectedEndpoint, expectedMessage)
	defer testSrv.Close()

	URL := ParseToURL(testSrv.URL)
	SUT := NewMessageClient(URL)

	// when:
	message, err := SUT.MessageWithID(messageID)

	// then:
	assertions.Nil(err)
	assertions.Equal(expectedMessage, message)
}

func newTestHTTPServer(assertions *assert.Assertions, expectedEndpoint string, expectedMessage Message) *httptest.Server {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		URL := req.URL.Path
		assertions.Equal(URL, expectedEndpoint)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		enc := json.NewEncoder(w)
		enc.SetIndent("", "\t")
		err := enc.Encode(expectedMessage)
		if err != nil {
			panic(err)
		}
	}))
	return s
}
