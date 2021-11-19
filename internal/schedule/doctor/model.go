package doctor

import (
	"github.com/pkg/errors"
	"gopkg.in/guregu/null.v4"
)

// ErrNoSchedule представляет ошибку отсутсвия расписания
var ErrNoSchedule = errors.New("no schedule")

// Schedule представляет расписание
type Schedule struct {
	Notd       null.String `db:"notd"`
	Nmpp       null.String `db:"nmpp"`
	Rname      null.String `db:"rname"`
	Even_Day   null.String `db:"even_day"`
	Noeven_Day null.String `db:"noeven_day"`
	Saturnday  null.String `db:"saturday"`
	Sunday     null.String `db:"dunday"`
}
