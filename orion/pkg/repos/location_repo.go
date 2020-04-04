package repos

import (
	"database/sql"
	"github.com/ahsu1230/mathnavigatorSite/orion/pkg/domains"
	"time"
)

// Global variable
var LocationRepo LocationRepoInterface = &locationRepo{}

type locationRepo struct {
	db *sql.DB
}

type LocationRepoInterface interface {
	Initialize(db *sql.DB)
	SelectAll() ([]domains.Location, error)
	SelectByLocationId(string) (domains.Location, error)
	Insert(domains.Location) error
	Update(string, domains.Location) error
	Delete(string) error
	SelectAllUnpublished() ([]string, error)
	Publish([]string) error
}

func (lr *locationRepo) Initialize(db *sql.DB) {
	lr.db = db
}

func (lr *locationRepo) SelectAll() ([]domains.Location, error) {
	results := make([]domains.Location, 0)

	stmt, err := lr.db.Prepare("SELECT * FROM locations")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var location domains.Location
		if errScan := rows.Scan(
			&location.Id,
			&location.CreatedAt,
			&location.UpdatedAt,
			&location.DeletedAt,
			&location.PublishedAt,
			&location.LocId,
			&location.Street,
			&location.City,
			&location.State,
			&location.Zipcode,
			&location.Room); errScan != nil {
			return results, errScan
		}
		results = append(results, location)
	}

	return results, nil
}

func (lr *locationRepo) SelectByLocationId(locId string) (domains.Location, error) {
	stmt, err := lr.db.Prepare("SELECT * FROM locations WHERE loc_id=?")
	if err != nil {
		return domains.Location{}, err
	}
	defer stmt.Close()

	var location domains.Location
	row := stmt.QueryRow(locId)
	errScan := row.Scan(
		&location.Id,
		&location.CreatedAt,
		&location.UpdatedAt,
		&location.DeletedAt,
		&location.PublishedAt,
		&location.LocId,
		&location.Street,
		&location.City,
		&location.State,
		&location.Zipcode,
		&location.Room)

	return location, errScan
}

func (lr *locationRepo) Insert(location domains.Location) error {
	stmt, err := lr.db.Prepare("INSERT INTO locations (" +
		"created_at, " +
		"updated_at, " +
		"loc_id, " +
		"street, " +
		"city, " +
		"state, " +
		"zipcode, " +
		"room" +
		") VALUES (?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	now := time.Now().UTC()
	result, err := stmt.Exec(
		now,
		now,
		location.LocId,
		location.Street,
		location.City,
		location.State,
		location.Zipcode,
		location.Room)
	if err != nil {
		return err
	}

	return handleSqlExecResult(result, 1, "location was not inserted")
}

func (lr *locationRepo) Update(locId string, location domains.Location) error {
	stmt, err := lr.db.Prepare("UPDATE locations SET " +
		"updated_at=?, " +
		"loc_id=?, " +
		"street=?, " +
		"city=?, " +
		"state=?, " +
		"zipcode=?, " +
		"room=? " +
		"WHERE loc_id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	now := time.Now().UTC()
	result, err := stmt.Exec(
		now,
		location.LocId,
		location.Street,
		location.City,
		location.State,
		location.Zipcode,
		location.Room,
		locId)
	if err != nil {
		return err
	}

	return handleSqlExecResult(result, 1, "location was not updated")
}

func (lr *locationRepo) Delete(locId string) error {
	stmt, err := lr.db.Prepare("DELETE FROM locations WHERE loc_id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(locId)
	if err != nil {
		return err
	}

	return handleSqlExecResult(result, 1, "location was not deleted")
}

func (lr *locationRepo) SelectAllUnpublished() ([]string, error) {
	results := make([]string, 0)

	stmt, err := lr.db.Prepare("SELECT loc_id FROM locations WHERE published_at IS NULL")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var locId string
		if errScan := rows.Scan(&locId); errScan != nil {
			return results, errScan
		}
		results = append(results, locId)
	}
	return results, nil
}

func (lr *locationRepo) Publish(locIds []string) error {
	for _, locId := range locIds {
		location, err := lr.SelectByLocationId(locId)
		if err != nil {
			return err
		}
		if !location.PublishedAt.Valid {
			now := time.Now().UTC()
			location.PublishedAt.Scan(now)
			lr.Update(locId, location)
		}
	}
	return nil
}

func CreateTestLocationRepo(db *sql.DB) LocationRepoInterface {
	lr := &locationRepo{}
	lr.Initialize(db)
	return lr
}
