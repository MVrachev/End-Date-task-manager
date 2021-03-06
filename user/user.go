package user

import "github.com/jinzhu/gorm"

// User struct will handle user input and
// will send requests to the server
type User struct {
	gorm.Model
	Name     string `json:"name"`
	Password string `json:"password"`
}

// NewUser creates a new User
func NewUser(name string, pass string) User {
	return User{
		Name:     name,
		Password: pass,
	}
}
