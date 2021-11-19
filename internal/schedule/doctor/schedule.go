package doctor

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"
)

// FindSchedule поиск расписания
func (s *Store) FindSchedule(ctx context.Context) ([]*Schedule, error) {
	var q = "select * from it_rasp"
	schedule := make([]*Schedule, 0)
	if err := s.db.SelectContext(ctx, &schedule, q); err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrNoSchedule
		}

		return nil, errors.Wrap(err, "error run sql query")
	}

	return schedule, nil
}
