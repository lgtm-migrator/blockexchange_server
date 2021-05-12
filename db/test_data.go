package db

import (
	"blockexchange/types"
	"time"

	"github.com/jmoiron/sqlx"
)

func PopulateTestData(_db *sqlx.DB) error {
	userrepo := DBUserRepository{DB: _db}
	tokenrepo := DBAccessTokenRepository{DB: _db}

	user, err := userrepo.GetUserByName("Testuser")
	if err != nil {
		return err
	}
	if user != nil {
		return nil
	}

	user = &types.User{
		Name: "Testuser",
		Type: types.UserTypeLocal,
		Hash: "",
	}
	err = userrepo.CreateUser(user)
	if err != nil {
		return err
	}

	token := &types.AccessToken{
		Name:     "Default",
		Token:    "default",
		UserID:   user.ID,
		Created:  time.Now().Unix(),
		Expires:  time.Now().Unix() + time.Hour.Microseconds(),
		UseCount: 0,
	}

	err = tokenrepo.CreateAccessToken(token)
	return err
}
