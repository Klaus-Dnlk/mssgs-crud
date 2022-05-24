package Models


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
