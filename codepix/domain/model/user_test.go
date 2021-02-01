package model

import (
	"testing"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

func TestNewUser(t *testing.T) {
	name := "User001"
	email := "user001@email.com"
	bank, err := model. .NewBank(code, name)

	accountNumber := "abcnumber"
	ownerName := "Wesley"
	account, err := model.NewAccount(bank, accountNumber, ownerName)
}
