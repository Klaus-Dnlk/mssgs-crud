package Models

type User struct {
	Id   uint   `json: "user_id" gorm: "primaryKey"`
	Name string `json: "name"`
}

type Message struct {
	Id          uint   `json: "message_id" gorm: "primaryKey"`
	SenderId    uint   `json: "sender_id" gorm:"foreignKey: sender_fk; referencs: Id"`
	RecipientId uint   `json: "recipient_id" gorm:"foreignKey: recipient_fk; referencs: Id"`
	Message     string `json: "message"`
}
