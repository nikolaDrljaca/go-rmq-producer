package models

type Message struct {
	UID     int
	Content string
}

type messagePayload struct {
	UID     int    `json:"uid"`
	Content string `json:"content"`
	Tag     string `json:"tag"`
}

func NewMessagePayload(content string, uid int) messagePayload {
	return messagePayload{
		Content: content,
		UID:     uid,
		Tag:     "msg",
	}
}
