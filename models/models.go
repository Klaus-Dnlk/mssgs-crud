package Models

type User struct {
	Id   uint   `json: "id" gorm: "primaryKey`
	Name string `json: "name"`
}

type Message struct {
	Id      uint   `json: "message_id"`
	Message string `json: "message"`
}
