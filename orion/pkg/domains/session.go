package domains

import (
	"database/sql"
	"errors"
	"regexp"
	"time"
)

var TABLE_SESSIONS = "sessions"

type Session struct {
	Id        uint
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt time.Time    `db:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at"`
	ClassId   string       `db:"class_id" json:"classId"`
	StartsAt  time.Time    `db:"starts_at" json:"startsAt"`
	EndsAt    time.Time    `db:"ends_at" json:"endsAt"`
	Canceled  bool         `db:"canceled" json:"canceled"`
	Notes     NullString   `db:"notes" json:"notes"`
}

func (session *Session) Validate() error {
	// Retrieves the inputted values
	notes := session.Notes

	// Notes validation
	if notes.Valid {
		if matches, _ := regexp.MatchString(REGEX_LETTER, notes.String); !matches {
			return errors.New("invalid notes")
		}
	}

	return nil
}