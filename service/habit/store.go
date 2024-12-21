package habit

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sakeththota/habit-tracker-go/types"
)

type Store struct {
	db *pgxpool.Pool
}

func NewStore(db *pgxpool.Pool) *Store {
	return &Store{db: db}
}

func (s *Store) GetHabits(user_id int) ([]types.Habit, error) {
	rows, err := s.db.Query(context.Background(), "SELECT * FROM habits WHERE user_id = $1", user_id)
	if err != nil {
		return nil, err
	}

	habits := make([]types.Habit, 0)
	for rows.Next() {
		h := new(types.Habit)
		err := rows.Scan(&h.ID, &h.UserID, &h.Title, &h.Description, &h.CreatedAt)
		if err != nil {
			return nil, err
		}
		habits = append(habits, *h)
	}

	return habits, nil
}

func (s *Store) CreateHabit(habit types.Habit) error {
	_, err := s.db.Exec(context.Background(), "INSERT INTO habits (user_id, title, description) VALUES ($1, $2, $3)", habit.UserID, habit.Title, habit.Description)
	if err != nil {
		return err
	}

	return nil
}
