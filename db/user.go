package db

import (
	"blockexchange/types"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type UserRepository interface {
	GetUserById(id int64) (*types.User, error)
	GetUserByName(name string) (*types.User, error)
	GetUserByExternalId(external_id string) (*types.User, error)
	GetUsers() ([]types.User, error)
	CreateUser(user *types.User) error
	UpdateUser(user *types.User) error
}

type DBUserRepository struct {
	DB *sqlx.DB
}

func (repo DBUserRepository) GetUserById(id int64) (*types.User, error) {
	user := []types.User{}
	err := repo.DB.Select(&user, "select * from public.user where id = $1", id)
	if err != nil {
		return nil, err
	} else if len(user) == 1 {
		return &user[0], nil
	} else {
		return nil, nil
	}
}

func (repo DBUserRepository) GetUserByName(name string) (*types.User, error) {
	user := []types.User{}
	err := repo.DB.Select(&user, "select * from public.user where name = $1", name)
	if err != nil {
		return nil, err
	} else if len(user) == 1 {
		return &user[0], nil
	} else {
		return nil, nil
	}
}

func (repo DBUserRepository) GetUserByExternalId(external_id string) (*types.User, error) {
	user := []types.User{}
	err := repo.DB.Select(&user, "select * from public.user where external_id = $1", external_id)
	if err != nil {
		return nil, err
	} else if len(user) == 1 {
		return &user[0], nil
	} else {
		return nil, nil
	}
}

func (repo DBUserRepository) GetUsers() ([]types.User, error) {
	logrus.Trace("db.GetUsers")
	list := []types.User{}
	err := repo.DB.Select(&list, "select * from public.user")
	if err != nil {
		return nil, err
	} else {
		return list, nil
	}
}

func (repo DBUserRepository) CreateUser(user *types.User) error {
	logrus.Trace("db.CreateUser", user)
	query := `
		insert into
		public.user(
			created, name, hash, type,
			external_id, mail
		)
		values(
			$1, $2, $3, $4,
			$5, $6
		)
		returning id
	`
	row := repo.DB.QueryRow(query,
		user.Created, user.Name, user.Hash, user.Type,
		user.ExternalID, user.Mail,
	)
	return row.Scan(&user.ID)
}

func (repo DBUserRepository) UpdateUser(user *types.User) error {
	logrus.Trace("db.UpdateUser", user)
	query := `
		update public.user
		set name = $2, mail = $3
		where id = $1
	`
	_, err := repo.DB.Exec(query, user.ID, user.Name, user.Mail)
	return err
}
