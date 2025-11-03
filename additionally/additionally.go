package additionally

import (
	"fmt"
)
var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
	"user3": "password3",
}

func login(username, password string) (bool, error) {
	pass, userExist := users[username]
	if username == "admin" {
		panic("CRITICAL: доступ администратора заблокирован")
	} else if !userExist {
		return false, &UserNotFoundError{username}
	} else if password != pass {
		return false, &WrongPasswordError{username}
	}
	return true, nil
}

type UserNotFoundError struct {
	Username string
}
func (U *UserNotFoundError) Error() string {
	return fmt.Sprintf("Не найден юзер с логином: %s", U.Username)
}

type WrongPasswordError struct {
	Username string
}
func (W *WrongPasswordError) Error() string {
	return fmt.Sprintf("Неверный пароль для пользователя: %s", W.Username)
}

func additionally() {
	fmt.Println(login("user2", "password3"))
}