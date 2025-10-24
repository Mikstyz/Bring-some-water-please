package entities

type User struct {
	UserId     int `db:"Id"`
	TelegramId int `db:"tgId"`
}
