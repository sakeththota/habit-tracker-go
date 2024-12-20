package user

import (
	"context"

	"github.com/sakeththota/habit-tracker-go/types"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Store struct {
	db *pgxpool.Pool
}

func NewStore(db *pgxpool.Pool) *Store {
	return &Store{db: db}
}

func (s *Store) GetUserByEmail(email string) (*types.User, error) {
	user := new(types.User)
	err := s.db.QueryRow(context.Background(), "SELECT * FROM users WHERE email = $1", email).Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.CreatedAt)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Store) GetUserById(id int) (*types.User, error) {
	user := new(types.User)
	err := s.db.QueryRow(context.Background(), "SELECT * FROM users WHERE user_id = $1", id).Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.CreatedAt)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Store) CreateUser(user types.User) error {
	_, err := s.db.Exec(context.Background(), "INSERT INTO users (username, email, password_hash) VALUES ($1, $2, $3)", user.Username, user.Email, user.PasswordHash)
	if err != nil {
		return err
	}

	return nil
}
