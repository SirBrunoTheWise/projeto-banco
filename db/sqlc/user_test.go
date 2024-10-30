package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
	arg := CreateUserParams{
		Alias:       "+a",
		Email:       "maisasilva@gmail.com",
		Password:    "hashing",
		FirstName:   "Maisa",
		LastName:    "Silva",
		Cpf:         "09263012945",
		PhoneNumber: "+55940028922",
	}

	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Alias, user.Alias)
	require.Equal(t, arg.Cpf, user.Cpf)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.FirstName, user.FirstName)
	require.Equal(t, arg.LastName, user.LastName)
	require.Equal(t, arg.Password, user.Password)
	require.Equal(t, arg.PhoneNumber, user.PhoneNumber)

	require.NotZero(t, user.UserID)
	require.NotZero(t, user.CreateTime)
}

func TestGetUser(t *testing.T) {
	arg := CreateUserParams{
		Alias:       "+a",
		Email:       "maisasilva@gmail.com",
		Password:    "hashing",
		FirstName:   "Maisa",
		LastName:    "Silva",
		Cpf:         "09263012945",
		PhoneNumber: "+55940028922",
	}

	user, err := testQueries.GetUser(context.Background(), 1)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.Alias, user.Alias)
	require.Equal(t, arg.Cpf, user.Cpf)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.FirstName, user.FirstName)
	require.Equal(t, arg.LastName, user.LastName)
	require.Equal(t, arg.Password, user.Password)
	require.Equal(t, arg.PhoneNumber, user.PhoneNumber)

	require.NotZero(t, user.UserID)
	require.NotZero(t, user.CreateTime)
}
