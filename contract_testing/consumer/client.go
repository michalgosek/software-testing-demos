package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
)

const (
	MessageEndpoint  = "/messages/%s"
	MessagesEndpoint = "/messages/"
)

type Message struct {
	ID      string `json:"id"`
	Author  string `json:"author"`
	Content string `json:"content"`
}

type MessageClient struct {
	cli *http.Client
	url *url.URL
}

func NewMessageClient(URL *url.URL) *MessageClient {
	return &MessageClient{
		url: URL,
		cli: http.DefaultClient,
	}
}

func (m *MessageClient) getRequest(endpoint string) (*http.Response, error) {
	ref := &url.URL{Path: endpoint}
	URL := m.url.ResolveReference(ref).String()
	req, err := http.NewRequest(http.MethodGet, URL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	res, err := m.cli.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode == http.StatusNotFound {
		return nil, errors.New(http.StatusText(http.StatusNotFound))
	}
	return res, nil
}

func (m *MessageClient) MessageWithID(ID string) (Message, error) {
	URL := fmt.Sprintf(MessageEndpoint, ID)
	res, err := m.getRequest(URL)
	if err != nil {
		return Message{}, err
	}
	defer func() {
		err := res.Body.Close()
		if err != nil {
			panic(err)
		}
	}()

	var dst Message
	dec := json.NewDecoder(res.Body)
	err = dec.Decode(&dst)
	if err != nil {
		return Message{}, err
	}
	return dst, nil
}

func (m *MessageClient) Messages() ([]Message, error) {
	res, err := m.getRequest(MessagesEndpoint)
	if err != nil {
		return nil, err
	}
	defer func() {
		err := res.Body.Close()
		if err != nil {
			panic(err)
		}
	}()

	var dst []Message
	dec := json.NewDecoder(res.Body)
	err = dec.Decode(&dst)
	if err != nil {
		return nil, err
	}
	return dst, nil
}

func ParseToURL(s string) *url.URL {
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	return u
}
