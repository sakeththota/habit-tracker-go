package progress

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sakeththota/habit-tracker-go/types"
)

type Store struct {
	db *pgxpool.Pool
}

func NewStore(db *pgxpool.Pool) *Store {
	return &Store{db: db}
}

func (s *Store) GetProgress(habit_id int) ([]types.ProgressEntry, error) {
	rows, err := s.db.Query(context.Background(), "SELECT * FROM progress WHERE habit_id = $1", habit_id)
	if err != nil {
		return nil, err
	}

	progress := make([]types.ProgressEntry, 0)
	for rows.Next() {
		p := new(types.ProgressEntry)
		err := rows.Scan(&p.ID, &p.HabitID, &p.Date, &p.CreatedAt)
		if err != nil {
			return nil, err
		}
		progress = append(progress, *p)
	}

	return progress, nil
}

func (s *Store) CreateCompletion(habit_id int, date string) error {
	createdAt := time.Now().UTC().Format("2006-01-02")
	_, err := s.db.Exec(context.Background(), "INSERT INTO progress (habit_id, date, created_at) VALUES ($1, $2, $3)", habit_id, date, createdAt)
	if err != nil {
		return err
	}

	return nil
}
