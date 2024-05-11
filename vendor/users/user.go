package users

import "errors"

type User struct {
	Username    string `json:"username"`
	DisplayName string `json:"displayname"`
	UUID        string `json:"id"`
}

type UserHandler struct {
	Database map[string]User
}

func (uh UserHandler) GetUserByUsername(username string) (*User, error) {
	for _, user := range uh.Database {
		if user.Username == username {
			return &user, nil
		}
	}
	return nil, errors.New("No user found for " + username)
}

func (uh UserHandler) ListUsers() ([]User, error) {
	userArray := []User{}
	for _, user := range uh.Database {
		userArray = append(userArray, user)
	}
	return userArray, nil
}

func (uh UserHandler) RegisterUser(user User) (error) {
	uh.Database[user.UUID] = user
	return nil
}
