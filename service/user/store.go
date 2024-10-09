package user

import (
	"database/sql"
	"fmt"

	"github.com/sysronnie/go-user-auth/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) GetUserByUsername(username string) (*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM user_auth WHERE USERNAME = ?", username)
	if err != nil {
		return nil, err
	}

	User := new(types.User)

	for rows.Next() {
		User, err = scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	if User.USER_ID == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return User, nil
}

func scanRowIntoUser(rows *sql.Rows) (*types.User, error) {
	// Setting an instance of the User type and assinging it to user
	User := new(types.User)

	err := rows.Scan(
		&User.USER_ID,
		&User.EMAIL,
		&User.USERNAME,
		&User.PASSWORD,
		&User.CREATED_AT,
	)

	if err != nil {
		return nil, err
	}

	return User, nil
}

func (s *Store) GetUserByID(id int) (*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM user_auth where id = ?", id)
	if err != nil {
		return nil, err
	}

	User := new(types.User)

	for rows.Next() {
		User, err = scanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	if User.USER_ID == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return User, nil
}

func (s *Store) CreateUser(User types.User) error {
	_, err := s.db.Exec("INSERT into user_auth (USER_ID, EMAIL, USERNAME, PASSWORD) VALUES (?, ?, ?, ?)", User.USER_ID, User.EMAIL, User.USERNAME, User.PASSWORD)
	if err != nil {
		return err
	}

	return nil
}