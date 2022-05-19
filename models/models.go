package Models

type User struct {
	Id   uint   `json: "user_id" gorm: "primaryKey"`
	Name string `json: "name"`
}

type Message struct {
	Id      uint   `json: "message_id"  gorm: "primaryKey"`
	Message string `json: "message"`
}
