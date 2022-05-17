package models

type Users struct {
	User_id   uint   `json: "user_id"`
	User_name string `json: "user_name"`
}

type Messages struct {
	Message_id uint   `json: "message_id"`
	Message    string `json: "message"`
}

type error struct {
	Error string `json: "error"`
}
