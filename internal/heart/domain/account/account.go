package account

import "net/mail"

type Account struct {
	Email    string
	Password string
}

func IsValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
