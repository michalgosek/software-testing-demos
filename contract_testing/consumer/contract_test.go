package main

import (
	"fmt"
	"github.com/pact-foundation/pact-go/dsl"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestContract_MessageServiceShouldReturnMessageWithID10(t *testing.T) {
	assertions := assert.New(t)

	// given:
	pact := &dsl.Pact{
		Host:     "localhost",
		Consumer: "MESSAGE_SERVICE_CLIENT",
		Provider: "MESSAGE_SERVICE",
		LogLevel: "INFO",
	}

	pact.Setup(true)
	defer pact.Teardown()

	URL := ParseToURL(fmt.Sprintf("http://localhost:%d", pact.Server.Port))
	SUT := NewMessageClient(URL)

	const messageID = "10"

	pact.AddInteraction().
		Given("Message with ID 10 is available").
		UponReceiving("GET request for message with ID 10").
		WithRequest(dsl.Request{
			Method:  http.MethodGet,
			Path:    dsl.Term("api/v1/messages/10", "api/v1/messages/[0-9]+"),
			Headers: dsl.MapMatcher{"Content-Type": dsl.String("application/json")},
		}).
		WillRespondWith(dsl.Response{
			Status:  http.StatusOK,
			Headers: dsl.MapMatcher{"Content-Type": dsl.String("application/json")},
			Body:    dsl.Like(Message{}),
		})

	var expectedMessage Message

	// when:
	err := pact.Verify(func() error {
		message, err := SUT.MessageWithID(messageID)
		if err != nil {
			return err
		}
		assertions.Nil(err)
		assertions.Equal(expectedMessage, message)
		return nil
	})

	// then:
	assertions.Nil(err)
}
