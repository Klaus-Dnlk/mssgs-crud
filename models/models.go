package Models

// import (
// 	"gorm.io/gorm"
// )
// MODEL WITH ID REFERENCE

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

// -----------------------------------


// MODEL WITH NAME REFERENCE


type User struct {
	UserId   	uint   `gorm:"primary_key"`
	Name 		string `gorm:"index:,unique"`
	SentMessages []Message `gorm:"foreignKey:SenderId"`
	ReceivedMessages []Message `gorm:"foreignKey:RecipientId"`
	// Messages []*Message `gorm:"many2many:user_messages;"`
 }

type Message struct {
	MessageId   uint `gorm:"primary_key"`
	Message     string 
	SenderId 	uint `gorm:"index:"`
	RecipientId uint `gorm:"index:"`
	// Sender User `gorm:"foreignKey:UserId;references:SenderId"`
	// Recipient User `gorm:"foreignKey:UserId;references:RecipientId"`
	// Users []*User `gorm:"many2many:user_messages;"`
}
