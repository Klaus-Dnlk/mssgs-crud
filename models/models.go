package Models

// type User struct {
// 	Id   uint   `gorm: "primaryKey"`
// 	Name string `gorm: "unique"`
// }

// type Message struct {
// 	Id          uint `gorm: "primaryKey"`
// 	SenderId    uint `gorm:"foreignKey: sender_fk; references: Id"`
// 	RecipientId uint `gorm:"foreignKey: recipient_fk; references: Id"`
// 	Message     string 
// }



type User struct {
	Id   uint   `gorm: "primaryKey"`
	Name string `gorm: "unique"`
}

type Message struct {
	Id          uint `gorm: "primaryKey"`
	SenderId    uint `gorm:"foreignKey: sender_fk; references: Name"`
	RecipientId uint `gorm:"foreignKey: recipient_fk; references: Name"`
	Message     string 
}
