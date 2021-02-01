package model

import (
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewUser(t *testing.T) {
	name := "User001"
	email := "user001@email.com"
	user, err := NewUser(name, email)

	require.Nil(t, err)
	require.NotNil(t, uuid.FromStringOrNil(user.ID))
	require.Equal(t, user.Name, name)
	require.Equal(t, user.Email, email)
}
