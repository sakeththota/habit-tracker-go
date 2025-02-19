package progress

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

	// check if current date is latest progress entry
	// if not insert new entry and return full set

	return progress, nil
}

func (s *Store) MarkComplete(habit_id int) error {
	return nil
}
