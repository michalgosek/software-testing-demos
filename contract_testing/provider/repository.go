package main

type Message struct {
	ID      string `json:"id"`
	Author  string `json:"author"`
	Content string `json:"content"`
}

func (m *Message) IsEmpty() bool {
	ss := []string{m.Content, m.Author, m.ID}
	for _, s := range ss {
		if s == "" {
			return true
		}
	}
	return false
}

type MessageCache struct {
	messages map[string]Message
}

func (m *MessageCache) InsertMessages(messages ...Message) {
	for _, s := range messages {
		m.messages[s.ID] = s
	}
}

func (m *MessageCache) MessageWithID(ID string) Message {
	v, ok := m.messages[ID]
	if !ok {
		return Message{}
	}
	return v
}

func (m *MessageCache) Messages() []Message {
	if len(m.messages) == 0 {
		return []Message{}
	}
	var messages []Message
	for _, m := range m.messages {
		messages = append(messages, m)
	}
	return messages
}

func NewMessageCache() *MessageCache {
	return &MessageCache{messages: make(map[string]Message)}
}
