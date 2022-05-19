package models

type User struct {
	id   uint   `json: "id" gorm: "primaryKey`
	name string `json: "name"`
}

type Message struct {
	id      uint   `json: "message_id"`
	message string `json: "message"`
}
