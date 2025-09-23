package userService

import (
	"JHETBackend/common/exception"
	"JHETBackend/models"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// Verify Password
func VerifyPwd(user *models.AccountInfo, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		return exception.UsrPasswordErr
	}
	return err
}
