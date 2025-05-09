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

func (s *Store) GetHabitById(habit_id int) (*types.Habit, error) {
	habit := new(types.Habit)
	err := s.db.QueryRow(context.Background(), "SELECT * FROM habits WHERE habit_id = $1", habit_id).Scan(&habit.ID, &habit.UserID, &habit.Title, &habit.Description, &habit.CreatedAt)
	if err != nil {
		return nil, err
	}

	return habit, nil
}

func (s *Store) CreateHabit(habit types.Habit) error {
	_, err := s.db.Exec(context.Background(), "INSERT INTO habits (user_id, title, description) VALUES ($1, $2, $3)", habit.UserID, habit.Title, habit.Description)
	if err != nil {
		return err
	}

	return nil
}

func (s *Store) DeleteHabit(user_id int, habit_id int) error {
	_, err := s.db.Exec(context.Background(), "DELETE FROM habits WHERE user_id = $1 AND habit_id = $2", user_id, habit_id)
	if err != nil {
		return err
	}

	return nil
}
