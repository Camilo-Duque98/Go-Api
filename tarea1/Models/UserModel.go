package Models

type User struct {
	Id      uint   `json:"Id"`
	Name    string `json:"Name"`
	Password   string `json:"Password"`
}

func (b *User) TableName() string {
	return "user"
}